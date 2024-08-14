package rortypes

import "github.com/NorskHelsenett/ror/pkg/config/globalconfig"

// (r *ResourceAppProject) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceAppProject) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalAppProjects[cr.Metadata.Name] {
		cr.RorMeta.Internal = true
	}
	return nil
}
