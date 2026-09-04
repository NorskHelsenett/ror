package aclstore

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// NotifyingStore decorates a Store, publishing an ACL-change signal after every
// successful write so other processes refresh their snapshots. Reads pass
// through unchanged.
type NotifyingStore struct {
	inner     Store
	publisher ChangePublisher
}

// compile-time assurance that NotifyingStore satisfies the Store interface.
var _ Store = (*NotifyingStore)(nil)

// NewNotifyingStore wraps inner so successful writes publish a change signal. A
// nil publisher is treated as a no-op.
func NewNotifyingStore(inner Store, publisher ChangePublisher) *NotifyingStore {
	if publisher == nil {
		publisher = NoopPublisher{}
	}
	return &NotifyingStore{inner: inner, publisher: publisher}
}

// --- Reader: pass-through ---

func (s *NotifyingStore) GetAll(ctx context.Context) (aclmodels.AclV3List, error) {
	return s.inner.GetAll(ctx)
}

func (s *NotifyingStore) GetById(ctx context.Context, id string) (*aclmodels.AclV3ListItem, error) {
	return s.inner.GetById(ctx, id)
}

func (s *NotifyingStore) GetByGroups(ctx context.Context, groups []string) (aclmodels.AclV3List, error) {
	return s.inner.GetByGroups(ctx, groups)
}

func (s *NotifyingStore) GetByScopeSubject(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject) (aclmodels.AclV3List, error) {
	return s.inner.GetByScopeSubject(ctx, scope, subject)
}

// --- Writer: delegate, then publish on success ---

func (s *NotifyingStore) Create(ctx context.Context, item aclmodels.AclV3ListItem) (*aclmodels.AclV3ListItem, error) {
	created, err := s.inner.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	s.publish(ctx)
	return created, nil
}

func (s *NotifyingStore) Update(ctx context.Context, id string, item aclmodels.AclV3ListItem) (*aclmodels.AclV3ListItem, *aclmodels.AclV3ListItem, error) {
	updated, previous, err := s.inner.Update(ctx, id, item)
	if err != nil {
		return nil, nil, err
	}
	s.publish(ctx)
	return updated, previous, nil
}

func (s *NotifyingStore) Delete(ctx context.Context, id string) (*aclmodels.AclV3ListItem, error) {
	deleted, err := s.inner.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	s.publish(ctx)
	return deleted, nil
}

// publish broadcasts a change signal. A failure is logged but never fails the
// write: the change is already persisted and the periodic reload will converge
// other processes.
func (s *NotifyingStore) publish(ctx context.Context) {
	if err := s.publisher.PublishChange(ctx); err != nil {
		rlog.Warnc(ctx, "failed to publish ACL change signal", rlog.Any("error", err))
	}
}
