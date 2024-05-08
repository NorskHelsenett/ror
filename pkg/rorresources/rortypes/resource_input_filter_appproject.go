package rortypes

import "github.com/NorskHelsenett/ror/pkg/config/globalconfig"

// (r *ResourceAppProject) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceAppProject) ApplyInputFilter() error {
	if globalconfig.InternalAppProjects[r.Metadata.Name] {
		r.RorMeta.Internal = true
	}
	return nil
}
