package acl

import (
	"context"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/NorskHelsenett/ror/pkg/telemetry/rortracer"

	"go.opentelemetry.io/otel/attribute"
)

// CachedScopeExpander wraps a ScopeExpander with an in-memory TTL cache.
// The cache is keyed by scope+subject and stores the expanded descendant list.
// It is safe for concurrent use.
type CachedScopeExpander struct {
	backend ScopeExpander
	ttl     time.Duration

	mu    sync.RWMutex
	cache map[Ownerref]cachedExpansion
}

type cachedExpansion struct {
	refs      []Ownerref
	expiresAt time.Time
}

// NewCachedScopeExpander creates a cached wrapper around the given expander.
func NewCachedScopeExpander(backend ScopeExpander, ttl time.Duration) *CachedScopeExpander {
	return &CachedScopeExpander{
		backend: backend,
		ttl:     ttl,
		cache:   make(map[Ownerref]cachedExpansion),
	}
}

// ExpandScope returns cached descendants or falls through to the backend.
func (c *CachedScopeExpander) ExpandScope(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject) ([]Ownerref, error) {
	ctx, span := rortracer.StartSpan(ctx, "acl.CachedScopeExpander.ExpandScope")
	defer span.End()
	span.SetAttributes(
		attribute.String("acl.scope", string(scope)),
		attribute.String("acl.subject", string(subject)),
	)

	key := Ownerref{Scope: scope, Subject: subject}

	c.mu.RLock()
	if entry, ok := c.cache[key]; ok && time.Now().Before(entry.expiresAt) {
		c.mu.RUnlock()
		span.SetAttributes(attribute.Bool("acl.cache_hit", true))
		return entry.refs, nil
	}
	c.mu.RUnlock()
	span.SetAttributes(attribute.Bool("acl.cache_hit", false))

	refs, err := c.backend.ExpandScope(ctx, scope, subject)
	if err != nil {
		return nil, rortracer.SpanError(span, err)
	}

	c.mu.Lock()
	c.cache[key] = cachedExpansion{
		refs:      refs,
		expiresAt: time.Now().Add(c.ttl),
	}
	c.mu.Unlock()

	return refs, nil
}

// ExpandScopes expands several seeds, serving each from the in-memory cache when
// possible and resolving all cache misses through the backend in a single batched
// call. Results are keyed by seed ownerref.
func (c *CachedScopeExpander) ExpandScopes(ctx context.Context, seeds []Ownerref) (map[Ownerref][]Ownerref, error) {
	ctx, span := rortracer.StartSpan(ctx, "acl.CachedScopeExpander.ExpandScopes")
	defer span.End()
	span.SetAttributes(attribute.Int("acl.seeds", len(seeds)))

	result := make(map[Ownerref][]Ownerref, len(seeds))
	var misses []Ownerref

	c.mu.RLock()
	for _, seed := range seeds {
		if _, done := result[seed]; done {
			continue
		}
		if entry, ok := c.cache[seed]; ok && time.Now().Before(entry.expiresAt) {
			result[seed] = entry.refs
		} else {
			result[seed] = nil
			misses = append(misses, seed)
		}
	}
	c.mu.RUnlock()

	uniqueSeeds := len(result)
	span.SetAttributes(
		attribute.Int("acl.cache_hits", uniqueSeeds-len(misses)),
		attribute.Int("acl.cache_misses", len(misses)),
	)

	if len(misses) == 0 {
		return result, nil
	}

	expanded, err := c.backend.ExpandScopes(ctx, misses)
	if err != nil {
		return nil, rortracer.SpanError(span, err)
	}

	now := time.Now()
	c.mu.Lock()
	for _, seed := range misses {
		refs := expanded[seed]
		c.cache[seed] = cachedExpansion{refs: refs, expiresAt: now.Add(c.ttl)}
		result[seed] = refs
	}
	c.mu.Unlock()

	return result, nil
}

// Invalidate removes the cached expansion for a specific scope+subject.
func (c *CachedScopeExpander) Invalidate(scope aclscope.Scope, subject aclscope.Subject) {
	key := Ownerref{Scope: scope, Subject: subject}
	c.mu.Lock()
	delete(c.cache, key)
	c.mu.Unlock()
}

// InvalidateAll clears the entire cache.
func (c *CachedScopeExpander) InvalidateAll() {
	c.mu.Lock()
	c.cache = make(map[Ownerref]cachedExpansion)
	c.mu.Unlock()
}
