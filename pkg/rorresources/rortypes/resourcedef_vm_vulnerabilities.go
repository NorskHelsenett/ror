package rortypes

type CVEId string

type ResourceVirtualMachineVulnerabilityInfo struct {
	Id       string                                    `json:"id"`
	HostName string                                    `json:"hostName"`
	Status   ResourceVirtualMachineVulnerabilityStatus `json:"status"`
	Spec     ResourceVirtualMachineVulnerabilitySpec   `json:"spec"`
}

type ResourceVirtualMachineVulnerabilityStatus struct {
	Id                  string  `json:"id"`
	HostSeverity        string  `json:"hostSeverity"`
	Severity            string  `json:"severity"`
	SeverityScore       float32 `json:"severityScore"`
	CVEs                []CVEId `json:"cves"`
	LastCalculationTime int64   `json:"lastCalculationTime"`
	LastReportTime      int64   `json:"lastReportTime"`
}

type ResourceVirtualMachineVulnerabilitySpec struct {
}
