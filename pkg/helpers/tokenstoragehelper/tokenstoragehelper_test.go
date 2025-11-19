package tokenstoragehelper

import (
	"crypto"
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

type stubJWKKey struct {
	jwk.Key
	failName string
	err      error
}

func (s *stubJWKKey) Set(name string, value interface{}) error {
	if name == s.failName {
		return s.err
	}
	return s.Key.Set(name, value)
}

func resetGlobals() {
	nowFunc = time.Now
	generateKeyFn = GenerateKey
	randomIntFunc = rand.Int
	sleepFunc = time.Sleep
	jwkFromRawFn = jwk.FromRaw
	newJWKSetFn = jwk.NewSet
	addKeyFunc = func(set jwk.Set, key jwk.Key) error {
		return set.AddKey(key)
	}
	rotateErrorHandler = func(string, error) {}
	rsaGenerateKeyFunc = rsa.GenerateKey
	jwkThumbprintFunc = func(key jwk.Key, hash crypto.Hash) ([]byte, error) {
		return key.Thumbprint(hash)
	}
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

func TestKeyStorageProviderRotateGenerateKeyError(t *testing.T) {
	resetGlobals()
	fixedNow := time.Unix(800, 0)
	nowFunc = func() time.Time { return fixedNow }
	generateKeyFn = func() (Key, error) {
		return Key{}, errors.New("generate failure")
	}
	defer func() {
		nowFunc = time.Now
		generateKeyFn = GenerateKey
	}()

	provider := KeyStorageProvider{
		LastRotation:     fixedNow.Add(-time.Hour),
		RotationInterval: time.Minute,
		NumKeys:          1,
		Keys: map[int]Key{
			0: {KeyID: "active"},
			1: {},
		},
	}

	rotated := provider.Rotate(false)
	require.True(t, rotated)
	require.Equal(t, fixedNow, provider.LastRotation)
	require.Empty(t, provider.Keys[0].KeyID)
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

func TestKeyStorageProviderVerifySuccess(t *testing.T) {
	resetGlobals()
	privKey := generateRSAKey(t, 2048)
	provider := &KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "kid",
				PrivateKey:   privKey,
				AlgorithmKey: "RS256",
			},
		},
	}

	tokenString, err := provider.Sign(jwt.MapClaims{"sub": "user"})
	require.NoError(t, err)

	token, err := provider.Verify(tokenString)
	require.NoError(t, err)
	require.NotNil(t, token)

	claims, ok := token.Claims.(jwt.MapClaims)
	require.True(t, ok)
	require.Equal(t, "user", claims["sub"])
}

func TestKeyStorageProviderVerifyMissingKid(t *testing.T) {
	resetGlobals()
	privKey := generateRSAKey(t, 1024)
	provider := &KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "kid",
				PrivateKey:   privKey,
				AlgorithmKey: "RS256",
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{"sub": "user"})
	tokenString, err := token.SignedString(privKey)
	require.NoError(t, err)

	_, err = provider.Verify(tokenString)
	require.ErrorContains(t, err, "token missing kid header")
}

func TestKeyStorageProviderVerifyUnknownKid(t *testing.T) {
	resetGlobals()
	privKey := generateRSAKey(t, 1024)
	provider := &KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "kid",
				PrivateKey:   privKey,
				AlgorithmKey: "RS256",
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{"sub": "user"})
	token.Header["kid"] = "unknown"
	tokenString, err := token.SignedString(privKey)
	require.NoError(t, err)

	_, err = provider.Verify(tokenString)
	require.ErrorContains(t, err, "no matching key for kid unknown")
}

func TestKeyStorageProviderVerifyUnexpectedAlgorithm(t *testing.T) {
	resetGlobals()
	privKey := generateRSAKey(t, 1024)
	provider := &KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "kid",
				PrivateKey:   privKey,
				AlgorithmKey: "RS256",
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{"sub": "user"})
	token.Header["kid"] = "kid"
	tokenString, err := token.SignedString(privKey)
	require.NoError(t, err)

	_, err = provider.Verify(tokenString)
	require.ErrorContains(t, err, "signing method RS512 is invalid")
}

func TestKeyStorageProviderVerifyEmptyToken(t *testing.T) {
	resetGlobals()
	provider := &KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "kid",
				PrivateKey:   generateRSAKey(t, 1024),
				AlgorithmKey: "RS256",
			},
		},
	}

	_, err := provider.Verify("")
	require.EqualError(t, err, "token string is empty")
}

func TestKeyStorageProviderVerifyNilProvider(t *testing.T) {
	resetGlobals()
	var provider *KeyStorageProvider

	_, err := provider.Verify("token")
	require.EqualError(t, err, "keystorage provider is nil")
}

