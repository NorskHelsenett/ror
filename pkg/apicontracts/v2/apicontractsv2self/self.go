package apicontractsv2self

import (
	identitymodels "ror/internal/identity/models"
	"time"
)

type SelfData struct {
	Auth identitymodels.AuthInfo     `json:"auth,omitempty"`
	Type identitymodels.IdentityType `json:"type"`
	User SelfUser                    `json:"user,omitempty"`
}

type SelfUser struct {
	Name   string   `json:"name,omitempty"`
	Email  string   `json:"email,omitempty"`
	Groups []string `json:"groups,omitempty"`
}

type CreateOrRenewApikeyRequest struct {
	Name string `json:"name" validate:"required"`
	Ttl  int64  `json:"ttl" validate:"required"`
}

type CreateOrRenewApikeyResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}
