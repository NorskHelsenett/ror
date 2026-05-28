package aclmodels

import (
	"slices"
	"strings"
)

// HasAccess checks if the given access type is granted in this ACL entry.
func (a *AclV3ListItem) HasAccess(access AccessTypeV3) bool {
	return slices.Contains(a.Access, access)
}

// MergeAccess returns the union of two access slices, deduplicated.
func MergeAccess(a, b []AccessTypeV3) []AccessTypeV3 {
	seen := make(map[AccessTypeV3]struct{}, len(a)+len(b))
	for _, v := range a {
		seen[v] = struct{}{}
	}
	for _, v := range b {
		seen[v] = struct{}{}
	}
	result := make([]AccessTypeV3, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}
	return result
}

// MatchPrefix returns all access types from the slice that start with the given prefix.
// Example: MatchPrefix(access, "resource:") returns all resource-kind entries.
func MatchPrefix(access []AccessTypeV3, prefix string) []AccessTypeV3 {
	var result []AccessTypeV3
	for _, a := range access {
		if strings.HasPrefix(string(a), prefix) {
			result = append(result, a)
		}
	}
	return result
}

// CanAccessKind checks if the access list grants the given verb on a resource kind.
// Returns true if either the wildcard "resource:*:<verb>" or the specific
// "resource:<kind>:<verb>" is present in the access list.
func CanAccessKind(access []AccessTypeV3, kind string, verb string) bool {
	wildcard := AccessTypeV3("resource:*:" + verb)
	specific := AccessTypeV3("resource:" + kind + ":" + verb)
	return slices.Contains(access, wildcard) || slices.Contains(access, specific)
}

// CompileAccess merges access from multiple ACL entries that match the given scope and subject,
// returning the union of all granted access types.
func CompileAccess(entries []AclV3ListItem, scope Acl3Scope, subject Acl3Subject) []AccessTypeV3 {
	seen := make(map[AccessTypeV3]struct{})
	for _, entry := range entries {
		if entry.Scope == scope && entry.Subject == subject {
			for _, a := range entry.Access {
				seen[a] = struct{}{}
			}
		}
	}
	result := make([]AccessTypeV3, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}
	return result
}

// AllowedKinds extracts all explicitly granted resource kinds for a given verb from the access list.
// Returns nil if wildcard access is granted (resource:*:<verb>), meaning all kinds are allowed.
// Returns an empty slice if no resource kind access is granted for the verb.
func AllowedKinds(access []AccessTypeV3, verb string) []string {
	wildcardPrefix := "resource:*:" + verb
	if slices.Contains(access, AccessTypeV3(wildcardPrefix)) {
		return nil // nil means all kinds allowed
	}

	suffix := ":" + verb
	kinds := make([]string, 0)
	for _, a := range access {
		s := string(a)
		if strings.HasPrefix(s, "resource:") && strings.HasSuffix(s, suffix) {
			parts := strings.Split(s, ":")
			if len(parts) == 3 && parts[1] != "*" {
				kinds = append(kinds, parts[1])
			}
		}
	}
	return kinds
}
