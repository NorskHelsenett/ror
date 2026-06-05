package aclmodels_test

import (
	"testing"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/stretchr/testify/assert"
)

func TestValidateAccess_ValidTypes(t *testing.T) {
	valid := []aclmodels.AccessTypeV3{
		"ror:read",
		"ror:write",
		"ror:owner",
		"ror:metadata:write",
		"ror:vulnerability:read",
		"ror:vulnerability:write",
		"ror:config:read",
		"ror:config:write",
		"kubernetes:logon",
		"kubernetes:admin",
		"kubernetes:readonly",
		"kubernetes:argocd:admin",
		"kubernetes:argocd:project:admin",
		"kubernetes:grafana:admin",
		"resource:Deployment:read",
		"resource:Pod:write",
		"resource:*:read",
		"resource:*:delete",
		"resource:VulnerabilityReport:read",
		"virtualmachine:delete",
	}
	for _, v := range valid {
		t.Run(string(v), func(t *testing.T) {
			err := aclmodels.ValidateAccess(v)
			assert.NoError(t, err)
		})
	}
}

func TestValidateAccess_InvalidTypes(t *testing.T) {
	tests := []struct {
		name   string
		access aclmodels.AccessTypeV3
	}{
		{"single segment", "read"},
		{"unknown system", "foo:bar"},
		{"wrong verb at path", "ror:metadata:read"},
		{"unknown verb", "ror:execute"},
		{"verb not allowed at level", "resource:Deployment:admin"},
		{"unknown sub path", "ror:unknown:write"},
		{"empty string", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := aclmodels.ValidateAccess(tt.access)
			assert.Error(t, err)
		})
	}
}

func TestValidScope_Systems(t *testing.T) {
	assert.NoError(t, aclmodels.ValidScope("ror"))
	assert.NoError(t, aclmodels.ValidScope("all"))
	assert.NoError(t, aclmodels.ValidScope("spam"))
	assert.NoError(t, aclmodels.ValidScope("alarm"))
}

func TestValidScope_Kinds(t *testing.T) {
	assert.NoError(t, aclmodels.ValidScope("KubernetesCluster"))
	assert.NoError(t, aclmodels.ValidScope("Deployment"))
	assert.NoError(t, aclmodels.ValidScope("Project"))
	assert.NoError(t, aclmodels.ValidScope("Datacenter"))
	assert.NoError(t, aclmodels.ValidScope("VirtualMachine"))
	assert.NoError(t, aclmodels.ValidScope("Machine"))
	assert.NoError(t, aclmodels.ValidScope("BackupJob"))
}

func TestValidScope_Invalid(t *testing.T) {
	assert.Error(t, aclmodels.ValidScope("unknown"))
	assert.Error(t, aclmodels.ValidScope(""))
	assert.Error(t, aclmodels.ValidScope("cluster")) // old V2 scope, not valid in V3
	assert.Error(t, aclmodels.ValidScope("notakind"))
}

func TestHasAccess(t *testing.T) {
	entry := aclmodels.AclV3ListItem{
		Access: []aclmodels.AccessTypeV3{"ror:read", "kubernetes:logon"},
	}
	assert.True(t, entry.HasAccess("ror:read"))
	assert.True(t, entry.HasAccess("kubernetes:logon"))
	assert.False(t, entry.HasAccess("ror:write"))
	assert.False(t, entry.HasAccess("kubernetes:admin"))
}

func TestMergeAccess(t *testing.T) {
	a := []aclmodels.AccessTypeV3{"ror:read", "ror:write"}
	b := []aclmodels.AccessTypeV3{"ror:read", "kubernetes:logon"}
	result := aclmodels.MergeAccess(a, b)
	assert.Len(t, result, 3)
	assert.Contains(t, result, aclmodels.AccessTypeV3("ror:read"))
	assert.Contains(t, result, aclmodels.AccessTypeV3("ror:write"))
	assert.Contains(t, result, aclmodels.AccessTypeV3("kubernetes:logon"))
}

