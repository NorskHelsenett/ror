package acl

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

type AclInterface interface {
	Create(ctx context.Context, item aclmodels.AclV2ListItem) error
	Update(ctx context.Context, id string, item aclmodels.AclV2ListItem) error
	Delete(ctx context.Context, id string) error
	CheckAccess(ctx context.Context, scope, subject, access string) bool
	GetById(ctx context.Context, id string) (*aclmodels.AclV2ListItem, error)
	GetByFilter(ctx context.Context, filter apicontracts.Filter) (*apicontracts.PaginatedResult[aclmodels.AclV2ListItem], error)
}
