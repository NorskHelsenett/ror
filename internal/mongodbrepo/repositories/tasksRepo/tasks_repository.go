package tasksrepo

import (
	"context"
	"errors"
	"fmt"
	mongoHelper "github.com/NorskHelsenett/ror/internal/helpers/mongoHelper"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionName = "tasks"
)

func FindOne(ctx context.Context, property string, propertyValue string) (*apicontracts.Task, error) {
	db := mongodb.GetMongoDb()
	var taskResult apicontracts.Task
	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{property: propertyValue}).Decode(&taskResult); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		msg := "could not find task"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	if taskResult.Name == "" {
		return nil, nil
	}

	return &taskResult, nil
}

func GetById(ctx context.Context, taskId string) (*apicontracts.Task, error) {
	db := mongodb.GetMongoDb()
	id, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return nil, fmt.Errorf("invalid task id: %v", err)
	}

	var result apicontracts.Task
	if err := db.Collection(CollectionName).
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&result); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("could not find task: %v", err)
	}

	return &result, nil
}

func GetByFilter(ctx context.Context, filter *apicontracts.Filter) (*apicontracts.PaginatedResult[apicontracts.Task], error) {
	db := mongodb.GetMongoDb()
	aggregatePipeline := mongoHelper.CreateAggregationPipeline(filter, apicontracts.SortMetadata{SortField: "name", SortOrder: 1}, []string{"name"})

	var (
		query []bson.M
	)
	query = append(query, aggregatePipeline...)
	var totalCountQuery []bson.M

	configCollection := db.Collection(CollectionName)
	results, err := db.Collection(CollectionName).Aggregate(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("could not fetch tasks: %v", err)
	}

	totalCountResult, err := configCollection.Aggregate(ctx, totalCountQuery)
	if err != nil {
		return nil, fmt.Errorf("could not fetch tasks: %v", err)
	}

	var totalCountAcc []bson.M
	if err = totalCountResult.All(ctx, &totalCountAcc); err != nil {
		return nil, fmt.Errorf("could not fetch total count for tasks: %v", err)
	}

	totalCount := len(totalCountAcc)

	//reading from the db in an optimal way
	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)
	if results.RemainingBatchLength() == 0 {
		emptyResult := apicontracts.PaginatedResult[apicontracts.Task]{}
		return &emptyResult, nil
	}

	tasks := make([]apicontracts.Task, 0)
	paginatedResult := apicontracts.PaginatedResult[apicontracts.Task]{}
	for results.Next(ctx) {
		var element apicontracts.Task
		if err = results.Decode(&element); err != nil {
			return nil, fmt.Errorf("could not fetch tasks: %v", err)
		}

		tasks = append(tasks, element)
	}

	paginatedResult.Data = tasks
	paginatedResult.DataCount = int64(len(tasks))
	paginatedResult.Offset = int64(filter.Skip)
	paginatedResult.TotalCount = int64(totalCount)

	return &paginatedResult, nil
}

func Create(ctx context.Context, task *apicontracts.Task) (*apicontracts.Task, error) {
	db := mongodb.GetMongoDb()

	insertResult, err := db.Collection(CollectionName).InsertOne(ctx, task)
	if err != nil {
		msg := "could not create task"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	var taskResult apicontracts.Task
	if insertResult.InsertedID == nil {
		return nil, errors.New("could not create task")
	}

	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&taskResult); err != nil {
		rlog.Error("Could not find task", err)
		return nil, errors.New("could not find task after creation")
	}

	return &taskResult, nil
}

func GetAll(ctx context.Context) (*[]apicontracts.Task, error) {
	db := mongodb.GetMongoDb()
	var query []bson.M
	cursor, err := db.Collection(CollectionName).Aggregate(ctx, query)
	if err != nil {
		return nil, errors.New("could not fetch all tasks")
	}

	tasks := make([]apicontracts.Task, 0)
	if cursor.RemainingBatchLength() == 0 {
		return &tasks, nil
	}

	for cursor.Next(ctx) {
		var task apicontracts.Task
		if err = cursor.Decode(&task); err != nil {
			rlog.Error("could not parse task to api contract definition", err)
			continue
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil
}

func Update(ctx context.Context, taskId string, taskInput apicontracts.Task) (*apicontracts.Task, *apicontracts.Task, error) {
	db := mongodb.GetMongoDb()
	mongoId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return nil, nil, fmt.Errorf("could not convert taskId: %v", err)
	}

	var originalTask apicontracts.Task
	originalSingleResult := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoId})
	err = originalSingleResult.Decode(&originalTask)
	if err != nil {
		rlog.Error("could not get original task for auditlog", err)
	}

	updateResult, err := db.Collection(CollectionName).ReplaceOne(ctx, bson.M{"_id": mongoId}, taskInput)
	if err != nil {
		return nil, nil, fmt.Errorf("could not update task: %v", err)
	}

	if updateResult.MatchedCount == 0 {
		return nil, nil, fmt.Errorf("could not find task")
	}

	if updateResult.ModifiedCount == 0 {
		return nil, nil, fmt.Errorf("could not update task")
	}

	var taskResult apicontracts.Task
	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoId}).Decode(&taskResult); err != nil {
		return nil, nil, fmt.Errorf("could not find task after creation: %v", err)
	}

	return &taskResult, &originalTask, nil
}

func Delete(ctx context.Context, taskId string) (bool, *apicontracts.Task, error) {
	db := mongodb.GetMongoDb()
	mongoId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return false, nil, fmt.Errorf("could not convert taskId: %v", err)
	}

	var originalTask apicontracts.Task
	originalSingleResult := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoId})
	err = originalSingleResult.Decode(&originalTask)
	if err != nil {
		rlog.Error("could not get original task for auditlog", err)
	}

	deleteResult, err := db.Collection(CollectionName).DeleteOne(ctx, bson.M{"_id": mongoId})
	if err != nil {
		return false, nil, fmt.Errorf("could not delete task: %v", err)
	}

	if deleteResult.DeletedCount == 0 {
		return false, nil, fmt.Errorf("could not delete task")
	}

	return true, &originalTask, nil
}
