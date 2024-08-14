package mongodb

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (rc MongodbCon) Aggregate(ctx context.Context, col string, query []bson.M, value interface{}) error {
	db := rc.GetMongoDb()
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

func (rc MongodbCon) UpdateOne(ctx context.Context, col string, filter bson.M, update bson.M) (mongo.UpdateResult, error) {
	db := rc.GetMongoDb()
	updateResult, err := db.Collection(col).UpdateOne(ctx, filter, update)
	if err != nil {
		return mongo.UpdateResult{}, fmt.Errorf("could not updateOne object: %v", err)
	}

	return *updateResult, nil
}

func (rc MongodbCon) InsertOne(ctx context.Context, col string, input interface{}) (mongo.InsertOneResult, error) {
	db := rc.GetMongoDb()
	insertResult, err := db.Collection(col).InsertOne(ctx, input)
	if err != nil {
		return mongo.InsertOneResult{}, fmt.Errorf("could not insertOne object: %v", err)
	}

	return *insertResult, nil
}

func (rc MongodbCon) UpsertOne(ctx context.Context, col string, filter bson.M, update interface{}) (mongo.UpdateResult, error) {
	db := rc.GetMongoDb()

	upsert := true
	opts := &options.UpdateOptions{
		Upsert: &upsert,
	}
	updateResult, err := db.Collection(col).UpdateOne(ctx, filter, update, opts)

	if err != nil {
		return mongo.UpdateResult{}, fmt.Errorf("could not updateOne object: %v", err)
	}

	return *updateResult, nil
}

func (rc MongodbCon) FindOne(ctx context.Context, col string, query bson.M, value interface{}) error {
	db := rc.GetMongoDb()
	err := db.Collection(col).FindOne(ctx, query).Decode(value)
	if err != nil {
		return fmt.Errorf("could not findOne object: %v", err)
	}

	return nil
}

func (rc MongodbCon) DeleteOne(ctx context.Context, col string, query bson.M) (mongo.DeleteResult, error) {
	db := rc.GetMongoDb()
	result, err := db.Collection(col).DeleteOne(ctx, query)
	if err != nil {
		return mongo.DeleteResult{}, fmt.Errorf("could not insertOne object: %v", err)
	}
	return *result, nil
}

func (rc MongodbCon) Count(ctx context.Context, col string) (int64, error) {
	db := rc.GetMongoDb()
	opts := options.Count().SetHint("_id_")
	count, err := db.Collection(col).CountDocuments(ctx, bson.D{}, opts)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (rc MongodbCon) CountWithFilter(ctx context.Context, col string, filter interface{}) (int64, error) {
	db := rc.GetMongoDb()
	opts := options.Count().SetHint("_id_")
	count, err := db.Collection(col).CountDocuments(ctx, filter, opts)

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (rc MongodbCon) CountWithQuery(ctx context.Context, col string, query any) (int, error) {
	db := rc.GetMongoDb()
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

func (rc MongodbCon) GenerateAggregateQuery(rorResourceQuery *rorresources.ResourceQuery) []bson.M {
	query := make([]bson.M, 0)
	match := bson.M{}
	if rorResourceQuery == nil {
		return query
	}
	// Add filters

	if !rorResourceQuery.VersionKind.Empty() {
		apiversion, kind := rorResourceQuery.VersionKind.ToAPIVersionAndKind()
		if apiversion != "" {
			match["apiversion"] = apiversion
		}
		if kind != "" {
			match["kind"] = kind
		}
	}

	if len(rorResourceQuery.Uids) > 0 {
		match["uid"] = bson.M{"$in": rorResourceQuery.Uids}
	}
	query = append(query, bson.M{"$match": match})

	// Add sorting
	sort := bson.M{}
	if len(rorResourceQuery.Order) != 0 {
		sort["rorResourceQuery.SortBy"] = 1
	} else {
		sort["uid"] = 1
	}
	query = append(query, bson.M{"$sort": sort})
	// Add projection
	if len(rorResourceQuery.Fields) != 0 {
		project := bson.M{}
		for _, field := range rorResourceQuery.Fields {
			project[field] = 1
		}
		query = append(query, bson.M{"$project": project})
	}
	stringhelper.PrettyprintStruct(query)
	return query
}
