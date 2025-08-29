package apiresourcecontracts

// K8s StorageClass struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceStorageClass struct {
	ApiVersion           string           `json:"apiVersion"`
	Kind                 string           `json:"kind"`
	Metadata             ResourceMetadata `json:"metadata"`
	AllowVolumeExpansion bool             `json:"allowVolumeExpansion"`
	Provisioner          string           `json:"provisioner"`
	ReclaimPolicy        string           `json:"reclaimPolicy"`
	VolumeBindingMode    string           `json:"volumeBindingMode"`
}
