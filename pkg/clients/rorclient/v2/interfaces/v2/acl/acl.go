package acl

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

type AclInterface interface {
	// Lookup resolves the scope+subject pairs the caller has the given access
	// type for, using the V3 ACL backend. The optional scopes and subjects
	// narrow the result to the given scopes and/or subjects (uids); pass nil for
	// no filtering.
	Lookup(ctx context.Context, access aclmodels.AccessTypeV3, scopes []aclscope.Scope, subjects []aclscope.Subject) (*aclmodels.AclV3LookupResponse, error)

	// LookupByScopeSubject resolves the access groups the caller has for the
	// given scope+subject pair, using the V3 ACL backend. Unlike Lookup, scope
	// and subject are required and identify a single resource.
	LookupByScopeSubject(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject) (*aclmodels.Acl3LookupByScopeSubjectResponse, error)

	CheckAccess(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject, access aclmodels.AccessTypeV3) bool

	// Create persists a new V3 ACL entry via POST /v2/acl and returns the
	// server-assigned entry. Unlike the V1 client it speaks AclV3ListItem, so
	// V3-only capabilities (resource:*, ror:config:*, ...) are preserved.
	Create(ctx context.Context, item aclmodels.AclV3ListItem) (*aclmodels.AclV3ListItem, error)
	// Update replaces the V3 ACL entry with the given id via PUT /v2/acl/{id}.
	Update(ctx context.Context, id string, item aclmodels.AclV3ListItem) (*aclmodels.AclV3ListItem, error)
	// Delete removes the V3 ACL entry with the given id via DELETE /v2/acl/{id}.
	Delete(ctx context.Context, id string) error
	// GetById returns the V3 ACL entry with the given id via GET /v2/acl/{id}.
	GetById(ctx context.Context, id string) (*aclmodels.AclV3ListItem, error)
	// GetByFilter returns a page of V3 ACL entries via POST /v2/acl/filter.
	GetByFilter(ctx context.Context, filter apicontracts.Filter) (*apicontracts.PaginatedResult[aclmodels.AclV3ListItem], error)
	// GetAll pages through GetByFilter and returns every V3 ACL entry.
	GetAll(ctx context.Context) (*[]aclmodels.AclV3ListItem, error)
}
