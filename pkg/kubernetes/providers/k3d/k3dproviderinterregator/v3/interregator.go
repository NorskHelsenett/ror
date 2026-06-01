package k3dproviderinterregator

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/k3d/k3dclusternamehelper"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type K3dProviderinterregator struct {
	client       *kubernetes.Clientset
	nodes        []v1.Node
	clusterName  string
	workspace    string
	instanceType string
}

// Detect creates a K3dProviderinterregator from a Kubernetes client.
// Returns the initialized provider if this cluster matches K3D, or nil otherwise.
func Detect(client *kubernetes.Clientset) any {
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), v1meta.ListOptions{})
	if err != nil || len(nodes.Items) == 0 {
		return nil
	}

	if !strings.Contains(nodes.Items[0].Status.NodeInfo.KubeletVersion, "k3s") {
		return nil
	}

	labels := nodes.Items[0].GetLabels()
	hostname := labels["kubernetes.io/hostname"]

	return &K3dProviderinterregator{
		client:       client,
		nodes:        nodes.Items,
		clusterName:  k3dclusternamehelper.GetClusternameFromHostname(hostname),
		instanceType: labels["beta.kubernetes.io/instance-type"],
	}
}

func (t *K3dProviderinterregator) GetProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeK3d
}

func (t *K3dProviderinterregator) GetClusterName() string {
	return t.clusterName
}

func (t *K3dProviderinterregator) GetClusterWorkspace() string {
	return fmt.Sprintf("%s-%s", "local", t.instanceType)
}

func (t *K3dProviderinterregator) GetDatacenter() string {
	return t.GetRegion() + " " + t.GetAz()
}

func (t *K3dProviderinterregator) GetAz() string {
	return "local"
}

func (t *K3dProviderinterregator) GetRegion() string {
	return "k3d"
}

func (t *K3dProviderinterregator) GetMachineProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeK3d
}

func (t *K3dProviderinterregator) GetKubernetesProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeK3d
}

func (v *K3dProviderinterregator) GetKubernetesCA() string {
	return v.fetchCACert()
}

// fetchCACert retrieves the cluster CA certificate from the kube-root-ca.crt ConfigMap.
func (v *K3dProviderinterregator) fetchCACert() string {
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
