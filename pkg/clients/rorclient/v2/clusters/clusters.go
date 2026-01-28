package clusters

import "github.com/NorskHelsenett/ror/pkg/apicontracts/clustersapi/v2"

type ClustersInterface interface {
	Register(data clustersapi.RegisterClusterRequest) (clustersapi.RegisterClusterResponse, error)
}
