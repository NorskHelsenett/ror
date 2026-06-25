package acl

import (
	"context"
	"slices"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/NorskHelsenett/ror/pkg/telemetry/rortracer"

	"go.opentelemetry.io/otel/attribute"
)

// Ownerref represents a scope+subject pair that a user has access to.
type Ownerref struct {
	Scope   aclscope.Scope
	Subject aclscope.Subject
}

// OwnerrefFilter narrows the ownerrefs returned by ResolveOwnerrefs.
// Each dimension is optional: an empty slice means "no restriction" for that
// dimension. When Scopes is non-empty, only ownerrefs whose scope is in the set
// are returned. When Subjects is non-empty, only ownerrefs whose subject is in
// the set are returned. Both must match when both are set.
type OwnerrefFilter struct {
	Scopes   []aclscope.Scope
	Subjects []aclscope.Subject
}

// IsEmpty reports whether the filter imposes no restriction.
func (f OwnerrefFilter) IsEmpty() bool {
	return len(f.Scopes) == 0 && len(f.Subjects) == 0
}

// Matches reports whether the given ownerref passes the filter.
func (f OwnerrefFilter) Matches(ref Ownerref) bool {
	if len(f.Scopes) > 0 && !slices.Contains(f.Scopes, ref.Scope) {
		return false
	}
	if len(f.Subjects) > 0 && !slices.Contains(f.Subjects, ref.Subject) {
		return false
	}
	return true
}

// skipExpansion reports whether descendant expansion of an entry with the given
// scope can be skipped without affecting the filtered result.
//
// Descendants are strictly lower in the ownership tree than their seed, so a
// descendant can never share the seed's scope (a kind does not contain another
// resource of the same kind in the RoR model). Therefore, when the filter
// restricts results to exactly the entry's own scope and imposes no subject
// restriction, no descendant can match and the (potentially expensive) expansion
// is pure overhead — e.g. a kubernetes:logon lookup filtered to KubernetesCluster
// never needs to walk into each cluster's resources.
func (f OwnerrefFilter) skipExpansion(entryScope aclscope.Scope) bool {
	if len(f.Subjects) > 0 || len(f.Scopes) == 0 {
		return false
	}
	for _, s := range f.Scopes {
		if s != entryScope {
			return false
		}
	}
	return true
}

// Resolver resolves access for a set of groups using a Store.
// It loads all entries in one batch call, then compiles access in-memory.
// An optional ScopeExpander enables hierarchical scope resolution:
// if a user has access to a Project, the expander resolves all descendant
// ownerrefs (Workspaces, Clusters, etc.) so they are included in the result.
type Resolver struct {
	store    Store
	expander ScopeExpander
}

