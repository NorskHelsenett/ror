package acl

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

// AncestorResolver resolves the ancestor ownerrefs of a resource by walking the
// ownerref chain in resourcesv2 upward (child -> parent). It is the inverse of
// ScopeExpander: where the expander finds a scope's descendants, this finds a
// resource's ancestors. No hardcoded hierarchy — the tree emerges from
// rormeta.ownerref data on each resource.
//
// Example: if Cluster "cluster-abc" has ownerref {Workspace, ws-dev} and
// Workspace "ws-dev" has ownerref {Project, proj-1}, then
// Ancestors(ctx, "KubernetesCluster", "cluster-abc") returns, nearest-first:
//
//	[{Workspace, ws-dev}, {Project, proj-1}]
//
// The queried scope+subject is NOT included, and an empty slice is returned when
// the resource has no ancestors (or does not exist).
type AncestorResolver interface {
	Ancestors(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject) ([]Ownerref, error)
}
