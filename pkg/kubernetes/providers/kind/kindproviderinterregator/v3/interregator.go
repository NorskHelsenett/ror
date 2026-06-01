package kindproviderinterregator

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/kind/kindclusternamehelper"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type KindProviderinterregator struct {
	client       *kubernetes.Clientset
	nodes        []v1.Node
	clusterName  string
	instanceType string
}

// Detect creates a KindProviderinterregator from a Kubernetes client.
// Returns the initialized provider if this cluster matches Kind, or nil otherwise.
func Detect(client *kubernetes.Clientset) any {
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), v1meta.ListOptions{})
	if err != nil || len(nodes.Items) == 0 {
		return nil
	}

	if !strings.HasPrefix(nodes.Items[0].Spec.ProviderID, "kind") {
		return nil
	}

	labels := nodes.Items[0].GetLabels()
	hostname := labels["kubernetes.io/hostname"]

	return &KindProviderinterregator{
		client:       client,
		nodes:        nodes.Items,
		clusterName:  kindclusternamehelper.GetClusternameFromHostname(hostname),
		instanceType: labels["beta.kubernetes.io/instance-type"],
	}
}

func (t *KindProviderinterregator) GetProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeKind
}

func (t *KindProviderinterregator) GetClusterName() string {
	return t.clusterName
}

func (t *KindProviderinterregator) GetClusterWorkspace() string {
	return fmt.Sprintf("%s-%s", "local", t.instanceType)
}

func (t *KindProviderinterregator) GetDatacenter() string {
	return t.GetRegion() + " " + t.GetAz()
}

func (t *KindProviderinterregator) GetAz() string {
	return "local"
}

func (t *KindProviderinterregator) GetRegion() string {
	return "kind"
}

func (t *KindProviderinterregator) GetMachineProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeKind
}

func (t *KindProviderinterregator) GetKubernetesProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeKind
}

func (v *KindProviderinterregator) GetKubernetesCA() string {
	return v.fetchCACert()
}

// fetchCACert retrieves the cluster CA certificate from the kube-root-ca.crt ConfigMap.
func (v *KindProviderinterregator) fetchCACert() string {
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
