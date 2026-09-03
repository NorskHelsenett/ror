package aclstore

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// SnapshotStore is an in-memory, read-optimized Store. It serves every read from
// an atomically-swapped snapshot of all ACL entries and delegates writes to a
// backend Store, reloading afterwards so the local process reflects its own
// changes. Reload is also meant to be driven externally (a change-notification
// subscriber and a periodic ticker) so the snapshot converges with writes made
// by other processes.
//
// The full ACL set is small (bounded, ~thousands of entries), so holding it in
// memory and rebuilding indexes on every reload is cheap.
type SnapshotStore struct {
	backend Store
	snap    atomic.Pointer[snapshot]
}

// scopeSubject is the composite key for the scope+subject index.
type scopeSubject struct {
	scope   aclscope.Scope
	subject aclscope.Subject
}

// snapshot is an immutable, indexed view of all ACL entries. A new snapshot is
// built on every reload and swapped in atomically; readers never mutate it.
type snapshot struct {
	all            aclmodels.AclV3List
	byID           map[string]aclmodels.AclV3ListItem
	byGroup        aclmodels.AclV3ListByGroup
	byScopeSubject map[scopeSubject][]aclmodels.AclV3ListItem
}

// compile-time assurance that SnapshotStore satisfies the Store interface.
var _ Store = (*SnapshotStore)(nil)

// NewSnapshotStore creates a snapshot store backed by the given Store. The
// snapshot starts empty; call Reload before serving reads.
func NewSnapshotStore(backend Store) *SnapshotStore {
	s := &SnapshotStore{backend: backend}
	s.snap.Store(newSnapshot(nil))
	return s
}

// newSnapshot builds the indexed view from a flat list of entries.
func newSnapshot(entries []aclmodels.AclV3ListItem) *snapshot {
	snap := &snapshot{
		all:            entries,
		byID:           make(map[string]aclmodels.AclV3ListItem, len(entries)),
		byGroup:        make(aclmodels.AclV3ListByGroup),
		byScopeSubject: make(map[scopeSubject][]aclmodels.AclV3ListItem),
	}
	for _, e := range entries {
		if e.Id != "" {
			snap.byID[e.Id] = e
		}
		snap.byGroup[e.Group] = append(snap.byGroup[e.Group], e)
		key := scopeSubject{scope: e.Scope, subject: e.Subject}
		snap.byScopeSubject[key] = append(snap.byScopeSubject[key], e)
	}
	return snap
}

// Reload rebuilds the snapshot from the backend and swaps it in atomically.
func (s *SnapshotStore) Reload(ctx context.Context) error {
	entries, err := s.backend.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("failed to reload ACL snapshot: %w", err)
	}
	s.snap.Store(newSnapshot(entries))
	return nil
}

// current returns the live snapshot for lock-free reads.
func (s *SnapshotStore) current() *snapshot {
	return s.snap.Load()
}

// --- Reader: served entirely from the in-memory snapshot ---

// GetAll returns a copy of every entry in the current snapshot.
func (s *SnapshotStore) GetAll(_ context.Context) (aclmodels.AclV3List, error) {
	all := s.current().all
	out := make([]aclmodels.AclV3ListItem, len(all))
	copy(out, all)
	return out, nil
}

// GetById returns the entry with the given id, or nil if absent.
func (s *SnapshotStore) GetById(_ context.Context, id string) (*aclmodels.AclV3ListItem, error) {
	entry, ok := s.current().byID[id]
	if !ok {
		return nil, nil
	}
	return &entry, nil
}

// GetByGroups returns entries for the given groups keyed by group; groups with
// no entries are omitted.
func (s *SnapshotStore) GetByGroups(_ context.Context, groups []string) (aclmodels.AclV3List, error) {
	snap := s.current()
	result := make(aclmodels.AclV3ListByGroup, len(groups))
	for _, g := range groups {
		if entries := snap.byGroup[g]; len(entries) > 0 {
			cp := make([]aclmodels.AclV3ListItem, len(entries))
			copy(cp, entries)
			result[g] = cp
		}
	}
	return result.Flatten(), nil
}

// GetByScopeSubject returns all entries for the given scope+subject pair.
func (s *SnapshotStore) GetByScopeSubject(_ context.Context, scope aclscope.Scope, subject aclscope.Subject) (aclmodels.AclV3List, error) {
	entries := s.current().byScopeSubject[scopeSubject{scope: scope, subject: subject}]
	out := make([]aclmodels.AclV3ListItem, len(entries))
	copy(out, entries)
	return out, nil
}

// --- Writer: delegated to the backend, then a best-effort local reload ---

// Create persists via the backend and refreshes the local snapshot.
func (s *SnapshotStore) Create(ctx context.Context, item aclmodels.AclV3ListItem) (*aclmodels.AclV3ListItem, error) {
	created, err := s.backend.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	s.reloadAfterWrite(ctx)
	return created, nil
}

// Update persists via the backend and refreshes the local snapshot.
func (s *SnapshotStore) Update(ctx context.Context, id string, item aclmodels.AclV3ListItem) (*aclmodels.AclV3ListItem, *aclmodels.AclV3ListItem, error) {
	updated, previous, err := s.backend.Update(ctx, id, item)
	if err != nil {
		return nil, nil, err
	}
	s.reloadAfterWrite(ctx)
	return updated, previous, nil
}

// Delete persists via the backend and refreshes the local snapshot.
func (s *SnapshotStore) Delete(ctx context.Context, id string) (*aclmodels.AclV3ListItem, error) {
	deleted, err := s.backend.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	s.reloadAfterWrite(ctx)
	return deleted, nil
}

// reloadAfterWrite refreshes the snapshot after a successful write. A failure
// here does not fail the write: the entry is already persisted, and the periodic
// reload / change-notification path will converge the snapshot.
func (s *SnapshotStore) reloadAfterWrite(ctx context.Context) {
	if err := s.Reload(ctx); err != nil {
		rlog.Warnc(ctx, "failed to refresh ACL snapshot after write", rlog.Any("error", err))
	}
}
