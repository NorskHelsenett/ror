package rortypes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResourceKubernetesCluster struct {
	Spec   KubernetesClusterSpec   `json:"spec"`
	Status KubernetesClusterStatus `json:"status,omitempty"`
}

type KubernetesClusterSpec struct {
	Cluster  KubernetesClusterSpecData     `json:"data,omitzero"`
	Topology KubernetesClusterSpecTopology `json:"topology,omitzero"`
}

type KubernetesClusterSpecData struct {
	ClusterUID       string               `json:"clusterUid"` // ClusterUID is a unique identifier for the cluster, e.g., "12345678-1234-1234-1234-123456789012"
	ClusterId        string               `json:"clusterId"`
	Provider         RorResourceReference `json:"provider"` // Provider is the infrastructure provider for the cluster, e.g., "aws", "azure", "gcp", "onprem"
	Region           RorResourceReference `json:"region"`
	AvailabilityZone RorResourceReference `json:"az"`
	Project          RorResourceReference `json:"project"`
	Workspace        RorResourceReference `json:"workspace"`
	Workorder        RorResourceReference `json:"workorder"`
	Environment      RorResourceReference `json:"environment"`
}

type KubernetesClusterSpecTopology struct {
	Version      string                            `json:"version"`      // Kubernetes version, e.g., "1.23.0"
	ControlPlane KubernetesClusterSpecControlPlane `json:"controlplane"` // ControlPlane contains the control plane configuration.
	Workers      KubernetesClusterWorkers          `json:"workers"`      // Workers contains the worker nodes configuration.
}

type KubernetesClusterSpecControlPlane struct {
	Replicas     int                                  `json:"replicas"`
	Version      string                               `json:"version"`      // Kubernetes version, e.g., "1.23.0"
	Provider     string                               `json:"provider"`     // Provider is the infrastructure provider for the control plane, e.g., "aws", "azure", "gcp", "onprem"
	MachineClass string                               `json:"machineClass"` // MachineClass is the machine class for the control plane nodes, e.g., "c5.large", "m5.xlarge"
	Metadata     KubernetesClusterSpecMetadataDetails `json:"metadata"`
	Storage      []KubernetesClusterStorage           `json:"storage"` // Storage contains the storage configuration for the control plane nodes. The provider must support empty or wrong storage configuration.
}

type KubernetesClusterSpecMetadataDetails struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type KubernetesClusterStorage struct {
	Class string `json:"class"`
	Path  string `json:"path"`
	Size  string `json:"size"`
}

type KubernetesClusterWorkers struct {
	NodePools []KubernetesClusterNodePool `json:"nodePools"`
}

type KubernetesClusterNodePool struct {
	MachineClass string                               `json:"machineClass"`
	Provider     string                               `json:"provider"`
	Version      string                               `json:"version"` // Kubernetes version, e.g., "1.23.0"
	Name         string                               `json:"name"`
	Replicas     int                                  `json:"replicas"`
	Autoscaling  KubernetesClusterAutoscalingSpec     `json:"autoscaling"`
	Metadata     KubernetesClusterSpecMetadataDetails `json:"metadata"`
	Taint        []KubernetesClusterTaint             `json:"taint"`
}

type KubernetesClusterTaint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

type KubernetesClusterAutoscalingConfig struct {
	Enabled     bool `json:"enabled"`
	MinReplicas int  `json:"minReplicas"`
	MaxReplicas int  `json:"maxReplicas"`
}
type KubernetesClusterAutoscalingSpec struct {
	KubernetesClusterAutoscalingConfig `json:",inline"`
	ScalingRules                       []string `json:"scalingRules"`
}

// KubernetesClusterStatus represents the status of a Kubernetes cluster.
// It contains the current state, phase, and conditions of the cluster.
type KubernetesClusterStatus struct {
	State      KubernetesClusterClusterState `json:"state"`
	Phase      string                        `json:"phase"` // Provisioning, Running, Deleting, Failed, Updating
	Conditions []KubernetesClusterCondition  `json:"conditions"`
}

type KubernetesClusterClusterState struct {
	Cluster       KubernetesClusterClusterDetails `json:"cluster"`
	Versions      []KubernetesClusterVersion      `json:"versions"`
	Endpoints     []KubernetesClusterEndpoint     `json:"endpoints"`
	EgressIP      string                          `json:"egressIP"`
	LastUpdated   metav1.Time                     `json:"lastUpdated"`
	LastUpdatedBy string                          `json:"lastUpdatedBy"`
	Created       metav1.Time                     `json:"created"`
}

type KubernetesClusterEndpoint struct {
	Name    string `json:"name"`    // Name is the name of the endpoint, e.g., "controllplane", "kubernetes", "api", "dashboard, grafana, argocd", "datacenter"
	Address string `json:"address"` // Address is the address of the endpoint, e.g., "https://api.example.com", "http://dashboard.example.com"
}

