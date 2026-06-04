package acl_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/acl"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/stretchr/testify/assert"
)

// mockStore implements acl.Store for testing.
type mockStore struct {
	entries map[string][]aclmodels.AclV3ListItem
	err     error
}

func (m *mockStore) GetByGroups(ctx context.Context, groups []string) (map[string][]aclmodels.AclV3ListItem, error) {
	if m.err != nil {
		return nil, m.err
	}
	result := make(map[string][]aclmodels.AclV3ListItem)
	for _, g := range groups {
		if entries, ok := m.entries[g]; ok {
			result[g] = entries
		}
	}
	return result, nil
}

func newMockStore(entries ...aclmodels.AclV3ListItem) *mockStore {
	m := &mockStore{entries: make(map[string][]aclmodels.AclV3ListItem)}
	for _, e := range entries {
		m.entries[e.Group] = append(m.entries[e.Group], e)
	}
	return m
}

func TestResolver_ResolveAccess_ExactMatch(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read", "kubernetes:logon"},
		},
	)
	resolver := acl.NewResolver(store)

	access, err := resolver.ResolveAccess(context.Background(), []string{"dev-team"}, "KubernetesCluster", "cluster-1")
	assert.NoError(t, err)
	assert.Len(t, access, 2)
	assert.Contains(t, access, aclmodels.AccessTypeV3("ror:read"))
	assert.Contains(t, access, aclmodels.AccessTypeV3("kubernetes:logon"))
}

func TestResolver_ResolveAccess_NoMatch(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	resolver := acl.NewResolver(store)

	access, err := resolver.ResolveAccess(context.Background(), []string{"dev-team"}, "KubernetesCluster", "cluster-2")
	assert.NoError(t, err)
	assert.Empty(t, access)
}

func TestResolver_ResolveAccess_MultipleGroups(t *testing.T) {
	store := newMockStore(
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
			Access:  []aclmodels.AccessTypeV3{"ror:read", "ror:write", "kubernetes:admin"},
		},
	)
	resolver := acl.NewResolver(store)

	access, err := resolver.ResolveAccess(context.Background(), []string{"dev-team", "ops-team"}, "KubernetesCluster", "cluster-1")
	assert.NoError(t, err)
	assert.Len(t, access, 3)
	assert.Contains(t, access, aclmodels.AccessTypeV3("ror:read"))
	assert.Contains(t, access, aclmodels.AccessTypeV3("ror:write"))
	assert.Contains(t, access, aclmodels.AccessTypeV3("kubernetes:admin"))
}

func TestResolver_ResolveAccess_GlobalScope(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "admins",
			Scope:   "all",
			Subject: "All",
			Access:  []aclmodels.AccessTypeV3{"ror:read", "ror:write", "ror:owner"},
		},
	)
	resolver := acl.NewResolver(store)

	// Global entry should match any scope+subject
	access, err := resolver.ResolveAccess(context.Background(), []string{"admins"}, "KubernetesCluster", "random-cluster")
	assert.NoError(t, err)
	assert.Len(t, access, 3)
}

func TestResolver_ResolveAccess_RorScopeWithSubjectAsScope(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "cluster-admins",
			Scope:   "ror",
			Subject: "KubernetesCluster",
			Access:  []aclmodels.AccessTypeV3{"ror:read", "ror:write"},
		},
	)
	resolver := acl.NewResolver(store)

	// scope=ror, subject=KubernetesCluster should match requests for scope KubernetesCluster
	access, err := resolver.ResolveAccess(context.Background(), []string{"cluster-admins"}, "KubernetesCluster", "any-cluster")
	assert.NoError(t, err)
	assert.Len(t, access, 2)
}

func TestResolver_ResolveAccess_RorScopeGlobalSubject(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "super-admins",
			Scope:   "ror",
			Subject: "globalscope",
			Access:  []aclmodels.AccessTypeV3{"ror:owner"},
		},
	)
	resolver := acl.NewResolver(store)

	// scope=ror, subject=globalscope should match everything
	access, err := resolver.ResolveAccess(context.Background(), []string{"super-admins"}, "Project", "my-project")
	assert.NoError(t, err)
	assert.Contains(t, access, aclmodels.AccessTypeV3("ror:owner"))
}

func TestResolver_ResolveAccess_StoreError(t *testing.T) {
	store := &mockStore{err: fmt.Errorf("connection refused")}
	resolver := acl.NewResolver(store)

	access, err := resolver.ResolveAccess(context.Background(), []string{"dev-team"}, "KubernetesCluster", "cluster-1")
	assert.Error(t, err)
	assert.Nil(t, access)
}

func TestResolver_ResolveAccess_EmptyGroups(t *testing.T) {
	store := newMockStore()
	resolver := acl.NewResolver(store)

	access, err := resolver.ResolveAccess(context.Background(), []string{}, "KubernetesCluster", "cluster-1")
	assert.NoError(t, err)
	assert.Empty(t, access)
}

