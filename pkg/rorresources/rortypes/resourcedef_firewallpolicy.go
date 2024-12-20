package rortypes

import "time"

type ResourceFirewallPolicy struct {
	Spec   ResourceFirewallPolicySpec
	Status ResourceFirewallPolicyStatus
}

type ResourceFirewallPolicySpec struct {

	// Display name of the policy
	Name string

	// Free text description
	Description string

	// Unique id for the instance of a policy
	Id string

	// Any applicable rules to this policy
	Rules []ResourceFirewallRule
}

type ResourceFirewallPolicyStatus struct {

	// External references to this policy
	ExternalIds map[string]string

	CreatedAt      time.Time
	CreatedBy      string
	lastModifiedAt time.Time
	LastModifiedBy string

	ResourceFirewallPolicySpec
}
