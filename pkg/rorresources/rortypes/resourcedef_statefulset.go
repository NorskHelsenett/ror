package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// K8s deployment struct
type ResourceStatefulSet struct {
	Status ResourceStatefulSetStatus `json:"status"`
}
type ResourceStatefulSetStatus struct {
	Replicas          int `json:"replicas"`
	AvailableReplicas int `json:"availableReplicas"`
	CurrentReplicas   int `json:"currentReplicas"`
	ReadyReplicas     int `json:"readyReplicas"`
	UpdatedReplicas   int `json:"updatedReplicas"`
}

// (r *ResourceStatefulSet) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceStatefulSet) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Namespace] {
		cr.RorMeta.Internal = true
	}
	return nil
}
