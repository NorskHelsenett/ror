package resources

import (
	"context"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2resources"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
)

type ResourcesInterface interface {
	Get(query rorresources.ResourceQuery) (rorresources.ResourceSet, error)
	Update(res *rorresources.ResourceSet) (*rorresources.ResourceUpdateResults, error)
	Delete(uid string) (*rorresources.ResourceUpdateResults, error)
	Exists(uid string) (bool, error)
	GetOwnHashes(clusterId string) (apicontractsv2resources.HashList, error)
	GetWithContext(ctx context.Context, query rorresources.ResourceQuery) (rorresources.ResourceSet, error)
	UpdateWithContext(ctx context.Context, res *rorresources.ResourceSet) (*rorresources.ResourceUpdateResults, error)
	DeleteWithContext(ctx context.Context, uid string) (*rorresources.ResourceUpdateResults, error)
	ExistsWithContext(ctx context.Context, uid string) (bool, error)
	GetOwnHashesWithContext(ctx context.Context, clusterId string) (apicontractsv2resources.HashList, error)
}
