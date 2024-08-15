package rortypes

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// (r *ResourcePolicyReport) ApplyInputFilter Applies the input filter to the resource
func (r *ResourcePolicyReport) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Namespace] {
		cr.RorMeta.Internal = true
	}

	cr.RorMeta.LastReported = time.Now().Local().String()

	return nil
}
