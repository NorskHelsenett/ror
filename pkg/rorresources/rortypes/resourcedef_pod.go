package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// ResourcePod
// K8s namepace struct
type ResourcePod struct {
	Spec   ResourcePodSpec   `json:"spec" bson:"spec,omitempty"`
	Status ResourcePodStatus `json:"status" bson:"status,omitempty"`
}

type ResourcePodSpec struct {
	Containers         []ResourcePodSpecContainers `json:"containers,omitempty"`
	ServiceAccountName string                      `json:"serviceAccountName,omitempty"`
	NodeName           string                      `json:"nodeName,omitempty"`
}
type ResourcePodSpecContainers struct {
	Name  string                           `json:"name,omitempty"`
	Image string                           `json:"image,omitempty"`
	Ports []ResourcePodSpecContainersPorts `json:"ports,omitempty"`
}
type ResourcePodSpecContainersPorts struct {
	Name          string `json:"name,omitempty"`
	ContainerPort int    `json:"containerPort,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
}

type ResourcePodStatus struct {
	Message   string `json:"message,omitempty"`
	Phase     string `json:"phase,omitempty"`
	Reason    string `json:"reason,omitempty"`
	StartTime string `json:"startTime,omitempty"`
}

// (r *ResourcePod) ApplyInputFilter Applies the input filter to the resource
func (r *ResourcePod) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Namespace] {
		cr.RorMeta.Internal = true
	}
	return nil
}

// (r ResourcePod) Get returns a pointer to the resource of type ResourcePod
func (r *ResourcePod) Get() *ResourcePod {
	return r
}

// Podinterface represents the interface for resources of the type pod
type Podinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourcePod
}
