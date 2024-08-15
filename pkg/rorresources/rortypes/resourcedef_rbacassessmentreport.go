package rortypes

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
