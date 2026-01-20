package gkeproviderinterregator

import (
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type GkeProviderinterregator struct {
	nodes []v1.Node
}
type Interregator struct{}

func (i Interregator) NewInterregator(nodes []v1.Node) interregatortypes.ClusterInterregator {
	interregator := &GkeProviderinterregator{
		nodes: nodes,
	}
	if !interregator.IsTypeOf() {
		return nil
	}
	return interregator
}

func (t GkeProviderinterregator) IsTypeOf() bool {
	return t.nodes[0].GetLabels()["cloud.google.com/gke-container-runtime"] != ""
}
func (t GkeProviderinterregator) GetProvider() providermodels.ProviderType {
	if t.IsTypeOf() {
		return providermodels.ProviderTypeGke
	}
	return providermodels.ProviderTypeUnknown
}

func (t GkeProviderinterregator) GetClusterId() string {
	return t.nodes[0].GetLabels()["kubernetes.io/cluster-id"]
}

func (t GkeProviderinterregator) GetClusterName() string {
	//gk3-roger-cluster-1-pool-2-22ae7c65-3ohs
	hostname := t.nodes[0].Labels["kubernetes.io/hostname"]
	hostname = strings.Replace(hostname, fmt.Sprintf("%s%s", "-", t.nodes[0].Labels["cloud.google.com/gke-nodepool"]), ":", -1)
	hostnameSplit := strings.Split(hostname, ":")
	hostname = hostnameSplit[0]
	hostname = strings.Replace(hostname, "gk3-", "", 1)
	return hostname

}
func (t GkeProviderinterregator) GetClusterWorkspace() string {
	return "Gke"
}
func (t GkeProviderinterregator) GetDatacenter() string {
	dataCenter := t.GetRegion() + " " + t.GetAz()
	return dataCenter
}

func (t GkeProviderinterregator) GetAz() string {
	return t.nodes[0].GetLabels()["topology.kubernetes.io/zone"]
}

func (t GkeProviderinterregator) GetRegion() string {
	return t.nodes[0].GetLabels()["topology.kubernetes.io/region"]
}

func (t GkeProviderinterregator) GetMachineProvider() string {
	return t.nodes[0].GetLabels()["kubernetes.azure.com/role"]
}

func (t GkeProviderinterregator) GetKubernetesProvider() string {
	return "GKE"
}
