package resourcesservice

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

// Function to filter incomming resources of type RbacAssessmentReport.
func filterInRbacAssessmentReport(input apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRbacAssessmentReport]) apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRbacAssessmentReport] {
	if globalconfig.InternalNamespaces[input.Resource.Metadata.Namespace] {
		input.Internal = true
	}
	return input
}

// Function to filter outgoing resources of type RbacAssessmentReport.
// func filterOutRbacAssessmentReport(unfiltered apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRbacAssessmentReport])apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRbacAssessmentReport]{
//	return unfiltered
//}
