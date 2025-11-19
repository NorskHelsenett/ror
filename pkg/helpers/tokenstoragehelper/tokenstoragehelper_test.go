package tokenstoragehelper

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"io"
	"math/big"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/stretchr/testify/require"
)

type mockStorageAdapter struct {
	setFn func(*KeyStorageProvider) error
	getFn func() (KeyStorageProvider, error)
}

func (m *mockStorageAdapter) Set(ks *KeyStorageProvider) error {
	if m.setFn == nil {
		panic("unexpected call to Set")
	}
	return m.setFn(ks)
}

func (m *mockStorageAdapter) Get() (KeyStorageProvider, error) {
	if m.getFn == nil {
		panic("unexpected call to Get")
	}
	return m.getFn()
}

func resetGlobals() {
	nowFunc = time.Now
	generateKeyFn = GenerateKey
	randomIntFunc = rand.Int
	sleepFunc = time.Sleep
	keyStorage = KeyStorageProvider{}
}

func generateRSAKey(t *testing.T, bits int) *rsa.PrivateKey {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, bits)
	require.NoError(t, err)
	return key
}

func TestNewTokenKeyStorageSuccess(t *testing.T) {
	resetGlobals()
	fixedNow := time.Unix(100, 0)
	nowFunc = func() time.Time { return fixedNow }
	defer func() { nowFunc = time.Now }()
	privKey := generateRSAKey(t, 1024)

	adapter := &mockStorageAdapter{
		getFn: func() (KeyStorageProvider, error) {
			return KeyStorageProvider{
				LastRotation:     fixedNow,
				RotationInterval: time.Hour,
				NumKeys:          1,
				Keys: map[int]Key{
					1: {
						KeyID:        "kid",
						PrivateKey:   privKey,
						AlgorithmKey: "RS256",
					},
				},
			}, nil
		},
		setFn: func(*KeyStorageProvider) error {
			return nil
		},
	}

	storage, err := NewTokenKeyStorage(adapter)
	require.NoError(t, err)

	provider, ok := storage.(*KeyStorageProvider)
	require.True(t, ok)
	require.Equal(t, fixedNow, provider.LastRotation)
	require.Equal(t, time.Hour, provider.RotationInterval)
	require.Equal(t, "kid", provider.Keys[1].KeyID)
}

func TestNewTokenKeyStorageLoadFails(t *testing.T) {
	resetGlobals()
	adapter := &mockStorageAdapter{
		getFn: func() (KeyStorageProvider, error) {
			return KeyStorageProvider{}, errors.New("load failed")
		},
		setFn: func(*KeyStorageProvider) error {
			return nil
		},
	}

	storage, err := NewTokenKeyStorage(adapter)
	require.Nil(t, storage)
	require.EqualError(t, err, "load failed")
}

func TestKeyStorageProviderSaveNoAdapter(t *testing.T) {
	resetGlobals()
	var provider KeyStorageProvider
	err := provider.Save()
	require.EqualError(t, err, "no storage provider set")
}

func TestKeyStorageProviderSaveSuccess(t *testing.T) {
	resetGlobals()
	provider := KeyStorageProvider{}
	adapter := &mockStorageAdapter{
		setFn: func(ks *KeyStorageProvider) error {
			require.Equal(t, &provider, ks)
			return nil
		},
		getFn: func() (KeyStorageProvider, error) {
			return KeyStorageProvider{}, nil
		},
	}
	provider.storageAdapter = adapter

	err := provider.Save()
	require.NoError(t, err)
}

func TestKeyStorageProviderLoadNoAdapter(t *testing.T) {
	resetGlobals()
	var provider KeyStorageProvider
	err := provider.Load()
	require.EqualError(t, err, "no storage provider set")
}

func TestKeyStorageProviderLoadSuccess(t *testing.T) {
	resetGlobals()
	provider := KeyStorageProvider{}
	expected := KeyStorageProvider{
		LastRotation:     time.Unix(50, 0),
		RotationInterval: time.Minute,
		NumKeys:          2,
		Keys: map[int]Key{
			0: {KeyID: "zero"},
			1: {KeyID: "one"},
			2: {KeyID: "two"},
		},
	}
	adapter := &mockStorageAdapter{
		setFn: func(*KeyStorageProvider) error {
			return nil
		},
		getFn: func() (KeyStorageProvider, error) {
			return expected, nil
		},
	}
	provider.storageAdapter = adapter

	err := provider.Load()
	require.NoError(t, err)
	require.Equal(t, expected.LastRotation, provider.LastRotation)
	require.Equal(t, expected.RotationInterval, provider.RotationInterval)
	require.Equal(t, expected.NumKeys, provider.NumKeys)
	require.Equal(t, expected.Keys, provider.Keys)
}

