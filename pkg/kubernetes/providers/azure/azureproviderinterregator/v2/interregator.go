package azureproviderinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type Azuretypes struct {
	nodes []v1.Node
}
type Interregator struct{}

func (i Interregator) NewInterregator(nodes []v1.Node) interregatortypes.ClusterInterregator {
	interregator := &Azuretypes{
		nodes: nodes,
	}
	if !interregator.IsTypeOf() {
		return nil
	}
	return interregator
}

func (t Azuretypes) IsTypeOf() bool {
	return t.nodes[0].GetLabels()["kubernetes.azure.com/role"] != ""
}

func (t Azuretypes) GetProvider() providermodels.ProviderType {
	if t.IsTypeOf() {
		return providermodels.ProviderTypeAks
	}
	return providermodels.ProviderTypeUnknown
}

func (t Azuretypes) GetClusterId() string {
	return t.nodes[0].GetLabels()["kubernetes.io/cluster-id"]
}

func (t Azuretypes) GetClusterName() string {
	return t.nodes[0].GetLabels()["aks-cluster-name"]
}

func (t Azuretypes) GetClusterWorkspace() string {
	return "Azure"
}

func (t Azuretypes) GetDatacenter() string {
	dataCenter := t.GetRegion() + " " + t.GetAz()
	return dataCenter
}

func (t Azuretypes) GetAz() string {
	return t.nodes[0].GetLabels()["topology.kubernetes.io/zone"]
}

func (t Azuretypes) GetRegion() string {
	return t.nodes[0].GetLabels()["topology.kubernetes.io/region"]
}

func (t Azuretypes) GetMachineProvider() string {
	return "AzureVM"
}

func (t Azuretypes) GetKubernetesProvider() string {
	return "Azure Kubernetes Service"
}
