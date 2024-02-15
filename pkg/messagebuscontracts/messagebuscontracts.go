package messagebuscontracts

import (
	"ror/internal/models/tanzu"
	"ror/pkg/apicontracts/apiresourcecontracts"
)

type ClusterCreatedEvent struct {
	EventBase
	EventClusterBase
	ClusterName   string `json:"clusterName"`
	WorkspaceName string `json:"workspaceName"`
}

type ClusterUpdatedEvent struct {
	EventBase
	EventClusterBase
}

type AclUpdateEvent struct {
	EventBase
	Action string `json:"action"`
}

type ResourceUpdatedEvent struct {
	ResourceNamespace apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNamespace] `json:"resourceNamespace"`
}

type ClusterCreate struct {
	ClusterInput tanzu.TanzuKubernetesClusterInput `json:"clusterInput"`
}

type ClusterModify struct {
	ClusterInput tanzu.TanzuKubernetesClusterInput `json:"clusterInput"`
}

type ClusterDelete struct {
	ClusterInput tanzu.TanzuKubernetesClusterInput `json:"clusterInput"`
}

type OperatorOrder struct {
	TanzuKubernetesCluster apiresourcecontracts.ResourceTanzuKubernetesCluster `json:"tanzuKubernetesCluster"`
}
