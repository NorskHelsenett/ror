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

// recordingExpander records the seeds passed to ExpandScopes so tests can prove
// that expansion was (or was not) performed for a given seed. It can also be made
// to fail to exercise error-propagation paths.
type recordingExpander struct {
	expansions map[acl.Ownerref][]acl.Ownerref
	err        error

	scopeCalls  [][]acl.Ownerref // seeds for each ExpandScopes call
	singleCalls []acl.Ownerref   // args for each ExpandScope call
}

func (m *recordingExpander) ExpandScope(_ context.Context, scope aclscope.Scope, subject aclscope.Subject) ([]acl.Ownerref, error) {
	if m.err != nil {
		return nil, m.err
	}
	key := acl.Ownerref{Scope: scope, Subject: subject}
	m.singleCalls = append(m.singleCalls, key)
	return m.expansions[key], nil
}

func (m *recordingExpander) ExpandScopes(_ context.Context, seeds []acl.Ownerref) (map[acl.Ownerref][]acl.Ownerref, error) {
	if m.err != nil {
		return nil, m.err
	}
	m.scopeCalls = append(m.scopeCalls, append([]acl.Ownerref(nil), seeds...))
	result := make(map[acl.Ownerref][]acl.Ownerref, len(seeds))
	for _, s := range seeds {
		result[s] = m.expansions[s]
	}
	return result, nil
}

// expandedSeeds returns the flattened set of seeds that actually reached the backend.
func (m *recordingExpander) expandedSeeds() []acl.Ownerref {
	var all []acl.Ownerref
	for _, c := range m.scopeCalls {
		all = append(all, c...)
	}
	return all
}

// --- OwnerrefFilter predicate unit tests ---

func TestOwnerrefFilter_IsEmpty(t *testing.T) {
	assert.True(t, acl.OwnerrefFilter{}.IsEmpty())
	assert.False(t, acl.OwnerrefFilter{Scopes: []aclscope.Scope{"KubernetesCluster"}}.IsEmpty())
	assert.False(t, acl.OwnerrefFilter{Subjects: []aclscope.Subject{"cluster-1"}}.IsEmpty())
	assert.False(t, acl.OwnerrefFilter{
		Scopes:   []aclscope.Scope{"KubernetesCluster"},
		Subjects: []aclscope.Subject{"cluster-1"},
	}.IsEmpty())
}

func TestOwnerrefFilter_Matches(t *testing.T) {
	ref := acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-1"}

	// Empty filter matches everything.
	assert.True(t, acl.OwnerrefFilter{}.Matches(ref))

	// Scope-only.
	assert.True(t, acl.OwnerrefFilter{Scopes: []aclscope.Scope{"KubernetesCluster"}}.Matches(ref))
	assert.False(t, acl.OwnerrefFilter{Scopes: []aclscope.Scope{"Project"}}.Matches(ref))

	// Subject-only.
	assert.True(t, acl.OwnerrefFilter{Subjects: []aclscope.Subject{"cluster-1"}}.Matches(ref))
	assert.False(t, acl.OwnerrefFilter{Subjects: []aclscope.Subject{"cluster-2"}}.Matches(ref))

	// Both set — both dimensions MUST match (guards against too-wide results).
	assert.True(t, acl.OwnerrefFilter{
		Scopes:   []aclscope.Scope{"KubernetesCluster"},
		Subjects: []aclscope.Subject{"cluster-1"},
	}.Matches(ref))
	// Scope matches but subject does not → reject.
	assert.False(t, acl.OwnerrefFilter{
		Scopes:   []aclscope.Scope{"KubernetesCluster"},
		Subjects: []aclscope.Subject{"cluster-2"},
	}.Matches(ref))
	// Subject matches but scope does not → reject.
	assert.False(t, acl.OwnerrefFilter{
		Scopes:   []aclscope.Scope{"Project"},
		Subjects: []aclscope.Subject{"cluster-1"},
	}.Matches(ref))
}

// --- ResolveOwnerrefs with a non-empty filter ---

