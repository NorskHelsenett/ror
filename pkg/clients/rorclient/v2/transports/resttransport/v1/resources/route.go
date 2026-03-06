package resources

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

func (c *V1Client) GetRoutesByOwner(ctx context.Context, owner string, scope aclmodels.Acl2Scope) ([]apiresourcecontracts.ResourceRoute, error) {
	kind := "Route"
	apiVersion := "general.ror.internal/v1alpha1"
	var result []apiresourcecontracts.ResourceRoute
	err := c.Client.GetJSON(ctx, c.basePath+"?ownerScope="+string(scope)+"&ownerSubject="+string(owner)+"&apiversion="+apiVersion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