func TestMergeAccess_Empty(t *testing.T) {
	result := aclmodels.MergeAccess(nil, nil)
	assert.Empty(t, result)
}

func TestMatchPrefix(t *testing.T) {
	access := []aclmodels.AccessTypeV3{
		"ror:read",
		"resource:Deployment:read",
		"resource:Pod:write",
		"kubernetes:logon",
	}
	result := aclmodels.MatchPrefix(access, "resource:")
	assert.Len(t, result, 2)
	assert.Contains(t, result, aclmodels.AccessTypeV3("resource:Deployment:read"))
	assert.Contains(t, result, aclmodels.AccessTypeV3("resource:Pod:write"))
}

func TestCanAccessKind(t *testing.T) {
	access := []aclmodels.AccessTypeV3{
		"resource:Deployment:read",
		"resource:Pod:read",
	}
	assert.True(t, aclmodels.CanAccessKind(access, "Deployment", aclmodels.VerbRead))
	assert.True(t, aclmodels.CanAccessKind(access, "Pod", aclmodels.VerbRead))
	assert.False(t, aclmodels.CanAccessKind(access, "Service", aclmodels.VerbRead))
	assert.False(t, aclmodels.CanAccessKind(access, "Deployment", aclmodels.VerbWrite))
}

func TestCanAccessKind_Wildcard(t *testing.T) {
	access := []aclmodels.AccessTypeV3{"resource:*:read"}
	assert.True(t, aclmodels.CanAccessKind(access, "Deployment", aclmodels.VerbRead))
	assert.True(t, aclmodels.CanAccessKind(access, "Service", aclmodels.VerbRead))
	assert.True(t, aclmodels.CanAccessKind(access, "VulnerabilityReport", aclmodels.VerbRead))
	assert.False(t, aclmodels.CanAccessKind(access, "Deployment", aclmodels.VerbWrite))
}

func TestAllowedKinds(t *testing.T) {
	access := []aclmodels.AccessTypeV3{
		"resource:Deployment:read",
		"resource:Pod:read",
		"resource:Service:write",
	}
	kinds := aclmodels.AllowedKinds(access, aclmodels.VerbRead)
	assert.Len(t, kinds, 2)
	assert.Contains(t, kinds, "Deployment")
	assert.Contains(t, kinds, "Pod")
}

func TestAllowedKinds_Wildcard(t *testing.T) {
	access := []aclmodels.AccessTypeV3{"resource:*:read", "resource:Deployment:read"}
	kinds := aclmodels.AllowedKinds(access, aclmodels.VerbRead)
	assert.Nil(t, kinds) // nil means all kinds allowed
}

func TestAllowedKinds_NoAccess(t *testing.T) {
	access := []aclmodels.AccessTypeV3{"ror:read"}
	kinds := aclmodels.AllowedKinds(access, aclmodels.VerbRead)
	assert.NotNil(t, kinds)
	assert.Empty(t, kinds)
}

func TestCompileAccess(t *testing.T) {
	entries := []aclmodels.AclV3ListItem{
		{
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read"},
		},
		{
			Scope:   "KubernetesCluster",
			Subject: "cluster-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read", "kubernetes:logon"},
		},
		{
			Scope:   "KubernetesCluster",
			Subject: "cluster-2",
			Access:  []aclmodels.AccessTypeV3{"ror:write"},
		},
	}
	result := aclmodels.CompileAccess(entries, "KubernetesCluster", "cluster-1")
	assert.Len(t, result, 2)
	assert.Contains(t, result, aclmodels.AccessTypeV3("ror:read"))
	assert.Contains(t, result, aclmodels.AccessTypeV3("kubernetes:logon"))
}

func TestValidateACLEntry_Valid(t *testing.T) {
	entry := aclmodels.AclV3ListItem{
		Group:   "dev-team",
		Scope:   "ror",
		Subject: "Global",
		Access:  []aclmodels.AccessTypeV3{"ror:read", "ror:write"},
	}
	assert.NoError(t, aclmodels.ValidateACLEntry(entry))
}

