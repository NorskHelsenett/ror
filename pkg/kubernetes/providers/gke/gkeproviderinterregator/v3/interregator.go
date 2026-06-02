package gkeproviderinterregator

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type GkeProviderinterregator struct {
	client      *kubernetes.Clientset
	nodes       []v1.Node
	clusterName string
	region      string
	az          string
}

// Detect creates a GkeProviderinterregator from a Kubernetes client.
// Returns the initialized provider if this cluster matches GKE, or nil otherwise.
func Detect(client *kubernetes.Clientset) any {
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), v1meta.ListOptions{})
	if err != nil || len(nodes.Items) == 0 {
		return nil
	}

	labels := nodes.Items[0].GetLabels()
	if labels["cloud.google.com/gke-container-runtime"] == "" {
		return nil
	}

	// Parse cluster name from hostname: gk3-roger-cluster-1-pool-2-22ae7c65-3ohs
	hostname := labels["kubernetes.io/hostname"]
	hostname = strings.Replace(hostname, fmt.Sprintf("%s%s", "-", labels["cloud.google.com/gke-nodepool"]), ":", -1)
	hostnameSplit := strings.Split(hostname, ":")
	hostname = hostnameSplit[0]
	hostname = strings.Replace(hostname, "gk3-", "", 1)

	return &GkeProviderinterregator{
		client:      client,
		nodes:       nodes.Items,
		clusterName: hostname,
		region:      labels["topology.kubernetes.io/region"],
		az:          labels["topology.kubernetes.io/zone"],
	}
}

func (t *GkeProviderinterregator) GetProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeGke
}

func (t *GkeProviderinterregator) GetClusterId() string {
	return providermodels.UNKNOWN_CLUSTER_ID
}

func (t *GkeProviderinterregator) GetClusterName() string {
	return t.clusterName
}

func (t *GkeProviderinterregator) GetClusterWorkspace() string {
	return "Gke"
}

func (t *GkeProviderinterregator) GetDatacenter() string {
	return t.region + " " + t.az
}

func (t *GkeProviderinterregator) GetAz() string {
	return t.az
}

func (t *GkeProviderinterregator) GetRegion() string {
	return t.region
}

func (t *GkeProviderinterregator) GetMachineProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeGke
}

func (t *GkeProviderinterregator) GetKubernetesProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeGke
}

func (v *GkeProviderinterregator) GetKubernetesCA() string {
	return v.fetchCACert()
}

// fetchCACert retrieves the cluster CA certificate from the kube-root-ca.crt ConfigMap.
func (v *GkeProviderinterregator) fetchCACert() string {
	if v.client == nil {
		return ""
	}
	cm, err := v.client.CoreV1().ConfigMaps("default").Get(context.TODO(), "kube-root-ca.crt", v1meta.GetOptions{})
	if err != nil {
		return ""
	}
	caCert, ok := cm.Data["ca.crt"]
	if !ok {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(caCert))
}
