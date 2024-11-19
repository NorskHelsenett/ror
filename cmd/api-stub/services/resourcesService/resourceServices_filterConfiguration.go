package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type VulnerabilityReport.
func filterInConfiguration(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfiguration]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfiguration] {
	return input
}

// Function to filter outgoing resources of type VulnerabilityReport.
// func filterOutVulnerabilityReport(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityReport])apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityReport]{
//	return unfiltered
//}
