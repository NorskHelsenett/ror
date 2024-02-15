package apiresourcecontracts

// K8s namepace struct
type ResourceConfigAuditReport struct {
	ApiVersion string                            `json:"apiVersion"`
	Kind       string                            `json:"kind"`
	Metadata   ResourceMetadata                  `json:"metadata"`
	Report     ResourceVulnerabilityReportReport `json:"report"`
}
type ResourceConfigAuditReportReport struct {
	Summary AquaReportSummary `json:"summary"`
}
