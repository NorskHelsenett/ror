package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

func filterInVirtualMachine(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachine]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachine] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}
