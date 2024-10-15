package configuration

import (
	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/helpers/jsonhelper"
)

// NewSecretMapLoader creates a ConfigLoaderInterface from a map of SecretStruct
func NewSecretMapLoader(inn []SecretStruct, vaultClient *vaultclient.VaultClient) ConfigLoaderInterface {
	jsonmap := make(map[string]string)
	for _, secretStruct := range inn {
		//TODO: Replave with method from vaultclient
		secret, err := vaultClient.GetSecretValue(secretStruct.VaultPath, secretStruct.VaultKey)
		if err != nil {
			rlog.Error("Could not get file from vault", err, rlog.String("vaultPath", secretStruct.VaultPath), rlog.String("vaultKey", secretStruct.VaultKey))
			continue
		}
		jsonmap[secretStruct.JsonPath] = secret
	}

	data := jsonhelper.MapToJson(jsonmap)
	parserinterface := NewJsonParser(data)
	ret := Loader{
		Parser: parserinterface,
		Secret: true,
	}

	return ret
}