func TestResolver_HasAccess_True(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read", "kubernetes:logon"},
		},
	)
	resolver := acl.NewResolver(store)

	ok, err := resolver.HasAccess(context.Background(), []string{"dev-team"}, "KubernetesCluster", "cluster-1", "ror:read")
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestResolver_HasAccess_False(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	resolver := acl.NewResolver(store)

	ok, err := resolver.HasAccess(context.Background(), []string{"dev-team"}, "KubernetesCluster", "cluster-1", "ror:write")
	assert.NoError(t, err)
	assert.False(t, ok)
}

func TestResolver_ResolveOwnerrefs_Basic(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-2",
			Access:  []aclmodels.AccessTypeV3{"ror:read", "ror:write"},
		},
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "Project",
			Subject: "project-1",
			Access:  []aclmodels.AccessTypeV3{"ror:write"},
		},
	)
	resolver := acl.NewResolver(store)

	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"dev-team"}, "ror:read")
	assert.NoError(t, err)
	assert.Len(t, refs, 2)
	assert.Contains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-1"})
	assert.Contains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-2"})
}

func TestResolver_ResolveOwnerrefs_GlobalReturnsNil(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "admins",
			Scope:   "all",
			Subject: "All",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	resolver := acl.NewResolver(store)

	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"admins"}, "ror:read")
	assert.NoError(t, err)
	assert.Nil(t, refs) // nil = unrestricted
}

func TestResolver_ResolveOwnerrefs_Deduplicates(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "team-a",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
		aclmodels.AclV3ListItem{
			Group:   "team-b",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	resolver := acl.NewResolver(store)

	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"team-a", "team-b"}, "ror:read")
	assert.NoError(t, err)
	assert.Len(t, refs, 1) // same scope+subject, deduplicated
}

func TestResolver_ResolveOwnerrefs_NoMatchingAccess(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	resolver := acl.NewResolver(store)

	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"dev-team"}, "ror:write")
	assert.NoError(t, err)
	assert.Empty(t, refs)
	assert.NotNil(t, refs) // empty but not nil — nil means unrestricted
}

func TestResolver_ResolveOwnerrefs_SubjectAll(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "ops",
			Scope:   "KubernetesCluster",
			Subject: "all",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	resolver := acl.NewResolver(store)

	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"ops"}, "ror:read")
	assert.NoError(t, err)
	assert.Nil(t, refs) // subject=all → unrestricted
}

func TestResolver_ResolveAccess_ManyGroups(t *testing.T) {
	// Simulates a user with many groups, each contributing different access
	var entries []aclmodels.AclV3ListItem
	for i := range 50 {
		entries = append(entries, aclmodels.AclV3ListItem{
			Group:   fmt.Sprintf("group-%d", i),
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{aclmodels.AccessTypeV3(fmt.Sprintf("ror:test%d:read", i))},
		})
	}
	store := newMockStore(entries...)
	resolver := acl.NewResolver(store)

	groups := make([]string, 50)
	for i := range 50 {
		groups[i] = fmt.Sprintf("group-%d", i)
	}

	access, err := resolver.ResolveAccess(context.Background(), groups, "KubernetesCluster", "cluster-1")
	assert.NoError(t, err)
	assert.Len(t, access, 50)
}

// --- ScopeExpander tests ---

// mockExpander implements acl.ScopeExpander for testing.
type mockExpander struct {
	expansions map[acl.Ownerref][]acl.Ownerref
	calls      int
}

func (m *mockExpander) ExpandScope(_ context.Context, scope aclscope.Scope, subject aclscope.Subject) ([]acl.Ownerref, error) {
	m.calls++
	key := acl.Ownerref{Scope: scope, Subject: subject}
	return m.expansions[key], nil
}

func TestResolver_ResolveOwnerrefs_WithExpander_ProjectExpands(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "Project",
			Subject: "proj-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	expander := &mockExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Project", Subject: "proj-1"}: {
				{Scope: "Workspace", Subject: "ws-dev"},
				{Scope: "KubernetesCluster", Subject: "cluster-abc"},
				{Scope: "KubernetesCluster", Subject: "cluster-def"},
			},
		},
	}
	resolver := acl.NewResolver(store, acl.WithScopeExpander(expander))

	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"dev-team"}, "ror:read")
	assert.NoError(t, err)
	// Original + 3 descendants
	assert.Len(t, refs, 4)
	assert.Contains(t, refs, acl.Ownerref{Scope: "Project", Subject: "proj-1"})
	assert.Contains(t, refs, acl.Ownerref{Scope: "Workspace", Subject: "ws-dev"})
	assert.Contains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-abc"})
	assert.Contains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-def"})
}

func TestResolver_ResolveOwnerrefs_WithExpander_LeafScope_NoExpansion(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	expander := &mockExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{},
	}
	resolver := acl.NewResolver(store, acl.WithScopeExpander(expander))

	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"dev-team"}, "ror:read")
	assert.NoError(t, err)
	assert.Len(t, refs, 1)
	assert.Contains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-1"})
}

