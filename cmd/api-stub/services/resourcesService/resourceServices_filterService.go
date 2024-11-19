package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type Service.
func filterInService(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceService]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceService] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type Service.
// func filterOutService(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceService]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceService] {
// 	return unfiltered
// }
