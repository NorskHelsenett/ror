package configuration

// Common structure for configuration loaders
type Loader struct {
	Parser ConfigParserInterface
	Secret bool
}

// Parse runs the parser to return a json
func (c Loader) Parse() ([]byte, error) {
	return c.Parser.Parse()
}

func (c Loader) IsSecret() bool {
	return c.Secret
}
