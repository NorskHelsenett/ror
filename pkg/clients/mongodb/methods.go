package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbQuery []bson.M

func (m MongodbQuery) PrettyPrint() {
	var prettyDocs []bson.M
	for _, doc := range m {
		bsonDoc, err := bson.Marshal(doc)
		if err != nil {
			rlog.Error("error marshalling bson", err)
		}
		var prettyDoc bson.M
		err = bson.Unmarshal(bsonDoc, &prettyDoc)
		if err != nil {
			rlog.Error("error unmarshalling bson", err)
		}
		prettyDocs = append(prettyDocs, prettyDoc)
	}
	prettyJSON, err := json.MarshalIndent(prettyDocs, "", "  ")
	if err != nil {
		rlog.Error("error marshalling json", err)
	}

	fmt.Println(string(prettyJSON))
}

func (m MongodbQuery) MongoshPrint(collection string) {
	var prettyDocs []bson.M
	for _, doc := range m {
		bsonDoc, err := bson.Marshal(doc)
		if err != nil {
			rlog.Error("error marshalling bson", err)
		}
		var prettyDoc bson.M
		err = bson.Unmarshal(bsonDoc, &prettyDoc)
		if err != nil {
			rlog.Error("error unmarshalling bson", err)
		}
		prettyDocs = append(prettyDocs, prettyDoc)
	}
	prettyJSON, err := json.Marshal(prettyDocs)
	if err != nil {
		rlog.Error("error marshalling json", err)
	}

	fmt.Printf("db.%s.aggregate(%s)", collection, string(prettyJSON))
}

func NewMongodbQuery(input []bson.M) MongodbQuery {
	return MongodbQuery(input)
}

func (rc MongodbCon) Aggregate(ctx context.Context, col string, query []bson.M, value interface{}) error {
	db := rc.GetMongoDb()
	resourceCollection := db.Collection(col)
	results, err := resourceCollection.Aggregate(ctx, query)
	if err != nil {
		return fmt.Errorf("could not decode mongo document: %w", err)
	}
	defer func(cursor *mongo.Cursor, closeCtx context.Context) {
		_ = cursor.Close(closeCtx)
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
	defer func(cursor *mongo.Cursor, closeCtx context.Context) {
		_ = cursor.Close(closeCtx)
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
			match["typemeta.apiversion"] = apiversion
		}
		if kind != "" {
			match["typemeta.kind"] = kind
		}
	}

	if len(rorResourceQuery.Uids) > 0 {
		match["uid"] = bson.M{"$in": rorResourceQuery.Uids}
	}

	if len(rorResourceQuery.OwnerRefs) > 0 {
		match["rormeta.ownerref"] = bson.M{"$in": rorResourceQuery.OwnerRefs}
	}

	if len(rorResourceQuery.Filters) == 1 {
		addMatchFilter(rorResourceQuery.Filters[0], match)
	} else if len(rorResourceQuery.Filters) > 1 {
		filterCount := map[string]int{}
		for _, filter := range rorResourceQuery.Filters {
			filterCount[filter.Field]++
		}

		for key, value := range filterCount {
			if value == 1 {
				for _, filter := range rorResourceQuery.Filters {
					if filter.Field == key {
						addMatchFilter(filter, match)
					}
				}
			} else {
				var filterList []string
				for _, filter := range rorResourceQuery.Filters {
					if filter.Field == key {
						filterList = append(filterList, filter.Value)
					}
				}
				match[key] = bson.M{"$in": filterList}
			}
		}
	}
	query = append(query, bson.M{"$match": match})

	// Add sorting
	sortaggregate := bson.M{}
	if len(rorResourceQuery.Order) != 0 {
		for _, orderline := range rorResourceQuery.GetOrderSorted() {
			if orderline.Descending {
				sortaggregate[orderline.Field] = -1
			} else {
				sortaggregate[orderline.Field] = 1
			}
		}

	} else {
		sortaggregate["metadata.name"] = 1
	}
	query = append(query, bson.M{"$sort": sortaggregate})
	// Add projection
	if len(rorResourceQuery.Fields) != 0 {
		project := bson.M{}
		project["metadata"] = 1
		project["rormeta"] = 1
		project["typemeta"] = 1
		for _, field := range rorResourceQuery.Fields {
			project[field] = 1
		}
		query = append(query, bson.M{"$project": project})
	}

	// Add offset and limit
	if rorResourceQuery.Offset != 0 {
		query = append(query, bson.M{"$skip": rorResourceQuery.Offset})
	}

	if rorResourceQuery.Limit > 1000 {
		rorResourceQuery.Limit = 1000
	}

	if rorResourceQuery.Limit == 0 {
		query = append(query, bson.M{"$limit": 100})
	}

	if rorResourceQuery.Limit != -1 {
		query = append(query, bson.M{"$limit": rorResourceQuery.Limit})
	}

	NewMongodbQuery(query).PrettyPrint()
	return query
}

func addMatchFilter(filter rorresources.ResourceQueryFilter, match bson.M) {
	switch filter.Type {
	case rorresources.FilterTypeString:
		if filter.Operator == "eq" {
			match[filter.Field] = bson.M{"$eq": filter.Value}
		}
		if filter.Operator == "ne" {
			match[filter.Field] = bson.M{"$ne": filter.Value}
		}
		if filter.Operator == "regexp" {
			match[filter.Field] = bson.M{"$regex": filter.Value, "$options": "i"}
		}
	case rorresources.FilterTypeInt:
		if filter.Operator == "eq" {
			match[filter.Field] = bson.M{"$eq": filter.Value}
		}
		if filter.Operator == "gt" {
			match[filter.Field] = bson.M{"$gt": filter.Value}
		}
		if filter.Operator == "lt" {
			match[filter.Field] = bson.M{"$lt": filter.Value}
		}
		if filter.Operator == "ge" {
			match[filter.Field] = bson.M{"$gte": filter.Value}
		}
		if filter.Operator == "le" {
			match[filter.Field] = bson.M{"$lte": filter.Value}
		}
	case rorresources.FilterTypeBool:
		if filter.Operator == "eq" {
			boolfilter, err := strconv.ParseBool(filter.Value)
			if err == nil {
				match[filter.Field] = bson.M{"$eq": boolfilter}
			}
		}
		// case rorresources.FilterTypeTime:
		// 	format := "2006-01-02 15:04:05.999999999 -0700 MST"
		// 	timevalue, err := time.Parse(format, filter.Value)
		// 	if err == nil {
		// 		//timevalue := time.ParseTimestamps(filter.Value)
		// 		if filter.Operator == "eq" {
		// 			match[filter.Field] = bson.M{"$eq": timevalue}
		// 		}
		// 		if filter.Operator == "gt" {
		// 			match[filter.Field] = bson.M{"$gt": timevalue}
		// 		}
		// 		if filter.Operator == "lt" {
		// 			match[filter.Field] = bson.M{"$lt": timevalue}
		// 		}
		// 		if filter.Operator == "ge" {
		// 			match[filter.Field] = bson.M{"$gte": timevalue}
		// 		}
		// 		if filter.Operator == "le" {
		// 			match[filter.Field] = bson.M{"$lte": timevalue}
		// 		}
		// 	}
	}
}
