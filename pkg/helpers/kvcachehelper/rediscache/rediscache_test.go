package rediscache

import (
	"context"
	"reflect"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/clients/redisdb"
	"github.com/NorskHelsenett/ror/pkg/helpers/kvcachehelper/memorycache"
	"github.com/dotse/go-health"
)

func TestNewRedisCache(t *testing.T) {
	type args struct {
		redisDb redisdb.RedisDB
	}
	tests := []struct {
		name string
		args args
		want *RedisCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedisCache(tt.args.redisDb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedisCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
