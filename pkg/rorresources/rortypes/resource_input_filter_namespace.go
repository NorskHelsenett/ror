package rortypes

import "github.com/NorskHelsenett/ror/pkg/config/globalconfig"

// (r *ResourceNamespace) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceNamespace) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Name] {
		cr.RorMeta.Internal = true
	}
	return nil
}
