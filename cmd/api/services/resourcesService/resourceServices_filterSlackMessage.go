package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type SlackMessage.
func filterInSlackMessage(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceSlackMessage]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceSlackMessage] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}
