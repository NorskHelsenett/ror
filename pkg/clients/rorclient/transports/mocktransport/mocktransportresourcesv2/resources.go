package mocktransportresourcesv2

import (
	"context"
	"errors"

	"github.com/NorskHelsenett/ror/pkg/helpers/resourcecache/resourcecachehashlist"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type V2Client struct{}

func NewV2Client() *V2Client {
	return &V2Client{}
}

func (c *V2Client) Get(ctx context.Context, query rorresources.ResourceQuery) (*rorresources.ResourceSet, error) {
	resourceSet := &rorresources.ResourceSet{
		Resources: []*rorresources.Resource{
			{
				CommonResource: rortypes.CommonResource{
					Metadata: metav1.ObjectMeta{
						UID: types.UID("mock-resource-001"),
					},
				},
			},
		},
	}
	return resourceSet, nil
}

func (c *V2Client) Update(ctx context.Context, res *rorresources.ResourceSet) (*rorresources.ResourceUpdateResults, error) {
	if res == nil {
		return nil, errors.New("resource set cannot be nil")
	}

	results := &rorresources.ResourceUpdateResults{
		Results: map[string]rorresources.ResourceUpdateResult{
			"mock-update-result-001": {
				Status:  200,
				Message: "success",
			},
		},
	}
	return results, nil
}

func (c *V2Client) Delete(ctx context.Context, uid string) (*rorresources.ResourceUpdateResults, error) {
	if uid == "" {
		return nil, errors.New("UID cannot be empty")
	}

	results := &rorresources.ResourceUpdateResults{
		Results: map[string]rorresources.ResourceUpdateResult{
			uid: {
				Status:  200,
				Message: "deleted",
			},
		},
	}
	return results, nil
}

func (c *V2Client) Exists(ctx context.Context, uid string) (bool, error) {
	if uid == "" {
		return false, errors.New("UID cannot be empty")
	}
	// Mock implementation - always return true
	return true, nil
}

func (c *V2Client) GetOwnHashes(ctx context.Context, clientId rorresourceowner.RorResourceOwnerReference) (*resourcecachehashlist.HashList, error) {
	hashList := &resourcecachehashlist.HashList{
		Items: []resourcecachehashlist.HashItem{
			{
				Uid:  "mock-hash-001",
				Hash: "mock-hash-value-001",
			},
		},
	}
	return hashList, nil
}

func (c *V2Client) GetByUid(ctx context.Context, uid string) (*rorresources.ResourceSet, error) {
	if uid == "" {
		return nil, errors.New("UID cannot be empty")
	}

	resourceSet := &rorresources.ResourceSet{
		Resources: []*rorresources.Resource{
			{
				CommonResource: rortypes.CommonResource{
					Metadata: metav1.ObjectMeta{
						UID: types.UID(uid),
					},
				},
			},
		},
	}
	return resourceSet, nil
}

func (c *V2Client) UpdateOne(ctx context.Context, resource *rorresources.Resource) (*rorresources.ResourceUpdateResults, error) {
	if resource == nil {
		return nil, errors.New("resource cannot be nil")
	}

	uid := string(resource.Metadata.UID)
	if uid == "" {
		uid = "unknown"
	}

	results := &rorresources.ResourceUpdateResults{
		Results: map[string]rorresources.ResourceUpdateResult{
			uid: {
				Status:  200,
				Message: "updated",
			},
		},
	}
	return results, nil
}
