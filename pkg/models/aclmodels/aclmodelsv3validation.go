package aclmodels

import (
	"fmt"
	"strings"
)

// AccessNode represents a node in the access type validation tree.
// Each node can have valid verbs and child components.
type AccessNode struct {
	Verbs    map[string]bool        // valid verbs at this level
	Children map[string]*AccessNode // sub-components
}

// ValidVerbs is the set of all valid verbs that can appear as the last segment of an AccessTypeV3.
var ValidVerbs = map[Verb]bool{
	VerbRead:     true,
	VerbWrite:    true,
	VerbCreate:   true,
	VerbUpdate:   true,
	VerbDelete:   true,
	VerbAdmin:    true,
	VerbLogon:    true,
	VerbOwner:    true,
	VerbReadonly: true,
}

// accessTree defines the valid access type paths and their allowed verbs.
// To extend, add child nodes or verbs at the appropriate level.
var accessTree = &AccessNode{
	Children: map[string]*AccessNode{
		"ror": {
			Verbs: map[string]bool{"read": true, "write": true, "create": true, "update": true, "delete": true, "owner": true},
			Children: map[string]*AccessNode{
				"metadata": {
					Verbs: map[string]bool{"write": true},
				},
				"vulnerability": {
					Verbs: map[string]bool{"read": true, "write": true},
				},
				"config": {
					Verbs: map[string]bool{"read": true, "write": true},
				},
			},
		},
		"kubernetes": {
			Verbs: map[string]bool{"logon": true, "admin": true, "readonly": true},
			Children: map[string]*AccessNode{
				"argocd": {
					Verbs: map[string]bool{"admin": true},
					Children: map[string]*AccessNode{
						"project": {
							Verbs: map[string]bool{"admin": true},
						},
					},
				},
				"grafana": {
					Verbs: map[string]bool{"admin": true},
				},
			},
		},
		"resource": {
			Children: map[string]*AccessNode{
				"*": { // wildcard node — accepts any component name (resource kind)
					Verbs: map[string]bool{"read": true, "write": true, "delete": true},
				},
			},
		},
		"virtualmachine": {
			Verbs: map[string]bool{"delete": true},
		},
	},
}

// ValidateAccess validates that an AccessTypeV3 string follows the system:component:verb
// convention and that the path and verb are registered in the access tree.
func ValidateAccess(access AccessTypeV3) error {
	cap, verb := access.Parse()
	if verb == "" {
		return fmt.Errorf("access type must have at least system:verb, got %q", access)
	}

	if !ValidVerbs[verb] {
		return fmt.Errorf("unknown verb %q in %q", verb, access)
	}

	path := strings.Split(string(cap), ":")

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

	if node.Verbs == nil || !node.Verbs[string(verb)] {
		return fmt.Errorf("verb %q not valid at path %q in %q", verb, cap, access)
	}
	return nil
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
