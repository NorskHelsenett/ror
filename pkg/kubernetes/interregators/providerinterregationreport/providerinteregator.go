package providerinterregationreport

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/clusterinterregator/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"

	v1 "k8s.io/api/core/v1"
)

type InterregationReport struct {
	Nodes              []v1.Node
	Interregator       providermodels.ProviderType
	Provider           providermodels.ProviderType
	KubernetesProvider providermodels.ProviderType
	MachineProvider    providermodels.ProviderType
	ClusterName        string
	Workspace          string
	Datacenter         string
	AvailabilityZone   string
	Region             string
	Country            string
}

func (c *InterregationReport) GetProvider() providermodels.ProviderType {
	return c.Provider
}
func (c *InterregationReport) GetMachineProvider() providermodels.ProviderType {
	return c.MachineProvider

}
func (c *InterregationReport) GetKubernetesProvider() providermodels.ProviderType {
	return c.KubernetesProvider
}

// InterregateCluster interregates the cluster and returns a report
// containing the provider, cluster name, workspace and datacenter
// of the cluster
//
// Parameters:
// - nodes []v1.Node: the nodes in the cluster v1.Node from k8s.io/api/core/v1
func NewInterregationReport(nodes []v1.Node) (*InterregationReport, error) {
	report := &InterregationReport{}

	if len(nodes) == 0 {
		return nil, fmt.Errorf("no nodes found")
	}

	report.Nodes = nodes
	interregator := clusterinterregator.NewClusterInterregator(nodes)

	getInterregationReport(report, interregator)

	return report, nil
}

func GetInterregationReport(interregator interregatortypes.ClusterInterregator) (InterregationReport, error) {

	if interregator == nil {
		return InterregationReport{}, fmt.Errorf("interregator is nil")
	}

	report := InterregationReport{}
	getInterregationReport(&report, interregator)
	return report, nil
}

func getInterregationReport(report *InterregationReport, interregator interregatortypes.ClusterInterregator) {
	report.Interregator = interregator.GetProvider()
	report.Provider = interregator.GetProvider()
	report.MachineProvider = interregator.GetMachineProvider()
	report.KubernetesProvider = interregator.GetKubernetesProvider()
	report.ClusterName = interregator.GetClusterName()
	report.Workspace = interregator.GetClusterWorkspace()
	report.Datacenter = interregator.GetDatacenter()
	report.AvailabilityZone = interregator.GetAz()
	report.Region = interregator.GetRegion()
	report.Country = interregator.GetCountry()
}
