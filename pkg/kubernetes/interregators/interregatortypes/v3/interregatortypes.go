package interregatortypes

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

// ProviderDetectFunc is a function that attempts to detect a provider from
// a Kubernetes cluster. It receives a client, inspects cluster state (e.g. node
// annotations), and returns an initialized provider instance or nil if the
// cluster does not match this provider. The returned value should implement
// some or all of the ClusterInterregator getter methods; the factory wraps it
// and falls back to defaults for any unimplemented methods.
type ProviderDetectFunc func(client *kubernetes.Clientset) any

// ClusterMetadata provides provider-specific cluster metadata.
type ClusterMetadata interface {
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
	GetEnvironment() string
	GetKubernetesApiServer() string
	GetKubernetesCA() string
}

// ClusterNodeLister provides access to cluster node information.
type ClusterNodeLister interface {
	Nodes() ClusterNodeReport
}

// ClusterInterregator combines cluster metadata with node listing.
type ClusterInterregator interface {
	ClusterMetadata
	ClusterNodeLister
}

type ClusterNodeReport interface {
	Get() []v1.Node
	GetByName(name string) *v1.Node
	GetByUid(uid string) *v1.Node
	GetByHostname(hostname string) *v1.Node
	GetByMachineProvider(machineProvider providermodels.ProviderType) []v1.Node
}
