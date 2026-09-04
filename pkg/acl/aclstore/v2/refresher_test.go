package aclstore

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakeReloadable counts Reload calls and can be made to fail.
type fakeReloadable struct {
	mu    sync.Mutex
	count int
	err   error
}

func (f *fakeReloadable) Reload(context.Context) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.count++
	return f.err
}

func (f *fakeReloadable) calls() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.count
}

func (f *fakeReloadable) setErr(err error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.err = err
}

func TestRefresher_StartDoesInitialLoad(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	store := &fakeReloadable{}
	r := NewRefresher(store, WithReloadInterval(time.Hour), WithDebounce(time.Hour))

	require.NoError(t, r.Start(ctx))
	assert.Equal(t, 1, store.calls())
	assert.False(t, r.LastReload().IsZero())
}

func TestRefresher_StartInitialLoadError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	store := &fakeReloadable{err: fmt.Errorf("db down")}
	r := NewRefresher(store)

	err := r.Start(ctx)
	require.Error(t, err)
	assert.Error(t, r.LastError())
}

// A burst of Notify calls must collapse into a single extra reload.
func TestRefresher_NotifyCoalesces(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	store := &fakeReloadable{}
	// Long interval so the ticker never fires during the test; short debounce.
	r := NewRefresher(store, WithReloadInterval(time.Hour), WithDebounce(30*time.Millisecond))
	require.NoError(t, r.Start(ctx)) // count == 1

	for i := 0; i < 5; i++ {
		r.Notify()
	}

	assert.Eventually(t, func() bool { return store.calls() == 2 }, time.Second, 5*time.Millisecond,
		"burst of notifies should trigger exactly one extra reload")

	// No further reloads should happen without new signals.
	time.Sleep(120 * time.Millisecond)
	assert.Equal(t, 2, store.calls())
}

// The ticker must reload unconditionally even without any Notify.
func TestRefresher_TickerReloads(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	store := &fakeReloadable{}
	r := NewRefresher(store, WithReloadInterval(20*time.Millisecond), WithDebounce(time.Hour))
	require.NoError(t, r.Start(ctx)) // count == 1

	assert.Eventually(t, func() bool { return store.calls() >= 3 }, time.Second, 5*time.Millisecond,
		"ticker should keep reloading without any notifies")
}

// A failed reload records the error but does not stop future reloads.
func TestRefresher_RecoversAfterReloadError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	store := &fakeReloadable{}
	r := NewRefresher(store, WithReloadInterval(20*time.Millisecond), WithDebounce(time.Hour))
	require.NoError(t, r.Start(ctx))

	store.setErr(fmt.Errorf("transient"))
	assert.Eventually(t, func() bool { return r.LastError() != nil }, time.Second, 5*time.Millisecond)

	store.setErr(nil)
	assert.Eventually(t, func() bool { return r.LastError() == nil }, time.Second, 5*time.Millisecond,
		"refresher should recover once reloads succeed again")
}
