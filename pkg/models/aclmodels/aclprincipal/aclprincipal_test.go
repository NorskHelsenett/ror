package aclprincipal_test

import (
	"testing"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclprincipal"
	"github.com/stretchr/testify/assert"
)

func TestPrincipalNames(t *testing.T) {
	assert.Equal(t, "abc-123@cluster.ror.system", aclprincipal.Cluster("abc-123"))
	assert.Equal(t, "*@cluster.ror.system", aclprincipal.AllClusters())
	assert.Equal(t, "vulnerability@service.ror.system", aclprincipal.Service("vulnerability"))
	assert.Equal(t, "service-vulnerability@ror.system", aclprincipal.LegacyService("vulnerability"))
	assert.Equal(t, "*@service.ror.system", aclprincipal.AllServices())
	assert.Equal(t, "deployer@team-a.abc-123.sa.ror.system",
		aclprincipal.ServiceAccount("abc-123", "team-a", "deployer"))
	assert.Equal(t, "*@team-a.abc-123.sa.ror.system",
		aclprincipal.AllServiceAccountsInNamespace("abc-123", "team-a"))
	assert.Equal(t, "*@abc-123.sa.ror.system", aclprincipal.AllServiceAccountsInCluster("abc-123"))
	assert.Equal(t, "*@sa.ror.system", aclprincipal.AllServiceAccounts())
}

func TestGroupMemberships(t *testing.T) {
	assert.Equal(t, []string{
		"abc-123@cluster.ror.system",
		"*@cluster.ror.system",
	}, aclprincipal.ClusterGroups("abc-123"))

	assert.Equal(t, []string{
		"nhn@service.ror.system",
		"service-nhn@ror.system",
		"*@service.ror.system",
	}, aclprincipal.ServiceGroups("nhn"))

	assert.Equal(t, []string{
		"deployer@team-a.abc-123.sa.ror.system",
		"*@team-a.abc-123.sa.ror.system",
		"*@abc-123.sa.ror.system",
		"*@sa.ror.system",
	}, aclprincipal.ServiceAccountGroups("abc-123", "team-a", "deployer"))
}

// A ServiceAccount named "all" must not collide with an aggregate group, or
// anyone able to create that ServiceAccount would inherit the namespace grants.
func TestAggregateLocalPartIsNotAValidKubernetesName(t *testing.T) {
	named := aclprincipal.ServiceAccount("abc-123", "team-a", "all")
	aggregate := aclprincipal.AllServiceAccountsInNamespace("abc-123", "team-a")
	assert.NotEqual(t, aggregate, named)
	assert.NotContains(t, "abcdefghijklmnopqrstuvwxyz0123456789-.", aclprincipal.Wildcard)
}

func TestIsReserved(t *testing.T) {
	reserved := []string{
		"abc-123@cluster.ror.system",
		"nhn@service.ror.system",
		"service-nhn@ror.system",
		"deployer@team-a.abc-123.sa.ror.system",
		"*@cluster.ror.system",
		"Someone@Cluster.ROR.System",
	}
	for _, g := range reserved {
		assert.Truef(t, aclprincipal.IsReserved(g), "%q must be reserved", g)
	}

	external := []string{
		"A-T1-SDI-DevOps-Operators@ror.dev",
		"cluster-admin@ror.io",
		"team-a@example.com",
		"plain-group",
		"",
		"notror.system",
	}
	for _, g := range external {
		assert.Falsef(t, aclprincipal.IsReserved(g), "%q must not be reserved", g)
	}
}

func TestSanitizeExternalGroups(t *testing.T) {
	got := aclprincipal.SanitizeExternalGroups([]string{
		"team-a@example.com",
		"abc-123@cluster.ror.system", // spoofed cluster principal
		"A-T1-SDI-DevOps-Operators@ror.dev",
		"*@sa.ror.system", // spoofed aggregate
	})
	assert.Equal(t, []string{"team-a@example.com", "A-T1-SDI-DevOps-Operators@ror.dev"}, got)
}
