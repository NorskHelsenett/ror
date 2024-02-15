package apiresourcecontracts

// K8s StorageClass struct
type ResourceStorageClass struct {
	ApiVersion           string           `json:"apiVersion"`
	Kind                 string           `json:"kind"`
	Metadata             ResourceMetadata `json:"metadata"`
	AllowVolumeExpansion bool             `json:"allowVolumeExpansion"`
	Provisioner          string           `json:"provisioner"`
	ReclaimPolicy        string           `json:"reclaimPolicy"`
	VolumeBindingMode    string           `json:"volumeBindingMode"`
}
