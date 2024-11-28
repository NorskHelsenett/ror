package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type Node.
func filterInNode(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNode]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNode] {
	return unfiltered
}

// Function to filter outgoing resources of type Node.
// func filterOutNode(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNode]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNode] {
// 	return unfiltered
// }
