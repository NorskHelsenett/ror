package oidchelper

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// unverifiedToken holds claims extracted from a JWT without verification.
type unverifiedToken struct {
	Issuer   string   `json:"iss"`
	Audience audience `json:"aud"`
}

// audience handles both string and []string JSON representations of the aud claim.
type audience []string

func (a *audience) UnmarshalJSON(data []byte) error {
	var single string
	if err := json.Unmarshal(data, &single); err == nil {
		*a = audience{single}
		return nil
	}
	var multi []string
	if err := json.Unmarshal(data, &multi); err != nil {
		return err
	}
	*a = multi
	return nil
}

// matchAudience returns the first matching audience from the provided client IDs.
func (u *unverifiedToken) matchAudience(clientIDs ...string) (string, bool) {
	for _, cid := range clientIDs {
		for _, aud := range u.Audience {
			if aud == cid {
				return aud, true
			}
		}
	}
	return "", false
}

// extractUnverifiedClaims decodes the JWT payload without verification
// to extract issuer and audience claims for provider lookup.
func extractUnverifiedClaims(token string) (unverifiedToken, error) {
	var claims unverifiedToken

	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return claims, fmt.Errorf("invalid token format")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return claims, fmt.Errorf("could not decode token payload: %w", err)
	}

	if err := json.Unmarshal(payload, &claims); err != nil {
		return claims, fmt.Errorf("could not parse token claims: %w", err)
	}

	if claims.Issuer == "" {
		return claims, fmt.Errorf("issuer claim is missing in token")
	}

	if len(claims.Audience) == 0 {
		return claims, fmt.Errorf("audience claim is missing in token")
	}

	return claims, nil
}

// ExtractGroups appends the email domain to each group name.
func ExtractGroups(email string, groups []string) ([]string, error) {
	if email == "" {
		return nil, fmt.Errorf("email is empty")
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 || parts[1] == "" {
		return nil, fmt.Errorf("could not extract domain from email")
	}

	domain := parts[1]
	result := make([]string, 0, len(groups))
	for _, g := range groups {
		result = append(result, fmt.Sprintf("%s@%s", g, domain))
	}

	return result, nil
}
