package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type AppProject.
func filterInAppProject(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceAppProject]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceAppProject] {
	if globalconfig.InternalAppProjects[input.Resource.Metadata.Name] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type AppProject.
// func filterOutAppProject(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceAppProject]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceAppProject] {
// 	return unfiltered
// }
