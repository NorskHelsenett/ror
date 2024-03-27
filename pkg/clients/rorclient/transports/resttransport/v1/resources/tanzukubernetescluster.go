package resources

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
)

func (c *V1Client) GetTanzuKubernetesClusterByUid(uid, ownerSubject string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceTanzuKubernetesCluster, error) {
	kind := "TanzuKubernetesCluster"
	apiversion := "run.tanzu.vmware.com/v1alpha2"
	var result *apiresourcecontracts.ResourceTanzuKubernetesCluster
	err := c.Client.GetJSON(c.basePath+"/uid/"+uid+"?ownerScope="+string(scope)+"&ownerSubject="+ownerSubject+"&apiversion="+apiversion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
