package rorclientv2self

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2self"
)

type SelfInterface interface {
	Get(ctx context.Context) (apicontractsv2self.SelfData, error)
	CreateOrUpdateApiKey(ctx context.Context, name string, ttl int64) (string, error)
}
