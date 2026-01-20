package vitistackinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

const (
	ClusterNameKey        = "vitistack.io/clustername"        // The name of the cluster
	ClusterWorkspaceKey   = "vitistack.io/clusterworkspace"   // The workspace of the cluster
	RegionKey             = "vitistack.io/region"             // The region of the cluster
	AzKey                 = "vitistack.io/az"                 // The availability zone of the cluster
	MachineProviderKey    = "vitistack.io/vmprovider"         // The VM provider of the cluster
	KubernetesProviderKey = "vitistack.io/kubernetesprovider" // The Kubernetes provider of the cluster
	ClusterIdKey          = "vitistack.io/clusterid"          // The ID of the cluster, this is the uuid in ror

)

var (
	MustBeSet = []string{
		ClusterNameKey,
		ClusterWorkspaceKey,
		RegionKey,
		AzKey,
		MachineProviderKey,
		KubernetesProviderKey,
		ClusterIdKey,
	}
)

type Interregator struct{}

type VitistackProviderinterregator struct {
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
}

func (i Interregator) NewInterregator(nodes []v1.Node) interregatortypes.ClusterInterregator {
	interregator := &VitistackProviderinterregator{
		nodes: nodes,
	}
	if !interregator.MustInitialize() {
		return nil
	}
	return interregator
}

func (v *VitistackProviderinterregator) MustInitialize() bool {
	if v.isOfType {
		return true
	}

	if v.initialized {
		return false
	}

	for _, node := range v.nodes {
		if v.checkIfValid(&node) {
			v.clustername = getValueByKey(&node, ClusterNameKey)
			v.clusterworkspace = getValueByKey(&node, ClusterWorkspaceKey)
			v.region = getValueByKey(&node, RegionKey)
			v.az = getValueByKey(&node, AzKey)
			v.machineprovider = getValueByKey(&node, MachineProviderKey)
			v.kubernetesprovider = getValueByKey(&node, KubernetesProviderKey)
			v.clusterId = getValueByKey(&node, ClusterIdKey)
			v.isOfType = true
			v.initialized = true
			return true
		}
	}

	v.initialized = true
	v.isOfType = false
	return false
}

func (v VitistackProviderinterregator) checkIfValid(node *v1.Node) bool {

	for _, key := range MustBeSet {
		if !checkIfKeyPresent(node, key) {
			return false
		}
	}
	return true
}

func checkIfKeyPresent(node *v1.Node, key string) bool {
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
func getValueByKey(node *v1.Node, key string) string {
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
func (v VitistackProviderinterregator) IsTypeOf() bool {
	return v.MustInitialize()
}

// GetProvider returns the provider type of the nodes.
func (v VitistackProviderinterregator) GetProvider() providermodels.ProviderType {
	if !v.MustInitialize() {
		return providermodels.ProviderTypeUnknown
	}
	return providermodels.ProviderTypeVitistack

}

// GetClusterId returns the cluster ID of the nodes.
func (v VitistackProviderinterregator) GetClusterId() string {
	if !v.MustInitialize() {
		return providermodels.UNKNOWN_CLUSTER_ID
	}
	return v.clusterId
}

// GetClusterName returns the cluster name of the nodes.
func (v VitistackProviderinterregator) GetClusterName() string {
	if !v.MustInitialize() {
		return providermodels.UNKNOWN_CLUSTER
	}
	return v.clustername
}

// GetClusterWorkspace returns the cluster workspace of the nodes.
func (v VitistackProviderinterregator) GetClusterWorkspace() string {
	if !v.MustInitialize() {
		return "Vitistack"
	}
	return v.clusterworkspace
}

// GetDatacenter returns the datacenter of the cluster.
func (v VitistackProviderinterregator) GetDatacenter() string {
	if !v.MustInitialize() {
		return providermodels.UNKNOWN_DATACENTER
	}

	return v.GetRegion() + " " + v.GetAz()

}

// GetRegion returns the region of the cluster.
func (v VitistackProviderinterregator) GetRegion() string {
	if !v.MustInitialize() {
		return providermodels.UNKNOWN_REGION
	}
	return v.region
}

// GetAz returns the availability zone of the cluster.
func (v VitistackProviderinterregator) GetAz() string {
	if !v.MustInitialize() {
		return providermodels.UNKNOWN_AZ
	}
	return v.az
}

// GetVMProvider returns the VM provider of the cluster.
func (v VitistackProviderinterregator) GetMachineProvider() string {
	if !v.MustInitialize() {
		return providermodels.UNKNOWN_MACHINE_PROVIDER
	}
	return v.machineprovider
}

// GetKubernetesProvider returns the Kubernetes provider of the cluster.
func (v VitistackProviderinterregator) GetKubernetesProvider() string {
	if !v.MustInitialize() {
		return providermodels.UNKNOWN_KUBERNETES_PROVIDER
	}
	return v.kubernetesprovider
}
