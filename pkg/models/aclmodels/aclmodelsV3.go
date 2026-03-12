package aclmodels

import "time"

type AclV3ListItemAccess map[string]bool

// system:[subsystem]:verb/access, e.g. ror:read, kubernetes:logon, ror:metadata:write

// ror:read true
// ror:write true
// ror:owner true
// kubernetes:logon true
// kubernetes:admin true
// kubernetes:readonly true
// kubernetes:argocd:admin true
// kubernetes:grafana:admin true
// ror:metadata:write true
// virtualmachine:delete true

type AccessTypeV3 string

type Acl3Scope Acl2Scope
type Acl3Subject Acl2Subject

// AccessType ror:read, ror:write, ror:owner, ror:metadata, ror:vulnerability, kubernetes:logon, kubernetes:admin, kubernetes:readonly, argocd:admin
// kubernetes:test:read

// Full acl v3 model
type AclV3ListItem struct {
	Id      string                `json:"id" bson:"_id,omitempty"`                   // Id
	Version int                   `json:"version" default:"3" validate:"eq=3" `      // Acl Version, must be 3
	Group   string                `json:"group" validate:"required,min=1,rortext" `  // The group which the acces is granted
	Scope   Acl3Scope             `json:"scope" validate:"required,min=1,rortext"`   // Type of object ['cluster','project'] that access is granted to
	Subject Acl3Subject           `json:"subject" validate:"required,min=1,rortext"` // The subject eg. clusterid, projectid (can be 'All') that access is granted to
	Access  map[AccessTypeV3]bool `json:"access" validate:"required"`                // v3 access model for ror api and kubernetes
	//	Access     AclV2ListItemAccess     `json:"access" validate:"required"`                // v2 access model for ror api
	//	Kubernetes AclV2ListItemKubernetes `json:"kubernetes" validate:""`                    // v2 access model for kubernetes
	Created  time.Time `json:"created"`
	IssuedBy string    `json:"issuedBy,omitempty" validate:"email"` // expects an email
}
