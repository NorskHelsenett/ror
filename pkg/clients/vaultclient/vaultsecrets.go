package vaultclient

import (
	"errors"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/hashicorp/vault-client-go"
)

// Deprecated: Use the VaultClient method instead
func GetSecret(secretPath string) (map[string]interface{}, error) {
	if secretPath == "" {
		return nil, errors.New("secret path is nil or empty")
	}

	data, err := vaultClient.Client.Read(vaultClient.Context, secretPath)
	if err != nil {
		var vaultError *vault.ResponseError
		errors.As(err, &vaultError)
		msg := fmt.Sprintf("Could not get secret, StatusCode: %d", vaultError.StatusCode)
		if vaultError.StatusCode == 404 {
			rlog.Info(msg)
		} else {
			rlog.Error(msg, err)
		}
		return nil, fmt.Errorf(msg)
	}
	if data != nil {
		return data.Data, nil
	}

	return nil, nil
}

// Deprecated: Use the VaultClient method instead
func GetSecretValue(secretPath string, key string) (string, error) {
	if secretPath == "" {
		return "", errors.New("secret path is nil or empty")
	}

	data, err := vaultClient.Client.Read(vaultClient.Context, secretPath)
	if err != nil {
		var err2 *vault.ResponseError
		errors.As(err, &err2)
		msg := fmt.Sprintf("Could not get secret, StatusCode: %d", err2.StatusCode)
		if err2.StatusCode == 404 {
			rlog.Info(msg)
		} else {
			rlog.Error(msg, err)
		}
		return "", fmt.Errorf(msg)
	}
	if data != nil {
		vaultval, _ := data.Data["data"].(map[string]interface{})
		vaultkey, _ := vaultval[key].(string)
		return vaultkey, nil
	}

	return "", nil
}

func (vc VaultClient) GetSecret(secretPath string) (map[string]interface{}, error) {
	if secretPath == "" {
		return nil, errors.New("secret path is nil or empty")
	}

	data, err := vc.Client.Read(vc.Context, secretPath)
	if err != nil {
		var vaultError *vault.ResponseError
		errors.As(err, &vaultError)
		msg := fmt.Sprintf("Could not get secret, StatusCode: %d", vaultError.StatusCode)
		if vaultError.StatusCode == 404 {
			rlog.Info(msg)
		} else {
			rlog.Error(msg, err)
		}
		return nil, fmt.Errorf(msg)
	}
	if data != nil {
		return data.Data, nil
	}

	return nil, nil
}

func (vc VaultClient) GetSecretValue(secretPath string, key string) (string, error) {
	if secretPath == "" {
		return "", errors.New("secret path is nil or empty")
	}

	data, err := vc.Client.Read(vc.Context, secretPath)
	if err != nil {
		var err2 *vault.ResponseError
		errors.As(err, &err2)
		msg := fmt.Sprintf("Could not get secret, StatusCode: %d", err2.StatusCode)
		if err2.StatusCode == 404 {
			rlog.Info(msg)
		} else {
			rlog.Error(msg, err)
		}
		return "", fmt.Errorf(msg)
	}
	if data != nil {
		vaultval, _ := data.Data["data"].(map[string]interface{})
		vaultkey, _ := vaultval[key].(string)
		return vaultkey, nil
	}

	return "", nil
}

func (vc VaultClient) SetSecret(secretPath string, value []byte) (bool, error) {
	if len(secretPath) < 1 {
		return false, fmt.Errorf("could not set secret, secret path is empty")
	}

	secret, err := vc.Client.WriteFromBytes(vc.Context, secretPath, value)
	if err != nil {
		msg := fmt.Sprintf("could not set secret on path: %s", secretPath)
		rlog.Error(msg, err)
		return false, errors.New(msg)
	}

	if secret.Data != nil {
		return true, nil
	}

	return false, errors.New("could not set secret")
}
