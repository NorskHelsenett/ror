package providerinterregator

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/azure/azureproviderinterregator"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providerinterregator/types"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/tanzu/tanzuproviderinterregator"

	v1 "k8s.io/api/core/v1"
)

// InterregateCluster interregates the cluster and returns a report
// containing the provider, cluster name, workspace and datacenter
// of the cluster
//
// Parameters:
// - nodes []v1.Node: the nodes in the cluster v1.Node from k8s.io/api/core/v1
func NewInterregationReport(nodes []v1.Node) (*types.InterregationReport, error) {
	report := types.InterregationReport{}

	report.Nodes = nodes
	interregator := report.GetInterregator(GetProviderInterregators())

	if interregator == nil {
		return nil, fmt.Errorf("no interregator found")
	}
	report.Provider = interregator.GetProvider(&report)
	report.ClusterName = interregator.GetClusterName(&report)
	report.Workspace = interregator.GetWorkspace(&report)
	report.Datacenter = interregator.GetDatacenter(&report)

	return &report, nil
}

// GetProviderInterregators returns a list of all the provider interregators implemented
// TODO: move to a config provider?
func GetProviderInterregators() []types.ClusterProviderinterregator {
	return []types.ClusterProviderinterregator{
		tanzuproviderinterregator.NewInterregator(),
		azureproviderinterregator.NewInterregator(),
	}
}
