// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package rortypes

import "context"

// CommonResourceInterface represents the minimum interface for all resources
type CommonResourceInterface interface {
	GetRorHash() string
	ApplyInputFilter(cr *CommonResource) error
	ApplyOutputFilter(ctx context.Context, cr *CommonResource) error
}
