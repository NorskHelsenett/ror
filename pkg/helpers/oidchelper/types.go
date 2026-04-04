package oidchelper

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

// IssuerConfig holds configuration for a single OIDC issuer.
type IssuerConfig struct {
	IssuerURL  string   `json:"issuerUrl"`
	ClientIDs  []string `json:"clientIds"`
	SkipVerify bool     `json:"skipVerify,omitempty"`
}

// TokenClaims represents the claims extracted from a validated token.
type TokenClaims struct {
	Subject          string    `json:"sub"`
	Issuer           string    `json:"iss"`
	Email            string    `json:"email"`
	EmailVerified    bool      `json:"email_verified"`
	Name             string    `json:"name"`
	Groups           []string  `json:"groups"`
	Audience         string    `json:"aud"`
	ExpirationTime   time.Time `json:"exp"`
	IssuedAt         time.Time `json:"iat"`
	ClusterID        string    `json:"clusterID,omitempty"`
	ProviderISS      string    `json:"providerISS,omitempty"`
	ProviderAudience string    `json:"providerAudience,omitempty"`
}

// TokenValidator validates OIDC tokens from one or more issuers.
type TokenValidator interface {
	ValidateToken(ctx context.Context, tokenString string) (*TokenClaims, error)
}

// TokenSigner signs JWTs and exposes a JWKS endpoint.
type TokenSigner interface {
	SignToken(claims TokenClaims) (string, error)
	SignMapClaims(claims jwt.MapClaims) (string, error)
	GetJWKS() (jwk.Set, error)
}

// TokenManager combines validation and signing.
type TokenManager interface {
	TokenValidator
	TokenSigner
}
