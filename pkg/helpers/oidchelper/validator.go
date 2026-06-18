package oidchelper

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	gooidc "github.com/coreos/go-oidc/v3/oidc"
)

const (
	// oidcInitialBackoff is the wait time before the first discovery retry.
	oidcInitialBackoff = 1 * time.Second
	// oidcMaxBackoff caps the exponential backoff between discovery retries.
	oidcMaxBackoff = 30 * time.Second
	// oidcDiscoveryTimeout bounds a single OIDC discovery attempt so a slow or
	// unreachable issuer cannot block indefinitely.
	oidcDiscoveryTimeout = 30 * time.Second
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
//
// Called with no configs it returns an empty, ready-to-use validator that
// issuers can be added to later (synchronously via AddIssuer or asynchronously
// via LoadIssuersAsync), which is useful when issuer discovery must not block
// startup.
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

// newOidcProvider performs OIDC discovery for cfg using the supplied context.
// When a separate discovery URL is configured (or verification is skipped) it
// discovers against that URL while keeping cfg.IssuerURL as the expected issuer.
func newOidcProvider(ctx context.Context, cfg IssuerConfig) (*gooidc.Provider, error) {
	discoveryURL := cfg.IssuerURL
	if cfg.DiscoveryURL != "" {
		discoveryURL = cfg.DiscoveryURL
	}

	if cfg.DiscoveryURL != "" || cfg.SkipVerify {
		// Use InsecureIssuerURLContext so that discovery is performed against
		// discoveryURL while tokens are validated against cfg.IssuerURL.
		insecureCtx := gooidc.InsecureIssuerURLContext(ctx, cfg.IssuerURL)
		return gooidc.NewProvider(insecureCtx, discoveryURL)
	}
	return gooidc.NewProvider(ctx, cfg.IssuerURL)
}

// AddIssuer registers a new OIDC issuer for token validation. Discovery is
// performed once with a bounded timeout (oidcDiscoveryTimeout) so it cannot
// block indefinitely. Use LoadIssuersAsync to add issuers with retry in the
// background.
func (v *MultiIssuerValidator) AddIssuer(cfg IssuerConfig) error {
	if cfg.IssuerURL == "" {
		return fmt.Errorf("issuer URL is empty")
	}
	if len(cfg.ClientIDs) == 0 {
		return fmt.Errorf("no client IDs configured for issuer %s", cfg.IssuerURL)
	}

	ctx, cancel := context.WithTimeout(context.Background(), oidcDiscoveryTimeout)
	defer cancel()

	provider, err := newOidcProvider(ctx, cfg)
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

// AddIssuerWithRetry registers an issuer, retrying discovery with exponential
// backoff until it succeeds or ctx is cancelled. Each discovery attempt is
// bounded by oidcDiscoveryTimeout. It is safe to call at runtime to add an
// issuer after startup.
func (v *MultiIssuerValidator) AddIssuerWithRetry(ctx context.Context, cfg IssuerConfig) error {
	if cfg.IssuerURL == "" {
		return fmt.Errorf("issuer URL is empty")
	}
	if len(cfg.ClientIDs) == 0 {
		return fmt.Errorf("no client IDs configured for issuer %s", cfg.IssuerURL)
	}

	backoff := oidcInitialBackoff
	for attempt := 1; ; attempt++ {
		attemptCtx, cancel := context.WithTimeout(ctx, oidcDiscoveryTimeout)
		provider, err := newOidcProvider(attemptCtx, cfg)
		cancel()
		if err == nil {
			v.mu.Lock()
			v.issuers[cfg.IssuerURL] = &issuerEntry{config: cfg, provider: provider}
			v.mu.Unlock()
			if attempt > 1 {
				rlog.Info("discovered OIDC issuer",
					rlog.String("issuer", cfg.IssuerURL),
					rlog.Int("attempts", attempt))
			}
			return nil
		}

		rlog.Error("could not discover OIDC issuer, retrying", err,
			rlog.String("issuer", cfg.IssuerURL),
			rlog.Int("attempt", attempt),
			rlog.String("retryIn", backoff.String()))

		select {
		case <-ctx.Done():
			return fmt.Errorf("giving up adding OIDC issuer %s: %w", cfg.IssuerURL, ctx.Err())
		case <-time.After(backoff):
		}

		backoff *= 2
		if backoff > oidcMaxBackoff {
			backoff = oidcMaxBackoff
		}
	}
}

// LoadIssuersAsync registers each issuer in its own goroutine, retrying
// discovery in the background. A slow or unreachable issuer therefore delays
// only its own registration, not startup or the other issuers. Tokens from an
// issuer can only be validated once it has finished registering.
func (v *MultiIssuerValidator) LoadIssuersAsync(ctx context.Context, configs ...IssuerConfig) {
	for _, cfg := range configs {
		go func(cfg IssuerConfig) {
			if err := v.AddIssuerWithRetry(ctx, cfg); err != nil {
				rlog.Error("failed to load OIDC issuer", err, rlog.String("issuer", cfg.IssuerURL))
				return
			}
			rlog.Info("OIDC issuer loaded", rlog.String("issuer", cfg.IssuerURL))
		}(cfg)
	}
}

// oidcStartupChecker gates /health/ready during the initial OIDC issuer load.
// It reports StatusFail until every expected issuer has loaded or the startup
// deadline passes, whichever comes first, then reports StatusPass permanently.
// It only reflects the initial startup load: issuers added later at runtime do
// not affect it.
type oidcStartupChecker struct {
	mu       sync.Mutex
	expected int
	loaded   int
	deadline time.Time
}

func (c *oidcStartupChecker) markLoaded() {
	c.mu.Lock()
	c.loaded++
	c.mu.Unlock()
}

func (c *oidcStartupChecker) CheckHealth(_ context.Context) []rorhealth.Check {
	c.mu.Lock()
	loaded, expected := c.loaded, c.expected
	c.mu.Unlock()

	if loaded >= expected || time.Now().After(c.deadline) {
		return []rorhealth.Check{{Status: rorhealth.StatusPass}}
	}
	return []rorhealth.Check{{
		Status: rorhealth.StatusFail,
		Output: fmt.Sprintf("Loading OIDC issuers (%d/%d)", loaded, expected),
	}}
}

// LoadIssuersForStartup behaves like LoadIssuersAsync but also registers a
// health check named "oidc" that holds /health/ready unready until all issuers
// have loaded or readyTimeout elapses, whichever comes first. After that point
// the check passes permanently. Use this for the initial startup load only;
// issuers added later via AddIssuerWithRetry do not affect readiness.
func (v *MultiIssuerValidator) LoadIssuersForStartup(ctx context.Context, readyTimeout time.Duration, configs ...IssuerConfig) {
	checker := &oidcStartupChecker{
		expected: len(configs),
		deadline: time.Now().Add(readyTimeout),
	}
	rorhealth.Register(ctx, "oidc", checker)

	for _, cfg := range configs {
		go func(cfg IssuerConfig) {
			if err := v.AddIssuerWithRetry(ctx, cfg); err != nil {
				rlog.Error("failed to load OIDC issuer", err, rlog.String("issuer", cfg.IssuerURL))
				return
			}
			checker.markLoaded()
			rlog.Info("OIDC issuer loaded", rlog.String("issuer", cfg.IssuerURL))
		}(cfg)
	}
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
