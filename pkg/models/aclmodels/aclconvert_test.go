package aclmodels_test

import (
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/stretchr/testify/assert"
)

func TestV2ToV3_AllFieldsTrue(t *testing.T) {
	v2 := aclmodels.AclV2ListItem{
		Id:      "abc123",
		Version: 2,
		Group:   "dev-team",
		Scope:   "KubernetesCluster",
		Subject: "cluster-1",
		Access: aclmodels.AclV2ListItemAccess{
			Read:   true,
			Create: true,
			Update: true,
			Delete: true,
			Owner:  true,
		},
		Kubernetes: aclmodels.AclV2ListItemKubernetes{
			Logon: true,
		},
		Created:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		IssuedBy: "admin@example.com",
	}

	v3 := aclmodels.V2ToV3(v2)

	assert.Equal(t, "abc123", v3.Id)
	assert.Equal(t, 3, v3.Version)
	assert.Equal(t, "dev-team", v3.Group)
	assert.Equal(t, aclscope.Scope("KubernetesCluster"), v3.Scope)
	assert.Equal(t, aclscope.Subject("cluster-1"), v3.Subject)
	assert.Equal(t, v2.Created, v3.Created)
	assert.Equal(t, "admin@example.com", v3.IssuedBy)

	assert.Len(t, v3.Access, 6)
	assert.Contains(t, v3.Access, aclmodels.AccessTypeV3("ror:read"))
	assert.Contains(t, v3.Access, aclmodels.AccessTypeV3("ror:create"))
	assert.Contains(t, v3.Access, aclmodels.AccessTypeV3("ror:update"))
	assert.Contains(t, v3.Access, aclmodels.AccessTypeV3("ror:delete"))
	assert.Contains(t, v3.Access, aclmodels.AccessTypeV3("ror:owner"))
	assert.Contains(t, v3.Access, aclmodels.AccessTypeV3("kubernetes:logon"))
}

func TestV2ToV3_ReadOnly(t *testing.T) {
	v2 := aclmodels.AclV2ListItem{
		Version: 2,
		Group:   "viewers",
		Scope:   "Project",
		Subject: "proj-1",
		Access:  aclmodels.NewAclV2ListItemAccessReadOnly(),
	}

	v3 := aclmodels.V2ToV3(v2)
	assert.Len(t, v3.Access, 1)
	assert.Contains(t, v3.Access, aclmodels.AccessTypeV3("ror:read"))
}

func TestV2ToV3_NoAccess(t *testing.T) {
	v2 := aclmodels.AclV2ListItem{
		Version: 2,
		Group:   "empty",
		Scope:   "Project",
		Subject: "proj-1",
		Access:  aclmodels.AclV2ListItemAccess{},
	}

	v3 := aclmodels.V2ToV3(v2)
	assert.Empty(t, v3.Access)
}

func TestV3ToV2_MappableCapabilities(t *testing.T) {
	v3 := aclmodels.AclV3ListItem{
		Id:      "xyz789",
		Version: 3,
		Group:   "ops-team",
		Scope:   "KubernetesCluster",
		Subject: "cluster-prod",
		Access: []aclmodels.AccessTypeV3{
			"ror:read",
			"ror:create",
			"ror:update",
			"ror:delete",
			"ror:owner",
			"kubernetes:logon",
		},
		Created:  time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC),
		IssuedBy: "admin@example.com",
	}

	v2 := aclmodels.V3ToV2(v3)

	assert.Equal(t, "xyz789", v2.Id)
	assert.Equal(t, 2, v2.Version)
	assert.Equal(t, "ops-team", v2.Group)
	assert.True(t, v2.Access.Read)
	assert.True(t, v2.Access.Create)
	assert.True(t, v2.Access.Update)
	assert.True(t, v2.Access.Delete)
	assert.True(t, v2.Access.Owner)
	assert.True(t, v2.Kubernetes.Logon)
}

func TestV3ToV2_V3OnlyCapabilitiesDropped(t *testing.T) {
	v3 := aclmodels.AclV3ListItem{
		Version: 3,
		Group:   "dev-team",
		Scope:   "KubernetesCluster",
		Subject: "cluster-1",
		Access: []aclmodels.AccessTypeV3{
			"ror:read",
			"kubernetes:admin",              // V3-only, no V2 equivalent
			"kubernetes:argocd:admin",       // V3-only
			"resource:Deployment:read",      // V3-only
			"ror:vulnerability:read",        // V3-only
		},
	}

	v2 := aclmodels.V3ToV2(v3)

	assert.True(t, v2.Access.Read)
	assert.False(t, v2.Access.Create)
	assert.False(t, v2.Access.Update)
	assert.False(t, v2.Access.Delete)
	assert.False(t, v2.Access.Owner)
	assert.False(t, v2.Kubernetes.Logon)
}

func TestV3ToV2_EmptyAccess(t *testing.T) {
	v3 := aclmodels.AclV3ListItem{
		Version: 3,
		Group:   "empty",
		Scope:   "Project",
		Subject: "proj-1",
		Access:  []aclmodels.AccessTypeV3{},
	}

	v2 := aclmodels.V3ToV2(v3)
	assert.False(t, v2.Access.Read)
	assert.False(t, v2.Access.Create)
}

func TestV2ToV3_Roundtrip(t *testing.T) {
	original := aclmodels.AclV2ListItem{
		Id:      "round1",
		Version: 2,
		Group:   "team-a",
		Scope:   "KubernetesCluster",
		Subject: "cluster-1",
		Access: aclmodels.AclV2ListItemAccess{
			Read:   true,
			Create: false,
			Update: true,
			Delete: false,
			Owner:  false,
		},
		Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: true},
		Created:    time.Date(2025, 3, 15, 0, 0, 0, 0, time.UTC),
		IssuedBy:   "user@example.com",
	}

	// V2 → V3 → V2 should preserve all V2-representable fields
	v3 := aclmodels.V2ToV3(original)
	roundtripped := aclmodels.V3ToV2(v3)

	assert.Equal(t, original.Id, roundtripped.Id)
	assert.Equal(t, original.Group, roundtripped.Group)
	assert.Equal(t, original.Scope, roundtripped.Scope)
	assert.Equal(t, original.Subject, roundtripped.Subject)
	assert.Equal(t, original.Access.Read, roundtripped.Access.Read)
	assert.Equal(t, original.Access.Create, roundtripped.Access.Create)
	assert.Equal(t, original.Access.Update, roundtripped.Access.Update)
	assert.Equal(t, original.Access.Delete, roundtripped.Access.Delete)
	assert.Equal(t, original.Access.Owner, roundtripped.Access.Owner)
	assert.Equal(t, original.Kubernetes.Logon, roundtripped.Kubernetes.Logon)
}
