package apiresourcecontracts

// K8s namepace struct
type ResourceExposedSecretReport struct {
	ApiVersion string                            `json:"apiVersion"`
	Kind       string                            `json:"kind"`
	Metadata   ResourceMetadata                  `json:"metadata"`
	Report     ResourceVulnerabilityReportReport `json:"report"`
}
type ResourceExposedSecretReportReport struct {
	Summary AquaReportSummary `json:"summary"`
}
