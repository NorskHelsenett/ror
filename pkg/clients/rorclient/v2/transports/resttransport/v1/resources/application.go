package resources

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

func (c *V1Client) GetApplicationByUid(uid string, ownerSubject string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceApplication, error) {
	kind := "Application"
	apiVersion := "argoproj.io/v1alpha1"

	var result *apiresourcecontracts.ResourceApplication
	err := c.Client.GetJSON(c.basePath+"/uid/"+uid+"?ownerScope="+string(scope)+"&ownerSubject="+ownerSubject+"&apiversion="+apiVersion+"&kind="+kind, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
