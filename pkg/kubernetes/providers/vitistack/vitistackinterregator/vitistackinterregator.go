// Package vitistackinterregator provides utilities to identify Vitistack Kubernetes clusters
// and extract cluster metadata from node annotations. It enables detection of Vitistack clusters
// and retrieval of information such as cluster ID, name, workspace, region, availability zone,
// VM provider, and Kubernetes provider from Kubernetes node objects.
//
// This package is used to classify and extract Vitistack-specific information for use in
// higher-level cluster management and reporting features.
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
	ClusterUidAnnotation         = "vitistack.io/clusteruid"         // The ID of the cluster, this is the uuid in ror
)

// Vitistacktypes provides methods to interrogate Vitistack cluster node metadata.
type Vitistacktypes struct {
}

// NewInterregator creates a new instance of Vitistacktypes for cluster interrogation.
func NewInterregator() *Vitistacktypes {
	return &Vitistacktypes{}
}

// IsTypeOf checks if the provided nodes belong to a Vitistack cluster by inspecting annotations.
func (v Vitistacktypes) IsTypeOf(nodes []v1.Node) bool {
	if _, ok := nodes[0].GetAnnotations()[ClusterNameAnnotation]; ok {
		return true
	}
	return false
}

// GetProvider returns the provider type for the given nodes (Vitistack or Unknown).
func (v Vitistacktypes) GetProvider(nodes []v1.Node) providermodels.ProviderType {
	if v.IsTypeOf(nodes) {
		return providermodels.ProviderTypeVitistack
	}
	return providermodels.ProviderTypeUnknown
}

// GetClusterId returns the cluster ID from the node annotations, or UNKNOWN_CLUSTER_ID if not found.
func (v Vitistacktypes) GetClusterId(nodes []v1.Node) string {
	clusterId, ok := nodes[0].GetAnnotations()[ClusterUidAnnotation]
	if !ok {
		clusterId = providermodels.UNKNOWN_CLUSTER_ID
	}
	return clusterId
}

// GetClusterName returns the cluster name from the node annotations, or UNKNOWN_CLUSTER if not found.
func (v Vitistacktypes) GetClusterName(nodes []v1.Node) string {
	clusterName, ok := nodes[0].GetAnnotations()[ClusterNameAnnotation]
	if !ok {
		clusterName = providermodels.UNKNOWN_CLUSTER
	}
	return clusterName
}

// GetClusterWorkspace returns the workspace name from the node annotations, or "Vitistack" if not found.
func (v Vitistacktypes) GetClusterWorkspace(nodes []v1.Node) string {
	workspace, ok := nodes[0].GetAnnotations()[ClusterWorkspaceAnnotation]
	if !ok {
		workspace = "Vitistack"
	}
	return workspace
}

// GetDatacenter returns a string combining the region and availability zone of the cluster.
func (v Vitistacktypes) GetDatacenter(nodes []v1.Node) string {
	dataCenter := v.GetRegion(nodes) + " " + v.GetAz(nodes)
	return dataCenter
}

// GetRegion returns the region from the node annotations, or UNKNOWN_REGION if not found.
func (v Vitistacktypes) GetRegion(nodes []v1.Node) string {
	region, ok := nodes[0].GetAnnotations()[RegionAnnotation]
	if !ok {
		region = providermodels.UNKNOWN_REGION
	}
	return region
}

// GetAz returns the availability zone from the node annotations, or UNKNOWN_AZ if not found.
func (v Vitistacktypes) GetAz(nodes []v1.Node) string {
	az, ok := nodes[0].GetAnnotations()[AzAnnotation]
	if !ok {
		az = providermodels.UNKNOWN_AZ
	}
	return az
}

// GetVMProvider returns the VM provider from the node annotations, or UNKNOWN_VMPROVIDER if not found.
func (v Vitistacktypes) GetVMProvider(nodes []v1.Node) string {
	vmProvider, ok := nodes[0].GetAnnotations()[VMProviderAnnotation]
	if !ok {
		vmProvider = providermodels.UNKNOWN_VMPROVIDER
	}
	return vmProvider
}

// GetKubernetesProvider returns the Kubernetes provider from the node annotations, or UNKNOWN_KUBERNETES_PROVIDER if not found.
func (v Vitistacktypes) GetKubernetesProvider(nodes []v1.Node) string {
	kubernetesProvider, ok := nodes[0].GetAnnotations()[KubernetesProviderAnnotation]
	if !ok {
		kubernetesProvider = providermodels.UNKNOWN_KUBERNETES_PROVIDER
	}
	return kubernetesProvider
}
