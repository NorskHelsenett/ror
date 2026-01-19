package k3dproviderinterregator

import (
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/helpers/providerclusternamehelper"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type Interregator struct{}

func (i Interregator) NewInterregator(nodes []v1.Node) interregatortypes.ClusterInterregator {
	interregator := &K3dtypes{
		nodes: nodes,
	}
	if !interregator.IsTypeOf() {
		return nil
	}
	return interregator
}

type K3dtypes struct {
	nodes []v1.Node
}

func (t K3dtypes) IsTypeOf() bool {
	return strings.Contains(t.nodes[0].Status.NodeInfo.KubeletVersion, "k3s")

}
func (t K3dtypes) GetProvider() providermodels.ProviderType {
	if t.IsTypeOf() {
		return providermodels.ProviderTypeK3d
	}
	return providermodels.ProviderTypeUnknown
}
func (t K3dtypes) GetClusterId() string {
	return t.nodes[0].GetLabels()["kubernetes.io/cluster-id"]
}
func (t K3dtypes) GetClusterName() string {
	hostname := t.nodes[0].GetLabels()["kubernetes.io/hostname"]
	return providerclusternamehelper.GetK3dClustername(hostname)
}
func (t K3dtypes) GetClusterWorkspace() string {
	return fmt.Sprintf("%s-%s", "local", t.nodes[0].GetLabels()["beta.kubernetes.io/instance-type"])
}
func (t K3dtypes) GetDatacenter() string {
	dataCenter := t.GetRegion() + " " + t.GetAz()
	return dataCenter
}

func (t K3dtypes) GetAz() string {
	return "local"
}

func (t K3dtypes) GetRegion() string {
	return "k3s"
}

func (t K3dtypes) GetMachineProvider() string {
	return "k3s"
}

func (t K3dtypes) GetKubernetesProvider() string {
	return "k3s"
}
