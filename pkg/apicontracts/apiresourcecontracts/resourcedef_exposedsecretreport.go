package apiresourcecontracts

// K8s namepace struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceExposedSecretReport struct {
	ApiVersion string                            `json:"apiVersion"`
	Kind       string                            `json:"kind"`
	Metadata   ResourceMetadata                  `json:"metadata"`
	Report     ResourceVulnerabilityReportReport `json:"report"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceExposedSecretReportReport struct {
	Summary AquaReportSummary `json:"summary"`
}
