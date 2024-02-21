package aclmodels

import (
	"time"
)

type AclV2ListItems struct {
	Scope   Acl2Scope           // Type of object ['cluster','project']
	Subject Acl2Subject         // The subject eg. clusterid, projectid (can be 'All')
	Global  AclV2ListItemAccess //If global access granted
	Items   []AclV2ListItem     // v2 access model for ror api
}

// Full acl v2 model
type AclV2ListItem struct {
	Id         string                  `json:"id" bson:"_id,omitempty"`                   // Id
	Version    int                     `json:"version" default:"2" validate:"eq=2" `      // Acl Version, must be 2
	Group      string                  `json:"group" validate:"required,min=1,rortext" `  // The group wich the acces is granted
	Scope      Acl2Scope               `json:"scope" validate:"required,min=1,rortext"`   // Type of object ['cluster','project']
	Subject    Acl2Subject             `json:"subject" validate:"required,min=1,rortext"` // The subject eg. clusterid, projectid (can be 'All')
	Access     AclV2ListItemAccess     `json:"access" validate:"required"`                // v2 access model for ror api
	Kubernetes AclV2ListItemKubernetes `json:"kubernetes" validate:""`                    // v2 access model for kubernetes
	Created    time.Time               `json:"created,omitempty"`
	IssuedBy   string                  `json:"issuedBy,omitempty" validate:"email"`
}

// v2 access model for ror api
type AclV2ListItemAccess struct {
	Read   bool `json:"read" validate:"boolean"`   // Read metadata of subject
	Create bool `json:"create" validate:"boolean"` // Write metadata of subject
	Update bool `json:"update" validate:"boolean"` // Update metadata of subject
	Delete bool `json:"delete" validate:"boolean"` // Delete metadata of subject
	Owner  bool `json:"owner" validate:"boolean"`  // Delete metadata of subject
}

// v2 access model for kubernetes
type AclV2ListItemKubernetes struct {
	Logon bool `json:"logon,omitempty" validate:"boolean"` // Logon to subject if 'cluster'
}
