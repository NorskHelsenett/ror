package interregatortypes

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type ClusterProviderInterregator interface {
	NewInterregator([]v1.Node) ClusterInterregator
}

type ClusterInterregator interface {
	GetProvider() providermodels.ProviderType
	GetClusterId() string
	GetClusterName() string
	GetClusterWorkspace() string
	GetDatacenter() string
	GetAz() string
	GetRegion() string
	GetMachineProvider() providermodels.ProviderType
	GetKubernetesProvider() providermodels.ProviderType
	GetCountry() string
	Nodes() ClusterNodeReport
}

type ClusterNodeReport interface {
	Get() []v1.Node
	GetByName(name string) *v1.Node
	GetByUid(uid string) *v1.Node
	GetByHostname(hostname string) *v1.Node
	GetByMachineProvider(machineProvider providermodels.ProviderType) []v1.Node
}
