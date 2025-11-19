package tokenstoragehelper

import (
	"crypto/rsa"
	"errors"
	"time"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

// TokenKeyStorage defines the interface for managing cryptographic keys used in token signing.
// Implementations of this interface handle key generation, rotation, storage, and JWT operations.
type TokenKeyStorage interface {
	// Load retrieves the key storage state from the underlying storage adapter
	Load() error
	// Save persists the current key storage state to the underlying storage adapter
	Save() error
	// Rotate rotates the keys if the rotation interval has elapsed or if force is true.
	// Returns true if rotation occurred, false otherwise.
	Rotate(force bool) bool
	// Sign creates and signs a JWT token with the provided claims using the current key
	Sign(claims jwt.MapClaims) (string, error)
	// GetJwks returns a JSON Web Key Set containing the public keys for token verification
	GetJwks() (jwk.Set, error)
}

// cast KeystorageProvider to TokenKeyStorage to test interface implementation
var _ TokenKeyStorage = &KeyStorageProvider{}

// StorageAdapter defines the interface for persisting and retrieving key storage.
// Implementations can use various backends such as Vault, databases, or file systems.
type StorageAdapter interface {
	// Set persists the key storage provider to the storage backend
	Set(*KeyStorageProvider) error
	// Get retrieves the key storage provider from the storage backend
	Get() (KeyStorageProvider, error)
}

// KeyStorageProvider manages the storage and rotation of cryptographic keys.
// It maintains multiple keys to support gradual rotation and uses a storage adapter
// for persistence.
type KeyStorageProvider struct {
	storageAdapter   StorageAdapter
	LastRotation     time.Time     `json:"last_rotation"`     // Timestamp of the last key rotation
	RotationInterval time.Duration `json:"rotation_interval"` // Duration between automatic rotations
	NumKeys          int           `json:"num_keys"`          // Total number of keys to maintain
	Keys             map[int]Key   `json:"keys"`              // Map of key index to Key
}

// Key represents an RSA keypair used for JWT signing.
type Key struct {
	KeyID        string          `json:"key_id"`        // Unique identifier for the key (SHA256 thumbprint)
	PrivateKey   *rsa.PrivateKey `json:"private_key"`   // RSA private key for signing
	AlgorithmKey string          `json:"algorithm_key"` // JWT signing algorithm (e.g., "RS512")
}

// NewTokenKeyStorage creates a new instance of the keystorage provider with the provided storage adapter.
// It initializes the key storage by loading existing keys from the adapter and performing an initial
// rotation check. This function is an alternative to using the singleton pattern via Init().
//
// Returns an error if the storage adapter fails to load existing keys.
func NewTokenKeyStorage(storageAdapter StorageAdapter) (TokenKeyStorage, error) {
	tokenStorage := &KeyStorageProvider{}
	tokenStorage.storageAdapter = storageAdapter
	err := tokenStorage.Load()
	if err != nil {
		rlog.Error("could not load keystorage from vault", err)
		return nil, err
	}
	tokenStorage.Rotate(false)
	return tokenStorage, nil
}

// GetTokenKeyStorage returns the singleton instance of the keystorage provider.
// This should only be called after Init() has been invoked to initialize the singleton.
//
// Example:
//
//	tokenStorage := tokenstoragehelper.GetTokenKeyStorage()
//	token, err := tokenStorage.Sign(claims)
func GetTokenKeyStorage() TokenKeyStorage {
	return &keyStorage
}

// getCurrentKey returns the current active key used for signing new tokens.
// The current key is always at index 1 in the Keys map.
func (k *KeyStorageProvider) getCurrentKey() Key {
	return k.Keys[1]
}

// Save persists the current key storage state to the underlying storage adapter.
// Returns an error if no storage adapter is configured or if the save operation fails.
func (k *KeyStorageProvider) Save() error {
	if k.storageAdapter != nil {
		return k.storageAdapter.Set(k)
	}
	return errors.New("no storage provider set")
}

// Load retrieves the key storage state from the underlying storage adapter.
// It updates the current instance with the loaded configuration including keys,
// rotation settings, and timestamps.
//
// Returns an error if no storage adapter is configured or if the load operation fails.
func (k *KeyStorageProvider) Load() error {
	if k.storageAdapter != nil {
		loaded, err := k.storageAdapter.Get()
		if err != nil {
			return err
		}
		k.LastRotation = loaded.LastRotation
		k.RotationInterval = loaded.RotationInterval
		k.NumKeys = loaded.NumKeys
		k.Keys = loaded.Keys
		return nil
	}
	return errors.New("no storage provider set")
}

// Rotate performs key rotation if needed based on the rotation interval or if forced.
// During rotation, keys are shifted down by one position and a new key is generated
// for the highest position. This ensures older keys remain available for verifying
// existing tokens while new tokens are signed with the fresh key.
//
// The force parameter can be set to true to bypass the interval check and force rotation.
// Returns true if rotation occurred, false otherwise.
func (k *KeyStorageProvider) Rotate(force bool) bool {
	if k.needRotate(force) {
		for i := 0; i < k.NumKeys; i++ {
			k.Keys[i] = k.Keys[i+1]
			if k.Keys[i].KeyID == "" {
				rlog.Info("generating new key for position", rlog.Int("position", i))
				newKey, err := GenerateKey()
				if err != nil {
					rlog.Error("could not generate new key", err)
				}
				k.Keys[i] = newKey
			}
		}
		k.LastRotation = time.Now()
		return true
	}
	return false
}

// needRotate determines whether key rotation is needed based on the rotation interval
// or the force parameter. Returns true if the current time exceeds the last rotation
// time plus the rotation interval, or if force is true.
func (k *KeyStorageProvider) needRotate(force bool) bool {
	return time.Now().Unix() > k.LastRotation.Add(k.RotationInterval).Unix() || force
}

// Sign creates and signs a JWT token with the provided claims using the current active key.
// The signed token includes a "kid" (key ID) header to identify which key was used for signing.
//
// Example:
//
//	claims := jwt.MapClaims{
//		"sub": "user123",
//		"exp": time.Now().Add(time.Hour).Unix(),
//		"iat": time.Now().Unix(),
//	}
//	token, err := keyStorage.Sign(claims)
//
// Returns the signed JWT token as a string, or an error if signing fails.
func (k *KeyStorageProvider) Sign(claims jwt.MapClaims) (string, error) {
	currentKey := k.getCurrentKey()

	token := jwt.NewWithClaims(jwt.GetSigningMethod(currentKey.AlgorithmKey), claims)
	token.Header["kid"] = currentKey.KeyID
	signedString, err := token.SignedString(currentKey.PrivateKey)
	if err != nil {
		return "", err
	}
	return signedString, nil

}

// GetJwks returns a JSON Web Key Set (JWKS) containing all public keys from the key storage.
// The JWKS can be published at a well-known endpoint (e.g., /.well-known/jwks.json) to allow
// token verifiers to retrieve public keys for signature validation.
//
// Each key in the set includes:
//   - kid: Key ID for matching with JWT header
//   - alg: Algorithm used for signing
//   - use: Set to "sig" for signature verification
//   - Public key material (n, e for RSA)
//
// Returns an error if no keys are available or if key conversion fails.
func (k *KeyStorageProvider) GetJwks() (jwk.Set, error) {
	if k == nil || len(k.Keys) == 0 {
		return nil, errors.New("no keys available in keystorage")
	}
	set := jwk.NewSet()
	for _, data := range k.Keys {
		pubKey := data.PrivateKey.Public().(*rsa.PublicKey)
		jwkKey, err := jwk.FromRaw(pubKey)
		if err != nil {
			return nil, err
		}
		if err := jwkKey.Set(jwk.KeyIDKey, data.KeyID); err != nil {
			return nil, err
		}
		if err := jwkKey.Set(jwk.AlgorithmKey, data.AlgorithmKey); err != nil {
			return nil, err
		}
		if err := jwkKey.Set(jwk.KeyUsageKey, "sig"); err != nil {
			return nil, err
		}

		if err := set.AddKey(jwkKey); err != nil {
			return nil, err
		}
	}

	return set, nil
}
