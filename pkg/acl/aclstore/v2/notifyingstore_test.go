package aclstore

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakePublisher records PublishChange calls and can be made to fail.
type fakePublisher struct {
	mu    sync.Mutex
	calls int
	err   error
}

func (p *fakePublisher) PublishChange(context.Context) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.calls++
	return p.err
}

func (p *fakePublisher) count() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.calls
}

func TestNotifyingStore_CreatePublishes(t *testing.T) {
	pub := &fakePublisher{}
	store := NewNotifyingStore(&fakeBackend{}, pub)

	_, err := store.Create(context.Background(), entry("", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	require.NoError(t, err)
	assert.Equal(t, 1, pub.count())
}

func TestNotifyingStore_UpdatePublishes(t *testing.T) {
	backend := &fakeBackend{}
	backend.setEntries(entry("1", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	pub := &fakePublisher{}
	store := NewNotifyingStore(backend, pub)

	_, _, err := store.Update(context.Background(), "1", entry("", "team-b", aclscope.ScopeCluster, "c1", "ror:read"))
	require.NoError(t, err)
	assert.Equal(t, 1, pub.count())
}

func TestNotifyingStore_DeletePublishes(t *testing.T) {
	backend := &fakeBackend{}
	backend.setEntries(entry("1", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	pub := &fakePublisher{}
	store := NewNotifyingStore(backend, pub)

	_, err := store.Delete(context.Background(), "1")
	require.NoError(t, err)
	assert.Equal(t, 1, pub.count())
}

// A failed write must not publish a change signal.
func TestNotifyingStore_WriteErrorDoesNotPublish(t *testing.T) {
	pub := &fakePublisher{}
	store := NewNotifyingStore(&fakeBackend{createErr: fmt.Errorf("insert failed")}, pub)

	_, err := store.Create(context.Background(), entry("", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	require.Error(t, err)
	assert.Equal(t, 0, pub.count())
}

// A publish failure must not fail the write.
func TestNotifyingStore_PublishErrorDoesNotFailWrite(t *testing.T) {
	pub := &fakePublisher{err: fmt.Errorf("bus down")}
	store := NewNotifyingStore(&fakeBackend{}, pub)

	created, err := store.Create(context.Background(), entry("", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.Equal(t, 1, pub.count())
}

// Reads must pass through and never publish.
func TestNotifyingStore_ReadsPassThroughWithoutPublishing(t *testing.T) {
	backend := &fakeBackend{}
	backend.setEntries(entry("1", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	pub := &fakePublisher{}
	store := NewNotifyingStore(backend, pub)

	got, err := store.GetByGroups(context.Background(), []string{"team-a"})
	require.NoError(t, err)
	assert.Len(t, got, 1)
	assert.Equal(t, 0, pub.count())
}

// A nil publisher is treated as a no-op and must not panic.
func TestNotifyingStore_NilPublisherIsNoop(t *testing.T) {
	store := NewNotifyingStore(&fakeBackend{}, nil)

	_, err := store.Create(context.Background(), entry("", "team-a", aclscope.ScopeCluster, "c1", "ror:read"))
	require.NoError(t, err)
}
