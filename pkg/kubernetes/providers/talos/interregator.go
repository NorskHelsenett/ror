package talosproviderinterregator

import (
	"strings"

	"github.com/NorskHelsenett/ror/pkg/helpers/providerclusternamehelper"
	"github.com/NorskHelsenett/ror/pkg/models/providers"
	v1 "k8s.io/api/core/v1"
)

type Talostypes struct {
}

func NewInterregator() *Talostypes {
	return &Talostypes{}
}

func (t Talostypes) IsTypeOf(nodes []v1.Node) bool {
	return strings.Contains(strings.ToLower(nodes[0].Status.NodeInfo.OSImage), "talos")
}

func (t Talostypes) GetProvider(nodes []v1.Node) providers.ProviderType {
	if t.IsTypeOf(nodes) {
		return providers.ProviderTypeTalos
	}
	return providers.ProviderTypeUnknown
}

func (t Talostypes) GetClusterName(nodes []v1.Node) string {
	hostname := nodes[0].GetLabels()["kubernetes.io/hostname"]
	return providerclusternamehelper.GetKindClustername(hostname)
}

func (t Talostypes) GetWorkspace(nodes []v1.Node) string {
	return "Talos"
}

func (t Talostypes) GetDatacenter(nodes []v1.Node) string {
	return "tempDatacenter"
}
