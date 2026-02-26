package azureproviderinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/factories/interregatorfactory"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type AzureProviderinterregator struct {
	nodes []v1.Node
}
type Interregator struct{}

func (i Interregator) NewInterregator(nodes []v1.Node) interregatortypes.ClusterInterregator {
	interregator := &AzureProviderinterregator{
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

func (t AzureProviderinterregator) IsTypeOf() bool {
	return t.nodes[0].GetLabels()["kubernetes.azure.com/role"] != ""
}

func (t AzureProviderinterregator) GetProvider() providermodels.ProviderType {
	if t.IsTypeOf() {
		return providermodels.ProviderTypeAks
	}
	return providermodels.ProviderTypeUnknown
}

func (t AzureProviderinterregator) GetClusterName() string {
	return t.nodes[0].GetLabels()["aks-cluster-name"]
}

func (t AzureProviderinterregator) GetClusterWorkspace() string {
	return "Azure"
}

func (t AzureProviderinterregator) GetDatacenter() string {
	dataCenter := t.GetRegion() + " " + t.GetAz()
	return dataCenter
}

func (t AzureProviderinterregator) GetAz() string {
	return t.nodes[0].GetLabels()["topology.kubernetes.io/zone"]
}

func (t AzureProviderinterregator) GetRegion() string {
	return t.nodes[0].GetLabels()["topology.kubernetes.io/region"]
}

func (t AzureProviderinterregator) GetMachineProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeAks
}

func (t AzureProviderinterregator) GetKubernetesProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeAks
}
