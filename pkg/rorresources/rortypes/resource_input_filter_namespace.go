package rortypes

import "github.com/NorskHelsenett/ror/pkg/config/globalconfig"

// (r *ResourceNamespace) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceNamespace) ApplyInputFilter() error {
	if globalconfig.InternalNamespaces[r.Metadata.Name] {
		r.RorMeta.Internal = true
	}
	return nil
}
