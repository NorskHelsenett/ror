package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type VulnerabilityReport.
func filterInKubernetesCluster(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceKubernetesCluster]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceKubernetesCluster] {
	input.Internal = true
	return input
}

// Function to filter outgoing resources of type VulnerabilityReport.
// func filterOutVulnerabilityReport(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityReport])apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityReport]{
//	return unfiltered
//}
