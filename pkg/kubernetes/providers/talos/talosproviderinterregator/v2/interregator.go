package talosproviderinterregator

import (
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/factories/interregatorfactory"
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
	return interregatorfactory.NewClusterInterregatorFactory(nodes, interregatorfactory.ClusterInterregatorFactoryConfig{
		GetProviderFunc: func() providermodels.ProviderType {
			return interregator.GetProvider()
		},
		GetClusterNameFunc: func() string {
			return interregator.GetClusterName()
		},
		GetClusterWorkspaceFunc: func() string {
			return interregator.GetClusterWorkspace()
		},
		GetDatacenterFunc: func() string {
			return interregator.GetDatacenter()
		},
		GetAzFunc: func() string {
			return interregator.GetAz()
		},
		GetRegionFunc: func() string {
			return interregator.GetRegion()
		},
		GetMachineProviderFunc: func() providermodels.ProviderType {
			return interregator.GetMachineProvider()
		},
		GetKubernetesProviderFunc: func() providermodels.ProviderType {
			return interregator.GetKubernetesProvider()
		},
	})
}

type TalosProviderinterregator struct {
	nodes []v1.Node
}

func (t TalosProviderinterregator) IsTypeOf() bool {
	if len(t.nodes) == 0 {
		return false
	}
	return strings.Contains(strings.ToLower(t.nodes[0].Status.NodeInfo.OSImage), "talos")
}

func (t TalosProviderinterregator) GetProvider() providermodels.ProviderType {
	if t.IsTypeOf() {
		return providermodels.ProviderTypeTalos
	}
	return providermodels.ProviderTypeUnknown
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

func (t TalosProviderinterregator) GetMachineProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeUnknown
}

func (t TalosProviderinterregator) GetRegion() string {
	dataCenter, ok := t.nodes[0].GetAnnotations()["ror.io/datacenter"]
	if !ok {
		dataCenter = "TalosDC"
	}

	return dataCenter
}

func (t TalosProviderinterregator) GetKubernetesProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeTalos
}
