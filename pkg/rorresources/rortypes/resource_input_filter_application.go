package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// (r *ResourceApplication) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceApplication) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalAppProjects[r.Spec.Project] {
		cr.RorMeta.Internal = true
	}
	return nil
}
