package k3dproviderinterregator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type K3dtypes struct {
}

func NewInterregator() *K3dtypes {
	return &K3dtypes{}
}

func (t K3dtypes) IsTypeOf(nodes []v1.Node) bool {
	return strings.Contains(nodes[0].Status.NodeInfo.KubeletVersion, "k3s")

}
func (t K3dtypes) GetProvider(nodes []v1.Node) providermodels.ProviderType {
	if t.IsTypeOf(nodes) {
		return providermodels.ProviderTypeK3d
	}
	return providermodels.ProviderTypeUnknown
}
func (t K3dtypes) GetClusterId(nodes []v1.Node) string {
	return nodes[0].GetLabels()["kubernetes.io/cluster-id"]
}
func (t K3dtypes) GetClusterName(nodes []v1.Node) string {
	hostname := nodes[0].GetLabels()["kubernetes.io/hostname"]
	return getClusterNameOfArray(hostname)
}
func (t K3dtypes) GetClusterWorkspace(nodes []v1.Node) string {
	return fmt.Sprintf("%s-%s", "local", nodes[0].GetLabels()["beta.kubernetes.io/instance-type"])
}
func (t K3dtypes) GetDatacenter(nodes []v1.Node) string {
	dataCenter := t.GetRegion(nodes) + " " + t.GetAz(nodes)
	return dataCenter
}

func (t K3dtypes) GetAz(nodes []v1.Node) string {
	return "local"
}

func (t K3dtypes) GetRegion(nodes []v1.Node) string {
	return "k3s"
}

func (t K3dtypes) GetVMProvider(nodes []v1.Node) string {
	return "k3s"
}

func (t K3dtypes) GetKubernetesProvider(nodes []v1.Node) string {
	return "k3s"
}

func getClusterNameOfArray(hostname string) string {
	hostnameArray := strings.Split(hostname, "-")
	lastblock := hostnameArray[len(hostnameArray)-1]
	var clusterName string
	var length int
	if _, err := strconv.Atoi(lastblock); err == nil {
		length = len(hostnameArray) - 2

	} else {
		length = len(hostnameArray) - 1
	}

	var separator string
	for i := 0; i < length; i++ {
		if hostnameArray[i] == "control" || hostnameArray[i] == "plane" {
			break
		}
		if hostnameArray[i] == "controlplane" {
			break
		}
		if len(clusterName) > 0 {
			separator = "-"
		}
		clusterName = clusterName + separator + hostnameArray[i]
	}
	return clusterName
}
