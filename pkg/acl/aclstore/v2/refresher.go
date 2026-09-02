package aclstore

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

const (
	defaultReloadInterval = 30 * time.Second
	defaultDebounce       = 250 * time.Millisecond
)

// Reloadable is the behavior the Refresher drives: rebuilding an in-memory view
// from the source of truth. *SnapshotStore satisfies it.
type Reloadable interface {
	Reload(ctx context.Context) error
}

// Refresher keeps a Reloadable converged with the source of truth. It coalesces
// change signals (debounced) into a single reload and reloads unconditionally on
// a ticker, so a missed signal cannot leave the view stale beyond the interval.
type Refresher struct {
	store    Reloadable
	interval time.Duration
	debounce time.Duration
	notify   chan struct{}

	mu         sync.Mutex
	lastReload time.Time
	lastErr    error
}

// RefresherOption configures a Refresher.
type RefresherOption func(*Refresher)

// WithReloadInterval sets the unconditional periodic reload interval.
func WithReloadInterval(d time.Duration) RefresherOption {
	return func(r *Refresher) {
		if d > 0 {
			r.interval = d
		}
	}
}

// WithDebounce sets the window used to coalesce a burst of change signals.
func WithDebounce(d time.Duration) RefresherOption {
	return func(r *Refresher) {
		if d > 0 {
			r.debounce = d
		}
	}
}

// NewRefresher creates a Refresher for the given store.
func NewRefresher(store Reloadable, opts ...RefresherOption) *Refresher {
	r := &Refresher{
		store:    store,
		interval: defaultReloadInterval,
		debounce: defaultDebounce,
		notify:   make(chan struct{}, 1),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// Notify requests a coalesced reload. It is non-blocking and safe to call from
// any goroutine; concurrent calls collapse into a single pending reload.
func (r *Refresher) Notify() {
	select {
	case r.notify <- struct{}{}:
	default: // a reload is already pending
	}
}

// Start performs an initial synchronous load, then runs the ticker and debounce
// loops until ctx is cancelled. It returns an error only if the initial load
// fails (so callers can gate readiness on it).
func (r *Refresher) Start(ctx context.Context) error {
	if err := r.reload(ctx); err != nil {
		return fmt.Errorf("initial ACL snapshot load failed: %w", err)
	}
	go r.loop(ctx)
	return nil
}

// loop runs until ctx is cancelled, reloading on the ticker and after a
// debounced burst of Notify calls.
func (r *Refresher) loop(ctx context.Context) {
	ticker := time.NewTicker(r.interval)
	defer ticker.Stop()

	var debounceC <-chan time.Time
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_ = r.reload(ctx)
		case <-r.notify:
			// Start the debounce window on the first signal of a burst;
			// further signals while it runs are already accounted for.
			if debounceC == nil {
				debounceC = time.After(r.debounce)
			}
		case <-debounceC:
			debounceC = nil
			_ = r.reload(ctx)
		}
	}
}

// reload runs a single reload and records the outcome for observability.
func (r *Refresher) reload(ctx context.Context) error {
	err := r.store.Reload(ctx)

	r.mu.Lock()
	if err != nil {
		r.lastErr = err
	} else {
		r.lastReload = time.Now()
		r.lastErr = nil
	}
	r.mu.Unlock()

	if err != nil {
		rlog.Warnc(ctx, "ACL snapshot reload failed", rlog.Any("error", err))
	}
	return err
}

// LastReload returns the time of the last successful reload (zero if none).
func (r *Refresher) LastReload() time.Time {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.lastReload
}

// LastError returns the error from the most recent reload, or nil if it succeeded.
func (r *Refresher) LastError() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.lastErr
}
