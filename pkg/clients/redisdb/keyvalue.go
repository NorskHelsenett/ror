package redisdb

import (
	"context"
	"time"
)

func (rc rediscon) Get(ctx context.Context, key string, output *string) error {
	var result string
	result, err := rc.Client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	*output = result
	return nil

}

func (rc rediscon) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := rc.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rc rediscon) Keys(ctx context.Context) ([]string, error) {
	keys, err := rc.Client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (rc rediscon) Delete(ctx context.Context, key string) error {
	err := rc.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rc rediscon) MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	return rc.Client.MGet(ctx, keys...).Result()
}

func (rc rediscon) SetPipelined(ctx context.Context, items []SetItem) error {
	pipe := rc.Client.Pipeline()
	for _, item := range items {
		pipe.Set(ctx, item.Key, item.Value, item.Expiration)
	}
	_, err := pipe.Exec(ctx)
	return err
}
