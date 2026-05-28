package aclmodels

import (
	"context"
	"slices"
)

// AclV3Resolver resolves access for a set of groups using an AclV3Store.
// It loads all entries in one batch call, then compiles access in-memory.
type AclV3Resolver struct {
	store AclV3Store
}

// NewAclV3Resolver creates a new resolver with the given store backend.
func NewAclV3Resolver(store AclV3Store) *AclV3Resolver {
	return &AclV3Resolver{store: store}
}

// ResolveAccess loads ACL entries for all given groups (single round-trip) and returns
// the union of access types that match the given scope and subject.
func (r *AclV3Resolver) ResolveAccess(ctx context.Context, groups []string, scope Acl3Scope, subject Acl3Subject) ([]AccessTypeV3, error) {
	entriesByGroup, err := r.store.GetByGroups(ctx, groups)
	if err != nil {
		return nil, err
	}

	seen := make(map[AccessTypeV3]struct{})
	for _, entries := range entriesByGroup {
		for _, entry := range entries {
			if matchesScopeSubject(entry, scope, subject) {
				for _, a := range entry.Access {
					seen[a] = struct{}{}
				}
			}
		}
	}

	result := make([]AccessTypeV3, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}
	return result, nil
}

// ResolveOwnerrefs returns all scope+subject pairs the groups have access to for the given access type.
// Returns nil (meaning unrestricted) if any entry grants global/all access.
func (r *AclV3Resolver) ResolveOwnerrefs(ctx context.Context, groups []string, requiredAccess AccessTypeV3) ([]AclV3Ownerref, error) {
	entriesByGroup, err := r.store.GetByGroups(ctx, groups)
	if err != nil {
		return nil, err
	}

	refs := make([]AclV3Ownerref, 0)
	seen := make(map[AclV3Ownerref]struct{})

	for _, entries := range entriesByGroup {
		for _, entry := range entries {
			if !slices.Contains(entry.Access, requiredAccess) {
				continue
			}
			// Global access — unrestricted
			if entry.Scope == Acl3Scope(Acl2ScopeAll) || entry.Subject == Acl3Subject(Acl2RorSubjectAll) {
				return nil, nil
			}
			ref := AclV3Ownerref{Scope: entry.Scope, Subject: entry.Subject}
			if _, ok := seen[ref]; !ok {
				seen[ref] = struct{}{}
				refs = append(refs, ref)
			}
		}
	}
	return refs, nil
}

// HasAccess checks if the groups have a specific access type for the given scope+subject.
func (r *AclV3Resolver) HasAccess(ctx context.Context, groups []string, scope Acl3Scope, subject Acl3Subject, requiredAccess AccessTypeV3) (bool, error) {
	access, err := r.ResolveAccess(ctx, groups, scope, subject)
	if err != nil {
		return false, err
	}
	return slices.Contains(access, requiredAccess), nil
}

// AclV3Ownerref represents a scope+subject pair that a user has access to.
type AclV3Ownerref struct {
	Scope   Acl3Scope
	Subject Acl3Subject
}

// matchesScopeSubject checks if an ACL entry applies to the given scope+subject.
// An entry matches if:
// - scope and subject match exactly, OR
// - entry has scope "all" or subject "All" (global grant), OR
// - entry has scope "ror" with subject matching the requested scope or "Global"
func matchesScopeSubject(entry AclV3ListItem, scope Acl3Scope, subject Acl3Subject) bool {
	// Exact match
	if entry.Scope == scope && entry.Subject == subject {
		return true
	}
	// Global scope
	if entry.Scope == Acl3Scope(Acl2ScopeAll) || entry.Subject == Acl3Subject(Acl2RorSubjectAll) {
		return true
	}
	// "ror" scope with subject matching the requested scope or "Global"
	if entry.Scope == Acl3Scope(Acl2ScopeRor) {
		if entry.Subject == Acl3Subject(scope) || entry.Subject == Acl3Subject(Acl2RorSubjectGlobal) {
			return true
		}
	}
	return false
}
