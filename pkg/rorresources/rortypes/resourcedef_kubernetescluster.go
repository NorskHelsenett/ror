package rortypes

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	vitiv1alpha1 "github.com/vitistack/common/pkg/v1alpha1"
)

// ResourceKubernetesCluster uses the external types from vitistack/common
type ResourceKubernetesCluster struct {
	Spec   ResourceKubernetesClusterSpec   `json:"spec"`
	Status ResourceKubernetesClusterStatus `json:"status"`
}

type ResourceKubernetesClusterSpec struct {
	SlackChannels []string              `json:"slackChannels"`
	VitiSpec      KubernetesClusterSpec `json:"vitiSpec,omitzero" bson:"vitispec,omitempty"`
}

type ResourceKubernetesClusterStatus struct {
	ProviderStatus KubernetesClusterStatus      `json:"providerstatus,omitzero" bson:"providerstatus,omitempty"`
	AgentStatus    KubernetesClusterAgentStatus `json:"agentstatus,omitzero" bson:"agentstatus,omitempty"`
}

type KubernetesClusterAgentStatus struct {
	ClusterId          string                            `json:"clusterId,omitempty" bson:"clusterid,omitempty"`
	ClusterName        string                            `json:"clusterName,omitempty" bson:"clustername,omitempty"`
	KubernetesProvider providermodels.ProviderType       `json:"kubernetesProvider,omitempty" bson:"kubernetesprovider,omitempty"`
	Az                 string                            `json:"az,omitempty" bson:"az,omitempty"`
	Region             string                            `json:"region,omitempty" bson:"region,omitempty"`
	Country            string                            `json:"country,omitempty" bson:"country,omitempty"`
	Workspace          string                            `json:"workspaceId,omitempty" bson:"workspaceid,omitempty"`
	Environment        string                            `json:"environment,omitempty" bson:"environment,omitempty"`
	Datacenter         string                            `json:"datacenter,omitempty" bson:"datacenter,omitempty"`
	Nodes              KubernetesClusterAgentStatusNodes `json:"nodes,omitzero" bson:"nodes,omitempty"`
	Versions           map[string]string                 `json:"versions,omitempty" bson:"versions,omitempty"`
	Urls               map[string]string                 `json:"urls,omitempty" bson:"urls,omitempty"`
	CreatedAt          time.Time                         `json:"createdAt,omitempty" bson:"createdat,omitempty"`
	LastSeen           time.Time                         `json:"lastSeen,omitempty" bson:"lastseen,omitempty"`
}

type KubernetesClusterAgentStatusNodes struct {
	ControllPlane []KubernetesClusterAgentStatusNodesNodepoolsNodes `json:"controlPlane,omitzero" bson:"controlplane,omitempty"`
	Nodepools     []KubernetesClusterAgentStatusNodesNodepools      `json:"nodepools,omitzero" bson:"nodepools,omitempty"`
}
type KubernetesClusterAgentStatusNodesNodepools struct {
	Name  string                                            `json:"name,omitempty" bson:"name,omitempty"`
	Nodes []KubernetesClusterAgentStatusNodesNodepoolsNodes `json:"nodes,omitzero" bson:"nodes,omitempty"`
}
type KubernetesClusterAgentStatusNodesNodepoolsNodes struct {
	Name              string   `json:"name,omitempty" bson:"name,omitempty"`
	Cpu               Quantity `json:"cpu,omitzero" bson:"cpu,omitempty"`
	Memory            Quantity `json:"memory,omitzero" bson:"memory,omitempty"`
	Architecture      string   `json:"architecture,omitempty" bson:"architecture,omitempty"`
	KubernetesVersion string   `json:"kubernetesVersion,omitempty" bson:"kubernetesversion,omitempty"`
}

func (r *KubernetesClusterAgentStatus) GetNodepoolCount() int {
	return len(r.Nodes.Nodepools)
}

func (r *KubernetesClusterAgentStatus) GetNodeCount() int {
	nodeCount := 0
	for _, nodepool := range r.Nodes.Nodepools {
		nodeCount += len(nodepool.Nodes)
	}
	return nodeCount
}

func (r *KubernetesClusterAgentStatus) GetKubernetesVersion() string {
	if len(r.Nodes.ControllPlane) > 0 && r.Nodes.ControllPlane[0].KubernetesVersion != "" {
		return r.Nodes.ControllPlane[0].KubernetesVersion
	}
	return "Unknown"
}
func (r *KubernetesClusterAgentStatus) GetVersionByKey(key string) string {
	version, _ := r.Versions[key]
	return nomalizeVersion(version)
}

// nomalizeVersion returns "Unknown" if the version string is empty, otherwise it returns the version string ensuring it starts with a v

func nomalizeVersion(version string) string {
	if version == "" {
		return "Unknown"
	}
	if version[0] != 'v' && version[0] != 'V' {
		return "v" + version
	}
	return version
}

