package rorclient

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/acl"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

type AclClient struct {
	Transport acl.AclInterface
}

func NewAclClient(transport acl.AclInterface) AclClient {
	return AclClient{Transport: transport}
}

func (c *AclClient) Create(ctx context.Context, item aclmodels.AclV2ListItem) error {
	return c.Transport.Create(ctx, item)
}

func (c *AclClient) Delete(ctx context.Context, id string) error {
	return c.Transport.Delete(ctx, id)
}

func (c *AclClient) GetById(ctx context.Context, id string) (*aclmodels.AclV2ListItem, error) {
	return c.Transport.GetById(ctx, id)
}

func (c *AclClient) GetByFilter(ctx context.Context, filter apicontracts.Filter) (*apicontracts.PaginatedResult[aclmodels.AclV2ListItem], error) {
	return c.Transport.GetByFilter(ctx, filter)
}
