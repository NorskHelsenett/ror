package tanzuproviderinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/models/providers"
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
func (t TanzuProviderinterregator) GetProvider(nodes []v1.Node) providers.ProviderType {
	if t.IsTypeOf(nodes) {
		return providers.ProviderTypeTanzu
	}
	return providers.ProviderTypeUnknown
}
func (t TanzuProviderinterregator) GetClusterName(nodes []v1.Node) string {
	return nodes[0].GetAnnotations()["cluster.x-k8s.io/cluster-name"]
}
func (t TanzuProviderinterregator) GetWorkspace(nodes []v1.Node) string {
	return nodes[0].GetAnnotations()["cluster.x-k8s.io/cluster-namespace"]
}
func (t TanzuProviderinterregator) GetDatacenter([]v1.Node) string {
	return ""
}
