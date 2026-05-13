package aclmodels_test

import (
	"testing"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/stretchr/testify/assert"
)

// TestNewAclV2ListItem tests the constructor works as intended.
//
// Creates a new object that has an ID generated automatically, and inserts the values as expected.
func TestNewAclV2ListItem(t *testing.T) {

	groupName := "group_name"
	scope := aclmodels.Acl2ScopeRor
	subject := aclmodels.Acl2RorSubjectPrice
	access := aclmodels.NewAclV2ListItemAccessAll()
	kuberneteslogon := false
	issuedBy := "Thomas Vifte"

	actual := aclmodels.NewAclV2ListItem(groupName, scope, subject, access, kuberneteslogon, issuedBy)
	expected := &aclmodels.AclV2ListItem{
		Id:         "",
		Version:    2,
		Group:      groupName,
		Scope:      scope,
		Subject:    subject,
		Access:     access,
		Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: kuberneteslogon},
		Created:    actual.Created,
		IssuedBy:   issuedBy,
	}

	// Checks that it autogenerates an ID.
	assert.Empty(t, expected.Id)

	// Checks that time.Time is created with an value and not set with time.Time{}.
	assert.Equal(t, expected.Created.IsZero(), false)
	assert.Equal(t, expected, actual)
}

// TestNewAclV2ListItemAccess tests the constructor works as intended.
func TestNewAclV2ListItemAccess(t *testing.T) {

	read := true
	create := true
	update := true
	deleteAccess := true
	owner := true

	expected := aclmodels.NewAclV2ListItemAccess(read, create, update, deleteAccess, owner)
	actual := aclmodels.AclV2ListItemAccess{
		Read:   read,
		Create: create,
		Update: update,
		Delete: deleteAccess,
		Owner:  owner,
	}

	assert.Equal(t, expected, actual)
}

// TestNewAclV2ListItemAccessPredefined tests that the predinfed ListItem sets output what is expected of them.
func TestNewAclV2ListItemAccessPredefined(t *testing.T) {

	expected := aclmodels.AclV2ListItemAccess{
		Read:   true,
		Create: false,
		Update: false,
		Delete: false,
		Owner:  false,
	}
	actual := aclmodels.NewAclV2ListItemAccessReadOnly()
	assert.Equal(t, expected, actual)

	expected = aclmodels.AclV2ListItemAccess{
		Read:   true,
		Create: true,
		Update: false,
		Delete: false,
		Owner:  false,
	}
	actual = aclmodels.NewAclV2ListItemAccessCreateOnly()
	assert.Equal(t, expected, actual)

	expected = aclmodels.AclV2ListItemAccess{
		Read:   true,
		Create: false,
		Update: true,
		Delete: false,
		Owner:  false,
	}
	actual = aclmodels.NewAclV2ListItemAccessEditor()
	assert.Equal(t, expected, actual)

	expected = aclmodels.AclV2ListItemAccess{
		Read:   true,
		Create: true,
		Update: true,
		Delete: false,
		Owner:  false,
	}
	actual = aclmodels.NewAclV2ListItemAccessContributor()
	assert.Equal(t, expected, actual)

	expected = aclmodels.AclV2ListItemAccess{
		Read:   true,
		Create: true,
		Update: true,
		Delete: true,
		Owner:  false,
	}
	actual = aclmodels.NewAclV2ListItemAccessOperator()
	assert.Equal(t, expected, actual)

	expected = aclmodels.AclV2ListItemAccess{
		Read:   true,
		Create: true,
		Update: true,
		Delete: true,
		Owner:  true,
	}
	actual = aclmodels.NewAclV2ListItemAccessAll()
	assert.Equal(t, expected, actual)
}

