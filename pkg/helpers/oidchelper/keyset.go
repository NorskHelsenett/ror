package oidchelper

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/lestrrat-go/jwx/v4/jwa"
	"github.com/lestrrat-go/jwx/v4/jwk"
)

// DefaultKeySetRefreshInterval limits how often unknown key IDs trigger a fetch.
const DefaultKeySetRefreshInterval = 10 * time.Second

// ErrUnknownKeyID is returned when a key ID is absent after any allowed refresh.
var ErrUnknownKeyID = errors.New("unknown verification key id")

// KeySetFetcher fetches a public JSON Web Key Set from a trusted, fixed source.
type KeySetFetcher func(ctx context.Context) (jwk.Set, error)

// KeySet caches public verification keys. It fetches keys only when a requested
// key ID is unknown, at most once per refresh interval, and coalesces concurrent
// fetches. A failed fetch or invalid key set never replaces the cached keys.
// Signature failures for known key IDs must not call KeyFor again expecting a refresh.
type KeySet struct {
	fetch    KeySetFetcher
	interval time.Duration
	now      func() time.Time

	fetching chan struct{}
	mu       sync.RWMutex
	keys     jwk.Set
	lastTry  time.Time
}

// NewKeySet creates a key cache. A non-positive interval uses DefaultKeySetRefreshInterval.
func NewKeySet(fetch KeySetFetcher, interval time.Duration) *KeySet {
	if interval <= 0 {
		interval = DefaultKeySetRefreshInterval
	}
	return &KeySet{fetch: fetch, interval: interval, now: time.Now, fetching: make(chan struct{}, 1)}
}

// Seed installs independently trusted, validated keys only into an empty cache.
func (k *KeySet) Seed(keys jwk.Set) error {
	if err := ValidatePublicKeySet(keys); err != nil {
		return err
	}
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.keys != nil {
		return errors.New("verification key cache is already initialized")
	}
	k.keys = keys
	return nil
}

// KeyFor returns the cached public key set when it contains kid, refreshing the
// cache first when kid is unknown and the refresh interval allows it.
func (k *KeySet) KeyFor(ctx context.Context, kid string) (jwk.Set, error) {
	if kid == "" {
		return nil, fmt.Errorf("%w: empty key id", ErrUnknownKeyID)
	}
	if keys, ok := k.cachedWith(kid); ok {
		return keys, nil
	}
	select {
	case k.fetching <- struct{}{}:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	defer func() { <-k.fetching }()

	// Another caller may have refreshed the keys while this caller waited.
	if keys, ok := k.cachedWith(kid); ok {
		return keys, nil
	}
	k.mu.RLock()
	recent := !k.lastTry.IsZero() && k.now().Sub(k.lastTry) < k.interval
	k.mu.RUnlock()
	if recent {
		return nil, fmt.Errorf("%w %q", ErrUnknownKeyID, kid)
	}

	k.mu.Lock()
	k.lastTry = k.now()
	k.mu.Unlock()
	if k.fetch == nil {
		return nil, errors.New("key set fetcher is required")
	}
	keys, err := k.fetch(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch verification keys: %w", err)
	}
	if err := ValidatePublicKeySet(keys); err != nil {
		return nil, err
	}
	k.mu.Lock()
	k.keys = keys
	k.mu.Unlock()
	if _, ok := keys.LookupKeyID(kid); !ok {
		return nil, fmt.Errorf("%w %q", ErrUnknownKeyID, kid)
	}
	return keys, nil
}

func (k *KeySet) cachedWith(kid string) (jwk.Set, bool) {
	k.mu.RLock()
	defer k.mu.RUnlock()
	if k.keys == nil {
		return nil, false
	}
	_, ok := k.keys.LookupKeyID(kid)
	return k.keys, ok
}

// ValidatePublicKeySet requires a nonempty set of public RSA signature keys
// with unique key IDs and explicit algorithms.
func ValidatePublicKeySet(keys jwk.Set) error {
	if keys == nil || keys.Len() == 0 {
		return errors.New("verification key set is empty")
	}
	seen := map[string]bool{}
	for _, key := range keys.All() {
		kid, ok := key.KeyID()
		if !ok || kid == "" {
			return errors.New("verification key is missing kid")
		}
		if seen[kid] {
			return fmt.Errorf("duplicate verification key id %q", kid)
		}
		seen[kid] = true
		if key.KeyType() != jwa.RSA() {
			return fmt.Errorf("verification key %q must be RSA", kid)
		}
		if private, err := jwk.IsPrivateKey(key); err != nil || private {
			return fmt.Errorf("verification key %q must be public", kid)
		}
		if _, ok := key.Algorithm(); !ok {
			return fmt.Errorf("verification key %q is missing alg", kid)
		}
		if use, ok := key.KeyUsage(); ok && use != "sig" {
			return fmt.Errorf("verification key %q is not a signature key", kid)
		}
	}
	return nil
}
