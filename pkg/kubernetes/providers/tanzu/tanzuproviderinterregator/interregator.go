package tanzuproviderinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type TanzuProviderinterregator struct {
}

func NewInterregator() *TanzuProviderinterregator {
	return &TanzuProviderinterregator{}
}

func (t TanzuProviderinterregator) IsTypeOf(nodes []v1.Node) bool {
	labels := nodes[0].GetLabels()
	return labels["run.tanzu.vmware.com/kubernetesDistributionVersion"] != ""
}
func (t TanzuProviderinterregator) GetProvider(nodes []v1.Node) providermodels.ProviderType {
	if t.IsTypeOf(nodes) {
		return providermodels.ProviderTypeTanzu
	}
	return providermodels.ProviderTypeUnknown
}
func (t TanzuProviderinterregator) GetClusterId(nodes []v1.Node) string {
	return nodes[0].GetAnnotations()["cluster.x-k8s.io/cluster-id"]
}
func (t TanzuProviderinterregator) GetClusterName(nodes []v1.Node) string {
	return nodes[0].GetAnnotations()["cluster.x-k8s.io/cluster-name"]
}
func (t TanzuProviderinterregator) GetClusterWorkspace(nodes []v1.Node) string {
	return nodes[0].GetAnnotations()["cluster.x-k8s.io/cluster-namespace"]
}
func (t TanzuProviderinterregator) GetDatacenter(nodes []v1.Node) string {
	dataCenter := t.GetRegion(nodes) + " " + t.GetAz(nodes)
	return dataCenter
}

func (t TanzuProviderinterregator) GetAz(nodes []v1.Node) string {
	return nodes[0].GetLabels()["vitistack.io/az"]
}

func (t TanzuProviderinterregator) GetRegion(nodes []v1.Node) string {
	return nodes[0].GetLabels()["vitistack.io/region"]
}

func (t TanzuProviderinterregator) GetVMProvider(nodes []v1.Node) string {
	return "VMwareESXi"
}

func (t TanzuProviderinterregator) GetKubernetesProvider(nodes []v1.Node) string {
	return providermodels.ProviderTypeTanzu.String()
}
