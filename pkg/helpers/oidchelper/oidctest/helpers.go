package oidctest

import (
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/oidchelper"
)

// DefaultTestClientID is the default audience used in test tokens.
const DefaultTestClientID = "test-client"

// NewTestManager creates a Manager wired to an in-memory test issuer.
func NewTestManager(t *testing.T, signerIssuer string) (*oidchelper.Manager, *TestIssuer, func()) {
	t.Helper()

	issuer, err := NewTestIssuer()
	if err != nil {
		t.Fatalf("could not create test issuer: %v", err)
	}

	adapter, err := NewMemoryStorageAdapter()
	if err != nil {
		issuer.Close()
		t.Fatalf("could not create memory storage adapter: %v", err)
	}

	manager, err := oidchelper.NewManager(
		signerIssuer,
		adapter,
		issuer.Config(DefaultTestClientID),
	)
	if err != nil {
		issuer.Close()
		t.Fatalf("could not create manager: %v", err)
	}

	cleanup := func() {
		issuer.Close()
	}

	return manager, issuer, cleanup
}

// MustSignToken signs a token with the test issuer or fails the test.
func MustSignToken(t *testing.T, issuer *TestIssuer, claims oidchelper.TokenClaims) string {
	t.Helper()
	token, err := issuer.SignToken(claims)
	if err != nil {
		t.Fatalf("could not sign token: %v", err)
	}
	return token
}

// DefaultUserClaims creates TokenClaims for a test user.
func DefaultUserClaims(email string, groups ...string) oidchelper.TokenClaims {
	return oidchelper.TokenClaims{
		Subject:        email,
		Email:          email,
		EmailVerified:  true,
		Name:           "Test User",
		Groups:         groups,
		Audience:       DefaultTestClientID,
		ExpirationTime: time.Now().Add(1 * time.Hour),
	}
}

// DefaultServiceClaims creates TokenClaims for a test service.
func DefaultServiceClaims(name string) oidchelper.TokenClaims {
	return oidchelper.TokenClaims{
		Subject:        name,
		Name:           name,
		Audience:       DefaultTestClientID,
		ExpirationTime: time.Now().Add(1 * time.Hour),
	}
}
