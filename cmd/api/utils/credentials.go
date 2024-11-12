package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/NorskHelsenett/ror/cmd/api/apiconnections"
	"github.com/NorskHelsenett/ror/internal/models/ldapmodels"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
)

func GetCredsFromVault() {
	_, _ = GetApikeySalt()
	_, _ = GetLdapConfigs()
}

func GetApikeySalt() (string, error) {
	secretPath := "secret/data/v1.0/ror/config/common" // #nosec G101 Jest the path to the token file in the secrets engine
	vaultData, err := apiconnections.VaultClient.GetSecret(secretPath)
	if err != nil {
		return "", errors.New("could not extract apikeySalt from vault")
	}
	commonConfig, ok := vaultData["data"].(map[string]interface{})
	if !ok {
		rlog.Error("could not get apikey salt", fmt.Errorf("data type assertion failed: %T %#v", vaultData["data"], vaultData["data"]))
		return "", errors.New("could not extract api key salt from vault.data")
	}

	salt, ok := commonConfig["apikeySalt"].(string)
	if !ok {
		rlog.Fatal("could not get apikey salt", fmt.Errorf("value apikeySalt not set in Vault [TODO: docref]"))
	}

	if len(salt) == 0 {
		rlog.Error("could not extract api key salt", fmt.Errorf("value is empty"))
		panic("missing apikey salt")
	}

	viper.Set(configconsts.API_KEY_SALT, salt)
	return salt, err
}

func GetLdapConfigs() (ldapmodels.LdapConfigs, error) {
	secretPath := "secret/data/v1.0/ror/config/ldap" // #nosec G101 Jest the path to the token file in the secrets engine
	vaultData, err := apiconnections.VaultClient.GetSecret(secretPath)
	if err != nil {
		return ldapmodels.LdapConfigs{}, errors.New("could not extract ldap-domains from vault")
	}

	ldapConfigsMap, ok := vaultData["data"].(map[string]interface{})
	if !ok {
		rlog.Error(fmt.Sprintf("data type assertion failed, vault data: %v", ldapConfigsMap), nil)
		return ldapmodels.LdapConfigs{}, errors.New("could not extract ldap-domains from vault.data")
	}

	ldapConfigsJson, err := json.Marshal(ldapConfigsMap["data"])
	if err != nil {
		rlog.Error("could not convert from map to json", nil)
		return ldapmodels.LdapConfigs{}, errors.New("could not convert from map to json")
	}

	var ldapConfigs ldapmodels.LdapConfigs
	err = json.Unmarshal(ldapConfigsJson, &ldapConfigs)
	if err != nil {
		rlog.Error("could not convert from json to ldapConfigs", nil)
		return ldapmodels.LdapConfigs{}, errors.New("could not convert from json to ldapConfigs")
	}

	viper.Set(configconsts.LDAP_CONFIGS, ldapConfigs)
	return ldapConfigs, nil
}
