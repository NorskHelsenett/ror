package oidchelper

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/helpers/tokenstoragehelper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

// Manager combines token validation and signing into a single entry point.
type Manager struct {
	Validator *MultiIssuerValidator
	Signer    *TokenSignerImpl
}

// Compile-time interface check.
var _ TokenManager = &Manager{}

// NewManager creates a Manager with signing capability and multiple issuer validation.
func NewManager(signerIssuer string, storageAdapter tokenstoragehelper.StorageAdapter, issuers ...IssuerConfig) (*Manager, error) {
	signingStorage, err := tokenstoragehelper.NewSigningTokenKeyStorage(storageAdapter)
	if err != nil {
		return nil, fmt.Errorf("could not initialize signing storage: %w", err)
	}

	validator, err := NewMultiIssuerValidator(issuers...)
	if err != nil {
		return nil, fmt.Errorf("could not initialize validator: %w", err)
	}

	return &Manager{
		Validator: validator,
		Signer:   NewTokenSigner(signingStorage, signerIssuer),
	}, nil
}

// NewManagerWithStorage creates a Manager using pre-initialized signing storage.
func NewManagerWithStorage(signerIssuer string, signingStorage tokenstoragehelper.SigningTokenKeyStorage, issuers ...IssuerConfig) (*Manager, error) {
	validator, err := NewMultiIssuerValidator(issuers...)
	if err != nil {
		return nil, fmt.Errorf("could not initialize validator: %w", err)
	}

	return &Manager{
		Validator: validator,
		Signer:   NewTokenSigner(signingStorage, signerIssuer),
	}, nil
}

// ValidateToken delegates to the multi-issuer validator.
func (m *Manager) ValidateToken(ctx context.Context, tokenString string) (*TokenClaims, error) {
	return m.Validator.ValidateToken(ctx, tokenString)
}

// SignToken delegates to the signer.
func (m *Manager) SignToken(claims TokenClaims) (string, error) {
	return m.Signer.SignToken(claims)
}

// SignMapClaims delegates to the signer.
func (m *Manager) SignMapClaims(claims jwt.MapClaims) (string, error) {
	return m.Signer.SignMapClaims(claims)
}

// GetJWKS delegates to the signer.
func (m *Manager) GetJWKS() (jwk.Set, error) {
	return m.Signer.GetJWKS()
}
