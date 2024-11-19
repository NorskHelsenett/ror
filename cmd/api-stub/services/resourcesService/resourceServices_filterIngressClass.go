package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type IngressClass.
func filterInIngressClass(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngressClass]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngressClass] {
	return unfiltered
}

// Function to filter outgoing resources of type IngressClass.
// func filterOutIngressClass(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngressClass]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngressClass] {
//   return unfiltered
// }
