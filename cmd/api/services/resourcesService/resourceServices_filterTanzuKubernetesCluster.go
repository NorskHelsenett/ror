package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type TanzuKubernetesCluster.
func filterInTanzuKubernetesCluster(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesCluster]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesCluster] {
	return unfiltered
}

// Function to filter outgoing resources of type TanzuKubernetesCluster.
// func filterOutTanzuKubernetesCluster(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesCluster])apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesCluster]{
//	return unfiltered
//}
