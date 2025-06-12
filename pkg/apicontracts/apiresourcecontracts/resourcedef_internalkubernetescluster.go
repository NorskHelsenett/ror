package apiresourcecontracts

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesCluster struct {
	ApiVersion string                          `json:"apiVersion"`
	Kind       string                          `json:"kind"`
	Metadata   ResourceMetadata                `json:"metadata"`
	Spec       ResourceKubernetesClusterSpec   `json:"spec"`
	Status     ResourceKubernetesClusterStatus `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
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
	Endpoints     []ResourceKubernetesClusterSpecEndpoint    `json:"endpoints"` // TODO: remove this field, use ResourceKubernetesClusterStatus instead
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesClusterSpecProviderSpec struct {
	TanzuSpec ResourceKubernetesClusterSpecProviderSpecTanzuSpec `json:"tanzuSpec"`
	AzureSpec ResourceKubernetesClusterSpecProviderSpecAzureSpec `json:"azureSpec"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesClusterSpecProviderSpecTanzuSpec struct {
	SupervisorClusterName string `json:"supervisorClusterName"`
	Namespace             string `json:"namespace"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesClusterSpecProviderSpecAzureSpec struct {
	SubscriptionId string `json:"subscriptionId"`
	ResourceGroup  string `json:"resourceGroup"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesClusterSpecToolingConfig struct {
	SplunkIndex string `json:"splunkIndex"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesClusterSpecEndpoint struct {
	Type    string `json:"type"`    // Type of the endpoint, e.g., "kubernetesApi", "grafana", "argocd",etc.
	Address string `json:"address"` // Address of the endpoint, e.g., "https://api.example.com", "https://grafana.example.com", etc.
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesClusterSpecTopology struct {
	ControlPlane ResourceKubernetesClusterSpecTopologyControlPlane `json:"controlPlane"`
	Workers      []ResourceKubernetesClusterSpecTopologyWorkers    `json:"workers"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesClusterSpecTopologyControlPlane struct {
	Replicas     int    `json:"replicas"`
	Version      string `json:"version"`
	MachineClass string `json:"machineClass"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesClusterSpecTopologyWorkers struct {
	Name         string `json:"name"`
	Replicas     int    `json:"replicas"`
	Version      string `json:"version"`
	MachineClass string `json:"machineClass"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesClusterStatus struct {
	Status            string                                       `json:"status"`
	Endpoints         []ResourceKubernetesClusterSpecEndpoint      `json:"endpoints"`
	Phase             string                                       `json:"phase"`
	Conditions        []ResourceKubernetesClusterStatusCondition   `json:"conditions"`
	KubernetesVersion string                                       `json:"kubernetesVersion"`
	ProviderStatus    map[string]interface{}                       `json:"providerStatus"`
	CreatedTime       string                                       `json:"createdTime"`
	UpdatedTime       string                                       `json:"updatedTime"`
	LastObservedTime  string                                       `json:"lastObservedTime"`
	ClusterStatus     ResourceKubernetesClusterStatusClusterStatus `json:"clusterStatus"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesClusterStatusCondition struct {
	Type               string `json:"type" example:"ClusterReady"`                                   // Type is the type of the condition. For example, "ready", "available", etc.
	Status             string `json:"status"  example:"ok" enums:"ok,warning,error,working,unknown"` // Status is the status of the condition. Valid vales are: ok, warning, error, working, unknown.
	LastTransitionTime string `json:"lastTransitionTime"`                                            // LastTransitionTime is the last time the condition transitioned from one status to another.
	Reason             string `json:"reason"`                                                        // Reason is a brief reason for the condition's last transition.
	Message            string `json:"message"`                                                       // Message is a human-readable message indicating details about the condition.
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesClusterStatusClusterStatus struct {
	Price     ResourceKubernetesClusterStatusClusterStatusPrice    `json:"price"`     // Price is the price of the cluster, e.g., "1000 NOK/month"
	NodePools int                                                  `json:"nodePools"` // NodePools is the number of node pools in the cluster.
	Nodes     int                                                  `json:"nodes"`     // Nodes is the number of nodes in the cluster.
	CPU       ResourceKubernetesClusterStatusClusterStatusResource `json:"cpu"`       // CPU is the total CPU capacity of the cluster, if not specified in millicores, e.g., "16 cores", "8000 millicores"
	Memory    ResourceKubernetesClusterStatusClusterStatusResource `json:"memory"`    // Memory is the total memory capacity of the cluster, if not specified in bytes, e.g., "64 GB", "128000 MB", "25600000000 bytes"
	GPU       ResourceKubernetesClusterStatusClusterStatusResource `json:"gpu"`       // GPU is the total GPU capacity of the cluster, if not specified in number of GPUs"
	Disk      ResourceKubernetesClusterStatusClusterStatusResource `json:"disk"`      // Disk is the total disk capacity of the cluster, if not specified in bytes"
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesClusterStatusClusterStatusResource struct {
	Capacity  string `json:"capacity"`   // Capacity is the total capacity of the resource, e.g., "16", "64Gi", "900m"
	Used      string `json:"used"`       // Used is the amount of the resource that is currently used, e.g., "16", "64Gi", "900m"
	Percetage int    `json:"percentage"` // Percentage is the percentage of the resource that is currently used, e.g., "50"
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceKubernetesClusterStatusClusterStatusPrice struct {
	Monthly int `json:"monthly"` // Monthly is the monthly price of the cluster in your currency, e.g., "1000"
	Yearly  int `json:"yearly"`  // Yearly is the yearly price of the cluster, e.g., "12000"
}
