package azureproviderinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providerinterregator/types"
)

type Azuretypes struct {
}

func NewInterregator() Azuretypes {
	return Azuretypes{}
}

func (t Azuretypes) IsOfType(cr *types.InterregationReport) bool {
	return cr.Nodes[0].GetLabels()["kubernetes.azure.com/role"] != ""
}
func (t Azuretypes) GetProvider(cr *types.InterregationReport) string {
	return "tanzu"
}
func (t Azuretypes) GetClusterName(cr *types.InterregationReport) string {
	return cr.Nodes[0].GetLabels()["aks-cluster-name"]
}
func (t Azuretypes) GetWorkspace(cr *types.InterregationReport) string {
	return "Azure"
}
func (t Azuretypes) GetDatacenter(cr *types.InterregationReport) string {
	return cr.Nodes[0].GetLabels()["topology.kubernetes.io/region"]
}
