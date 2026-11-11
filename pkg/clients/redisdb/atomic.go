package redisdb

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

// AtomicRedisDB adds bounded reads and persistent compare-and-swap writes.
type AtomicRedisDB interface {
	RedisDB
	GetBounded(ctx context.Context, key string, maxBytes int) (string, error)
	CompareAndSwap(ctx context.Context, key string, expected *string, value string) (bool, error)
}

var boundedGet = redis.NewScript(`
local size = redis.call('STRLEN', KEYS[1])
if size > tonumber(ARGV[1]) then
  return redis.error_reply('stored value exceeds byte limit')
end
return redis.call('GET', KEYS[1])
`)

func (rc rediscon) GetBounded(ctx context.Context, key string, maxBytes int) (string, error) {
	if maxBytes <= 0 {
		return "", errors.New("read byte limit must be positive")
	}
	return boundedGet.Run(ctx, rc.Client, []string{key}, maxBytes).Text()
}

var compareAndSwap = redis.NewScript(`
local current = redis.call('GET', KEYS[1])
if ARGV[1] == 'absent' then
  if current then return 0 end
elseif not current or current ~= ARGV[2] then
  return 0
end
redis.call('SET', KEYS[1], ARGV[3])
return 1
`)

// CompareAndSwap writes without expiry only if the expected bytes still match.
// A nil expected value requires that the key does not exist.
func (rc rediscon) CompareAndSwap(ctx context.Context, key string, expected *string, value string) (bool, error) {
	mode, previous := "absent", ""
	if expected != nil {
		mode, previous = "present", *expected
	}
	result, err := compareAndSwap.Run(ctx, rc.Client, []string{key}, mode, previous, value).Int()
	return result == 1, err
}
