package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type StorageClass.
func filterInStorageClass(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStorageClass]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStorageClass] {
	return unfiltered
}

// Function to filter outgoing resources of type StorageClass.
// func filterOutStorageClass(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStorageClass]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStorageClass] {
// 	return unfiltered
// }
