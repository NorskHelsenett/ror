// Package tokenstoragehelper provides a secure token storage and key management system
// for JWT token signing and rotation.
//
// # Overview
//
// This package implements a robust key management solution for JWT token signing with
// automatic key rotation capabilities. It maintains multiple RSA keypairs to ensure
// smooth rotation without invalidating existing tokens, and supports pluggable storage
// backends for persistence.
//
// # Architecture
//
// The package consists of several key components:
//
//   - TokenKeyStorage: Interface defining key management operations
//   - KeyStorageProvider: Main implementation managing multiple keys
//   - StorageAdapter: Interface for pluggable persistence backends
//   - VaultStorageAdapter: Vault implementation of StorageAdapter
//
// # Key Rotation Strategy
//
// Keys are rotated based on a configurable time interval. The system maintains multiple
// keys (configured via NumKeys) to ensure tokens signed with older keys remain valid
// during the rotation period. When rotation occurs:
//
//  1. Existing keys shift down by one position
//  2. A new key is generated for the highest position
//  3. The current signing key (index 1) is always the most recent
//  4. Older keys are retained for verification purposes
//
// # Singleton vs Instance Usage
//
// The package supports two usage patterns:
//
// Singleton Pattern (recommended for most applications):
//
//	// Initialize once at startup
//	adapter := vaulttokenadapter.NewVaultStorageAdapter(vaultClient, "secret/jwt-keys")
//	tokenstoragehelper.Init(adapter)
//
//	// Use throughout the application
//	storage := tokenstoragehelper.GetTokenKeyStorage()
//	token, err := storage.Sign(claims)
//
// Instance Pattern (for multiple key stores):
//
//	storage, err := tokenstoragehelper.NewTokenKeyStorage(adapter)
//	token, err := storage.Sign(claims)
//
// # Security Considerations
//
//   - Uses 4096-bit RSA keys for strong security
//   - RS512 algorithm (RSA with SHA-512) for signing
//   - Private keys are stored encrypted in the backend (e.g., Vault)
//   - Key IDs are SHA256 thumbprints of public keys
//   - Supports distributed deployments with coordinated rotation
//
// # JWKS Endpoint
//
// The package supports generating JWKS (JSON Web Key Set) for public key distribution:
//
//	storage := tokenstoragehelper.GetTokenKeyStorage()
//	jwks, err := storage.GetJwks()
//	// Serve jwks at /.well-known/jwks.json
//
// # Example: Complete Setup
//
//	package main
//
//	import (
//		"time"
//		"github.com/NorskHelsenett/ror/pkg/helpers/tokenstoragehelper"
//		"github.com/NorskHelsenett/ror/pkg/helpers/tokenstoragehelper/vaulttokenadapter"
//		"github.com/golang-jwt/jwt/v5"
//	)
//
//	func main() {
//		// Initialize storage adapter
//		vaultClient := initVaultClient() // Your Vault setup
//		adapter := vaulttokenadapter.NewVaultStorageAdapter(
//			vaultClient,
//			"secret/data/jwt-keys",
//		)
//
//		// Initialize token storage
//		tokenstoragehelper.Init(adapter)
//
//		// Schedule periodic rotation (e.g., every hour)
//		go func() {
//			ticker := time.NewTicker(1 * time.Hour)
//			for range ticker.C {
//				tokenstoragehelper.RotateKeys()
//			}
//		}()
//
//		// Sign tokens
//		storage := tokenstoragehelper.GetTokenKeyStorage()
//		claims := jwt.MapClaims{
//			"sub": "user123",
//			"exp": time.Now().Add(24 * time.Hour).Unix(),
//			"iat": time.Now().Unix(),
//		}
//		token, err := storage.Sign(claims)
//		if err != nil {
//			// Handle error
//		}
//
//		// Publish JWKS for verification
//		jwks, err := storage.GetJwks()
//		// Serve at /.well-known/jwks.json
//	}
//
// # Custom Storage Adapters
//
// To implement a custom storage backend, implement the StorageAdapter interface:
//
//	type CustomAdapter struct {
//		// Your fields
//	}
//
//	func (a *CustomAdapter) Set(ks *KeyStorageProvider) error {
//		// Persist ks to your backend
//		return nil
//	}
//
//	func (a *CustomAdapter) Get() (KeyStorageProvider, error) {
//		// Retrieve ks from your backend
//		return KeyStorageProvider{}, nil
//	}
package tokenstoragehelper
