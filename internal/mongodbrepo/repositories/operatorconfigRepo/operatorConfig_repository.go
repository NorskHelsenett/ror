package operatorconfigrepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/internal/helpers/mapping"
	mongoHelper "github.com/NorskHelsenett/ror/internal/helpers/mongoHelper"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/mongoTypes"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionName = "operatorconfigs"
)

func FindOne(ctx context.Context, property string, propertyValue string) (*apicontracts.OperatorConfig, error) {
	db := mongodb.GetMongoDb()
	var result mongoTypes.MongoOperatorConfig
	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{property: propertyValue}).Decode(&result); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		msg := "could not find operator config"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	if result.ApiVersion == "" || result.Kind == "" || result.Spec == nil {
		return nil, nil
	}

	var mapped apicontracts.OperatorConfig
	err := mapping.Map(result, &mapped)
	if err != nil {
		return nil, nil
	}

	return &mapped, nil
}

// func GetByFilter(ctx context.Context, filter *apicontracts.NewFilter) ([]mongoTypes.MongoOperatorConfig, int, error) {
// 	fields := []string{ "kind", "version" }
// 	pipeline := mongohelper.CreateAggregationPipeline(filter, apicontracts.NewSortMetadata{SortField: "kind", SortOrder: 1}, fields)
// 	totalCountPipeline := mongohelper.CreateFilterPipeline(filter, fields)

// 	collection := db.Collection(CollectionName)
// 	cursor, err := collection.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, 0, fmt.Errorf("error when finding operatorconfigs: %v", err)
// 	}

// 	totalCountResult, err := collection.Aggregate(ctx, totalCountPipeline)
// 	if err != nil {
// 		return nil, 0, fmt.Errorf("could not fetch operatorconfigs: %v", err)
// 	}

// 	var totalCountAcc []bson.M
// 	if err = totalCountResult.All(ctx, &totalCountAcc); err != nil {
// 		return nil, 0, fmt.Errorf("could not fetch total count for operatorconfig: %v", err)
// 	}

// 	totalCount := len(totalCountAcc)
// 	results := []mongoTypes.MongoOperatorConfig{}
// 	for cursor.Next(ctx) {
// 		var config mongoTypes.MongoOperatorConfig
// 		err := cursor.Decode(&config)
// 		if err != nil {
// 			return nil, 0, fmt.Errorf("could not decode operatorconfig: %v", err)
// 		}
// 		results = append(results, config)
// 	}
// 	err = cursor.Close(ctx)
// 	if err != nil {
// 		return nil, 0, fmt.Errorf("could not close cursor: %v", err)
// 	}

// 	return results, totalCount, nil
// }

func GetByFilter(ctx context.Context, filter *apicontracts.Filter) (*apicontracts.PaginatedResult[apicontracts.OperatorConfig], error) {
	db := mongodb.GetMongoDb()
	aggregatePipeline := mongoHelper.CreateAggregationPipeline(filter, apicontracts.SortMetadata{SortField: "apiversion", SortOrder: 1}, []string{"kind", "apiversion"})

	var query []bson.M
	query = append(query, aggregatePipeline...)
	var totalCountQuery []bson.M

	configCollection := db.Collection(CollectionName)
	results, err := db.Collection(CollectionName).Aggregate(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("could not fetch operator configs: %v", err)
	}

	totalCountResult, err := configCollection.Aggregate(ctx, totalCountQuery)
	if err != nil {
		return nil, fmt.Errorf("could not fetch operator configs: %v", err)
	}

	var totalCountAcc []bson.M
	if err = totalCountResult.All(ctx, &totalCountAcc); err != nil {
		return nil, fmt.Errorf("could not fetch total count for operator configs: %v", err)
	}

	totalCount := len(totalCountAcc)

	//reading from the db in an optimal way
	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)
	if results.RemainingBatchLength() == 0 {
		emptyResult := apicontracts.PaginatedResult[apicontracts.OperatorConfig]{}
		return &emptyResult, nil
	}

	var operatorConfigs []apicontracts.OperatorConfig = make([]apicontracts.OperatorConfig, 0)
	paginatedResult := apicontracts.PaginatedResult[apicontracts.OperatorConfig]{}
	for results.Next(ctx) {
		var element apicontracts.OperatorConfig
		if err = results.Decode(&element); err != nil {
			return nil, fmt.Errorf("could not fetch operator configs: %v", err)
		}

		operatorConfigs = append(operatorConfigs, element)
	}

	paginatedResult.Data = operatorConfigs
	paginatedResult.DataCount = int64(len(operatorConfigs))
	paginatedResult.Offset = int64(filter.Skip)
	paginatedResult.TotalCount = int64(totalCount)

	return &paginatedResult, nil
}

