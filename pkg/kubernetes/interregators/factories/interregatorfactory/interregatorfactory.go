package interregatorfactory

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/factories/nodereportfactory"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/unknown/unknownproviderinterregator/v2"
	v1 "k8s.io/api/core/v1"
)

type ClusterInterregatorFactory struct {
	nodes               []v1.Node
	unknowninterregator interregatortypes.ClusterInterregator
	defaultNodesFunc    interregatortypes.ClusterNodeReport
	config              ClusterInterregatorFactoryConfig
}

type ClusterInterregatorFactoryConfig struct {
	GetProviderFunc           func() providermodels.ProviderType
	GetClusterIdFunc          func() string
	GetClusterNameFunc        func() string
	GetClusterWorkspaceFunc   func() string
	GetDatacenterFunc         func() string
	GetAzFunc                 func() string
	GetRegionFunc             func() string
	GetCountryFunc            func() string
	GetMachineProviderFunc    func() providermodels.ProviderType
	GetKubernetesProviderFunc func() providermodels.ProviderType
	NodesFunc                 func() interregatortypes.ClusterNodeReport
}

func NewClusterInterregatorFactory(nodes []v1.Node, config ClusterInterregatorFactoryConfig) interregatortypes.ClusterInterregator {
	clusterInterregator := ClusterInterregatorFactory{
		nodes:               nodes,
		unknowninterregator: unknownproviderinterregator.NewInterregator(),
		config:              config,
	}
	return &clusterInterregator

}

func (c ClusterInterregatorFactory) GetProvider() providermodels.ProviderType {
	if c.config.GetProviderFunc != nil {
		return c.config.GetProviderFunc()
	}
	return c.unknowninterregator.GetProvider()
}

func (c ClusterInterregatorFactory) GetClusterId() string {
	if c.config.GetClusterIdFunc != nil {
		return c.config.GetClusterIdFunc()
	}
	return c.unknowninterregator.GetClusterId()
}

func (c ClusterInterregatorFactory) GetClusterName() string {
	if c.config.GetClusterNameFunc != nil {
		return c.config.GetClusterNameFunc()
	}
	return c.unknowninterregator.GetClusterName()
}

func (c ClusterInterregatorFactory) GetClusterWorkspace() string {
	if c.config.GetClusterWorkspaceFunc != nil {
		return c.config.GetClusterWorkspaceFunc()
	}
	return c.unknowninterregator.GetClusterWorkspace()
}

func (c ClusterInterregatorFactory) GetDatacenter() string {
	if c.config.GetDatacenterFunc != nil {
		return c.config.GetDatacenterFunc()
	}
	return c.unknowninterregator.GetDatacenter()
}

func (c ClusterInterregatorFactory) GetAz() string {
	if c.config.GetAzFunc != nil {
		return c.config.GetAzFunc()
	}
	return c.unknowninterregator.GetAz()
}

func (c ClusterInterregatorFactory) GetRegion() string {
	if c.config.GetRegionFunc != nil {
		return c.config.GetRegionFunc()
	}
	return c.unknowninterregator.GetRegion()
}

func (c ClusterInterregatorFactory) GetCountry() string {
	if c.config.GetCountryFunc != nil {
		return c.config.GetCountryFunc()
	}
	return c.unknowninterregator.GetCountry()
}

func (c ClusterInterregatorFactory) GetMachineProvider() providermodels.ProviderType {
	if c.config.GetMachineProviderFunc != nil {
		return c.config.GetMachineProviderFunc()
	}
	return c.unknowninterregator.GetMachineProvider()
}

func (c ClusterInterregatorFactory) GetKubernetesProvider() providermodels.ProviderType {
	if c.config.GetKubernetesProviderFunc != nil {
		return c.config.GetKubernetesProviderFunc()
	}
	return c.unknowninterregator.GetKubernetesProvider()
}

func (c *ClusterInterregatorFactory) Nodes() interregatortypes.ClusterNodeReport {
	if c.config.NodesFunc != nil {
		return c.config.NodesFunc()
	}
	if c.defaultNodesFunc == nil {
		c.defaultNodesFunc = nodereportfactory.NewClusterNodeReport(c.nodes)
	}
	return c.defaultNodesFunc
}
