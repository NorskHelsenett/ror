package resources

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

const (
	apiversion = "general.ror.internal/v1alpha1"
	kind       = "ClusterOrder"
)

func (c *V1Client) GetClusterOrderByUid(uid string, ownerSubject aclscope.Subject, ownerScope aclscope.Scope) (*apiresourcecontracts.ResourceClusterOrder, error) {
	var result apiresourcecontracts.ResourceClusterOrder
	err := c.Client.GetJSON(c.basePath+"/uid/"+uid+"?ownerScope="+string(ownerScope)+"&ownerSubject="+string(ownerSubject)+"&apiversion="+apiversion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *V1Client) GetClusterOrders(ownerSubject aclscope.Subject, ownerScope aclscope.Scope) ([]*apiresourcecontracts.ResourceClusterOrder, error) {
	var result []*apiresourcecontracts.ResourceClusterOrder
	err := c.Client.GetJSON(c.basePath+"?ownerScope="+string(ownerScope)+"&ownerSubject="+string(ownerSubject)+"&apiversion="+apiversion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *V1Client) UpdateClusterOrder(updateModel *apiresourcecontracts.ResourceUpdateModel) error {
	var result bool
	err := c.Client.PutJSON(c.basePath+"/uid/"+updateModel.Uid, updateModel, &result)
	if err != nil {
		return err
	}
	return nil
}
