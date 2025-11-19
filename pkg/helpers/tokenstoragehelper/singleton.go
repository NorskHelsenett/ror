package tokenstoragehelper

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"math/big"
	"time"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

var (
	// keyStorage is the singleton instance of the key storage provider
	keyStorage    KeyStorageProvider
	randomIntFunc = rand.Int
	sleepFunc     = time.Sleep
)

// RotateKeys performs a key rotation operation on the singleton key storage instance.
// It implements a distributed lock mechanism using random sleep intervals to prevent
// concurrent rotation attempts in multi-instance deployments.
//
// The rotation process:
//  1. Checks if rotation is needed based on the rotation interval
//  2. Sleeps for a random duration (0-5 seconds) to avoid race conditions
//  3. Reloads the key storage from persistent storage
//  4. Performs the rotation if still needed
//  5. Saves the updated keys back to storage
//
// This function is designed to be called periodically (e.g., via a cron job or ticker)
// and is safe to call from multiple instances of an application.
func RotateKeys() {
	if keyStorage.needRotate(false) {

		randomInterval, err := randomIntFunc(rand.Reader, big.NewInt(5000))
		if err != nil {
			rlog.Error("could not generate random interval for key rotation", err)
			return
		}
		sleepFunc(time.Duration(time.Duration(randomInterval.Int64()) * time.Millisecond))
		err = keyStorage.Load()
		if err != nil {
			rlog.Error("could not load keystorage from vault", err)
			return
		}
		rotated := keyStorage.Rotate(true)
		if rotated {
			err := keyStorage.Save()
			if err != nil {
				rlog.Error("could not save keystorage to vault", err)
			}
		}
		rlog.Info("Key rotation completed")
	}
}

// GenerateKey creates a new RSA keypair for JWT signing.
// It generates a 4096-bit RSA key and calculates a SHA256 thumbprint to use as the key ID.
//
// The generated key uses the RS512 algorithm (RSA signature with SHA-512) which provides
// strong security guarantees for JWT token signing.
//
// Returns a Key struct containing:
//   - KeyID: Hex-encoded SHA256 thumbprint of the public key
//   - PrivateKey: The generated RSA private key
//   - AlgorithmKey: "RS512" signing algorithm identifier
//
// Returns an error if key generation or thumbprint calculation fails.
func GenerateKey() (Key, error) {
	newPrivatekey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return Key{}, err
	}

	thumprint, err := jwk.FromRaw(newPrivatekey.PublicKey)
	if err != nil {
		return Key{}, err
	}
	keyid, err := thumprint.Thumbprint(crypto.SHA256)
	if err != nil {
		return Key{}, err
	}
	key := Key{
		KeyID:        fmt.Sprintf("%x", keyid),
		PrivateKey:   newPrivatekey,
		AlgorithmKey: "RS512",
	}
	return key, nil
}

// Init initializes the singleton token service with the provided storage provider.
// This function should be called once at application startup before using any other
// functions in this package.
//
// The initialization process:
//  1. Sets the storage adapter for the singleton instance
//  2. Loads existing keys from storage (or initializes if none exist)
//  3. Performs an initial rotation check to ensure keys are current
//
// Example:
//
//	adapter := vaulttokenadapter.NewVaultStorageAdapter(vaultClient, "secret/jwt-keys")
//	tokenstoragehelper.Init(adapter)
//
// Parameters:
//   - storageProvider: Implementation of StorageAdapter for persisting keys
func Init(storageProvider StorageAdapter) {
	keyStorage.storageAdapter = storageProvider
	err := keyStorage.Load()
	if err != nil {
		rlog.Error("could not load keystorage from vault", err)
	}
	keyStorage.Rotate(false)
}