func TestValidateACLEntry_ValidKindScope(t *testing.T) {
	entry := aclmodels.AclV3ListItem{
		Group:   "dev-team",
		Scope:   "KubernetesCluster",
		Subject: "prod-cluster-1",
		Access:  []aclmodels.AccessTypeV3{"ror:read", "resource:Deployment:read"},
	}
	assert.NoError(t, aclmodels.ValidateACLEntry(entry))
}

func TestValidateACLEntry_InvalidScope(t *testing.T) {
	entry := aclmodels.AclV3ListItem{
		Group:   "dev-team",
		Scope:   "invalid-scope",
		Subject: "something",
		Access:  []aclmodels.AccessTypeV3{"ror:read"},
	}
	assert.Error(t, aclmodels.ValidateACLEntry(entry))
}

func TestValidateACLEntry_InvalidAccess(t *testing.T) {
	entry := aclmodels.AclV3ListItem{
		Group:   "dev-team",
		Scope:   "ror",
		Subject: "Global",
		Access:  []aclmodels.AccessTypeV3{"ror:read", "invalid:nonsense"},
	}
	assert.Error(t, aclmodels.ValidateACLEntry(entry))
}

func TestAccessTypeV3Constants(t *testing.T) {
	// Verify all constants pass validation
	constants := []aclmodels.AccessTypeV3{
		aclmodels.AccessRorRead,
		aclmodels.AccessRorWrite,
		aclmodels.AccessRorOwner,
		aclmodels.AccessRorMetadataWrite,
		aclmodels.AccessRorVulnerabilityRead,
		aclmodels.AccessRorVulnerabilityWrite,
		aclmodels.AccessRorConfigRead,
		aclmodels.AccessRorConfigWrite,
		aclmodels.AccessKubernetesLogon,
		aclmodels.AccessKubernetesAdmin,
		aclmodels.AccessKubernetesReadonly,
		aclmodels.AccessKubernetesArgocdAdmin,
		aclmodels.AccessKubernetesArgocdProjectAdmin,
		aclmodels.AccessKubernetesGrafanaAdmin,
		aclmodels.AccessVirtualmachineDelete,
	}
	for _, c := range constants {
		t.Run(string(c), func(t *testing.T) {
			assert.NoError(t, aclmodels.ValidateAccess(c))
		})
	}
}

func TestCapabilityWithVerb(t *testing.T) {
	assert.Equal(t, aclmodels.AccessTypeV3("ror:vulnerability:read"), aclmodels.CapRorVulnerability.WithVerb(aclmodels.VerbRead))
	assert.Equal(t, aclmodels.AccessTypeV3("kubernetes:argocd:admin"), aclmodels.CapKubernetesArgocd.WithVerb(aclmodels.VerbAdmin))
	assert.Equal(t, aclmodels.AccessTypeV3("ror:config:write"), aclmodels.CapRorConfig.WithVerb(aclmodels.VerbWrite))
}

func TestAccessTypeV3Parse(t *testing.T) {
	cap, verb := aclmodels.AccessRorVulnerabilityRead.Parse()
	assert.Equal(t, aclmodels.CapRorVulnerability, cap)
	assert.Equal(t, aclmodels.VerbRead, verb)

	cap, verb = aclmodels.AccessKubernetesArgocdProjectAdmin.Parse()
	assert.Equal(t, aclmodels.CapKubernetesArgocdProject, cap)
	assert.Equal(t, aclmodels.VerbAdmin, verb)

	cap, verb = aclmodels.AccessRorRead.Parse()
	assert.Equal(t, aclmodels.CapRor, cap)
	assert.Equal(t, aclmodels.VerbRead, verb)
}

func TestParseWithVerbRoundtrip(t *testing.T) {
	original := aclmodels.AccessRorConfigWrite
	cap, verb := original.Parse()
	roundtripped := cap.WithVerb(verb)
	assert.Equal(t, original, roundtripped)
}
