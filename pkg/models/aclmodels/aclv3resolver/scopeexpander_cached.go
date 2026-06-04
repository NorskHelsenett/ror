package aclv3resolver

import (
	"context"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

// CachedScopeExpander wraps a ScopeExpander with an in-memory TTL cache.
// The cache is keyed by scope+subject and stores the expanded descendant list.
// It is safe for concurrent use.
type CachedScopeExpander struct {
	backend ScopeExpander
	ttl     time.Duration

	mu    sync.RWMutex
	cache map[AclV3Ownerref]cachedExpansion
}

type cachedExpansion struct {
	refs      []AclV3Ownerref
	expiresAt time.Time
}

// NewCachedScopeExpander creates a cached wrapper around the given expander.
func NewCachedScopeExpander(backend ScopeExpander, ttl time.Duration) *CachedScopeExpander {
	return &CachedScopeExpander{
		backend: backend,
		ttl:     ttl,
		cache:   make(map[AclV3Ownerref]cachedExpansion),
	}
}

// ExpandScope returns cached descendants or falls through to the backend.
func (c *CachedScopeExpander) ExpandScope(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject) ([]AclV3Ownerref, error) {
	key := AclV3Ownerref{Scope: scope, Subject: subject}

	c.mu.RLock()
	if entry, ok := c.cache[key]; ok && time.Now().Before(entry.expiresAt) {
		c.mu.RUnlock()
		return entry.refs, nil
	}
	c.mu.RUnlock()

	refs, err := c.backend.ExpandScope(ctx, scope, subject)
	if err != nil {
		return nil, err
	}

	c.mu.Lock()
	c.cache[key] = cachedExpansion{
		refs:      refs,
		expiresAt: time.Now().Add(c.ttl),
	}
	c.mu.Unlock()

	return refs, nil
}

// Invalidate removes the cached expansion for a specific scope+subject.
func (c *CachedScopeExpander) Invalidate(scope aclscope.Scope, subject aclscope.Subject) {
	key := AclV3Ownerref{Scope: scope, Subject: subject}
	c.mu.Lock()
	delete(c.cache, key)
	c.mu.Unlock()
}

// InvalidateAll clears the entire cache.
func (c *CachedScopeExpander) InvalidateAll() {
	c.mu.Lock()
	c.cache = make(map[AclV3Ownerref]cachedExpansion)
	c.mu.Unlock()
}
