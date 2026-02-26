package vitistackinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/helpers/kubernetes/metadatahelper"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/factories/interregatorfactory"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"

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
	country            string
}

func (i Interregator) NewInterregator(nodes []v1.Node) interregatortypes.ClusterInterregator {
	interregator := &VitistackProviderinterregator{
		nodes: nodes,
	}
	if !interregator.MustInitialize() {
		return nil
	}

	return interregatorfactory.NewClusterInterregatorFactory(nodes, interregatorfactory.ClusterInterregatorFactoryConfig{
		GetProviderFunc: func() providermodels.ProviderType {
			return interregator.GetProvider()
		},
		GetClusterIdFunc: func() string {
			return interregator.GetClusterId()
		},
		GetClusterNameFunc: func() string {
			return interregator.GetClusterName()
		},
		GetClusterWorkspaceFunc: func() string {
			return interregator.GetClusterWorkspace()
		},
		GetDatacenterFunc: func() string {
			return interregator.GetDatacenter()
		},
		GetAzFunc: func() string {
			return interregator.GetAz()
		},
		GetRegionFunc: func() string {
			return interregator.GetRegion()
		},
		GetCountryFunc: func() string {
			return interregator.GetCountry()
		},
		GetMachineProviderFunc: func() providermodels.ProviderType {
			return interregator.GetMachineProvider()
		},
		GetKubernetesProviderFunc: func() providermodels.ProviderType {
			return interregator.GetKubernetesProvider()
		},
	})

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
			v.clustername, _ = getValueByKey(&node, ClusterNameKey)
			v.clusterworkspace, _ = getValueByKey(&node, ClusterWorkspaceKey)
			v.region, _ = getValueByKey(&node, RegionKey)
			v.az, _ = getValueByKey(&node, AzKey)
			v.machineprovider, _ = getValueByKey(&node, MachineProviderKey)
			v.kubernetesprovider, _ = getValueByKey(&node, KubernetesProviderKey)
			v.clusterId, _ = getValueByKey(&node, ClusterIdKey)
			v.country, _ = getValueByKey(&node, CountryKey)
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
	return metadatahelper.CheckAnnotationOrLabel(node.ObjectMeta, key)
}
func getValueByKey(node *v1.Node, key string) (string, bool) {
	return metadatahelper.GetAnnotationOrLabel(node.ObjectMeta, key)
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

	return v.GetAz() + "." + v.GetRegion() + "." + v.GetCountry()

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

// GetCountry returns the country of the cluster.
func (v VitistackProviderinterregator) GetCountry() string {
	if !v.MustInitialize() {
		return providermodels.UNKNOWN_UNDEFINED
	}
	if v.country == "" {
		return providermodels.DefaultCountry
	}
	return v.country
}

// GetMachineProvider returns the machine provider of the cluster.
func (v VitistackProviderinterregator) GetMachineProvider() providermodels.ProviderType {
	if !v.MustInitialize() {
		return providermodels.ProviderTypeUnknown
	}
	return providermodels.ProviderType(v.machineprovider)
}

// GetKubernetesProvider returns the Kubernetes provider of the cluster.
func (v VitistackProviderinterregator) GetKubernetesProvider() providermodels.ProviderType {
	if !v.MustInitialize() {
		return providermodels.ProviderTypeUnknown
	}
	return providermodels.ProviderType(v.kubernetesprovider)
}
