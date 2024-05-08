package rortypes

// K8s StorageClass struct
type ResourceStorageClass struct {
	CommonResource       `json:",inline"`
	AllowVolumeExpansion bool   `json:"allowVolumeExpansion"`
	Provisioner          string `json:"provisioner"`
	ReclaimPolicy        string `json:"reclaimPolicy"`
	VolumeBindingMode    string `json:"volumeBindingMode"`
}
