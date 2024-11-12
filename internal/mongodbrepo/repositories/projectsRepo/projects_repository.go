package projectsrepo

import (
	"context"
	"errors"
	"fmt"
	mongohelper "github.com/NorskHelsenett/ror/internal/helpers/mongoHelper"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/mongoTypes"
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionName = "projects"
)

func Create(ctx context.Context, projectInput *mongoTypes.MongoProject) (*mongoTypes.MongoProject, error) {
	db := mongodb.GetMongoDb()
	projectInput.Created = time.Now()
	projectInput.Updated = time.Now()

	insertResult, err := db.Collection(CollectionName).InsertOne(ctx, projectInput)
	if err != nil {
		return nil, fmt.Errorf("could not insert project: %v", err)
	}

	var projectResult mongoTypes.MongoProject
	err = db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&projectResult)
	if err != nil {
		return nil, fmt.Errorf("could not find inserted project: %v", err)
	}

	return &projectResult, nil
}

func Update(ctx context.Context, updatedObject mongoTypes.MongoProject, projectId string) (*mongoTypes.MongoProject, *mongoTypes.MongoProject, error) {
	db := mongodb.GetMongoDb()
	mongoID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return nil, nil, fmt.Errorf("could not convert ID: %v", err)
	}

	var originalObject mongoTypes.MongoProject
	originalSingleResult := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoID})
	err = originalSingleResult.Decode(&originalObject)
	if err != nil {
		rlog.Error("could not get original object for auditlog", err)
	}

	updatedObject.Created = originalObject.Created
	updatedObject.Updated = time.Now()

	updateResult, err := db.Collection(CollectionName).ReplaceOne(ctx, bson.M{"_id": mongoID}, updatedObject)
	if err != nil {
		return nil, nil, fmt.Errorf("could not update object: %v", err)
	}

	if updateResult.MatchedCount == 0 {
		return nil, nil, fmt.Errorf("could not find object")
	}

	if updateResult.ModifiedCount == 0 {
		return nil, nil, fmt.Errorf("could not update object")
	}

	var result mongoTypes.MongoProject
	err = db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoID}).Decode(&result)
	if err != nil {
		return nil, nil, fmt.Errorf("could not find object after creation: %v", err)
	}

	return &result, &originalObject, nil
}

func Delete(ctx context.Context, id string) (bool, *mongoTypes.MongoProject, error) {
	db := mongodb.GetMongoDb()
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, nil, fmt.Errorf("could not convert ID: %v", err)
	}

	var originalObject mongoTypes.MongoProject
	originalSingleResult := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoId})
	err = originalSingleResult.Decode(&originalObject)
	if err != nil {
		rlog.Error("could not get original object for auditlog", err)
	}

	deleteResult, err := db.Collection(CollectionName).DeleteOne(ctx, bson.M{"_id": mongoId})
	if err != nil {
		return false, nil, fmt.Errorf("could not delete object: %v", err)
	}

	if deleteResult.DeletedCount == 0 {
		return false, nil, fmt.Errorf("could not delete object")
	}

	return true, &originalObject, nil
}

func GetByFilter(ctx context.Context, filter *apicontracts.Filter) ([]mongoTypes.MongoProject, int, error) {
	db := mongodb.GetMongoDb()
	aggregationPipeline := mongohelper.CreateAggregationPipeline(filter, apicontracts.SortMetadata{SortField: "created", SortOrder: 1}, []string{})
	totalCountPipeline := mongohelper.CreateFilterPipeline(filter, []string{})

	collection := db.Collection(CollectionName)
	cursor, err := collection.Aggregate(ctx, aggregationPipeline)
	if err != nil {
		return nil, 0, fmt.Errorf("error when finding projects: %v", err)
	}

	totalCountResult, err := collection.Aggregate(ctx, totalCountPipeline)
	if err != nil {
		return nil, 0, fmt.Errorf("could not fetch projects: %v", err)
	}

	var totalCountAcc []bson.M
	if err = totalCountResult.All(ctx, &totalCountAcc); err != nil {
		return nil, 0, fmt.Errorf("could not fetch total count for projects: %v", err)
	}

	totalCount := len(totalCountAcc)

	var results []mongoTypes.MongoProject
	for cursor.Next(ctx) {
		var project mongoTypes.MongoProject
		err := cursor.Decode(&project)
		if err != nil {
			return nil, 0, fmt.Errorf("could not decode project: %v", err)
		}
		results = append(results, project)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)
	defer func(totalCountResult *mongo.Cursor, ctx context.Context) {
		_ = totalCountResult.Close(ctx)
	}(totalCountResult, ctx)

	return results, totalCount, nil
}

func GetById(ctx context.Context, projectId string) (*mongoTypes.MongoProject, error) {
	db := mongodb.GetMongoDb()
	id, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return nil, fmt.Errorf("invalid ID: %v", err)
	}

	var object mongoTypes.MongoProject
	err = db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&object)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("could not find object: %v", err)
	}

	return &object, nil
}
