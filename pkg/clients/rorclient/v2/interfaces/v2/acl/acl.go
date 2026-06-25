package acl

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

type AclInterface interface {
	// Lookup resolves the scope+subject pairs the caller has the given access
	// type for, using the V3 ACL backend. The optional scopes and subjects
	// narrow the result to the given scopes and/or subjects (uids); pass nil for
	// no filtering.
	Lookup(ctx context.Context, access string, scopes []string, subjects []string) (*aclmodels.AclV3LookupResponse, error)
}
