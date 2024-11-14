package datacenterRepo

import (
	"context"
	"errors"

	"github.com/NorskHelsenett/ror/internal/helpers/mapping"
	mongoHelper "github.com/NorskHelsenett/ror/internal/helpers/mongoHelper"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/mongoTypes"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

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
	CollectionName        = "datacenters"
	ClusterCollectionName = "clusters"
)

func GetAllByUser(ctx context.Context) (*[]apicontracts.Datacenter, error) {
	db := mongodb.GetMongoDb()
	var datacentersQuery []bson.M

	datacentersCursor, err := db.Collection(CollectionName).Aggregate(ctx, datacentersQuery)
	accessLists := aclrepo.GetACL2ByIdentityQuery(ctx, aclmodels.AclV2QueryAccessScope{Scope: aclmodels.Acl2ScopeCluster})
	accessQuery := mongoHelper.CreateClusterACLFilter(accessLists)

	empty := make([]apicontracts.Datacenter, 0)
	if err != nil {
		return &empty, nil
	}

	var datacentersUnfiltered []apicontracts.Datacenter
	if err := datacentersCursor.All(ctx, &datacentersUnfiltered); err != nil {
		return &empty, nil
	}

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
			"$lookup": bson.M{
				"from":         "datacenters",
				"localField":   "workspaces.datacenterid",
				"foreignField": "_id",
				"as":           "datacenters",
			},
		},
		{
			"$project": bson.M{
				"datacenter": bson.M{
					"$arrayElemAt": bson.A{
						"$datacenters",
						0,
					},
				},
			},
		},
	}

	clusterCursor, err := db.Collection(ClusterCollectionName).Aggregate(ctx, clusterQuery)
	if err != nil {
		return nil, errors.New("could not fetch all datacenters")
	}

	if clusterCursor.RemainingBatchLength() == 0 {
		return &empty, nil
	}

	var clusterPrimitivMap []bson.M
	if err = clusterCursor.All(ctx, &clusterPrimitivMap); err != nil {
		return &empty, nil
	}

	defer func(datacentersCursor *mongo.Cursor, ctx context.Context) {
		_ = datacentersCursor.Close(ctx)
	}(datacentersCursor, ctx)
	defer func(clusterCursor *mongo.Cursor, ctx context.Context) {
		_ = clusterCursor.Close(ctx)
	}(clusterCursor, ctx)

	datacenterIds := make([]string, 0)
	for i := 0; i < len(clusterPrimitivMap); i++ {
		c := clusterPrimitivMap[i]
		if c["datacenter"] == nil {
			continue
		}
		datacenter := c["datacenter"].(primitive.M)

		id := datacenter["_id"].(primitive.ObjectID)
		datacenterIds = append(datacenterIds, id.Hex())
	}

	datacenters := make([]apicontracts.Datacenter, 0)
	for i := 0; i < len(datacentersUnfiltered); i++ {
		du := datacentersUnfiltered[i]
		for k := 0; k < len(datacenterIds); k++ {
			id := datacenterIds[k]
			if du.ID == id {
				datacenters = append(datacenters, du)
				break
			}
		}
	}

	return &datacenters, nil
}

func GetTotalCount(ctx context.Context) (int64, error) {
	db := mongodb.GetMongoDb()
	opts := options.Count().SetHint("_id_")
	datacenterCount, err := db.Collection(CollectionName).CountDocuments(ctx, bson.D{}, opts)
	if err != nil {
		return 0, err
	}

	return datacenterCount, nil
}

func FindByName(ctx context.Context, name string) (*apicontracts.Datacenter, error) {
	db := mongodb.GetMongoDb()
	var datacenterResult mongoTypes.MongoDatacenter
	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"name": name}).Decode(&datacenterResult); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		msg := "could not find datacenter"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	if datacenterResult.Name == "" {
		return nil, nil
	}

	var mapped apicontracts.Datacenter
	err := mapping.Map(datacenterResult, &mapped)
	if err != nil {
		return nil, nil
	}
	rlog.Debug("", rlog.Any("mapped", mapped))

	return &mapped, nil
}

func GetById(ctx context.Context, id string) (*apicontracts.Datacenter, error) {
	db := mongodb.GetMongoDb()
	var datacenterResult mongoTypes.MongoDatacenter
	mongoid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		msg := "could not convert datacenter id"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoid}).Decode(&datacenterResult); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		msg := "could not find datacenter"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	if datacenterResult.Name == "" {
		return nil, nil
	}

	var mapped apicontracts.Datacenter
	err = mapping.Map(datacenterResult, &mapped)
	if err != nil {
		return nil, nil
	}
	rlog.Debug("", rlog.Any("mapped", mapped))

	return &mapped, nil
}

func Create(ctx context.Context, datacenterInput *apicontracts.DatacenterModel, user *identitymodels.User) (*apicontracts.Datacenter, error) {
	db := mongodb.GetMongoDb()
	var mongoInput mongoTypes.MongoDatacenter
	err := mapping.Map(datacenterInput, &mongoInput)

	if err != nil {
		return nil, errors.New("could not parse input to mongo model")
	}

	insertResult, err := db.Collection(CollectionName).InsertOne(ctx, mongoInput)
	if err != nil {
		msg := "could not create datacenter"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	var datacenterResult mongoTypes.MongoDatacenter
	if insertResult.InsertedID == nil {

		return nil, errors.New("could not create datacenter")
	}

	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&datacenterResult); err != nil {
		rlog.Error("Could not find datacenter", err)
		return nil, errors.New("could not find datacenter after creation")
	}

	var mapped apicontracts.Datacenter
	_ = mapping.Map(datacenterResult, &mapped)

	return &mapped, nil
}

func Update(ctx context.Context, datacenterInput *apicontracts.DatacenterModel, user *identitymodels.User) (*apicontracts.Datacenter, error) {
	db := mongodb.GetMongoDb()
	var mongoInput mongoTypes.MongoDatacenter
	err := mapping.Map(datacenterInput, &mongoInput)

	mongoInput.ID = primitive.NilObjectID

	if err != nil {
		return nil, errors.New("could not parse input to mongo model")
	}

	mongoId, err := primitive.ObjectIDFromHex(datacenterInput.ID)
	if err != nil {
		return nil, errors.New("could not convert datacenter.ID")
	}

	var originalMongoDatacenter mongoTypes.MongoDatacenter
	originalSingleResult := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoId})
	err = originalSingleResult.Decode(&originalMongoDatacenter)
	if err != nil {
		rlog.Error("could not get original datacenter for auditlog", err)
	}

	updateResult, err := db.Collection(CollectionName).ReplaceOne(ctx, bson.M{"_id": mongoId}, mongoInput)
	if err != nil {
		msg := "could not update datacenter"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	var datacenterResult mongoTypes.MongoDatacenter
	if updateResult.ModifiedCount == 0 {
		return nil, errors.New("could not update datacenter")
	}

	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"_id": mongoId}).Decode(&datacenterResult); err != nil {
		rlog.Error("Could not find datacenter", err)
		return nil, errors.New("could not find datacenter after creation")
	}

	var mapped apicontracts.Datacenter
	_ = mapping.Map(datacenterResult, &mapped)

	return &mapped, nil
}
