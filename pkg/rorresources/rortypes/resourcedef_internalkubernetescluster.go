package rortypes

type ResourceKubernetesCluster struct {
	Spec   ResourceKubernetesClusterSpec   `json:"spec"`
	Status ResourceKubernetesClusterStatus `json:"status"`
}

type ResourceKubernetesClusterSpec struct {
	ClusterId     string                                     `json:"clusterId"`
	ClusterName   string                                     `json:"clusterName"`
	Description   string                                     `json:"description"`
	Project       string                                     `json:"project"`
	Provider      string                                     `json:"provider"`
	CreatedBy     string                                     `json:"createdBy"`
	ToolingConfig ResourceKubernetesClusterSpecToolingConfig `json:"toolingConfig"`
	Environment   string                                     `json:"environment"`
	ProviderSpec  ResourceKubernetesClusterSpecProviderSpec  `json:"providerSpec"`
	Topology      ResourceKubernetesClusterSpecTopology      `json:"topology"`
	Endpoints     []ResourceKubernetesClusterSpecEndpoint    `json:"endpoints"`
}

type ResourceKubernetesClusterSpecProviderSpec struct {
	TanzuSpec ResourceKubernetesClusterSpecProviderSpecTanzuSpec `json:"tanzuSpec"`
	AzureSpec ResourceKubernetesClusterSpecProviderSpecAzureSpec `json:"azureSpec"`
}

type ResourceKubernetesClusterSpecProviderSpecTanzuSpec struct {
	SupervisorClusterName string `json:"supervisorClusterName"`
	Namespace             string `json:"namespace"`
}

type ResourceKubernetesClusterSpecProviderSpecAzureSpec struct {
	SubscriptionId string `json:"subscriptionId"`
	ResourceGroup  string `json:"resourceGroup"`
}

type ResourceKubernetesClusterSpecToolingConfig struct {
	SplunkIndex string `json:"splunkIndex"`
}

type ResourceKubernetesClusterSpecEndpoint struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}
type ResourceKubernetesClusterSpecTopology struct {
	ControlPlane ResourceKubernetesClusterSpecTopologyControlPlane `json:"controlPlane"`
	Workers      []ResourceKubernetesClusterSpecTopologyWorkers    `json:"workers"`
}

type ResourceKubernetesClusterSpecTopologyControlPlane struct {
	Replicas     int    `json:"replicas"`
	Version      string `json:"version"`
	MachineClass string `json:"machineClass"`
}
type ResourceKubernetesClusterSpecTopologyWorkers struct {
	Name         string `json:"name"`
	Replicas     int    `json:"replicas"`
	Version      string `json:"version"`
	MachineClass string `json:"machineClass"`
}

type ResourceKubernetesClusterStatus struct {
	Status            string                                     `json:"status"`
	Phase             string                                     `json:"phase"`
	Conditions        []ResourceKubernetesClusterStatusCondition `json:"conditions"`
	KubernetesVersion string                                     `json:"kubernetesVersion"`
	ProviderStatus    map[string]interface{}                     `json:"providerStatus"`
	CreatedTime       string                                     `json:"createdTime"`
	UpdatedTime       string                                     `json:"updatedTime"`
	LastObservedTime  string                                     `json:"lastObservedTime"`
}

type ResourceKubernetesClusterStatusCondition struct {
	Type               string `json:"type"`
	Status             string `json:"status"`
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Message            string `json:"message"`
}