func TestParseAcl2AccessType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantType aclmodels.AccessType
		wantOk   bool
	}{
		{"read", "read", aclmodels.AccessTypeRead, true},
		{"create", "create", aclmodels.AccessTypeCreate, true},
		{"write maps to create", "write", aclmodels.AccessTypeCreate, true},
		{"update", "update", aclmodels.AccessTypeUpdate, true},
		{"delete", "delete", aclmodels.AccessTypeDelete, true},
		{"owner", "owner", aclmodels.AccessTypeOwner, true},
		{"rormetadata", "rormetadata", aclmodels.AccessTypeRorMetadata, true},
		{"rorvulnerability", "rorvulnerability", aclmodels.AccessTypeRorVulnerability, true},
		{"clusterlogon", "clusterlogon", aclmodels.AccessTypeClusterLogon, true},
		{"empty string", "", "", false},
		{"unknown value", "admin", "", false},
		{"uppercase READ", "READ", "", false},
		{"mixed case Read", "Read", "", false},
		{"trailing space", "read ", "", false},
		{"leading space", " read", "", false},
		{"numeric", "123", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := aclmodels.ParseAcl2AccessType(tt.input)
			assert.Equal(t, tt.wantOk, ok)
			assert.Equal(t, tt.wantType, got)
		})
	}
}

func TestHasAccessType(t *testing.T) {
	allAccess := aclmodels.AclV2ListItemAccess{
		Read:            true,
		Create:          true,
		Update:          true,
		Delete:          true,
		Owner:           true,
		KubernetesLogon: true,
	}
	noAccess := aclmodels.AclV2ListItemAccess{}

	t.Run("all true returns true for every valid type", func(t *testing.T) {
		assert.True(t, allAccess.HasAccessType(aclmodels.AccessTypeRead))
		assert.True(t, allAccess.HasAccessType(aclmodels.AccessTypeCreate))
		assert.True(t, allAccess.HasAccessType(aclmodels.AccessTypeUpdate))
		assert.True(t, allAccess.HasAccessType(aclmodels.AccessTypeDelete))
		assert.True(t, allAccess.HasAccessType(aclmodels.AccessTypeOwner))
		assert.True(t, allAccess.HasAccessType(aclmodels.AccessTypeClusterLogon))
	})

	t.Run("all false returns false for every valid type", func(t *testing.T) {
		assert.False(t, noAccess.HasAccessType(aclmodels.AccessTypeRead))
		assert.False(t, noAccess.HasAccessType(aclmodels.AccessTypeCreate))
		assert.False(t, noAccess.HasAccessType(aclmodels.AccessTypeUpdate))
		assert.False(t, noAccess.HasAccessType(aclmodels.AccessTypeDelete))
		assert.False(t, noAccess.HasAccessType(aclmodels.AccessTypeOwner))
		assert.False(t, noAccess.HasAccessType(aclmodels.AccessTypeClusterLogon))
	})

	t.Run("unknown access type returns false", func(t *testing.T) {
		assert.False(t, allAccess.HasAccessType(aclmodels.AccessType("nonexistent")))
		assert.False(t, allAccess.HasAccessType(aclmodels.AccessType("")))
		assert.False(t, allAccess.HasAccessType(aclmodels.AccessType("ADMIN")))
	})

	t.Run("rormetadata and rorvulnerability have no field and return false", func(t *testing.T) {
		assert.False(t, allAccess.HasAccessType(aclmodels.AccessTypeRorMetadata))
		assert.False(t, allAccess.HasAccessType(aclmodels.AccessTypeRorVulnerability))
	})

	t.Run("partial access only matches set fields", func(t *testing.T) {
		readOnly := aclmodels.AclV2ListItemAccess{Read: true}
		assert.True(t, readOnly.HasAccessType(aclmodels.AccessTypeRead))
		assert.False(t, readOnly.HasAccessType(aclmodels.AccessTypeCreate))
		assert.False(t, readOnly.HasAccessType(aclmodels.AccessTypeDelete))
		assert.False(t, readOnly.HasAccessType(aclmodels.AccessTypeOwner))
	})
}
