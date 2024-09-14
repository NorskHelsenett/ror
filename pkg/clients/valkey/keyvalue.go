package valkey

import (
	"context"
	"errors"
	"time"
)

func (vc valkeycon) Get(ctx context.Context, key string, output *string) error {
	result, err := vc.Client.Do(ctx, vc.Client.B().Get().Key(key).Build()).ToString()
	if err != nil {
		return err
	}
	*output = result
	return nil
}

func (vc valkeycon) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	setstr, ok := value.(string)
	if !ok {
		return errors.New("value is not a string")
	}
	err := vc.Client.Do(ctx, vc.Client.B().Set().Key(key).Value(setstr).Ex(expiration).Build()).Error()
	if err != nil {
		return err
	}
	return nil
}

func (vc valkeycon) Delete(ctx context.Context, key string) error {
	err := vc.Client.Do(ctx, vc.Client.B().Del().Key(key).Build()).Error()
	if err != nil {
		return err
	}
	return nil
}
