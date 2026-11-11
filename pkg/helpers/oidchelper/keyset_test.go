package oidchelper

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/lestrrat-go/jwx/v4/jwk"
)

func testPublicKey(t *testing.T, private *rsa.PrivateKey, kid string) jwk.Key {
	t.Helper()
	key, err := jwk.Import[jwk.Key](&private.PublicKey)
	if err != nil {
		t.Fatal(err)
	}
	for name, value := range map[string]any{jwk.KeyIDKey: kid, jwk.AlgorithmKey: "RS512", jwk.KeyUsageKey: "sig"} {
		if err := key.Set(name, value); err != nil {
			t.Fatal(err)
		}
	}
	return key
}

func testKeySet(t *testing.T, keys ...jwk.Key) jwk.Set {
	t.Helper()
	set := jwk.NewSet()
	for _, key := range keys {
		if err := set.AddKey(key); err != nil {
			t.Fatal(err)
		}
	}
	return set
}

func newTestRSAKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	return key
}

type fakeClock struct{ now time.Time }

func (c *fakeClock) Now() time.Time { return c.now }

func TestKeySetSeed(t *testing.T) {
	private := newTestRSAKey(t)
	first := testKeySet(t, testPublicKey(t, private, "a"))
	next := testKeySet(t, testPublicKey(t, private, "b"))
	var calls int
	cache := NewKeySet(func(context.Context) (jwk.Set, error) {
		calls++
		return next, nil
	}, time.Minute)
	if err := cache.Seed(jwk.NewSet()); err == nil {
		t.Fatal("invalid seed accepted")
	}
	if err := cache.Seed(first); err != nil {
		t.Fatal(err)
	}
	if _, err := cache.KeyFor(context.Background(), "a"); err != nil || calls != 0 {
		t.Fatalf("seeded lookup fetched: %v, %d", err, calls)
	}
	if _, err := cache.KeyFor(context.Background(), "b"); err != nil || calls != 1 {
		t.Fatalf("unknown key did not refresh: %v, %d", err, calls)
	}
	if err := cache.Seed(first); err == nil {
		t.Fatal("seed replaced fetched keys")
	}
	if _, ok := cache.cachedWith("b"); !ok {
		t.Fatal("fetched keys lost")
	}
}

func TestKeySetSeedWaitsForFetch(t *testing.T) {
	private := newTestRSAKey(t)
	seeded := testKeySet(t, testPublicKey(t, private, "seeded"))
	fetched := testKeySet(t, testPublicKey(t, private, "fetched"))

	for _, fail := range []bool{false, true} {
		t.Run(fmt.Sprintf("fetch-fails=%t", fail), func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				release := make(chan struct{})
				cache := NewKeySet(func(context.Context) (jwk.Set, error) {
					<-release
					if fail {
						return nil, errors.New("fetch failed")
					}
					return fetched, nil
				}, time.Minute)
				fetchDone := make(chan error, 1)
				go func() {
					_, err := cache.KeyFor(context.Background(), "fetched")
					fetchDone <- err
				}()
				synctest.Wait()

				seedDone := make(chan error, 1)
				go func() { seedDone <- cache.Seed(seeded) }()
				synctest.Wait()
				select {
				case err := <-seedDone:
					close(release)
					t.Fatalf("Seed completed during an in-flight fetch: %v", err)
				default:
				}

				close(release)
				if err := <-fetchDone; (err != nil) != fail {
					t.Fatalf("fetch error = %v, want failure %t", err, fail)
				}
				err := <-seedDone
				if fail {
					if err != nil {
						t.Fatalf("Seed after failed fetch: %v", err)
					}
					// The failed fetch's refresh throttle must not block seeded keys.
					if _, err := cache.KeyFor(context.Background(), "seeded"); err != nil {
						t.Fatalf("seeded key lookup: %v", err)
					}
				} else {
					if err == nil {
						t.Fatal("Seed replaced the completed fetch")
					}
					if _, ok := cache.cachedWith("fetched"); !ok {
						t.Fatal("fetched keys were lost")
					}
				}
			})
		})
	}
}

func TestKeySetRefreshesOnlyForUnknownKeyIDs(t *testing.T) {
	private := newTestRSAKey(t)
	first := testKeySet(t, testPublicKey(t, private, "a"))
	rotated := testKeySet(t, testPublicKey(t, private, "a"), testPublicKey(t, private, "b"))
	var calls atomic.Int32
	current := first
	clock := &fakeClock{now: time.Unix(1000, 0)}
	cache := NewKeySet(func(context.Context) (jwk.Set, error) {
		calls.Add(1)
		return current, nil
	}, time.Minute)
	cache.now = clock.Now
	ctx := context.Background()

	if _, err := cache.KeyFor(ctx, "a"); err != nil || calls.Load() != 1 {
		t.Fatalf("initial load: err %v calls %d", err, calls.Load())
	}
	if _, err := cache.KeyFor(ctx, "a"); err != nil || calls.Load() != 1 {
		t.Fatalf("known kid must not refetch: err %v calls %d", err, calls.Load())
	}
	current = rotated
	if _, err := cache.KeyFor(ctx, "b"); !errors.Is(err, ErrUnknownKeyID) || calls.Load() != 1 {
		t.Fatalf("refresh within interval must be suppressed: err %v calls %d", err, calls.Load())
	}
	clock.now = clock.now.Add(time.Minute)
	keys, err := cache.KeyFor(ctx, "b")
	if err != nil || calls.Load() != 2 || keys.Len() != 2 {
		t.Fatalf("rotation: err %v calls %d", err, calls.Load())
	}
	if _, err := cache.KeyFor(ctx, "a"); err != nil || calls.Load() != 2 {
		t.Fatalf("retained key: err %v calls %d", err, calls.Load())
	}
}

