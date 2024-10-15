package configuration

import (
	"github.com/NorskHelsenett/ror/internal/clients/helsegitlab"

	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// NewHelsegitlabLoader creates a ConfigLoaderInterface from a git repo
//
// Parameters:
//
//		parser: the parser to use
//	 projectid: the id in helsegitlab
//		path: the path to the file
//	 branch: the branch to download from
func NewHelsegitlabLoader(parser ParserType, projectid int, path string, branch string, vaultClient *vaultclient.VaultClient) ConfigLoaderInterface {
	data, err := helsegitlab.GetFileContent(projectid, path, branch, vaultClient)
	if err != nil {
		rlog.Error("Could not get file from helsegitlab", err)
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
