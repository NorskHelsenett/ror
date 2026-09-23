package identitymodels_test

import (
	"errors"
	"testing"
	"time"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclprincipal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testAuth = identitymodels.AuthInfo{
	AuthProvider:   identitymodels.IdentityProviderOidc,
	AuthProviderID: "alice@example.com",
	ExpirationTime: time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
}

// assertResolvesTo checks every getter of an identity in one place.
func assertResolvesTo(t *testing.T, identity identitymodels.Identity, subject, name, id string, groups []string) {
	t.Helper()
	require.NoError(t, identity.Validate())

	gotSubject, err := identity.GetSubject()
	require.NoError(t, err)
	assert.Equal(t, subject, gotSubject)

	gotName, err := identity.GetName()
	require.NoError(t, err)
	assert.Equal(t, name, gotName)

	gotGroups, err := identity.GetGroups()
	require.NoError(t, err)
	assert.Equal(t, groups, gotGroups)

	assert.Equal(t, id, identity.GetId())
}

// Each constructor must resolve to the subject, name, id and groups its
// principal type is authorized by.
func TestConstructors_Resolve(t *testing.T) {
	const (
		email     = "alice@example.com"
		userName  = "Alice Example"
		clusterID = "prod-cluster-1"
		clusterUI = "6f1c2b7e-0000-4000-8000-000000000001"
		serviceID = "vulnerability-scanner"
	)
	idpGroups := []string{"team-blue@example.com"}

	t.Run("user", func(t *testing.T) {
		built, err := identitymodels.NewUserIdentity(testAuth, email, userName, idpGroups, nil)
		require.NoError(t, err)

		assertResolvesTo(t, built, email, userName, email, idpGroups)
	})

	t.Run("cluster", func(t *testing.T) {
		built, err := identitymodels.NewClusterIdentity(testAuth, clusterID, clusterUI)
		require.NoError(t, err)

		// A cluster's subject is its uid while its id is what it is known by.
		groups := []string{clusterUI + "@cluster.ror.system", "*@cluster.ror.system"}
		assertResolvesTo(t, built, clusterUI, clusterID, clusterID, groups)
	})

	t.Run("service", func(t *testing.T) {
		built, err := identitymodels.NewServiceIdentity(testAuth, serviceID)
		require.NoError(t, err)

		assertResolvesTo(t, built, serviceID, serviceID, serviceID, aclprincipal.ServiceGroups(serviceID))
	})
}

// Reserved ROR-owned groups must never be honoured from an identity provider,
// otherwise a token could impersonate a cluster or service.
func TestNewUserIdentity_DropsReservedGroups(t *testing.T) {
	identity, err := identitymodels.NewUserIdentity(testAuth, "alice@example.com", "Alice",
		[]string{"team-blue@example.com", "abc-123@cluster.ror.system"}, nil)
	require.NoError(t, err)

	groups, err := identity.GetGroups()
	require.NoError(t, err)
	assert.Equal(t, []string{"team-blue@example.com"}, groups)
}

func TestConstructors_RejectIncompleteInput(t *testing.T) {
	tests := []struct {
		name string
		call func() (identitymodels.Identity, error)
	}{
		{"user without email", func() (identitymodels.Identity, error) {
			return identitymodels.NewUserIdentity(testAuth, "", "Alice", nil, nil)
		}},
		{"cluster without uid", func() (identitymodels.Identity, error) {
			return identitymodels.NewClusterIdentity(testAuth, "prod-cluster-1", "")
		}},
		{"cluster without cluster id", func() (identitymodels.Identity, error) {
			return identitymodels.NewClusterIdentity(testAuth, "", "6f1c2b7e")
		}},
		{"service without id", func() (identitymodels.Identity, error) {
			return identitymodels.NewServiceIdentity(testAuth, "")
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.call()
			require.Error(t, err)
			assert.ErrorIs(t, err, identitymodels.ErrIncompleteIdentity)
		})
	}
}

// An identity built as a struct literal has no resolvable state and must fail
// closed rather than panic.
func TestIdentity_MalformedFailsClosed(t *testing.T) {
	tests := []struct {
		name     string
		identity identitymodels.Identity
		wantErr  error
	}{
		{"user literal", identitymodels.Identity{Type: identitymodels.IdentityTypeUser}, identitymodels.ErrIncompleteIdentity},
		{"cluster literal", identitymodels.Identity{Type: identitymodels.IdentityTypeCluster}, identitymodels.ErrIncompleteIdentity},
		{"service literal", identitymodels.Identity{Type: identitymodels.IdentityTypeService}, identitymodels.ErrIncompleteIdentity},
		{"unknown type", identitymodels.Identity{Type: "Robot"}, identitymodels.ErrUnknownIdentityType},
		{"zero value", identitymodels.Identity{}, identitymodels.ErrUnknownIdentityType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() { _ = tt.identity.GetId() })
			assert.Empty(t, tt.identity.GetId())

			err := tt.identity.Validate()
			require.Error(t, err)
			assert.True(t, errors.Is(err, tt.wantErr), "got %v, want %v", err, tt.wantErr)

			_, err = tt.identity.GetGroups()
			assert.Error(t, err)
			_, err = tt.identity.GetSubject()
			assert.Error(t, err)
		})
	}
}

