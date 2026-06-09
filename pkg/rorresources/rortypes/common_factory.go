package rortypes

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mitchellh/hashstructure/v2"
)

type CommonFactory struct {
	resource any
}

type CommonFactoryOutputFilter interface {
	ApplyOutputFilter(ctx context.Context, cr *CommonResource) error
}

type CommonFactoryInputFilter interface {
	ApplyInputFilter(cr *CommonResource) error
}

type CommonFactoryHashGetter interface {
	GetRorHash() string
}

func NewCommonFactory(resource any) CommonResourceInterface {
	return &CommonFactory{
		resource: resource,
	}
}

func (cf *CommonFactory) GetRorHash() string {
	if p, ok := cf.resource.(CommonFactoryHashGetter); ok && !reflect.ValueOf(p).IsNil() {
		return p.GetRorHash()
	}
	hash, err := hashstructure.Hash(cf.resource, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)

}

func (cf *CommonFactory) ApplyInputFilter(cr *CommonResource) error {
	if p, ok := cf.resource.(CommonFactoryInputFilter); ok && !reflect.ValueOf(p).IsNil() {
		return p.ApplyInputFilter(cr)
	}
	return nil
}

func (cf *CommonFactory) ApplyOutputFilter(ctx context.Context, cr *CommonResource) error {
	if p, ok := cf.resource.(CommonFactoryOutputFilter); ok && !reflect.ValueOf(p).IsNil() {
		return p.ApplyOutputFilter(ctx, cr)
	}
	return nil
}
