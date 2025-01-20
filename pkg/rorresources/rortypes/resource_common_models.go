// package delivers apicontracts for resources
package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResourceAction string

const (
	K8sActionAdd    ResourceAction = "Add"
	K8sActionDelete ResourceAction = "Delete"
	K8sActionUpdate ResourceAction = "Update"
)

// Commonresource defines the minimum resource definition.
type CommonResource struct {
	metav1.TypeMeta `json:",inline"`
	Metadata        metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	RorMeta         ResourceRorMeta   `json:"rormeta"`
}

// ResourceRorMeta represents the metadata stored by ror
type ResourceRorMeta struct {
	Version      string                                     `json:"version,omitempty"`
	LastReported string                                     `json:"lastReported,omitempty"`
	Internal     bool                                       `json:"internal,omitempty"`
	Hash         string                                     `json:"hash,omitempty"`
	Ownerref     rorresourceowner.RorResourceOwnerReference `json:"ownerref,omitempty"`
	Action       ResourceAction                             `json:"action,omitempty"`
}

// GetName returns the name of the resource
func (r *CommonResource) GetName() string {
	return r.Metadata.Name
}

// GetUID returns the UID of the resource
func (r CommonResource) GetUID() string {
	return string(r.Metadata.UID)
}

// GetKind returns the kind of the resource
func (r CommonResource) GetKind() string {
	return r.TypeMeta.Kind
}

// GetAPIVersion returns the API version of the resource
func (r CommonResource) GetAPIVersion() string {
	return r.TypeMeta.APIVersion
}

// GetMetadata returns the metadata of the resource
func (r CommonResource) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// GetRorMeta returns the RorMeta of the resource
func (r CommonResource) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// SetMetadata sets the metadata of the resource
func (r *CommonResource) SetMetadata(meta metav1.ObjectMeta) {
	r.Metadata = meta
}

// SetRorMeta sets the RorMeta of the resource
func (r *CommonResource) SetRorMeta(meta ResourceRorMeta) error {
	r.RorMeta = meta
	return nil
}
