package redisdb

import (
	"context"
	"encoding/json"

	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson/v4"
)

func (rc rediscon) GetJSON(ctx context.Context, key string, path string, output interface{}) error {
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClientWithContext(ctx, rc.Client)
	value, err := redis.Bytes(rh.JSONGet(key, path))
	if err != nil {
		return err
	}
	err = json.Unmarshal(value, output)
	if err != nil {
		return err
	}
	return nil
}

func (rc rediscon) SetJSON(ctx context.Context, key string, path string, value interface{}) error {
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClientWithContext(ctx, rc.Client)

	_, err := rh.JSONSet(key, path, value)
	if err != nil {
		return err
	}
	return nil

}
