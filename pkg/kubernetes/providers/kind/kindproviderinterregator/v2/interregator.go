package kindproviderinterregator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type Interregator struct{}

func (i Interregator) NewInterregator(nodes []v1.Node) interregatortypes.ClusterInterregator {
	interregator := &KindProviderinterregator{
		nodes: nodes,
	}
	if !interregator.IsTypeOf() {
		return nil
	}
	return interregator
}

type KindProviderinterregator struct {
	nodes []v1.Node
}

func (t KindProviderinterregator) IsTypeOf() bool {
	return strings.HasPrefix(t.nodes[0].Spec.ProviderID, "kind")
}
func (t KindProviderinterregator) GetProvider() providermodels.ProviderType {
	if t.IsTypeOf() {
		return providermodels.ProviderTypeKind
	}
	return providermodels.ProviderTypeUnknown
}
func (t KindProviderinterregator) GetClusterId() string {
	return providermodels.UNKNOWN_CLUSTER_ID
}
func (t KindProviderinterregator) GetClusterName() string {
	hostname := t.nodes[0].GetLabels()["kubernetes.io/hostname"]
	return getClusterNameOfArray(hostname)
}
func (t KindProviderinterregator) GetClusterWorkspace() string {
	return fmt.Sprintf("%s-%s", "local", t.nodes[0].GetLabels()["beta.kubernetes.io/instance-type"])
}
func (t KindProviderinterregator) GetDatacenter() string {
	dataCenter := t.GetRegion() + " " + t.GetAz()
	return dataCenter
}
func (t KindProviderinterregator) GetAz() string {
	return "local"
}
func (t KindProviderinterregator) GetRegion() string {
	return "kind"
}
func (t KindProviderinterregator) GetMachineProvider() string {
	return "kind"
}
func (t KindProviderinterregator) GetKubernetesProvider() string {
	return "kind"
}
func getClusterNameOfArray(hostname string) string {
	hostnameArray := strings.Split(hostname, "-")
	lastblock := hostnameArray[len(hostnameArray)-1]
	var length int
	if _, err := strconv.Atoi(lastblock); err == nil {
		length = len(hostnameArray) - 2

	} else {
		length = len(hostnameArray) - 1
	}

	var builder strings.Builder
	for i := 0; i < length; i++ {
		if hostnameArray[i] == "control" || hostnameArray[i] == "plane" {
			break
		}
		if hostnameArray[i] == "controlplane" {
			break
		}
		if builder.Len() > 0 {
			builder.WriteString("-")
		}
		builder.WriteString(hostnameArray[i])
	}
	return builder.String()
}
