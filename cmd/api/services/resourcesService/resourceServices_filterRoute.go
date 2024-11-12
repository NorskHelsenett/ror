package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type VulnerabilityEvent.
func filterInRoute(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRoute]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRoute] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type VulnerabilityEvent.
// func filterOutVulnerabilityEvent(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityEvent])apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityEvent]{
//	return unfiltered
//}
