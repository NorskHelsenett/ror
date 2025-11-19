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
	expected := aclmodels.AclV2ListItem{
		Id:         actual.Id,
		Group:      groupName,
		Scope:      scope,
		Subject:    subject,
		Access:     access,
		Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: kuberneteslogon},
		Created:    actual.Created,
		IssuedBy:   issuedBy,
	}

	// Checks that it autogenerates an ID.
	assert.NotEmpty(t, expected.Id)

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
