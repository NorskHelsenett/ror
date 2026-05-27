// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package rortypes

// CommonResourceInterface represents the minimum interface for all resources
type CommonResourceInterface interface {
	GetRorHash() string
	ApplyInputFilter(cr *CommonResource) error
	ApplyOutputFilter(cr *CommonResource) error
}
