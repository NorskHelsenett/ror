package aclmodels

import (
	"time"
)

type AccessType string

const (
	AccessTypeRead             AccessType = "read"
	AccessTypeCreate           AccessType = "create"
	AccessTypeUpdate           AccessType = "update"
	AccessTypeDelete           AccessType = "delete"
	AccessTypeOwner            AccessType = "owner"
	AccessTypeRorMetadata      AccessType = "rormetadata"
	AccessTypeRorVulnerability AccessType = "rorvulnerability"
	AccessTypeClusterLogon     AccessType = "clusterlogon"
)

type AclV2ListItems struct {
	Scope   Acl2Scope           // Type of object ['cluster','project']
	Subject Acl2Subject         // The subject eg. clusterid, projectid (can be 'All')
	Global  AclV2ListItemAccess //If global access granted
	Items   []AclV2ListItem     // v2 access model for ror api
}

// Full acl v2 model
type AclV2ListItem struct {
	Id      string              `json:"id" bson:"_id,omitempty"`                   // Id
	Version int                 `json:"version" default:"2" validate:"eq=2" `      // Acl Version, must be 2
	Group   string              `json:"group" validate:"required,min=1,rortext" `  // The group which the acces is granted
	Scope   Acl2Scope           `json:"scope" validate:"required,min=1,rortext"`   // Type of object ['cluster','project']
	Subject Acl2Subject         `json:"subject" validate:"required,min=1,rortext"` // The subject eg. clusterid, projectid (can be 'All')
	Access  AclV2ListItemAccess `json:"access" validate:"required"`                // v2 access model for ror api
	//	Accessv2   []map[AccessType]bool    `json:"accessv2" validate:""`                      // v2 access model for ror api
	Kubernetes AclV2ListItemKubernetes `json:"kubernetes" validate:""` // v2 access model for kubernetes
	Created    time.Time               `json:"created"`
	IssuedBy   string                  `json:"issuedBy,omitempty" validate:"email"`
}

func NewAclV2ListItem(group string,
	scope Acl2Scope,
	subject Acl2Subject,
	access AclV2ListItemAccess,
	kubernetesLogon bool,
	issuedBy string,
) *AclV2ListItem {
	return &AclV2ListItem{
		Id:      "",
		Version: 2,
		Group:   group,
		Scope:   scope,
		Subject: subject,
		Access:  access,
		Kubernetes: AclV2ListItemKubernetes{
			Logon: kubernetesLogon,
		},
		Created:  time.Now(),
		IssuedBy: issuedBy,
	}
}

// v2 access model for ror api
type AclV2ListItemAccess struct {
	Read   bool `json:"read" validate:"boolean"`   // Read metadata of subject
	Create bool `json:"create" validate:"boolean"` // Write metadata of subject
	Update bool `json:"update" validate:"boolean"` // Update metadata of subject
	Delete bool `json:"delete" validate:"boolean"` // Delete metadata of subject
	Owner  bool `json:"owner" validate:"boolean"`  // Delete metadata of subject
}

// NewAclV2ListItemAccess construct a new AclV2ListItemAccess object.
func NewAclV2ListItemAccess(read, create, update, delete, owner bool) AclV2ListItemAccess {
	return AclV2ListItemAccess{
		Read:   read,
		Create: create,
		Update: update,
		Delete: delete,
		Owner:  owner,
	}
}

// NewAclV2ListItemAccessReadOnly gives you Read access.
func NewAclV2ListItemAccessReadOnly() AclV2ListItemAccess {
	return NewAclV2ListItemAccess(true, false, false, false, false)
}

// NewAclV2ListItemAccessCreateOnly gives you Read and Create  access.
func NewAclV2ListItemAccessCreateOnly() AclV2ListItemAccess {
	return NewAclV2ListItemAccess(true, true, false, false, false)
}

// NewAclV2ListItemAccessEditor gives you Read and Update access.
func NewAclV2ListItemAccessEditor() AclV2ListItemAccess {
	return NewAclV2ListItemAccess(true, false, true, false, false)
}

// NewAclV2ListItemAccessContributor gives you Read, Create, and Update access.
func NewAclV2ListItemAccessContributor() AclV2ListItemAccess {
	return NewAclV2ListItemAccess(true, true, true, false, false)
}

// NewAclV2ListItemAccessContributor gives you Read, Create, Update, and Delete access.
func NewAclV2ListItemAccessOperator() AclV2ListItemAccess {
	return NewAclV2ListItemAccess(true, true, true, true, false)
}

// NewAclV2ListItemAccessAll gives you Read, Create, Update, Delete, and Owner access.
func NewAclV2ListItemAccessAll() AclV2ListItemAccess {
	return NewAclV2ListItemAccess(true, true, true, true, true)
}

// v2 access model for kubernetes
type AclV2ListItemKubernetes struct {
	Logon bool `json:"logon,omitempty" validate:"boolean"` // Logon to subject if 'cluster'
}