// A scope filter must drop every ref whose scope is not requested, INCLUDING
// descendants produced by expansion. Without expansion the project seed would be
// dropped but the user must still reach its clusters, so the expander must run.
func TestResolver_ResolveOwnerrefs_ScopeFilter_KeepsOnlyMatchingScope(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "Project",
			Subject: "proj-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	expander := &recordingExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Project", Subject: "proj-1"}: {
				{Scope: "Workspace", Subject: "ws-dev"},
				{Scope: "KubernetesCluster", Subject: "cluster-abc"},
				{Scope: "KubernetesCluster", Subject: "cluster-def"},
			},
		},
	}
	resolver := acl.NewResolver(store, acl.WithScopeExpander(expander))

	filter := acl.OwnerrefFilter{Scopes: []aclscope.Scope{"KubernetesCluster"}}
	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"dev-team"}, "ror:read", filter)
	assert.NoError(t, err)

	// Only the two clusters survive; the Project seed and Workspace are filtered out.
	assert.Len(t, refs, 2)
	assert.Contains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-abc"})
	assert.Contains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-def"})
	assert.NotContains(t, refs, acl.Ownerref{Scope: "Project", Subject: "proj-1"})
	assert.NotContains(t, refs, acl.Ownerref{Scope: "Workspace", Subject: "ws-dev"})

	// proj-1 differs from the filter scope, so it had to be expanded.
	assert.Contains(t, expander.expandedSeeds(), acl.Ownerref{Scope: "Project", Subject: "proj-1"})
}

// When the filter targets exactly the entry's own scope, expansion is skipped.
// This must NOT pull in any descendants (which would be a too-wide result) and
// must NOT call the backend for that seed.
func TestResolver_ResolveOwnerrefs_ScopeFilter_SkipsExpansionNoLeak(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"kubernetes:logon"},
		},
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-2",
			Access:  []aclmodels.AccessTypeV3{"kubernetes:logon"},
		},
	)
	// If the expander were (incorrectly) consulted, it would leak a bogus child.
	expander := &recordingExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "KubernetesCluster", Subject: "cluster-1"}: {
				{Scope: "KubernetesCluster", Subject: "LEAKED-CHILD"},
			},
		},
	}
	resolver := acl.NewResolver(store, acl.WithScopeExpander(expander))

	filter := acl.OwnerrefFilter{Scopes: []aclscope.Scope{"KubernetesCluster"}}
	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"dev-team"}, "kubernetes:logon", filter)
	assert.NoError(t, err)

	assert.Len(t, refs, 2)
	assert.Contains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-1"})
	assert.Contains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-2"})
	assert.NotContains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "LEAKED-CHILD"})

	// Expansion was skipped entirely.
	assert.Empty(t, expander.expandedSeeds())
}

// A subject filter restricts the result to specific subjects, even across
// expansion. Expansion still runs because a subject restriction can match a
// descendant.
func TestResolver_ResolveOwnerrefs_SubjectFilter_KeepsOnlyMatchingSubject(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "Project",
			Subject: "proj-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	expander := &recordingExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Project", Subject: "proj-1"}: {
				{Scope: "KubernetesCluster", Subject: "cluster-abc"},
				{Scope: "KubernetesCluster", Subject: "cluster-def"},
			},
		},
	}
	resolver := acl.NewResolver(store, acl.WithScopeExpander(expander))

	filter := acl.OwnerrefFilter{Subjects: []aclscope.Subject{"cluster-abc"}}
	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"dev-team"}, "ror:read", filter)
	assert.NoError(t, err)

	assert.Len(t, refs, 1)
	assert.Contains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-abc"})
	assert.NotContains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-def"})
}

