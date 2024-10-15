// TODO: Remove this file once the rorconfig is implemented
// authservice implements authorization helpers for the agent
package authservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

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

func CreateRorOwnerref() rortypes.RorResourceOwnerReference {
	return rortypes.RorResourceOwnerReference{
		Scope:   aclmodels.Acl2ScopeCluster,
		Subject: aclmodels.Acl2Subject(viper.GetString(configconsts.CLUSTER_ID)),
	}

}
