package apikeys

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apikeystypes/v2"
)

type ApiKeysInterface interface {
	RegisterAgent(ctx context.Context, data apikeystypes.RegisterClusterRequest) (apikeystypes.RegisterClusterResponse, error)
}