// Both dimensions set: only the single ref matching scope AND subject survives.
func TestResolver_ResolveOwnerrefs_ScopeAndSubjectFilter(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "Project",
			Subject: "proj-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	expander := &recordingExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Project", Subject: "proj-1"}: {
				{Scope: "Workspace", Subject: "cluster-def"}, // subject matches but scope does not
				{Scope: "KubernetesCluster", Subject: "cluster-def"},
				{Scope: "KubernetesCluster", Subject: "cluster-abc"},
			},
		},
	}
	resolver := acl.NewResolver(store, acl.WithScopeExpander(expander))

	filter := acl.OwnerrefFilter{
		Scopes:   []aclscope.Scope{"KubernetesCluster"},
		Subjects: []aclscope.Subject{"cluster-def"},
	}
	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"dev-team"}, "ror:read", filter)
	assert.NoError(t, err)

	assert.Len(t, refs, 1)
	assert.Contains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-def"})
}

// Mixed seeds: a leaf seed matching the filter scope is NOT expanded, while a
// higher seed of a different scope IS expanded to surface its matching children.
func TestResolver_ResolveOwnerrefs_ScopeFilter_MixedSeeds(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"kubernetes:logon"},
		},
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "Project",
			Subject: "proj-1",
			Access:  []aclmodels.AccessTypeV3{"kubernetes:logon"},
		},
	)
	expander := &recordingExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Project", Subject: "proj-1"}: {
				{Scope: "KubernetesCluster", Subject: "cluster-abc"},
			},
			// cluster-1 has a (bogus) expansion that must never be consulted.
			{Scope: "KubernetesCluster", Subject: "cluster-1"}: {
				{Scope: "KubernetesCluster", Subject: "LEAKED-CHILD"},
			},
		},
	}
	resolver := acl.NewResolver(store, acl.WithScopeExpander(expander))

	filter := acl.OwnerrefFilter{Scopes: []aclscope.Scope{"KubernetesCluster"}}
	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"dev-team"}, "kubernetes:logon", filter)
	assert.NoError(t, err)

	assert.Len(t, refs, 2)
	assert.Contains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-1"})
	assert.Contains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-abc"})
	assert.NotContains(t, refs, acl.Ownerref{Scope: "KubernetesCluster", Subject: "LEAKED-CHILD"})

	// Only proj-1 was expanded; cluster-1 was skipped.
	seeds := expander.expandedSeeds()
	assert.Contains(t, seeds, acl.Ownerref{Scope: "Project", Subject: "proj-1"})
	assert.NotContains(t, seeds, acl.Ownerref{Scope: "KubernetesCluster", Subject: "cluster-1"})
}

// A filter that matches nothing must return an empty, NON-nil slice. Returning
// nil here would be read as "unrestricted" downstream — the worst possible
// too-wide result.
func TestResolver_ResolveOwnerrefs_Filter_NoMatch_EmptyNotNil(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	resolver := acl.NewResolver(store)

	filter := acl.OwnerrefFilter{Scopes: []aclscope.Scope{"Project"}}
	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"dev-team"}, "ror:read", filter)
	assert.NoError(t, err)
	assert.Empty(t, refs)
	assert.NotNil(t, refs)
}

// Global/unrestricted grants short-circuit and return nil REGARDLESS of the
// filter. This pins the (intentional) behaviour: a global grant means the caller
// is unrestricted, so a narrowing filter cannot make the result safer here.
func TestResolver_ResolveOwnerrefs_GlobalGrant_IgnoresFilter(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "admins",
			Scope:   "all",
			Subject: aclscope.SubjectAll,
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	expander := &recordingExpander{}
	resolver := acl.NewResolver(store, acl.WithScopeExpander(expander))

	filter := acl.OwnerrefFilter{Scopes: []aclscope.Scope{"KubernetesCluster"}}
	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"admins"}, "ror:read", filter)
	assert.NoError(t, err)
	assert.Nil(t, refs) // nil = unrestricted
	assert.Empty(t, expander.expandedSeeds())
}

