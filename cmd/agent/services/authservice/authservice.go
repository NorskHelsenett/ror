// authservice implements authorization helpers for the agent
package authservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/spf13/viper"
)

// creaters a ownerref object for the agent
func CreateOwnerref() apiresourcecontracts.ResourceOwnerReference {
	return apiresourcecontracts.ResourceOwnerReference{
		Scope:   aclmodels.Acl2ScopeCluster,
		Subject: viper.GetString(configconsts.CLUSTER_ID),
	}
}
