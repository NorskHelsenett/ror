package valkey

import (
	"context"
	"errors"
)

// Deprecated: ReJSON is no longer supported
func (vc valkeycon) GetJSON(ctx context.Context, key string, path string, output interface{}) error {
	return errors.New("ReJSON is no longer supported")
}

// Deprecated: ReJSON is no longer supported
func (vc valkeycon) SetJSON(ctx context.Context, key string, path string, value interface{}) error {
	return errors.New("ReJSON is no longer supported")

}
