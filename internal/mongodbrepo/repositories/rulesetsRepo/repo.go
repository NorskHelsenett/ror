package rulesetsRepo

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/messages"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"
)

func Create(ctx context.Context, set *messages.RulesetModel) error {
	db := mongodb.GetMongoDb()
	coll := db.Collection("messagerulesets")

	set.Resources = make([]messages.RulesetResourceModel, 0)

	result, err := coll.InsertOne(ctx, set)
	if err != nil {
		return fmt.Errorf("could not insert cluster: %v", err)
	}

	var newBoard messages.RulesetModel
	if err := coll.FindOne(ctx, bson.D{{Key: "_id", Value: result.InsertedID}}).Decode(&newBoard); err != nil {
		return err
	}

	set.ID = newBoard.ID

	return nil
}

func FindInternal(ctx context.Context) (*messages.RulesetModel, error) {
	db := mongodb.GetMongoDb()
	coll := db.Collection("messagerulesets")

	set := new(messages.RulesetModel)

	filter := bson.D{
		{Key: "identity.id", Value: "internal-primary"},
		{Key: "identity.type", Value: messages.RulesetIdentityTypeInternal},
	}

	if err := coll.FindOne(ctx, filter).Decode(set); err != nil {
		return nil, err
	}

	return set, nil
}

func FindAll(ctx context.Context) ([]*messages.RulesetModel, error) {
	db := mongodb.GetMongoDb()
	coll := db.Collection("messagerulesets")

	var messagerulesets []*messages.RulesetModel

	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &messagerulesets); err != nil {
		return nil, err
	}

	return messagerulesets, nil
}

func FindCluster(ctx context.Context, clusterId string) (*messages.RulesetModel, error) {
	db := mongodb.GetMongoDb()
	coll := db.Collection("messagerulesets")

	set := new(messages.RulesetModel)

	filter := bson.D{
		{Key: "identity.id", Value: clusterId},
		{Key: "identity.type", Value: messages.RulesetIdentityTypeCluster},
	}

	if err := coll.FindOne(ctx, filter).Decode(set); err != nil {
		return nil, err
	}

	return set, nil
}

func FindById(ctx context.Context, hexId string) (*messages.RulesetModel, error) {
	db := mongodb.GetMongoDb()
	coll := db.Collection("messagerulesets")

	id, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return nil, err
	}

	set := new(messages.RulesetModel)
	if err := coll.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(set); err != nil {
		return nil, err
	}

	return set, nil
}

func AddResource(ctx context.Context, set *messages.RulesetModel, resource *messages.RulesetResourceModel) error {
	db := mongodb.GetMongoDb()
	coll := db.Collection("messagerulesets")

	resource.Rules = make([]messages.RulesetRuleModel, 0)

	id, err := primitive.ObjectIDFromHex(set.ID)
	if err != nil {
		return err
	}

	filter := bson.D{
		{Key: "_id", Value: id},
	}

	update := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "resources", Value: resource},
		}},
	}

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

func PullResource(ctx context.Context, set *messages.RulesetModel, resourceId string) error {
	db := mongodb.GetMongoDb()
	coll := db.Collection("messagerulesets")

	id, err := primitive.ObjectIDFromHex(set.ID)
	if err != nil {
		return err
	}

	filter := bson.D{
		{Key: "_id", Value: id},
	}

	update := bson.D{
		{Key: "$pull", Value: bson.D{
			{Key: "resources", Value: bson.D{
				{Key: "id", Value: resourceId},
			}},
		}},
	}

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

func AddResourceRule(ctx context.Context, set *messages.RulesetModel, resource *messages.RulesetResourceModel, rule *messages.RulesetRuleModel) error {
	db := mongodb.GetMongoDb()
	coll := db.Collection("messagerulesets")

	setId, err := primitive.ObjectIDFromHex(set.ID)
	if err != nil {
		return err
	}

	filter := bson.D{
		{Key: "_id", Value: setId},
	}

	update := bson.M{
		"$push": bson.M{
			"resources.$[resource].rules": rule,
		},
	}

	// Specify the array filter for the positional $ operator to identify the resource
	arrayFilters := options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"resource.id": resource.Id},
		},
	}

	// Set the array filters option
	opts := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
	}

	if _, err := coll.UpdateOne(ctx, filter, update, &opts); err != nil {
		return err
	}

	return nil
}

func PullResourceRule(ctx context.Context, set *messages.RulesetModel, resource *messages.RulesetResourceModel, ruleId string) error {
	db := mongodb.GetMongoDb()
	coll := db.Collection("messagerulesets")

	setId, err := primitive.ObjectIDFromHex(set.ID)
	if err != nil {
		return err
	}

	filter := bson.D{
		{Key: "_id", Value: setId},
	}

	update := bson.M{
		"$pull": bson.M{
			"resources.$[resource].rules": bson.M{
				"id": ruleId,
			},
		},
	}

	arrayFilters := options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"resource.id": resource.Id},
		},
	}

	opts := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
	}

	if _, err := coll.UpdateOne(ctx, filter, update, &opts); err != nil {
		return err
	}

	return nil
}
