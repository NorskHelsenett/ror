package k3dproviderinterregator

import (
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/helpers/providerclusternamehelper"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type K3dtypes struct {
}

func NewInterregator() *K3dtypes {
	return &K3dtypes{}
}

func (t K3dtypes) IsTypeOf(nodes []v1.Node) bool {
	return strings.Contains(nodes[0].Status.NodeInfo.KubeletVersion, "k3s")

}
func (t K3dtypes) GetProvider(nodes []v1.Node) providermodels.ProviderType {
	if t.IsTypeOf(nodes) {
		return providermodels.ProviderTypeK3d
	}
	return providermodels.ProviderTypeUnknown
}
func (t K3dtypes) GetClusterId(nodes []v1.Node) string {
	return nodes[0].GetLabels()["kubernetes.io/cluster-id"]
}
func (t K3dtypes) GetClusterName(nodes []v1.Node) string {
	hostname := nodes[0].GetLabels()["kubernetes.io/hostname"]
	return providerclusternamehelper.GetK3dClustername(hostname)
}
func (t K3dtypes) GetClusterWorkspace(nodes []v1.Node) string {
	return fmt.Sprintf("%s-%s", "local", nodes[0].GetLabels()["beta.kubernetes.io/instance-type"])
}
func (t K3dtypes) GetDatacenter(nodes []v1.Node) string {
	dataCenter := t.GetRegion(nodes) + " " + t.GetAz(nodes)
	return dataCenter
}

func (t K3dtypes) GetAz(nodes []v1.Node) string {
	return "local"
}

func (t K3dtypes) GetRegion(nodes []v1.Node) string {
	return "k3s"
}

func (t K3dtypes) GetVMProvider(nodes []v1.Node) string {
	return "k3s"
}

func (t K3dtypes) GetKubernetesProvider(nodes []v1.Node) string {
	return "k3s"
}
