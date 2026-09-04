package aclstore

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakeBackend is an in-memory Store used to drive SnapshotStore tests without a
// database. It records how often GetAll is called and can be made to fail.
type fakeBackend struct {
	mu          sync.Mutex
	entries     []aclmodels.AclV3ListItem
	getAllCalls int
	getAllErr   error
	createErr   error
	updateErr   error
	deleteErr   error
	nextID      int
}

func (f *fakeBackend) setEntries(entries ...aclmodels.AclV3ListItem) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.entries = entries
}

func (f *fakeBackend) GetAll(_ context.Context) (aclmodels.AclV3List, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.getAllCalls++
	if f.getAllErr != nil {
		return nil, f.getAllErr
	}
	out := make([]aclmodels.AclV3ListItem, len(f.entries))
	copy(out, f.entries)
	return out, nil
}

func (f *fakeBackend) GetById(_ context.Context, id string) (*aclmodels.AclV3ListItem, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	for i := range f.entries {
		if f.entries[i].Id == id {
			e := f.entries[i]
			return &e, nil
		}
	}
	return nil, nil
}

func (f *fakeBackend) GetByGroups(_ context.Context, groups []string) (aclmodels.AclV3List, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	want := make(map[string]bool, len(groups))
	for _, g := range groups {
		want[g] = true
	}
	var result aclmodels.AclV3List
	for _, e := range f.entries {
		if want[e.Group] {
			result = append(result, e)
		}
	}
	return result, nil
}

// countGroup counts the entries in a flat list that belong to the given group.
func countGroup(l aclmodels.AclV3List, group string) int {
	n := 0
	for _, e := range l {
		if e.Group == group {
			n++
		}
	}
	return n
}

func (f *fakeBackend) GetByScopeSubject(_ context.Context, scope aclscope.Scope, subject aclscope.Subject) (aclmodels.AclV3List, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	var result []aclmodels.AclV3ListItem
	for _, e := range f.entries {
		if e.Scope == scope && e.Subject == subject {
			result = append(result, e)
		}
	}
	return result, nil
}

func (f *fakeBackend) Create(_ context.Context, item aclmodels.AclV3ListItem) (*aclmodels.AclV3ListItem, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.createErr != nil {
		return nil, f.createErr
	}
	f.nextID++
	item.Id = fmt.Sprintf("id-%d", f.nextID)
	item.Version = 3
	f.entries = append(f.entries, item)
	return &item, nil
}

func (f *fakeBackend) Update(_ context.Context, id string, item aclmodels.AclV3ListItem) (*aclmodels.AclV3ListItem, *aclmodels.AclV3ListItem, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.updateErr != nil {
		return nil, nil, f.updateErr
	}
	for i := range f.entries {
		if f.entries[i].Id == id {
			previous := f.entries[i]
			item.Id = id
			item.Version = 3
			f.entries[i] = item
			return &item, &previous, nil
		}
	}
	return nil, nil, fmt.Errorf("entry %q not found", id)
}

func (f *fakeBackend) Delete(_ context.Context, id string) (*aclmodels.AclV3ListItem, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.deleteErr != nil {
		return nil, f.deleteErr
	}
	for i := range f.entries {
		if f.entries[i].Id == id {
			deleted := f.entries[i]
			f.entries = append(f.entries[:i], f.entries[i+1:]...)
			return &deleted, nil
		}
	}
	return nil, fmt.Errorf("entry %q not found", id)
}

func (f *fakeBackend) callCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.getAllCalls
}

func entry(id, group string, scope aclscope.Scope, subject aclscope.Subject, access ...aclmodels.AccessTypeV3) aclmodels.AclV3ListItem {
	return aclmodels.AclV3ListItem{
		Id:      id,
		Version: 3,
		Group:   group,
		Scope:   scope,
		Subject: subject,
		Access:  access,
	}
}

// --- Snapshot lifecycle ---

func TestSnapshotStore_EmptyBeforeReload(t *testing.T) {
	store := NewSnapshotStore(&fakeBackend{})

	all, err := store.GetAll(context.Background())
	require.NoError(t, err)
	assert.Empty(t, all)
}

