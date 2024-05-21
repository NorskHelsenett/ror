package unknownproviderinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/models/providers"
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
func (t UnknownProviderinterregator) GetProvider(nodes []v1.Node) providers.ProviderType {
	return providers.ProviderTypeUnknown
}
func (t UnknownProviderinterregator) GetClusterName(nodes []v1.Node) string {
	return "unknown"
}
func (t UnknownProviderinterregator) GetWorkspace(nodes []v1.Node) string {
	return "unknown"
}
func (t UnknownProviderinterregator) GetDatacenter(nodes []v1.Node) string {
	return "unknown"
}
