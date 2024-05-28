package rediscache

import (
	"context"
	"errors"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/helpers/kvcachehelper/memorycache"
	"github.com/dotse/go-health"
)

type mockRedisDB struct {
	memorycache *memorycache.KvCache
}

func NewMockRedisDb() *mockRedisDB {
	return &mockRedisDB{
		memorycache: memorycache.NewKvCache(),
	}
}

func (m *mockRedisDB) Get(ctx context.Context, key string, output *string) error {
	value, exists := m.memorycache.Get(ctx, key)
	if !exists {
		return nil
	}
	*output = value
	return nil
}

func (m *mockRedisDB) Set(ctx context.Context, key string, value interface{}) error {
	m.memorycache.Set(ctx, key, value.(string))
	return nil
}

func (m *mockRedisDB) Delete(ctx context.Context, key string) error {
	if key == "" {
		return errors.New("Key is empty")
	}
	if _, exists := m.memorycache.Get(ctx, key); !exists {
		return errors.New("Key does not exist")
	}
	m.memorycache.Remove(ctx, key)
	return nil
}

func (m *mockRedisDB) GetJSON(context.Context, string, string, interface{}) error {
	m.memorycache.Get(context.Background(), "test-key")
	return nil
}

func (m *mockRedisDB) SetJSON(ctx context.Context, key string, path string, value interface{}) error {
	m.memorycache.Set(ctx, key, value.(string))
	return nil
}

func (m *mockRedisDB) CheckHealth() []health.Check {
	results := []health.Check{}
	return results
}

func (m *mockRedisDB) Ping() bool {
	return true
}

func TestNewRedisCache(t *testing.T) {
	redisDb := &mockRedisDB{}
	cache := NewRedisCache(redisDb)

	if cache.redisDb != redisDb {
		t.Errorf("Expected RedisDB to be set to mockRedisDB, but got %v", cache.redisDb)
	}
}

func TestRedisCache_Get(t *testing.T) {
	redisDb := NewMockRedisDb()
	cache := NewRedisCache(redisDb)

	// Test case 1: Key exists in cache
	expectedValue := "test-value"
	err := redisDb.Set(context.Background(), "test-key", expectedValue)
	if err != nil {
		t.Errorf("Failed to set value in Redis cache: %v", err)
	}

	value, exists := cache.Get(context.Background(), "test-key")
	if !exists {
		t.Errorf("Expected key 'test-key' to exist in cache, but it doesn't")
	}
	if value != expectedValue {
		t.Errorf("Expected value '%s' for key 'test-key', but got '%s'", expectedValue, value)
	}

	// Test case 2: Key does not exist in cache
	value, exists = cache.Get(context.Background(), "non-existent-key")
	if exists {
		t.Errorf("Expected key 'non-existent-key' to not exist in cache, but it does")
	}
	if value != "" {
		t.Errorf("Expected empty value for non-existent key, but got '%s'", value)
	}
}

func TestRedisCache_Set(t *testing.T) {
	redisDb := NewMockRedisDb()
	cache := NewRedisCache(redisDb)

	// Test case 1: Set value in cache
	key := "test-key"
	value := "test-value"
	cache.Set(context.Background(), key, value)

	// Verify that the value is set in the cache
	cacheValue, exists := cache.Get(context.Background(), key)
	if !exists {
		t.Errorf("Expected key '%s' to exist in cache, but it doesn't", key)
	}
	if cacheValue != value {
		t.Errorf("Expected value '%s' for key '%s', but got '%s'", value, key, cacheValue)
	}

	// Test case 2: Set value with empty key
	emptyKey := ""
	emptyValue := "empty-value"
	cache.Set(context.Background(), emptyKey, emptyValue)

	// Verify that the value is not set in the cache
	emptyCacheValue, exists := cache.Get(context.Background(), emptyKey)
	if exists {
		t.Errorf("Expected key '%s' to not exist in cache, but it does", emptyKey)
	}
	if emptyCacheValue != "" {
		t.Errorf("Expected empty value for key '%s', but got '%s'", emptyKey, emptyCacheValue)
	}
}

func TestRedisCache_Remove(t *testing.T) {
	redisDb := NewMockRedisDb()
	cache := NewRedisCache(redisDb)

	// Test case 1: Remove existing key
	key := "test-key"
	value := "test-value"
	err := redisDb.Set(context.Background(), key, value)
	if err != nil {
		t.Errorf("Failed to set value in Redis cache: %v", err)
	}

	removed := cache.Remove(context.Background(), key)
	if !removed {
		t.Errorf("Expected key '%s' to be removed from cache, but it wasn't", key)
	}

	_, exists := cache.Get(context.Background(), key)
	if exists {
		t.Errorf("Expected key '%s' to not exist in cache after removal, but it does", key)
	}

	// Test case 2: Remove non-existent key
	nonExistentKey := "non-existent-key"
	removed = cache.Remove(context.Background(), nonExistentKey)
	if removed {
		t.Errorf("Expected key '%s' to not be removed from cache, but it was", nonExistentKey)
	}
}
