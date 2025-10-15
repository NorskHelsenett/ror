package clusterinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/unknown/unknownproviderinterregator"
	v1 "k8s.io/api/core/v1"
)

type clusterinteregator struct {
	nodes                []v1.Node
	providerinterregator ClusterProviderInterregator
}

// IsTypeOf(nodes []v1.Node) bool

type ClusterProviderInterregator interface {
	IsTypeOf([]v1.Node) bool
	GetProvider([]v1.Node) providermodels.ProviderType
	GetClusterId(nodes []v1.Node) string
	GetClusterName([]v1.Node) string
	GetClusterWorkspace([]v1.Node) string
	GetDatacenter([]v1.Node) string
	GetAz(nodes []v1.Node) string
	GetRegion(nodes []v1.Node) string
	GetVMProvider(nodes []v1.Node) string
	GetKubernetesProvider(nodes []v1.Node) string
}

type ClusterInterregator interface {
	GetProvider() providermodels.ProviderType
	GetClusterId() string
	GetClusterName() string
	GetClusterWorkspace() string
	GetDatacenter() string
	GetAz() string
	GetRegion() string
	GetVMProvider() string
	GetKubernetesProvider() string
}

func NewClusterInterregator(nodes []v1.Node) ClusterInterregator {
	clusterinterregator := clusterinteregator{
		nodes: nodes,
	}

	for _, inter := range interregators {
		if inter.IsTypeOf(nodes) {
			clusterinterregator.providerinterregator = inter
			break
		}
	}
	if clusterinterregator.providerinterregator == nil {
		clusterinterregator.providerinterregator = unknownproviderinterregator.NewInterregator()
	}
	return clusterinterregator

}

func (c clusterinteregator) GetProvider() providermodels.ProviderType {
	return c.providerinterregator.GetProvider(c.nodes)
}
func (c clusterinteregator) GetClusterId() string {
	return c.providerinterregator.GetClusterId(c.nodes)
}
func (c clusterinteregator) GetClusterName() string {
	return c.providerinterregator.GetClusterName(c.nodes)
}
func (c clusterinteregator) GetClusterWorkspace() string {
	return c.providerinterregator.GetClusterWorkspace(c.nodes)
}
func (c clusterinteregator) GetDatacenter() string {
	return c.providerinterregator.GetDatacenter(c.nodes)
}
func (c clusterinteregator) GetAz() string {
	return c.providerinterregator.GetAz(c.nodes)
}
func (c clusterinteregator) GetRegion() string {
	return c.providerinterregator.GetRegion(c.nodes)
}
func (c clusterinteregator) GetVMProvider() string {
	return c.providerinterregator.GetVMProvider(c.nodes)
}
func (c clusterinteregator) GetKubernetesProvider() string {
	return c.providerinterregator.GetKubernetesProvider(c.nodes)
}
