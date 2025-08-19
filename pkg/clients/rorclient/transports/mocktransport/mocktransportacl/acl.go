package mocktransportacl

import (
	"context"
	"errors"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

type V1Client struct{}

func NewV1Client() *V1Client {
	return &V1Client{}
}

func (c *V1Client) Create(ctx context.Context, item aclmodels.AclV2ListItem) error {
	// Mock implementation - just return nil to simulate success
	return nil
}

func (c *V1Client) Update(ctx context.Context, id string, item aclmodels.AclV2ListItem) error {
	if id == "" {
		return errors.New("ID cannot be empty")
	}
	// Mock implementation - just return nil to simulate success
	return nil
}

func (c *V1Client) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("ID cannot be empty")
	}
	// Mock implementation - just return nil to simulate success
	return nil
}

func (c *V1Client) CheckAccess(ctx context.Context, scope, subject, access string) bool {
	// Mock implementation - always return true for now
	return true
}

func (c *V1Client) GetById(ctx context.Context, id string) (*aclmodels.AclV2ListItem, error) {
	if id == "" {
		return nil, errors.New("ID cannot be empty")
	}

	item := &aclmodels.AclV2ListItem{
		Id: id,
	}
	return item, nil
}

func (c *V1Client) GetByFilter(ctx context.Context, filter apicontracts.Filter) (*apicontracts.PaginatedResult[aclmodels.AclV2ListItem], error) {
	result := &apicontracts.PaginatedResult[aclmodels.AclV2ListItem]{
		Data: []aclmodels.AclV2ListItem{
			{
				Id: "mock-acl-001",
			},
			{
				Id: "mock-acl-002",
			},
		},
		TotalCount: 2,
	}
	return result, nil
}
