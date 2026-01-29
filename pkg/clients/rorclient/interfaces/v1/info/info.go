package info

import "context"

type InfoInterface interface {
	GetVersion(ctx context.Context) (string, error)
}
