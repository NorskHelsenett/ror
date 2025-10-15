package providerinterregationreport

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/clusterinterregator"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"

	v1 "k8s.io/api/core/v1"
)

type InterregationReport struct {
	Nodes              []v1.Node
	Interregator       providermodels.ProviderType
	Provider           providermodels.ProviderType
	KubernetesProvider providermodels.ProviderType
	VMProvider         providermodels.ProviderType
	ClusterName        string
	Workspace          string
	Datacenter         string
}

func (c *InterregationReport) GetProvider() providermodels.ProviderType {
	return c.Provider
}
func (c *InterregationReport) GetVMProvider() providermodels.ProviderType {
	return c.Provider

}
func (c *InterregationReport) GetKubernetesProvider() providermodels.ProviderType {
	return c.Provider
}

// InterregateCluster interregates the cluster and returns a report
// containing the provider, cluster name, workspace and datacenter
// of the cluster
//
// Parameters:
// - nodes []v1.Node: the nodes in the cluster v1.Node from k8s.io/api/core/v1
func NewInterregationReport(nodes []v1.Node) (*InterregationReport, error) {
	report := InterregationReport{}

	if len(nodes) == 0 {
		return nil, fmt.Errorf("no nodes found")
	}

	report.Nodes = nodes
	interregator := clusterinterregator.NewClusterInterregator(nodes)

	report.Interregator = interregator.GetProvider()
	report.Provider = providermodels.ProviderType(interregator.GetKubernetesProvider())
	report.VMProvider = providermodels.ProviderType(interregator.GetVMProvider())
	report.KubernetesProvider = providermodels.ProviderType(interregator.GetKubernetesProvider())
	report.ClusterName = interregator.GetClusterName()
	report.Workspace = interregator.GetClusterWorkspace()
	report.Datacenter = interregator.GetDatacenter()

	return &report, nil
}
