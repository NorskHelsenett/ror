package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incoming resources of type ClusterComplianceReport.
func filterInClusterComplianceReport(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterComplianceReport]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterComplianceReport] {
	return unfiltered
}
