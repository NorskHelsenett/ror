package resources

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

func (c *V1Client) GetApplicationByUid(ctx context.Context, uid string, ownerSubject string, scope aclscope.Scope) (*apiresourcecontracts.ResourceApplication, error) {
	kind := "Application"
	apiVersion := "argoproj.io/v1alpha1"

	var result *apiresourcecontracts.ResourceApplication
	err := c.Client.GetJSON(ctx, c.basePath+"/uid/"+uid+"?ownerScope="+string(scope)+"&ownerSubject="+ownerSubject+"&apiversion="+apiVersion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
