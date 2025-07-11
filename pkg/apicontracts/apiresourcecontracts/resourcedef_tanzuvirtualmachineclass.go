package apiresourcecontracts

// ResourceTanzuKubernetesCluster
// K8s node struct
// Tanzu kubernetes release struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachineClass struct {
	ApiVersion string                              `json:"apiVersion"`
	Kind       string                              `json:"kind"`
	Metadata   ResourceVirtualMachineClassMetadata `json:"metadata"`
	Spec       ResourceVirtualMachineClassSpec     `json:"spec"`
	Status     map[string]string                   `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachineClassMetadata struct {
	Annotations map[string]string `json:"annotations"`
	ClusterName string            `json:"clusterName"`
	//CreationTimestamp     string                                                  `json:"creationTimestamp"`
	//DeletionGracePeriodSeconds int                                                     `json:"deletionGracePeriodSeconds"`
	//DeletionTimestamp          string                                                  `json:"deletionTimestamp"`
	Finalizers      []string                                            `json:"finalizers"`
	GenerateName    string                                              `json:"generateName"`
	Generation      int                                                 `json:"generation"`
	Labels          map[string]string                                   `json:"labels"`
	ManagedFields   []ResourceVirtualMachineClassMetadataManagedField   `json:"managedFields"`
	Name            string                                              `json:"name"`
	Namespace       string                                              `json:"namespace"`
	OwnerReferences []ResourceVirtualMachineClassMetadataOwnerReference `json:"ownerReferences"`
	//ResourceVersion string                                              `json:"resourceVersion"`
	SelfLink string `json:"selfLink"`
	Uid      string `json:"uid"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachineClassMetadataManagedField struct {
	ApiVersion string `json:"apiVersion"`
	FieldsType string `json:"fieldsType"`
	//FieldsV1    map[string]interface{} `json:"fieldsV1"`
	Manager     string `json:"manager"`
	Operation   string `json:"operation"`
	Subresource string `json:"subresource"`
	Time        string `json:"time"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachineClassMetadataOwnerReference struct {
	ApiVersion         string `json:"apiVersion"`
	BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
	Controller         bool   `json:"controller"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Uid                string `json:"uid"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachineClassSpec struct {
	Description string                                  `json:"description"`
	Hardware    ResourceVirtualMachineClassSpecHardware `json:"hardware"`
	//Policies ResourceVirtualMachineClassSpecPolicies `json:"policies"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachineClassSpecHardware struct {
	Cpus int `json:"cpus"`
	//Devices ResourceVirtualMachineClassSpecHardwareDevice `json:"devices"`
	InstanceStorage ResourceVirtualMachineClassSpecHardwareInstanceStorage `json:"instanceStorage"`
}

// type ResourceVirtualMachineClassSpecHardwareDevice struct {
// }
// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachineClassSpecHardwareInstanceStorage struct {
	StorageClass string `json:"storageClass"`
	//Volumes 	[]ResourceVirtualMachineClassSpecHardwareInstanceStorageVolumes `json:"volumes"`
}

// type ResourceVirtualMachineClassSpecHardwareInstanceStorageVolumes struct {
// 	Capacity map[string]string `json:"capacity"`
// 	Name string `json:"name"`
// }

// type ResourceVirtualMachineClassSpecPolicies struct {
// 	Resources ResourceVirtualMachineClassSpecPoliciesResources `json:"resources"`
// }

// type ResourceVirtualMachineClassSpecPoliciesResources struct {
// 	Limits ResourceVirtualMachineClassSpecPoliciesResourcesLimits `json:"limits"`
// 	Requests ResourceVirtualMachineClassSpecPoliciesResourcesRequests `json:"requests"`
// }

// type ResourceVirtualMachineClassSpecPoliciesResourcesLimits struct {

// }

// type ResourceVirtualMachineClassSpecPoliciesResourcesRequests struct {
// }
