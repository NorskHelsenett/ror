package rortypes

import "github.com/NorskHelsenett/ror/pkg/config/globalconfig"

// (r *ResourcePersistentVolumeClaim) ApplyInputFilter Applies the input filter to the resource
func (r *ResourcePersistentVolumeClaim) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Namespace] {
		cr.RorMeta.Internal = true
	}
	return nil
}
