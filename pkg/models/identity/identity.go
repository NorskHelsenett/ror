// Package implements models representing identity
package identitymodels

import (
	"errors"
	"time"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclprincipal"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Type used to set identity in context
type ContexIdentityType string

const ContexIdentity ContexIdentityType = "ror-identity"

// Type to hold the identitytype eg. user, cluster,service...
type IdentityType string
type IdentityProvider string

const (
	IdentityTypeUser    IdentityType = "User"
	IdentityTypeCluster IdentityType = "Cluster"
	IdentityTypeService IdentityType = "Service"

	IdentityProviderOidc   IdentityProvider = "OIDC"
	IdentityProviderApiKey IdentityProvider = "APIKEY"
)

// Identity is a representation of the consumers identity kept in the context for authentication
type Identity struct {
	Auth            AuthInfo     `json:"auth"`
	Type            IdentityType `json:"type,omitempty"`
	User            *User        `json:"user,omitempty"`
	token           string
	ClusterIdentity *ServiceIdentity `json:"clusterIdentity,omitempty"`
	ServiceIdentity *ServiceIdentity `json:"serviceIdentity,omitempty"`
}

type AuthInfo struct {
	AuthProvider   IdentityProvider `json:"authProvider,omitempty"`
	AuthProviderID string           `json:"authProviderId,omitempty"`
	ExpirationTime time.Time        `json:"expirationTime"`
}

// The type is a representation of a user identity.
//
// The json fields corresponds with the values provided in an oidc token.
type User struct {
	Email           string   `json:"email"`
	IsEmailVerified bool     `json:"email_verified"`
	Name            string   `json:"name"`
	Groups          []string `json:"groups"`
	Audience        string   `json:"aud"`
	Issuer          string   `json:"iss"`
	ExpirationTime  int      `json:"exp"`
}

// Function returns the id of the identity.
//
// User is represented by email, cluster by clusterid and service by service name
func (identity *Identity) GetId() string {
	switch identity.Type {
	case IdentityTypeUser:
		return identity.User.Email
	case IdentityTypeCluster:
		return identity.ClusterIdentity.Id
	case IdentityTypeService:
		return identity.ServiceIdentity.Id
	default:
		return ""
	}
}

// Function returns true if identity is an user
func (identity *Identity) IsUser() bool {
	return identity.Type == IdentityTypeUser
}

// Function returns true if identity is a cluster
func (identity *Identity) IsCluster() bool {
	return identity.Type == IdentityTypeCluster
}

// Function returns true if identity is a service
func (identity *Identity) IsService() bool {
	return identity.Type == IdentityTypeService
}

// GetGroups returns the ACL groups of the identity. Every identity type
// resolves through the same group mechanism: user groups come from the identity
// provider, while cluster and service groups are ROR-owned principal names.
func (identity *Identity) GetGroups() ([]string, error) {
	switch identity.Type {
	case IdentityTypeUser:
		if identity.User == nil {
			return nil, errors.New("user identity has nil user")
		}
		// Groups in the ROR-owned domain grant authorization directly, so one
		// supplied by the identity provider would allow impersonating a cluster,
		// service or service account.
		return aclprincipal.SanitizeExternalGroups(identity.User.Groups), nil
	case IdentityTypeCluster:
		if identity.ClusterIdentity == nil || identity.ClusterIdentity.Uid == "" {
			return nil, errors.New("cluster identity has no uid")
		}
		return aclprincipal.ClusterGroups(identity.ClusterIdentity.Uid), nil
	case IdentityTypeService:
		if identity.ServiceIdentity == nil || identity.ServiceIdentity.Id == "" {
			return nil, errors.New("service identity has no id")
		}
		return aclprincipal.ServiceGroups(identity.ServiceIdentity.Id), nil
	default:
		return nil, errors.New("type not implemented")
	}
}

// Function returns a bson.A containing the groups of an identity. To be used in filtering in mongodb.
func (identity Identity) ReturnGroupQuery() (bson.A, error) {
	groups, err := identity.GetGroups()
	if err != nil {
		return nil, err
	}

	filterGroups := bson.A{}
	for _, group := range groups {
		filterGroups = append(filterGroups, group)
	}
	return filterGroups, nil
}

// Function returns the auth info of the identity
func (identity *Identity) GetAuthInfo() AuthInfo {
	return identity.Auth
}

func (identity *Identity) SetToken(token string) {
	identity.token = token
}

func (identity *Identity) GetToken() string {
	return identity.token
}

// The type is a representation of a cluster or service identity. May be splited if needed.
type ServiceIdentity struct {
	Id  string `json:"id"`
	Uid string `json:"uid,omitempty"`
}
