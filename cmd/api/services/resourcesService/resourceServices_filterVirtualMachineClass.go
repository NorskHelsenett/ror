package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type VirtualMachineClass.
func filterInVirtualMachineClass(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClass]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClass] {
	return unfiltered
}

// Function to filter outgoing resources of type VirtualMachineClass.
// func filterOutVirtualMachineClass(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClass])apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClass]{
//	return unfiltered
//}
