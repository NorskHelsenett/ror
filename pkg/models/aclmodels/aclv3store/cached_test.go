package aclv3store_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients/redisdb"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclv3store"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// mockBackend implements aclmodels.AclV3Store as a test backend.
type mockBackend struct {
	entries map[string][]aclmodels.AclV3ListItem
	calls   int // counts how many times GetByGroups was called
}

func newMockBackend(entries ...aclmodels.AclV3ListItem) *mockBackend {
	m := &mockBackend{entries: make(map[string][]aclmodels.AclV3ListItem)}
	for _, e := range entries {
		m.entries[e.Group] = append(m.entries[e.Group], e)
	}
	return m
}

func (m *mockBackend) GetByGroups(ctx context.Context, groups []string) (map[string][]aclmodels.AclV3ListItem, error) {
	m.calls++
	result := make(map[string][]aclmodels.AclV3ListItem)
	for _, g := range groups {
		if entries, ok := m.entries[g]; ok {
			result[g] = entries
		}
	}
	return result, nil
}

func setupTest(t *testing.T, entries ...aclmodels.AclV3ListItem) (*aclv3store.CachedAclV3Store, *mockBackend, *miniredis.Miniredis) {
	t.Helper()
	mr := miniredis.RunT(t)
	rc := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	rdb := redisdb.NewFromClient(rc)
	backend := newMockBackend(entries...)
	store := aclv3store.NewCachedAclV3Store(backend, rdb, 5*time.Minute)
	return store, backend, mr
}

func TestCached_CacheMiss_BackfillsFromBackend(t *testing.T) {
	store, backend, _ := setupTest(t,
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)

	result, err := store.GetByGroups(context.Background(), []string{"dev-team"})
	assert.NoError(t, err)
	assert.Len(t, result["dev-team"], 1)
	assert.Equal(t, 1, backend.calls)
}

func TestCached_CacheHit_SkipsBackend(t *testing.T) {
	store, backend, _ := setupTest(t,
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)

	// First call — cache miss, hits backend
	_, err := store.GetByGroups(context.Background(), []string{"dev-team"})
	assert.NoError(t, err)
	assert.Equal(t, 1, backend.calls)

	// Second call — cache hit, should NOT hit backend
	result, err := store.GetByGroups(context.Background(), []string{"dev-team"})
	assert.NoError(t, err)
	assert.Len(t, result["dev-team"], 1)
	assert.Equal(t, 1, backend.calls) // still 1
}

func TestCached_PartialHit_OnlyBackfillsMisses(t *testing.T) {
	store, backend, _ := setupTest(t,
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
		aclmodels.AclV3ListItem{
			Group:   "ops-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:write"},
		},
	)

	// Warm cache for dev-team only
	_, err := store.GetByGroups(context.Background(), []string{"dev-team"})
	assert.NoError(t, err)
	assert.Equal(t, 1, backend.calls)

	// Request both — dev-team from cache, ops-team from backend
	result, err := store.GetByGroups(context.Background(), []string{"dev-team", "ops-team"})
	assert.NoError(t, err)
	assert.Len(t, result["dev-team"], 1)
	assert.Len(t, result["ops-team"], 1)
	assert.Equal(t, 2, backend.calls) // second call only for ops-team
}

func TestCached_EmptyGroups(t *testing.T) {
	store, backend, _ := setupTest(t)

	result, err := store.GetByGroups(context.Background(), []string{})
	assert.NoError(t, err)
	assert.Empty(t, result)
	assert.Equal(t, 0, backend.calls)
}

func TestCached_EmptyGroupCached(t *testing.T) {
	store, backend, _ := setupTest(t) // no entries in backend

	// First call — miss, backend returns nothing
	result, err := store.GetByGroups(context.Background(), []string{"empty-group"})
	assert.NoError(t, err)
	assert.Empty(t, result["empty-group"])
	assert.Equal(t, 1, backend.calls)

	// Second call — cached empty, should not hit backend
	_, err = store.GetByGroups(context.Background(), []string{"empty-group"})
	assert.NoError(t, err)
	assert.Equal(t, 1, backend.calls) // still 1
}

func TestCached_Invalidate(t *testing.T) {
	store, backend, _ := setupTest(t,
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)

	// Warm cache
	_, err := store.GetByGroups(context.Background(), []string{"dev-team"})
	assert.NoError(t, err)
	assert.Equal(t, 1, backend.calls)

	// Invalidate
	err = store.Invalidate(context.Background(), "dev-team")
	assert.NoError(t, err)

	// Next call should hit backend again
	_, err = store.GetByGroups(context.Background(), []string{"dev-team"})
	assert.NoError(t, err)
	assert.Equal(t, 2, backend.calls)
}

