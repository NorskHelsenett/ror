package kvcachehelper

import "context"

type CacheInterface interface {
	Set(ctx context.Context, key string, value string)
	Get(ctx context.Context, key string) (string, bool)
	Remove(ctx context.Context, key string) bool
}
