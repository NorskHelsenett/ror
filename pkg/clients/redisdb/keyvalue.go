package redisdb

import "context"

func (rc rediscon) Get(ctx context.Context, key string, output *string) error {
	var result string
	result, err := rc.Client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	*output = result
	return nil

}

func (rc rediscon) Set(ctx context.Context, key string, value interface{}) error {
	err := rc.Client.Set(ctx, key, value, 0).Err()
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
