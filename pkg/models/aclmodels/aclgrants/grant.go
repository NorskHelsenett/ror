// Package aclgrants defines the shared ACL grant token claim.
package aclgrants

import (
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclcaps"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

// ClaimKey is the JWT claim name used to carry grants.
const ClaimKey = "ror_grants"

// Grant is a raw, unexpanded ACL grant.
type Grant struct {
	Scope   aclscope.Scope         `json:"scope"`
	Subject aclscope.Subject       `json:"subject"`
	Access  []aclcaps.AccessTypeV3 `json:"access"`
}

// Grants is the value of the ClaimKey claim.
type Grants []Grant
