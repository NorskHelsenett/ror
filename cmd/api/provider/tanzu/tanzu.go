package tanzu

import providertypes "github.com/NorskHelsenett/ror/cmd/api/provider/types"

type provider struct {
}

const (
	ProviderName = "Tanzu"
)

func NewTanzuProvider() *provider {
	return &provider{}
}

func (p *provider) GetName() string {
	return ProviderName
}
func (p *provider) GetConfigurations(page string) map[string]providertypes.ProviderConfig {
	if page == "cluster.create.1" {
		return map[string]providertypes.ProviderConfig{
			"datacenter": {
				Name:     "Datacenter",
				Order:    1,
				Query:    "",
				Disabled: false,
			},
			"tanzuNamespace": {
				Name:     "Tanzu Namespace",
				Order:    2,
				Query:    "datacenter",
				Disabled: false,
			},
			"Machine-class": {
				Name:     "Machine Class",
				Order:    3,
				Query:    "tanzuNamespace",
				Disabled: false,
			},
			"Storage-class": {
				Name:     "Storage Class",
				Order:    4,
				Query:    "tanzuNamespace",
				Disabled: false,
			},
		}
	}
	return nil
}

func (p *provider) GetConfigOptions(configname string, options ...string) []providertypes.ProviderConfigOptions {
	//machineClass

	if configname == "Machine-class" {
		return []providertypes.ProviderConfigOptions{
			{
				Name:     "Small ",
				Default:  false,
				Disabled: false,
			},
			{
				Name:     "Medium (2x cpu 8gb ram)",
				Value:    "best-effort-medium",
				Default:  true,
				Disabled: false,
			},
			{
				Name:     "Large (4x cpu 16gb ram)",
				Value:    "best-effort-large",
				Default:  false,
				Disabled: false,
			},
		}
	}
	return nil
}