func TestKeyStorageProviderVerifyNoKeys(t *testing.T) {
	resetGlobals()
	provider := &KeyStorageProvider{}

	_, err := provider.Verify("token")
	require.EqualError(t, err, "no keys available in keystorage")
}

func TestKeyStorageProviderVerifyNoKeyIDs(t *testing.T) {
	resetGlobals()
	provider := &KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				PrivateKey:   generateRSAKey(t, 1024),
				AlgorithmKey: "RS256",
			},
		},
	}

	_, err := provider.Verify("token")
	require.EqualError(t, err, "no keys with key ids available in keystorage")
}

func TestKeyStorageProviderVerifyMissingPrivateKey(t *testing.T) {
	resetGlobals()
	signingKey := generateRSAKey(t, 1024)
	provider := &KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "kid",
				AlgorithmKey: "RS256",
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{"sub": "user"})
	token.Header["kid"] = "kid"
	tokenString, err := token.SignedString(signingKey)
	require.NoError(t, err)

	_, err = provider.Verify(tokenString)
	require.ErrorContains(t, err, "key material missing for kid kid")
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

func TestKeyStorageProviderGetJwksFromRawError(t *testing.T) {
	resetGlobals()
	privKey := generateRSAKey(t, 1024)
	jwkFromRawFn = func(interface{}) (jwk.Key, error) {
		return nil, errors.New("from raw failure")
	}
	defer func() { jwkFromRawFn = jwk.FromRaw }()

	provider := &KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "kid",
				PrivateKey:   privKey,
				AlgorithmKey: "RS256",
			},
		},
	}

	_, err := provider.GetJwks()
	require.EqualError(t, err, "from raw failure")
}

func TestKeyStorageProviderGetJwksKeyIDSetError(t *testing.T) {
	resetGlobals()
	privKey := generateRSAKey(t, 1024)
	original := jwkFromRawFn
	jwkFromRawFn = func(v interface{}) (jwk.Key, error) {
		key, err := original(v)
		if err != nil {
			return nil, err
		}
		return &stubJWKKey{Key: key, failName: jwk.KeyIDKey, err: errors.New("key id failure")}, nil
	}
	defer func() { jwkFromRawFn = jwk.FromRaw }()

	provider := &KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "kid",
				PrivateKey:   privKey,
				AlgorithmKey: "RS256",
			},
		},
	}

	_, err := provider.GetJwks()
	require.EqualError(t, err, "key id failure")
}

func TestKeyStorageProviderGetJwksAlgorithmSetError(t *testing.T) {
	resetGlobals()
	privKey := generateRSAKey(t, 1024)
	original := jwkFromRawFn
	jwkFromRawFn = func(v interface{}) (jwk.Key, error) {
		key, err := original(v)
		if err != nil {
			return nil, err
		}
		return &stubJWKKey{Key: key, failName: jwk.AlgorithmKey, err: errors.New("algorithm failure")}, nil
	}
	defer func() { jwkFromRawFn = jwk.FromRaw }()

	provider := &KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "kid",
				PrivateKey:   privKey,
				AlgorithmKey: "RS256",
			},
		},
	}

	_, err := provider.GetJwks()
	require.EqualError(t, err, "algorithm failure")
}

func TestKeyStorageProviderGetJwksUsageSetError(t *testing.T) {
	resetGlobals()
	privKey := generateRSAKey(t, 1024)
	original := jwkFromRawFn
	jwkFromRawFn = func(v interface{}) (jwk.Key, error) {
		key, err := original(v)
		if err != nil {
			return nil, err
		}
		return &stubJWKKey{Key: key, failName: jwk.KeyUsageKey, err: errors.New("usage failure")}, nil
	}
	defer func() { jwkFromRawFn = jwk.FromRaw }()

	provider := &KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "kid",
				PrivateKey:   privKey,
				AlgorithmKey: "RS256",
			},
		},
	}

	_, err := provider.GetJwks()
	require.EqualError(t, err, "usage failure")
}

