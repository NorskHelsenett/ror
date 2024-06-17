package resources

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
)

func (c *V1Client) GetSlackMessageByUid(uid, owner string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceSlackMessage, error) {
	kind := "SlackMessage"
	apiVersion := "general.ror.internal/v1alpha1"
	var result apiresourcecontracts.ResourceSlackMessage
	err := c.Client.GetJSON(c.basePath+"/uid/"+uid+"?ownerScope="+string(scope)+"&ownerSubject="+string(owner)+"&apiversion="+apiVersion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *V1Client) CreateSlackMessage(sm *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceSlackMessage, error) {
	var result apiresourcecontracts.ResourceSlackMessage
	err := c.Client.PostJSON(c.basePath, sm, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *V1Client) UpdateSlackMessageByUid(sm *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceSlackMessage, error) {
	var result apiresourcecontracts.ResourceSlackMessage
	err := c.Client.PutJSON(c.basePath+"/uid/"+sm.Uid, sm, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
