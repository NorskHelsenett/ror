package aclmodels

import "github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"

// AclV3LookupOwnerref is a single scope+subject pair the caller has access to.
type AclV3LookupOwnerref struct {
	Scope   aclscope.Scope   `json:"scope"`
	Subject aclscope.Subject `json:"subject"`
}

// AclV3LookupResponse is the response for GET /v2/acl/lookup. It lists the
// scope+subject pairs the caller has the requested access type for, resolved
// through the V3 ACL backend. When Unrestricted is true the caller has global
// access for the requested access type and Ownerrefs is empty.
type AclV3LookupResponse struct {
	Access       AccessTypeV3          `json:"access"`
	Unrestricted bool                  `json:"unrestricted"`
	Ownerrefs    []AclV3LookupOwnerref `json:"ownerrefs"`
}
