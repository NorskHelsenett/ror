package apikeys

import "github.com/NorskHelsenett/ror/pkg/apicontracts/clustersapi/v2"

type ApiKeysInterface interface {
	RegisterAgent(data clustersapi.RegisterClusterRequest) (clustersapi.RegisterClusterResponse, error)
}
