package talosproviderinterregator

import (
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type Interregator struct{}

func (i Interregator) NewInterregator(nodes []v1.Node) interregatortypes.ClusterInterregator {
	interregator := &Talostypes{
		nodes: nodes,
	}
	if !interregator.IsTypeOf() {
		return nil
	}
	return interregator
}

type Talostypes struct {
	nodes []v1.Node
}

func (t Talostypes) IsTypeOf() bool {
	return strings.Contains(strings.ToLower(t.nodes[0].Status.NodeInfo.OSImage), "talos")
}

func (t Talostypes) GetProvider() providermodels.ProviderType {
	if t.IsTypeOf() {
		return providermodels.ProviderTypeTalos
	}
	return providermodels.ProviderTypeUnknown
}

func (t Talostypes) GetClusterId() string {
	clusterId := t.nodes[0].GetAnnotations()["ror.io/cluster-id"]
	return clusterId
}

func (t Talostypes) GetClusterName() string {
	clusterName := t.nodes[0].GetAnnotations()["ror.io/name"]
	return clusterName
}

func (t Talostypes) GetClusterWorkspace() string {
	workspace, ok := t.nodes[0].GetAnnotations()["ror.io/namespace"]
	if !ok {
		workspace = "Talos"
	}

	return workspace
}

func (t Talostypes) GetDatacenter() string {
	dataCenter := t.GetRegion() + " " + t.GetAz()
	return dataCenter
}

func (t Talostypes) GetAz() string {
	return "TalosAZ"
}

func (t Talostypes) GetMachineProvider() string {
	return "TalosVM"
}

func (t Talostypes) GetRegion() string {
	dataCenter, ok := t.nodes[0].GetAnnotations()["ror.io/datacenter"]
	if !ok {
		dataCenter = "TalosDC"
	}

	return dataCenter
}

func (t Talostypes) GetKubernetesProvider() string {
	return "Talos"
}
