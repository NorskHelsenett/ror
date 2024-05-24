package kindproviderinterregator

import (
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/helpers/providerclusternamehelper"
	"github.com/NorskHelsenett/ror/pkg/models/providers"
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
func (t Kindtypes) GetProvider(nodes []v1.Node) providers.ProviderType {
	if t.IsTypeOf(nodes) {
		return providers.ProviderTypeKind
	}
	return providers.ProviderTypeUnknown
}
func (t Kindtypes) GetClusterName(nodes []v1.Node) string {
	hostname := nodes[0].GetLabels()["kubernetes.io/hostname"]
	return providerclusternamehelper.GetKindClustername(hostname)
}
func (t Kindtypes) GetWorkspace(nodes []v1.Node) string {
	return fmt.Sprintf("%s-%s", "local", nodes[0].GetLabels()["beta.kubernetes.io/instance-type"])
}
func (t Kindtypes) GetDatacenter(nodes []v1.Node) string {
	return "local"
}
