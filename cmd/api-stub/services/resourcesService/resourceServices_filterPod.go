package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type Pod.
func filterInPod(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePod]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePod] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type Pod.
// func filterOutPod(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePod]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePod] {
// 	return unfiltered
// }
