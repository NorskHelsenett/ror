package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type ExposedSecretReport.
func filterInExposedSecretReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceExposedSecretReport]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceExposedSecretReport] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type ExposedSecretReport.
// func filterOutExposedSecretReport(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceExposedSecretReport])apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceExposedSecretReport]{
//	return unfiltered
//}
