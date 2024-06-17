package resources

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
)

func (c *V1Client) CreateNotification(v *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceNotification, error) {
	var result apiresourcecontracts.ResourceNotification
	err := c.Client.PostJSON(c.basePath, v, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *V1Client) GetNotificationByUid(uid string, ownerSubject string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceNotification, error) {
	kind := "Notification"
	apiVersion := "general.ror.internal/v1alpha1"
	var result apiresourcecontracts.ResourceNotification
	err := c.Client.GetJSON(c.basePath+"/uid/"+uid+"?ownerScope="+string(scope)+"&ownerSubject="+string(ownerSubject)+"&apiversion="+apiVersion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
