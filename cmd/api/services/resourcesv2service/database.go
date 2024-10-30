package resourcesv2service

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"
)

type ResourceDBProvider interface {
	Set(ctx context.Context, resource *rorresources.Resource) error
	Get(ctx context.Context, rorResourceQuery *rorresources.ResourceQuery) (*rorresources.ResourceSet, error)
	Del(ctx context.Context, resource *rorresources.Resource) error
	GetHashList(ctx context.Context, owner rortypes.RorResourceOwnerReference) (apiresourcecontracts.HashList, error)
}

// Mongodb implementation of ResourceDBProvider

type ResourceMongoDB struct {
	db *mongodb.MongodbCon
}

func NewResourceMongoDB(db *mongodb.MongodbCon) ResourceDBProvider {
	return &ResourceMongoDB{db: db}
}

func (r *ResourceMongoDB) Set(ctx context.Context, resource *rorresources.Resource) error {
	filter := bson.M{"uid": resource.GetUID()}
	update := bson.M{"$set": resource}
	_, err := r.db.UpsertOne(ctx, "resourcesv2", filter, update)
	if err != nil {
		rlog.Errorc(ctx, "Failed to upsert resource", err)
		return err
	}
	return nil
}

func (r *ResourceMongoDB) Get(ctx context.Context, rorResourceQuery *rorresources.ResourceQuery) (*rorresources.ResourceSet, error) {
	query := r.db.GenerateAggregateQuery(rorResourceQuery)
	var resources = make([]rorresources.Resource, 0)
	err := r.db.Aggregate(ctx, "resourcesv2", query, &resources)
	if err != nil {
		return nil, err
	}
	resourceSet := rorresources.NewResourceSet()
	if len(resources) > 0 {
		for _, resource := range resources {
			resourceSet.Add(rorresources.NewResourceFromStruct(resource))
		}
		return resourceSet, nil
	}
	return nil, nil
}

func (r *ResourceMongoDB) Del(ctx context.Context, resource *rorresources.Resource) error {
	filter := bson.M{"uid": resource.GetUID()}
	_, err := r.db.DeleteOne(ctx, "resourcesv2", filter)
	if err != nil {
		return err
	}
	return nil
}

func (r *ResourceMongoDB) GetHashList(ctx context.Context, owner rortypes.RorResourceOwnerReference) (apiresourcecontracts.HashList, error) {
	query := r.db.GenerateHashQuery(owner)
	var hashes = make([]apiresourcecontracts.HashItem, 0)
	err := r.db.Aggregate(ctx, "resourcesv2", query, &hashes)
	if err != nil {
		return apiresourcecontracts.HashList{}, err
	}
	return apiresourcecontracts.HashList{Items: hashes}, nil
}
