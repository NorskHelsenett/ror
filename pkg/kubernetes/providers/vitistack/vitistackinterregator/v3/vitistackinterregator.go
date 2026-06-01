package vitistackinterregator

import (
	"context"
	"encoding/base64"

	"github.com/NorskHelsenett/ror/pkg/helpers/kubernetes/metadatahelper"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	vitiv1alpha1 "github.com/vitistack/common/pkg/v1alpha1"
)

const (
	ClusterNameKey        = vitiv1alpha1.ClusterNameAnnotation        // The name of the cluster
	ClusterWorkspaceKey   = vitiv1alpha1.ClusterWorkspaceAnnotation   // The workspace of the cluster
	RegionKey             = vitiv1alpha1.RegionAnnotation             // The region of the cluster
	AzKey                 = vitiv1alpha1.AzAnnotation                 // The availability zone of the cluster
	MachineProviderKey    = vitiv1alpha1.VMProviderAnnotation         // The VM provider of the cluster
	KubernetesProviderKey = vitiv1alpha1.KubernetesProviderAnnotation // The Kubernetes provider of the cluster
	ClusterIdKey          = vitiv1alpha1.ClusterIdAnnotation          // The ID of the cluster, this is the uuid in ror
	CountryKey            = vitiv1alpha1.CountryAnnotation            // The country of the cluster
	EnvironmentKey        = vitiv1alpha1.EnvironmentAnnotation        // The environment of the cluster
	ApiEndpointKey        = vitiv1alpha1.K8sEndpointAnnotation        // The API endpoint of the cluster
)

var (
	mustBeSet = []string{
		ClusterNameKey,
		ClusterWorkspaceKey,
		RegionKey,
		AzKey,
		MachineProviderKey,
		KubernetesProviderKey,
		ClusterIdKey,
	}
)

type VitistackProviderinterregator struct {
	client             *kubernetes.Clientset
	nodes              []v1.Node
	initialized        bool
	isOfType           bool
	clustername        string
	clusterworkspace   string
	region             string
	az                 string
	machineprovider    string
	kubernetesprovider string
	clusterId          string
	country            string
	environment        string
	apiEndpoint        string
	caCert              string
}

// Detect creates a VitistackProviderinterregator from a Kubernetes client.
// Returns the initialized provider if this cluster matches Vitistack, or nil otherwise.
func Detect(client *kubernetes.Clientset) any {
	interregator := &VitistackProviderinterregator{
		client: client,
	}
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), v1meta.ListOptions{})
	if err != nil {
		return nil
	}
	interregator.nodes = nodes.Items

	if !interregator.DetectProvider() {
		return nil
	}

	return interregator
}

// DetectProvider attempts to detect whether this cluster is a Vitistack provider
// by inspecting node annotations/labels. Returns true if detection succeeds.
func (v *VitistackProviderinterregator) DetectProvider() bool {
	if v.isOfType {
		return true
	}

	if v.initialized {
		return false
	}
	for _, node := range v.nodes {
		if v.checkIfValid(&node) {
			v.clustername, _ = getValueByKey(&node, ClusterNameKey)
			v.clusterworkspace, _ = getValueByKey(&node, ClusterWorkspaceKey)
			v.region, _ = getValueByKey(&node, RegionKey)
			v.az, _ = getValueByKey(&node, AzKey)
			v.machineprovider, _ = getValueByKey(&node, MachineProviderKey)
			v.kubernetesprovider, _ = getValueByKey(&node, KubernetesProviderKey)
			v.clusterId, _ = getValueByKey(&node, ClusterIdKey)
			v.country, _ = getValueByKey(&node, CountryKey)
			v.environment = getValueByKeyWithDefault(&node, EnvironmentKey, providermodels.UNKNOWN_ENVIRONMENT)
			v.apiEndpoint, _ = getValueByKey(&node, ApiEndpointKey)
			v.isOfType = true
			v.initialized = true
			return true
		}
	}

	v.initialized = true
	v.isOfType = false
	return false
}

func (v *VitistackProviderinterregator) checkIfValid(node *v1.Node) bool {

	for _, key := range mustBeSet {
		if !checkIfKeyPresent(node, key) {
			return false
		}
	}
	return true
}

func checkIfKeyPresent(node *v1.Node, key string) bool {
	return metadatahelper.CheckAnnotationOrLabel(node.ObjectMeta, key)
}

func getValueByKeyWithDefault(node *v1.Node, key string, def string) string {
	res, ok := metadatahelper.GetAnnotationOrLabel(node.ObjectMeta, key)
	if !ok {
		return def
	}
	return res
}

func getValueByKey(node *v1.Node, key string) (string, bool) {
	return metadatahelper.GetAnnotationOrLabel(node.ObjectMeta, key)
}

// GetProvider returns the provider type.
func (v *VitistackProviderinterregator) GetProvider() providermodels.ProviderType {
	return providermodels.ProviderTypeVitistack
}

// GetClusterId returns the cluster ID.
func (v *VitistackProviderinterregator) GetClusterId() string {
	return v.clusterId
}

// GetClusterName returns the cluster name.
func (v *VitistackProviderinterregator) GetClusterName() string {
	return v.clustername
}

// GetClusterWorkspace returns the cluster workspace.
func (v *VitistackProviderinterregator) GetClusterWorkspace() string {
	return v.clusterworkspace
}

// GetDatacenter returns the datacenter of the cluster.
func (v *VitistackProviderinterregator) GetDatacenter() string {
	return v.az + "." + v.region + "." + v.GetCountry()
}

// GetRegion returns the region of the cluster.
func (v *VitistackProviderinterregator) GetRegion() string {
	return v.region
}

// GetAz returns the availability zone of the cluster.
func (v *VitistackProviderinterregator) GetAz() string {
	return v.az
}

// GetCountry returns the country of the cluster.
func (v *VitistackProviderinterregator) GetCountry() string {
	if v.country == "" {
		return providermodels.DefaultCountry
	}
	return v.country
}

// GetMachineProvider returns the machine provider of the cluster.
func (v *VitistackProviderinterregator) GetMachineProvider() providermodels.ProviderType {
	return providermodels.ProviderType(v.machineprovider)
}

// GetKubernetesProvider returns the Kubernetes provider of the cluster.
func (v *VitistackProviderinterregator) GetKubernetesProvider() providermodels.ProviderType {
	return providermodels.ProviderType(v.kubernetesprovider)
}

// GetEnvironment returns the environment of the cluster.
func (v *VitistackProviderinterregator) GetEnvironment() string {
	return v.environment
}

func (v *VitistackProviderinterregator) GetKubernetesApiServer() string {
	return v.apiEndpoint
}

func (v *VitistackProviderinterregator) GetKubernetesCA() string {
	return v.fetchCACert()
}

// fetchCACert retrieves the cluster CA certificate from the kube-root-ca.crt ConfigMap.
func (v *VitistackProviderinterregator) fetchCACert() string {
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