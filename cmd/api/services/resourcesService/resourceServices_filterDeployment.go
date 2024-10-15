package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type Deployment.
func filterInDeployment(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDeployment]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDeployment] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}

// // Function to filter outgoing resources of type Deployment.
// func filterOutDeployment(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDeployment]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDeployment] {
// 	return unfiltered
// }
