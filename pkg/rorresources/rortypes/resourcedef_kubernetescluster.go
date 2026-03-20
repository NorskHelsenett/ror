package rortypes

import (
	"fmt"
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
	Name              string `json:"name,omitempty" bson:"name,omitempty"`
	Cpu               int    `json:"cpu,omitempty" bson:"cpu,omitempty"`
	Memory            int64  `json:"memory,omitempty" bson:"memory,omitempty"`
	Architecture      string `json:"architecture,omitempty" bson:"architecture,omitempty"`
	KubernetesVersion string `json:"kubernetesVersion,omitempty" bson:"kubernetesversion,omitempty"`
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
	if version, ok := r.Versions[key]; ok {
		return version
	}
	return "Unknown"
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

func (r *KubernetesClusterAgentStatus) GetCpu() int {
	cpu := 0
	for _, nodepool := range r.Nodes.Nodepools {
		for _, node := range nodepool.Nodes {
			cpu += node.Cpu
		}
	}
	return cpu
}

func (r *KubernetesClusterAgentStatus) GetMemory() string {
	var memory int64 = 0
	for _, nodepool := range r.Nodes.Nodepools {
		for _, node := range nodepool.Nodes {
			memory += node.Memory
		}
	}
	return formatMemory(memory)
}

func formatMemory(bytes int64) string {
	const (
		KiB = 1024
		MiB = 1024 * KiB
		GiB = 1024 * MiB
		TiB = 1024 * GiB
	)
	switch {
	case bytes >= TiB:
		return fmt.Sprintf("%.1f TiB", float64(bytes)/float64(TiB))
	case bytes >= GiB:
		return fmt.Sprintf("%.1f GiB", float64(bytes)/float64(GiB))
	case bytes >= MiB:
		return fmt.Sprintf("%.1f MiB", float64(bytes)/float64(MiB))
	case bytes >= KiB:
		return fmt.Sprintf("%.1f KiB", float64(bytes)/float64(KiB))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
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
