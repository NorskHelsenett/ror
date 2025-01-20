package rorresourceowner

import (
	"errors"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

var (
	// aclmodels.ErrInvalidScope is returned when the scope is invalid
	ErrInvalidScope   = errors.New("invalid scope")
	ErrInvalidSubject = errors.New("invalid subject")
)

// The RorResourceOwnerReference or ownereref references the owner og a resource.
// Its used to chek acl and select resources for valid Scopes.
type RorResourceOwnerReference struct {
	Scope   aclmodels.Acl2Scope   `json:"scope"`   // cluster, workspace,...
	Subject aclmodels.Acl2Subject `json:"subject"` // ror id eg clusterId or workspaceName
}

// Validate validates the ResourceOwnerReference
func (r *RorResourceOwnerReference) Validate() (bool, error) {
	if r.Scope == "" {
		return false, ErrInvalidScope
	}
	if r.Subject == "" {
		return false, ErrInvalidSubject
	}
	if !r.Scope.IsValid() {
		return false, ErrInvalidScope
	}
	if !r.Subject.HasValidScope(r.Scope) {
		return false, ErrInvalidScope
	}
	return true, nil
}

func (r RorResourceOwnerReference) String() string {
	return string(r.Scope) + ":" + string(r.Subject)
}

func (r RorResourceOwnerReference) GetQueryParams() map[string]string {
	response := make(map[string]string)
	response["ownerScope"] = string(r.Scope)
	response["ownerSubject"] = string(r.Subject)
	return response
}
