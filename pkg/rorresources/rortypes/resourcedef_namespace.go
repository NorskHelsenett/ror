package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// K8s namepace struct

type ResourceNamespace struct {
}

// (r *ResourceNamespace) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceNamespace) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Name] {
		cr.RorMeta.Internal = true
	}
	return nil
}

// (r ResourceNamespace) Get returns a pointer to the resource of type ResourceNamespace
func (r *ResourceNamespace) Get() *ResourceNamespace {
	return r
}

// Namespaceinterface represents the interface for resources of the type namespace
type Namespaceinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceNamespace
}
