package unknownproviderinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type UnknownProviderinterregator struct {
}

func NewInterregator() *UnknownProviderinterregator {
	return &UnknownProviderinterregator{}
}

func (t UnknownProviderinterregator) IsTypeOf(nodes []v1.Node) bool {
	return false
}
func (t UnknownProviderinterregator) GetProvider(nodes []v1.Node) providermodels.ProviderType {
	return providermodels.ProviderTypeUnknown
}
func (t UnknownProviderinterregator) GetClusterId(nodes []v1.Node) string {
	return "unknown"
}
func (t UnknownProviderinterregator) GetClusterName(nodes []v1.Node) string {
	return "unknown"
}
func (t UnknownProviderinterregator) GetClusterWorkspace(nodes []v1.Node) string {
	return "unknown"
}
func (t UnknownProviderinterregator) GetDatacenter(nodes []v1.Node) string {
	dataCenter := t.GetRegion(nodes) + " " + t.GetAz(nodes)
	return dataCenter
}

func (t UnknownProviderinterregator) GetAz(nodes []v1.Node) string {
	return "unknown"
}

func (t UnknownProviderinterregator) GetRegion(nodes []v1.Node) string {
	return "unknown"
}

func (t UnknownProviderinterregator) GetVMProvider(nodes []v1.Node) string {
	return "unknown"
}

func (t UnknownProviderinterregator) GetKubernetesProvider(nodes []v1.Node) string {
	return "unknown"
}
