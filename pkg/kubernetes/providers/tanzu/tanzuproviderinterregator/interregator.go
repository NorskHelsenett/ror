package tanzuproviderinterregator

import "github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providerinterregator/types"

type TanzuProviderinterregator struct {
}

func NewInterregator() TanzuProviderinterregator {
	return TanzuProviderinterregator{}
}

func (t TanzuProviderinterregator) IsOfType(cr *types.InterregationReport) bool {
	labels := cr.Nodes[0].GetLabels()
	return labels["run.tanzu.vmware.com/kubernetesDistributionVersion"] != ""
}
func (t TanzuProviderinterregator) GetProvider(cr *types.InterregationReport) string {
	return "tanzu"
}
func (t TanzuProviderinterregator) GetClusterName(cr *types.InterregationReport) string {
	return cr.Nodes[0].GetAnnotations()["cluster.x-k8s.io/cluster-name"]
}
func (t TanzuProviderinterregator) GetWorkspace(cr *types.InterregationReport) string {
	return cr.Nodes[0].GetAnnotations()["cluster.x-k8s.io/cluster-namespace"]
}
func (t TanzuProviderinterregator) GetDatacenter(cr *types.InterregationReport) string {
	return ""
}
