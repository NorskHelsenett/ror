package memorycache

import (
	"sync"
	"time"
)

// Path: cache/cache.go

type CacheValue struct {
	// The cache's expiration time
	ExpirationTime time.Time
	// The cache's value
	Value string
}

type KvCache struct {
	lock           sync.RWMutex
	values         map[string]CacheValue
	expirationTime time.Time
}
type CacheOption map[string]interface{}

// NewKvCache instantiates a new Cache and sets the default values.
// The default expiration time is set to 6 hours from the current time.
// The expiration time can be overridden by passing an expirationTime option.
// Example:
// cache := NewKvCache(ExpirationTime: time.Now().Add(1 * time.Hour))
func NewKvCache(opts ...CacheOption) *KvCache {
	c := &KvCache{values: make(map[string]CacheValue)}
	if len(opts) > 0 {
		for _, opt := range opts {
			if expirationTime, ok := opt["expirationTime"].(time.Time); ok {
				c.expirationTime = expirationTime
			}
		}
	}

	if c.expirationTime.IsZero() {
		c.expirationTime = time.Now().Add(6 * time.Hour)
	}
	return c
}

// Add adds a new key-value pair to the cache.
// If the key already exists, it will be overwritten.
func (c *KvCache) Add(key string, value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.values[key] = CacheValue{
		ExpirationTime: c.expirationTime,
		Value:          value,
	}
}

// Get retrieves a value from the cache.
// If the key does not exist or the value has expired, it will return false.
func (c *KvCache) Get(key string) (string, bool) {
	if val, ok := c.values[key]; ok {
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
func (c *KvCache) Remove(key string) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.values, key)
	return true
}
