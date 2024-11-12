package types

type ProviderConfig struct {
	Name     string
	Page     []string
	Order    int
	Query    string
	Disabled bool
}

type ProviderConfigOptions struct {
	Name     string
	Value    string
	Default  bool
	Disabled bool
}

type Provider interface {
	GetName() string
	GetConfigurations(page string) map[string]ProviderConfig
	GetConfigOptions(configname string, options ...string) []ProviderConfigOptions
}
