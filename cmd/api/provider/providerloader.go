package provider

import (
	"github.com/NorskHelsenett/ror/cmd/api/provider/tanzu"
	providertypes "github.com/NorskHelsenett/ror/cmd/api/provider/types"
	"slices"

	"github.com/NorskHelsenett/ror/pkg/models/providers"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

const (
	ProviderType = "provider"
)

type providerLoader struct {
	modules  []providers.ProviderType
	Provider map[providers.ProviderType]providertypes.Provider
}

func NewProviderloader(modules []providers.ProviderType) *providerLoader {
	providerloader := &providerLoader{
		modules:  modules,
		Provider: make(map[providers.ProviderType]providertypes.Provider),
	}

	if len(modules) == 0 {
		rlog.Info("no provider modules to load")
		return providerloader
	}

	if slices.Contains(modules, providers.ProviderTypeTanzu) {
		providerloader.Provider[providers.ProviderTypeTanzu] = tanzu.NewTanzuProvider()
		rlog.Info("loading provider", rlog.Any("provider", providerloader.Provider[providers.ProviderTypeTanzu].GetName()), rlog.Any("providerId", providers.ProviderTypeTanzu))
	}

	return providerloader
}

func (pl *providerLoader) GetProvider(providerType providers.ProviderType) (providertypes.Provider, bool) {
	provider, ok := pl.Provider[providerType]
	return provider, ok
}

func (pl *providerLoader) GetProviders() map[providers.ProviderType]providertypes.Provider {
	return pl.Provider
}
func (pl *providerLoader) GetProviderIds() []providers.ProviderType {
	return pl.modules
}
func (pl *providerLoader) IsProviderLoaded(module providers.ProviderType) bool {
	return slices.Contains(pl.modules, module)
}