// NewResolver creates a new resolver with the given store backend.
// Use WithScopeExpander to enable hierarchical scope resolution.
func NewResolver(store Store, opts ...ResolverOption) *Resolver {
	r := &Resolver{store: store}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// ResolverOption configures a Resolver.
type ResolverOption func(*Resolver)

// WithScopeExpander enables hierarchical scope resolution.
func WithScopeExpander(expander ScopeExpander) ResolverOption {
	return func(r *Resolver) {
		r.expander = expander
	}
}

// ResolveAccess loads ACL entries for all given groups (single round-trip) and returns
// the union of access types that match the given scope and subject.
func (r *Resolver) ResolveAccess(ctx context.Context, groups []string, scope aclscope.Scope, subject aclscope.Subject) ([]aclmodels.AccessTypeV3, error) {
	ctx, span := rortracer.StartSpan(ctx, "acl.Resolver.ResolveAccess")
	defer span.End()
	span.SetAttributes(
		attribute.Int("acl.groups", len(groups)),
		attribute.String("acl.scope", string(scope)),
		attribute.String("acl.subject", string(subject)),
	)

	entriesByGroup, err := r.store.GetByGroups(ctx, groups)
	if err != nil {
		return nil, rortracer.SpanError(span, err)
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
	span.SetAttributes(attribute.Int("acl.access_types", len(result)))
	return result, nil
}

// ResolveOwnerrefs returns all scope+subject pairs the groups have access to for the given access type.
// Returns nil (meaning unrestricted) if any entry grants global/all access.
// When a ScopeExpander is configured, non-leaf scopes (e.g. Project, Workspace) are expanded
// to include all descendant ownerrefs alongside the original entry.
//
// The optional filter narrows the result to specific scopes and/or subjects.
// An empty filter returns every resolved ownerref.
func (r *Resolver) ResolveOwnerrefs(ctx context.Context, groups []string, requiredAccess aclmodels.AccessTypeV3, filter OwnerrefFilter) ([]Ownerref, error) {
	ctx, span := rortracer.StartSpan(ctx, "acl.Resolver.ResolveOwnerrefs")
	defer span.End()
	span.SetAttributes(
		attribute.Int("acl.groups", len(groups)),
		attribute.String("acl.required_access", string(requiredAccess)),
		attribute.Bool("acl.expander_enabled", r.expander != nil),
		attribute.Int("acl.filter_scopes", len(filter.Scopes)),
		attribute.Int("acl.filter_subjects", len(filter.Subjects)),
	)

	entriesByGroup, err := r.store.GetByGroups(ctx, groups)
	if err != nil {
		return nil, rortracer.SpanError(span, err)
	}

	refs := make([]Ownerref, 0)
	seen := make(map[Ownerref]struct{})

	addRef := func(ref Ownerref) {
		if !filter.Matches(ref) {
			return
		}
		if _, ok := seen[ref]; !ok {
			seen[ref] = struct{}{}
			refs = append(refs, ref)
		}
	}

	// Collect the seeds to expand. Expansion is skipped for an entry when the
	// filter can only be satisfied by the entry's own scope (see skipExpansion):
	// descendants are strictly lower in the tree, so they can never match.
	var seeds []Ownerref
	seedSeen := make(map[Ownerref]struct{})

	for _, entries := range entriesByGroup {
		for _, entry := range entries {
			if !slices.Contains(entry.Access, requiredAccess) {
				continue
			}
			// Global access — unrestricted. Mirrors the global semantics in
			// matchesScopeSubject: the "all" scope or subject, or the "ror" scope
			// with the global subject.
			if entry.Scope == aclscope.ScopeAll || entry.Subject == aclscope.SubjectAll ||
				(entry.Scope == aclscope.ScopeRor && entry.Subject == aclscope.SubjectGlobal) {
				span.SetAttributes(attribute.Bool("acl.unrestricted", true))
				return nil, nil
			}

			ref := Ownerref{Scope: entry.Scope, Subject: entry.Subject}
			addRef(ref)

			if r.expander != nil && !filter.skipExpansion(entry.Scope) {
				if _, dup := seedSeen[ref]; !dup {
					seedSeen[ref] = struct{}{}
					seeds = append(seeds, ref)
				}
			}
		}
	}

	// Expand all seeds in a single batched round-trip and add the descendants.
	if r.expander != nil && len(seeds) > 0 {
		expanded, err := r.expander.ExpandScopes(ctx, seeds)
		if err != nil {
			return nil, rortracer.SpanError(span, err)
		}
		for _, descendants := range expanded {
			for _, d := range descendants {
				addRef(d)
			}
		}
	}

	span.SetAttributes(
		attribute.Int("acl.expanded_seeds", len(seeds)),
		attribute.Int("acl.ownerrefs", len(refs)),
	)
	return refs, nil
}

// HasAccess checks if the groups have a specific access type for the given scope+subject.
func (r *Resolver) HasAccess(ctx context.Context, groups []string, scope aclscope.Scope, subject aclscope.Subject, requiredAccess aclmodels.AccessTypeV3) (bool, error) {
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
