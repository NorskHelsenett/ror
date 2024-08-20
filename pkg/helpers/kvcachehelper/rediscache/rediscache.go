package rediscache

import (
	"context"
	"fmt"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients/kvdbclient"
	"github.com/NorskHelsenett/ror/pkg/helpers/kvcachehelper"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type RedisCache struct {
	redisDb    kvdbclient.KvdbClient
	prefix     string
	expiration time.Duration
}

func NewRedisCache(redisDb kvdbclient.KvdbClient, opts ...kvcachehelper.CacheOptions) *RedisCache {
	rc := RedisCache{}
	rlog.Debug("Creating new RedisCache")
	if redisDb == nil {
		rlog.Error("RedisDB is nil", nil)
		return nil
	}

	if len(opts) == 1 {
		for _, opt := range opts {
			if opt.Timeout.Seconds() != 0 {
				rc.expiration = opt.Timeout
			}
			if opt.Prefix != "" {
				rc.prefix = opt.Prefix
			}

		}
	}

	if rc.expiration.Seconds() == 0 {
		rc.expiration = 6 * time.Hour
	}
	rc.redisDb = redisDb

	return &rc
}

func (c *RedisCache) Set(ctx context.Context, key string, value string) {
	if key == "" {
		rlog.Warnc(ctx, "Key is empty")
		return
	}

	rkey := c.prefix + key
	err := c.redisDb.Set(ctx, rkey, value, c.expiration)
	if err != nil {
		rlog.Debugc(ctx, fmt.Sprintf("Error adding value to redis cache by key: %s", rkey))
		return
	}
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, bool) {
	if key == "" {
		rlog.Warnc(ctx, "Key is empty")
		return "", false
	}
	var cacheValue string
	rkey := c.prefix + key
	err := c.redisDb.Get(ctx, rkey, &cacheValue)
	if err != nil {
		rlog.Debugc(ctx, fmt.Sprintf("Error getting value from redis cache by key: %s", rkey))
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
	rkey := c.prefix + key
	err := c.redisDb.Delete(ctx, rkey)
	if err != nil {
		rlog.Debugc(ctx, fmt.Sprintf("Error deleting value from redis cache by key: %s", rkey))
		return false
	}

	return true
}
