package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// K8s deployment struct
type ResourceDeployment struct {
	Status ResourceDeploymentStatus `json:"status"`
}
type ResourceDeploymentStatus struct {
	Replicas          int `json:"replicas"`
	AvailableReplicas int `json:"availableReplicas"`
	ReadyReplicas     int `json:"readyReplicas"`
	UpdatedReplicas   int `json:"updatedReplicas"`
}

// (r *ResourceDeployment) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceDeployment) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Namespace] {
		cr.RorMeta.Internal = true
	}
	return nil
}

// (r ResourceDeployment) Get returns a pointer to the resource of type ResourceDeployment
func (r *ResourceDeployment) Get() *ResourceDeployment {
	return r
}

// Deploymentinterface represents the interface for resources of the type deployment
type Deploymentinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceDeployment
}
