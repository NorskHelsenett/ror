package tokenhelper

import (
	"context"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v5"
)

type TokenVerifier interface {
	// Verify verifies the provided JWT token string and returns the parsed token if valid
	Verify(tokenString string) (*jwt.Token, error)
}

func IsTokenValid(ctx context.Context, issuer, clientId, token string) bool {

	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return false
	}
	return IsTokenValidByProvider(ctx, provider, clientId, token)
}

func IsTokenValidByProvider(ctx context.Context, provider *oidc.Provider, clientId, token string) bool {

	tokenVerifier := provider.Verifier(&oidc.Config{
		ClientID: clientId,
	})

	return IsTokenValidByVerifier(ctx, tokenVerifier, token)
}

func IsTokenValidByVerifier(ctx context.Context, verifier *oidc.IDTokenVerifier, token string) bool {
	_, err := verifier.Verify(ctx, token)
	if err != nil {
		return false
	}
	return true
}

// IsTokenExpired checks if the provided JWT token string is expired
// IT IS THE CALLER'S RESPONSIBILITY TO ENSURE THE TOKEN IS VALID BEFORE CALLING THIS FUNCTION
func IsTokenExpired(tokenString string) bool {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return true
	}

	expiry, err := token.Claims.GetExpirationTime()
	if err != nil {
		return true
	}

	return expiry.Before(time.Now())
}

func GetTokenExpiration(tokenString string) time.Time {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return time.Time{}
	}

	expiry, err := token.Claims.GetExpirationTime()
	if err != nil {
		return time.Time{}
	}

	return expiry.Time
}

func GetNameFromToken(tokenString string) string {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "Could not extract name"
	}

	claims := token.Claims.(jwt.MapClaims)

	nameClaim, ok := claims["name"]
	if !ok {
		return "Name not found"
	}

	name, ok := nameClaim.(string)
	if !ok {
		return "Name is not a string"
	}

	return name

}

func GetEmailFromToken(tokenString string) string {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "Could not extract email"
	}

	claims := token.Claims.(jwt.MapClaims)

	emailClaim, ok := claims["email"]
	if !ok {
		return "Email not found"
	}

	email, ok := emailClaim.(string)
	if !ok {
		return "Email is not a string"
	}

	return email
}
