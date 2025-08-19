package kindproviderinterregator

import (
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/helpers/providerclusternamehelper"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type Kindtypes struct {
}

func NewInterregator() *Kindtypes {
	return &Kindtypes{}
}

func (t Kindtypes) IsTypeOf(nodes []v1.Node) bool {
	return strings.HasPrefix(nodes[0].Spec.ProviderID, "kind")
}
func (t Kindtypes) GetProvider(nodes []v1.Node) providermodels.ProviderType {
	if t.IsTypeOf(nodes) {
		return providermodels.ProviderTypeKind
	}
	return providermodels.ProviderTypeUnknown
}
func (t Kindtypes) GetClusterId(nodes []v1.Node) string {
	return nodes[0].GetLabels()["kubernetes.io/cluster-id"]
}
func (t Kindtypes) GetClusterName(nodes []v1.Node) string {
	hostname := nodes[0].GetLabels()["kubernetes.io/hostname"]
	return providerclusternamehelper.GetKindClustername(hostname)
}
func (t Kindtypes) GetClusterWorkspace(nodes []v1.Node) string {
	return fmt.Sprintf("%s-%s", "local", nodes[0].GetLabels()["beta.kubernetes.io/instance-type"])
}
func (t Kindtypes) GetDatacenter(nodes []v1.Node) string {
	dataCenter := t.GetRegion(nodes) + " " + t.GetAz(nodes)
	return dataCenter
}
func (t Kindtypes) GetAz(nodes []v1.Node) string {
	return "local"
}
func (t Kindtypes) GetRegion(nodes []v1.Node) string {
	return "kind"
}
func (t Kindtypes) GetVMProvider(nodes []v1.Node) string {
	return "kind"
}
func (t Kindtypes) GetKubernetesProvider(nodes []v1.Node) string {
	return "kind"
}
