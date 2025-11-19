// Package vaulttokenadapter provides a HashiCorp Vault implementation of the StorageAdapter
// interface for persisting JWT signing keys.
//
// This adapter stores the KeyStorageProvider in Vault's key-value store, allowing for
// secure, centralized key management across multiple application instances.
//
// # Usage
//
//	vaultClient := vaultclient.NewVaultClient(config)
//	adapter := vaulttokenadapter.NewVaultStorageAdapter(vaultClient, "secret/data/jwt-keys")
//	tokenstoragehelper.Init(adapter)
//
// The keys are stored as JSON in Vault under the specified secret path.
package vaulttokenadapter

import (
	"encoding/json"
	"errors"

	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"
	"github.com/NorskHelsenett/ror/pkg/helpers/tokenstoragehelper"
)

// VaultStorageAdapter implements the StorageAdapter interface using HashiCorp Vault
// as the persistence backend. It stores and retrieves KeyStorageProvider data from
// Vault's key-value secrets engine.
type vaultClient interface {
	SetSecret(secretPath string, value []byte) (bool, error)
	GetSecret(secretPath string) (map[string]interface{}, error)
}

// ensure the HashiCorp vault client satisfies the interface used by the adapter
var _ vaultClient = (*vaultclient.VaultClient)(nil)

type VaultStorageAdapter struct {
	vaultclient vaultClient
	secretPath  string
}

// NewVaultStorageAdapter creates a new VaultStorageAdapter instance.
//
// Parameters:
//   - vaultclient: An initialized VaultClient for interacting with Vault
//   - secretPath: The path in Vault where the key storage will be persisted
//     (e.g., "secret/data/jwt-keys")
//
// Returns a configured VaultStorageAdapter ready for use.
func NewVaultStorageAdapter(vaultclient vaultClient, secretPath string) *VaultStorageAdapter {
	return &VaultStorageAdapter{
		vaultclient: vaultclient,
		secretPath:  secretPath,
	}
}

// Set persists the KeyStorageProvider to Vault.
// The key storage data is marshaled to JSON and stored in a nested structure
// compatible with Vault's KV v2 secrets engine.
//
// Returns an error if:
//   - The vault client is not initialized
//   - JSON marshaling fails
//   - The Vault write operation fails
var marshalJSON = json.Marshal

func (v *VaultStorageAdapter) Set(ks *tokenstoragehelper.KeyStorageProvider) error {
	if v.vaultclient == nil {
		return errors.New("vault client not initialized")
	}

	payload, err := marshalJSON(ks)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"data": map[string]interface{}{
			"config": string(payload),
		},
	}

	body, err := marshalJSON(data)
	if err != nil {
		return err
	}

	_, err = v.vaultclient.SetSecret(v.secretPath, body)
	return err
}

// Get retrieves the KeyStorageProvider from Vault.
// It reads the secret from the configured path, extracts the JSON data,
// and unmarshals it into a KeyStorageProvider struct.
//
// Returns an error if:
//   - The vault client is not initialized
//   - The secret cannot be read from Vault
//   - The secret data format is invalid
//   - JSON unmarshaling fails
func (v *VaultStorageAdapter) Get() (tokenstoragehelper.KeyStorageProvider, error) {
	if v.vaultclient == nil {
		return tokenstoragehelper.KeyStorageProvider{}, errors.New("vault client not initialized")
	}
	secretData, err := v.vaultclient.GetSecret(v.secretPath)
	if err != nil {
		return tokenstoragehelper.KeyStorageProvider{}, err
	}
	data, ok := secretData["data"].(map[string]interface{})
	if !ok {
		return tokenstoragehelper.KeyStorageProvider{}, errors.New("invalid data format in vault secret")
	}

	dataStr, ok := data["config"].(string)
	if !ok {
		return tokenstoragehelper.KeyStorageProvider{}, errors.New("invalid data format in vault secret")
	}
	var ks tokenstoragehelper.KeyStorageProvider
	err = json.Unmarshal([]byte(dataStr), &ks)
	if err != nil {
		return tokenstoragehelper.KeyStorageProvider{}, err
	}
	return ks, nil
}
