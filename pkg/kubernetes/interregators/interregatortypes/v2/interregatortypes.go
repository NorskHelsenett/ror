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
	GetMachineProvider() string
	GetKubernetesProvider() string
}
