package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Aggregate(ctx context.Context, col string, query []bson.M, value interface{}) error {
	db := GetMongoDb()
	resourceCollection := db.Collection(col)
	results, err := resourceCollection.Aggregate(ctx, query)
	if err != nil {
		return fmt.Errorf("could not decode mongo document: %w", err)
	}
	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	err = results.All(ctx, value)
	if err != nil {
		return fmt.Errorf("could not fetch resource: %v", err)
	}

	return nil
}

func UpdateOne(ctx context.Context, col string, filter bson.M, update bson.M) (mongo.UpdateResult, error) {
	db := GetMongoDb()
	updateResult, err := db.Collection(col).UpdateOne(ctx, filter, update)
	if err != nil {
		return mongo.UpdateResult{}, fmt.Errorf("could not updateOne object: %v", err)
	}

	return *updateResult, nil
}

func InsertOne(ctx context.Context, col string, input interface{}) (mongo.InsertOneResult, error) {
	db := GetMongoDb()
	insertResult, err := db.Collection(col).InsertOne(ctx, input)
	if err != nil {
		return mongo.InsertOneResult{}, fmt.Errorf("could not insertOne object: %v", err)
	}

	return *insertResult, nil
}

func FindOne(ctx context.Context, col string, query bson.M, value interface{}) error {
	db := GetMongoDb()
	err := db.Collection(col).FindOne(ctx, query).Decode(value)
	if err != nil {
		return fmt.Errorf("could not findOne object: %v", err)
	}

	return nil
}

func DeleteOne(ctx context.Context, col string, query bson.M) (mongo.DeleteResult, error) {
	db := GetMongoDb()
	result, err := db.Collection(col).DeleteOne(ctx, query)
	if err != nil {
		return mongo.DeleteResult{}, fmt.Errorf("could not insertOne object: %v", err)
	}
	return *result, nil
}

func Count(ctx context.Context, col string) (int64, error) {
	db := GetMongoDb()
	opts := options.Count().SetHint("_id_")
	count, err := db.Collection(col).CountDocuments(ctx, bson.D{}, opts)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountWithFilter(ctx context.Context, col string, filter interface{}) (int64, error) {
	db := GetMongoDb()
	opts := options.Count().SetHint("_id_")
	count, err := db.Collection(col).CountDocuments(ctx, filter, opts)

	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountWithQuery(ctx context.Context, col string, query any) (int, error) {
	db := GetMongoDb()
	resourceCollection := db.Collection(col)
	results, err := resourceCollection.Aggregate(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("could not decode mongo document: %w", err)
	}
	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)
	returnvalue := results.RemainingBatchLength()
	return returnvalue, nil
}
