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
	set, err := r.Transport.Get(query)
	if err != nil {
		return nil, err
	}

	return &set, nil
}

func (r *ResourceClient) GetByUid(ctx context.Context, uid string) (*rorresources.ResourceSet, error) {
	query := rorresources.NewResourceQuery().WithUID(uid)

	set, err := r.Get(ctx, *query)
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (r *ResourceClient) Update(ctx context.Context, set rorresources.ResourceSet) (*rorresources.ResourceUpdateResults, error) {
	res, err := r.Transport.Update(&set)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *ResourceClient) UpdateOne(ctx context.Context, resource rorresources.Resource) (*rorresources.ResourceUpdateResults, error) {
	set := rorresources.NewResourceSet()
	set.Add(&resource)

	res, err := r.Update(ctx, *set)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *ResourceClient) Delete(ctx context.Context, uid string) (*rorresources.ResourceUpdateResults, error) {
	res, err := r.Transport.Delete(uid)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *ResourceClient) Exists(ctx context.Context, uid string) (bool, error) {
	res, err := r.Transport.Exists(uid)
	if err != nil {
		return false, err
	}

	return res, nil
}

func (r *ResourceClient) GetOwnHashes() (*apicontractsv2resources.HashList, error) {
	res, err := r.Transport.GetOwnHashes()
	if err != nil {
		return nil, err
	}

	return &res, nil
}
