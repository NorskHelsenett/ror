package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type Certificate.
func filterInCertificate(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceCertificate]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceCertificate] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type Certificate.
// func filterOutCertificate(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceCertificate]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceCertificate] {
// 	return unfiltered
// }