func GetById(ctx context.Context, operatorConfigId string) (*apicontracts.OperatorConfig, error) {
	db := mongodb.GetMongoDb()
	id, err := primitive.ObjectIDFromHex(operatorConfigId)
	if err != nil {
		return nil, fmt.Errorf("invalid operator config id: %v", err)
	}

	var result mongoTypes.MongoOperatorConfig
	if err := db.Collection(CollectionName).
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&result); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("could not find operator config: %v", err)
	}

	var mapped apicontracts.OperatorConfig
	err = mapping.Map(result, &mapped)
	if err != nil {
		return nil, fmt.Errorf("could not map operator config: %v", err)
	}

	return &mapped, nil
}

func Create(ctx context.Context, operatorConfig *apicontracts.OperatorConfig) (*apicontracts.OperatorConfig, error) {
	db := mongodb.GetMongoDb()
	var mongoInput mongoTypes.MongoOperatorConfig
	err := mapping.Map(operatorConfig, &mongoInput)

	if err != nil {
		return nil, errors.New("could not parse input to mongo model")
	}

	insertResult, err := db.Collection(CollectionName).InsertOne(ctx, mongoInput)
	if err != nil {
		msg := "could not create operator config"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	var result mongoTypes.MongoOperatorConfig
	if insertResult.InsertedID == nil {
		return nil, errors.New("could not create operator config")
	}

	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&result); err != nil {
		rlog.Error("Could not find operator config", err)
		return nil, errors.New("could not find operator config after creation")
	}

	var mapped apicontracts.OperatorConfig
	_ = mapping.Map(result, &mapped)

	return &mapped, nil
}

func GetAll(ctx context.Context) (*[]apicontracts.OperatorConfig, error) {
	db := mongodb.GetMongoDb()
	var query []bson.M
	cursor, err := db.Collection(CollectionName).Aggregate(ctx, query)
	if err != nil {
		return nil, errors.New("could not fetch all operator configs")
	}

	var result []apicontracts.OperatorConfig = make([]apicontracts.OperatorConfig, 0)
	if cursor.RemainingBatchLength() == 0 {
		return &result, nil
	}

	for cursor.Next(ctx) {
		var mongoP mongoTypes.MongoOperatorConfig
		if err = cursor.Decode(&mongoP); err != nil {
			rlog.Error("could not parse operator config to api contract definition", err)
			continue
		}

		var operatorConfig apicontracts.OperatorConfig
		err = mapping.Map(mongoP, &operatorConfig)
		if err != nil {
			rlog.Error("could not map mongoOperatorConfig to OperatorConfig", err)
			continue
		}

		result = append(result, operatorConfig)
	}

	return &result, nil
}

func Update(ctx context.Context, id string, input mongoTypes.MongoOperatorConfig) (*mongoTypes.MongoOperatorConfig, *mongoTypes.MongoOperatorConfig, error) {
	db := mongodb.GetMongoDb()
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, nil, fmt.Errorf("could not convert operatorConfigId: %v", err)
	}

	var original mongoTypes.MongoOperatorConfig
	originalSingleResult := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoID})
	err = originalSingleResult.Decode(&original)
	if err != nil {
		rlog.Error("could not get original operatorconfig for auditlog", err)
	}

	updateResult, err := db.Collection(CollectionName).ReplaceOne(ctx, bson.M{"_id": mongoID}, input)
	if err != nil {
		return nil, nil, fmt.Errorf("could not update operatorConfig: %v", err)
	}

	if updateResult.MatchedCount == 0 {
		return nil, nil, fmt.Errorf("could not find operatorConfig")
	}

	if updateResult.ModifiedCount == 0 {
		return nil, nil, fmt.Errorf("could not update operatorConfig")
	}

	var result mongoTypes.MongoOperatorConfig
	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoID}).Decode(&result); err != nil {
		return nil, nil, fmt.Errorf("could not find operatorConfig after creation: %v", err)
	}

	return &result, &original, nil
}

func Delete(ctx context.Context, id string) (bool, error) {
	db := mongodb.GetMongoDb()
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, fmt.Errorf("could not convert string id to mongoid: %v", err)
	}

	var existing mongoTypes.MongoOperatorConfig
	originalSingleResult := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoID})
	err = originalSingleResult.Decode(&existing)
	if err != nil {
		rlog.Error("could not get original operatorconfig for auditlog", err)
	}

	deleteResult, err := db.Collection(CollectionName).DeleteOne(ctx, bson.M{"_id": mongoID})
	if err != nil {
		return false, fmt.Errorf("could not delete operatorconfig: %v", err)
	}

	if deleteResult.DeletedCount == 0 {
		return false, fmt.Errorf("could not delete operatorconfig")
	}

	return true, nil
}
