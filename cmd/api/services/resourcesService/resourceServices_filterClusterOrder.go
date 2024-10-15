package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type ResourceClusterOrder.
func filterInClusterOrder(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterOrder]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterOrder] {
	//input.Internal = true
	return input
}

// Function to filter outgoing resources of type ResourceClusterOrder.
// func filterOutClusterOrder(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterOrder])apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityReport]{
//	return unfiltered
//}
