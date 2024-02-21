package providers

type Provider struct {
	Name     string       `json:"name"`
	Type     ProviderType `json:"type"`
	Disabled bool         `json:"disabled"`
}

type ProviderKubernetesVersion struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Disabled bool   `json:"disabled"`
}