func TestCached_TTLExpiry(t *testing.T) {
	store, backend, mr := setupTest(t,
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)

	// Warm cache
	_, err := store.GetByGroups(context.Background(), []string{"dev-team"})
	assert.NoError(t, err)
	assert.Equal(t, 1, backend.calls)

	// Fast-forward past TTL
	mr.FastForward(6 * time.Minute)

	// Should miss cache, hit backend
	_, err = store.GetByGroups(context.Background(), []string{"dev-team"})
	assert.NoError(t, err)
	assert.Equal(t, 2, backend.calls)
}

func TestCached_RedisDown_FallsThrough(t *testing.T) {
	store, backend, mr := setupTest(t,
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)

	// Stop Redis
	mr.Close()

	// Should fall through to backend
	result, err := store.GetByGroups(context.Background(), []string{"dev-team"})
	assert.NoError(t, err)
	assert.Len(t, result["dev-team"], 1)
	assert.Equal(t, 1, backend.calls)
}

func TestCached_CorruptedCache_TreatedAsMiss(t *testing.T) {
	mr := miniredis.RunT(t)
	rc := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	rdb := redisdb.NewFromClient(rc)
	backend := newMockBackend(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	store := aclv3store.NewCachedAclV3Store(backend, rdb, 5*time.Minute)

	// Manually put garbage in cache
	mr.Set("acl:v3:group:dev-team", "not-valid-json{{{")

	// Should treat as miss and fetch from backend
	result, err := store.GetByGroups(context.Background(), []string{"dev-team"})
	assert.NoError(t, err)
	assert.Len(t, result["dev-team"], 1)
	assert.Equal(t, 1, backend.calls)
}

func TestCached_ManyGroups(t *testing.T) {
	var entries []aclmodels.AclV3ListItem
	for i := range 50 {
		entries = append(entries, aclmodels.AclV3ListItem{
			Group:   fmt.Sprintf("group-%d", i),
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		})
	}
	store, backend, _ := setupTest(t, entries...)

	groups := make([]string, 50)
	for i := range 50 {
		groups[i] = fmt.Sprintf("group-%d", i)
	}

	// First call — all miss
	result, err := store.GetByGroups(context.Background(), groups)
	assert.NoError(t, err)
	assert.Len(t, result, 50)
	assert.Equal(t, 1, backend.calls)

	// Second call — all hit
	result, err = store.GetByGroups(context.Background(), groups)
	assert.NoError(t, err)
	assert.Len(t, result, 50)
	assert.Equal(t, 1, backend.calls) // still 1
}

func TestCached_PreservesAccessData(t *testing.T) {
	original := aclmodels.AclV3ListItem{
		Group:   "dev-team",
		Scope:   "KubernetesCluster",
		Subject: "cluster-1",
		Access:  []aclmodels.AccessTypeV3{"ror:read", "kubernetes:logon", "resource:Deployment:read"},
	}
	store, _, _ := setupTest(t, original)

	// Warm cache
	_, err := store.GetByGroups(context.Background(), []string{"dev-team"})
	assert.NoError(t, err)

	// Fetch from cache and verify data integrity
	result, err := store.GetByGroups(context.Background(), []string{"dev-team"})
	assert.NoError(t, err)

	entries := result["dev-team"]
	assert.Len(t, entries, 1)
	assert.Equal(t, aclmodels.Acl3Scope("KubernetesCluster"), entries[0].Scope)
	assert.Equal(t, aclmodels.Acl3Subject("cluster-1"), entries[0].Subject)
	assert.Len(t, entries[0].Access, 3)
	assert.Contains(t, entries[0].Access, aclmodels.AccessTypeV3("ror:read"))
	assert.Contains(t, entries[0].Access, aclmodels.AccessTypeV3("kubernetes:logon"))
	assert.Contains(t, entries[0].Access, aclmodels.AccessTypeV3("resource:Deployment:read"))
}

func TestCached_CacheKeyFormat(t *testing.T) {
	store, _, mr := setupTest(t,
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)

	_, err := store.GetByGroups(context.Background(), []string{"dev-team"})
	assert.NoError(t, err)

	// Verify the Redis key format
	val, err := mr.Get("acl:v3:group:dev-team")
	assert.NoError(t, err)

	var entries []aclmodels.AclV3ListItem
	assert.NoError(t, json.Unmarshal([]byte(val), &entries))
	assert.Len(t, entries, 1)
}
