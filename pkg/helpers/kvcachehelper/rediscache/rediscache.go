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
		rlog.Error("Key is empty", nil)
		return
	}
	err := c.redisDb.Set(context.Background(), key, value)
	if err != nil {
		rlog.Error(fmt.Sprintf("Error adding value to redis cache by key: %s", key), nil)
		return
	}
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, bool) {
	if key == "" {
		rlog.Error("Key is empty", nil)
		return "", false
	}
	var cacheValue string
	err := c.redisDb.Get(context.Background(), key, &cacheValue)
	if err != nil {
		rlog.Error(fmt.Sprintf("Error getting value from redis cache by key: %s", key), nil)
		return "", false
	}

	if cacheValue == "" {
		return "", false
	}

	return cacheValue, true
}

func (c *RedisCache) Remove(ctx context.Context, key string) bool {
	if key == "" {
		rlog.Error("Key is empty", nil)
		return false
	}
	err := c.redisDb.Delete(context.Background(), key)
	if err != nil {
		rlog.Error(fmt.Sprintf("Error deleting value from redis cache by key: %s", key), nil)
		return false
	}

	return true
}
