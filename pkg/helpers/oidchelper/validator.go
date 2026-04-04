package oidchelper

import (
	"context"
	"fmt"
	"sync"
	"time"

	gooidc "github.com/coreos/go-oidc/v3/oidc"
)

type issuerEntry struct {
	config   IssuerConfig
	provider *gooidc.Provider
}

// MultiIssuerValidator validates tokens from multiple OIDC issuers.
type MultiIssuerValidator struct {
	mu      sync.RWMutex
	issuers map[string]*issuerEntry
}

// NewMultiIssuerValidator creates a validator that supports multiple OIDC issuers.
func NewMultiIssuerValidator(configs ...IssuerConfig) (*MultiIssuerValidator, error) {
	v := &MultiIssuerValidator{
		issuers: make(map[string]*issuerEntry),
	}
	for _, cfg := range configs {
		if err := v.AddIssuer(cfg); err != nil {
			return nil, fmt.Errorf("failed to add issuer %s: %w", cfg.IssuerURL, err)
		}
	}
	return v, nil
}

// AddIssuer registers a new OIDC issuer for token validation.
func (v *MultiIssuerValidator) AddIssuer(cfg IssuerConfig) error {
	if cfg.IssuerURL == "" {
		return fmt.Errorf("issuer URL is empty")
	}
	if len(cfg.ClientIDs) == 0 {
		return fmt.Errorf("no client IDs configured for issuer %s", cfg.IssuerURL)
	}

	ctx := context.Background()
	var provider *gooidc.Provider
	var err error

	if cfg.SkipVerify {
		insecureCtx := gooidc.InsecureIssuerURLContext(ctx, cfg.IssuerURL)
		provider, err = gooidc.NewProvider(insecureCtx, cfg.IssuerURL)
	} else {
		provider, err = gooidc.NewProvider(ctx, cfg.IssuerURL)
	}
	if err != nil {
		return fmt.Errorf("could not create OIDC provider for %s: %w", cfg.IssuerURL, err)
	}

	v.mu.Lock()
	defer v.mu.Unlock()
	v.issuers[cfg.IssuerURL] = &issuerEntry{
		config:   cfg,
		provider: provider,
	}
	return nil
}

// AddIssuerWithProvider registers a pre-created OIDC provider (useful for testing).
func (v *MultiIssuerValidator) AddIssuerWithProvider(cfg IssuerConfig, provider *gooidc.Provider) {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.issuers == nil {
		v.issuers = make(map[string]*issuerEntry)
	}
	v.issuers[cfg.IssuerURL] = &issuerEntry{
		config:   cfg,
		provider: provider,
	}
}

// RemoveIssuer unregisters an OIDC issuer.
func (v *MultiIssuerValidator) RemoveIssuer(issuerURL string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	delete(v.issuers, issuerURL)
}

// ValidateToken validates a JWT token against registered issuers.
func (v *MultiIssuerValidator) ValidateToken(ctx context.Context, tokenString string) (*TokenClaims, error) {
	unverified, err := extractUnverifiedClaims(tokenString)
	if err != nil {
		return nil, fmt.Errorf("could not extract claims: %w", err)
	}

	v.mu.RLock()
	entry, exists := v.issuers[unverified.Issuer]
	v.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("no OIDC provider found for issuer: %s", unverified.Issuer)
	}

	clientID, matched := unverified.matchAudience(entry.config.ClientIDs...)
	if !matched {
		return nil, fmt.Errorf("token audience does not match any configured client IDs for issuer %s", unverified.Issuer)
	}

	verifier := entry.provider.Verifier(&gooidc.Config{
		ClientID:                   clientID,
		SkipIssuerCheck:            entry.config.SkipVerify,
		InsecureSkipSignatureCheck: entry.config.SkipVerify,
	})

	idToken, err := verifier.Verify(ctx, tokenString)
	if err != nil {
		return nil, fmt.Errorf("token verification failed: %w", err)
	}

	var rawClaims struct {
		Email         string   `json:"email"`
		EmailVerified bool     `json:"email_verified"`
		Name          string   `json:"name"`
		Groups        []string `json:"groups"`
		Audience      string   `json:"aud"`
		Issuer        string   `json:"iss"`
		Exp           int64    `json:"exp"`
	}
	if err := idToken.Claims(&rawClaims); err != nil {
		return nil, fmt.Errorf("could not extract claims: %w", err)
	}

	return &TokenClaims{
		Subject:        rawClaims.Email,
		Issuer:         rawClaims.Issuer,
		Email:          rawClaims.Email,
		EmailVerified:  rawClaims.EmailVerified,
		Name:           rawClaims.Name,
		Groups:         rawClaims.Groups,
		Audience:       clientID,
		ExpirationTime: time.Unix(rawClaims.Exp, 0),
	}, nil
}
