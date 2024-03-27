package httpauthprovider

import (
	"fmt"
	"net/http"
)

type AuthProviderType string

const (
	AuthPoviderTypeAPIKey  AuthProviderType = "APIKEY"
	AuthProviderTypeBearer AuthProviderType = "BEARER"
)

type AuthProvider struct {
	Type   AuthProviderType
	Secret string
}

func NewAuthProvider(providertype AuthProviderType, secret string) *AuthProvider {
	return &AuthProvider{
		Type:   providertype,
		Secret: secret,
	}
}

func (a *AuthProvider) AddAuthHeaders(req *http.Request) {
	switch a.Type {
	case AuthPoviderTypeAPIKey:
		req.Header.Add("X-API-KEY", a.Secret)
	case AuthProviderTypeBearer:
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Secret))
		req.Header.Add("Flow", "device")
	}
}
