package configuration

type LayerConfig struct {
	Name    string
	Tier    int
	Order   int
	Content []byte
	Secret  bool
}

// NewConfigurationLayer creates a new configuration layer
//
// Parameters:
//
//	name: the name of the layer
//	tier: the tier of the layer
//	order: the order of the layer
//	content: the content of the layer represented as a ConfigLoaderInterface
func NewConfigurationLayer(name string, tier int, order int, content ConfigLoaderInterface) ConfigLayerInterface {
	if content == nil {
		return nil
	}
	secret := content.IsSecret()

	conf, err := content.Parse()
	if err != nil {
		return nil
	}
	ret := LayerConfig{
		Name:    name,
		Tier:    tier,
		Order:   order,
		Content: conf,
		Secret:  secret,
	}
	return ret
}

// GetName returns the name of the layer
func (c LayerConfig) GetName() string {
	return c.Name
}

// GetTier returns the tier of the layer
func (c LayerConfig) GetTier() int {
	return c.Tier
}

// GetOrder returns the order of the layer
func (c LayerConfig) GetOrder() int {
	return c.Order
}

// GetContent returns the content of the layer
func (c LayerConfig) GetContent() []byte {
	return c.Content
}

// IsSecret returns is the layer contains secrets
func (c LayerConfig) IsSecret() bool {
	return c.Secret
}
