package oidchelper

import (
	"fmt"
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/tokenstoragehelper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

// TokenSignerImpl wraps tokenstoragehelper with a configurable issuer.
type TokenSignerImpl struct {
	storage   tokenstoragehelper.SigningTokenKeyStorage
	issuerURL string
}

// NewTokenSigner creates a signer with a configurable issuer URL.
func NewTokenSigner(storage tokenstoragehelper.SigningTokenKeyStorage, issuerURL string) *TokenSignerImpl {
	return &TokenSignerImpl{
		storage:   storage,
		issuerURL: issuerURL,
	}
}

// SignToken signs a TokenClaims struct into a JWT string.
func (s *TokenSignerImpl) SignToken(claims TokenClaims) (string, error) {
	mapClaims := jwt.MapClaims{
		"sub":   claims.Subject,
		"iss":   s.issuerURL,
		"email": claims.Email,
		"name":  claims.Name,
		"iat":   time.Now().Unix(),
		"nbf":   time.Now().Add(-1 * time.Minute).Unix(),
	}

	if claims.Email != "" {
		mapClaims["email_verified"] = claims.EmailVerified
	}
	if len(claims.Groups) > 0 {
		mapClaims["groups"] = claims.Groups
	}
	if !claims.ExpirationTime.IsZero() {
		mapClaims["exp"] = claims.ExpirationTime.Unix()
	}
	if claims.Audience != "" {
		mapClaims["aud"] = claims.Audience
	}
	if claims.ClusterID != "" {
		mapClaims["clusterID"] = claims.ClusterID
	}
	if claims.ProviderISS != "" {
		mapClaims["providerISS"] = claims.ProviderISS
	}
	if claims.ProviderAudience != "" {
		mapClaims["providerAudience"] = claims.ProviderAudience
	}

	return s.storage.Sign(mapClaims)
}

// SignMapClaims signs raw jwt.MapClaims. The issuer claim is set automatically.
func (s *TokenSignerImpl) SignMapClaims(claims jwt.MapClaims) (string, error) {
	if claims == nil {
		return "", fmt.Errorf("claims cannot be nil")
	}
	claims["iss"] = s.issuerURL
	return s.storage.Sign(claims)
}

// GetJWKS returns the JSON Web Key Set for token verification.
func (s *TokenSignerImpl) GetJWKS() (jwk.Set, error) {
	return s.storage.GetJwks()
}

// GetIssuerURL returns the configured issuer URL.
func (s *TokenSignerImpl) GetIssuerURL() string {
	return s.issuerURL
}
