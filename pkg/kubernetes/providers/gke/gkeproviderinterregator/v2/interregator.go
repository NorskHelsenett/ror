package gkeproviderinterregator

import (
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type Gketypes struct {
	nodes []v1.Node
}
type Interregator struct{}

func (i Interregator) NewInterregator(nodes []v1.Node) interregatortypes.ClusterInterregator {
	interregator := &Gketypes{
		nodes: nodes,
	}
	if !interregator.IsTypeOf() {
		return nil
	}
	return interregator
}

func (t Gketypes) IsTypeOf() bool {
	return t.nodes[0].GetLabels()["cloud.google.com/gke-container-runtime"] != ""
}
func (t Gketypes) GetProvider() providermodels.ProviderType {
	if t.IsTypeOf() {
		return providermodels.ProviderTypeGke
	}
	return providermodels.ProviderTypeUnknown
}

func (t Gketypes) GetClusterId() string {
	return t.nodes[0].GetLabels()["kubernetes.io/cluster-id"]
}

func (t Gketypes) GetClusterName() string {
	//gk3-roger-cluster-1-pool-2-22ae7c65-3ohs
	hostname := t.nodes[0].Labels["kubernetes.io/hostname"]
	hostname = strings.Replace(hostname, fmt.Sprintf("%s%s", "-", t.nodes[0].Labels["cloud.google.com/gke-nodepool"]), ":", -1)
	hostnameSplit := strings.Split(hostname, ":")
	hostname = hostnameSplit[0]
	hostname = strings.Replace(hostname, "gk3-", "", 1)
	return hostname

}
func (t Gketypes) GetClusterWorkspace() string {
	return "Gke"
}
func (t Gketypes) GetDatacenter() string {
	dataCenter := t.GetRegion() + " " + t.GetAz()
	return dataCenter
}

func (t Gketypes) GetAz() string {
	return t.nodes[0].GetLabels()["topology.kubernetes.io/zone"]
}

func (t Gketypes) GetRegion() string {
	return t.nodes[0].GetLabels()["topology.kubernetes.io/region"]
}

func (t Gketypes) GetMachineProvider() string {
	return t.nodes[0].GetLabels()["kubernetes.azure.com/role"]
}

func (t Gketypes) GetKubernetesProvider() string {
	return "GKE"
}
