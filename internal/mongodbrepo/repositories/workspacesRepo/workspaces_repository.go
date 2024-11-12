package workspacesrepo

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

	aclrepo "github.com/NorskHelsenett/ror/internal/acl/repositories"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CollectionName         = "workspaces"
	ClustersCollectionName = "clusters"
)

func GetAllByIdentity(ctx context.Context) (*[]apicontracts.Workspace, error) {
	db := mongodb.GetMongoDb()
	accessLists := aclrepo.GetACL2ByIdentityQuery(ctx, aclmodels.AclV2QueryAccessScope{Scope: aclmodels.Acl2ScopeCluster})
	accessQuery := mongoHelper.CreateClusterACLFilter(accessLists)

	workspacesQuery := []bson.M{
		accessQuery,
		{
			"$project": bson.M{
				"workspaceid": 1,
			},
		},
		{
			"$sort": bson.M{"workspaceid": 1},
		},
		{"$group": bson.M{
			"_id": "$workspaceid",
		},
		},
		{
			"$lookup": bson.M{
				"as":           "workspaces",
				"foreignField": "_id",
				"from":         "workspaces",
				"localField":   "_id",
			},
		},
		{
			"$lookup": bson.M{
				"as":           "datacenters",
				"foreignField": "_id",
				"from":         "datacenters",
				"localField":   "workspaces.datacenterid",
			},
		},
		{
			"$project": bson.M{
				"_id":          "$_id",
				"name":         bson.M{"$first": "$workspaces.name"},
				"datacenterid": bson.M{"$first": "$datacenters._id"},
				"datacenter":   bson.M{"$first": "$datacenters"},
			},
		},
	}

	results, err := db.Collection(ClustersCollectionName).Aggregate(ctx, workspacesQuery)
	if err != nil {
		return nil, fmt.Errorf("could not fetch workspaces: %v", err)
	}
	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)
	var workspaces []apicontracts.Workspace = make([]apicontracts.Workspace, 0)

	if results.RemainingBatchLength() == 0 {
		return &workspaces, nil
	}

	for results.Next(ctx) {
		var singleWorkspace apicontracts.Workspace
		if err = results.Decode(&singleWorkspace); err != nil {
			return nil, fmt.Errorf("could not fetch workspace: %v", err)
		}
		workspaces = append(workspaces, singleWorkspace)
	}

	return &workspaces, nil
}

func GetTotalCount(ctx context.Context) (int64, error) {
	db := mongodb.GetMongoDb()
	opts := options.Count().SetHint("_id_")
	workspacesCount, err := db.Collection(CollectionName).CountDocuments(ctx, bson.D{}, opts)
	if err != nil {
		return 0, nil
	}

	return workspacesCount, nil
}

func GetByName(ctx context.Context, workspaceName string) (*apicontracts.Workspace, error) {
	db := mongodb.GetMongoDb()
	workspacesQuery := []bson.M{
		{
			"$match": bson.M{
				"name": workspaceName,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "datacenters",
				"localField":   "datacenterid",
				"foreignField": "_id",
				"as":           "datacenters",
			},
		},
		{
			"$set": bson.M{
				"datacenter": bson.M{
					"$first": "$datacenters",
				},
			},
		},
	}

	workspacesCursor, err := db.Collection(CollectionName).Aggregate(ctx, workspacesQuery)
	if err != nil {
		return nil, nil
	}

	if workspacesCursor.RemainingBatchLength() == 0 || workspacesCursor.RemainingBatchLength() > 1 {
		return nil, nil
	}

	var workspace apicontracts.Workspace
	if workspacesCursor.Next(ctx) {
		if err = workspacesCursor.Decode(&workspace); err != nil {
			return nil, errors.New("could not find workspace")
		}
	}
	accessLists := aclrepo.GetACL2ByIdentityQuery(ctx, aclmodels.AclV2QueryAccessScope{Scope: aclmodels.Acl2ScopeCluster})
	accessQuery := mongoHelper.CreateClusterACLFilter(accessLists)

	clusterQuery := []bson.M{
		accessQuery,
		{
			"$lookup": bson.M{
				"from":         "workspaces",
				"localField":   "workspaceid",
				"foreignField": "_id",
				"as":           "workspaces",
			},
		},
		{
			"$set": bson.M{
				"workspace": bson.M{
					"$first": "$workspaces",
				},
			},
		},
		{
			"$match": bson.M{
				"workspace.name": workspaceName,
			},
		},
	}

	clusterCursor, err := db.Collection(ClustersCollectionName).Aggregate(ctx, clusterQuery)
	if err != nil {
		return nil, errors.New("could not fetch all workspaces")
	}

	if clusterCursor.RemainingBatchLength() == 0 {
		return nil, nil
	}

	var clusterPrimitivMap []bson.M
	if err = clusterCursor.All(ctx, &clusterPrimitivMap); err != nil {
		return nil, nil
	}

	workspaceIds := make([]string, 0)
	for i := 0; i < len(clusterPrimitivMap); i++ {
		c := clusterPrimitivMap[i]
		workspace := c["workspace"].(primitive.M)

		id := workspace["_id"].(primitive.ObjectID)
		workspaceIds = append(workspaceIds, id.Hex())
	}

	for k := 0; k < len(workspaceIds); k++ {
		id := workspaceIds[k]
		if workspace.ID == id {
			break
		}
	}

	defer func(clusterCursor *mongo.Cursor, ctx context.Context) {
		_ = clusterCursor.Close(ctx)
	}(clusterCursor, ctx)
	defer func(workspacesCursor *mongo.Cursor, ctx context.Context) {
		_ = workspacesCursor.Close(ctx)
	}(workspacesCursor, ctx)

	return &workspace, nil
}

