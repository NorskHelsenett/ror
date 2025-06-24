package rorclient

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2resources"
	v2resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/resources"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
)

type ResourceClient struct {
	Transport v2resources.ResourcesInterface
}

func NewResourceClient(transport v2resources.ResourcesInterface) ResourceClient {
	client := ResourceClient{
		Transport: transport,
	}

	return client
}

func (r *ResourceClient) Get(ctx context.Context, query rorresources.ResourceQuery) (*rorresources.ResourceSet, error) {
	return r.Transport.Get(ctx, query)
}

func (r *ResourceClient) GetByUid(ctx context.Context, uid string) (*rorresources.ResourceSet, error) {
	query := rorresources.NewResourceQuery().WithUID(uid)

	return r.Get(ctx, *query)
}

func (r *ResourceClient) Update(ctx context.Context, set *rorresources.ResourceSet) (*rorresources.ResourceUpdateResults, error) {
	return r.Transport.Update(ctx, set)
}

func (r *ResourceClient) UpdateOne(ctx context.Context, resource *rorresources.Resource) (*rorresources.ResourceUpdateResults, error) {
	set := rorresources.NewResourceSet()
	set.Add(resource)

	return r.Update(ctx, set)
}

func (r *ResourceClient) Delete(ctx context.Context, uid string) (*rorresources.ResourceUpdateResults, error) {
	return r.Transport.Delete(ctx, uid)
}

func (r *ResourceClient) Exists(ctx context.Context, uid string) (bool, error) {
	return r.Transport.Exists(ctx, uid)
}

func (r *ResourceClient) GetOwnHashes(ctx context.Context, clusterId string) (*apicontractsv2resources.HashList, error) {
	return r.Transport.GetOwnHashes(ctx, clusterId)
}
