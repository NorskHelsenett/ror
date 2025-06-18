package apiresourcecontracts

// K8s namepace struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceClusterComplianceReport struct {
	ApiVersion string                                `json:"apiVersion"`
	Kind       string                                `json:"kind"`
	Metadata   ResourceMetadata                      `json:"metadata"`
	Status     ResourceClusterComplianceReportStatus `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceClusterComplianceReportStatus struct {
	Summary       ResourceClusterComplianceReportSummary       `json:"summary"`
	SummaryReport ResourceClusterComplianceReportSummaryReport `json:"summaryReport"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceClusterComplianceReportSummary struct {
	FailCount int `json:"failCount"`
	PassCount int `json:"passCount"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceClusterComplianceReportSummaryReport struct {
	ControlCheck []any  `json:"controlCheck"`
	Id           string `json:"id"`
	Title        string `json:"title"`
}