func TestSnapshotStore_ReloadLoadsFromBackend(t *testing.T) {
	backend := &fakeBackend{}
	backend.setEntries(
		entry("1", "team-a", aclscope.ScopeCluster, "c1", "ror:read"),
		entry("2", "team-b", aclscope.ScopeProject, "p1", "ror:read"),
	)
	store := NewSnapshotStore(backend)

	require.NoError(t, store.Reload(context.Background()))

	all, err := store.GetAll(context.Background())
	require.NoError(t, err)
	assert.Len(t, all, 2)
}

// Reads must be served from the snapshot, never the backend: after a reload,
// mutating the backend must not change reads until the next reload.
func TestSnapshotStore_ReadsServedFromSnapshot(t *testing.T) {
	backend := &fakeBackend{}
	backend.setEntries(entry("1", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	store := NewSnapshotStore(backend)
	require.NoError(t, store.Reload(context.Background()))

	callsAfterReload := backend.callCount()

	// Change the backend behind the store's back.
	backend.setEntries(
		entry("1", "team-a", aclscope.ScopeCluster, "c1", "ror:read"),
		entry("2", "team-a", aclscope.ScopeCluster, "c2", "ror:read"),
	)

	// Several reads should not hit the backend and should reflect the old snapshot.
	for i := 0; i < 3; i++ {
		got, err := store.GetByGroups(context.Background(), []string{"team-a"})
		require.NoError(t, err)
		assert.Len(t, got, 1, "reads must reflect the loaded snapshot, not the backend")
	}
	assert.Equal(t, callsAfterReload, backend.callCount(), "reads must not call the backend")

	// After an explicit reload the new state is visible.
	require.NoError(t, store.Reload(context.Background()))
	got, err := store.GetByGroups(context.Background(), []string{"team-a"})
	require.NoError(t, err)
	assert.Len(t, got, 2)
}

// A failed reload must not clobber the existing snapshot.
func TestSnapshotStore_ReloadErrorKeepsPreviousSnapshot(t *testing.T) {
	backend := &fakeBackend{}
	backend.setEntries(entry("1", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	store := NewSnapshotStore(backend)
	require.NoError(t, store.Reload(context.Background()))

	backend.mu.Lock()
	backend.getAllErr = fmt.Errorf("db down")
	backend.mu.Unlock()

	err := store.Reload(context.Background())
	require.Error(t, err)

	all, err := store.GetAll(context.Background())
	require.NoError(t, err)
	assert.Len(t, all, 1, "previous snapshot must survive a failed reload")
}

// --- Read access patterns ---

func TestSnapshotStore_GetById(t *testing.T) {
	backend := &fakeBackend{}
	backend.setEntries(entry("42", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	store := NewSnapshotStore(backend)
	require.NoError(t, store.Reload(context.Background()))

	got, err := store.GetById(context.Background(), "42")
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, "team-a", got.Group)

	missing, err := store.GetById(context.Background(), "nope")
	require.NoError(t, err)
	assert.Nil(t, missing)
}

func TestSnapshotStore_GetByGroups(t *testing.T) {
	backend := &fakeBackend{}
	backend.setEntries(
		entry("1", "team-a", aclscope.ScopeCluster, "c1", "ror:read"),
		entry("2", "team-a", aclscope.ScopeCluster, "c2", "ror:read"),
		entry("3", "team-b", aclscope.ScopeProject, "p1", "ror:read"),
	)
	store := NewSnapshotStore(backend)
	require.NoError(t, store.Reload(context.Background()))

	byA, err := store.GetByGroups(context.Background(), []string{"team-a"})
	require.NoError(t, err)
	assert.Len(t, byA, 2)
	assert.Equal(t, 0, countGroup(byA, "team-b"), "unrequested group must be absent")

	both, err := store.GetByGroups(context.Background(), []string{"team-a", "team-b"})
	require.NoError(t, err)
	assert.Equal(t, 2, countGroup(both, "team-a"))
	assert.Equal(t, 1, countGroup(both, "team-b"))
}

func TestSnapshotStore_GetByScopeSubject(t *testing.T) {
	backend := &fakeBackend{}
	backend.setEntries(
		entry("1", "team-a", aclscope.ScopeCluster, "c1", "ror:read"),
		entry("2", "team-b", aclscope.ScopeCluster, "c1", "ror:write"),
		entry("3", "team-a", aclscope.ScopeCluster, "c2", "ror:read"),
	)
	store := NewSnapshotStore(backend)
	require.NoError(t, store.Reload(context.Background()))

	got, err := store.GetByScopeSubject(context.Background(), aclscope.ScopeCluster, "c1")
	require.NoError(t, err)
	assert.Len(t, got, 2)

	none, err := store.GetByScopeSubject(context.Background(), aclscope.ScopeCluster, "absent")
	require.NoError(t, err)
	assert.Empty(t, none)
}

// Returned slices must be copies; mutating them must not corrupt the snapshot.
func TestSnapshotStore_ReadsReturnCopies(t *testing.T) {
	backend := &fakeBackend{}
	backend.setEntries(
		entry("1", "team-a", aclscope.ScopeCluster, "c1", "ror:read"),
		entry("2", "team-a", aclscope.ScopeCluster, "c2", "ror:read"),
	)
	store := NewSnapshotStore(backend)
	require.NoError(t, store.Reload(context.Background()))

	got, err := store.GetByGroups(context.Background(), []string{"team-a"})
	require.NoError(t, err)
	got[0] = entry("x", "hacked", aclscope.ScopeCluster, "x", "ror:read")

	again, err := store.GetByGroups(context.Background(), []string{"team-a"})
	require.NoError(t, err)
	assert.NotEqual(t, "hacked", again[0].Group, "snapshot must be unaffected by caller mutation")
}

// --- Writes: delegate + local reload ---

func TestSnapshotStore_CreateDelegatesAndRefreshes(t *testing.T) {
	backend := &fakeBackend{}
	store := NewSnapshotStore(backend)

	created, err := store.Create(context.Background(), entry("", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.NotEmpty(t, created.Id)

	// The write must be visible via the snapshot without an explicit reload.
	got, err := store.GetById(context.Background(), created.Id)
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, "team-a", got.Group)
}

func TestSnapshotStore_UpdateDelegatesAndRefreshes(t *testing.T) {
	backend := &fakeBackend{}
	store := NewSnapshotStore(backend)
	created, err := store.Create(context.Background(), entry("", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	require.NoError(t, err)

	updated, previous, err := store.Update(context.Background(), created.Id,
		entry("", "team-b", aclscope.ScopeCluster, "c1", "ror:read"))
	require.NoError(t, err)
	assert.Equal(t, "team-a", previous.Group)
	assert.Equal(t, "team-b", updated.Group)

	// Snapshot reflects the moved group.
	old, err := store.GetByGroups(context.Background(), []string{"team-a"})
	require.NoError(t, err)
	assert.Empty(t, old)
	moved, err := store.GetByGroups(context.Background(), []string{"team-b"})
	require.NoError(t, err)
	assert.Len(t, moved, 1)
}

func TestSnapshotStore_DeleteDelegatesAndRefreshes(t *testing.T) {
	backend := &fakeBackend{}
	store := NewSnapshotStore(backend)
	created, err := store.Create(context.Background(), entry("", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	require.NoError(t, err)

	deleted, err := store.Delete(context.Background(), created.Id)
	require.NoError(t, err)
	assert.Equal(t, created.Id, deleted.Id)

	got, err := store.GetById(context.Background(), created.Id)
	require.NoError(t, err)
	assert.Nil(t, got, "deleted entry must be gone from the snapshot")
}

// A backend write error must be surfaced and must not change the snapshot.
func TestSnapshotStore_WriteErrorLeavesSnapshotUnchanged(t *testing.T) {
	backend := &fakeBackend{createErr: fmt.Errorf("insert failed")}
	backend.setEntries(entry("1", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	store := NewSnapshotStore(backend)
	require.NoError(t, store.Reload(context.Background()))

	_, err := store.Create(context.Background(), entry("", "team-b", aclscope.ScopeCluster, "c2", "ror:read"))
	require.Error(t, err)

	all, err := store.GetAll(context.Background())
	require.NoError(t, err)
	assert.Len(t, all, 1, "failed write must not alter the snapshot")
}

// Exercise concurrent reads against reloads; run with -race to catch data races.
func TestSnapshotStore_ConcurrentReadsAndReloads(t *testing.T) {
	backend := &fakeBackend{}
	backend.setEntries(entry("1", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	store := NewSnapshotStore(backend)
	require.NoError(t, store.Reload(context.Background()))

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 200; j++ {
				_, _ = store.GetByGroups(context.Background(), []string{"team-a"})
			}
		}()
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				_ = store.Reload(context.Background())
			}
		}()
	}
	wg.Wait()
}
