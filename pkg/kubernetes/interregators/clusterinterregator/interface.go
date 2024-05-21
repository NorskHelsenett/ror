package clusterinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/unknown/unknownproviderinterregator"
	"github.com/NorskHelsenett/ror/pkg/models/providers"
	v1 "k8s.io/api/core/v1"
)

type clusterinteregator struct {
	nodes                []v1.Node
	providerinterregator ClusterProviderInterregator
}

type ClusterProviderInterregator interface {
	IsTypeOf([]v1.Node) bool
	GetProvider([]v1.Node) providers.ProviderType
	GetClusterName([]v1.Node) string
	GetWorkspace([]v1.Node) string
	GetDatacenter([]v1.Node) string
}

type ClusterInterregator interface {
	GetProvider() providers.ProviderType
	GetClusterName() string
	GetWorkspace() string
	GetDatacenter() string
}

func NewClusterInterregator(nodes []v1.Node) ClusterInterregator {
	clusterinterregator := clusterinteregator{
		nodes: nodes,
	}

	for _, inter := range interregators {
		if inter.IsTypeOf(nodes) {
			clusterinterregator.providerinterregator = inter
		}
	}
	if clusterinterregator.providerinterregator == nil {
		clusterinterregator.providerinterregator = unknownproviderinterregator.NewInterregator()
	}
	return clusterinterregator

}

func (c clusterinteregator) GetProvider() providers.ProviderType {
	return c.providerinterregator.GetProvider(c.nodes)
}
func (c clusterinteregator) GetClusterName() string {
	return c.providerinterregator.GetClusterName(c.nodes)
}
func (c clusterinteregator) GetWorkspace() string {
	return c.providerinterregator.GetWorkspace(c.nodes)
}
func (c clusterinteregator) GetDatacenter() string {
	return c.providerinterregator.GetDatacenter(c.nodes)
}
