package aclmodels

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rordefs"
)

// ValidScope checks if a scope is either a known resource kind or a known system.
// System scopes (e.g. "ror", "all", "spam") are owned by the aclscope package;
// resource kinds are resolved at runtime from rordefs.Resourcedefs.
func ValidScope(scope aclscope.Scope) error {
	if scope.IsValid() {
		return nil
	}

	for _, r := range rordefs.Resourcedefs {
		if r.GetKind() == string(scope) {
			return nil
		}
	}

	return fmt.Errorf("unknown scope %q: must be a known resource kind or system", scope)
}
