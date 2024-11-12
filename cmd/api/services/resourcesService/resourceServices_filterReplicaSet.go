package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type ReplicaSet.
func filterInReplicaSet(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceReplicaSet]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceReplicaSet] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type ReplicaSet.
// func filterOutReplicaSet(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceReplicaSet]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceReplicaSet] {
//	return unfiltered
// }
