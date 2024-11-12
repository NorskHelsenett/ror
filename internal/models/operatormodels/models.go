package operatormodels

import "github.com/NorskHelsenett/ror/pkg/models/providers"

type ClusterInfo struct {
	Id             string                 `json:"id"`
	ClusterName    string                 `json:"clusterName"`
	DatacenterName string                 `json:"datacenterName"`
	WorkspaceName  string                 `json:"workspaceName"`
	Provider       providers.ProviderType `json:"provider"`
}
