package aclmodels

import (
	"slices"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
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
	slices.Sort(result)
	return result
}

// CompileAccess merges access from multiple ACL entries that match the given scope and subject,
// returning the union of all granted access types.
func CompileAccess(entries []AclV3ListItem, scope aclscope.Scope, subject aclscope.Subject) []AccessTypeV3 {
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
	slices.Sort(result)
	return result
}
