package token

import (
	"context"
)

type TokenInterface interface {
	Exchange(ctx context.Context, token string, clusterId string, admin bool) (string, error)
}
