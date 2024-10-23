package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

func filterInVm(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVm]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVm] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}
