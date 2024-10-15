package auditlogrepo

import (
	"context"
	"errors"
	"fmt"
	mongoHelper "github.com/NorskHelsenett/ror/internal/helpers/mongoHelper"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/mongoTypes"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionName = "auditlogs"
)

func Create(ctx context.Context, auditLog mongoTypes.MongoAuditLog) (*mongoTypes.MongoAuditLog, error) {
	db := mongodb.GetMongoDb()
	collection := db.Collection(collectionName)

	insertResult, err := collection.InsertOne(ctx, auditLog)
	if err != nil {
		return nil, fmt.Errorf("unable to save auditlog: %v", err)
	}

	if insertResult.InsertedID == nil {
		return nil, fmt.Errorf("unable to get id of saved auditlog")
	}

	insertedAuditlog := collection.FindOne(ctx, bson.M{"_id": insertResult.InsertedID})
	var decodedAuditlog mongoTypes.MongoAuditLog
	err = insertedAuditlog.Decode(&decodedAuditlog)
	if err != nil {
		return nil, fmt.Errorf("unable to decode auditlog: %v", err)
	}

	return &decodedAuditlog, nil
}

func GetByFilter(ctx context.Context, filter *apicontracts.Filter) ([]mongoTypes.MongoAuditLog, int, error) {
	db := mongodb.GetMongoDb()
	aggregationPipeline := mongoHelper.CreateAggregationPipeline(filter, apicontracts.SortMetadata{SortField: "metadata.timestamp", SortOrder: -1}, []string{"metadata"})
	totalCountPipeline := mongoHelper.CreateFilterPipeline(filter, []string{"metadata"})

	collection := db.Collection(collectionName)
	cursor, err := collection.Aggregate(ctx, aggregationPipeline)
	if err != nil {
		return nil, 0, fmt.Errorf("error when finding auditlogs: %v", err)
	}

	totalCountResult, err := collection.Aggregate(ctx, totalCountPipeline)
	if err != nil {
		return nil, 0, fmt.Errorf("could not fetch auditlogs: %v", err)
	}

	var totalCountAcc []bson.M
	if err = totalCountResult.All(ctx, &totalCountAcc); err != nil {
		return nil, 0, fmt.Errorf("could not fetch total count for auditlogs: %v", err)
	}

	totalCount := len(totalCountAcc)

	var results []mongoTypes.MongoAuditLog
	for cursor.Next(ctx) {
		var auditLog mongoTypes.MongoAuditLog
		err := cursor.Decode(&auditLog)
		if err != nil {
			return nil, 0, fmt.Errorf("could not decode audit log: %v", err)
		}
		results = append(results, auditLog)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)
	defer func(totalCountResult *mongo.Cursor, ctx context.Context) {
		_ = totalCountResult.Close(ctx)
	}(totalCountResult, ctx)

	return results, totalCount, nil
}

func GetByID(ctx context.Context, ID string) (*mongoTypes.MongoAuditLog, error) {
	db := mongodb.GetMongoDb()
	mongoID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, fmt.Errorf("could not get ID from string: %v", err)
	}

	collection := db.Collection(collectionName)

	var auditlog mongoTypes.MongoAuditLog
	result := collection.FindOne(ctx, bson.M{"_id": mongoID})
	err = result.Decode(&auditlog)
	if err != nil {
		return nil, fmt.Errorf("could not decode mongoauditlog: %v", err)
	}

	return &auditlog, nil
}

func GetMetadata(ctx context.Context) (map[string][]string, error) {
	db := mongodb.GetMongoDb()
	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, bson.M{
		"$group": bson.M{
			"_id": nil,
			"actions": bson.M{
				"$addToSet": "$metadata.action",
			},
			"categories": bson.M{
				"$addToSet": "$metadata.category",
			},
		},
	})
	aggregationPipeline = append(aggregationPipeline, bson.M{
		"$project": bson.M{
			"_id":        false,
			"actions":    "$actions",
			"categories": "$categories",
		},
	})
	clusterCollection := db.Collection(collectionName)
	cursor, err := clusterCollection.Aggregate(ctx, aggregationPipeline)
	if err != nil {
		return nil, fmt.Errorf("could not perform aggregation: %v", err)
	}

	if cursor.RemainingBatchLength() == 0 {
		return nil, nil
	}

	var metadataSlice []map[string][]string
	for cursor.Next(ctx) {
		var metadata map[string][]string
		err = cursor.Decode(&metadata)
		if err != nil {
			return nil, fmt.Errorf("could not decode mongo document: %v", err)
		}
		metadataSlice = append(metadataSlice, metadata)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	if len(metadataSlice) > 0 {
		return metadataSlice[0], nil
	}

	return nil, errors.New("missing metadata")
}
