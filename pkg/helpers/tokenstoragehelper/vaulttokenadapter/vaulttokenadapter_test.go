package vaulttokenadapter

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/tokenstoragehelper"
	"github.com/stretchr/testify/require"
)

type mockVaultClient struct {
	setSecretFn func(string, []byte) (bool, error)
	getSecretFn func(string) (map[string]interface{}, error)
}

func (m mockVaultClient) SetSecret(secretPath string, value []byte) (bool, error) {
	if m.setSecretFn == nil {
		panic("unexpected call to SetSecret")
	}
	return m.setSecretFn(secretPath, value)
}

func (m mockVaultClient) GetSecret(secretPath string) (map[string]interface{}, error) {
	if m.getSecretFn == nil {
		panic("unexpected call to GetSecret")
	}
	return m.getSecretFn(secretPath)
}

func TestSetReturnsErrorWhenVaultClientNil(t *testing.T) {
	adapter := NewVaultStorageAdapter(nil, "secret/path")
	err := adapter.Set(&tokenstoragehelper.KeyStorageProvider{})
	require.EqualError(t, err, "vault client not initialized")
}

func TestSetReturnsErrorWhenMarshalPayloadFails(t *testing.T) {
	originalMarshal := marshalJSON
	t.Cleanup(func() { marshalJSON = originalMarshal })

	marshalJSON = func(interface{}) ([]byte, error) {
		return nil, errors.New("marshal payload failure")
	}

	adapter := NewVaultStorageAdapter(mockVaultClient{
		setSecretFn: func(string, []byte) (bool, error) {
			panic("SetSecret should not be called when marshal fails")
		},
	}, "secret/path")

	err := adapter.Set(&tokenstoragehelper.KeyStorageProvider{})
	require.EqualError(t, err, "marshal payload failure")
}

func TestSetReturnsErrorWhenMarshalBodyFails(t *testing.T) {
	originalMarshal := marshalJSON
	t.Cleanup(func() { marshalJSON = originalMarshal })

	callCount := 0
	marshalJSON = func(v interface{}) ([]byte, error) {
		callCount++
		if callCount == 2 {
			return nil, errors.New("marshal body failure")
		}
		return originalMarshal(v)
	}

	adapter := NewVaultStorageAdapter(mockVaultClient{
		setSecretFn: func(string, []byte) (bool, error) {
			panic("SetSecret should not be called when marshal fails")
		},
	}, "secret/path")

	err := adapter.Set(&tokenstoragehelper.KeyStorageProvider{})
	require.EqualError(t, err, "marshal body failure")
}

func TestSetPropagatesVaultError(t *testing.T) {
	adapter := NewVaultStorageAdapter(mockVaultClient{
		setSecretFn: func(secretPath string, value []byte) (bool, error) {
			require.Equal(t, "secret/path", secretPath)
			require.NotEmpty(t, value)
			return false, errors.New("vault failure")
		},
	}, "secret/path")

	err := adapter.Set(&tokenstoragehelper.KeyStorageProvider{})
	require.EqualError(t, err, "vault failure")
}

func TestSetSucceeds(t *testing.T) {
	captured := struct {
		path string
		body map[string]interface{}
	}{}

	adapter := NewVaultStorageAdapter(mockVaultClient{
		setSecretFn: func(secretPath string, value []byte) (bool, error) {
			captured.path = secretPath
			require.NoError(t, json.Unmarshal(value, &captured.body))
			return true, nil
		},
	}, "secret/path")

	lastRotation := time.Unix(1697040000, 0).UTC()

	ks := &tokenstoragehelper.KeyStorageProvider{
		LastRotation:     lastRotation,
		RotationInterval: time.Hour,
		NumKeys:          1,
		Keys: map[int]tokenstoragehelper.Key{
			1: {
				KeyID:        "kid",
				AlgorithmKey: "RS256",
			},
		},
	}

	err := adapter.Set(ks)
	require.NoError(t, err)
	require.Equal(t, "secret/path", captured.path)

	configValue, ok := captured.body["data"].(map[string]interface{})
	require.True(t, ok)
	storedConfig, ok := configValue["config"].(string)
	require.True(t, ok)

	var restored tokenstoragehelper.KeyStorageProvider
	require.NoError(t, json.Unmarshal([]byte(storedConfig), &restored))
	require.Equal(t, ks.LastRotation, restored.LastRotation)
	require.Equal(t, ks.RotationInterval, restored.RotationInterval)
	require.Equal(t, ks.NumKeys, restored.NumKeys)
	require.Equal(t, ks.Keys, restored.Keys)
}

