package kvcachehelper

import (
	"context"
	"time"
)

type CacheInterface interface {
	Set(ctx context.Context, key string, value string)
	Get(ctx context.Context, key string) (string, bool)
	// Keys retrieves all keys currently stored in the cache.
	// It returns an error if the operation fails, such as due to a connection issue or an internal error.
	Keys(ctx context.Context) ([]string, error)
	Remove(ctx context.Context, key string) bool
}

type CacheOptions struct {
	Prefix       string
	Timeout      time.Duration
	CronSchedule time.Duration
}
