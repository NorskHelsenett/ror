package rortypes

// ResourceTanzuKubernetesCluster
// K8s node struct
type ResourceTanzuKubernetesCluster struct {
	Spec   ResourceTanuzKuberntesClusterSpec    `json:"spec"`
	Status ResourceTanzuKubernetesClusterStatus `json:"status,omitempty"`
}

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

type ResourceTanuzKuberntesMetadataManagedFields struct {
	ApiVersion string         `json:"apiVersion"`
	FieldsType string         `json:"fieldsType"`
	FieldsV1   map[string]any `json:"fieldsV1"`
	Manager    string         `json:"manager"`
	Operation  string         `json:"operation"`
	Time       string         `json:"time"`
}

type ResourceTanuzKuberntesMetadataOwnerReferences struct {
	ApiVersion         string `json:"apiVersion"`
	BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
	Controller         bool   `json:"controller"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Uid                string `json:"uid"`
}

type ResourceTanuzKuberntesClusterSpec struct {
	Distribution ResourceTanzuKubernetesClusterSpecDistribution `json:"distribution"`
	Settings     ResourceTanzuKubernetesClusterSpecSettings     `json:"settings"`
	Topology     ResourceTanzuKubernetesClusterSpecTopology     `json:"topology"`
}

type ResourceTanzuKubernetesClusterSpecSettings struct {
	Network ResourceTanzuKubernetesClusterSpecSettingsNetwork `json:"network"`
	Storage ResourceTanzuKubernetesClusterSpecSettingsStorage `json:"storage"`
}

type ResourceTanzuKubernetesClusterSpecSettingsStorage struct {
	Classes      []string `json:"classes"`
	DefaultClass string   `json:"defaultClass"`
}

type ResourceTanzuKubernetesClusterSpecSettingsNetwork struct {
	Cni           ResourceTanzuKubernetesClusterSpecSettingsNetworkCni      `json:"cni"`
	Pods          ResourceTanzuKubernetesClusterSpecSettingsNetworkPods     `json:"pods"`
	Proxy         ResourceTanzuKubernetesClusterSpecSettingsNetworkProxy    `json:"proxy"`
	ServiceDomain string                                                    `json:"serviceDomain"`
	Services      ResourceTanzuKubernetesClusterSpecSettingsNetworkServices `json:"services"`
	Trust         ResourceTanzuKubernetesClusterSpecSettingsNetworkTrust    `json:"trust"`
}

type ResourceTanzuKubernetesClusterSpecSettingsNetworkCni struct {
	Name string `json:"name"`
}

type ResourceTanzuKubernetesClusterSpecSettingsNetworkPods struct {
	CidrBlocks []string `json:"cidrBlocks"`
}

type ResourceTanzuKubernetesClusterSpecSettingsNetworkProxy struct {
	HTTPProxy  string `json:"httpProxy"`
	HTTPSProxy string `json:"httpsProxy"`
	NoProxy    string `json:"noProxy"`
}

type ResourceTanzuKubernetesClusterSpecSettingsNetworkServices struct {
	CidrBlocks []string `json:"cidrBlocks"`
}

type ResourceTanzuKubernetesClusterSpecSettingsNetworkTrust struct {
	AdditionalTrusCAs []ResourceTanzuKubernetesClusterSpecSettingsNetworkTrustAdditionalTrustedCA `json:"additionalTrustedCAs"`
}

type ResourceTanzuKubernetesClusterSpecSettingsNetworkTrustAdditionalTrustedCA struct {
	Data string `json:"data"`
	Name string `json:"name"`
}

type ResourceTanzuKubernetesClusterSpecTopology struct {
	ControlPlane ResourceTanzuKubernetesClusterSpecTopologyControlPlane `json:"controlPlane"`
	NodePools    []ResourceTanzuKubernetesClusterSpecTopologyNodePools  `json:"nodePools"`
}

type ResourceTanzuKubernetesClusterSpecTopologyControlPlane struct {
	NodeDrainTimeout string                                                    `json:"nodeDrainTimeout"`
	Replicas         int                                                       `json:"replicas"`
	StorageClass     string                                                    `json:"storageClass"`
	Tkr              ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkr `json:"tkr"`
	VmClass          string                                                    `json:"vmClass"`
}

type ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkr struct {
	Reference ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkrReference `json:"reference"`
}

type ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkrReference struct {
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
	//ResourceVersion string `json:"resourceVersion"`
	Uid string `json:"uid"`
}

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

type ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTaints struct {
	Effect    string `json:"effect"`
	Key       string `json:"key"`
	TimeAdded string `json:"timeAdded"`
	Value     string `json:"value"`
}

type ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkr struct {
	Reference ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkrReference `json:"reference"`
}

type ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkrReference struct {
	FieldPath string `json:"fieldPath"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	//ResourceVersion string `json:"resourceVersion"`
	Uid string `json:"uid"`
}

type ResourceTanzuKubernetesClusterSpecTopologyNodePoolsVolumes struct {
	Capasity     map[string]string `json:"capasity"`
	MountPath    string            `json:"mountPath"`
	Name         string            `json:"name"`
	StorageClass string            `json:"storageClass"`
}

type ResourceTanzuKubernetesClusterSpecDistribution struct {
	FullVersion string `json:"fullVersion"`
	Version     string `json:"version"`
}

type ResourceTanzuKubernetesClusterStatus struct {
	//Addons []ResourceTanzuKubernetesClusterStatusAddons `json:"addons"`
	APIEndpoints        []ResourceTanzuKubernetesClusterStatusAPIEndpoints `json:"apiEndpoints"`
	Conditions          []ResourceTanzuKubernetesClusterStatusConditions   `json:"conditions"`
	Phase               string                                             `json:"phase"`
	TotalWorkerReplicas int                                                `json:"totalWorkerReplicas"`
	Version             string                                             `json:"version"`
}

type ResourceTanzuKubernetesClusterStatusConditions struct {
	LastTransitionTime string `json:"lastTransitionTime"`
	Message            string `json:"message"`
	Reason             string `json:"reason"`
	Severity           string `json:"severity"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}

type ResourceTanzuKubernetesClusterStatusAPIEndpoints struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type ResourceTanzuKubernetesClusterStatusAddons struct {
	Conditions []ResourceTanzuKubernetesClusterStatusAddonsConditions `json:"conditions"`
	Name       string                                                 `json:"name"`
	Type       string                                                 `json:"type"`
	Version    string                                                 `json:"version"`
}

type ResourceTanzuKubernetesClusterStatusAddonsConditions struct {
	LastTransitionTime string `json:"lastTransitionTime"`
	Message            string `json:"message"`
	Reason             string `json:"reason"`
	Severity           string `json:"severity"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}
