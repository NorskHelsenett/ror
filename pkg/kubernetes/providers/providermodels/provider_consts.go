package providermodels

type ProviderType string

const (
	UNKNOWN_REGION              string = "unknown-region"
	UNKNOWN_DATACENTER          string = "unknown-datacenter"
	UNKNOWN_WORKSPACE           string = "unknown-workspace"
	UNKNOWN_VMPROVIDER          string = "unknown-virtualization-provider"
	UNKNOWN_MACHINE_PROVIDER    string = "unknown-machine-provider"
	UNKNOWN_KUBERNETES_PROVIDER string = "unknown-kubernetes-provider"
	UNKNOWN_CLUSTER             string = "unknown-cluster"
	UNKNOWN_CLUSTER_ID          string = "unknown-cluster-id"
	UNKNOWN_AZ                  string = "unknown-az"

	ProviderTypeUnknown   ProviderType = "unknown"
	ProviderTypeTanzu     ProviderType = "tanzu"
	ProviderTypeAks       ProviderType = "aks"
	ProviderTypeK3d       ProviderType = "k3d"
	ProviderTypeKind      ProviderType = "kind"
	ProviderTypeGke       ProviderType = "gke"
	ProviderTypeTalos     ProviderType = "talos"
	ProviderTypeVitistack ProviderType = "vitistack"
)

// String returns the string representation of the ProviderType.
func (p ProviderType) String() string {
	return string(p)
}
