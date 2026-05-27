package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// K8s namepace struct
type ResourceConfigAuditReport struct {
	Report ResourceVulnerabilityReportReport `json:"report"`
}
type ResourceConfigAuditReportReport struct {
	Summary AquaReportSummary `json:"summary"`
}

// (r *ResourceConfigAuditReport) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceConfigAuditReport) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Namespace] {
		cr.RorMeta.Internal = true
	}
	return nil
}
