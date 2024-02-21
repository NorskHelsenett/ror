package providers

type ProviderType string

const (
	ProviderTypeUnknown ProviderType = "Unknown"
	ProviderTypeTanzu   ProviderType = "Tanzu"
	ProviderTypeAks     ProviderType = "AKS"
	ProviderTypeK3d     ProviderType = "K3D"
	ProviderTypeKind    ProviderType = "Kind"
)
