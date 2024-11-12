package configuration

// ConfigurationsGenerator is the main interface for generating configurations
type ConfigGeneratorInterface interface {
	GenerateConfig() ([]byte, error)
	GenerateConfigYaml() ([]byte, error)
	AddConfiguration(configuration ConfigLayerInterface)
}

// ConfigLayerInterface is the interface for a configuration layer
type ConfigLayerInterface interface {
	GetName() string
	GetTier() int
	GetOrder() int
	GetContent() []byte
	IsSecret() bool
}

// ConfigLoaderInterface is the interface for loading a configuration
type ConfigLoaderInterface interface {
	Parse() ([]byte, error)
	IsSecret() bool
}

// ConfigParserInterface is the interface for parsing a configuration to json
type ConfigParserInterface interface {
	Parse() ([]byte, error)
}

type SecretStruct struct {
	VaultPath string
	VaultKey  string
	JsonPath  string
}

type ParserType string

const (
	ParserTypeJson ParserType = "json"
	ParserTypeYaml ParserType = "yaml"
)