func TestKeyStorageProviderRotateNoNeed(t *testing.T) {
	resetGlobals()
	fixedNow := time.Unix(200, 0)
	nowFunc = func() time.Time { return fixedNow }
	defer func() { nowFunc = time.Now }()

	provider := KeyStorageProvider{
		LastRotation:     fixedNow,
		RotationInterval: time.Hour,
		NumKeys:          2,
		Keys: map[int]Key{
			0: {KeyID: "zero"},
			1: {KeyID: "one"},
			2: {KeyID: "two"},
		},
	}

	rotated := provider.Rotate(false)
	require.False(t, rotated)
	require.Equal(t, fixedNow, provider.LastRotation)
}

func TestKeyStorageProviderRotateWithGeneration(t *testing.T) {
	resetGlobals()
	fixedNow := time.Unix(500, 0)
	nowFunc = func() time.Time { return fixedNow }
	generateKeyFn = func() (Key, error) {
		return Key{KeyID: "generated", AlgorithmKey: "RS256"}, nil
	}
	defer func() {
		nowFunc = time.Now
		generateKeyFn = GenerateKey
	}()

	provider := KeyStorageProvider{
		LastRotation:     fixedNow.Add(-time.Hour),
		RotationInterval: time.Minute,
		NumKeys:          2,
		Keys: map[int]Key{
			0: {KeyID: "zero"},
			1: {KeyID: "one"},
			2: {},
		},
	}

	rotated := provider.Rotate(false)
	require.True(t, rotated)
	require.Equal(t, fixedNow, provider.LastRotation)
	require.Equal(t, "one", provider.Keys[0].KeyID)
	require.Equal(t, "generated", provider.Keys[1].KeyID)
}

func TestKeyStorageProviderNeedRotate(t *testing.T) {
	resetGlobals()
	baseTime := time.Unix(10, 0)
	nowFunc = func() time.Time { return baseTime }
	defer func() { nowFunc = time.Now }()

	provider := KeyStorageProvider{
		LastRotation:     baseTime.Add(-2 * time.Minute),
		RotationInterval: time.Minute,
	}

	require.True(t, provider.needRotate(false))
	require.True(t, provider.needRotate(true))
}

func TestKeyStorageProviderSignSuccess(t *testing.T) {
	resetGlobals()
	privKey := generateRSAKey(t, 2048)
	provider := KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "kid",
				PrivateKey:   privKey,
				AlgorithmKey: "RS256",
			},
		},
	}

	token, err := provider.Sign(jwt.MapClaims{"sub": "user"})
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestKeyStorageProviderSignError(t *testing.T) {
	resetGlobals()
	privKey := generateRSAKey(t, 1024)
	provider := KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "kid",
				PrivateKey:   privKey,
				AlgorithmKey: "ES256",
			},
		},
	}

	_, err := provider.Sign(jwt.MapClaims{"sub": "user"})
	require.Error(t, err)
}

func TestKeyStorageProviderGetJwksNoKeys(t *testing.T) {
	resetGlobals()
	var provider *KeyStorageProvider
	_, err := provider.GetJwks()
	require.EqualError(t, err, "no keys available in keystorage")
}

func TestKeyStorageProviderGetJwksSuccess(t *testing.T) {
	resetGlobals()
	privKey := generateRSAKey(t, 1024)
	provider := &KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "jwks",
				PrivateKey:   privKey,
				AlgorithmKey: "RS256",
			},
		},
	}

	set, err := provider.GetJwks()
	require.NoError(t, err)
	require.Equal(t, 1, set.Len())

	key, ok := set.Key(0)
	require.True(t, ok)

	kidValue, ok := key.Get(jwk.KeyIDKey)
	require.True(t, ok)
	kidStr, ok2 := kidValue.(string)
	require.True(t, ok2)
	require.Equal(t, "jwks", kidStr)
}

func TestGetTokenKeyStorage(t *testing.T) {
	resetGlobals()
	storage := GetTokenKeyStorage()
	require.NotNil(t, storage)
	_, ok := storage.(*KeyStorageProvider)
	require.True(t, ok)
}

func TestInit(t *testing.T) {
	resetGlobals()
	fixedNow := time.Unix(1000, 0)
	nowFunc = func() time.Time { return fixedNow }
	defer func() { nowFunc = time.Now }()

	loadCalled := false
	adapter := &mockStorageAdapter{
		getFn: func() (KeyStorageProvider, error) {
			loadCalled = true
			return KeyStorageProvider{
				LastRotation:     fixedNow,
				RotationInterval: time.Hour,
				NumKeys:          1,
				Keys: map[int]Key{
					1: {KeyID: "kid"},
				},
			}, nil
		},
		setFn: func(*KeyStorageProvider) error {
			return nil
		},
	}

	Init(adapter)

	require.True(t, loadCalled)
	require.NotNil(t, keyStorage.storageAdapter)
	require.Equal(t, "kid", keyStorage.Keys[1].KeyID)
}

