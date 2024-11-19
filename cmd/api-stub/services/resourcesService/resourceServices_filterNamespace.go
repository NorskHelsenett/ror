package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type Namespace.
func filterInNamespace(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNamespace]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNamespace] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Name] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type Namespace.
// func filterOutNamespace(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNamespace]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNamespace] {
// 	return unfiltered
// }
