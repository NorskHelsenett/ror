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

func (rc rediscon) Delete(ctx context.Context, key string) error {
	err := rc.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
