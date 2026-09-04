package aclmodels

import (
	"fmt"
	"strings"
)

// accessNode represents a node in the access type validation tree.
// Each node can have valid verbs and child components.
type accessNode struct {
	Verbs    map[Verb]struct{}      // valid verbs at this level
	Children map[string]*accessNode // sub-components
}

// validVerbs is the set of all valid verbs that can appear as the last segment of an AccessTypeV3.
var validVerbs = map[Verb]struct{}{
	VerbRead:     {},
	VerbWrite:    {},
	VerbCreate:   {},
	VerbUpdate:   {},
	VerbDelete:   {},
	VerbAdmin:    {},
	VerbLogon:    {},
	VerbOwner:    {},
	VerbReadonly: {},
}

// accessTree defines the valid access type paths and their allowed verbs.
// To extend, add child nodes or verbs at the appropriate level.
var accessTree = &accessNode{
	Children: map[string]*accessNode{
		"ror": {
			Verbs: map[Verb]struct{}{"read": {}, "write": {}, "create": {}, "update": {}, "delete": {}, "owner": {}},
			Children: map[string]*accessNode{
				"metadata": {
					Verbs: map[Verb]struct{}{"write": {}},
				},
				"vulnerability": {
					Verbs: map[Verb]struct{}{"read": {}, "write": {}},
				},
				"config": {
					Verbs: map[Verb]struct{}{"read": {}, "write": {}},
				},
			},
		},
		"kubernetes": {
			Verbs: map[Verb]struct{}{"logon": {}, "admin": {}, "readonly": {}},
			Children: map[string]*accessNode{
				"argocd": {
					Verbs: map[Verb]struct{}{"admin": {}},
					Children: map[string]*accessNode{
						"project": {
							Verbs: map[Verb]struct{}{"admin": {}},
						},
					},
				},
				"grafana": {
					Verbs: map[Verb]struct{}{"admin": {}},
				},
			},
		},
		"resource": {
			Children: map[string]*accessNode{
				"*": { // wildcard node — accepts any component name (resource kind)
					Verbs: map[Verb]struct{}{"read": {}, "write": {}, "delete": {}},
				},
			},
		},
		"virtualmachine": {
			Verbs: map[Verb]struct{}{"delete": {}},
		},
	},
}

// ValidateAccess validates that an AccessTypeV3 string follows the system:component:verb
// convention and that the path and verb are registered in the access tree.
func ValidateAccess(access AccessTypeV3) error {
	capability, verb := access.Parse()
	if verb == "" {
		return fmt.Errorf("access type must have at least system:verb, got %q", access)
	}

	if _, ok := validVerbs[verb]; !ok {
		return fmt.Errorf("unknown verb %q in %q", verb, access)
	}

	path := strings.Split(string(capability), ":")

	node := accessTree
	for i, segment := range path {
		// Check for exact match
		if child, ok := node.Children[segment]; ok {
			node = child
			continue
		}
		// Check for wildcard node (accepts any value at this level)
		if wild, ok := node.Children["*"]; ok {
			node = wild
			continue
		}
		return fmt.Errorf("unknown path segment %q at position %d in %q", segment, i, access)
	}

	if _, ok := node.Verbs[verb]; !ok {
		return fmt.Errorf("verb %q not valid at path %q in %q", verb, capability, access)
	}
	return nil
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
	for _, a := range entry.Access {
		if err := ValidateAccess(a); err != nil {
			return fmt.Errorf("invalid ACL entry: %w", err)
		}
	}
	return nil
}
