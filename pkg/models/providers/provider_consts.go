package providers

type ProviderType string

const (
	ProviderTypeUnknown ProviderType = "unknown"
	ProviderTypeTanzu   ProviderType = "tanzu"
	ProviderTypeAks     ProviderType = "aks"
	ProviderTypeK3d     ProviderType = "k3d"
	ProviderTypeKind    ProviderType = "kind"
	ProviderTypeGke     ProviderType = "gke"
	ProviderTypeTalos   ProviderType = "talos"
)

// String returns the string representation of the ProviderType.
func (p ProviderType) String() string {
	return string(p)
}
