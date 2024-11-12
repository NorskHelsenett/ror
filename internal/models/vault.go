package models

type VaultTokenModel struct {
	Auth          AuthModel `json:"auth"`
	LeaseDuration int       `json:"lease_duration"`
	LeaseId       string    `json:"lease_id"`
	Renewable     bool      `json:"renewable"`
	RequestId     string    `json:"request_id"`
	Warnings      []string  `json:"warnings"`
	WrapInfo      any       `json:"wrap_info"`
}

type AuthModel struct {
	Accessor       string        `json:"accessor"`
	Client_Token   string        `json:"client_token"`
	Entity_Id      string        `json:"entity_id"`
	Orphan         bool          `json:"orphan"`
	Lease_Duration int           `json:"lease_duration"`
	Metadata       MetadataModel `json:"metadata"`
	Renewable      bool          `json:"renewable"`
	Policies       []string      `json:"policies"`
	Token_Policies []string      `json:"token_policies"`
	Token_Type     string        `json:"token_type"`
}

type VaultRequestTokenModel struct {
	Jwt  string `json:"jwt"`
	Role string `json:"role"`
}

type VaultLoginModel struct {
	// array of strings - A list of policies for the token. This must be a subset of the policies belonging to the token making the request, unless the calling token is root or contains sudo capabilities to auth/token/create. If not specified, defaults to all the policies of the calling token.
	Policies []string `json:"policies"`
	// string - The TTL period of the token, provided as "1h", where hour is the largest suffix. If not provided, the token is valid for the default lease TTL, or indefinitely if the root policy is used.
	TimeToLeave string `json:"ttl"`

	// default is true
	Renewable              bool   `json:"renewable,omitempty"`
	ExplicitMaxTimeToLeave string `json:"explicit_max_ttl,omitempty"`
	DisplayName            string `json:"display_name,omitempty"`
	Period                 string `json:"period,omitempty"`
	MaxNumberOfUses        int    `json:"num_uses,omitempty"`
}

type VaultRenewModel struct {
	Increment string `json:"increment"`
	Token     string `json:"token,omitempty"`    // The token to renew. Only required if iusing /auth/token/renew endpoint
	Accessor  string `json:"accessor,omitempty"` // The accessor of the token to renew. Only required if using /auth/token/renew-accessor endpoint
}

type MetadataModel struct {
	Role                     string `json:"role"`
	ServiceAccountName       string `json:"service_account_name"`
	ServiceAccountNamespace  string `json:"service_account_namespace"`
	ServiceAccountSecretName string `json:"service_account_secret_name"`
	ServiceAccountUid        string `json:"service_account_uid"`
}
