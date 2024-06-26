package providerinterregationreport

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/clusterinterregator"
	"github.com/NorskHelsenett/ror/pkg/models/providers"

	v1 "k8s.io/api/core/v1"
)

type InterregationReport struct {
	Nodes       []v1.Node
	Provider    providers.ProviderType
	ClusterName string
	Workspace   string
	Datacenter  string
}

func (c *InterregationReport) GetProvider() providers.ProviderType {
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

	report.Provider = interregator.GetProvider()
	report.ClusterName = interregator.GetClusterName()
	report.Workspace = interregator.GetWorkspace()
	report.Datacenter = interregator.GetDatacenter()

	return &report, nil
}
