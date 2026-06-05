package aclmodels

import (
	"strings"
	"time"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

// Capability represents the system:component path of an access type, without the verb.
// Example: "ror", "ror:vulnerability", "kubernetes:argocd", "resource:Deployment"
type Capability string

// WithVerb builds a full AccessTypeV3 by appending the verb.
// Example: CapRorConfig.WithVerb(VerbRead) → "ror:config:read"
func (c Capability) WithVerb(v Verb) AccessTypeV3 {
	return AccessTypeV3(string(c) + ":" + string(v))
}

// Verb represents the action part of an access type.
type Verb string

// Well-known verbs.
const (
	VerbRead     Verb = "read"
	VerbWrite    Verb = "write"
	VerbCreate   Verb = "create"
	VerbUpdate   Verb = "update"
	VerbDelete   Verb = "delete"
	VerbAdmin    Verb = "admin"
	VerbLogon    Verb = "logon"
	VerbOwner    Verb = "owner"
	VerbReadonly Verb = "readonly"
)

// Well-known capabilities (without verb).
const (
	CapRor              Capability = "ror"
	CapRorMetadata      Capability = "ror:metadata"
	CapRorVulnerability Capability = "ror:vulnerability"
	CapRorConfig        Capability = "ror:config"

	CapKubernetes              Capability = "kubernetes"
	CapKubernetesArgocd        Capability = "kubernetes:argocd"
	CapKubernetesArgocdProject Capability = "kubernetes:argocd:project"
	CapKubernetesGrafana       Capability = "kubernetes:grafana"

	CapVirtualmachine Capability = "virtualmachine"
)

// AccessTypeV3 represents a hierarchical capability string.
// Format: system:component[:subcomponent...]:verb
// The last segment is always the verb. Everything before it is the path.
type AccessTypeV3 string

// Parse splits an AccessTypeV3 into its Capability and Verb parts.
// The verb is the last colon-separated segment; everything before it is the capability.
func (a AccessTypeV3) Parse() (Capability, Verb) {
	s := string(a)
	i := strings.LastIndex(s, ":")
	if i < 0 {
		return Capability(s), ""
	}
	return Capability(s[:i]), Verb(s[i+1:])
}

// Access type constants for the ror system
const (
	AccessRorRead  AccessTypeV3 = "ror:read"
	AccessRorWrite AccessTypeV3 = "ror:write"
	AccessRorOwner AccessTypeV3 = "ror:owner"

	AccessRorMetadataWrite      AccessTypeV3 = "ror:metadata:write"
	AccessRorVulnerabilityRead  AccessTypeV3 = "ror:vulnerability:read"
	AccessRorVulnerabilityWrite AccessTypeV3 = "ror:vulnerability:write"

	AccessRorConfigRead  AccessTypeV3 = "ror:config:read"
	AccessRorConfigWrite AccessTypeV3 = "ror:config:write"
)

// Access type constants for kubernetes
const (
	AccessKubernetesLogon    AccessTypeV3 = "kubernetes:logon"
	AccessKubernetesAdmin    AccessTypeV3 = "kubernetes:admin"
	AccessKubernetesReadonly AccessTypeV3 = "kubernetes:readonly"

	AccessKubernetesArgocdAdmin        AccessTypeV3 = "kubernetes:argocd:admin"
	AccessKubernetesArgocdProjectAdmin AccessTypeV3 = "kubernetes:argocd:project:admin"
	AccessKubernetesGrafanaAdmin       AccessTypeV3 = "kubernetes:grafana:admin"
)

// Access type constants for virtual machines
const (
	AccessVirtualmachineDelete AccessTypeV3 = "virtualmachine:delete"
)

// AclV3ListItem is the full ACL v3 model.
//
// Scope is a resource kind or system identifier.
// Subject is the name/id of the object, e.g. clusterid, projectid, "All".
// Access is a list of granted capabilities — presence means granted, absence means denied.
//
// Example:
//
//	Group: "dev-team", Scope: "KubernetesCluster", Subject: "prod-cluster-1",
//	Access: ["ror:read", "ror:write", "kubernetes:logon", "resource:Deployment:read"]
type AclV3ListItem struct {
	Id       string           `json:"id" bson:"_id,omitempty"`
	Version  int              `json:"version" default:"3" validate:"eq=3"`
	Group    string           `json:"group" validate:"required,min=1,rortext"`
	Scope    aclscope.Scope   `json:"scope" validate:"required,min=1,rortext"`
	Subject  aclscope.Subject `json:"subject" validate:"required,min=1,rortext"`
	Access   []AccessTypeV3   `json:"access" bson:"access" validate:"required"`
	Created  time.Time        `json:"created"`
	IssuedBy string           `json:"issuedBy,omitempty" validate:"email"`
}
