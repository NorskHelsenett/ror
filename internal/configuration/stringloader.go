package configuration

// StringLoader is a ConfigLoaderInterface that loads a config from a string
//
// Parameters:
//
//	Parser: the parser to use
//	Data: the string to parse
func NewStringLoader(parser ParserType, data string) ConfigLoaderInterface {
	var parserinterface ConfigParserInterface
	switch parser {
	case ParserTypeJson:
		parserinterface = NewJsonParser([]byte(data))
	case ParserTypeYaml:
		parserinterface = NewYamlParser([]byte(data))
	default:
		return nil
	}
	ret := Loader{
		Parser: parserinterface,
		Secret: false,
	}

	return ret
}