func FindByName(ctx context.Context, name string) (*apicontracts.Workspace, error) {
	db := mongodb.GetMongoDb()
	var wsResult mongoTypes.MongoWorkspace
	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"name": name}).Decode(&wsResult); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		msg := "could not find workspace"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	if wsResult.Name == "" {
		return nil, nil
	}

	var mapped apicontracts.Workspace
	err := mapping.Map(wsResult, &mapped)
	if err != nil {
		return nil, nil
	}

	return &mapped, nil
}

func GetById(ctx context.Context, id string) (*apicontracts.Workspace, error) {
	if id == "" {
		return nil, nil
	}
	db := mongodb.GetMongoDb()
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("could not get object id from hex: %w", err)
	}

	var object mongoTypes.MongoWorkspace
	err = db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoId}).Decode(&object)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("could not find object: %w", err)
	}

	var mapped apicontracts.Workspace
	err = mapping.Map(object, &mapped)
	if err != nil {
		return nil, fmt.Errorf("could not map object: %w", err)
	}

	return &mapped, nil
}

func Create(ctx context.Context, workspaceInput *apicontracts.Workspace) (*apicontracts.Workspace, error) {
	db := mongodb.GetMongoDb()
	var mongoInput mongoTypes.MongoWorkspace
	err := mapping.Map(workspaceInput, &mongoInput)
	if err != nil {
		return nil, errors.New("could not create workspace")
	}

	insertResult, err := db.Collection(CollectionName).InsertOne(ctx, mongoInput)
	if err != nil {
		msg := "could not create workspace"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	var workspaceResult mongoTypes.MongoWorkspace
	if insertResult.InsertedID == nil {
		return nil, errors.New("could not create workspace")
	}

	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&workspaceResult); err != nil {
		rlog.Error("could not find workspace", err)
		return nil, errors.New("could not find workspace after creation")
	}

	var mapped apicontracts.Workspace
	err = mapping.Map(workspaceResult, &mapped)
	if err != nil {
		return nil, errors.New("could not map workspace")
	}

	return &mapped, nil
}

func Update(ctx context.Context, input *mongoTypes.MongoWorkspace, id string) (*mongoTypes.MongoWorkspace, *mongoTypes.MongoWorkspace, error) {
	db := mongodb.GetMongoDb()
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, nil, fmt.Errorf("could not convert id: %w", err)
	}

	var originalObject mongoTypes.MongoWorkspace
	originalSingleResult := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoId})
	err = originalSingleResult.Decode(&originalObject)
	if err != nil {
		rlog.Error("could not get original object for auditlog", err)
	}

	updateResult, err := db.Collection(CollectionName).ReplaceOne(ctx, bson.M{"_id": mongoId}, input)
	if err != nil {
		return nil, nil, fmt.Errorf("could not update object: %w", err)
	}

	if updateResult.MatchedCount == 0 {
		return nil, nil, fmt.Errorf("could not find object")
	}

	var updatedObject mongoTypes.MongoWorkspace
	err = db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoId}).Decode(&updatedObject)
	if err != nil {
		return nil, nil, fmt.Errorf("could not find object after creation: %w", err)
	}

	return &updatedObject, &originalObject, nil
}
