package kvdbclient

import (
	"context"
	"time"

	"github.com/dotse/go-health"
)

type KvdbClient interface {
	Get(ctx context.Context, key string, output *string) error
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
	CheckHealth() []health.Check
}
