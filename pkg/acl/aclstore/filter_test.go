package aclstore_test

import (
	"fmt"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/acl"
	"github.com/NorskHelsenett/ror/pkg/acl/aclstore"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rordefs"

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
			"rormeta.ownerref.scope":   "KubernetesCluster",
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

// --- ResourceTypeFilter ---

// testProtectedTypes is the registry used by all ResourceTypeFilter tests.
// Tests swap this into ProtectedResourceTypes so they are independent of
// the production registry.
var testProtectedTypes = map[aclmodels.Capability][]string{
	aclmodels.CapRorVulnerability: {
		rordefs.ResourceVulnerabilityReport.Kind,
		rordefs.ResourceExposedSecretReport.Kind,
		rordefs.ResourceConfigAuditReport.Kind,
		rordefs.ResourceRbacAssessmentReport.Kind,
	},
	aclmodels.CapRorConfig: {
		rordefs.ResourceConfiguration.Kind,
	},
}

// withTestRegistry swaps ProtectedResourceTypes for the duration of a test.
func withTestRegistry(t *testing.T) {
	t.Helper()
	orig := aclstore.ProtectedResourceTypes
	aclstore.ProtectedResourceTypes = testProtectedTypes
	t.Cleanup(func() { aclstore.ProtectedResourceTypes = orig })
}

func TestResourceTypeFilter_AllCapabilities_NoRestriction(t *testing.T) {
	withTestRegistry(t)
	access := []aclmodels.AccessTypeV3{
		"ror:vulnerability:read",
		"ror:config:read",
	}
	result := aclstore.ResourceTypeFilter(access)
	assert.Equal(t, bson.M{}, result)
}

func TestResourceTypeFilter_NoCapabilities_AllProtectedExcluded(t *testing.T) {
	withTestRegistry(t)
	result := aclstore.ResourceTypeFilter([]aclmodels.AccessTypeV3{})

	match, ok := result["$match"].(bson.M)
	assert.True(t, ok)
	nin := match["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.Contains(t, nin, "VulnerabilityReport")
	assert.Contains(t, nin, "ExposedSecretReport")
	assert.Contains(t, nin, "ConfigAuditReport")
	assert.Contains(t, nin, "RbacAssessmentReport")
	assert.Contains(t, nin, "Configuration")
}

func TestResourceTypeFilter_NilAccess_AllProtectedExcluded(t *testing.T) {
	withTestRegistry(t)
	result := aclstore.ResourceTypeFilter(nil)

	match, ok := result["$match"].(bson.M)
	assert.True(t, ok)
	nin := match["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.Len(t, nin, 5)
}

func TestResourceTypeFilter_StandardReadOnly_ExcludesProtected(t *testing.T) {
	withTestRegistry(t)
	// ror:read is not a protected prefix — all protected kinds excluded
	access := []aclmodels.AccessTypeV3{"ror:read", "ror:write"}
	result := aclstore.ResourceTypeFilter(access)

	match := result["$match"].(bson.M)
	nin := match["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.Contains(t, nin, "VulnerabilityReport")
	assert.Contains(t, nin, "Configuration")
}

func TestResourceTypeFilter_VulnReadOnly_ExcludesConfig(t *testing.T) {
	withTestRegistry(t)
	access := []aclmodels.AccessTypeV3{"ror:vulnerability:read"}
	result := aclstore.ResourceTypeFilter(access)

	match := result["$match"].(bson.M)
	nin := match["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.NotContains(t, nin, "VulnerabilityReport")
	assert.Contains(t, nin, "Configuration")
}

func TestResourceTypeFilter_ConfigReadOnly_ExcludesVuln(t *testing.T) {
	withTestRegistry(t)
	access := []aclmodels.AccessTypeV3{"ror:config:read"}
	result := aclstore.ResourceTypeFilter(access)

	match := result["$match"].(bson.M)
	nin := match["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.Contains(t, nin, "VulnerabilityReport")
	assert.NotContains(t, nin, "Configuration")
}

func TestResourceTypeFilter_SortedOutput(t *testing.T) {
	withTestRegistry(t)
	result := aclstore.ResourceTypeFilter([]aclmodels.AccessTypeV3{})

	match := result["$match"].(bson.M)
	nin := match["typemeta.kind"].(bson.M)["$nin"].([]string)
	for i := 1; i < len(nin); i++ {
		assert.LessOrEqual(t, nin[i-1], nin[i], "excluded kinds must be sorted")
	}
}

func TestResourceTypeFilter_UnrelatedCapabilities_StillExcludes(t *testing.T) {
	withTestRegistry(t)
	access := []aclmodels.AccessTypeV3{"ror:read", "kubernetes:logon", "kubernetes:admin"}
	result := aclstore.ResourceTypeFilter(access)

	match := result["$match"].(bson.M)
	nin := match["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.Contains(t, nin, "VulnerabilityReport")
	assert.Contains(t, nin, "Configuration")
}

func TestResourceTypeFilter_DuplicateCapabilities_StillWorks(t *testing.T) {
	withTestRegistry(t)
	access := []aclmodels.AccessTypeV3{
		"ror:vulnerability:read",
		"ror:vulnerability:read",
		"ror:config:read",
	}
	result := aclstore.ResourceTypeFilter(access)
	assert.Equal(t, bson.M{}, result)
}

func TestResourceTypeFilter_NoMatchStage_WhenNoExclusion(t *testing.T) {
	withTestRegistry(t)
	access := []aclmodels.AccessTypeV3{"ror:vulnerability:read", "ror:config:read"}
	result := aclstore.ResourceTypeFilter(access)
	assert.Equal(t, bson.M{}, result)
	_, hasMatch := result["$match"]
	assert.False(t, hasMatch)
}

func TestResourceTypeFilter_ProtectedRegistryCoverage(t *testing.T) {
	withTestRegistry(t)
	for prefix, kinds := range aclstore.ProtectedResourceTypes {
		assert.NotEmpty(t, kinds, "prefix %s must protect at least one kind", prefix)
		for _, k := range kinds {
			assert.NotEmpty(t, k, "kind must not be empty for prefix %s", prefix)
		}
	}
}

func TestResourceTypeFilter_ResultIsValidPipelineStage(t *testing.T) {
	withTestRegistry(t)
	result := aclstore.ResourceTypeFilter([]aclmodels.AccessTypeV3{"ror:read"})

	if len(result) == 0 {
		return
	}

	match, ok := result["$match"]
	assert.True(t, ok, "non-empty result must have $match key")

	matchDoc, ok := match.(bson.M)
	assert.True(t, ok, "$match must be a bson.M")

	kindFilter, ok := matchDoc["typemeta.kind"]
	assert.True(t, ok, "$match must contain typemeta.kind")

	ninDoc, ok := kindFilter.(bson.M)
	assert.True(t, ok, "kind filter must be a bson.M")

	_, hasNin := ninDoc["$nin"]
	assert.True(t, hasNin, "kind filter must use $nin operator")
}

func TestResourceTypeFilter_ComposesWithOwnerrefsToFilter(t *testing.T) {
	withTestRegistry(t)
	refs := []acl.Ownerref{
		{Scope: "cluster", Subject: "cluster-1"},
	}
	ownerFilter := aclstore.OwnerrefsToFilter(refs)

	typeFilter := aclstore.ResourceTypeFilter([]aclmodels.AccessTypeV3{"ror:read"})

	pipeline := bson.A{}
	if len(ownerFilter) > 0 {
		pipeline = append(pipeline, ownerFilter)
	}
	if len(typeFilter) > 0 {
		pipeline = append(pipeline, typeFilter)
	}

	assert.Len(t, pipeline, 2)

	stage1 := pipeline[0].(bson.M)["$match"].(bson.M)
	stage2 := pipeline[1].(bson.M)["$match"].(bson.M)

	_, hasOwnerref := stage1["$or"]
	assert.True(t, hasOwnerref, "first stage filters on ownerrefs")

	_, hasKind := stage2["typemeta.kind"]
	assert.True(t, hasKind, "second stage filters on kind")
}

func TestResourceTypeFilter_UnrestrictedOwnerPlusTypeFilter(t *testing.T) {
	withTestRegistry(t)
	ownerFilter := aclstore.OwnerrefsToFilter(nil)
	assert.Equal(t, bson.M{}, ownerFilter)

	typeFilter := aclstore.ResourceTypeFilter([]aclmodels.AccessTypeV3{"ror:read"})

	pipeline := bson.A{}
	if len(ownerFilter) > 0 {
		pipeline = append(pipeline, ownerFilter)
	}
	if len(typeFilter) > 0 {
		pipeline = append(pipeline, typeFilter)
	}

	assert.Len(t, pipeline, 1, "only type filter stage when ownerrefs unrestricted")
}

func TestResourceTypeFilter_DenyAllOwnerPlusTypeFilter(t *testing.T) {
	withTestRegistry(t)
	ownerFilter := aclstore.OwnerrefsToFilter([]acl.Ownerref{})
	assert.Equal(t, aclstore.DenyAllFilter, ownerFilter)

	// All read capabilities present — type filter is no-op
	access := []aclmodels.AccessTypeV3{"ror:vulnerability:read", "ror:config:read"}
	typeFilter := aclstore.ResourceTypeFilter(access)
	assert.Equal(t, bson.M{}, typeFilter)

	pipeline := bson.A{ownerFilter}
	if len(typeFilter) > 0 {
		pipeline = append(pipeline, typeFilter)
	}

	assert.Len(t, pipeline, 1, "only deny-all stage, type filter is no-op")
}

func TestResourceTypeFilter_ExcludesExactKinds(t *testing.T) {
	withTestRegistry(t)
	result := aclstore.ResourceTypeFilter([]aclmodels.AccessTypeV3{})

	match := result["$match"].(bson.M)
	nin := match["typemeta.kind"].(bson.M)["$nin"].([]string)

	var allProtected []string
	for _, kinds := range testProtectedTypes {
		allProtected = append(allProtected, kinds...)
	}

	assert.ElementsMatch(t, allProtected, nin)
}

func TestResourceTypeFilter_ClusterIdentityPlusTypeFilter(t *testing.T) {
	withTestRegistry(t)
	clusterFilter := aclstore.ClusterIdentityFilter("my-cluster")

	typeFilter := aclstore.ResourceTypeFilter([]aclmodels.AccessTypeV3{"ror:read"})

	pipeline := bson.A{clusterFilter}
	if len(typeFilter) > 0 {
		pipeline = append(pipeline, typeFilter)
	}

	assert.Len(t, pipeline, 2)

	s1 := pipeline[0].(bson.M)["$match"].(bson.M)
	assert.Equal(t, "KubernetesCluster", s1["rormeta.ownerref.scope"])
	assert.Equal(t, "my-cluster", s1["rormeta.ownerref.subject"])

	s2 := pipeline[1].(bson.M)["$match"].(bson.M)
	nin := s2["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.Contains(t, nin, "VulnerabilityReport")
	assert.Contains(t, nin, "Configuration")
}

func TestResourceTypeFilter_EmptyRegistry_NoFilter(t *testing.T) {
	orig := aclstore.ProtectedResourceTypes
	aclstore.ProtectedResourceTypes = map[aclmodels.Capability][]string{}
	t.Cleanup(func() { aclstore.ProtectedResourceTypes = orig })

	result := aclstore.ResourceTypeFilter([]aclmodels.AccessTypeV3{})
	assert.Equal(t, bson.M{}, result)
}

func TestResourceTypeFilter_SinglePrefixMultipleKinds(t *testing.T) {
	orig := aclstore.ProtectedResourceTypes
	aclstore.ProtectedResourceTypes = map[aclmodels.Capability][]string{
		"ror:secret": {"SecretA", "SecretB", "SecretC"},
	}
	t.Cleanup(func() { aclstore.ProtectedResourceTypes = orig })

	// Without capability — all 3 excluded
	result := aclstore.ResourceTypeFilter([]aclmodels.AccessTypeV3{})
	nin := result["$match"].(bson.M)["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.ElementsMatch(t, []string{"SecretA", "SecretB", "SecretC"}, nin)

	// With capability — none excluded
	result = aclstore.ResourceTypeFilter([]aclmodels.AccessTypeV3{"ror:secret:read"})
	assert.Equal(t, bson.M{}, result)
}

// --- ResourceTypeWriteFilter ---

func TestResourceTypeWriteFilter_AllWriteCapabilities_NoRestriction(t *testing.T) {
	withTestRegistry(t)
	access := []aclmodels.AccessTypeV3{"ror:vulnerability:write", "ror:config:write"}
	result := aclstore.ResourceTypeWriteFilter(access)
	assert.Equal(t, bson.M{}, result)
}

func TestResourceTypeWriteFilter_NoCapabilities_AllExcluded(t *testing.T) {
	withTestRegistry(t)
	result := aclstore.ResourceTypeWriteFilter([]aclmodels.AccessTypeV3{})

	match := result["$match"].(bson.M)
	nin := match["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.Contains(t, nin, "VulnerabilityReport")
	assert.Contains(t, nin, "Configuration")
}

func TestResourceTypeWriteFilter_ReadDoesNotGrantWrite(t *testing.T) {
	withTestRegistry(t)
	access := []aclmodels.AccessTypeV3{"ror:vulnerability:read", "ror:config:read"}
	result := aclstore.ResourceTypeWriteFilter(access)

	match := result["$match"].(bson.M)
	nin := match["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.Contains(t, nin, "VulnerabilityReport")
	assert.Contains(t, nin, "Configuration")
}

func TestResourceTypeWriteFilter_WriteDoesNotGrantRead(t *testing.T) {
	withTestRegistry(t)
	access := []aclmodels.AccessTypeV3{"ror:vulnerability:write", "ror:config:write"}
	result := aclstore.ResourceTypeFilter(access) // read filter

	match := result["$match"].(bson.M)
	nin := match["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.Contains(t, nin, "VulnerabilityReport")
	assert.Contains(t, nin, "Configuration")
}

func TestResourceTypeWriteFilter_PartialWrite(t *testing.T) {
	withTestRegistry(t)
	access := []aclmodels.AccessTypeV3{"ror:vulnerability:write"}
	result := aclstore.ResourceTypeWriteFilter(access)

	match := result["$match"].(bson.M)
	nin := match["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.NotContains(t, nin, "VulnerabilityReport")
	assert.Contains(t, nin, "Configuration")
}

func TestResourceTypeWriteFilter_SortedOutput(t *testing.T) {
	withTestRegistry(t)
	result := aclstore.ResourceTypeWriteFilter([]aclmodels.AccessTypeV3{})

	match := result["$match"].(bson.M)
	nin := match["typemeta.kind"].(bson.M)["$nin"].([]string)
	for i := 1; i < len(nin); i++ {
		assert.LessOrEqual(t, nin[i-1], nin[i], "excluded kinds must be sorted")
	}
}

func TestResourceTypeWriteFilter_SameRegistryAsReadFilter(t *testing.T) {
	withTestRegistry(t)
	readResult := aclstore.ResourceTypeFilter([]aclmodels.AccessTypeV3{})
	writeResult := aclstore.ResourceTypeWriteFilter([]aclmodels.AccessTypeV3{})

	readNin := readResult["$match"].(bson.M)["typemeta.kind"].(bson.M)["$nin"].([]string)
	writeNin := writeResult["$match"].(bson.M)["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.ElementsMatch(t, readNin, writeNin)
}

func TestResourceTypeFilter_ReadAndWriteIndependent(t *testing.T) {
	withTestRegistry(t)
	access := []aclmodels.AccessTypeV3{"ror:vulnerability:read", "ror:config:write"}

	readFilter := aclstore.ResourceTypeFilter(access)
	writeFilter := aclstore.ResourceTypeWriteFilter(access)

	// Read: vuln ok, config excluded
	readNin := readFilter["$match"].(bson.M)["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.NotContains(t, readNin, "VulnerabilityReport")
	assert.Contains(t, readNin, "Configuration")

	// Write: config ok, vuln excluded
	writeNin := writeFilter["$match"].(bson.M)["typemeta.kind"].(bson.M)["$nin"].([]string)
	assert.Contains(t, writeNin, "VulnerabilityReport")
	assert.NotContains(t, writeNin, "Configuration")
}
