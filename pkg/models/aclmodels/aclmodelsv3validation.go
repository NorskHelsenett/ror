package aclmodels

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclcaps"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

// ValidateAccess validates that an AccessTypeV3 string follows the system:component:verb
// convention and that the path and verb are registered in the capability registry.
func ValidateAccess(access AccessTypeV3) error {
	return aclcaps.Validate(access)
}

// ParseAccessTypeV3 casts a string to an AccessTypeV3 and validates it against the access tree.
func ParseAccessTypeV3(s string) (AccessTypeV3, error) {
	access := AccessTypeV3(s)
	if err := ValidateAccess(access); err != nil {
		return "", err
	}
	return access, nil
}

// ValidateACLEntry validates the scope and all access entries of an AclV3ListItem.
func ValidateACLEntry(entry AclV3ListItem) error {
	if err := ValidScope(entry.Scope); err != nil {
		return fmt.Errorf("invalid ACL entry: %w", err)
	}
	// "all" is a lookup wildcard, not a grant value: a global grant is written as
	// scope "ror" with subject "globalscope".
	if entry.Scope == aclscope.ScopeAll || entry.Subject == aclscope.SubjectAll {
		return fmt.Errorf("invalid ACL entry: scope/subject \"all\" is not a valid grant value; use scope %q subject %q for a global grant", aclscope.ScopeRor, aclscope.SubjectGlobal)
	}
	for _, a := range entry.Access {
		if err := ValidateAccess(a); err != nil {
			return fmt.Errorf("invalid ACL entry: %w", err)
		}
	}
	return nil
}
