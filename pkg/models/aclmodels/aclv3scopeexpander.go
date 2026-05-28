package aclmodels

import "context"

// ScopeExpander resolves hierarchical scope relationships by walking the
// ownerref chain in resourcesv2. No business logic or hardcoded hierarchy —
// the tree emerges from rormeta.ownerref data on each resource.
//
// ExpandScope recursively finds all resourcesv2 resources whose ownerref
// matches {scope, subject}, then their children, and so on.
//
// Example: if Workspace "ws-dev" has ownerref {Project, proj-1} and
// Cluster "cluster-abc" has ownerref {Workspace, ws-dev}, then
// ExpandScope(ctx, "Project", "proj-1") returns:
//
//	[{Workspace, ws-dev}, {KubernetesCluster, cluster-abc}]
//
// Returns nil if no resources have the given ownerref (leaf scope).
// The original scope+subject is NOT included in the result.
type ScopeExpander interface {
	ExpandScope(ctx context.Context, scope Acl3Scope, subject Acl3Subject) ([]AclV3Ownerref, error)
}