func TestResolver_ResolveOwnerrefs_WithExpander_DeduplicatesAcrossEntries(t *testing.T) {
	// Two ACL entries that expand to overlapping clusters
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "Workspace",
			Subject: "ws-dev",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-abc",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	expander := &mockExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Workspace", Subject: "ws-dev"}: {
				{Scope: "KubernetesCluster", Subject: "cluster-abc"}, // overlaps with direct entry
				{Scope: "KubernetesCluster", Subject: "cluster-def"},
			},
		},
	}
	resolver := acl.NewResolver(store, acl.WithScopeExpander(expander))

	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"dev-team"}, "ror:read")
	assert.NoError(t, err)
	// ws-dev + cluster-abc (deduped) + cluster-def
	assert.Len(t, refs, 3)
}

func TestResolver_ResolveOwnerrefs_WithExpander_GlobalStillReturnsNil(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "admins",
			Scope:   "all",
			Subject: "All",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	expander := &mockExpander{}
	resolver := acl.NewResolver(store, acl.WithScopeExpander(expander))

	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"admins"}, "ror:read")
	assert.NoError(t, err)
	assert.Nil(t, refs) // nil = unrestricted, expander should not be called
	assert.Equal(t, 0, expander.calls)
}

func TestResolver_ResolveOwnerrefs_WithoutExpander_BackwardsCompatible(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "Project",
			Subject: "proj-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	// No expander — old behavior
	resolver := acl.NewResolver(store)

	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"dev-team"}, "ror:read")
	assert.NoError(t, err)
	assert.Len(t, refs, 1)
	assert.Contains(t, refs, acl.Ownerref{Scope: "Project", Subject: "proj-1"})
}

// --- CachedScopeExpander tests ---

func TestCachedScopeExpander_CachesResult(t *testing.T) {
	backend := &mockExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Project", Subject: "proj-1"}: {
				{Scope: "Workspace", Subject: "ws-dev"},
			},
		},
	}
	cached := acl.NewCachedScopeExpander(backend, 5*time.Minute)

	// First call — hits backend
	refs, err := cached.ExpandScope(context.Background(), "Project", "proj-1")
	assert.NoError(t, err)
	assert.Len(t, refs, 1)
	assert.Equal(t, 1, backend.calls)

	// Second call — from cache
	refs, err = cached.ExpandScope(context.Background(), "Project", "proj-1")
	assert.NoError(t, err)
	assert.Len(t, refs, 1)
	assert.Equal(t, 1, backend.calls) // still 1
}

func TestCachedScopeExpander_Invalidate(t *testing.T) {
	backend := &mockExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Project", Subject: "proj-1"}: {
				{Scope: "Workspace", Subject: "ws-dev"},
			},
		},
	}
	cached := acl.NewCachedScopeExpander(backend, 5*time.Minute)

	_, _ = cached.ExpandScope(context.Background(), "Project", "proj-1")
	assert.Equal(t, 1, backend.calls)

	cached.Invalidate("Project", "proj-1")

	_, _ = cached.ExpandScope(context.Background(), "Project", "proj-1")
	assert.Equal(t, 2, backend.calls) // re-fetched
}

func TestCachedScopeExpander_InvalidateAll(t *testing.T) {
	backend := &mockExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Project", Subject: "proj-1"}: {
				{Scope: "Workspace", Subject: "ws-dev"},
			},
			{Scope: "Workspace", Subject: "ws-dev"}: {
				{Scope: "KubernetesCluster", Subject: "cluster-1"},
			},
		},
	}
	cached := acl.NewCachedScopeExpander(backend, 5*time.Minute)

	_, _ = cached.ExpandScope(context.Background(), "Project", "proj-1")
	_, _ = cached.ExpandScope(context.Background(), "Workspace", "ws-dev")
	assert.Equal(t, 2, backend.calls)

	cached.InvalidateAll()

	_, _ = cached.ExpandScope(context.Background(), "Project", "proj-1")
	_, _ = cached.ExpandScope(context.Background(), "Workspace", "ws-dev")
	assert.Equal(t, 4, backend.calls)
}

func TestCachedScopeExpander_NilExpansion_Cached(t *testing.T) {
	backend := &mockExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{},
	}
	cached := acl.NewCachedScopeExpander(backend, 5*time.Minute)

	// Leaf scope — nil expansion
	refs, err := cached.ExpandScope(context.Background(), "KubernetesCluster", "cluster-1")
	assert.NoError(t, err)
	assert.Nil(t, refs)
	assert.Equal(t, 1, backend.calls)

	// Should be cached
	_, _ = cached.ExpandScope(context.Background(), "KubernetesCluster", "cluster-1")
	assert.Equal(t, 1, backend.calls)
}
