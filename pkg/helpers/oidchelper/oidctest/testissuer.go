package oidctest

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/oidchelper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

// TestIssuer is an in-memory OIDC issuer for testing.
type TestIssuer struct {
	IssuerURL  string
	PrivateKey *rsa.PrivateKey
	KeyID      string
	server     *httptest.Server
}

// NewTestIssuer creates a test issuer backed by an httptest.Server.
func NewTestIssuer() (*TestIssuer, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("could not generate RSA key: %w", err)
	}

	issuer := &TestIssuer{
		PrivateKey: privateKey,
		KeyID:      "test-key-1",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/.well-known/openid-configuration", issuer.handleDiscovery)
	mux.HandleFunc("/jwks", issuer.handleJWKS)

	issuer.server = httptest.NewServer(mux)
	issuer.IssuerURL = issuer.server.URL

	return issuer, nil
}

// Close shuts down the test OIDC server.
func (ti *TestIssuer) Close() {
	if ti.server != nil {
		ti.server.Close()
	}
}

// SignToken creates a signed JWT with the given claims.
func (ti *TestIssuer) SignToken(claims oidchelper.TokenClaims) (string, error) {
	mapClaims := jwt.MapClaims{
		"sub":            claims.Subject,
		"iss":            ti.IssuerURL,
		"email":          claims.Email,
		"email_verified": claims.EmailVerified,
		"name":           claims.Name,
		"iat":            time.Now().Unix(),
		"nbf":            time.Now().Add(-1 * time.Minute).Unix(),
	}

	if len(claims.Groups) > 0 {
		mapClaims["groups"] = claims.Groups
	}
	if !claims.ExpirationTime.IsZero() {
		mapClaims["exp"] = claims.ExpirationTime.Unix()
	} else {
		mapClaims["exp"] = time.Now().Add(1 * time.Hour).Unix()
	}
	if claims.Audience != "" {
		mapClaims["aud"] = claims.Audience
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, mapClaims)
	token.Header["kid"] = ti.KeyID

	return token.SignedString(ti.PrivateKey)
}

// Config returns an IssuerConfig for use with the validator.
func (ti *TestIssuer) Config(clientIDs ...string) oidchelper.IssuerConfig {
	return oidchelper.IssuerConfig{
		IssuerURL:  ti.IssuerURL,
		ClientIDs:  clientIDs,
		SkipVerify: true,
	}
}

func (ti *TestIssuer) handleDiscovery(w http.ResponseWriter, _ *http.Request) {
	doc := map[string]interface{}{
		"issuer":                                ti.IssuerURL,
		"jwks_uri":                              ti.IssuerURL + "/jwks",
		"subject_types_supported":               []string{"public"},
		"id_token_signing_alg_values_supported": []string{"RS256"},
		"response_types_supported":              []string{"id_token"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doc)
}

func (ti *TestIssuer) handleJWKS(w http.ResponseWriter, _ *http.Request) {
	pubKey := ti.PrivateKey.Public().(*rsa.PublicKey)
	jwkKey, err := jwk.FromRaw(pubKey)
	if err != nil {
		http.Error(w, "could not create JWK", http.StatusInternalServerError)
		return
	}
	_ = jwkKey.Set(jwk.KeyIDKey, ti.KeyID)
	_ = jwkKey.Set(jwk.AlgorithmKey, "RS256")
	_ = jwkKey.Set(jwk.KeyUsageKey, "sig")

	set := jwk.NewSet()
	_ = set.AddKey(jwkKey)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(set)
}
