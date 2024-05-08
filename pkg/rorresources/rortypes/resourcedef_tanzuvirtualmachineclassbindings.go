package rortypes

// ResourceTanzuKubernetesCluster
// K8s node struct
// Tanzu kubernetes release struct
type ResourceVirtualMachineClassBinding struct {
	CommonResource `json:",inline"`
	ClassRef       ResourceVirtualMachineClassBindingClassRef `json:"classRef"`
}

type ResourceVirtualMachineClassBindingMetadata struct {
	Annotations                map[string]string                                          `json:"annotations"`
	ClusterName                string                                                     `json:"clusterName"`
	CreationTimestamp          string                                                     `json:"creationTimestamp"`
	DeletionGracePeriodSeconds int                                                        `json:"deletionGracePeriodSeconds"`
	DeletionTimestamp          string                                                     `json:"deletionTimestamp"`
	Finalizers                 []string                                                   `json:"finalizers"`
	GenerateName               string                                                     `json:"generateName"`
	Generation                 int                                                        `json:"generation"`
	Labels                     map[string]string                                          `json:"labels"`
	ManagedFields              []ResourceVirtualMachineClassBindingMetadataManagedField   `json:"managedFields"`
	Name                       string                                                     `json:"name"`
	Namespace                  string                                                     `json:"namespace"`
	OwnerReferences            []ResourceVirtualMachineClassBindingMetadataOwnerReference `json:"ownerReferences"`
	//ResourceVersion            string                                                     `json:"resourceVersion"`
	SelfLink string `json:"selfLink"`
	Uid      string `json:"uid"`
}

type ResourceVirtualMachineClassBindingMetadataManagedField struct {
	ApiVersion string `json:"apiVersion"`
	FieldsType string `json:"fieldsType"`
	//FieldsV1    map[string]interface{} `json:"fieldsV1"`
	Manager     string `json:"manager"`
	Operation   string `json:"operation"`
	Subresource string `json:"subresource"`
	Time        string `json:"time"`
}

type ResourceVirtualMachineClassBindingMetadataOwnerReference struct {
	ApiVersion         string `json:"apiVersion"`
	BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
	Controller         bool   `json:"controller"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Uid                string `json:"uid"`
}

type ResourceVirtualMachineClassBindingClassRef struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
}
