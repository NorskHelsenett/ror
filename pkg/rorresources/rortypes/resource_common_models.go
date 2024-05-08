// package delivers apicontracts for resources
package rortypes

import (
	"errors"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResourceAction string

const (
	K8sActionAdd    ResourceAction = "Add"
	K8sActionDelete ResourceAction = "Delete"
	K8sActionUpdate ResourceAction = "Update"
)

var (
	ErrInvalidScope   = errors.New("invalid scope")
	ErrInvalidSubject = errors.New("invalid subject")
)

// Commonresource defines the minimum resource definition.
type CommonResource struct {
	metav1.TypeMeta `json:",inline"`
	Metadata        metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	RorMeta         ResourceRorMeta   `json:"rormeta"`
}

// ResourceRorMeta represents the metadata stored by ror
type ResourceRorMeta struct {
	Version      string                    `json:"version,omitempty"`
	LastReported string                    `json:"lastReported,omitempty"`
	Internal     bool                      `json:"internal,omitempty"`
	Hash         string                    `json:"hash,omitempty"`
	Ownerref     RorResourceOwnerReference `json:"ownerref,omitempty"`
	Action       ResourceAction            `json:"action,omitempty"`
}

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

// aclmodels.ErrInvalidScope is returned when the scope is invalid
