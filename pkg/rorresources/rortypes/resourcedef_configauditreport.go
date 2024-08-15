package rortypes

// K8s namepace struct
type ResourceConfigAuditReport struct {
	Report ResourceVulnerabilityReportReport `json:"report"`
}
type ResourceConfigAuditReportReport struct {
	Summary AquaReportSummary `json:"summary"`
}
