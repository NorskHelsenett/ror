package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type PersistentVolumeClaim.
func filterInPersistentVolumeClaim(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePersistentVolumeClaim]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePersistentVolumeClaim] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type PersistentVolumeClaim.
// func filterOutPersistentVolumeClaim(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePersistentVolumeClaim]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePersistentVolumeClaim] {
// 	return unfiltered
// }
