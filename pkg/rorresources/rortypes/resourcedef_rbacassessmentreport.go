package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// K8s namepace struct
type ResourceRbacAssessmentReport struct {
	Report ResourceVulnerabilityReportReport `json:"report"`
}
type ResourceRbacAssessmentReportReport struct {
	Summary AquaReportSummary `json:"summary"`
	Scanner AquaReportScanner `json:"scanner"`
}

type AquaReportSummary struct {
	CriticalCount int `json:"criticalCount"`
	HighCount     int `json:"highCount"`
	LowCount      int `json:"lowCount"`
	MediumCount   int `json:"mediumCount"`
	Total         int `json:"total,omitempty"`
}

type AquaReportScanner struct {
	Name    string `json:"name"`
	Vendor  string `json:"vendor"`
	Version string `json:"version"`
}

// (r *ResourceRbacAssessmentReport) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceRbacAssessmentReport) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Namespace] {
		cr.RorMeta.Internal = true
	}
	return nil
}

// (r ResourceRbacAssessmentReport) Get returns a pointer to the resource of type ResourceRbacAssessmentReport
func (r *ResourceRbacAssessmentReport) Get() *ResourceRbacAssessmentReport {
	return r
}
