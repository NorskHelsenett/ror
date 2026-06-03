package azureproviderinterregator

import (
	"context"
	"encoding/base64"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type AzureProviderinterregator struct {
	client      *kubernetes.Clientset
	nodes       []v1.Node
	clusterName string
	region      string
	az          string
}

// Detect creates an AzureProviderinterregator from a Kubernetes client.
// Returns the initialized provider if this cluster matches Azure/AKS, or nil otherwise.
func Detect(client *kubernetes.Clientset) any {
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), v1meta.ListOptions{})
	if err != nil || len(nodes.Items) == 0 {
		return nil
	}

	labels := nodes.Items[0].GetLabels()
	if labels["kubernetes.azure.com/role"] == "" {
		return nil
	}

	return &AzureProviderinterregator{
		client:      client,
		nodes:       nodes.Items,
		clusterName: labels["aks-cluster-name"],
		region:      labels["topology.kubernetes.io/region"],
		az:          labels["topology.kubernetes.io/zone"],
	}
}

func (t *AzureProviderinterregator) GetProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeAks
}

func (t *AzureProviderinterregator) GetClusterName() string {
	return t.clusterName
}

func (t *AzureProviderinterregator) GetClusterWorkspace() string {
	return "Azure"
}

func (t *AzureProviderinterregator) GetDatacenter() string {
	return t.region + " " + t.az
}

func (t *AzureProviderinterregator) GetAz() string {
	return t.az
}

func (t *AzureProviderinterregator) GetRegion() string {
	return t.region
}

func (t *AzureProviderinterregator) GetMachineProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeAks
}

func (t *AzureProviderinterregator) GetKubernetesProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeAks
}

func (v *AzureProviderinterregator) GetKubernetesCA() string {
	return v.fetchCACert()
}

// fetchCACert retrieves the cluster CA certificate from the kube-root-ca.crt ConfigMap.
func (v *AzureProviderinterregator) fetchCACert() string {
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
