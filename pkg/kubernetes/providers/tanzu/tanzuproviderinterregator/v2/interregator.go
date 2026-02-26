package tanzuproviderinterregator

import (
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/factories/interregatorfactory"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/factories/nodereportfactory"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type Interregator struct{}

type TanzuProviderinterregator struct {
	nodes []v1.Node
}

func (i Interregator) NewInterregator(nodes []v1.Node) interregatortypes.ClusterInterregator {
	interregator := &TanzuProviderinterregator{
		nodes: nodes,
	}
	if !interregator.IsTypeOf() {
		return nil
	}
	return interregatorfactory.NewClusterInterregatorFactory(nodes, interregatorfactory.ClusterInterregatorFactoryConfig{
		GetProviderFunc: func() providermodels.ProviderType {
			return interregator.GetProvider()
		},
		GetClusterNameFunc: func() string {
			return interregator.GetClusterName()
		},
		GetClusterWorkspaceFunc: func() string {
			return interregator.GetClusterWorkspace()
		},
		GetDatacenterFunc: func() string {
			return interregator.GetDatacenter()
		},
		GetMachineProviderFunc: func() providermodels.ProviderType {
			return interregator.GetMachineProvider()
		},
		GetKubernetesProviderFunc: func() providermodels.ProviderType {
			return interregator.GetKubernetesProvider()
		},
		NodesFunc: func() interregatortypes.ClusterNodeReport {
			return nodereportfactory.NewNodeReportFactory(nodes)
		},
	})
}

func (t TanzuProviderinterregator) IsTypeOf() bool {
	labels := t.nodes[0].GetLabels()
	return labels["run.tanzu.vmware.com/kubernetesDistributionVersion"] != ""
}
func (t TanzuProviderinterregator) GetProvider() providermodels.ProviderType {
	if t.IsTypeOf() {
		return providermodels.ProviderTypeTanzu
	}
	return providermodels.ProviderTypeUnknown
}
func (t TanzuProviderinterregator) GetClusterName() string {
	return t.nodes[0].GetAnnotations()["cluster.x-k8s.io/cluster-name"]
}
func (t TanzuProviderinterregator) GetClusterWorkspace() string {
	return t.nodes[0].GetAnnotations()["cluster.x-k8s.io/cluster-namespace"]
}
func (t TanzuProviderinterregator) GetDatacenter() string {
	ws := t.nodes[0].GetAnnotations()["cluster.x-k8s.io/cluster-namespace"]
	workspaceArray := strings.Split(ws, "-")
	if len(workspaceArray) > 0 {
		return workspaceArray[0]
	}
	return providermodels.UNKNOWN_DATACENTER
}

func (t TanzuProviderinterregator) GetMachineProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeVmware
}

func (t TanzuProviderinterregator) GetKubernetesProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeTanzu
}
