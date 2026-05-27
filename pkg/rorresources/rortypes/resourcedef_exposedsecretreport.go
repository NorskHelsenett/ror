package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// K8s namepace struct
type ResourceExposedSecretReport struct {
	Report ResourceVulnerabilityReportReport `json:"report"`
}
type ResourceExposedSecretReportReport struct {
	Summary AquaReportSummary `json:"summary"`
}

// (r *ResourceExposedSecretReport) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceExposedSecretReport) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Namespace] {
		cr.RorMeta.Internal = true
	}
	return nil
}

// (r ResourceExposedSecretReport) Get returns a pointer to the resource of type ResourceExposedSecretReport
func (r *ResourceExposedSecretReport) Get() *ResourceExposedSecretReport {
	return r
}
