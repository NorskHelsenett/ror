package rortypes

type CortexId string
type CVEId string

type ResourceVirtualMachineVulnerabilityInfo struct {
	// ID cortex assigns to a host
	CortexiD            CortexId `json:"cortexId"`
	HostName            string   `json:"hostName"`
	HostSeverity        string   `json:"hostSeverity"`
	Severity            string   `json:"severity"`
	SeverityScore       float32  `json:"severityScore"`
	LastCalculationTime int64    `json:"lastCalculationTime"`
	LastReportTime      int64    `json:"lastReportTime"`
	// If beneficcial ResourceCVE can be referred instead
	CVEs []CVEId `json:"cves"`
}
