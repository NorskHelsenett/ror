package unknownproviderinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
)

type UnknownProviderinterregator struct {
}

func NewInterregator() *UnknownProviderinterregator {
	return &UnknownProviderinterregator{}
}

func (t UnknownProviderinterregator) IsTypeOf() bool {
	return false
}
func (t UnknownProviderinterregator) GetProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeUnknown
}
func (t UnknownProviderinterregator) GetClusterId() string {
	return "unknown"
}
func (t UnknownProviderinterregator) GetClusterName() string {
	return "unknown"
}
func (t UnknownProviderinterregator) GetClusterWorkspace() string {
	return "unknown"
}
func (t UnknownProviderinterregator) GetDatacenter() string {
	dataCenter := t.GetRegion() + " " + t.GetAz()
	return dataCenter
}

func (t UnknownProviderinterregator) GetAz() string {
	return "unknown"
}

func (t UnknownProviderinterregator) GetRegion() string {
	return "unknown"
}

func (t UnknownProviderinterregator) GetMachineProvider() string {
	return "unknown"
}

func (t UnknownProviderinterregator) GetKubernetesProvider() string {
	return "unknown"
}
