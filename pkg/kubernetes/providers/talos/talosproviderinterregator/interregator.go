package talosproviderinterregator

import (
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type Talostypes struct {
}

func NewInterregator() *Talostypes {
	return &Talostypes{}
}

func (t Talostypes) IsTypeOf(nodes []v1.Node) bool {
	return strings.Contains(strings.ToLower(nodes[0].Status.NodeInfo.OSImage), "talos")
}

func (t Talostypes) GetProvider(nodes []v1.Node) providermodels.ProviderType {
	if t.IsTypeOf(nodes) {
		return providermodels.ProviderTypeTalos
	}
	return providermodels.ProviderTypeUnknown
}

func (t Talostypes) GetClusterId(nodes []v1.Node) string {
	clusterId := nodes[0].GetAnnotations()["ror.io/cluster-id"]
	return clusterId
}

func (t Talostypes) GetClusterName(nodes []v1.Node) string {
	clusterName := nodes[0].GetAnnotations()["ror.io/name"]
	return clusterName
}

func (t Talostypes) GetClusterWorkspace(nodes []v1.Node) string {
	workspace, ok := nodes[0].GetAnnotations()["ror.io/namespace"]
	if !ok {
		workspace = "Talos"
	}

	return workspace
}

func (t Talostypes) GetDatacenter(nodes []v1.Node) string {
	dataCenter := t.GetRegion(nodes) + " " + t.GetAz(nodes)
	return dataCenter
}

func (t Talostypes) GetAz(nodes []v1.Node) string {
	return "TalosAZ"
}

func (t Talostypes) GetVMProvider(nodes []v1.Node) string {
	return "TalosVM"
}

func (t Talostypes) GetRegion(nodes []v1.Node) string {
	dataCenter, ok := nodes[0].GetAnnotations()["ror.io/datacenter"]
	if !ok {
		dataCenter = "TalosDC"
	}

	return dataCenter
}

func (t Talostypes) GetKubernetesProvider(nodes []v1.Node) string {
	return string(providermodels.ProviderTypeTalos)
}
