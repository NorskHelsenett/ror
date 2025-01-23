package rortypes

import "time"

type ResourceFirewallPolicy struct {
	Spec   ResourceFirewallPolicySpec   `json:"spec"`
	Status ResourceFirewallPolicyStatus `json:"status"`
}

type ResourceFirewallPolicySpec struct {

	// Display name of the policy
	Name string `json:"name"`

	// Free text description
	Description string `json:"description"`

	// Unique id for the instance of a policy
	Id string `json:"id"`

	// Any applicable rules to this policy
	Rules []ResourceFirewallRule `json:"rules"`
}

type ResourceFirewallPolicyStatus struct {

	// External references to this policy
	ExternalIds map[string]string `json:"externalIds"`

	ResourceCommonFieldsStatus
	ResourceFirewallPolicySpec
}

type ResourceFirewallRule struct {
	Spec   ResourceFirewallRuleSpec   `json:"spec"`
	Status ResourceFirewallRuleStatus `json:"status"`
}

type ResourceCommonFieldsStatus struct {
	CreatedAt      time.Time `json:"createdAt"`
	CreatedBy      string    `json:"createdBy"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	LastModifiedBy string    `json:"lastModifiedBy"`
}

type ResourceFirewallRuleSpec struct {

	// Display name of the rule
	Name string `json:"name"`

	// Free text description
	Description string `json:"description"`

	// IPv4 (4), IPv6 (6), IPv4 && IPv6 (46)
	IpType int `json:"ipType"`

	// ALLOW, DROP, DENY
	Action string `json:"action"`

	// The direction of the traffic flow
	Direction string `json:"direction"`

	// TCP, UDP, ANY
	Protocol string `json:"protocol"`

	// The unique id for the instance of a rule
	Id string `json:"id"`

	// In what order is this rule applied, the lower the earlier
	Sequence int `json:"sequence"`

	// Ip addresses or group references
	Sources []string `json:"sources"`

	// Ip addresses or group references
	Destinations []string `json:"destinations"`

	// Port number or service
	Services []string `json:"services"`

	// Inclusion group that will have this rule applied to
	Scope []string `json:"scope"`

	// Whether this rule is active or deactivated
	Disabled bool `json:"disabled"`

	// NSX, Fortigate, Checkpoint
	Provider string `json:"provider"`
}

type ResourceFirewallRuleStatus struct {

	// External references to this rule
	ExternalIds map[string]string `json:"externalIds"`

	ResourceCommonFieldsStatus

	ResourceFirewallRuleSpec
}
