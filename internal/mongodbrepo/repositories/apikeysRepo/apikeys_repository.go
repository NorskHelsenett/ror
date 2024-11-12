package apikeyrepo

import (
	"context"
	"fmt"
	mongoHelper "github.com/NorskHelsenett/ror/internal/helpers/mongoHelper"
	"time"

	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionName = "apikeys"
)

func GetByHash(ctx context.Context, hashedapikey string) ([]apicontracts.ApiKey, error) {
	var aggregationPipeline = []bson.M{
		{"$match": bson.M{"hash": hashedapikey}},
	}
	var results = make([]apicontracts.ApiKey, 0)
	mongoctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()
	err := mongodb.Aggregate(mongoctx, collectionName, aggregationPipeline, &results)
	if err != nil {
		return results, fmt.Errorf("error finding apikeys: %v", err)
	}
	return results, nil
}

func GetByIdentifier(ctx context.Context, identifier string) ([]apicontracts.ApiKey, error) {
	var aggregationPipeline = []bson.M{
		{"$match": bson.M{"identifier": identifier}},
	}
	var results = make([]apicontracts.ApiKey, 0)
	mongoctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()
	err := mongodb.Aggregate(mongoctx, collectionName, aggregationPipeline, &results)
	if err != nil {
		return results, fmt.Errorf("error finding apikeys: %v", err)
	}
	return results, nil
}

func GetOwnByName(ctx context.Context, name string) (*apicontracts.ApiKey, error) {
	identity := rorcontext.GetIdentityFromRorContext(ctx)

	var aggregationPipeline = []bson.M{
		{"$match": bson.M{"displayname": name, "identifier": identity.GetId()}},
	}
	var results = make([]apicontracts.ApiKey, 0)
	mongoctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()
	mongoHelper.PrettyprintBSON(aggregationPipeline)
	err := mongodb.Aggregate(mongoctx, collectionName, aggregationPipeline, &results)
	if err != nil {
		return nil, fmt.Errorf("error finding apikeys: %v", err)
	}
	if len(results) > 1 {
		return nil, fmt.Errorf("found more than one apikey with description: %s for user %s", name, identity.GetId())
	}
	if len(results) == 0 {
		return nil, nil
	}
	return &results[0], nil
}

func GetByFilter(ctx context.Context, filter *apicontracts.Filter) ([]apicontracts.ApiKey, int, error) {
	aggregationPipeline := mongoHelper.CreateAggregationPipeline(filter, apicontracts.SortMetadata{SortField: "created", SortOrder: 1}, []string{})

	var results = make([]apicontracts.ApiKey, 0)
	err := mongodb.Aggregate(ctx, collectionName, aggregationPipeline, &results)
	if err != nil {
		return nil, 0, fmt.Errorf("error finding apikeys: %v", err)
	}

	totalCount, err := mongodb.CountWithQuery(ctx, collectionName, aggregationPipeline)
	if err != nil {
		return nil, 0, fmt.Errorf("could not fetch total count for apikey: %v", err)
	}
	return results, totalCount, nil
}

func UpdateOwnByName(ctx context.Context, name string, hash string, expires time.Time) error {

	existing, err := GetOwnByName(ctx, name)
	if err != nil {
		return fmt.Errorf("failed to query existing apikey: %v", err)
	}

	identity := rorcontext.GetIdentityFromRorContext(ctx)
	identifier := identity.GetId()
	objectId, err := primitive.ObjectIDFromHex(existing.Id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId, "identifier": identifier}
	update := bson.M{"$set": bson.M{"hash": hash, "expires": expires}}

	_, err = mongodb.UpdateOne(ctx, collectionName, filter, update)

	if err != nil {
		return err
	}
	return nil
}

func Delete(ctx context.Context, ID string) (bool, apicontracts.ApiKey, error) {
	var originalObject apicontracts.ApiKey
	mongoID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return false, originalObject, fmt.Errorf("could not convert ID: %v", err)
	}

	query := bson.M{"_id": mongoID}
	err = mongodb.FindOne(ctx, collectionName, query, &originalObject)
	if err != nil {
		rlog.Error("could not get original object for auditlog", err)
	}

	deleteQuery := bson.M{"_id": mongoID}
	deleteResult, err := mongodb.DeleteOne(ctx, collectionName, deleteQuery)
	if err != nil {
		return false, originalObject, fmt.Errorf("could not delete object: %v", err)
	}

	if deleteResult.DeletedCount == 0 {
		return false, originalObject, fmt.Errorf("could not delete object")
	}

	return true, originalObject, nil
}

func Create(ctx context.Context, input apicontracts.ApiKey) error {
	input.Created = time.Now()
	input.Id = ""

	_, err := mongodb.InsertOne(ctx, collectionName, input)
	if err != nil {
		return fmt.Errorf("could not insert project: %v", err)
	}

	return nil
}

func UpdateLastUsed(ctx context.Context, apikeyId string, identifier string) error {
	mongoID, err := primitive.ObjectIDFromHex(apikeyId)
	if err != nil {
		return fmt.Errorf("could not convert ID: %v", err)
	}

	filter := bson.M{"_id": mongoID, "identifier": identifier}
	update := bson.M{"$set": bson.M{"lastUsed": time.Now()}}

	updateResult, err := mongodb.UpdateOne(ctx, collectionName, filter, update)

	if err != nil {
		return err
	}

	if updateResult.MatchedCount == 0 {
		return fmt.Errorf("could not update object")
	}

	return nil
}
