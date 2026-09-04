package acl

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

type AclInterface interface {
	// Lookup resolves the scope+subject pairs the caller has the given access
	// type for, using the V3 ACL backend. The optional scopes and subjects
	// narrow the result to the given scopes and/or subjects (uids); pass nil for
	// no filtering.
	Lookup(ctx context.Context, access string, scopes []string, subjects []string) (*aclmodels.AclV3LookupResponse, error)

	// LookupByScopeSubject resolves the access groups the caller has for the
	// given scope+subject pair, using the V3 ACL backend. Unlike Lookup, scope
	// and subject are required and identify a single resource.
	LookupByScopeSubject(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject) (*aclmodels.Acl3LookupByScopeSubjectResponse, error)
}
