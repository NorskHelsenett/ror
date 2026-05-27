package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// K8s deployment struct
type ResourceDaemonSet struct {
	Status ResourceDaemonSetStatus `json:"status"`
}
type ResourceDaemonSetStatus struct {
	NumberReady            int `json:"numberReady"`
	NumberUnavailable      int `json:"numberUnavailable"`
	NumberMisscheduled     int `json:"currentReplicas"`
	NumberAvailable        int `json:"numberAvailable"`
	UpdatedNumberScheduled int `json:"updatedNumberScheduled"`
	DesiredNumberScheduled int `json:"desiredNumberScheduled"`
	CurrentNumberScheduled int `json:"currentNumberScheduled"`
}

// (r *ResourceDaemonSet) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceDaemonSet) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Namespace] {
		cr.RorMeta.Internal = true
	}
	return nil
}
