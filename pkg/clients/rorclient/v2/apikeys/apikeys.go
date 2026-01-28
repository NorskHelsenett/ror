package apikeys

import "github.com/NorskHelsenett/ror/pkg/apicontracts/apikeystypes/v2"

type ApiKeysInterface interface {
	RegisterAgent(data apikeystypes.RegisterClusterRequest) (apikeystypes.RegisterClusterResponse, error)
}
