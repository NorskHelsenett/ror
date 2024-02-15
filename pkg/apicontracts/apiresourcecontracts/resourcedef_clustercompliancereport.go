package apiresourcecontracts

// K8s namepace struct
type ResourceClusterComplianceReport struct {
	ApiVersion string                                `json:"apiVersion"`
	Kind       string                                `json:"kind"`
	Metadata   ResourceMetadata                      `json:"metadata"`
	Status     ResourceClusterComplianceReportStatus `json:"status"`
}

type ResourceClusterComplianceReportStatus struct {
	Summary       ResourceClusterComplianceReportSummary       `json:"summary"`
	SummaryReport ResourceClusterComplianceReportSummaryReport `json:"summaryReport"`
}

type ResourceClusterComplianceReportSummary struct {
	FailCount int `json:"failCount"`
	PassCount int `json:"passCount"`
}

type ResourceClusterComplianceReportSummaryReport struct {
	ControlCheck []any  `json:"controlCheck"`
	Id           string `json:"id"`
	Title        string `json:"title"`
}
