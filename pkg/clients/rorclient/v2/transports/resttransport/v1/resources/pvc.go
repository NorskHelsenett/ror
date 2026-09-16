package resources

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

func (c *V1Client) GetPVCByUid(ctx context.Context, uid string, ownerSubject string, scope aclscope.Scope) (*apiresourcecontracts.ResourcePersistentVolumeClaim, error) {
	kind := "PersistentVolumeClaim"
	apiVersion := "v1"
	var result *apiresourcecontracts.ResourcePersistentVolumeClaim
	err := c.Client.GetJSON(ctx, c.basePath+"/uid/"+uid+"?ownerScope="+string(scope)+"&ownerSubject="+ownerSubject+"&apiversion="+apiVersion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
