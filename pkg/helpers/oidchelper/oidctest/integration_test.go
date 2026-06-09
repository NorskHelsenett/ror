package oidctest

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/oidchelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTestIssuer_SignAndValidate(t *testing.T) {
	issuer, err := NewTestIssuer()
	require.NoError(t, err)
	defer issuer.Close()

	validator, err := oidchelper.NewMultiIssuerValidator(issuer.Config(DefaultTestClientID))
	require.NoError(t, err)

	claims := DefaultUserClaims("alice@example.com", "admins", "devs")
	token := MustSignToken(t, issuer, claims)

	validated, err := validator.ValidateToken(context.Background(), token)
	require.NoError(t, err)
	assert.Equal(t, "alice@example.com", validated.Email)
	assert.Equal(t, []string{"admins", "devs"}, validated.Groups)
	assert.Equal(t, issuer.IssuerURL, validated.Issuer)
}

func TestTestIssuer_UnknownIssuer(t *testing.T) {
	issuer1, err := NewTestIssuer()
	require.NoError(t, err)
	defer issuer1.Close()

	issuer2, err := NewTestIssuer()
	require.NoError(t, err)
	defer issuer2.Close()

	validator, err := oidchelper.NewMultiIssuerValidator(issuer1.Config(DefaultTestClientID))
	require.NoError(t, err)

	claims := DefaultUserClaims("bob@example.com")
	token := MustSignToken(t, issuer2, claims)

	_, err = validator.ValidateToken(context.Background(), token)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no OIDC provider found")
}

func TestTestIssuer_WrongAudience(t *testing.T) {
	issuer, err := NewTestIssuer()
	require.NoError(t, err)
	defer issuer.Close()

	validator, err := oidchelper.NewMultiIssuerValidator(issuer.Config("other-client"))
	require.NoError(t, err)

	claims := DefaultUserClaims("alice@example.com")
	token := MustSignToken(t, issuer, claims)

	_, err = validator.ValidateToken(context.Background(), token)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "audience does not match")
}

func TestTestIssuer_ExpiredToken(t *testing.T) {
	issuer, err := NewTestIssuer()
	require.NoError(t, err)
	defer issuer.Close()

	validator, err := oidchelper.NewMultiIssuerValidator(issuer.Config(DefaultTestClientID))
	require.NoError(t, err)

	claims := DefaultUserClaims("alice@example.com")
	claims.ExpirationTime = time.Now().Add(-1 * time.Hour)
	token := MustSignToken(t, issuer, claims)

	_, err = validator.ValidateToken(context.Background(), token)
	assert.Error(t, err)
}

func TestMultipleIssuers(t *testing.T) {
	issuer1, err := NewTestIssuer()
	require.NoError(t, err)
	defer issuer1.Close()

	issuer2, err := NewTestIssuer()
	require.NoError(t, err)
	defer issuer2.Close()

	validator, err := oidchelper.NewMultiIssuerValidator(
		issuer1.Config(DefaultTestClientID),
		issuer2.Config(DefaultTestClientID),
	)
	require.NoError(t, err)

	token1 := MustSignToken(t, issuer1, DefaultUserClaims("alice@example.com"))
	validated1, err := validator.ValidateToken(context.Background(), token1)
	require.NoError(t, err)
	assert.Equal(t, "alice@example.com", validated1.Email)
	assert.Equal(t, issuer1.IssuerURL, validated1.Issuer)

	token2 := MustSignToken(t, issuer2, DefaultUserClaims("bob@example.com"))
	validated2, err := validator.ValidateToken(context.Background(), token2)
	require.NoError(t, err)
	assert.Equal(t, "bob@example.com", validated2.Email)
	assert.Equal(t, issuer2.IssuerURL, validated2.Issuer)
}

func TestNewTestManager_SignAndValidate(t *testing.T) {
	manager, issuer, cleanup := NewTestManager(t, "https://test-ror.local")
	defer cleanup()

	claims := DefaultUserClaims("alice@example.com", "admins")
	token := MustSignToken(t, issuer, claims)

	validated, err := manager.ValidateToken(context.Background(), token)
	require.NoError(t, err)
	assert.Equal(t, "alice@example.com", validated.Email)

	signedClaims := oidchelper.TokenClaims{
		Subject:        "alice@example.com",
		Email:          "alice@example.com",
		Name:           "Alice",
		Groups:         []string{"admins"},
		Audience:       DefaultTestClientID,
		ExpirationTime: time.Now().Add(1 * time.Hour),
	}

	signed, err := manager.SignToken(signedClaims)
	require.NoError(t, err)
	assert.NotEmpty(t, signed)

	jwks, err := manager.GetJWKS()
	require.NoError(t, err)
	assert.True(t, jwks.Len() > 0)
}

func TestMemoryStorageAdapter(t *testing.T) {
	adapter, err := NewMemoryStorageAdapter()
	require.NoError(t, err)

	stored, err := adapter.Get()
	require.NoError(t, err)
	assert.Equal(t, 3, stored.NumKeys)
	assert.NotEmpty(t, stored.Keys[1].KeyID)
}

func TestAddAndRemoveIssuer(t *testing.T) {
	issuer, err := NewTestIssuer()
	require.NoError(t, err)
	defer issuer.Close()

	validator, err := oidchelper.NewMultiIssuerValidator()
	require.NoError(t, err)

	err = validator.AddIssuer(issuer.Config(DefaultTestClientID))
	require.NoError(t, err)

	claims := DefaultUserClaims("alice@example.com")
	token := MustSignToken(t, issuer, claims)

	validated, err := validator.ValidateToken(context.Background(), token)
	require.NoError(t, err)
	assert.Equal(t, "alice@example.com", validated.Email)

	validator.RemoveIssuer(issuer.IssuerURL)

	_, err = validator.ValidateToken(context.Background(), token)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no OIDC provider found")
}

func TestDiscoveryURL_DifferentFromIssuerURL(t *testing.T) {
	// Create the real issuer — its IssuerURL is its own server address
	// and tokens carry that address as the "iss" claim.
	issuer, err := NewTestIssuer()
	require.NoError(t, err)
	defer issuer.Close()

	// Stand up a proxy server at a different address that forwards
	// discovery and JWKS requests to the real issuer. This simulates
	// a Docker-internal hostname (e.g. http://dex:6556) that serves
	// the same OIDC endpoints.
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(issuer.IssuerURL + r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()
		for k, vs := range resp.Header {
			for _, v := range vs {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}))
	defer proxy.Close()

	// IssuerURL matches the token's iss claim (the real server).
	// DiscoveryURL points to the proxy (the Docker-internal address).
	cfg := oidchelper.IssuerConfig{
		IssuerURL:    issuer.IssuerURL,
		DiscoveryURL: proxy.URL,
		ClientIDs:    []string{DefaultTestClientID},
		SkipVerify:   true,
	}

	validator, err := oidchelper.NewMultiIssuerValidator(cfg)
	require.NoError(t, err)

	claims := DefaultUserClaims("alice@example.com", "admins")
	token := MustSignToken(t, issuer, claims)

	validated, err := validator.ValidateToken(context.Background(), token)
	require.NoError(t, err)
	assert.Equal(t, "alice@example.com", validated.Email)
	assert.Equal(t, issuer.IssuerURL, validated.Issuer)
	assert.Equal(t, []string{"admins"}, validated.Groups)
}
