package gkeproviderinterregator

import (
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type Gketypes struct {
}

func NewInterregator() *Gketypes {
	return &Gketypes{}
}

func (t Gketypes) IsTypeOf(nodes []v1.Node) bool {
	return nodes[0].GetLabels()["cloud.google.com/gke-container-runtime"] != ""
}
func (t Gketypes) GetProvider(nodes []v1.Node) providermodels.ProviderType {
	if t.IsTypeOf(nodes) {
		return providermodels.ProviderTypeGke
	}
	return providermodels.ProviderTypeUnknown
}

func (t Gketypes) GetClusterId(nodes []v1.Node) string {
	return nodes[0].GetLabels()["kubernetes.io/cluster-id"]
}

func (t Gketypes) GetClusterName(nodes []v1.Node) string {
	//gk3-roger-cluster-1-pool-2-22ae7c65-3ohs
	hostname := nodes[0].Labels["kubernetes.io/hostname"]
	hostname = strings.Replace(hostname, fmt.Sprintf("%s%s", "-", nodes[0].Labels["cloud.google.com/gke-nodepool"]), ":", -1)
	hostnameSplit := strings.Split(hostname, ":")
	hostname = hostnameSplit[0]
	hostname = strings.Replace(hostname, "gk3-", "", 1)
	return hostname

}
func (t Gketypes) GetClusterWorkspace(nodes []v1.Node) string {
	return "Gke"
}
func (t Gketypes) GetDatacenter(nodes []v1.Node) string {
	dataCenter := t.GetRegion(nodes) + " " + t.GetAz(nodes)
	return dataCenter
}

func (t Gketypes) GetAz(nodes []v1.Node) string {
	return nodes[0].GetLabels()["topology.kubernetes.io/zone"]
}

func (t Gketypes) GetRegion(nodes []v1.Node) string {
	return nodes[0].GetLabels()["topology.kubernetes.io/region"]
}

func (t Gketypes) GetVMProvider(nodes []v1.Node) string {
	return nodes[0].GetLabels()["kubernetes.azure.com/role"]
}

func (t Gketypes) GetKubernetesProvider(nodes []v1.Node) string {
	return "GKE"
}
