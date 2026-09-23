package aclcaps_test

import (
	"testing"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclcaps"
	"github.com/stretchr/testify/assert"
)

// TestCapabilityConstsInRegistry guards against drift between the exported Cap*
// constants and the capability Registry: every declared capability must resolve
// to a registered node that can carry verbs.
func TestCapabilityConstsInRegistry(t *testing.T) {
	caps := []aclcaps.Capability{
		aclcaps.CapRor,
		aclcaps.CapRorMetadata,
		aclcaps.CapRorVulnerability,
		aclcaps.CapRorConfig,
		aclcaps.CapKubernetes,
		aclcaps.CapKubernetesArgocd,
		aclcaps.CapKubernetesArgocdProject,
		aclcaps.CapKubernetesGrafana,
		aclcaps.CapVirtualmachine,
	}
	for _, c := range caps {
		assert.Truef(t, aclcaps.ValidCapability(c), "capability %q not in registry", c)
	}
}

func TestValidate_Valid(t *testing.T) {
	valid := []aclcaps.AccessTypeV3{
		"ror:read",
		"ror:config:write",
		"kubernetes:argocd:project:admin",
		"virtualmachine:delete",
		"monitoring:read",
		"monitoring:write",
		"dns:read",
		"dns:write",
	}
	for _, a := range valid {
		assert.NoErrorf(t, aclcaps.Validate(a), "expected %q to be valid", a)
	}
}

func TestValidate_Invalid(t *testing.T) {
	invalid := []aclcaps.AccessTypeV3{
		"read",
		"foo:bar",
		"ror:metadata:read",
		"ror:execute",
		"monitoring:admin",
		"dns:delete",
		"",
	}
	for _, a := range invalid {
		assert.Errorf(t, aclcaps.Validate(a), "expected %q to be invalid", a)
	}
}

func TestRegistryVerbsAreKnown(t *testing.T) {
	var walk func(n *aclcaps.Namespace)
	walk = func(n *aclcaps.Namespace) {
		if n == nil {
			return
		}
		for v := range n.Verbs {
			_, ok := aclcaps.AllVerbs[v]
			assert.Truef(t, ok, "verb %q not in AllVerbs", v)
		}
		for _, c := range n.Children {
			walk(c)
		}
	}
	walk(aclcaps.Registry)
}
