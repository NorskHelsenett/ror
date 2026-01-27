package azureproviderinterregator

import (
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
	return interregator
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

func (t AzureProviderinterregator) GetClusterId() string {
	return providermodels.UNKNOWN_CLUSTER_ID
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

func (t AzureProviderinterregator) GetMachineProvider() string {
	return "AzureVM"
}

func (t AzureProviderinterregator) GetKubernetesProvider() string {
	return "Azure Kubernetes Service"
}
