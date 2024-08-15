package rortypes

// K8s namepace struct
type ResourceExposedSecretReport struct {
	Report ResourceVulnerabilityReportReport `json:"report"`
}
type ResourceExposedSecretReportReport struct {
	Summary AquaReportSummary `json:"summary"`
}
