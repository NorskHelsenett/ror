package resources

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

const (
	apiversion = "general.ror.internal/v1alpha1"
	kind       = "ClusterOrder"
)

func (c *V1Client) GetClusterOrderByUid(ctx context.Context, uid string, ownerSubject aclmodels.Acl2Subject, ownerScope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceClusterOrder, error) {
	var result apiresourcecontracts.ResourceClusterOrder
	err := c.Client.GetJSON(ctx, c.basePath+"/uid/"+uid+"?ownerScope="+string(ownerScope)+"&ownerSubject="+string(ownerSubject)+"&apiversion="+apiversion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *V1Client) GetClusterOrders(ctx context.Context, ownerSubject aclmodels.Acl2Subject, ownerScope aclmodels.Acl2Scope) ([]*apiresourcecontracts.ResourceClusterOrder, error) {
	var result []*apiresourcecontracts.ResourceClusterOrder
	err := c.Client.GetJSON(ctx, c.basePath+"?ownerScope="+string(ownerScope)+"&ownerSubject="+string(ownerSubject)+"&apiversion="+apiversion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *V1Client) UpdateClusterOrder(ctx context.Context, updateModel *apiresourcecontracts.ResourceUpdateModel) error {
	var result bool
	err := c.Client.PutJSON(ctx, c.basePath+"/uid/"+updateModel.Uid, updateModel, &result)
	if err != nil {
		return err
	}
	return nil
}
