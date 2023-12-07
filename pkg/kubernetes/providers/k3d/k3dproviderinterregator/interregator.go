package k3dproviderinterregator

import (
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providerinterregator/types"
)

type K3dtypes struct {
}

func NewK3dtypes() K3dtypes {
	return K3dtypes{}
}

func (t K3dtypes) IsOfType(cr *types.InterregationReport) bool {
	return strings.Contains(cr.Nodes[0].Status.NodeInfo.KubeletVersion, "k3s")

}
func (t K3dtypes) GetProvider(cr *types.InterregationReport) string {
	return "k3d"
}
func (t K3dtypes) GetClusterName(cr *types.InterregationReport) string {
	hostname := cr.Nodes[0].GetLabels()["kubernetes.io/hostname"]
	hostnameArray := strings.Split(hostname, "-")
	return fmt.Sprintf("%s-%s", hostnameArray[0], hostnameArray[1])
}
func (t K3dtypes) GetWorkspace(cr *types.InterregationReport) string {
	return fmt.Sprintf("%s-%s", "local", cr.Nodes[0].GetLabels()["beta.kubernetes.io/instance-type"])
}
func (t K3dtypes) GetDatacenter(cr *types.InterregationReport) string {
	return "local"
}
