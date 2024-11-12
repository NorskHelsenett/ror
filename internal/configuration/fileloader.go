package configuration

import (
	"os"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// NewFileLoader creates a ConfigLoaderInterface from a file
//
// Parameters:
//
//	parser: the parser to use
//	path: the path to the file
//	WARNING: this function should not be used with user input
func NewFileLoader(parser ParserType, path string) ConfigLoaderInterface {

	_, err := os.Stat(path)
	if err != nil {
		rlog.Error("File does not exist", err, rlog.Any("path", path))
		return nil
	}
	data, err := os.ReadFile(path) // #nosec G304 - we are not using user input
	if err != nil {
		return nil
	}
	var parserinterface ConfigParserInterface
	switch parser {
	case ParserTypeJson:
		parserinterface = NewJsonParser(data)
	case ParserTypeYaml:
		parserinterface = NewYamlParser(data)
	default:
		return nil
	}
	ret := Loader{
		Parser: parserinterface,
		Secret: false,
	}

	return ret
}
