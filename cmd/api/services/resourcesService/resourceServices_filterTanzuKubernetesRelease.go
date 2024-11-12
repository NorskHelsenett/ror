package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type TanzuKubernetesRelease.
func filterInTanzuKubernetesRelease(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesRelease]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesRelease] {
	return unfiltered
}

// Function to filter outgoing resources of type TanzuKubernetesRelease.
// func filterOutTanzuKubernetesRelease(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesRelease])apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesRelease]{
//	return unfiltered
//}
