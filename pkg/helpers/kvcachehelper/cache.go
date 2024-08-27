package kvcachehelper

import (
	"context"
	"time"
)

type CacheInterface interface {
	Set(ctx context.Context, key string, value string)
	Get(ctx context.Context, key string) (string, bool)
	Remove(ctx context.Context, key string) bool
}

type CacheOptions struct {
	Prefix       string
	Timeout      time.Duration
	CronSchedule time.Duration
}
