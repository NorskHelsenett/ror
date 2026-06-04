package aclv3resolver

import (
	"context"
	"slices"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

// AclV3Ownerref represents a scope+subject pair that a user has access to.
type AclV3Ownerref struct {
	Scope   aclscope.Scope
	Subject aclscope.Subject
}

// AclV3Resolver resolves access for a set of groups using an AclV3Store.
// It loads all entries in one batch call, then compiles access in-memory.
// An optional ScopeExpander enables hierarchical scope resolution:
// if a user has access to a Project, the expander resolves all descendant
// ownerrefs (Workspaces, Clusters, etc.) so they are included in the result.
type AclV3Resolver struct {
	store    AclV3Store
	expander ScopeExpander
}

// NewAclV3Resolver creates a new resolver with the given store backend.
// Use WithScopeExpander to enable hierarchical scope resolution.
func NewAclV3Resolver(store AclV3Store, opts ...AclV3ResolverOption) *AclV3Resolver {
	r := &AclV3Resolver{store: store}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// AclV3ResolverOption configures an AclV3Resolver.
type AclV3ResolverOption func(*AclV3Resolver)

// WithScopeExpander enables hierarchical scope resolution.
func WithScopeExpander(expander ScopeExpander) AclV3ResolverOption {
	return func(r *AclV3Resolver) {
		r.expander = expander
	}
}

// ResolveAccess loads ACL entries for all given groups (single round-trip) and returns
// the union of access types that match the given scope and subject.
func (r *AclV3Resolver) ResolveAccess(ctx context.Context, groups []string, scope aclscope.Scope, subject aclscope.Subject) ([]aclmodels.AccessTypeV3, error) {
	entriesByGroup, err := r.store.GetByGroups(ctx, groups)
	if err != nil {
		return nil, err
	}

	seen := make(map[aclmodels.AccessTypeV3]struct{})
	for _, entries := range entriesByGroup {
		for _, entry := range entries {
			if matchesScopeSubject(entry, scope, subject) {
				for _, a := range entry.Access {
					seen[a] = struct{}{}
				}
			}
		}
	}

	result := make([]aclmodels.AccessTypeV3, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}
	return result, nil
}

// ResolveOwnerrefs returns all scope+subject pairs the groups have access to for the given access type.
// Returns nil (meaning unrestricted) if any entry grants global/all access.
// When a ScopeExpander is configured, non-leaf scopes (e.g. Project, Workspace) are expanded
// to include all descendant ownerrefs alongside the original entry.
func (r *AclV3Resolver) ResolveOwnerrefs(ctx context.Context, groups []string, requiredAccess aclmodels.AccessTypeV3) ([]AclV3Ownerref, error) {
	entriesByGroup, err := r.store.GetByGroups(ctx, groups)
	if err != nil {
		return nil, err
	}

	refs := make([]AclV3Ownerref, 0)
	seen := make(map[AclV3Ownerref]struct{})

	addRef := func(ref AclV3Ownerref) {
		if _, ok := seen[ref]; !ok {
			seen[ref] = struct{}{}
			refs = append(refs, ref)
		}
	}

	for _, entries := range entriesByGroup {
		for _, entry := range entries {
			if !slices.Contains(entry.Access, requiredAccess) {
				continue
			}
			// Global access — unrestricted
			if entry.Scope == aclscope.ScopeAll || entry.Subject == aclscope.SubjectAll {
				return nil, nil
			}

			ref := AclV3Ownerref{Scope: entry.Scope, Subject: entry.Subject}
			addRef(ref)

			// Expand to descendants if expander is available
			if r.expander != nil {
				descendants, err := r.expander.ExpandScope(ctx, entry.Scope, entry.Subject)
				if err != nil {
					return nil, err
				}
				for _, d := range descendants {
					addRef(d)
				}
			}
		}
	}
	return refs, nil
}

// HasAccess checks if the groups have a specific access type for the given scope+subject.
func (r *AclV3Resolver) HasAccess(ctx context.Context, groups []string, scope aclscope.Scope, subject aclscope.Subject, requiredAccess aclmodels.AccessTypeV3) (bool, error) {
	access, err := r.ResolveAccess(ctx, groups, scope, subject)
	if err != nil {
		return false, err
	}
	return slices.Contains(access, requiredAccess), nil
}

// matchesScopeSubject checks if an ACL entry applies to the given scope+subject.
// An entry matches if:
// - scope and subject match exactly, OR
// - entry has scope "all" or subject "All" (global grant), OR
// - entry has scope "ror" with subject matching the requested scope or "Global"
func matchesScopeSubject(entry aclmodels.AclV3ListItem, scope aclscope.Scope, subject aclscope.Subject) bool {
	// Exact match
	if entry.Scope == scope && entry.Subject == subject {
		return true
	}
	// Global scope
	if entry.Scope == aclscope.ScopeAll || entry.Subject == aclscope.SubjectAll {
		return true
	}
	// "ror" scope with subject matching the requested scope or "Global"
	if entry.Scope == aclscope.ScopeRor {
		if entry.Subject == aclscope.Subject(scope) || entry.Subject == aclscope.SubjectGlobal {
			return true
		}
	}
	return false
}
