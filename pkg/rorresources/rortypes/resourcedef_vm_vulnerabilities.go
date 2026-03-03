package rortypes

type ResourceVirtualMachineVulnerabilityInfo struct {
	Id       string                                    `json:"id"`
	HostName string                                    `json:"hostName"`
	Status   ResourceVirtualMachineVulnerabilityStatus `json:"status"`
	Spec     ResourceVirtualMachineVulnerabilitySpec   `json:"spec"`
}

type ResourceVirtualMachineVulnerabilityStatus struct {
	HostSeverity        string  `json:"hostSeverity"`
	Severity            string  `json:"severity"`
	SeverityScore       float32 `json:"severityScore"`
	CVEs                []CVE   `json:"cves"`
	LastCalculationTime int64   `json:"lastCalculationTime"`
	LastReportTime      int64   `json:"lastReportTime"`
}

type ResourceVirtualMachineVulnerabilitySpec struct {
}

type CVE struct {
	Id                 string              `json:"id"`
	CVSS               string              `json:"csvss"`
	Title              string              `json:"title"`
	Description        string              `json:"description"`
	References         []string            `json:"references"`
	VulnerableVersions []VulnerableVersion `json:"vulnerableVersions"`
}

type VulnerableVersion struct {
	Version     string `json:"version"`
	PackageName string `json:"packageName"`
}
