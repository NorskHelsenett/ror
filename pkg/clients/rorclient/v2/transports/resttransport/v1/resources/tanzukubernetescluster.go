package resources

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

func (c *V1Client) GetTanzuKubernetesClusterByUid(ctx context.Context, uid string, ownerSubject string, ownerScope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceTanzuKubernetesCluster, error) {
	kind := "TanzuKubernetesCluster"
	apiversion := "run.tanzu.vmware.com/v1alpha2"
	var result apiresourcecontracts.ResourceTanzuKubernetesCluster
	err := c.Client.GetJSON(ctx, c.basePath+"/uid/"+uid+"?ownerScope="+string(ownerScope)+"&ownerSubject="+ownerSubject+"&apiversion="+apiversion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
