package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type Ingress.
func filterInIngress(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngress]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngress] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type Ingress.
// func filterOutIngress(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngress]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngress] {
//   return unfiltered
// }
