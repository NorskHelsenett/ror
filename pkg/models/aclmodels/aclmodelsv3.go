package aclmodels

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclcaps"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

// Capability, Verb and AccessTypeV3 live in the aclcaps leaf package so they can
// be shared with rordefs without an import cycle. They are aliased here so the
// existing aclmodels.* API is unchanged.
type (
	Capability   = aclcaps.Capability
	Verb         = aclcaps.Verb
	AccessTypeV3 = aclcaps.AccessTypeV3
)

// Well-known verbs.
const (
	VerbRead     = aclcaps.VerbRead
	VerbWrite    = aclcaps.VerbWrite
	VerbCreate   = aclcaps.VerbCreate
	VerbUpdate   = aclcaps.VerbUpdate
	VerbDelete   = aclcaps.VerbDelete
	VerbAdmin    = aclcaps.VerbAdmin
	VerbLogon    = aclcaps.VerbLogon
	VerbOwner    = aclcaps.VerbOwner
	VerbReadonly = aclcaps.VerbReadonly
)

// Well-known capabilities (without verb).
const (
	CapRor              = aclcaps.CapRor
	CapRorMetadata      = aclcaps.CapRorMetadata
	CapRorVulnerability = aclcaps.CapRorVulnerability
	CapRorConfig        = aclcaps.CapRorConfig

	CapKubernetes              = aclcaps.CapKubernetes
	CapKubernetesArgocd        = aclcaps.CapKubernetesArgocd
	CapKubernetesArgocdProject = aclcaps.CapKubernetesArgocdProject
	CapKubernetesGrafana       = aclcaps.CapKubernetesGrafana

	CapVirtualmachine = aclcaps.CapVirtualmachine
)

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

type AclV3List []AclV3ListItem

type AclV3ListByGroup map[string]AclV3List
type AclV3ListByScopeSubject map[aclscope.Scope]map[aclscope.Subject]AclV3List

func (a AclV3ListByGroup) Flatten() AclV3List {
	out := AclV3List{}
	for _, entries := range a {
		for _, entry := range entries {
			out = append(out, entry)
		}
	}
	return out
}

func (a AclV3List) ByGroup() AclV3ListByGroup {
	out := AclV3ListByGroup{}
	for _, entry := range a {
		out[entry.Group] = append(out[entry.Group], entry)
	}
	return out
}

func (a AclV3List) ByScopeSubject() AclV3ListByScopeSubject {
	out := AclV3ListByScopeSubject{}
	for _, entry := range a {
		if _, ok := out[entry.Scope]; !ok {
			out[entry.Scope] = make(map[aclscope.Subject]AclV3List)
		}
		out[entry.Scope][entry.Subject] = append(out[entry.Scope][entry.Subject], entry)
	}
	return out
}
