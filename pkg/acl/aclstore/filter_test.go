package aclstore_test

import (
	"fmt"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/acl"
	"github.com/NorskHelsenett/ror/pkg/acl/aclstore"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// --- OwnerrefsToFilter ---

func TestOwnerrefsToFilter_Nil_Unrestricted(t *testing.T) {
	// nil means the user has global/all access — no filter needed
	result := aclstore.OwnerrefsToFilter(nil)
	assert.Equal(t, bson.M{}, result)
}

func TestOwnerrefsToFilter_Empty_DenyAll(t *testing.T) {
	// Empty slice means the user has no matching access
	result := aclstore.OwnerrefsToFilter([]acl.Ownerref{})
	assert.Equal(t, aclstore.DenyAllFilter, result)
}

func TestOwnerrefsToFilter_SingleCluster(t *testing.T) {
	refs := []acl.Ownerref{
		{Scope: "cluster", Subject: "cluster-1"},
	}
	result := aclstore.OwnerrefsToFilter(refs)

	expected := bson.M{
		"$match": bson.M{
			"$or": bson.A{
				bson.M{
					"rormeta.ownerref": bson.M{"$in": bson.A{
						bson.D{{Key: "scope", Value: "cluster"}, {Key: "subject", Value: "cluster-1"}},
					}},
				},
			},
		},
	}
	assert.Equal(t, expected, result)
}

func TestOwnerrefsToFilter_MultipleClusters(t *testing.T) {
	refs := []acl.Ownerref{
		{Scope: "cluster", Subject: "cluster-1"},
		{Scope: "cluster", Subject: "cluster-2"},
		{Scope: "cluster", Subject: "cluster-3"},
	}
	result := aclstore.OwnerrefsToFilter(refs)

	match, ok := result["$match"].(bson.M)
	assert.True(t, ok)
	or, ok := match["$or"].(bson.A)
	assert.True(t, ok)
	assert.Len(t, or, 1) // single $in clause

	inClause := or[0].(bson.M)["rormeta.ownerref"].(bson.M)["$in"].(bson.A)
	assert.Len(t, inClause, 3)
}

func TestOwnerrefsToFilter_MixedScopes(t *testing.T) {
	refs := []acl.Ownerref{
		{Scope: "cluster", Subject: "cluster-1"},
		{Scope: "project", Subject: "proj-a"},
		{Scope: "cluster", Subject: "cluster-2"},
	}
	result := aclstore.OwnerrefsToFilter(refs)

	match := result["$match"].(bson.M)
	or := match["$or"].(bson.A)
	assert.Len(t, or, 1) // all in one $in

	inClause := or[0].(bson.M)["rormeta.ownerref"].(bson.M)["$in"].(bson.A)
	assert.Len(t, inClause, 3)

	// Verify sorted order: cluster before project
	first := inClause[0].(bson.D)
	assert.Equal(t, "cluster", first[0].Value)
}

func TestOwnerrefsToFilter_RorScopeGrant(t *testing.T) {
	// scope=ror, subject=cluster → grants access to ALL resources with ownerref.scope=cluster
	refs := []acl.Ownerref{
		{Scope: "ror", Subject: "cluster"},
	}
	result := aclstore.OwnerrefsToFilter(refs)

	expected := bson.M{
		"$match": bson.M{
			"$or": bson.A{
				bson.M{"rormeta.ownerref.scope": "cluster"},
			},
		},
	}
	assert.Equal(t, expected, result)
}

func TestOwnerrefsToFilter_RorScopeGrantMultiple(t *testing.T) {
	refs := []acl.Ownerref{
		{Scope: "ror", Subject: "cluster"},
		{Scope: "ror", Subject: "project"},
	}
	result := aclstore.OwnerrefsToFilter(refs)

	match := result["$match"].(bson.M)
	or := match["$or"].(bson.A)
	assert.Len(t, or, 2)
	assert.Equal(t, bson.M{"rormeta.ownerref.scope": "cluster"}, or[0])
	assert.Equal(t, bson.M{"rormeta.ownerref.scope": "project"}, or[1])
}

func TestOwnerrefsToFilter_RorScopeDeduplicatesSpecific(t *testing.T) {
	// If scope=ror grants cluster access, specific cluster refs should be excluded
	refs := []acl.Ownerref{
		{Scope: "ror", Subject: "cluster"},
		{Scope: "cluster", Subject: "cluster-1"}, // should be excluded — covered by ror grant
		{Scope: "cluster", Subject: "cluster-2"}, // same
		{Scope: "project", Subject: "proj-a"},    // NOT covered — should appear in $in
	}
	result := aclstore.OwnerrefsToFilter(refs)

	match := result["$match"].(bson.M)
	or := match["$or"].(bson.A)
	assert.Len(t, or, 2) // ror scope-level grant + $in for project

	// First: ror scope-level grant
	assert.Equal(t, bson.M{"rormeta.ownerref.scope": "cluster"}, or[0])

	// Second: $in for uncovered refs
	inClause := or[1].(bson.M)["rormeta.ownerref"].(bson.M)["$in"].(bson.A)
	assert.Len(t, inClause, 1)
	entry := inClause[0].(bson.D)
	assert.Equal(t, "project", entry[0].Value)
	assert.Equal(t, "proj-a", entry[1].Value)
}

func TestOwnerrefsToFilter_RorGlobalscope_Unrestricted(t *testing.T) {
	// scope=ror, subject=globalscope → unrestricted (defensive; resolver normally returns nil)
	refs := []acl.Ownerref{
		{Scope: "ror", Subject: "globalscope"},
		{Scope: "cluster", Subject: "cluster-1"},
	}
	result := aclstore.OwnerrefsToFilter(refs)
	assert.Equal(t, bson.M{}, result)
}

func TestOwnerrefsToFilter_RorGlobalscope_Alone(t *testing.T) {
	refs := []acl.Ownerref{
		{Scope: "ror", Subject: "globalscope"},
	}
	result := aclstore.OwnerrefsToFilter(refs)
	assert.Equal(t, bson.M{}, result)
}

func TestOwnerrefsToFilter_AllCoveredByRorScope(t *testing.T) {
	// All specific refs are covered by ror-level grants — only scope-level clauses
	refs := []acl.Ownerref{
		{Scope: "ror", Subject: "cluster"},
		{Scope: "cluster", Subject: "cluster-1"},
		{Scope: "cluster", Subject: "cluster-2"},
	}
	result := aclstore.OwnerrefsToFilter(refs)

	match := result["$match"].(bson.M)
	or := match["$or"].(bson.A)
	assert.Len(t, or, 1) // only the scope-level grant, no $in
	assert.Equal(t, bson.M{"rormeta.ownerref.scope": "cluster"}, or[0])
}

func TestOwnerrefsToFilter_SortedOutput(t *testing.T) {
	// Verify deterministic sorted output regardless of input order
	refs := []acl.Ownerref{
		{Scope: "project", Subject: "proj-b"},
		{Scope: "cluster", Subject: "cluster-z"},
		{Scope: "cluster", Subject: "cluster-a"},
		{Scope: "project", Subject: "proj-a"},
	}
	result := aclstore.OwnerrefsToFilter(refs)

	match := result["$match"].(bson.M)
	or := match["$or"].(bson.A)
	inClause := or[0].(bson.M)["rormeta.ownerref"].(bson.M)["$in"].(bson.A)

	assert.Len(t, inClause, 4)
	// cluster-a, cluster-z, proj-a, proj-b
	assert.Equal(t, "cluster", inClause[0].(bson.D)[0].Value)
	assert.Equal(t, "cluster-a", inClause[0].(bson.D)[1].Value)
	assert.Equal(t, "cluster", inClause[1].(bson.D)[0].Value)
	assert.Equal(t, "cluster-z", inClause[1].(bson.D)[1].Value)
	assert.Equal(t, "project", inClause[2].(bson.D)[0].Value)
	assert.Equal(t, "proj-a", inClause[2].(bson.D)[1].Value)
	assert.Equal(t, "project", inClause[3].(bson.D)[0].Value)
	assert.Equal(t, "proj-b", inClause[3].(bson.D)[1].Value)
}

func TestOwnerrefsToFilter_WithExpandedHierarchy(t *testing.T) {
	// Simulates resolver output after scope expansion:
	// User has access to Project "proj-1", which expands to Workspace "ws-dev"
	// and Clusters "cluster-abc", "cluster-def"
	refs := []acl.Ownerref{
		{Scope: "project", Subject: "proj-1"},
		{Scope: "Workspace", Subject: "ws-dev"},
		{Scope: "KubernetesCluster", Subject: "cluster-abc"},
		{Scope: "KubernetesCluster", Subject: "cluster-def"},
	}
	result := aclstore.OwnerrefsToFilter(refs)

	match := result["$match"].(bson.M)
	or := match["$or"].(bson.A)
	assert.Len(t, or, 1)

	inClause := or[0].(bson.M)["rormeta.ownerref"].(bson.M)["$in"].(bson.A)
	assert.Len(t, inClause, 4)
}

func TestOwnerrefsToFilter_OnlyRorScopeRefs(t *testing.T) {
	// Only ror-scope refs, no specific scope+subject pairs
	refs := []acl.Ownerref{
		{Scope: "ror", Subject: "cluster"},
		{Scope: "ror", Subject: "project"},
		{Scope: "ror", Subject: "datacenter"},
	}
	result := aclstore.OwnerrefsToFilter(refs)

	match := result["$match"].(bson.M)
	or := match["$or"].(bson.A)
	assert.Len(t, or, 3)
	// All are scope-level grants
	for _, entry := range or {
		m := entry.(bson.M)
		_, hasScope := m["rormeta.ownerref.scope"]
		assert.True(t, hasScope)
	}
}

func TestOwnerrefsToFilter_SingleRef(t *testing.T) {
	refs := []acl.Ownerref{
		{Scope: "project", Subject: "proj-42"},
	}
	result := aclstore.OwnerrefsToFilter(refs)

	match := result["$match"].(bson.M)
	or := match["$or"].(bson.A)
	assert.Len(t, or, 1)

	inClause := or[0].(bson.M)["rormeta.ownerref"].(bson.M)["$in"].(bson.A)
	assert.Len(t, inClause, 1)
	entry := inClause[0].(bson.D)
	assert.Equal(t, "project", entry[0].Value)
	assert.Equal(t, "proj-42", entry[1].Value)
}

func TestOwnerrefsToFilter_ManyRefs(t *testing.T) {
	var refs []acl.Ownerref
	for i := range 100 {
		refs = append(refs, acl.Ownerref{
			Scope:   "cluster",
			Subject: aclscope.Subject(fmt.Sprintf("cluster-%03d", i)),
		})
	}
	result := aclstore.OwnerrefsToFilter(refs)

	match := result["$match"].(bson.M)
	or := match["$or"].(bson.A)
	assert.Len(t, or, 1)

	inClause := or[0].(bson.M)["rormeta.ownerref"].(bson.M)["$in"].(bson.A)
	assert.Len(t, inClause, 100)
}

func TestOwnerrefsToFilter_DuplicateRefsHandled(t *testing.T) {
	// Resolver deduplicates, but test that filter handles duplicates gracefully
	refs := []acl.Ownerref{
		{Scope: "cluster", Subject: "cluster-1"},
		{Scope: "cluster", Subject: "cluster-1"}, // duplicate
	}
	result := aclstore.OwnerrefsToFilter(refs)

	match := result["$match"].(bson.M)
	or := match["$or"].(bson.A)
	inClause := or[0].(bson.M)["rormeta.ownerref"].(bson.M)["$in"].(bson.A)
	// Both appear (dedup is resolver's job), but query is still valid
	assert.Len(t, inClause, 2)
}

func TestOwnerrefsToFilter_RorScopeWithMixedGlobalAndSpecific(t *testing.T) {
	// ror-scope grants for cluster + specific project + specific datacenter
	refs := []acl.Ownerref{
		{Scope: "ror", Subject: "cluster"},
		{Scope: "project", Subject: "proj-1"},
		{Scope: "datacenter", Subject: "dc-1"},
		{Scope: "cluster", Subject: "cluster-1"}, // covered by ror grant
	}
	result := aclstore.OwnerrefsToFilter(refs)

	match := result["$match"].(bson.M)
	or := match["$or"].(bson.A)
	assert.Len(t, or, 2) // ror scope-level + $in for project+datacenter

	// First: ror scope-level
	assert.Equal(t, bson.M{"rormeta.ownerref.scope": "cluster"}, or[0])

	// Second: $in
	inClause := or[1].(bson.M)["rormeta.ownerref"].(bson.M)["$in"].(bson.A)
	assert.Len(t, inClause, 2) // proj-1 + dc-1, cluster-1 excluded
}

// --- ClusterIdentityFilter ---

func TestClusterIdentityFilter(t *testing.T) {
	result := aclstore.ClusterIdentityFilter("my-cluster-id")

	expected := bson.M{
		"$match": bson.M{
			"rormeta.ownerref.scope":   "cluster",
			"rormeta.ownerref.subject": "my-cluster-id",
		},
	}
	assert.Equal(t, expected, result)
}

func TestClusterIdentityFilter_EmptyID(t *testing.T) {
	result := aclstore.ClusterIdentityFilter("")

	match := result["$match"].(bson.M)
	assert.Equal(t, "", match["rormeta.ownerref.subject"])
}

// --- DenyAllFilter ---

func TestDenyAllFilter_HasImpossibleMatch(t *testing.T) {
	match := aclstore.DenyAllFilter["$match"].(bson.M)
	assert.Equal(t, "NA-UNKNOWN", match["rormeta.ownerref.scope"])
	assert.Equal(t, "NA-UNKNOWN", match["rormeta.ownerref.subject"])
}
