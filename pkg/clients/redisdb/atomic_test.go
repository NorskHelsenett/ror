package redisdb

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func TestAtomicStorage(t *testing.T) {
	server := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: server.Addr()})
	t.Cleanup(func() { _ = client.Close() })
	db := NewFromClient(client)
	ctx := context.Background()
	if _, err := db.GetBounded(ctx, "state", 10); !errors.Is(err, redis.Nil) {
		t.Fatalf("missing key: %v", err)
	}
	if _, err := db.GetBounded(ctx, "state", 0); err == nil {
		t.Fatal("accepted zero bound")
	}
	if ok, err := db.CompareAndSwap(ctx, "state", nil, "first"); err != nil || !ok {
		t.Fatalf("initial CAS: %v, %v", ok, err)
	}
	if _, err := db.GetBounded(ctx, "state", 4); err == nil {
		t.Fatal("oversized value was returned")
	}
	if got, err := db.GetBounded(ctx, "state", 5); err != nil || got != "first" {
		t.Fatalf("bounded read: %q, %v", got, err)
	}
	wrong := "wrong"
	if ok, err := db.CompareAndSwap(ctx, "state", &wrong, "bad"); err != nil || ok {
		t.Fatalf("stale CAS: %v, %v", ok, err)
	}
	var winners atomic.Int32
	var wg sync.WaitGroup
	for range 16 {
		wg.Go(func() {
			old := "first"
			ok, err := db.CompareAndSwap(ctx, "state", &old, "next")
			if err != nil {
				t.Error(err)
			}
			if ok {
				winners.Add(1)
			}
		})
	}
	wg.Wait()
	if winners.Load() != 1 {
		t.Fatalf("CAS winners = %d", winners.Load())
	}
	if ttl := server.TTL("state"); ttl != 0 {
		t.Fatalf("unexpected expiry: %v", ttl)
	}
}
