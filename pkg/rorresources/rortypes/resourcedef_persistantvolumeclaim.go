package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// K8s PersistentVolumeClaim struct
type ResourcePersistentVolumeClaim struct {
	Spec   ResourcePersistentVolumeClaimSpec   `json:"spec"`
	Status ResourcePersistentVolumeClaimStatus `json:"status"`
}

type ResourcePersistentVolumeClaimSpec struct {
	AaccessModes     []string                                   `json:"accessModes"`
	Resources        ResourcePersistentVolumeClaimSpecResources `json:"resources"`
	StorageClassName string                                     `json:"storageClassName"`
	VolumeMode       string                                     `json:"volumeMode"`
	VolumeName       string                                     `json:"volumeName"`
}
type ResourcePersistentVolumeClaimSpecResources struct {
	Limits   map[string]string `json:"limits,omitempty"`
	Requests map[string]string `json:"requests"`
}
type ResourcePersistentVolumeClaimStatus struct {
	AaccessModes []string          `json:"accessModes"`
	Capacity     map[string]string `json:"capacity"`
	Phase        string            `json:"phase"`
}

// (r *ResourcePersistentVolumeClaim) ApplyInputFilter Applies the input filter to the resource
func (r *ResourcePersistentVolumeClaim) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Namespace] {
		cr.RorMeta.Internal = true
	}
	return nil
}

// (r ResourcePersistentVolumeClaim) Get returns a pointer to the resource of type ResourcePersistentVolumeClaim
func (r *ResourcePersistentVolumeClaim) Get() *ResourcePersistentVolumeClaim {
	return r
}

// PersistentVolumeClaiminterface represents the interface for resources of the type persistentvolumeclaim
type PersistentVolumeClaiminterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourcePersistentVolumeClaim
}
