package rediscache

import (
	"context"
	"fmt"
	"sync"

	"github.com/NorskHelsenett/ror/pkg/clients/redisdb"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type RedisCache struct {
	lock    sync.RWMutex
	redisDb redisdb.RedisDB
}

func NewRedisCache(redisDb redisdb.RedisDB) *RedisCache {
	rlog.Debug("Creating new RedisCache")
	return &RedisCache{
		redisDb: redisDb,
	}
}

func (c *RedisCache) Add(key string, value string) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	err := c.redisDb.Set(context.Background(), key, value)
	if err != nil {
		rlog.Error(fmt.Sprintf("Error adding value to redis cache by key: %s", key), nil)
		return
	}
}

func (c *RedisCache) Get(key string) (string, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	var cacheValue string
	err := c.redisDb.Get(context.Background(), key, &cacheValue)
	if err != nil {
		rlog.Error(fmt.Sprintf("Error getting value from redis cache by key: %s", key), nil)
		return "", false
	}
	return cacheValue, true
}

func (c *RedisCache) Remove(key string) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	err := c.redisDb.Delete(context.Background(), key)
	if err != nil {
		rlog.Error(fmt.Sprintf("Error deleting value from redis cache by key: %s", key), nil)
		return false
	}

	return true
}
