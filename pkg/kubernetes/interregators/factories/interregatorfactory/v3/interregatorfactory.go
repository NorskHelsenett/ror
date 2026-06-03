package interregatorfactory

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/factories/nodereportfactory"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v3"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/unknown/unknownproviderinterregator/v3"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Optional interfaces that providers can implement.
// The factory checks for each via type assertion and falls back to defaults.

type ProviderGetter interface {
	GetProvider() providermodels.ProviderType
}
type ClusterIdGetter interface{ GetClusterId() string }
type ClusterNameGetter interface{ GetClusterName() string }
type ClusterWorkspaceGetter interface{ GetClusterWorkspace() string }
type DatacenterGetter interface{ GetDatacenter() string }
type AzGetter interface{ GetAz() string }
type RegionGetter interface{ GetRegion() string }
type CountryGetter interface{ GetCountry() string }
type EnvironmentGetter interface{ GetEnvironment() string }
type MachineProviderGetter interface {
	GetMachineProvider() providermodels.ProviderType
}
type KubernetesProviderGetter interface {
	GetKubernetesProvider() providermodels.ProviderType
}
type NodesGetter interface {
	Nodes() interregatortypes.ClusterNodeReport
}
type KubernetesApiServerGetter interface{ GetKubernetesApiServer() string }
type KubernetesCAGetter interface{ GetKubernetesCA() string }

type ClusterInterregatorFactory struct {
	client              *kubernetes.Clientset
	unknowninterregator interregatortypes.ClusterInterregator
	defaultNodesFunc    interregatortypes.ClusterNodeReport
	provider            any
}

// NewClusterInterregatorFactory creates a ClusterInterregator by wrapping the given provider.
// It uses type assertions to discover which methods the provider implements.
// Any method from the ClusterInterregator interface that the provider does not implement
// will fall back to the unknown provider defaults.
func NewClusterInterregatorFactory(client *kubernetes.Clientset, provider any) interregatortypes.ClusterInterregator {
	return &ClusterInterregatorFactory{
		client:              client,
		unknowninterregator: unknownproviderinterregator.NewInterregator(),
		provider:            provider,
	}
}

func (c *ClusterInterregatorFactory) GetProvider() providermodels.ProviderType {
	if p, ok := c.provider.(ProviderGetter); ok {
		return p.GetProvider()
	}
	return c.unknowninterregator.GetProvider()
}

func (c *ClusterInterregatorFactory) GetClusterId() string {
	if p, ok := c.provider.(ClusterIdGetter); ok {
		return p.GetClusterId()
	}
	return c.unknowninterregator.GetClusterId()
}

func (c *ClusterInterregatorFactory) GetClusterName() string {
	if p, ok := c.provider.(ClusterNameGetter); ok {
		return p.GetClusterName()
	}
	return c.unknowninterregator.GetClusterName()
}

func (c *ClusterInterregatorFactory) GetClusterWorkspace() string {
	if p, ok := c.provider.(ClusterWorkspaceGetter); ok {
		return p.GetClusterWorkspace()
	}
	return c.unknowninterregator.GetClusterWorkspace()
}

func (c *ClusterInterregatorFactory) GetDatacenter() string {
	if p, ok := c.provider.(DatacenterGetter); ok {
		return p.GetDatacenter()
	}
	return c.unknowninterregator.GetDatacenter()
}

func (c *ClusterInterregatorFactory) GetAz() string {
	if p, ok := c.provider.(AzGetter); ok {
		return p.GetAz()
	}
	return c.unknowninterregator.GetAz()
}

func (c *ClusterInterregatorFactory) GetRegion() string {
	if p, ok := c.provider.(RegionGetter); ok {
		return p.GetRegion()
	}
	return c.unknowninterregator.GetRegion()
}

func (c *ClusterInterregatorFactory) GetCountry() string {
	if p, ok := c.provider.(CountryGetter); ok {
		return p.GetCountry()
	}
	return c.unknowninterregator.GetCountry()
}

func (c *ClusterInterregatorFactory) GetMachineProvider() providermodels.ProviderType {
	if p, ok := c.provider.(MachineProviderGetter); ok {
		return p.GetMachineProvider()
	}
	return c.unknowninterregator.GetMachineProvider()
}

func (c *ClusterInterregatorFactory) GetKubernetesProvider() providermodels.ProviderType {
	if p, ok := c.provider.(KubernetesProviderGetter); ok {
		return p.GetKubernetesProvider()
	}
	return c.unknowninterregator.GetKubernetesProvider()
}

func (c *ClusterInterregatorFactory) Nodes() interregatortypes.ClusterNodeReport {
	if p, ok := c.provider.(NodesGetter); ok {
		return p.Nodes()
	}
	if c.defaultNodesFunc == nil {
		nodes, _ := c.client.CoreV1().Nodes().List(context.TODO(), v1meta.ListOptions{})
		c.defaultNodesFunc = nodereportfactory.NewClusterNodeReport(nodes.Items)
	}
	return c.defaultNodesFunc
}

func (c *ClusterInterregatorFactory) GetEnvironment() string {
	if p, ok := c.provider.(EnvironmentGetter); ok {
		return p.GetEnvironment()
	}
	return c.unknowninterregator.GetEnvironment()
}

func (c *ClusterInterregatorFactory) GetKubernetesApiServer() string {
	if p, ok := c.provider.(KubernetesApiServerGetter); ok {
		return p.GetKubernetesApiServer()
	}
	return c.unknowninterregator.GetKubernetesApiServer()
}

func (c *ClusterInterregatorFactory) GetKubernetesCA() string {
	if p, ok := c.provider.(KubernetesCAGetter); ok {
		return p.GetKubernetesCA()
	}
	return c.unknowninterregator.GetKubernetesCA()
}
