package aclv3store

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients/redisdb"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

const cacheKeyPrefix = "acl:v3:group:"

// CachedAclV3Store wraps an AclV3Store backend with a Redis cache layer.
// Cache is per-group: each group's full ACL entry list is stored as a single JSON blob.
// On GetByGroups, it does MGET for all groups, backfills misses from the backend,
// and caches the backfilled results.
// If Redis is unavailable, it falls through to the backend transparently.
type CachedAclV3Store struct {
	backend aclmodels.AclV3Store
	redis   redisdb.RedisDB
	ttl     time.Duration
}

// NewCachedAclV3Store creates a cached store wrapping the given backend.
func NewCachedAclV3Store(backend aclmodels.AclV3Store, redis redisdb.RedisDB, ttl time.Duration) *CachedAclV3Store {
	return &CachedAclV3Store{
		backend: backend,
		redis:   redis,
		ttl:     ttl,
	}
}

// GetByGroups returns ACL entries for all groups, using Redis cache where possible.
// Cache misses are backfilled from the backend in a single call.
func (c *CachedAclV3Store) GetByGroups(ctx context.Context, groups []string) (map[string][]aclmodels.AclV3ListItem, error) {
	if len(groups) == 0 {
		return make(map[string][]aclmodels.AclV3ListItem), nil
	}

	result := make(map[string][]aclmodels.AclV3ListItem, len(groups))

	// Try cache first
	hits, misses, cacheErr := c.mget(ctx, groups)
	if cacheErr != nil {
		// Redis down — fall through to backend for all groups
		rlog.Warnc(ctx, "redis cache unavailable, falling through to backend", rlog.Any("error", cacheErr))
		return c.backend.GetByGroups(ctx, groups)
	}

	// Add hits to result
	for group, entries := range hits {
		result[group] = entries
	}

	// No misses — all from cache
	if len(misses) == 0 {
		return result, nil
	}

	// Backfill misses from backend
	backfilled, err := c.backend.GetByGroups(ctx, misses)
	if err != nil {
		return nil, err
	}

	// Cache backfilled results (including empty groups)
	c.setMany(ctx, misses, backfilled)

	// Merge backfilled into result
	for _, group := range misses {
		if entries, ok := backfilled[group]; ok {
			result[group] = entries
		}
		// Groups with no entries: leave absent from result (same as backend behavior)
	}

	return result, nil
}

// Invalidate removes the cached entries for a group.
// Call this when ACL entries for the group are created, updated, or deleted.
func (c *CachedAclV3Store) Invalidate(ctx context.Context, group string) error {
	key := cacheKeyPrefix + group
	return c.redis.Delete(ctx, key)
}

// mget fetches cached entries for all groups in a single MGET call.
// Returns hits (group → entries), misses (group names not in cache), and any error.
func (c *CachedAclV3Store) mget(ctx context.Context, groups []string) (map[string][]aclmodels.AclV3ListItem, []string, error) {
	keys := make([]string, len(groups))
	for i, g := range groups {
		keys[i] = cacheKeyPrefix + g
	}

	vals, err := c.redis.MGet(ctx, keys...)
	if err != nil {
		return nil, nil, err
	}

	hits := make(map[string][]aclmodels.AclV3ListItem)
	var misses []string

	for i, val := range vals {
		group := groups[i]
		if val == nil {
			misses = append(misses, group)
			continue
		}

		str, ok := val.(string)
		if !ok {
			misses = append(misses, group)
			continue
		}

		var entries []aclmodels.AclV3ListItem
		if err := json.Unmarshal([]byte(str), &entries); err != nil {
			rlog.Warnc(ctx, fmt.Sprintf("failed to unmarshal cached ACL for group %q, treating as miss", group), rlog.Any("error", err))
			misses = append(misses, group)
			continue
		}
		hits[group] = entries
	}

	return hits, misses, nil
}

// setMany caches entries for the given groups using a Redis pipeline.
// Groups with no entries in backfilled are cached as empty arrays.
func (c *CachedAclV3Store) setMany(ctx context.Context, groups []string, backfilled map[string][]aclmodels.AclV3ListItem) {
	var items []redisdb.SetItem

	for _, group := range groups {
		entries := backfilled[group]
		if entries == nil {
			entries = []aclmodels.AclV3ListItem{}
		}

		data, err := json.Marshal(entries)
		if err != nil {
			rlog.Warnc(ctx, fmt.Sprintf("failed to marshal ACL for group %q, skipping cache", group), rlog.Any("error", err))
			continue
		}

		items = append(items, redisdb.SetItem{
			Key:        cacheKeyPrefix + group,
			Value:      string(data),
			Expiration: c.ttl,
		})
	}

	if len(items) == 0 {
		return
	}

	if err := c.redis.SetPipelined(ctx, items); err != nil {
		rlog.Warnc(ctx, "failed to cache ACL entries in redis", rlog.Any("error", err))
	}
}
