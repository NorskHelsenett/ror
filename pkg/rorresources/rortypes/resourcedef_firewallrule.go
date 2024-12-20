package rortypes

import "time"

type ResourceFirewallRule struct {
	Spec   ResourceFirewallRuleSpec
	Status ResourceFirewallRuleStatus
}

type ResourceFirewallRuleSpec struct {

	// Display name of the rule
	Name string

	// Free text description
	Description string

	// IPv4 (4), IPv6 (6), IPv4 && IPv6 (46)
	IpType int

	// ALLOW, DROP, DENY
	Action string

	// The direction of the traffic flow
	Direction string

	// TCP, UDP, ANY
	Protocol string

	// The unique id for the instance of a rule
	Id string

	// In what order is this rule applied, the lower the earlier
	Sequence int

	// Ip addresses or group references
	Sources []string

	// Ip addresses or group references
	Destinations []string

	// Port number or service
	Services []string

	// Inclusion group that will have this rule applied to
	Scope []string

	// Whether this rule is active or deactivated
	Disabled bool

	// NSX, Fortigate, Checkpoint
	Provider string
}

type ResourceFirewallRuleStatus struct {

	// External references to this rule
	ExternalIds map[string]string

	CreatedAt      time.Time
	CreatedBy      string
	lastModifiedAt time.Time
	LastModifiedBy string

	ResourceFirewallRuleSpec
}
