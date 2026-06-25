package aclstore

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NorskHelsenett/ror/pkg/acl"
	"github.com/NorskHelsenett/ror/pkg/clients/redisdb"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/telemetry/rortracer"

	"go.opentelemetry.io/otel/attribute"
)

const (
	cacheKeyPrefix   = "acl:v3:group:"
	cacheKeyV2Prefix = "acl:v2:group:"
)

// CachedStore wraps an acl.Store backend with a Redis cache layer.
// Cache is per-group: each group's full ACL entry list is stored as a single JSON blob.
// On GetByGroups, it does MGET for all groups, backfills misses from the backend,
// and caches the backfilled results.
// If Redis is unavailable, it falls through to the backend transparently.
type CachedStore struct {
	backend acl.Store
	redis   redisdb.RedisDB
	ttl     time.Duration
}

// NewCachedStore creates a cached store wrapping the given backend.
func NewCachedStore(backend acl.Store, redis redisdb.RedisDB, ttl time.Duration) *CachedStore {
	return &CachedStore{
		backend: backend,
		redis:   redis,
		ttl:     ttl,
	}
}

// GetByGroups returns ACL entries for all groups, using Redis cache where possible.
// Cache misses are backfilled from the backend in a single call.
func (c *CachedStore) GetByGroups(ctx context.Context, groups []string) (map[string][]aclmodels.AclV3ListItem, error) {
	ctx, span := rortracer.StartSpan(ctx, "acl.CachedStore.GetByGroups")
	defer span.End()
	span.SetAttributes(attribute.Int("acl.groups", len(groups)))

	if len(groups) == 0 {
		return make(map[string][]aclmodels.AclV3ListItem), nil
	}

	result := make(map[string][]aclmodels.AclV3ListItem, len(groups))

	// Try cache first
	hits, misses, cacheErr := c.mget(ctx, groups)
	if cacheErr != nil {
		// Redis down — fall through to backend for all groups
		span.SetAttributes(attribute.Bool("acl.cache_available", false))
		rlog.Warnc(ctx, "redis cache unavailable, falling through to backend", rlog.Any("error", cacheErr))
		return c.backend.GetByGroups(ctx, groups)
	}
	span.SetAttributes(
		attribute.Bool("acl.cache_available", true),
		attribute.Int("acl.cache_hits", len(hits)),
		attribute.Int("acl.cache_misses", len(misses)),
	)

	// Add hits to result (preserve backend behavior: omit groups with no entries)
	for group, entries := range hits {
		if len(entries) > 0 {
			result[group] = entries
		}
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

// GetV2ByGroups returns ACL entries as V2 items, using Redis cache where possible.
func (c *CachedStore) GetV2ByGroups(ctx context.Context, groups []string) (map[string][]aclmodels.AclV2ListItem, error) {
	if len(groups) == 0 {
		return make(map[string][]aclmodels.AclV2ListItem), nil
	}

	result := make(map[string][]aclmodels.AclV2ListItem, len(groups))

	hits, misses, cacheErr := c.mgetV2(ctx, groups)
	if cacheErr != nil {
		rlog.Warnc(ctx, "redis cache unavailable, falling through to backend", rlog.Any("error", cacheErr))
		return c.backend.GetV2ByGroups(ctx, groups)
	}

	for group, entries := range hits {
		if len(entries) > 0 {
			result[group] = entries
		}
	}

	if len(misses) == 0 {
		return result, nil
	}

	backfilled, err := c.backend.GetV2ByGroups(ctx, misses)
	if err != nil {
		return nil, err
	}

	c.setManyV2(ctx, misses, backfilled)

	for _, group := range misses {
		if entries, ok := backfilled[group]; ok {
			result[group] = entries
		}
	}

	return result, nil
}

// Invalidate removes the cached entries for a group (both V2 and V3 caches).
// Call this when ACL entries for the group are created, updated, or deleted.
func (c *CachedStore) Invalidate(ctx context.Context, group string) error {
	keyV3 := cacheKeyPrefix + group
	keyV2 := cacheKeyV2Prefix + group
	// Delete both; ignore individual errors but return first failure
	errV3 := c.redis.Delete(ctx, keyV3)
	errV2 := c.redis.Delete(ctx, keyV2)
	if errV3 != nil {
		return errV3
	}
	return errV2
}

// mget fetches cached entries for all groups in a single MGET call.
// Returns hits (group → entries), misses (group names not in cache), and any error.
func (c *CachedStore) mget(ctx context.Context, groups []string) (map[string][]aclmodels.AclV3ListItem, []string, error) {
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
func (c *CachedStore) setMany(ctx context.Context, groups []string, backfilled map[string][]aclmodels.AclV3ListItem) {
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

// mgetV2 fetches cached V2 entries for all groups in a single MGET call.
func (c *CachedStore) mgetV2(ctx context.Context, groups []string) (map[string][]aclmodels.AclV2ListItem, []string, error) {
	keys := make([]string, len(groups))
	for i, g := range groups {
		keys[i] = cacheKeyV2Prefix + g
	}

	vals, err := c.redis.MGet(ctx, keys...)
	if err != nil {
		return nil, nil, err
	}

	hits := make(map[string][]aclmodels.AclV2ListItem)
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

		var entries []aclmodels.AclV2ListItem
		if err := json.Unmarshal([]byte(str), &entries); err != nil {
			rlog.Warnc(ctx, fmt.Sprintf("failed to unmarshal cached V2 ACL for group %q, treating as miss", group), rlog.Any("error", err))
			misses = append(misses, group)
			continue
		}
		hits[group] = entries
	}

	return hits, misses, nil
}

// setManyV2 caches V2 entries for the given groups using a Redis pipeline.
func (c *CachedStore) setManyV2(ctx context.Context, groups []string, backfilled map[string][]aclmodels.AclV2ListItem) {
	var items []redisdb.SetItem

	for _, group := range groups {
		entries := backfilled[group]
		if entries == nil {
			entries = []aclmodels.AclV2ListItem{}
		}

		data, err := json.Marshal(entries)
		if err != nil {
			rlog.Warnc(ctx, fmt.Sprintf("failed to marshal V2 ACL for group %q, skipping cache", group), rlog.Any("error", err))
			continue
		}

		items = append(items, redisdb.SetItem{
			Key:        cacheKeyV2Prefix + group,
			Value:      string(data),
			Expiration: c.ttl,
		})
	}

	if len(items) == 0 {
		return
	}

	if err := c.redis.SetPipelined(ctx, items); err != nil {
		rlog.Warnc(ctx, "failed to cache V2 ACL entries in redis", rlog.Any("error", err))
	}
}
