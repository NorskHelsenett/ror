package configuration

import (
	"encoding/json"
	"errors"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	jsonpatch "github.com/evanphx/json-patch/v5"
	"golang.org/x/exp/slices"
	k8syaml "sigs.k8s.io/yaml"
)

type ConfigurationsGenerator struct {
	Configurations []ConfigLayerInterface
}

// NewConfigurationsGenerator creates a new empty ConfigurationsGenerator
func NewConfigurationsGenerator() ConfigGeneratorInterface {
	ret := ConfigurationsGenerator{}
	return &ret
}

// AddConfiguration adds a configuration layer to the ConfigurationsGenerator
// The configuration layer will be added to the list, but its position will be determined by the tier and order
// The lower tier, the higher priority
// The lower order, the higher priority within the same tier
func (c *ConfigurationsGenerator) AddConfiguration(config ConfigLayerInterface) {
	config.IsSecret()
	if config == nil {
		rlog.Error("Configuration is nil", nil)
		return
	}
	if !c.detectDuplacateTierOrder(config) {
		c.Configurations = append(c.Configurations, config)
	}
}

// detectDuplacateTierOrder checks if there is a configuration with the same tier and order
func (c ConfigurationsGenerator) detectDuplacateTierOrder(config ConfigLayerInterface) bool {
	for _, co := range c.Configurations {
		if co.GetTier() == config.GetTier() && co.GetOrder() == config.GetOrder() {
			return true
		}
	}
	return false
}

// SortConfigurations sorts the configurations by tier and order
func (c *ConfigurationsGenerator) SortConfigurations() []ConfigLayerInterface {
	var tiers []int
	tierorders := map[int][]int{}
	sortingmap := make(map[int]map[int]ConfigLayerInterface)
	var returnslice []ConfigLayerInterface

	for _, configLayer := range c.Configurations {
		tier := configLayer.GetTier()
		order := configLayer.GetOrder()

		if !slices.Contains(tiers, tier) {
			tiers = append(tiers, tier)
		}

		if len(tierorders[tier]) == 0 {
			tierorders[tier] = make([]int, 0)
		}
		tierorders[tier] = append(tierorders[tier], order)

		if len(sortingmap[tier]) == 0 {
			sortingmap[tier] = make(map[int]ConfigLayerInterface)
		}
		sortingmap[tier][order] = configLayer
	}
	slices.Sort(tiers)
	for _, t := range tiers {
		tierorder := tierorders[t]
		slices.Sort(tierorder)
		for _, so := range tierorder {
			returnslice = append(returnslice, sortingmap[t][so])
		}
	}

	return returnslice
}

// GenerateConfig generates a json configuration from the configurations in the ConfigurationsGenerator
func (c *ConfigurationsGenerator) GenerateConfig() ([]byte, error) {
	conf := []byte("{}")
	// Finn ut rekkefølge
	sortedconfigs := c.SortConfigurations()
	for _, config := range sortedconfigs {
		conf, _ = Merge(conf, config.GetContent())
	}
	return conf, nil
}

// GenerateConfigYaml generates a yaml configuration from the configurations in the ConfigurationsGenerator
func (c *ConfigurationsGenerator) GenerateConfigYaml() ([]byte, error) {
	conf := []byte("{}")
	// Finn ut rekkefølge
	sortedconfigs := c.SortConfigurations()
	for _, config := range sortedconfigs {
		conf, _ = Merge(conf, config.GetContent())
	}
	returnyaml, err := k8syaml.JSONToYAML(conf)
	if err != nil {
		return nil, err
	}
	return returnyaml, nil
}

// Merge merges two jsons
func Merge(input []byte, patch []byte) ([]byte, error) {
	merged, err := jsonpatch.MergePatch(input, patch)

	if !json.Valid(merged) {
		return nil, errors.New("invalid json")
	}

	return merged, err
}