// A global grant expressed as the "ror" scope with the "globalscope" subject
// (the standard global access grant) must be treated as unrestricted, just like
// the "all" scope/subject. This mirrors matchesScopeSubject's global semantics.
func TestResolver_ResolveOwnerrefs_RorGlobalScope_IsUnrestricted(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "super-admins",
			Scope:   aclscope.ScopeRor,
			Subject: aclscope.SubjectGlobal,
			Access:  []aclmodels.AccessTypeV3{"kubernetes:logon"},
		},
	)
	expander := &recordingExpander{}
	resolver := acl.NewResolver(store, acl.WithScopeExpander(expander))

	filter := acl.OwnerrefFilter{Scopes: []aclscope.Scope{"KubernetesCluster"}}
	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"super-admins"}, "kubernetes:logon", filter)
	assert.NoError(t, err)
	assert.Nil(t, refs) // nil = unrestricted
	assert.Empty(t, expander.expandedSeeds())
}

// An error from the expander backend must propagate, never silently yielding a
// partial (and possibly misleading) result.
func TestResolver_ResolveOwnerrefs_ExpanderError_Propagates(t *testing.T) {
	store := newMockStore(
		aclmodels.AclV3ListItem{
			Group:   "dev-team",
			Scope:   "Project",
			Subject: "proj-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
	)
	expander := &recordingExpander{err: fmt.Errorf("mongo down")}
	resolver := acl.NewResolver(store, acl.WithScopeExpander(expander))

	refs, err := resolver.ResolveOwnerrefs(context.Background(), []string{"dev-team"}, "ror:read", acl.OwnerrefFilter{})
	assert.Error(t, err)
	assert.Nil(t, refs)
}

// --- CachedScopeExpander.ExpandScopes (batched) ---

func TestCachedScopeExpander_ExpandScopes_AllMiss_OneBackendCall(t *testing.T) {
	backend := &recordingExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Project", Subject: "proj-1"}:   {{Scope: "KubernetesCluster", Subject: "c1"}},
			{Scope: "Workspace", Subject: "ws-dev"}: {{Scope: "KubernetesCluster", Subject: "c2"}},
		},
	}
	cached := acl.NewCachedScopeExpander(backend, 5*time.Minute)

	seeds := []acl.Ownerref{
		{Scope: "Project", Subject: "proj-1"},
		{Scope: "Workspace", Subject: "ws-dev"},
	}
	result, err := cached.ExpandScopes(context.Background(), seeds)
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, []acl.Ownerref{{Scope: "KubernetesCluster", Subject: "c1"}}, result[seeds[0]])
	assert.Equal(t, []acl.Ownerref{{Scope: "KubernetesCluster", Subject: "c2"}}, result[seeds[1]])

	// All misses resolved in a single batched backend call.
	assert.Len(t, backend.scopeCalls, 1)
	assert.ElementsMatch(t, seeds, backend.scopeCalls[0])
}

func TestCachedScopeExpander_ExpandScopes_PartialHit_OnlyMissesHitBackend(t *testing.T) {
	backend := &recordingExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Project", Subject: "proj-1"}:   {{Scope: "KubernetesCluster", Subject: "c1"}},
			{Scope: "Workspace", Subject: "ws-dev"}: {{Scope: "KubernetesCluster", Subject: "c2"}},
		},
	}
	cached := acl.NewCachedScopeExpander(backend, 5*time.Minute)

	// Warm proj-1 only.
	_, err := cached.ExpandScopes(context.Background(), []acl.Ownerref{{Scope: "Project", Subject: "proj-1"}})
	assert.NoError(t, err)
	assert.Len(t, backend.scopeCalls, 1)

	// Request both — proj-1 from cache, ws-dev from backend.
	result, err := cached.ExpandScopes(context.Background(), []acl.Ownerref{
		{Scope: "Project", Subject: "proj-1"},
		{Scope: "Workspace", Subject: "ws-dev"},
	})
	assert.NoError(t, err)
	assert.Len(t, result, 2)

	// Second backend call contains only the miss.
	assert.Len(t, backend.scopeCalls, 2)
	assert.Equal(t, []acl.Ownerref{{Scope: "Workspace", Subject: "ws-dev"}}, backend.scopeCalls[1])
}

