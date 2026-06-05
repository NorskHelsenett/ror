package acl

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

// Store is the interface for loading ACL entries.
// Implementations can use MongoDB, Redis, or any other backend.
// Both methods load all entries (V2 and V3) in a single round-trip and convert
// to the requested format. V2 entries are converted to V3 using the v3 struct
// tags on AclV2ListItemAccess. V3 entries are converted to V2 with capabilities
// that have no V2 equivalent silently dropped.
type Store interface {
	// GetByGroups returns all ACL entries for the given group names as V3 items.
	// V2 entries in the database are converted to V3 using aclmodels.V2ToV3.
	GetByGroups(ctx context.Context, groups []string) (map[string][]aclmodels.AclV3ListItem, error)

	// GetV2ByGroups returns all ACL entries for the given group names as V2 items.
	// V3 entries in the database are converted to V2 using aclmodels.V3ToV2.
	// V3-only capabilities (e.g. "kubernetes:admin") are dropped in the conversion.
	GetV2ByGroups(ctx context.Context, groups []string) (map[string][]aclmodels.AclV2ListItem, error)
}