func TestGetReturnsErrorWhenVaultClientNil(t *testing.T) {
	adapter := NewVaultStorageAdapter(nil, "secret/path")
	_, err := adapter.Get()
	require.EqualError(t, err, "vault client not initialized")
}

func TestGetReturnsErrorWhenVaultClientFails(t *testing.T) {
	adapter := NewVaultStorageAdapter(mockVaultClient{
		getSecretFn: func(string) (map[string]interface{}, error) {
			return nil, errors.New("vault failure")
		},
	}, "secret/path")

	_, err := adapter.Get()
	require.EqualError(t, err, "vault failure")
}

func TestGetReturnsErrorOnInvalidDataMap(t *testing.T) {
	adapter := NewVaultStorageAdapter(mockVaultClient{
		getSecretFn: func(string) (map[string]interface{}, error) {
			return map[string]interface{}{
				"data": "not-a-map",
			}, nil
		},
	}, "secret/path")

	_, err := adapter.Get()
	require.EqualError(t, err, "invalid data format in vault secret")
}

func TestGetReturnsErrorOnInvalidConfigType(t *testing.T) {
	adapter := NewVaultStorageAdapter(mockVaultClient{
		getSecretFn: func(string) (map[string]interface{}, error) {
			return map[string]interface{}{
				"data": map[string]interface{}{
					"config": 123,
				},
			}, nil
		},
	}, "secret/path")

	_, err := adapter.Get()
	require.EqualError(t, err, "invalid data format in vault secret")
}

func TestGetReturnsErrorOnUnmarshalFailure(t *testing.T) {
	adapter := NewVaultStorageAdapter(mockVaultClient{
		getSecretFn: func(string) (map[string]interface{}, error) {
			return map[string]interface{}{
				"data": map[string]interface{}{
					"config": "not-json",
				},
			}, nil
		},
	}, "secret/path")

	_, err := adapter.Get()
	require.ErrorContains(t, err, "invalid character")
}

func TestGetSucceeds(t *testing.T) {
	lastRotation := time.Unix(1697040000, 0).UTC()

	expected := tokenstoragehelper.KeyStorageProvider{
		LastRotation:     lastRotation,
		RotationInterval: time.Hour,
		NumKeys:          2,
		Keys: map[int]tokenstoragehelper.Key{
			0: {
				KeyID:        "old",
				AlgorithmKey: "RS256",
			},
			1: {
				KeyID:        "current",
				AlgorithmKey: "RS256",
			},
		},
	}

	payload, err := json.Marshal(expected)
	require.NoError(t, err)

	adapter := NewVaultStorageAdapter(mockVaultClient{
		getSecretFn: func(string) (map[string]interface{}, error) {
			return map[string]interface{}{
				"data": map[string]interface{}{
					"config": string(payload),
				},
			}, nil
		},
	}, "secret/path")

	actual, err := adapter.Get()
	require.NoError(t, err)
	require.Equal(t, expected.LastRotation, actual.LastRotation)
	require.Equal(t, expected.RotationInterval, actual.RotationInterval)
	require.Equal(t, expected.NumKeys, actual.NumKeys)
	require.Equal(t, expected.Keys, actual.Keys)
}
