package resources

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

func (c *V1Client) GetSlackMessageByUid(ctx context.Context, uid, owner string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceSlackMessage, error) {
	kind := "SlackMessage"
	apiVersion := "general.ror.internal/v1alpha1"
	var result apiresourcecontracts.ResourceSlackMessage
	err := c.Client.GetJSON(ctx, c.basePath+"/uid/"+uid+"?ownerScope="+string(scope)+"&ownerSubject="+string(owner)+"&apiversion="+apiVersion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *V1Client) CreateSlackMessage(ctx context.Context, sm *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceSlackMessage, error) {
	var result apiresourcecontracts.ResourceSlackMessage
	err := c.Client.PostJSON(ctx, c.basePath, sm, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *V1Client) UpdateSlackMessageByUid(ctx context.Context, sm *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceSlackMessage, error) {
	var result apiresourcecontracts.ResourceSlackMessage
	err := c.Client.PutJSON(ctx, c.basePath+"/uid/"+sm.Uid, sm, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
