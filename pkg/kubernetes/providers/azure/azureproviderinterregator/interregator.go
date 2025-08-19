package azureproviderinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type Azuretypes struct {
}

func NewInterregator() *Azuretypes {
	return &Azuretypes{}
}

func (t Azuretypes) IsTypeOf(nodes []v1.Node) bool {
	return nodes[0].GetLabels()["kubernetes.azure.com/role"] != ""
}

func (t Azuretypes) GetProvider(nodes []v1.Node) providermodels.ProviderType {
	if t.IsTypeOf(nodes) {
		return providermodels.ProviderTypeAks
	}
	return providermodels.ProviderTypeUnknown
}

func (t Azuretypes) GetClusterId(nodes []v1.Node) string {
	return nodes[0].GetLabels()["kubernetes.io/cluster-id"]
}

func (t Azuretypes) GetClusterName(nodes []v1.Node) string {
	return nodes[0].GetLabels()["aks-cluster-name"]
}

func (t Azuretypes) GetClusterWorkspace(nodes []v1.Node) string {
	return "Azure"
}

func (t Azuretypes) GetDatacenter(nodes []v1.Node) string {
	dataCenter := t.GetRegion(nodes) + " " + t.GetAz(nodes)
	return dataCenter
}

func (t Azuretypes) GetAz(nodes []v1.Node) string {
	return nodes[0].GetLabels()["topology.kubernetes.io/zone"]
}

func (t Azuretypes) GetRegion(nodes []v1.Node) string {
	return nodes[0].GetLabels()["topology.kubernetes.io/region"]
}

func (t Azuretypes) GetVMProvider(nodes []v1.Node) string {
	return "AzureVM"
}

func (t Azuretypes) GetKubernetesProvider(nodes []v1.Node) string {
	return "Azure Kubernetes Service"
}