func (r *KubernetesClusterAgentStatus) GetUrlByKey(key string) string {
	if url, ok := r.Urls[key]; ok {
		return url
	}
	return "Unknown"
}

func (r *KubernetesClusterAgentStatus) GetStatus() string {
	diff := time.Since(r.LastSeen)
	if diff < 5*time.Minute {
		return "ok"
	} else if diff < 15*time.Minute {
		return "warning"
	} else {
		return "error"
	}
}

// GetTotalCpu returns the total CPU resources of the cluster by summing the CPU of all nodes in the control plane and node pools.
func (r *KubernetesClusterAgentStatus) GetTotalCpu() *Quantity {
	cpu := &Quantity{}
	cpu.Add(r.GetControlPlaneCpu().Quantity)
	cpu.Add(r.GetNodePoolCpu().Quantity)
	return cpu
}

// GetTotalMemory returns the total memory resources of the cluster by summing the memory of all nodes in the control plane and node pools.
func (r *KubernetesClusterAgentStatus) GetTotalMemory() *Quantity {
	memory := &Quantity{}
	memory.Add(r.GetControlPlaneMemory().Quantity)
	memory.Add(r.GetNodePoolMemory().Quantity)
	return memory
}

func (r *KubernetesClusterAgentStatus) GetControlPlaneCpu() *Quantity {
	cpu := &Quantity{}
	for _, node := range r.Nodes.ControllPlane {
		cpu.Add(node.Cpu.Quantity)
	}
	return cpu
}

func (r *KubernetesClusterAgentStatus) GetControlPlaneMemory() *Quantity {
	memory := &Quantity{}
	for _, node := range r.Nodes.ControllPlane {
		memory.Add(node.Memory.Quantity)
	}
	return memory
}

func (r *KubernetesClusterAgentStatus) GetNodePoolCpu() *Quantity {
	cpu := &Quantity{}
	for _, nodepool := range r.Nodes.Nodepools {
		for _, node := range nodepool.Nodes {
			cpu.Add(node.Cpu.Quantity)
		}
	}
	return cpu
}

func (r *KubernetesClusterAgentStatus) GetNodePoolMemory() *Quantity {
	memory := &Quantity{}
	for _, nodepool := range r.Nodes.Nodepools {
		for _, node := range nodepool.Nodes {
			memory.Add(node.Memory.Quantity)
		}
	}
	return memory
}

// Type aliases for convenience and backward compatibility
type KubernetesClusterSpec = vitiv1alpha1.KubernetesClusterSpec
type KubernetesClusterSpecData = vitiv1alpha1.KubernetesClusterSpecData
type KubernetesClusterSpecTopology = vitiv1alpha1.KubernetesClusterSpecTopology
type KubernetesClusterSpecControlPlane = vitiv1alpha1.KubernetesClusterSpecControlPlane
type KubernetesClusterSpecMetadataDetails = vitiv1alpha1.KubernetesClusterSpecMetadataDetails
type KubernetesClusterStorage = vitiv1alpha1.KubernetesClusterStorage
type KubernetesClusterWorkers = vitiv1alpha1.KubernetesClusterWorkers
type KubernetesClusterNodePool = vitiv1alpha1.KubernetesClusterNodePool
type KubernetesClusterTaint = vitiv1alpha1.KubernetesClusterTaint
type KubernetesClusterAutoscalingConfig = vitiv1alpha1.KubernetesClusterAutoscalingConfig
type KubernetesClusterAutoscalingSpec = vitiv1alpha1.KubernetesClusterAutoscalingSpec
type KubernetesClusterStatus = vitiv1alpha1.KubernetesClusterStatus
type KubernetesClusterClusterState = vitiv1alpha1.KubernetesClusterClusterState
type KubernetesClusterEndpoint = vitiv1alpha1.KubernetesClusterEndpoint
type KubernetesClusterStatusCondition = vitiv1alpha1.KubernetesClusterStatusCondition
type KubernetesClusterStatusPrice = vitiv1alpha1.KubernetesClusterStatusPrice
type KubernetesClusterClusterDetails = vitiv1alpha1.KubernetesClusterClusterDetails
type KubernetesClusterStatusClusterStatusResources = vitiv1alpha1.KubernetesClusterStatusClusterStatusResources
type KubernetesClusterStatusClusterStatusResource = vitiv1alpha1.KubernetesClusterStatusClusterStatusResource
type KubernetesClusterControlPlaneStatus = vitiv1alpha1.KubernetesClusterControlPlaneStatus
type KubernetesClusterNodePoolStatus = vitiv1alpha1.KubernetesClusterNodePoolStatus
type KubernetesClusterVersion = vitiv1alpha1.KubernetesClusterVersion
type KubernetesClusterCondition = vitiv1alpha1.KubernetesClusterCondition
