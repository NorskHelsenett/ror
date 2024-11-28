package resourcesservice

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type PolicyReport.
func filterInPolicyReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePolicyReport]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePolicyReport] {

	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	input.Resource.LastReported = time.Now().Local().String()
	return input
}

// Function to filter outgoing resources of type PolicyReport.
// func filterOutPolicyReport(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePolicyReport]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePolicyReport] {
// 	return unfiltered
// }