type KubernetesClusterStatusCondition struct {
	Type               string `json:"type" example:"ClusterReady"`                                   // Type is the type of the condition. For example, "ready", "available", etc.
	Status             string `json:"status"  example:"ok" enums:"ok,warning,error,working,unknown"` // Status is the status of the condition. Valid vales are: ok, warning, error, working, unknown.
	LastTransitionTime string `json:"lastTransitionTime"`                                            // LastTransitionTime is the last time the condition transitioned from one status to another.
	Reason             string `json:"reason"`                                                        // Reason is a brief reason for the condition's last transition.
	Message            string `json:"message"`                                                       // Message is a human-readable message indicating details about the condition.
}

type KubernetesClusterStatusPrice struct {
	Monthly int `json:"monthly"` // Monthly is the monthly price of the cluster in your currency, e.g., "1000"
	Yearly  int `json:"yearly"`  // Yearly is the yearly price of the cluster, e.g., "12000"
}

type KubernetesClusterClusterDetails struct {
	ExternalId         string                                        `json:"externalId"`
	Resources          KubernetesClusterStatusClusterStatusResources `json:"resources"`
	Price              KubernetesClusterStatusPrice                  `json:"price"` // Price is the price of the cluster, e.g., "1000 NOK/month"
	ControlPlaneStatus KubernetesClusterControlPlaneStatus           `json:"controlplane"`
	NodePools          []KubernetesClusterNodePoolStatus             `json:"nodepools"` // TODO
}

type KubernetesClusterStatusClusterStatusResources struct {
	CPU    KubernetesClusterStatusClusterStatusResource `json:"cpu,omitzero"`    // CPU is the total CPU capacity of the cluster, if not specified in millicores, e.g., "16 cores", "8000 millicores"
	Memory KubernetesClusterStatusClusterStatusResource `json:"memory,omitzero"` // Memory is the total memory capacity of the cluster, if not specified in bytes, e.g., "64 GB", "128000 MB", "25600000000 bytes"
	GPU    KubernetesClusterStatusClusterStatusResource `json:"gpu,omitzero"`    // GPU is the total GPU capacity of the cluster, if not specified in number of GPUs"
	Disk   KubernetesClusterStatusClusterStatusResource `json:"disk,omitzero"`   // Disk is the total disk capacity of the cluster, if not specified in bytes"
}

type KubernetesClusterStatusClusterStatusResource struct {
	Capacity  Quantity `json:"capacity"`   // Capacity is the total capacity of the resource."
	Used      Quantity `json:"used"`       // Used is the amount of the resource that is currently used."
	Percetage int      `json:"percentage"` // Percentage is the percentage of the resource that is currently used as an int.
}

type KubernetesClusterControlPlaneStatus struct {
	Status       string                                        `json:"status"`
	Message      string                                        `json:"message"`
	Scale        int                                           `json:"scale"`        // Scale is the number of replicas of the control plane.
	MachineClass string                                        `json:"machineClass"` // MachineClass is the machine class of the control plane, e.g., "c5.large", "m5.xlarge"
	Resources    KubernetesClusterStatusClusterStatusResources `json:"resources"`    // Resources is the resources of the control plane, e.g., CPU, Memory, Disk, GPU
	Nodes        []string                                      `json:"nodes"`        // Nodes is the list of the uuids of the nodes in the control plane
}

type KubernetesClusterNodePoolStatus struct {
	Name         string                                        `json:"name"`
	Status       string                                        `json:"status"`
	Message      string                                        `json:"message"`
	Scale        int                                           `json:"scale"`        // Scale is the number of replicas of the nodepool.
	MachineClass string                                        `json:"machineClass"` // MachineClass is the machine class of the nodepool, e.g., "c5.large", "m5.xlarge"
	Autoscaling  KubernetesClusterAutoscalingConfig            `json:"autoscaling"`  // Autoscaling is the autoscaling configuration of the node pool.
	Resources    KubernetesClusterStatusClusterStatusResources `json:"resources"`    // Resources is the resources of the node pool, e.g., CPU, Memory, Disk, GPU
	Nodes        []string                                      `json:"nodes"`        // Nodes is the list of the uuids of the nodes in the node pool
}

type KubernetesClusterVersion struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Branch  string `json:"branch"`
}

type KubernetesClusterCondition struct {
	Type               string `json:"type" example:"ClusterReady"`                                   // Type is the type of the condition. For example, "ready", "available", etc.
	Status             string `json:"status"  example:"ok" enums:"ok,warning,error,working,unknown"` // Status is the status of the condition. Valid vales are: ok, warning, error, working, unknown.
	LastTransitionTime string `json:"lastTransitionTime"`                                            // LastTransitionTime is the last time the condition transitioned from one status to another.
	Reason             string `json:"reason"`                                                        // Reason is a brief reason for the condition's last transition.
	Message            string `json:"message"`                                                       // Message is a human-readable message indicating details about the condition.
}
