package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type Application.
func filterInApplication(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceApplication]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceApplication] {
	if globalconfig.InternalAppProjects[input.Resource.Spec.Project] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type Application.
// func filterOutApplication(ctx context.Context, unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceApplication]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceApplication] {
//  	return unfiltered
//  }
