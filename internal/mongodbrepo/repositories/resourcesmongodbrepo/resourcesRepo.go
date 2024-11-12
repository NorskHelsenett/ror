// package implements repofunctions to do crud operations on resources
package resourcesmongodbrepo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ResourceCollectionName = "resources"
)

// Checks if resource exists by uid, return bool, error
func ResourceExistbyUid(uid string, ctx context.Context) bool {
	db := mongodb.GetMongoDb()
	filter := bson.M{"uid": uid}

	resourceCount, err := db.Collection(ResourceCollectionName).CountDocuments(ctx, filter)
	if err != nil {
		return false
	}

	if resourceCount == 1 {
		return true
	}

	return false
}

// function to map resource to provided resourcemodel
func MapToResourceModel[D any](input any) D {
	var returnStruct D
	jsonva, _ := json.Marshal(input)
	err := json.Unmarshal(jsonva, &returnStruct)
	if err != nil {
		rlog.Error("Could not cast to resourcetype", err)
	}
	return returnStruct
}

func GetResourcesByQuery[T apiresourcecontracts.Resourcetypes](ctx context.Context, query apiresourcecontracts.ResourceQuery) ([]T, error) {
	db := mongodb.GetMongoDb()
	var returnvalue []T
	match := bson.M{}
	if !query.Global {
		match["owner.scope"] = query.Owner.Scope
		match["owner.subject"] = query.Owner.Subject
	}
	match["kind"] = query.Kind
	match["apiversion"] = query.ApiVersion
	if !query.Internal {
		match["internal"] = false
	}
	if query.Uid != "" {
		match["uid"] = query.Uid
	}
	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, bson.M{"$match": match})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$sort": bson.M{"resource.metadata.namespace": 1, "resource.metadata.name": 1}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$replaceRoot": bson.M{"newRoot": "$resource"}})

	clusterCollection := db.Collection(ResourceCollectionName)
	results, err := clusterCollection.Aggregate(ctx, aggregationPipeline)
	if err != nil {
		return returnvalue, fmt.Errorf("could not decode mongo document: %w", err)
	}

	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	if results.RemainingBatchLength() == 0 {
		return returnvalue, nil
	}

	for results.Next(ctx) {
		var dbResult T
		if err = results.Decode(&dbResult); err != nil {
			return returnvalue, fmt.Errorf("could not fetch resource: %v", err)
		}
		returnvalue = append(returnvalue, dbResult)
	}

	return returnvalue, nil
}
func GetResourceByQuery[T apiresourcecontracts.Resourcetypes](ctx context.Context, query apiresourcecontracts.ResourceQuery) (T, error) {
	db := mongodb.GetMongoDb()
	var returnvalue T
	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, bson.M{"$match": bson.M{"owner.scope": query.Owner.Scope, "owner.subject": query.Owner.Subject, "kind": query.Kind, "apiversion": query.ApiVersion, "uid": query.Uid}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$sort": bson.M{"resource.metadata.namespace": 1, "resource.metadata.name": 1}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$replaceRoot": bson.M{"newRoot": "$resource"}})

	clusterCollection := db.Collection(ResourceCollectionName)
	results, err := clusterCollection.Aggregate(ctx, aggregationPipeline)
	if err != nil {
		return returnvalue, fmt.Errorf("could not decode mongo document: %w", err)
	}

	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	if results.RemainingBatchLength() == 0 {
		return returnvalue, nil
	}

	results.Next(ctx)
	var dbResult T
	if err = results.Decode(&dbResult); err != nil {
		return returnvalue, fmt.Errorf("could not fetch resource: %v", err)
	}
	return dbResult, nil

}

// DeleteResourceByUid Delete resource by uid
func DeleteResourceByUid(resourceUpdate apiresourcecontracts.ResourceUpdateModel, ctx context.Context) error {
	db := mongodb.GetMongoDb()
	_, err := db.Collection(ResourceCollectionName).DeleteOne(ctx, bson.D{{Key: "uid", Value: resourceUpdate.Uid}})
	if err != nil {
		msg := fmt.Sprintf("Could not delete resource %s/%s with uid %s", resourceUpdate.ApiVersion, resourceUpdate.Kind, resourceUpdate.Uid)
		rlog.Error(msg, err)
		return errors.New(msg)
	}

	return nil
}

// GetHashList return list of registerd hashes by ownerref
func GetHashList(ctx context.Context, owner apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.HashList, error) {
	var hashList apiresourcecontracts.HashList
	db := mongodb.GetMongoDb()
	query := []bson.M{
		{
			"$match": bson.M{
				"owner.scope":   string(owner.Scope),
				"owner.subject": owner.Subject,
			},
		},
		{
			"$project": bson.M{
				"_id":  0,
				"uid":  "$uid",
				"hash": "$hash",
			},
		},
	}

	results, err := db.Collection(ResourceCollectionName).Aggregate(ctx, query)
	if err != nil {
		return hashList, fmt.Errorf("could not fetch clusters: %v", err)
	}

	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	if results.RemainingBatchLength() == 0 {
		var hashItems []apiresourcecontracts.HashItem = make([]apiresourcecontracts.HashItem, 0)
		hashList.Items = hashItems
		return hashList, nil
	}

	for results.Next(ctx) {
		var hashItem apiresourcecontracts.HashItem
		if err = results.Decode(&hashItem); err != nil {
			return hashList, fmt.Errorf("could not fetch clusters: %v", err)
		}
		hashList.Items = append(hashList.Items, hashItem)
	}

	return hashList, nil
}

func PatchResource(ctx context.Context, uid string, resourceUpdate bson.M) (mongo.UpdateResult, error) {
	if uid == "" {
		return mongo.UpdateResult{}, errors.New("uid is required")
	}
	result, err := mongodb.UpdateOne(ctx, ResourceCollectionName, bson.M{"uid": uid}, resourceUpdate)
	if err != nil {
		msg := fmt.Sprintf("Could not update resource with uid %s", uid)
		rlog.Error(msg, err)
		return mongo.UpdateResult{}, errors.New(msg)
	}
	return result, nil
}

// Only internal use, does not implement owner check
func GetResourceByUid[T apiresourcecontracts.Resourcetypes](ctx context.Context, uid string) (T, error) {
	var returnvalue T

	err := mongodb.FindOne(ctx, ResourceCollectionName, bson.M{"uid": uid}, &returnvalue)
	if err != nil {
		return returnvalue, fmt.Errorf("could not decode mongo document: %w", err)
	}

	return returnvalue, nil
}
