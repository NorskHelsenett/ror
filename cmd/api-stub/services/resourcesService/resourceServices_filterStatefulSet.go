package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type StatefulSet.
func filterInStatefulSet(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStatefulSet]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStatefulSet] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type StatefulSet.
// func filterOutStatefulSet(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStatefulSet]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStatefulSet] {
//	return unfiltered
//}
