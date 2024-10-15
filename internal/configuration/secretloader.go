package configuration

import (
	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/helpers/jsonhelper"
)

// NewSecretLoader creates a ConfigLoaderInterface from a vault secret
//
// Parameters:
//
//	vaultpath: the path in vault
//	vaultkey: the key to use
//	jsonpath: the json path to set "test.auth.username"
func NewSecretLoader(vaultpath string, vaultkey string, jsonpath string, vaultClient *vaultclient.VaultClient) ConfigLoaderInterface {
	secret, err := vaultClient.GetSecretValue(vaultpath, vaultkey)
	if err != nil {
		rlog.Error("Could not secret from vault", err)
		return nil
	}
	data := jsonhelper.StringToJson(jsonpath, secret)

	parserinterface := NewJsonParser(data)
	ret := Loader{
		Parser: parserinterface,
		Secret: true,
	}

	return ret
}
