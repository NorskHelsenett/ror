package aclv3resolver

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

// AclV3Store is the interface for loading V3 ACL entries.
// Implementations can use MongoDB, Redis, or any other backend.
// The batch method (GetByGroups) allows efficient loading in a single round-trip.
type AclV3Store interface {
	// GetByGroups returns all V3 ACL entries for the given group names in a single call.
	// The result is a map from group name to the entries for that group.
	// This allows a single MongoDB $in query, or a Redis MGET + backfill pattern.
	GetByGroups(ctx context.Context, groups []string) (map[string][]aclmodels.AclV3ListItem, error)
}
