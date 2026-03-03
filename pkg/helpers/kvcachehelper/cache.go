package kvcachehelper

import (
	"context"
	"time"
)

type CacheInterface interface {
	Set(ctx context.Context, key string, value any, opts ...CacheSetOptions)
	Get(ctx context.Context, key string, opts ...CacheGetOptions) (any, bool)
	// Keys retrieves all keys currently stored in the cache.
	// It returns an error if the operation fails, such as due to a connection issue or an internal error.
	Keys(ctx context.Context, opts ...CacheKeysOptions) ([]string, error)
	Remove(ctx context.Context, key string, opts ...CacheRemoveOptions) bool
}

type CacheOptions struct {
	Prefix       string
	Timeout      time.Duration
	CronSchedule time.Duration
}

type CacheSetOptions struct {
	Prefix  string
	Timeout time.Duration
}

type CacheGetOptions struct {
	Prefix string
}

type CacheKeysOptions struct {
	Prefix string
}

type CacheRemoveOptions struct {
	Prefix string
}
