package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type DaemonSet.
func filterInDaemonSet(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDaemonSet]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDaemonSet] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type DaemonSet.
// func filterOutDaemonSet(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDaemonSet]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDaemonSet] {
//	return unfiltered
// }
