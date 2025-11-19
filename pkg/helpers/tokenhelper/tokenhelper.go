package tokenhelper

import "github.com/golang-jwt/jwt/v5"

type TokenVerifier interface {
	// Verify verifies the provided JWT token string and returns the parsed token if valid
	Verify(tokenString string) (*jwt.Token, error)
}
