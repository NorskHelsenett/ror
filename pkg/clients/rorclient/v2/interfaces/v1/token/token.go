package token

import (
	"context"
)

type TokenInterface interface {
	GetJWKS(ctx context.Context) (string, error)
}
