package rortypes

import (
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
	Connected bool `json:"connected"`
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
