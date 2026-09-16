// Package aclcaps holds the low-level capability, verb and access-type
// primitives shared by aclmodels and rordefs. It intentionally has no
// dependencies on other ror packages so it can be imported from both sides
// without creating an import cycle.
package aclcaps

import "strings"

// Capability represents the system:component path of an access type, without the verb.
// Example: "ror", "ror:vulnerability", "kubernetes:argocd", "resource:Deployment"
type Capability string

// WithVerb builds a full AccessTypeV3 by appending the verb.
// Example: CapRorConfig.WithVerb(VerbRead) → "ror:config:read"
func (c Capability) WithVerb(v Verb) AccessTypeV3 {
	return AccessTypeV3(c.String() + ":" + v.String())
}

func (c Capability) String() string {
	return string(c)
}

// Verb represents the action part of an access type.
type Verb string

func (v Verb) String() string {
	return string(v)
}

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

func (a AccessTypeV3) String() string {
	return string(a)
}
