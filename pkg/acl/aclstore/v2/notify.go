package aclstore

import "context"

// ChangePublisher broadcasts an ACL-change signal so other processes refresh
// their snapshots. Implementations are best-effort: a publish failure must not
// fail the originating write.
type ChangePublisher interface {
	PublishChange(ctx context.Context) error
}

// NoopPublisher is a ChangePublisher that does nothing. Use it when no change
// bus is configured (single instance, or tests).
type NoopPublisher struct{}

// PublishChange does nothing and never fails.
func (NoopPublisher) PublishChange(context.Context) error { return nil }
