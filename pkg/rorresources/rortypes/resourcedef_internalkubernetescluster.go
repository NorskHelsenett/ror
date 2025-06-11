package rortypes

type ResourceKubernetesCluster struct {
	Spec   KubernetesClusterSpec   `json:"spec"`
	Status KubernetesClusterStatus `json:"status,omitempty"`
}

type KubernetesClusterSpec struct {
	Data     KubernetesClusterSpecData     `json:"data,omitempty"`
	Topology KubernetesClusterSpecTopology `json:"topology,omitempty"`
}

type KubernetesClusterSpecData struct {
	ClusterId   string `json:"clusterId"`
	Provider    string `json:"provider"`
	Datacenter  string `json:"datacenter"`
	Region      string `json:"region"`
	Zone        string `json:"zone"`
	Project     string `json:"project"`
	Workspace   string `json:"workspace"`
	Workorder   string `json:"workorder"`
	Environment string `json:"environment"`
}

type KubernetesClusterSpecTopology struct {
	Version      string       `json:"version"`
	ControlPlane ControlPlane `json:"controlplane"`
	Workers      Workers      `json:"workers"`
}

type ControlPlane struct {
	Replicas     int             `json:"replicas"`
	Provider     string          `json:"provider"`
	MachineClass string          `json:"machineClass"`
	Metadata     MetadataDetails `json:"metadata"`
	Storage      []Storage       `json:"storage"`
}

type MetadataDetails struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type Storage struct {
	Class string `json:"class"`
	Path  string `json:"path"`
	Size  string `json:"size"`
}

type Workers struct {
	NodePools []NodePool `json:"nodePools"`
}

type NodePool struct {
	MachineClass string          `json:"machineClass"`
	Provider     string          `json:"provider"`
	Name         string          `json:"name"`
	Replicas     int             `json:"replicas"`
	Autoscaling  Autoscaling     `json:"autoscaling"`
	Metadata     MetadataDetails `json:"metadata"`
}

type Autoscaling struct {
	Enabled      bool     `json:"enabled"`
	MinReplicas  int      `json:"minReplicas"`
	MaxReplicas  int      `json:"maxReplicas"`
	ScalingRules []string `json:"scalingRules"`
}

type KubernetesClusterStatus struct {
	State      ClusterState `json:"state"`
	Phase      string       `json:"phase"` // Provisioning, Running, Deleting, Failed, Updating
	Conditions []Condition  `json:"conditions"`
}

type ClusterState struct {
	Cluster              ClusterDetails `json:"cluster"`
	Versions             []Version      `json:"versions"`
	ControlplaneEndpoint string         `json:"controlplaneendpoint"`
	EgressIP             string         `json:"egressIP"`
	LastUpdated          string         `json:"lastUpdated"`
	LastUpdatedBy        string         `json:"lastUpdatedBy"`
	Created              string         `json:"created"`
}

type ClusterDetails struct {
	ExternalId         string             `json:"externalId"`
	Resources          []Resource         `json:"resources"`
	ControlPlaneStatus ControlPlaneStatus `json:"controlplane"`
	Workers            WorkerStatus       `json:"workers"`
}

type Resource struct {
	Name      string `json:"name"`
	Allocated string `json:"allocated"`
	Usage     string `json:"usage"`
}

type ControlPlaneStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type WorkerStatus struct {
	NodePools []NodePoolStatus `json:"nodepools"`
}

type NodePoolStatus struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Version struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Branch  string `json:"branch"`
}

type Condition struct {
	Type               string `json:"type"`
	Status             string `json:"status"`
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Message            string `json:"message"`
}