func TestKeyStorageProviderGetJwksAddKeyError(t *testing.T) {
	resetGlobals()
	privKey := generateRSAKey(t, 1024)
	addKeyFunc = func(jwk.Set, jwk.Key) error {
		return errors.New("add key failure")
	}
	defer func() {
		addKeyFunc = func(set jwk.Set, key jwk.Key) error {
			return set.AddKey(key)
		}
	}()

	provider := &KeyStorageProvider{
		Keys: map[int]Key{
			1: {
				KeyID:        "kid",
				PrivateKey:   privKey,
				AlgorithmKey: "RS256",
			},
		},
	}

	_, err := provider.GetJwks()
	require.EqualError(t, err, "add key failure")
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
	var contexts []string
	rotateErrorHandler = func(ctx string, err error) {
		contexts = append(contexts, ctx)
	}
	randCalled := false
	randomIntFunc = func(_ io.Reader, _ *big.Int) (*big.Int, error) {
		randCalled = true
		return nil, errors.New("rand failure")
	}
	defer func() {
		sleepFunc = time.Sleep
		randomIntFunc = rand.Int
	}()

	require.True(t, keyStorage.needRotate(false))
	RotateKeys()
	require.False(t, sleepCalled)
	require.True(t, randCalled)
	require.Equal(t, []string{"random_interval"}, contexts)
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

	loadCalled := false
	var contexts []string
	rotateErrorHandler = func(ctx string, err error) {
		contexts = append(contexts, ctx)
	}
	adapter := &mockStorageAdapter{
		getFn: func() (KeyStorageProvider, error) {
			loadCalled = true
			return KeyStorageProvider{}, errors.New("load error")
		},
		setFn: func(*KeyStorageProvider) error { return nil },
	}
	keyStorage.storageAdapter = adapter

	require.True(t, keyStorage.needRotate(false))
	RotateKeys()
	require.True(t, loadCalled)
	require.Equal(t, []string{"load"}, contexts)
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

	setCalled := false
	var contexts []string
	rotateErrorHandler = func(ctx string, err error) {
		contexts = append(contexts, ctx)
	}
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
			setCalled = true
			return errors.New("save error")
		},
	}
	keyStorage.storageAdapter = adapter

	require.True(t, keyStorage.needRotate(false))
	RotateKeys()
	require.True(t, setCalled)
	require.Equal(t, []string{"save"}, contexts)
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

	setCalled := false
	var contexts []string
	rotateErrorHandler = func(ctx string, _ error) {
		contexts = append(contexts, ctx)
	}
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
			setCalled = true
			require.Equal(t, "generated", ks.Keys[0].KeyID)
			return nil
		},
	}
	keyStorage.storageAdapter = adapter

	require.True(t, keyStorage.needRotate(false))
	RotateKeys()
	require.Equal(t, fixedNow, keyStorage.LastRotation)
	require.Equal(t, "generated", keyStorage.Keys[0].KeyID)
	require.True(t, setCalled)
	require.Empty(t, contexts)
}

func TestGenerateKey(t *testing.T) {
	resetGlobals()
	key, err := GenerateKey()
	require.NoError(t, err)
	require.Equal(t, "RS512", key.AlgorithmKey)
	require.NotEmpty(t, key.KeyID)
	require.NotNil(t, key.PrivateKey)
}

func TestGenerateKeyRSAGenerateError(t *testing.T) {
	resetGlobals()
	rsaGenerateKeyFunc = func(io.Reader, int) (*rsa.PrivateKey, error) {
		return nil, errors.New("rsa failure")
	}
	defer func() { rsaGenerateKeyFunc = rsa.GenerateKey }()

	_, err := GenerateKey()
	require.EqualError(t, err, "rsa failure")
}

func TestGenerateKeyFromRawError(t *testing.T) {
	resetGlobals()
	rsaGenerateKeyFunc = func(reader io.Reader, bits int) (*rsa.PrivateKey, error) {
		return generateRSAKey(t, 1024), nil
	}
	original := jwkFromRawFn
	jwkFromRawFn = func(interface{}) (jwk.Key, error) {
		return nil, errors.New("from raw failure")
	}
	defer func() {
		rsaGenerateKeyFunc = rsa.GenerateKey
		jwkFromRawFn = original
	}()

	_, err := GenerateKey()
	require.EqualError(t, err, "from raw failure")
}

func TestGenerateKeyThumbprintError(t *testing.T) {
	resetGlobals()
	rsaGenerateKeyFunc = func(reader io.Reader, bits int) (*rsa.PrivateKey, error) {
		return generateRSAKey(t, 1024), nil
	}
	jwkThumbprintFunc = func(jwk.Key, crypto.Hash) ([]byte, error) {
		return nil, errors.New("thumbprint failure")
	}
	defer func() {
		rsaGenerateKeyFunc = rsa.GenerateKey
		jwkThumbprintFunc = func(key jwk.Key, hash crypto.Hash) ([]byte, error) {
			return key.Thumbprint(hash)
		}
	}()

	_, err := GenerateKey()
	require.EqualError(t, err, "thumbprint failure")
}

func TestDefaultAddKeyFunc(t *testing.T) {
	resetGlobals()
	privKey := generateRSAKey(t, 1024)
	set := newJWKSetFn()
	jwkKey, err := jwkFromRawFn(&privKey.PublicKey)
	require.NoError(t, err)

	err = addKeyFunc(set, jwkKey)
	require.NoError(t, err)
	require.Equal(t, 1, set.Len())
}
