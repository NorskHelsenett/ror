package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type VirtualMachineClassBinding.
func filterInVirtualMachineClassBinding(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClassBinding]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClassBinding] {
	return unfiltered
}

// Function to filter outgoing resources of type VirtualMachineClassBinding.
// func filterOutVirtualMachineClassBinding(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClassBinding])apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClassBinding]{
//	return unfiltered
//}
