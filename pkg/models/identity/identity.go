// Package implements models representing identity
package identitymodels

import (
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	Auth            AuthInfo         `json:"auth,omitempty"`
	Type            IdentityType     `json:"type,omitempty"`
	User            *User            `json:"user,omitempty"`
	ClusterIdentity *ServiceIdentity `json:"clusterIdentity,omitempty"`
	ServiceIdentity *ServiceIdentity `json:"serviceIdentity,omitempty"`
}

type AuthInfo struct {
	AuthProvider   IdentityProvider `json:"authProvider,omitempty"`
	AuthProviderID string           `json:"authProviderId,omitempty"`
	ExpirationTime time.Time        `json:"expirationTime,omitempty"`
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

// Function returns a bson.A containing the groups of an identity. To be used in filtering in mongodb.
func (identity Identity) ReturnGroupQuery() (bson.A, error) {
	filterGroups := bson.A{}

	switch identity.Type {
	case IdentityTypeCluster:
		return nil, errors.New("clusters dont have groups")
	case IdentityTypeUser:
		for i := 0; i < len(identity.User.Groups); i++ {
			filterGroups = append(filterGroups, identity.User.Groups[i])
		}

		return filterGroups, nil
	case IdentityTypeService:
		filterGroups = append(filterGroups, fmt.Sprintf("service-%s@ror.system", identity.GetId()))
		return filterGroups, nil
	default:
		return nil, errors.New("type not implemented")
	}
}

// Function returns the auth info of the identity
func (identity *Identity) GetAuthInfo() AuthInfo {
	return identity.Auth
}

// The type is a representation of a cluster or service identity. May be splited if needed.
type ServiceIdentity struct {
	Id string `json:"id"`
}
