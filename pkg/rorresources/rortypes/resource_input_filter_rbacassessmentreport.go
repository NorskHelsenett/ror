package rortypes

import "github.com/NorskHelsenett/ror/pkg/config/globalconfig"

// (r *ResourceRbacAssessmentReport) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceRbacAssessmentReport) ApplyInputFilter() error {
	if globalconfig.InternalNamespaces[r.Metadata.Namespace] {
		r.RorMeta.Internal = true
	}
	return nil
}
