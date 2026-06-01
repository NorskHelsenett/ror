package talosproviderinterregator

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type TalosProviderinterregator struct {
	client      *kubernetes.Clientset
	nodes       []v1.Node
	clusterName string
	workspace   string
	datacenter  string
}

// Detect creates a TalosProviderinterregator from a Kubernetes client.
// Returns the initialized provider if this cluster matches Talos, or nil otherwise.
func Detect(client *kubernetes.Clientset) any {
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), v1meta.ListOptions{})
	if err != nil || len(nodes.Items) == 0 {
		return nil
	}

	if !strings.Contains(strings.ToLower(nodes.Items[0].Status.NodeInfo.OSImage), "talos") {
		return nil
	}

	annotations := nodes.Items[0].GetAnnotations()

	workspace := annotations["ror.io/namespace"]
	if workspace == "" {
		workspace = "Talos"
	}

	datacenter := annotations["ror.io/datacenter"]
	if datacenter == "" {
		datacenter = "TalosDC"
	}

	return &TalosProviderinterregator{
		client:      client,
		nodes:       nodes.Items,
		clusterName: annotations["ror.io/name"],
		workspace:   workspace,
		datacenter:  datacenter,
	}
}

func (t *TalosProviderinterregator) GetProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeTalos
}

func (t *TalosProviderinterregator) GetClusterName() string {
	return t.clusterName
}

func (t *TalosProviderinterregator) GetClusterWorkspace() string {
	return t.workspace
}

func (t *TalosProviderinterregator) GetDatacenter() string {
	return t.GetRegion() + " " + t.GetAz()
}

func (t *TalosProviderinterregator) GetAz() string {
	return "TalosAZ"
}

func (t *TalosProviderinterregator) GetRegion() string {
	return t.datacenter
}

func (t *TalosProviderinterregator) GetMachineProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeUnknown
}

func (t *TalosProviderinterregator) GetKubernetesProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeTalos
}

func (v *TalosProviderinterregator) GetKubernetesCA() string {
	return v.fetchCACert()
}

// fetchCACert retrieves the cluster CA certificate from the kube-root-ca.crt ConfigMap.
func (v *TalosProviderinterregator) fetchCACert() string {
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
