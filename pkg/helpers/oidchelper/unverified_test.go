package oidchelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractUnverifiedClaims_ValidToken(t *testing.T) {
	// JWT payload: {"iss":"https://issuer.example.com","aud":"my-client","sub":"user@test.com"}
	token := "eyJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJodHRwczovL2lzc3Vlci5leGFtcGxlLmNvbSIsImF1ZCI6Im15LWNsaWVudCIsInN1YiI6InVzZXJAdGVzdC5jb20ifQ.fakesig"

	claims, err := extractUnverifiedClaims(token)
	require.NoError(t, err)
	assert.Equal(t, "https://issuer.example.com", claims.Issuer)
	assert.Equal(t, audience{"my-client"}, claims.Audience)
}

func TestExtractUnverifiedClaims_InvalidFormat(t *testing.T) {
	_, err := extractUnverifiedClaims("not-a-jwt")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid token format")
}

func TestExtractUnverifiedClaims_MissingIssuer(t *testing.T) {
	// payload: {"aud":"my-client"}
	token := "eyJhbGciOiJSUzI1NiJ9.eyJhdWQiOiJteS1jbGllbnQifQ.fakesig"
	_, err := extractUnverifiedClaims(token)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "issuer claim is missing")
}

func TestMatchAudience(t *testing.T) {
	u := &unverifiedToken{
		Audience: audience{"client-a", "client-b"},
	}

	matched, ok := u.matchAudience("client-b", "client-c")
	assert.True(t, ok)
	assert.Equal(t, "client-b", matched)

	_, ok = u.matchAudience("client-c")
	assert.False(t, ok)
}

func TestMatchAudience_NoClientIDs(t *testing.T) {
	u := &unverifiedToken{
		Audience: audience{"client-a"},
	}
	_, ok := u.matchAudience()
	assert.False(t, ok)
}

func TestExtractGroups(t *testing.T) {
	groups, err := ExtractGroups("alice@example.com", []string{"admins", "devs"})
	require.NoError(t, err)
	assert.Equal(t, []string{"admins@example.com", "devs@example.com"}, groups)
}

func TestExtractGroups_EmptyEmail(t *testing.T) {
	_, err := ExtractGroups("", []string{"admins"})
	assert.Error(t, err)
}

func TestExtractGroups_InvalidEmail(t *testing.T) {
	_, err := ExtractGroups("noatsign", []string{"admins"})
	assert.Error(t, err)
}

func TestExtractGroups_EmptyGroups(t *testing.T) {
	groups, err := ExtractGroups("alice@example.com", []string{})
	require.NoError(t, err)
	assert.Empty(t, groups)
}
