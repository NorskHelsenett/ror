package apiresourcecontracts

// K8s PersistentVolumeClaim struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourcePersistentVolumeClaim struct {
	ApiVersion string                              `json:"apiVersion"`
	Kind       string                              `json:"kind"`
	Metadata   ResourceMetadata                    `json:"metadata"`
	Spec       ResourcePersistentVolumeClaimSpec   `json:"spec"`
	Status     ResourcePersistentVolumeClaimStatus `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourcePersistentVolumeClaimSpec struct {
	AaccessModes     []string                                   `json:"accessModes"`
	Resources        ResourcePersistentVolumeClaimSpecResources `json:"resources"`
	StorageClassName string                                     `json:"storageClassName"`
	VolumeMode       string                                     `json:"volumeMode"`
	VolumeName       string                                     `json:"volumeName"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourcePersistentVolumeClaimSpecResources struct {
	Limits   map[string]string `json:"limits,omitempty"`
	Requests map[string]string `json:"requests"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourcePersistentVolumeClaimStatus struct {
	AaccessModes []string          `json:"accessModes"`
	Capacity     map[string]string `json:"capacity"`
	Phase        string            `json:"phase"`
}