func TestGetEmail_UserOnly(t *testing.T) {
	user, err := identitymodels.NewUserIdentity(testAuth, "alice@example.com", "Alice", nil, nil)
	require.NoError(t, err)
	email, err := user.GetEmail()
	require.NoError(t, err)
	assert.Equal(t, "alice@example.com", email)

	cluster, err := identitymodels.NewClusterIdentity(testAuth, "prod-cluster-1", "6f1c2b7e")
	require.NoError(t, err)
	_, err = cluster.GetEmail()
	assert.ErrorIs(t, err, identitymodels.ErrNotAUser)
}

// GetGroups hands out a copy so a caller cannot mutate the identity's groups.
func TestGetGroups_ReturnsCopy(t *testing.T) {
	identity, err := identitymodels.NewUserIdentity(testAuth, "alice@example.com", "Alice",
		[]string{"team-blue@example.com"}, nil)
	require.NoError(t, err)

	groups, err := identity.GetGroups()
	require.NoError(t, err)
	groups[0] = "admins@example.com"

	unchanged, err := identity.GetGroups()
	require.NoError(t, err)
	assert.Equal(t, []string{"team-blue@example.com"}, unchanged)
}

func TestGetClaim(t *testing.T) {
	claims := map[string]string{"email_verified": "true"}
	identity, err := identitymodels.NewUserIdentity(testAuth, "alice@example.com", "Alice", nil, claims)
	require.NoError(t, err)

	// Claims are copied on construction, so later edits do not leak in.
	claims["email_verified"] = "false"

	value, ok := identity.GetClaim("email_verified")
	assert.True(t, ok)
	assert.Equal(t, "true", value)

	_, ok = identity.GetClaim("missing")
	assert.False(t, ok)
}

func TestIdentity_PreservedAccessors(t *testing.T) {
	identity, err := identitymodels.NewUserIdentity(testAuth, "alice@example.com", "Alice", nil, nil)
	require.NoError(t, err)

	assert.True(t, identity.IsUser())
	assert.False(t, identity.IsCluster())
	assert.False(t, identity.IsService())
	assert.Equal(t, testAuth, identity.GetAuthInfo())

	identity.SetToken("token-value")
	assert.Equal(t, "token-value", identity.GetToken())
}

func TestReturnGroupQuery(t *testing.T) {
	identity, err := identitymodels.NewServiceIdentity(testAuth, "vulnerability-scanner")
	require.NoError(t, err)

	query, err := identity.ReturnGroupQuery()
	require.NoError(t, err)

	groups, err := identity.GetGroups()
	require.NoError(t, err)
	assert.Len(t, query, len(groups))
}