func TestKeySetFailurePreservesCache(t *testing.T) {
	private := newTestRSAKey(t)
	good := testKeySet(t, testPublicKey(t, private, "a"))
	clock := &fakeClock{now: time.Unix(1000, 0)}
	results := []func() (jwk.Set, error){
		func() (jwk.Set, error) { return good, nil },
		func() (jwk.Set, error) { return nil, errors.New("down") },
		func() (jwk.Set, error) { return jwk.NewSet(), nil },
	}
	cache := NewKeySet(func(context.Context) (jwk.Set, error) {
		next := results[0]
		results = results[1:]
		return next()
	}, time.Second)
	cache.now = clock.Now
	ctx := context.Background()
	if _, err := cache.KeyFor(ctx, "a"); err != nil {
		t.Fatal(err)
	}
	for range 2 {
		clock.now = clock.now.Add(time.Second)
		if _, err := cache.KeyFor(ctx, "missing"); err == nil {
			t.Fatal("expected refresh failure")
		}
		if _, err := cache.KeyFor(ctx, "a"); err != nil {
			t.Fatalf("cached key lost after failed refresh: %v", err)
		}
	}
}

func TestKeySetCoalescesConcurrentFetches(t *testing.T) {
	private := newTestRSAKey(t)
	keys := testKeySet(t, testPublicKey(t, private, "a"))
	var calls atomic.Int32
	release := make(chan struct{})
	cache := NewKeySet(func(context.Context) (jwk.Set, error) {
		calls.Add(1)
		<-release
		return keys, nil
	}, time.Minute)
	var wg sync.WaitGroup
	errs := make(chan error, 8)
	for range 8 {
		wg.Go(func() {
			_, err := cache.KeyFor(context.Background(), "a")
			errs <- err
		})
	}
	time.Sleep(20 * time.Millisecond)
	close(release)
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
	if calls.Load() != 1 {
		t.Fatalf("expected one fetch, got %d", calls.Load())
	}
}

func TestKeySetWaitHonorsContext(t *testing.T) {
	release := make(chan struct{})
	defer close(release)
	started := make(chan struct{})
	cache := NewKeySet(func(context.Context) (jwk.Set, error) {
		close(started)
		<-release
		return nil, errors.New("released")
	}, time.Minute)
	go func() { _, _ = cache.KeyFor(context.Background(), "a") }()
	<-started
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := cache.KeyFor(ctx, "a"); !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context cancellation, got %v", err)
	}
	if _, err := cache.KeyFor(context.Background(), ""); !errors.Is(err, ErrUnknownKeyID) {
		t.Fatalf("expected empty kid rejection, got %v", err)
	}
}

func TestValidatePublicKeySet(t *testing.T) {
	private := newTestRSAKey(t)
	publicKey := testPublicKey(t, private, "a")
	privateKey, err := jwk.Import[jwk.Key](private)
	if err != nil {
		t.Fatal(err)
	}
	_ = privateKey.Set(jwk.KeyIDKey, "p")
	_ = privateKey.Set(jwk.AlgorithmKey, "RS512")
	symmetric, err := jwk.Import[jwk.Key]([]byte("secret"))
	if err != nil {
		t.Fatal(err)
	}
	_ = symmetric.Set(jwk.KeyIDKey, "s")
	_ = symmetric.Set(jwk.AlgorithmKey, "HS512")
	noKid, _ := jwk.Import[jwk.Key](&private.PublicKey)
	_ = noKid.Set(jwk.AlgorithmKey, "RS512")
	noAlg, _ := jwk.Import[jwk.Key](&private.PublicKey)
	_ = noAlg.Set(jwk.KeyIDKey, "n")
	encryption := testPublicKey(t, private, "e")
	_ = encryption.Set(jwk.KeyUsageKey, "enc")
	duplicate := testPublicKey(t, newTestRSAKey(t), "a")

	if err := ValidatePublicKeySet(testKeySet(t, publicKey)); err != nil {
		t.Fatal(err)
	}
	for name, set := range map[string]jwk.Set{
		"nil":        nil,
		"empty":      jwk.NewSet(),
		"private":    testKeySet(t, privateKey),
		"symmetric":  testKeySet(t, symmetric),
		"no kid":     testKeySet(t, noKid),
		"no alg":     testKeySet(t, noAlg),
		"encryption": testKeySet(t, encryption),
		"duplicate":  testKeySet(t, publicKey, duplicate),
	} {
		t.Run(name, func(t *testing.T) {
			if err := ValidatePublicKeySet(set); err == nil {
				t.Fatal("expected invalid key set")
			}
		})
	}
}
