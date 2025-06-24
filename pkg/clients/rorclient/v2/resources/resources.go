package resources

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2resources"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
)

type ResourcesInterface interface {
	Get(ctx context.Context, query rorresources.ResourceQuery) (rorresources.ResourceSet, error)
	Update(ctx context.Context, res *rorresources.ResourceSet) (*rorresources.ResourceUpdateResults, error)
	Delete(ctx context.Context, uid string) (*rorresources.ResourceUpdateResults, error)
	Exists(ctx context.Context, uid string) (bool, error)
	GetOwnHashes(ctx context.Context, clusterId string) (apicontractsv2resources.HashList, error)
}
