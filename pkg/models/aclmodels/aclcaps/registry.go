package aclcaps

import (
	"fmt"
	"strings"
)

// Namespace is a node in the capability registry. Each node declares the verbs
// valid at its path and, optionally, child components or a wildcard child that
// matches any single segment (used for resource:<Kind>).
type Namespace struct {
	Verbs    map[Verb]struct{}
	Children map[string]*Namespace
	Wildcard *Namespace
}

// AllVerbs is the canonical set of verbs that may appear as the last segment of
// an AccessTypeV3. Every verb used in the Registry must be a member.
var AllVerbs = map[Verb]struct{}{
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

// Registry is the single source of truth for valid capability namespaces and the
// verbs allowed at each path. It replaces the previously hardcoded access tree
// and is the data other packages (aclmodels validation, egress audiences) derive
// from. Each top-level key is a capability namespace (the prefix of an audience).
var Registry = &Namespace{
	Children: map[string]*Namespace{
		"ror": {
			Verbs: verbs(VerbRead, VerbWrite, VerbCreate, VerbUpdate, VerbDelete, VerbOwner),
			Children: map[string]*Namespace{
				"metadata":      {Verbs: verbs(VerbWrite)},
				"vulnerability": {Verbs: verbs(VerbRead, VerbWrite)},
				"config":        {Verbs: verbs(VerbRead, VerbWrite)},
			},
		},
		"kubernetes": {
			Verbs: verbs(VerbLogon, VerbAdmin, VerbReadonly),
			Children: map[string]*Namespace{
				"argocd": {
					Verbs:    verbs(VerbAdmin),
					Children: map[string]*Namespace{"project": {Verbs: verbs(VerbAdmin)}},
				},
				"grafana": {Verbs: verbs(VerbAdmin)},
			},
		},
		"resource": {
			// Wildcard accepts any resource Kind: resource:<Kind>:<verb>.
			Wildcard: &Namespace{Verbs: verbs(VerbRead, VerbWrite, VerbDelete)},
		},
		"virtualmachine": {Verbs: verbs(VerbDelete)},
		"monitoring":     {Verbs: verbs(VerbRead, VerbWrite)},
		"dns":            {Verbs: verbs(VerbRead, VerbWrite)},
	},
}

// verbs builds a verb set from the given verbs.
func verbs(vs ...Verb) map[Verb]struct{} {
	m := make(map[Verb]struct{}, len(vs))
	for _, v := range vs {
		m[v] = struct{}{}
	}
	return m
}

// Validate checks that an AccessTypeV3 follows the system:component[:...]:verb
// convention and that the path and verb are registered in the Registry.
func Validate(a AccessTypeV3) error {
	capability, verb := a.Parse()
	if verb == "" {
		return fmt.Errorf("access type must have at least system:verb, got %q", a)
	}
	if _, ok := AllVerbs[verb]; !ok {
		return fmt.Errorf("unknown verb %q in %q", verb, a)
	}

	node := Registry
	path := strings.Split(string(capability), ":")
	for i, segment := range path {
		if child, ok := node.Children[segment]; ok {
			node = child
			continue
		}
		if node.Wildcard != nil {
			node = node.Wildcard
			continue
		}
		return fmt.Errorf("unknown path segment %q at position %d in %q", segment, i, a)
	}

	if _, ok := node.Verbs[verb]; !ok {
		return fmt.Errorf("verb %q not valid at path %q in %q", verb, capability, a)
	}
	return nil
}

// ValidCapability reports whether a Capability (without verb) resolves to a
// registered node that can carry verbs.
func ValidCapability(c Capability) bool {
	if c == "" {
		return false
	}
	node := Registry
	for _, segment := range strings.Split(string(c), ":") {
		if child, ok := node.Children[segment]; ok {
			node = child
			continue
		}
		if node.Wildcard != nil {
			node = node.Wildcard
			continue
		}
		return false
	}
	return len(node.Verbs) > 0 || node.Wildcard != nil
}
