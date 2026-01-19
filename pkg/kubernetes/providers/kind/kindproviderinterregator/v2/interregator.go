package kindproviderinterregator

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
	interregator := &Kindtypes{
		nodes: nodes,
	}
	if !interregator.IsTypeOf() {
		return nil
	}
	return interregator
}

type Kindtypes struct {
	nodes []v1.Node
}

func (t Kindtypes) IsTypeOf() bool {
	return strings.HasPrefix(t.nodes[0].Spec.ProviderID, "kind")
}
func (t Kindtypes) GetProvider() providermodels.ProviderType {
	if t.IsTypeOf() {
		return providermodels.ProviderTypeKind
	}
	return providermodels.ProviderTypeUnknown
}
func (t Kindtypes) GetClusterId() string {
	return t.nodes[0].GetLabels()["kubernetes.io/cluster-id"]
}
func (t Kindtypes) GetClusterName() string {
	hostname := t.nodes[0].GetLabels()["kubernetes.io/hostname"]
	return providerclusternamehelper.GetKindClustername(hostname)
}
func (t Kindtypes) GetClusterWorkspace() string {
	return fmt.Sprintf("%s-%s", "local", t.nodes[0].GetLabels()["beta.kubernetes.io/instance-type"])
}
func (t Kindtypes) GetDatacenter() string {
	dataCenter := t.GetRegion() + " " + t.GetAz()
	return dataCenter
}
func (t Kindtypes) GetAz() string {
	return "local"
}
func (t Kindtypes) GetRegion() string {
	return "kind"
}
func (t Kindtypes) GetMachineProvider() string {
	return "kind"
}
func (t Kindtypes) GetKubernetesProvider() string {
	return "kind"
}
