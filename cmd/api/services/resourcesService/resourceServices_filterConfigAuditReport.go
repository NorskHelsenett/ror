package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type ConfigAuditReport.
func filterInConfigAuditReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfigAuditReport]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfigAuditReport] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type ConfigAuditReport.
// func filterOutConfigAuditReport(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfigAuditReport])apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfigAuditReport]{
//	return unfiltered
//}
