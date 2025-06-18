package apiresourcecontracts

// ResourceTanzuKubernetesCluster
// K8s node struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesCluster struct {
	ApiVersion string                               `json:"apiVersion"`
	Kind       string                               `json:"kind"`
	Metadata   ResourceTanuzKuberntesMetadata       `json:"metadata"`
	Spec       ResourceTanuzKuberntesClusterSpec    `json:"spec"`
	Status     ResourceTanzuKubernetesClusterStatus `json:"status,omitempty"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanuzKuberntesMetadata struct {
	Annotations                map[string]string                               `json:"annotations"`
	ClusterName                string                                          `json:"clusterName"`
	CreationTimestamp          string                                          `json:"creationTimestamp"`
	DeletionGracePeriodSeconds int                                             `json:"deletionGracePeriodSeconds"`
	DeletionTimestamp          string                                          `json:"deletionTimestamp"`
	Finalizers                 []string                                        `json:"finalizers"`
	GenerateName               string                                          `json:"generateName"`
	Generation                 int                                             `json:"generation"`
	Labels                     map[string]string                               `json:"labels"`
	ManagedFields              []ResourceTanuzKuberntesMetadataManagedFields   `json:"managedFields"`
	Name                       string                                          `json:"name"`
	Namespace                  string                                          `json:"namespace"`
	OwnerReferences            []ResourceTanuzKuberntesMetadataOwnerReferences `json:"ownerReferences"`
	//ResourceVersion            string                                          `json:"resourceVersion"`
	SelfLink string `json:"selfLink"`
	Uid      string `json:"uid"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanuzKuberntesMetadataManagedFields struct {
	ApiVersion string         `json:"apiVersion"`
	FieldsType string         `json:"fieldsType"`
	FieldsV1   map[string]any `json:"fieldsV1"`
	Manager    string         `json:"manager"`
	Operation  string         `json:"operation"`
	Time       string         `json:"time"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanuzKuberntesMetadataOwnerReferences struct {
	ApiVersion         string `json:"apiVersion"`
	BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
	Controller         bool   `json:"controller"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Uid                string `json:"uid"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanuzKuberntesClusterSpec struct {
	Distribution ResourceTanzuKubernetesClusterSpecDistribution `json:"distribution"`
	Settings     ResourceTanzuKubernetesClusterSpecSettings     `json:"settings"`
	Topology     ResourceTanzuKubernetesClusterSpecTopology     `json:"topology"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecSettings struct {
	Network ResourceTanzuKubernetesClusterSpecSettingsNetwork `json:"network"`
	Storage ResourceTanzuKubernetesClusterSpecSettingsStorage `json:"storage"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecSettingsStorage struct {
	Classes      []string `json:"classes"`
	DefaultClass string   `json:"defaultClass"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecSettingsNetwork struct {
	Cni           ResourceTanzuKubernetesClusterSpecSettingsNetworkCni      `json:"cni"`
	Pods          ResourceTanzuKubernetesClusterSpecSettingsNetworkPods     `json:"pods"`
	Proxy         ResourceTanzuKubernetesClusterSpecSettingsNetworkProxy    `json:"proxy"`
	ServiceDomain string                                                    `json:"serviceDomain"`
	Services      ResourceTanzuKubernetesClusterSpecSettingsNetworkServices `json:"services"`
	Trust         ResourceTanzuKubernetesClusterSpecSettingsNetworkTrust    `json:"trust"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecSettingsNetworkCni struct {
	Name string `json:"name"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecSettingsNetworkPods struct {
	CidrBlocks []string `json:"cidrBlocks"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecSettingsNetworkProxy struct {
	HTTPProxy  string `json:"httpProxy"`
	HTTPSProxy string `json:"httpsProxy"`
	NoProxy    string `json:"noProxy"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecSettingsNetworkServices struct {
	CidrBlocks []string `json:"cidrBlocks"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecSettingsNetworkTrust struct {
	AdditionalTrusCAs []ResourceTanzuKubernetesClusterSpecSettingsNetworkTrustAdditionalTrustedCA `json:"additionalTrustedCAs"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecSettingsNetworkTrustAdditionalTrustedCA struct {
	Data string `json:"data"`
	Name string `json:"name"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecTopology struct {
	ControlPlane ResourceTanzuKubernetesClusterSpecTopologyControlPlane `json:"controlPlane"`
	NodePools    []ResourceTanzuKubernetesClusterSpecTopologyNodePools  `json:"nodePools"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecTopologyControlPlane struct {
	NodeDrainTimeout string                                                    `json:"nodeDrainTimeout"`
	Replicas         int                                                       `json:"replicas"`
	StorageClass     string                                                    `json:"storageClass"`
	Tkr              ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkr `json:"tkr"`
	VmClass          string                                                    `json:"vmClass"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkr struct {
	Reference ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkrReference `json:"reference"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkrReference struct {
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
	//ResourceVersion string `json:"resourceVersion"`
	Uid string `json:"uid"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecTopologyNodePools struct {
	FailureDomain    string                                                       `json:"failureDomain"`
	Labels           map[string]string                                            `json:"labels"`
	Name             string                                                       `json:"name"`
	NodeDrainTimeout string                                                       `json:"nodeDrainTimeout"`
	Replicas         int                                                          `json:"replicas"`
	StorageClass     string                                                       `json:"storageClass"`
	Taints           []ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTaints  `json:"taints"`
	Tkr              ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkr       `json:"tkr"`
	VmClass          string                                                       `json:"vmClass"`
	Volumes          []ResourceTanzuKubernetesClusterSpecTopologyNodePoolsVolumes `json:"volumes"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTaints struct {
	Effect    string `json:"effect"`
	Key       string `json:"key"`
	TimeAdded string `json:"timeAdded"`
	Value     string `json:"value"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkr struct {
	Reference ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkrReference `json:"reference"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkrReference struct {
	FieldPath string `json:"fieldPath"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	//ResourceVersion string `json:"resourceVersion"`
	Uid string `json:"uid"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecTopologyNodePoolsVolumes struct {
	Capasity     map[string]string `json:"capasity"`
	MountPath    string            `json:"mountPath"`
	Name         string            `json:"name"`
	StorageClass string            `json:"storageClass"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterSpecDistribution struct {
	FullVersion string `json:"fullVersion"`
	Version     string `json:"version"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterStatus struct {
	//Addons []ResourceTanzuKubernetesClusterStatusAddons `json:"addons"`
	APIEndpoints        []ResourceTanzuKubernetesClusterStatusAPIEndpoints `json:"apiEndpoints"`
	Conditions          []ResourceTanzuKubernetesClusterStatusConditions   `json:"conditions"`
	Phase               string                                             `json:"phase"`
	TotalWorkerReplicas int                                                `json:"totalWorkerReplicas"`
	Version             string                                             `json:"version"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterStatusConditions struct {
	LastTransitionTime string `json:"lastTransitionTime"`
	Message            string `json:"message"`
	Reason             string `json:"reason"`
	Severity           string `json:"severity"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterStatusAPIEndpoints struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterStatusAddons struct {
	Conditions []ResourceTanzuKubernetesClusterStatusAddonsConditions `json:"conditions"`
	Name       string                                                 `json:"name"`
	Type       string                                                 `json:"type"`
	Version    string                                                 `json:"version"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceTanzuKubernetesClusterStatusAddonsConditions struct {
	LastTransitionTime string `json:"lastTransitionTime"`
	Message            string `json:"message"`
	Reason             string `json:"reason"`
	Severity           string `json:"severity"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}
