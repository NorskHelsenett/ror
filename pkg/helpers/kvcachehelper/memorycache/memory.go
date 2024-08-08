package memorycache

import (
	"context"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/kvcachehelper"
)

// Path: cache/cache.go

type CacheValue struct {
	// The cache's expiration time
	ExpirationTime time.Time
	// The cache's value
	Value string
}

type KvCache struct {
	lock      sync.RWMutex
	values    map[string]CacheValue
	prefix    string
	expiresIn time.Duration
}
type CacheOption map[string]interface{}

// NewKvCache instantiates a new Cache and sets the default values.
// The default expiration time is set to 6 hours from the current time.
// The expiration time can be overridden by passing an expirationTime option.
// Example:
// cache := NewKvCache(ExpirationTime: time.Now().Add(1 * time.Hour))
func NewKvCache(opts ...kvcachehelper.CacheOptions) *KvCache {
	c := &KvCache{values: make(map[string]CacheValue)}
	if len(opts) == 1 {
		for _, opt := range opts {
			if opt.Timeout.Seconds() != 0 {
				c.expiresIn = opt.Timeout
			}
			if opt.Prefix != "" {
				c.prefix = opt.Prefix
			}

		}
	}

	if c.expiresIn.Seconds() == 0 {
		c.expiresIn = 6 * time.Hour
	}
	return c
}

// Set adds a new key-value pair to the cache.
// If the key already exists, it will be overwritten.
func (c *KvCache) Set(ctx context.Context, key string, value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.values[c.prefix+key] = CacheValue{
		ExpirationTime: time.Now().Add(c.expiresIn),
		Value:          value,
	}
}

// Get retrieves a value from the cache.
// If the key does not exist or the value has expired, it will return false.
func (c *KvCache) Get(ctx context.Context, key string) (string, bool) {
	if val, ok := c.values[c.prefix+key]; ok {
		if time.Now().After(val.ExpirationTime) {
			c.lock.Lock()
			defer c.lock.Unlock()
			delete(c.values, key)
			return "", false
		}
		c.lock.RLock()
		defer c.lock.RUnlock()
		return val.Value, true
	}
	return "", false
}

// Remove removes a key-value pair from the cache.
func (c *KvCache) Remove(ctx context.Context, key string) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.values, key)
	return true
}
