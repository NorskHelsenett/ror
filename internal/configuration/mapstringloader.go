package configuration

import "github.com/NorskHelsenett/ror/pkg/helpers/jsonhelper"

// MapStringLoader is a ConfigLoaderInterface that loads a config from a map[string]string
//
// Parameters:
//
//	Parser: the parser to use
//	Data: the string to parse
func NewMapStringLoader(inn map[string]string) ConfigLoaderInterface {
	data := jsonhelper.MapToJson(inn)
	parserinterface := NewJsonParser(data)
	ret := Loader{
		Parser: parserinterface,
		Secret: false,
	}

	return ret
}
