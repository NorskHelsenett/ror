package kindproviderinterregator

import (
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providerinterregator/types"
)

type Kindtypes struct {
}

func NewKindtypes() Kindtypes {
	return Kindtypes{}
}

func (t Kindtypes) IsOfType(cr *types.InterregationReport) bool {
	return strings.HasPrefix(cr.Nodes[0].Spec.ProviderID, "kind")
}
func (t Kindtypes) GetProvider(cr *types.InterregationReport) string {
	return "kind"
}
func (t Kindtypes) GetClusterName(cr *types.InterregationReport) string {
	hostname := cr.Nodes[0].GetLabels()["kubernetes.io/hostname"]
	hostnameArray := strings.Split(hostname, "-")
	return fmt.Sprintf("%s-%s", hostnameArray[0], hostnameArray[1])
}
func (t Kindtypes) GetWorkspace(cr *types.InterregationReport) string {
	return fmt.Sprintf("%s-%s", "local", cr.Nodes[0].GetLabels()["beta.kubernetes.io/instance-type"])
}
func (t Kindtypes) GetDatacenter(cr *types.InterregationReport) string {
	return "local"
}
