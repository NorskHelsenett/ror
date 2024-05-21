package azureproviderinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/models/providers"
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
func (t Azuretypes) GetProvider(nodes []v1.Node) providers.ProviderType {
	if t.IsTypeOf(nodes) {
		return providers.ProviderTypeAks
	}
	return providers.ProviderTypeUnknown
}
func (t Azuretypes) GetClusterName(nodes []v1.Node) string {
	return nodes[0].GetLabels()["aks-cluster-name"]
}
func (t Azuretypes) GetWorkspace(nodes []v1.Node) string {
	return "Azure"
}
func (t Azuretypes) GetDatacenter(nodes []v1.Node) string {
	return nodes[0].GetLabels()["topology.kubernetes.io/region"]
}
