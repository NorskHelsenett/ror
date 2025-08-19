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

type Vitistacktypes struct {
}

func NewInterregator() *Vitistacktypes {
	return &Vitistacktypes{}
}

// IsTypeOf checks if the nodes are of type Vitistack
func (v Vitistacktypes) IsTypeOf(nodes []v1.Node) bool {
	if _, ok := nodes[0].GetAnnotations()[ClusterNameAnnotation]; ok {
		return true
	}
	return false
}

// GetProvider returns the provider type of the nodes.
func (v Vitistacktypes) GetProvider(nodes []v1.Node) providermodels.ProviderType {
	if v.IsTypeOf(nodes) {
		return providermodels.ProviderTypeVitistack
	}
	return providermodels.ProviderTypeUnknown
}

// GetClusterId returns the cluster ID of the nodes.
func (v Vitistacktypes) GetClusterId(nodes []v1.Node) string {
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
