package desiredversionrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionName = "desired_versions"
)

func GetAll(ctx context.Context) ([]apicontracts.DesiredVersion, error) {
	db := mongodb.GetMongoDb()
	query := bson.D{}
	var result []apicontracts.DesiredVersion

	cursor, err := db.Collection(CollectionName).Find(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("could not find desired versions: %v", err)
	}

	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("could not find desired versions: %v", err)
	}

	return result, nil
}

func GetByKey(ctx context.Context, key string) (*apicontracts.DesiredVersion, error) {
	db := mongodb.GetMongoDb()
	query := bson.M{"key": key}
	var result apicontracts.DesiredVersion

	if err := db.Collection(CollectionName).
		FindOne(ctx, query).
		Decode(&result); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("could not find desired version: %v", err)
	}

	return &result, nil
}

func GetByID(ctx context.Context, id interface{}) (*apicontracts.DesiredVersion, error) {
	db := mongodb.GetMongoDb()
	query := bson.M{"_id": id}
	var result apicontracts.DesiredVersion

	if err := db.Collection(CollectionName).
		FindOne(ctx, query).
		Decode(&result); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("could not find desired version: %v", err)
	}

	return &result, nil
}

func Create(ctx context.Context, desiredversion apicontracts.DesiredVersion) (*mongo.InsertOneResult, error) {
	db := mongodb.GetMongoDb()
	query := bson.M{"key": desiredversion.Key, "value": desiredversion.Value}

	insertResult, err := db.Collection(CollectionName).InsertOne(ctx, query)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func UpdateByKey(ctx context.Context, key string, desiredversion apicontracts.DesiredVersion) (*mongo.UpdateResult, error) {
	db := mongodb.GetMongoDb()
	query := bson.M{"key": key}
	update := bson.M{"key": desiredversion.Key, "value": desiredversion.Value}

	updateResult, err := db.Collection(CollectionName).ReplaceOne(ctx, query, update)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func DeleteByKey(ctx context.Context, key string) (*mongo.DeleteResult, error) {
	db := mongodb.GetMongoDb()
	query := bson.M{"key": key}

	deleteResult, err := db.Collection(CollectionName).DeleteOne(ctx, query)
	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}
