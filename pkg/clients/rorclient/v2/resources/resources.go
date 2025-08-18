package resources

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/helpers/resourcecache/hashlist"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
)

type ResourcesInterface interface {
	Get(ctx context.Context, query rorresources.ResourceQuery) (*rorresources.ResourceSet, error)
	Update(ctx context.Context, res *rorresources.ResourceSet) (*rorresources.ResourceUpdateResults, error)
	Delete(ctx context.Context, uid string) (*rorresources.ResourceUpdateResults, error)
	Exists(ctx context.Context, uid string) (bool, error)
	GetOwnHashes(ctx context.Context, clientId rorresourceowner.RorResourceOwnerReference) (*hashlist.HashList, error)

	GetByUid(ctx context.Context, uid string) (*rorresources.ResourceSet, error)
	UpdateOne(ctx context.Context, resource *rorresources.Resource) (*rorresources.ResourceUpdateResults, error)
}
