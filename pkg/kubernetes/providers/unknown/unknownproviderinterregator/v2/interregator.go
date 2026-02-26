package unknownproviderinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/factories/nodereportfactory"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
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
	return providermodels.UNKNOWN_UNDEFINED
}
func (t UnknownProviderinterregator) GetClusterName() string {
	return providermodels.UNKNOWN_UNDEFINED
}
func (t UnknownProviderinterregator) GetClusterWorkspace() string {
	return providermodels.UNKNOWN_UNDEFINED
}
func (t UnknownProviderinterregator) GetDatacenter() string {
	return providermodels.UNKNOWN_UNDEFINED
}

func (t UnknownProviderinterregator) GetAz() string {
	return providermodels.UNKNOWN_UNDEFINED
}

func (t UnknownProviderinterregator) GetRegion() string {
	return providermodels.UNKNOWN_UNDEFINED
}
func (t UnknownProviderinterregator) GetCountry() string {
	return providermodels.DefaultCountry // Default to Norway, as ROR is developed and used primarily in Norway. This can be overridden by specific providers if they have better information.
}

func (t UnknownProviderinterregator) GetMachineProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeUnknown
}

func (t UnknownProviderinterregator) GetKubernetesProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeUnknown
}

func (t UnknownProviderinterregator) Nodes() interregatortypes.ClusterNodeReport {
	return nodereportfactory.NodeReportNotImplemented{}
}
