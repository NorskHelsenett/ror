package aclmodels

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rordefs"
)

var ValidSystems = map[string]bool{
	"ror":   true,
	"spam":  true,
	"alarm": true,
	"all":   true,
}

// ValidScope checks if a scope is either a known resource kind or a known system.
// Resource kinds are resolved at runtime from rordefs.Resourcedefs.
func ValidScope(scope aclscope.Scope) error {
	s := string(scope)

	if ValidSystems[s] {
		return nil
	}

	for _, r := range rordefs.Resourcedefs {
		if r.GetKind() == s {
			return nil
		}
	}

	return fmt.Errorf("unknown scope %q: must be a known resource kind or system", s)
}
