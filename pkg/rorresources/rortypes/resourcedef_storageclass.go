package rortypes

// K8s StorageClass struct
type ResourceStorageClass struct {
	AllowVolumeExpansion bool   `json:"allowVolumeExpansion"`
	Provisioner          string `json:"provisioner"`
	ReclaimPolicy        string `json:"reclaimPolicy"`
	VolumeBindingMode    string `json:"volumeBindingMode"`
}

// (r ResourceStorageClass) Get returns a pointer to the resource of type ResourceStorageClass
func (r *ResourceStorageClass) Get() *ResourceStorageClass {
	return r
}
