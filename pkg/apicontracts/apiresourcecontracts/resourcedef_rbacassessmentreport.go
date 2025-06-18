package apiresourcecontracts

// K8s namepace struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceRbacAssessmentReport struct {
	ApiVersion string                            `json:"apiVersion"`
	Kind       string                            `json:"kind"`
	Metadata   ResourceMetadata                  `json:"metadata"`
	Report     ResourceVulnerabilityReportReport `json:"report"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceRbacAssessmentReportReport struct {
	Summary AquaReportSummary `json:"summary"`
	Scanner AquaReportScanner `json:"scanner"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type AquaReportSummary struct {
	CriticalCount int `json:"criticalCount"`
	HighCount     int `json:"highCount"`
	LowCount      int `json:"lowCount"`
	MediumCount   int `json:"mediumCount"`
	Total         int `json:"total,omitempty"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type AquaReportScanner struct {
	Name    string `json:"name"`
	Vendor  string `json:"vendor"`
	Version string `json:"version"`
}
