package vitistackinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

const (
	ClusterNameAnnotation        = "vitistack.io/clustername"        // The name of the cluster
	ClusterWorkspaceAnnotation   = "vitistack.io/clusterworkspace"   // The workspace of the cluster
	RegionAnnotation             = "vitistack.io/region"             // The region of the cluster
	AzAnnotation                 = "vitistack.io/az"                 // The availability zone of the cluster
	VMProviderAnnotation         = "vitistack.io/vmprovider"         // The VM provider of the cluster
	KubernetesProviderAnnotation = "vitistack.io/kubernetesprovider" // The Kubernetes provider of the cluster
	ClusterIdAnnotation          = "vitistack.io/clusterid"          // The ID of the cluster, this is the uuid in ror

)

var (
	MustBeSet = []string{
		ClusterNameAnnotation,
		ClusterWorkspaceAnnotation,
		RegionAnnotation,
		AzAnnotation,
		VMProviderAnnotation,
		KubernetesProviderAnnotation,
		ClusterIdAnnotation,
	}
)

type Vitistacktypes struct {
	initialized        bool
	isOfType           bool
	clustername        string
	clusterworkspace   string
	region             string
	az                 string
	machineprovider    string
	kubernetesprovider string
	clusterId          string
}

func NewInterregator() *Vitistacktypes {
	return &Vitistacktypes{}
}

func (v Vitistacktypes) MustInitialize(nodes []v1.Node) bool {
	if v.isOfType {
		return true
	}

	if v.initialized {
		return false
	}

	for _, node := range nodes {
		if v.checkIfValid(node) {
			v.clustername = getValueByKey(node, ClusterNameAnnotation)
			v.clusterworkspace = getValueByKey(node, ClusterWorkspaceAnnotation)
			v.region = getValueByKey(node, RegionAnnotation)
			v.az = getValueByKey(node, AzAnnotation)
			v.machineprovider = getValueByKey(node, VMProviderAnnotation)
			v.kubernetesprovider = getValueByKey(node, KubernetesProviderAnnotation)
			v.clusterId = getValueByKey(node, ClusterIdAnnotation)
			v.isOfType = true
			v.initialized = true
			return true
		}
	}

	v.initialized = true
	v.isOfType = false
	return false
}

func (v Vitistacktypes) checkIfValid(node v1.Node) bool {

	for _, key := range MustBeSet {
		if !checkIfKeyPresent(node, key) {
			return false
		}
	}
	return true
}

func checkIfKeyPresent(node v1.Node, key string) bool {
	_, ok := node.GetAnnotations()[key]
	if ok {
		return true
	}

	_, ok = node.GetLabels()[key]
	if ok {
		return true
	}
	return false
}
func getValueByKey(node v1.Node, key string) string {
	value, ok := node.GetAnnotations()[key]
	if ok {
		return value
	}

	value, ok = node.GetLabels()[key]
	if ok {
		return value
	}
	return ""
}

// IsTypeOf checks if the nodes are of type Vitistack
// TODO: Improve detection logic
func (v Vitistacktypes) IsTypeOf(nodes []v1.Node) bool {
	return v.MustInitialize(nodes)
}

// GetProvider returns the provider type of the nodes.
func (v Vitistacktypes) GetProvider(nodes []v1.Node) providermodels.ProviderType {
	if !v.MustInitialize(nodes) {
		return providermodels.ProviderTypeUnknown
	}
	return providermodels.ProviderTypeVitistack

}

// GetClusterId returns the cluster ID of the nodes.
func (v Vitistacktypes) GetClusterId(nodes []v1.Node) string {
	v.MustInitialize(nodes)

	clusterId, ok := nodes[0].GetAnnotations()[ClusterIdAnnotation]
	if !ok {
		clusterId = providermodels.UNKNOWN_CLUSTER_ID
	}
	return clusterId
}

// GetClusterName returns the cluster name of the nodes.
func (v Vitistacktypes) GetClusterName(nodes []v1.Node) string {
	clusterName, ok := nodes[0].GetAnnotations()[ClusterNameAnnotation]
	if !ok {
		clusterName = providermodels.UNKNOWN_CLUSTER
	}
	return clusterName
}

// GetClusterWorkspace returns the cluster workspace of the nodes.
func (v Vitistacktypes) GetClusterWorkspace(nodes []v1.Node) string {
	workspace, ok := nodes[0].GetAnnotations()[ClusterWorkspaceAnnotation]
	if !ok {
		workspace = "Vitistack"
	}

	return workspace
}

// GetDatacenter returns the datacenter of the cluster.
func (v Vitistacktypes) GetDatacenter(nodes []v1.Node) string {
	dataCenter := v.GetRegion(nodes) + " " + v.GetAz(nodes)
	return dataCenter
}

// GetRegion returns the region of the cluster.
func (v Vitistacktypes) GetRegion(nodes []v1.Node) string {

	region, ok := nodes[0].GetAnnotations()[RegionAnnotation]
	if !ok {
		region = providermodels.UNKNOWN_REGION
	}

	return region
}

// GetAz returns the availability zone of the cluster.
func (v Vitistacktypes) GetAz(nodes []v1.Node) string {
	az, ok := nodes[0].GetAnnotations()[AzAnnotation]
	if !ok {
		az = providermodels.UNKNOWN_AZ
	}

	return az
}

// GetVMProvider returns the VM provider of the cluster.
func (v Vitistacktypes) GetVMProvider(nodes []v1.Node) string {

	vmProvider, ok := nodes[0].GetAnnotations()[VMProviderAnnotation]
	if !ok {
		vmProvider = providermodels.UNKNOWN_VMPROVIDER
	}

	return vmProvider
}

// GetKubernetesProvider returns the Kubernetes provider of the cluster.
func (v Vitistacktypes) GetKubernetesProvider(nodes []v1.Node) string {
	kubernetesProvider, ok := nodes[0].GetAnnotations()[KubernetesProviderAnnotation]
	if !ok {
		kubernetesProvider = providermodels.UNKNOWN_KUBERNETES_PROVIDER
	}
	return kubernetesProvider
}
