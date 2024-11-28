package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type VulnerabilityReport.
func filterInProject(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceProject]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceProject] {
	return input
}

// Function to filter outgoing resources of type VulnerabilityReport.
// func filterOutVulnerabilityReport(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityReport])apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityReport]{
//	return unfiltered
//}
