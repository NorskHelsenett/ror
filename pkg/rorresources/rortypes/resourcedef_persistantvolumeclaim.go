package rortypes

// K8s PersistentVolumeClaim struct
type ResourcePersistentVolumeClaim struct {
	CommonResource `json:",inline"`
	Spec           ResourcePersistentVolumeClaimSpec   `json:"spec"`
	Status         ResourcePersistentVolumeClaimStatus `json:"status"`
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
