package tanzuproviderinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type Interregator struct{}

type TanzuProviderinterregator struct {
	nodes []v1.Node
}

func (i Interregator) NewInterregator(nodes []v1.Node) interregatortypes.ClusterInterregator {
	return &TanzuProviderinterregator{
		nodes: nodes,
	}
}

func (t TanzuProviderinterregator) IsTypeOf() bool {
	labels := t.nodes[0].GetLabels()
	return labels["run.tanzu.vmware.com/kubernetesDistributionVersion"] != ""
}
func (t TanzuProviderinterregator) GetProvider() providermodels.ProviderType {
	if t.IsTypeOf() {
		return providermodels.ProviderTypeTanzu
	}
	return providermodels.ProviderTypeUnknown
}
func (t TanzuProviderinterregator) GetClusterId() string {
	return t.nodes[0].GetAnnotations()["cluster.x-k8s.io/cluster-id"]
}
func (t TanzuProviderinterregator) GetClusterName() string {
	return t.nodes[0].GetAnnotations()["cluster.x-k8s.io/cluster-name"]
}
func (t TanzuProviderinterregator) GetClusterWorkspace() string {
	return t.nodes[0].GetAnnotations()["cluster.x-k8s.io/cluster-namespace"]
}
func (t TanzuProviderinterregator) GetDatacenter() string {
	dataCenter := t.GetRegion() + " " + t.GetAz()
	return dataCenter
}

func (t TanzuProviderinterregator) GetAz() string {
	return t.nodes[0].GetLabels()["vitistack.io/az"]
}

func (t TanzuProviderinterregator) GetRegion() string {
	return t.nodes[0].GetLabels()["vitistack.io/region"]
}

func (t TanzuProviderinterregator) GetMachineProvider() string {
	return "VMwareESXi"
}

func (t TanzuProviderinterregator) GetKubernetesProvider() string {
	return providermodels.ProviderTypeTanzu.String()
}
