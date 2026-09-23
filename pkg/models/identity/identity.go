// Package implements models representing identity
package identitymodels

import (
	"errors"
	"fmt"
	"maps"
	"slices"
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
	Auth AuthInfo     `json:"auth"`
	Type IdentityType `json:"type,omitempty"`

	// Unexported so every read goes through a getter that verifies the identity
	// first. Build identities with NewUserIdentity, NewClusterIdentity or
	// NewServiceIdentity; a struct literal resolves to no access.
	subject string
	name    string
	email   string
	groups  []string
	claims  map[string]string

	token string
}

type AuthInfo struct {
	AuthProvider   IdentityProvider `json:"authProvider,omitempty"`
	AuthProviderID string           `json:"authProviderId,omitempty"`
	ExpirationTime time.Time        `json:"expirationTime"`
}

// Errors returned when an identity cannot be resolved. They are sentinels so
// callers can match with errors.Is rather than on message text.
var (
	ErrUnknownIdentityType = errors.New("unknown identity type")
	ErrIncompleteIdentity  = errors.New("incomplete identity")
	ErrNotAUser            = errors.New("identity is not a user")
)

// identityView is the normalised read model of an Identity. Every getter reads
// through it, so the rules for deriving subject, name, email and groups exist
// once regardless of how the identity was built.
type identityView struct {
	subject string
	name    string
	email   string
	groups  []string
}

// validateFor enforces the per-type invariants of a resolved identity.
func (v identityView) validateFor(identityType IdentityType) error {
	switch identityType {
	case IdentityTypeUser:
		if v.email == "" {
			return fmt.Errorf("%w: user identity requires an email", ErrIncompleteIdentity)
		}
	case IdentityTypeCluster:
		// Cluster grants are keyed by uid (the subject); the cluster id is what
		// the cluster is operationally known by. Both are required.
		if v.name == "" {
			return fmt.Errorf("%w: cluster identity requires a cluster id", ErrIncompleteIdentity)
		}
	case IdentityTypeService:
	default:
		return fmt.Errorf("%w: %q", ErrUnknownIdentityType, identityType)
	}
	if v.subject == "" {
		return fmt.Errorf("%w: %s identity requires a subject", ErrIncompleteIdentity, identityType)
	}
	return nil
}

// principalGroups returns the ACL groups a principal resolves to. User groups
// come from the identity provider and are sanitised; cluster and service groups
// are ROR-owned principal names derived from the subject.
func principalGroups(identityType IdentityType, subject string, idpGroups []string) ([]string, error) {
	switch identityType {
	case IdentityTypeUser:
		// Groups in the ROR-owned domain grant authorization directly, so one
		// supplied by the identity provider would allow impersonating a cluster,
		// service or service account.
		return aclprincipal.SanitizeExternalGroups(idpGroups), nil
	case IdentityTypeCluster:
		return aclprincipal.ClusterGroups(subject), nil
	case IdentityTypeService:
		return aclprincipal.ServiceGroups(subject), nil
	default:
		return nil, fmt.Errorf("%w: %q", ErrUnknownIdentityType, identityType)
	}
}

// resolve normalises the identity into the read model every getter reads
// through.
func (identity *Identity) resolve() (identityView, error) {
	v := identityView{
		subject: identity.subject,
		name:    identity.name,
		email:   identity.email,
		groups:  identity.groups,
	}
	if err := v.validateFor(identity.Type); err != nil {
		return identityView{}, err
	}
	return v, nil
}

// Validate reports whether the identity is complete enough to act on.
func (identity *Identity) Validate() error {
	_, err := identity.resolve()
	return err
}

// newIdentity completes and validates an identity produced by a constructor.
func newIdentity(identity Identity, idpGroups []string) (Identity, error) {
	groups, err := principalGroups(identity.Type, identity.subject, idpGroups)
	if err != nil {
		return Identity{}, err
	}
	identity.groups = groups
	if err := identity.Validate(); err != nil {
		return Identity{}, err
	}
	return identity, nil
}

// NewUserIdentity builds a user identity. idpGroups are the groups supplied by
// the identity provider; reserved ROR-owned names are dropped.
func NewUserIdentity(auth AuthInfo, email, name string, idpGroups []string, claims map[string]string) (Identity, error) {
	return newIdentity(Identity{
		Auth:    auth,
		Type:    IdentityTypeUser,
		subject: email,
		name:    name,
		email:   email,
		claims:  maps.Clone(claims),
	}, idpGroups)
}

// NewClusterIdentity builds a cluster identity. The uid is required: cluster
// grants are keyed by uid, so an identity without one resolves to no access.
func NewClusterIdentity(auth AuthInfo, clusterID, uid string) (Identity, error) {
	return newIdentity(Identity{
		Auth:    auth,
		Type:    IdentityTypeCluster,
		subject: uid,
		name:    clusterID,
	}, nil)
}

// NewServiceIdentity builds a service identity.
func NewServiceIdentity(auth AuthInfo, id string) (Identity, error) {
	return newIdentity(Identity{
		Auth:    auth,
		Type:    IdentityTypeService,
		subject: id,
		name:    id,
	}, nil)
}

// Function returns the id of the identity.
//
// User is represented by email, cluster by clusterid and service by service name
func (identity *Identity) GetId() string {
	v, err := identity.resolve()
	if err != nil {
		return ""
	}
	if identity.Type == IdentityTypeCluster {
		return v.name
	}
	return v.subject
}

// GetSubject returns the identifier authorization is derived from: a user's
// email, a cluster's uid or a service's id.
func (identity *Identity) GetSubject() (string, error) {
	v, err := identity.resolve()
	if err != nil {
		return "", err
	}
	return v.subject, nil
}

// GetName returns the human readable identifier: a user's display name, a
// cluster's cluster id or a service's id.
func (identity *Identity) GetName() (string, error) {
	v, err := identity.resolve()
	if err != nil {
		return "", err
	}
	return v.name, nil
}

// GetEmail returns the email of a user identity and fails for any other type.
func (identity *Identity) GetEmail() (string, error) {
	if identity.Type != IdentityTypeUser {
		return "", fmt.Errorf("%w: %q", ErrNotAUser, identity.Type)
	}
	v, err := identity.resolve()
	if err != nil {
		return "", err
	}
	return v.email, nil
}

// GetClaim returns an additional claim captured at authentication.
func (identity *Identity) GetClaim(name string) (string, bool) {
	value, ok := identity.claims[name]
	return value, ok
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
// The returned slice is a copy, so callers cannot mutate the identity.
func (identity *Identity) GetGroups() ([]string, error) {
	v, err := identity.resolve()
	if err != nil {
		return nil, err
	}
	return slices.Clone(v.groups), nil
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
