package k3dproviderinterregator

import (
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/k3d/k3dclusternamehelper"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type Interregator struct{}

func (i Interregator) NewInterregator(nodes []v1.Node) interregatortypes.ClusterInterregator {
	interregator := &K3dProviderinterregator{
		nodes: nodes,
	}
	if !interregator.IsTypeOf() {
		return nil
	}
	return interregator
}

type K3dProviderinterregator struct {
	nodes []v1.Node
}

func (t K3dProviderinterregator) IsTypeOf() bool {
	return strings.Contains(t.nodes[0].Status.NodeInfo.KubeletVersion, "k3s")

}
func (t K3dProviderinterregator) GetProvider() providermodels.ProviderType {
	if t.IsTypeOf() {
		return providermodels.ProviderTypeK3d
	}
	return providermodels.ProviderTypeUnknown
}
func (t K3dProviderinterregator) GetClusterId() string {
	return providermodels.UNKNOWN_CLUSTER_ID
}
func (t K3dProviderinterregator) GetClusterName() string {
	hostname := t.nodes[0].GetLabels()["kubernetes.io/hostname"]
	return k3dclusternamehelper.GetClusternameFromHostname(hostname)
}
func (t K3dProviderinterregator) GetClusterWorkspace() string {
	return fmt.Sprintf("%s-%s", "local", t.nodes[0].GetLabels()["beta.kubernetes.io/instance-type"])
}
func (t K3dProviderinterregator) GetDatacenter() string {
	dataCenter := t.GetRegion() + " " + t.GetAz()
	return dataCenter
}

func (t K3dProviderinterregator) GetAz() string {
	return "local"
}

func (t K3dProviderinterregator) GetRegion() string {
	return "k3s"
}

func (t K3dProviderinterregator) GetMachineProvider() string {
	return "k3s"
}

func (t K3dProviderinterregator) GetKubernetesProvider() string {
	return "k3s"
}
