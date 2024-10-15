package viewsmodels

type ComplianceReport struct {
	Clusterid string                   `json:"clusterid"`
	Metadata  ComplianceReportMetadata `json:"metadata"`
	Summary   ComplianceReportSummary  `json:"summary"`
	Reports   []ComplianceReportReport `json:"reports"`
}

type ComplianceReportMetadata struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type ComplianceReportSummary struct {
	Failcount int `json:"failcount"`
	Passcount int `json:"passcount"`
}

type ComplianceReportReport struct {
	Name      string                   `json:"name"`
	Severity  ComplianceReportSeverity `json:"severity"`
	Totalfail int                      `json:"totalfail"`
	Id        string                   `json:"id"`
}

type ComplianceReportSeverity string

const (
	CRITICAL ComplianceReportSeverity = "CRITICAL"
	HIGH     ComplianceReportSeverity = "HIGH"
	MEDIUM   ComplianceReportSeverity = "MEDIUM"
	LOW      ComplianceReportSeverity = "LOW"
	UNKNOWN  ComplianceReportSeverity = "UNKOWN"
)
