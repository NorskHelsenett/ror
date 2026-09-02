package aclstore

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

// Store is the single entry point for all ACL persistence: reads, writes and
// bulk load. V3 is the canonical representation used throughout; implementations
// own any storage-level version conversion so callers only deal with V3.
type Store interface {
	Reader
	Writer
}

// Reader exposes every ACL read access pattern, all served from the same
// underlying data set in the canonical V3 form.
type Reader interface {
	// GetByGroups returns ACL entries for the given groups as V3 items,
	// keyed by group. Groups with no entries are omitted from the map.
	GetByGroups(ctx context.Context, groups []string) (aclmodels.AclV3List, error)

	// GetByScopeSubject returns all ACL entries for the scope+subject pair as V3 items.
	GetByScopeSubject(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject) (aclmodels.AclV3List, error)

	// GetById returns a single ACL entry by its id as a V3 item, or nil if not found.
	GetById(ctx context.Context, id string) (*aclmodels.AclV3ListItem, error)

	// GetAll returns every ACL entry as V3 items. It is the bulk-load primitive
	// used to (re)build an in-memory snapshot; callers should not use it on a hot path.
	GetAll(ctx context.Context) (aclmodels.AclV3List, error)
}

// Writer mutates ACL entries. All writes are persisted in the canonical V3
// form. Implementations are responsible for triggering any cache invalidation
// or change notification after a write commits.
type Writer interface {
	// Create persists a new ACL entry and returns the stored item.
	Create(ctx context.Context, item aclmodels.AclV3ListItem) (*aclmodels.AclV3ListItem, error)

	// Update replaces the entry with the given id. It returns the updated item
	// and the previous item (the latter enabling callers to react to a changed
	// group/scope/subject, e.g. for invalidation).
	Update(ctx context.Context, id string, item aclmodels.AclV3ListItem) (updated *aclmodels.AclV3ListItem, previous *aclmodels.AclV3ListItem, err error)

	// Delete removes the entry with the given id and returns the deleted item.
	Delete(ctx context.Context, id string) (deleted *aclmodels.AclV3ListItem, err error)
}
