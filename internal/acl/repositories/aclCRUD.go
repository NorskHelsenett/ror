package aclrepository

import (
	"context"
	"errors"
	"fmt"

	mongoHelper "github.com/NorskHelsenett/ror/internal/helpers/mongoHelper"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetById(ctx context.Context, Id string) (*aclmodels.AclV2ListItem, error) {
	db := mongodb.GetMongoDb()
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	var object aclmodels.AclV2ListItem
	err = db.Collection(collectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&object)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("could not find object: %v", err)
	}

	return &object, nil
}

func GetByFilter(ctx context.Context, filter *apicontracts.Filter) ([]aclmodels.AclV2ListItem, int, error) {
	db := mongodb.GetMongoDb()
	aggregationPipeline := mongoHelper.CreateAggregationPipeline(filter, apicontracts.SortMetadata{SortField: "group", SortOrder: 1}, []string{})
	totalCountPipeline := mongoHelper.CreateFilterPipeline(filter, []string{"group"})

	collection := db.Collection(collectionName)
	cursor, err := collection.Aggregate(ctx, aggregationPipeline)
	if err != nil {
		return nil, 0, fmt.Errorf("error when finding acl: %v", err)
	}

	totalCountResult, err := collection.Aggregate(ctx, totalCountPipeline)
	if err != nil {
		return nil, 0, fmt.Errorf("could not fetch acl: %v", err)
	}

	var totalCountAcc []bson.M
	if err = totalCountResult.All(ctx, &totalCountAcc); err != nil {
		return nil, 0, fmt.Errorf("could not fetch total count for acl: %v", err)
	}

	totalCount := len(totalCountAcc)

	var results []aclmodels.AclV2ListItem
	for cursor.Next(ctx) {
		var acl aclmodels.AclV2ListItem
		err := cursor.Decode(&acl)
		if err != nil {
			return nil, 0, fmt.Errorf("could not decode acl log: %v", err)
		}
		results = append(results, acl)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)
	defer func(totalCountResult *mongo.Cursor, ctx context.Context) {
		_ = totalCountResult.Close(ctx)
	}(totalCountResult, ctx)

	return results, totalCount, nil
}

func Create(ctx context.Context, aclModel *aclmodels.AclV2ListItem) (*aclmodels.AclV2ListItem, error) {
	db := mongodb.GetMongoDb()
	aclModel.Version = 2
	insertResult, err := db.Collection(collectionName).InsertOne(ctx, aclModel)
	if err != nil {
		return nil, fmt.Errorf("could not insert acl: %v", err)
	}

	var result aclmodels.AclV2ListItem
	err = db.Collection(collectionName).FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("could not find inserted acl: %v", err)
	}

	return &result, nil
}

func Update(ctx context.Context, aclId string, aclModel *aclmodels.AclV2ListItem) (*aclmodels.AclV2ListItem, *aclmodels.AclV2ListItem, error) {
	db := mongodb.GetMongoDb()
	mongoId, err := primitive.ObjectIDFromHex(aclId)
	if err != nil {
		return nil, nil, fmt.Errorf("could not convert ID: %v", err)
	}

	aclModel.Version = 2

	var originalObject aclmodels.AclV2ListItem
	originalSingleResult := db.Collection(collectionName).FindOne(ctx, bson.M{"_id": mongoId})
	err = originalSingleResult.Decode(&originalObject)
	if err != nil {
		rlog.Error("could not get original object", err)
	}

	aclModel.Created = originalObject.Created
	aclModel.IssuedBy = originalObject.IssuedBy

	updateResult, err := db.Collection(collectionName).ReplaceOne(ctx, bson.M{"_id": mongoId}, aclModel)
	if err != nil {
		return nil, nil, fmt.Errorf("could not update object: %v", err)
	}

	if updateResult.MatchedCount == 0 {
		return nil, nil, fmt.Errorf("could not find object")
	}

	if updateResult.ModifiedCount == 0 {
		return nil, nil, fmt.Errorf("could not update object")
	}

	var result aclmodels.AclV2ListItem
	err = db.Collection(collectionName).FindOne(ctx, bson.M{"_id": mongoId}).Decode(&result)
	if err != nil {
		return nil, nil, fmt.Errorf("could not find object after creation: %v", err)
	}

	return &result, &originalObject, nil
}

func Delete(ctx context.Context, id string) (bool, *aclmodels.AclV2ListItem, error) {
	db := mongodb.GetMongoDb()
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, nil, fmt.Errorf("could not convert ID: %v", err)
	}

	var originalObject aclmodels.AclV2ListItem
	originalSingleResult := db.Collection(collectionName).FindOne(ctx, bson.M{"_id": mongoId})
	err = originalSingleResult.Decode(&originalObject)
	if err != nil {
		rlog.Error("could not get original object for auditlog", err)
	}

	deleteResult, err := db.Collection(collectionName).DeleteOne(ctx, bson.M{"_id": mongoId})
	if err != nil {
		return false, nil, fmt.Errorf("could not delete object: %v", err)
	}

	if deleteResult.DeletedCount == 0 {
		return false, nil, fmt.Errorf("could not delete object")
	}

	return true, &originalObject, nil
}
