package configuration

import "github.com/NorskHelsenett/ror/pkg/rlog"

// NewFuncLoader creates a ConfigLoaderInterface from a function
//
// Parameters:
//
//	parser: the parser to use
//	function: the function to call
//	  - the function must return a byte array and an error
func NewFuncLoader(parser ParserType, function func() ([]byte, error)) ConfigLoaderInterface {
	data, err := function()
	if err != nil {
		rlog.Error("func returned", err)
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
