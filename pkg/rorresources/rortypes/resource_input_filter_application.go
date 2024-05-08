package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// (r *ResourceApplication) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceApplication) ApplyInputFilter() error {
	if globalconfig.InternalAppProjects[r.Spec.Project] {
		r.RorMeta.Internal = true
	}
	return nil
}
