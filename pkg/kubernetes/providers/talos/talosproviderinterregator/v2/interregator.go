package talosproviderinterregator

import (
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type Interregator struct{}

func (i Interregator) NewInterregator(nodes []v1.Node) interregatortypes.ClusterInterregator {
	interregator := &TalosProviderinterregator{
		nodes: nodes,
	}
	if !interregator.IsTypeOf() {
		return nil
	}
	return interregator
}

type TalosProviderinterregator struct {
	nodes []v1.Node
}

func (t TalosProviderinterregator) IsTypeOf() bool {
	return strings.Contains(strings.ToLower(t.nodes[0].Status.NodeInfo.OSImage), "talos")
}

func (t TalosProviderinterregator) GetProvider() providermodels.ProviderType {
	if t.IsTypeOf() {
		return providermodels.ProviderTypeTalos
	}
	return providermodels.ProviderTypeUnknown
}

func (t TalosProviderinterregator) GetClusterId() string {
	clusterId := t.nodes[0].GetAnnotations()["ror.io/cluster-id"]
	return clusterId
}

func (t TalosProviderinterregator) GetClusterName() string {
	clusterName := t.nodes[0].GetAnnotations()["ror.io/name"]
	return clusterName
}

func (t TalosProviderinterregator) GetClusterWorkspace() string {
	workspace, ok := t.nodes[0].GetAnnotations()["ror.io/namespace"]
	if !ok {
		workspace = "Talos"
	}

	return workspace
}

func (t TalosProviderinterregator) GetDatacenter() string {
	dataCenter := t.GetRegion() + " " + t.GetAz()
	return dataCenter
}

func (t TalosProviderinterregator) GetAz() string {
	return "TalosAZ"
}

func (t TalosProviderinterregator) GetMachineProvider() string {
	return "TalosVM"
}

func (t TalosProviderinterregator) GetRegion() string {
	dataCenter, ok := t.nodes[0].GetAnnotations()["ror.io/datacenter"]
	if !ok {
		dataCenter = "TalosDC"
	}

	return dataCenter
}

func (t TalosProviderinterregator) GetKubernetesProvider() string {
	return "Talos"
}
