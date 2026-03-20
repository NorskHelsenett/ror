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
	CreatedAt          time.Time                         `json:"createdAt,omitempty" bson:"createdat,omitempty"`
	LastSeen           time.Time                         `json:"lastSeen,omitempty" bson:"lastseen,omitempty"`
}

type KubernetesClusterAgentStatusNodes struct {
	NodeCount     int                                               `json:"nodecount,omitempty" bson:"nodecount,omitempty"`
	NodePoolCount int                                               `json:"nodepoolcount,omitempty" bson:"nodepoolcount,omitempty"`
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
	Memory            int    `json:"memory,omitempty" bson:"memory,omitempty"`
	Architecture      string `json:"architecture,omitempty" bson:"architecture,omitempty"`
	KubernetesVersion string `json:"kubernetesVersion,omitempty" bson:"kubernetesversion,omitempty"`
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