func TestRotateKeysRandomError(t *testing.T) {
	resetGlobals()
	nowFunc = func() time.Time { return time.Unix(1000, 0) }
	defer func() { nowFunc = time.Now }()

	keyStorage.LastRotation = time.Unix(0, 0)
	keyStorage.RotationInterval = time.Second

	sleepCalled := false
	sleepFunc = func(d time.Duration) {
		sleepCalled = true
	}
	randomIntFunc = func(_ io.Reader, _ *big.Int) (*big.Int, error) {
		return nil, errors.New("rand failure")
	}
	defer func() {
		sleepFunc = time.Sleep
		randomIntFunc = rand.Int
	}()

	RotateKeys()
	require.False(t, sleepCalled)
}

func TestRotateKeysLoadError(t *testing.T) {
	resetGlobals()
	nowFunc = func() time.Time { return time.Unix(1000, 0) }
	defer func() { nowFunc = time.Now }()

	keyStorage.RotationInterval = time.Second
	keyStorage.LastRotation = time.Unix(0, 0)

	sleepFunc = func(time.Duration) {}
	randomIntFunc = func(_ io.Reader, _ *big.Int) (*big.Int, error) {
		return big.NewInt(0), nil
	}
	defer func() {
		sleepFunc = time.Sleep
		randomIntFunc = rand.Int
	}()

	adapter := &mockStorageAdapter{
		getFn: func() (KeyStorageProvider, error) {
			return KeyStorageProvider{}, errors.New("load error")
		},
		setFn: func(*KeyStorageProvider) error { return nil },
	}
	keyStorage.storageAdapter = adapter

	RotateKeys()
}

func TestRotateKeysSaveError(t *testing.T) {
	resetGlobals()
	fixedNow := time.Unix(5000, 0)
	nowFunc = func() time.Time { return fixedNow }
	generateKeyFn = func() (Key, error) {
		return Key{KeyID: "generated"}, nil
	}
	sleepFunc = func(time.Duration) {}
	randomIntFunc = func(_ io.Reader, _ *big.Int) (*big.Int, error) {
		return big.NewInt(0), nil
	}
	defer func() {
		nowFunc = time.Now
		generateKeyFn = GenerateKey
		sleepFunc = time.Sleep
		randomIntFunc = rand.Int
	}()

	adapter := &mockStorageAdapter{
		getFn: func() (KeyStorageProvider, error) {
			return KeyStorageProvider{
				LastRotation:     fixedNow.Add(-time.Hour),
				RotationInterval: time.Minute,
				NumKeys:          1,
				Keys: map[int]Key{
					0: {KeyID: "zero"},
					1: {},
				},
			}, nil
		},
		setFn: func(*KeyStorageProvider) error {
			return errors.New("save error")
		},
	}
	keyStorage.storageAdapter = adapter

	RotateKeys()
}

func TestRotateKeysSuccess(t *testing.T) {
	resetGlobals()
	fixedNow := time.Unix(7000, 0)
	nowFunc = func() time.Time { return fixedNow }
	generateKeyFn = func() (Key, error) {
		return Key{KeyID: "generated", AlgorithmKey: "RS256"}, nil
	}
	sleepFunc = func(time.Duration) {}
	randomIntFunc = func(_ io.Reader, _ *big.Int) (*big.Int, error) {
		return big.NewInt(0), nil
	}
	defer func() {
		nowFunc = time.Now
		generateKeyFn = GenerateKey
		sleepFunc = time.Sleep
		randomIntFunc = rand.Int
	}()

	adapter := &mockStorageAdapter{
		getFn: func() (KeyStorageProvider, error) {
			return KeyStorageProvider{
				LastRotation:     fixedNow.Add(-time.Hour),
				RotationInterval: time.Minute,
				NumKeys:          1,
				Keys: map[int]Key{
					0: {KeyID: "zero"},
					1: {},
				},
			}, nil
		},
		setFn: func(ks *KeyStorageProvider) error {
			require.Equal(t, "generated", ks.Keys[0].KeyID)
			return nil
		},
	}
	keyStorage.storageAdapter = adapter

	RotateKeys()
	require.Equal(t, fixedNow, keyStorage.LastRotation)
	require.Equal(t, "generated", keyStorage.Keys[0].KeyID)
}

func TestGenerateKey(t *testing.T) {
	resetGlobals()
	key, err := GenerateKey()
	require.NoError(t, err)
	require.Equal(t, "RS512", key.AlgorithmKey)
	require.NotEmpty(t, key.KeyID)
	require.NotNil(t, key.PrivateKey)
}
