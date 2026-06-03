package tanzuproviderinterregator

import (
	"context"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/goccy/go-yaml"
	v1 "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type TanzuProviderinterregator struct {
	client      *kubernetes.Clientset
	nodes       []v1.Node
	clusterName string
	workspace   string
	datacenter  string
}

type k8sClusterConfiguration struct {
	ControlPlaneEndpoint string `yaml:"controlPlaneEndpoint"`
}

// Detect creates a TanzuProviderinterregator from a Kubernetes client.
// Returns the initialized provider if this cluster matches Tanzu, or nil otherwise.
func Detect(client *kubernetes.Clientset) any {
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), v1meta.ListOptions{})
	if err != nil || len(nodes.Items) == 0 {
		return nil
	}

	labels := nodes.Items[0].GetLabels()
	if labels["run.tanzu.vmware.com/kubernetesDistributionVersion"] == "" {
		return nil
	}

	annotations := nodes.Items[0].GetAnnotations()
	clusterName := annotations["cluster.x-k8s.io/cluster-name"]
	workspace := annotations["cluster.x-k8s.io/cluster-namespace"]

	datacenter := providermodels.UNKNOWN_DATACENTER
	if workspace != "" {
		parts := strings.Split(workspace, "-")
		if len(parts) > 0 {
			datacenter = parts[0]
		}
	}

	return &TanzuProviderinterregator{
		client:      client,
		nodes:       nodes.Items,
		clusterName: clusterName,
		workspace:   workspace,
		datacenter:  datacenter,
	}
}

func (t *TanzuProviderinterregator) GetProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeTanzu
}

func (t *TanzuProviderinterregator) GetClusterName() string {
	return t.clusterName
}

func (t *TanzuProviderinterregator) GetClusterWorkspace() string {
	return t.workspace
}

func (t *TanzuProviderinterregator) GetDatacenter() string {
	return t.datacenter
}

func (t *TanzuProviderinterregator) GetMachineProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeVmware
}

func (t *TanzuProviderinterregator) GetKubernetesProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeTanzu
}

func (t *TanzuProviderinterregator) GetKubernetesApiServer() string {
	kubeadmConfigMap, err := t.client.CoreV1().ConfigMaps("kube-system").Get(context.TODO(), "kubeadm-config", v1meta.GetOptions{})
	if err != nil {
		errMsg := "getControlPlaneEndpoint: Could not get cluster config from kube-system/kubeadm-config, check rbac"
		rlog.Error(errMsg, err)
		return ""
	}

	if kubeadmConfigMap == nil {
		errMsg := "getControlPlaneEndpoint: get value 'ControlPlaneEndpoint' from yaml"
		rlog.Error(errMsg, nil)
		return ""
	}

	kubeadmClusterConfiguration := kubeadmConfigMap.Data["ClusterConfiguration"]

	var clusterConfigurationValues k8sClusterConfiguration

	err = yaml.Unmarshal([]byte(kubeadmClusterConfiguration), &clusterConfigurationValues)
	if err != nil {
		errMsg := "getControlPlaneEndpoint: Could not parse yaml string to stuct"
		rlog.Error(errMsg, errors.New(""))
		return ""
	}
	if clusterConfigurationValues.ControlPlaneEndpoint == "" {
		errMsg := "getControlPlaneEndpoint: ControlPlaneEndpoint is empty in configmap"
		rlog.Error(errMsg, errors.New(""))
		return ""
	}
	return clusterConfigurationValues.ControlPlaneEndpoint
}
func (v *TanzuProviderinterregator) GetKubernetesCA() string {
	return v.fetchCACert()
}

// fetchCACert retrieves the cluster CA certificate from the kube-root-ca.crt ConfigMap.
func (v *TanzuProviderinterregator) fetchCACert() string {
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
