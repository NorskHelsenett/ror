package rediscache

import (
	"context"
	"fmt"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients/redisdb"
	"github.com/NorskHelsenett/ror/pkg/helpers/kvcachehelper"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type RedisCache struct {
	redisDb    redisdb.RedisDB
	prefix     string
	expiration time.Duration
}

func NewRedisCache(redisDb redisdb.RedisDB, opts ...kvcachehelper.CacheOptions) *RedisCache {
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

func (c *RedisCache) Set(ctx context.Context, key string, value any, opts ...kvcachehelper.CacheSetOptions) {
	var expiresIn time.Duration = c.expiration
	var prefix string = c.prefix
	for _, opt := range opts {
		if opt.Timeout.Seconds() != 0 {
			expiresIn = opt.Timeout
		}
		if opt.Prefix != "" {
			prefix = opt.Prefix
		}
	}

	if key == "" {
		rlog.Warnc(ctx, "Key is empty")
		return
	}

	rkey := prefix + key
	err := c.redisDb.Set(ctx, rkey, value, expiresIn)
	if err != nil {
		rlog.Debugc(ctx, fmt.Sprintf("Error adding value to redis cache by key: %s", rkey))
		return
	}
}

func (c *RedisCache) Get(ctx context.Context, key string, opts ...kvcachehelper.CacheGetOptions) (any, bool) {
	var prefix string = c.prefix
	for _, opt := range opts {
		if opt.Prefix != "" {
			prefix = opt.Prefix
		}
	}
	if key == "" {
		rlog.Warnc(ctx, "Key is empty")
		return "", false
	}
	var cacheValue string
	rkey := prefix + key
	err := c.redisDb.Get(ctx, rkey, &cacheValue)
	if err != nil {
		rlog.Debugc(ctx, fmt.Sprintf("Error getting value from redis cache by key: %s", rkey))
		return "", false
	}

	return cacheValue, true
}

func (c *RedisCache) Keys(ctx context.Context) ([]string, error) {
	if c.redisDb == nil {
		rlog.Error("RedisDB is nil", nil)
		return nil, fmt.Errorf("RedisDB is nil")
	}
	keys, err := c.redisDb.Keys(ctx)
	if err != nil {
		rlog.Debugc(ctx, fmt.Sprintf("Error getting keys from redis cache: %s", err))
		return nil, err
	}

	for i := range keys {
		keys[i] = keys[i][len(c.prefix):]
	}

	return keys, nil
}

func (c *RedisCache) Remove(ctx context.Context, key string, opts ...kvcachehelper.CacheRemoveOptions) bool {
	prefix := c.prefix
	for _, opt := range opts {
		if opt.Prefix != "" {
			prefix = opt.Prefix
		}
	}
	if key == "" {
		rlog.Warnc(ctx, "Key is empty")
		return false
	}
	rkey := prefix + key
	err := c.redisDb.Delete(ctx, rkey)
	if err != nil {
		rlog.Debugc(ctx, fmt.Sprintf("Error deleting value from redis cache by key: %s", rkey))
		return false
	}

	return true
}
