package oidchelper

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/config/rorconfig"
)

// LoadFromEnv loads OIDC issuer configurations from environment variables.
// It first checks for OIDC_ISSUERS (JSON array), then falls back to the
// legacy single-issuer variables OIDC_PROVIDER + OIDC_CLIENT_ID.
func LoadFromEnv() ([]IssuerConfig, error) {
	issuersJSON := rorconfig.GetString(rorconfig.OIDC_ISSUERS)
	if issuersJSON != "" {
		configs, err := LoadFromJSON([]byte(issuersJSON))
		if err != nil {
			return nil, fmt.Errorf("could not parse OIDC_ISSUERS: %w", err)
		}
		return configs, nil
	}

	issuerURL := rorconfig.GetString(rorconfig.OIDC_PROVIDER)
	if issuerURL == "" {
		return nil, fmt.Errorf("no OIDC configuration found: set OIDC_ISSUERS or OIDC_PROVIDER")
	}

	var clientIDs []string
	if cid := rorconfig.GetString(rorconfig.OIDC_CLIENT_ID); cid != "" {
		clientIDs = append(clientIDs, cid)
	}
	if dcid := rorconfig.GetString(rorconfig.OIDC_DEVICE_CLIENT_ID); dcid != "" {
		clientIDs = append(clientIDs, dcid)
	}
	if len(clientIDs) == 0 {
		return nil, fmt.Errorf("no OIDC client IDs configured")
	}

	skipVerify := rorconfig.GetBool(rorconfig.OIDC_SKIP_ISSUER_VERIFY)

	return []IssuerConfig{
		{
			IssuerURL:  issuerURL,
			ClientIDs:  clientIDs,
			SkipVerify: skipVerify,
		},
	}, nil
}

// LoadFromJSON parses issuer configurations from a JSON byte slice.
// Expects a JSON array of IssuerConfig objects.
func LoadFromJSON(data []byte) ([]IssuerConfig, error) {
	trimmed := strings.TrimSpace(string(data))
	if trimmed == "" {
		return nil, fmt.Errorf("empty configuration data")
	}

	if !strings.HasPrefix(trimmed, "[") {
		return nil, fmt.Errorf("OIDC_ISSUERS must be a JSON array of issuer configurations")
	}

	var configs []IssuerConfig
	if err := json.Unmarshal([]byte(trimmed), &configs); err != nil {
		return nil, fmt.Errorf("could not parse JSON array: %w", err)
	}
	for i, cfg := range configs {
		if cfg.IssuerURL == "" {
			return nil, fmt.Errorf("issuer at index %d has empty URL", i)
		}
		if len(cfg.ClientIDs) == 0 {
			return nil, fmt.Errorf("issuer %s has no client IDs", cfg.IssuerURL)
		}
	}
	return configs, nil
}
