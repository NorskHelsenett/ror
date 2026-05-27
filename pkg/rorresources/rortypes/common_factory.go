package rortypes

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/hashstructure/v2"
)

type CommonFactory struct {
	resource any
}

func NewCommonFactory(resource any) CommonResourceInterface {
	return &CommonFactory{
		resource: resource,
	}
}

func (cf *CommonFactory) GetRorHash() string {
	if cf.resource == nil {
		return ""
	}
	// check if the underlying resource implements GetRorHash specifically,
	// without requiring the full CommonResourceInterface
	v := reflect.ValueOf(cf.resource)
	if m := v.MethodByName("GetRorHash"); m.IsValid() {
		results := m.Call(nil)
		if len(results) == 1 {
			if s, ok := results[0].Interface().(string); ok {
				return s
			}
		}
	}

	hash, err := hashstructure.Hash(cf.resource, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)

}

func (cf *CommonFactory) ApplyInputFilter(cr *CommonResource) error {
	if cf.resource == nil {
		return nil
	}
	// check if the underlying resource implements ApplyInputFilter specifically
	v := reflect.ValueOf(cf.resource)
	if m := v.MethodByName("ApplyInputFilter"); m.IsValid() {
		results := m.Call([]reflect.Value{reflect.ValueOf(cr)})
		if len(results) == 1 && !results[0].IsNil() {
			return results[0].Interface().(error)
		}
	}
	return nil
}

func (cf *CommonFactory) ApplyOutputFilter(cr *CommonResource) error {
	if cf.resource == nil {
		return nil
	}
	// check if the underlying resource implements ApplyOutputFilter specifically
	v := reflect.ValueOf(cf.resource)
	if m := v.MethodByName("ApplyOutputFilter"); m.IsValid() {
		results := m.Call([]reflect.Value{reflect.ValueOf(cr)})
		if len(results) == 1 && !results[0].IsNil() {
			return results[0].Interface().(error)
		}
	}
	return nil
}
