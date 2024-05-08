package rortypes

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// (r *ResourcePolicyReport) ApplyInputFilter Applies the input filter to the resource
func (r *ResourcePolicyReport) ApplyInputFilter() error {
	if globalconfig.InternalNamespaces[r.Metadata.Namespace] {
		r.RorMeta.Internal = true
	}

	r.RorMeta.LastReported = time.Now().Local().String()

	return nil
}
