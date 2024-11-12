package pricesrepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/internal/helpers/mapping"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/mongoTypes"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionName = "prices"
)

func GetAll(ctx context.Context) (*[]apicontracts.Price, error) {
	db := mongodb.GetMongoDb()
	var query []bson.M
	cursor, err := db.Collection(CollectionName).Aggregate(ctx, query)
	if err != nil {
		return nil, errors.New("could not fetch all prices")
	}

	var prices []apicontracts.Price = make([]apicontracts.Price, 0)
	if cursor.RemainingBatchLength() == 0 {
		return &prices, nil
	}

	for cursor.Next(ctx) {
		var mongoP mongoTypes.MongoPrice
		if err = cursor.Decode(&mongoP); err != nil {
			rlog.Error("could not parse price to api contract definition", err)
			continue
		}

		var price apicontracts.Price
		err = mapping.Map(mongoP, &price)
		if err != nil {
			rlog.Error("could not map mongoPrice to Price", err)
			continue
		}

		prices = append(prices, price)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	return &prices, nil
}

func FindOne(ctx context.Context, property string, propertyValue string) (*apicontracts.Price, error) {
	db := mongodb.GetMongoDb()
	var priceResult mongoTypes.MongoPrice
	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{property: propertyValue}).Decode(&priceResult); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		msg := "could not find price"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	if priceResult.MachineClass == "" {
		return nil, nil
	}

	var mapped apicontracts.Price
	err := mapping.Map(priceResult, &mapped)
	if err != nil {
		return nil, nil
	}

	return &mapped, nil
}

func GetById(ctx context.Context, priceId string) (*apicontracts.Price, error) {
	db := mongodb.GetMongoDb()
	id, err := primitive.ObjectIDFromHex(priceId)
	if err != nil {
		return nil, fmt.Errorf("invalid price id: %v", err)
	}

	var priceResult mongoTypes.MongoPrice
	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&priceResult); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("could not find price: %v", err)
	}

	var mapped apicontracts.Price
	err = mapping.Map(priceResult, &mapped)
	if err != nil {
		return nil, fmt.Errorf("could not map price: %v", err)
	}

	return &mapped, nil
}

func GetByProperty(ctx context.Context, property string, propertyValue string) (*[]apicontracts.Price, error) {
	db := mongodb.GetMongoDb()
	query := []bson.M{
		{"$match": bson.M{property: propertyValue}},
	}
	cursor, err := db.Collection(CollectionName).Aggregate(ctx, query)
	if err != nil {
		rlog.Error("could not find prices in mongodb", err)
		return nil, errors.New("could not find prices")
	}
	if cursor.RemainingBatchLength() == 0 {
		emptyList := make([]apicontracts.Price, 0)
		return &emptyList, nil
	}

	var prices []apicontracts.Price
	for cursor.Next(ctx) {
		var price apicontracts.Price
		if err = cursor.Decode(&price); err != nil {
			rlog.Debug("could not decode price")
		}

		prices = append(prices, price)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	return &prices, nil
}

func Create(ctx context.Context, priceInput *mongoTypes.MongoPrice) (*mongoTypes.MongoPrice, error) {
	db := mongodb.GetMongoDb()
	insertResult, err := db.Collection(CollectionName).InsertOne(ctx, priceInput)
	if err != nil {
		return nil, fmt.Errorf("could not insert price: %v", err)
	}

	if insertResult.InsertedID == nil {
		return nil, fmt.Errorf("could not create price")
	}

	var priceResult mongoTypes.MongoPrice
	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&priceResult); err != nil {
		return nil, fmt.Errorf("could not find inserted price: %v", err)
	}

	return &priceResult, nil
}

func Update(ctx context.Context, priceId string, priceInput mongoTypes.MongoPrice) (*mongoTypes.MongoPrice, *mongoTypes.MongoPrice, error) {
	db := mongodb.GetMongoDb()
	mongoID, err := primitive.ObjectIDFromHex(priceId)
	if err != nil {
		return nil, nil, fmt.Errorf("could not convert priceId: %v", err)
	}

	var originalMongoPrice mongoTypes.MongoPrice
	originalSingleResult := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoID})
	err = originalSingleResult.Decode(&originalMongoPrice)
	if err != nil {
		rlog.Error("could not get original price for auditlog", err)
	}

	updateResult, err := db.Collection(CollectionName).ReplaceOne(ctx, bson.M{"_id": mongoID}, priceInput)
	if err != nil {
		return nil, nil, fmt.Errorf("could not update price: %v", err)
	}

	if updateResult.MatchedCount == 0 {
		return nil, nil, fmt.Errorf("could not find price")
	}

	if updateResult.ModifiedCount == 0 {
		return nil, nil, fmt.Errorf("could not update price")
	}

	var priceResult mongoTypes.MongoPrice
	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoID}).Decode(&priceResult); err != nil {
		return nil, nil, fmt.Errorf("could not find price after creation: %v", err)
	}

	return &priceResult, &originalMongoPrice, nil
}

func Delete(ctx context.Context, priceId string) (bool, *mongoTypes.MongoPrice, error) {
	db := mongodb.GetMongoDb()
	mongoID, err := primitive.ObjectIDFromHex(priceId)
	if err != nil {
		return false, nil, fmt.Errorf("could not convert priceId: %v", err)
	}

	var originalMongoPrice mongoTypes.MongoPrice
	originalSingleResult := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoID})
	err = originalSingleResult.Decode(&originalMongoPrice)
	if err != nil {
		rlog.Error("could not get original price for auditlog", err)
	}

	deleteResult, err := db.Collection(CollectionName).DeleteOne(ctx, bson.M{"_id": mongoID})
	if err != nil {
		return false, nil, fmt.Errorf("could not delete price: %v", err)
	}

	if deleteResult.DeletedCount == 0 {
		return false, nil, fmt.Errorf("could not delete price")
	}

	return true, &originalMongoPrice, nil
}