func TestCachedScopeExpander_ExpandScopes_AllHit_NoBackendCall(t *testing.T) {
	backend := &recordingExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Project", Subject: "proj-1"}: {{Scope: "KubernetesCluster", Subject: "c1"}},
		},
	}
	cached := acl.NewCachedScopeExpander(backend, 5*time.Minute)

	seeds := []acl.Ownerref{{Scope: "Project", Subject: "proj-1"}}
	_, err := cached.ExpandScopes(context.Background(), seeds)
	assert.NoError(t, err)
	assert.Len(t, backend.scopeCalls, 1)

	_, err = cached.ExpandScopes(context.Background(), seeds)
	assert.NoError(t, err)
	assert.Len(t, backend.scopeCalls, 1) // still 1, served from cache
}

func TestCachedScopeExpander_ExpandScopes_DuplicateSeeds_BackendOnce(t *testing.T) {
	backend := &recordingExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Project", Subject: "proj-1"}: {{Scope: "KubernetesCluster", Subject: "c1"}},
		},
	}
	cached := acl.NewCachedScopeExpander(backend, 5*time.Minute)

	seed := acl.Ownerref{Scope: "Project", Subject: "proj-1"}
	result, err := cached.ExpandScopes(context.Background(), []acl.Ownerref{seed, seed, seed})
	assert.NoError(t, err)
	assert.Len(t, result, 1)

	// Duplicate seeds collapse to a single backend lookup.
	assert.Len(t, backend.scopeCalls, 1)
	assert.Equal(t, []acl.Ownerref{seed}, backend.scopeCalls[0])
}

func TestCachedScopeExpander_ExpandScopes_NilExpansion_Cached(t *testing.T) {
	backend := &recordingExpander{expansions: map[acl.Ownerref][]acl.Ownerref{}}
	cached := acl.NewCachedScopeExpander(backend, 5*time.Minute)

	seed := acl.Ownerref{Scope: "KubernetesCluster", Subject: "leaf"}
	result, err := cached.ExpandScopes(context.Background(), []acl.Ownerref{seed})
	assert.NoError(t, err)
	assert.Nil(t, result[seed])
	assert.Len(t, backend.scopeCalls, 1)

	// Second call served from cache despite the nil expansion.
	result, err = cached.ExpandScopes(context.Background(), []acl.Ownerref{seed})
	assert.NoError(t, err)
	assert.Nil(t, result[seed])
	assert.Len(t, backend.scopeCalls, 1)
}

func TestCachedScopeExpander_ExpandScopes_Empty(t *testing.T) {
	backend := &recordingExpander{expansions: map[acl.Ownerref][]acl.Ownerref{}}
	cached := acl.NewCachedScopeExpander(backend, 5*time.Minute)

	result, err := cached.ExpandScopes(context.Background(), nil)
	assert.NoError(t, err)
	assert.Empty(t, result)
	assert.Empty(t, backend.scopeCalls)
}

func TestCachedScopeExpander_ExpandScopes_BackendError_Propagates(t *testing.T) {
	backend := &recordingExpander{err: fmt.Errorf("mongo down")}
	cached := acl.NewCachedScopeExpander(backend, 5*time.Minute)

	result, err := cached.ExpandScopes(context.Background(), []acl.Ownerref{{Scope: "Project", Subject: "proj-1"}})
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCachedScopeExpander_ExpandScopes_TTLExpiry_Refetches(t *testing.T) {
	backend := &recordingExpander{
		expansions: map[acl.Ownerref][]acl.Ownerref{
			{Scope: "Project", Subject: "proj-1"}: {{Scope: "KubernetesCluster", Subject: "c1"}},
		},
	}
	cached := acl.NewCachedScopeExpander(backend, time.Nanosecond)

	seeds := []acl.Ownerref{{Scope: "Project", Subject: "proj-1"}}
	_, err := cached.ExpandScopes(context.Background(), seeds)
	assert.NoError(t, err)
	assert.Len(t, backend.scopeCalls, 1)

	time.Sleep(time.Millisecond) // let the entry expire

	_, err = cached.ExpandScopes(context.Background(), seeds)
	assert.NoError(t, err)
	assert.Len(t, backend.scopeCalls, 2) // re-fetched after expiry
}
