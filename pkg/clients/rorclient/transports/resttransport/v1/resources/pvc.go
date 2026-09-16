package resources

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

func (c *V1Client) GetPVCByUid(uid string, ownerSubject string, scope aclscope.Scope) (*apiresourcecontracts.ResourcePersistentVolumeClaim, error) {
	kind := "PersistentVolumeClaim"
	apiVersion := "v1"
	var result *apiresourcecontracts.ResourcePersistentVolumeClaim
	err := c.Client.GetJSON(c.basePath+"/uid/"+uid+"?ownerScope="+string(scope)+"&ownerSubject="+ownerSubject+"&apiversion="+apiVersion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
