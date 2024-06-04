package rediscache

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/clients/redisdb"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type RedisCache struct {
	redisDb redisdb.RedisDB
}

func NewRedisCache(redisDb redisdb.RedisDB) *RedisCache {
	rlog.Debug("Creating new RedisCache")
	return &RedisCache{
		redisDb: redisDb,
	}
}

func (c *RedisCache) Set(ctx context.Context, key string, value string) {
	if key == "" {
		rlog.Warnc(ctx, "Key is empty")
		return
	}
	err := c.redisDb.Set(ctx, key, value)
	if err != nil {
		rlog.Debugc(ctx, fmt.Sprintf("Error adding value to redis cache by key: %s", key))
		return
	}
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, bool) {
	if key == "" {
		rlog.Warnc(ctx, "Key is empty")
		return "", false
	}
	var cacheValue string
	err := c.redisDb.Get(ctx, key, &cacheValue)
	if err != nil {
		rlog.Debugc(ctx, fmt.Sprintf("Error getting value from redis cache by key: %s", key))
		return "", false
	}

	if cacheValue == "" {
		return "", false
	}

	return cacheValue, true
}

func (c *RedisCache) Remove(ctx context.Context, key string) bool {
	if key == "" {
		rlog.Warnc(ctx, "Key is empty")
		return false
	}
	err := c.redisDb.Delete(ctx, key)
	if err != nil {
		rlog.Debugc(ctx, fmt.Sprintf("Error deleting value from redis cache by key: %s", key))
		return false
	}

	return true
}
