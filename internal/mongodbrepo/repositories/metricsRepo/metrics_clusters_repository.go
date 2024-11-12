package metricsrepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/internal/helpers/mapping"
	mongoHelper "github.com/NorskHelsenett/ror/internal/helpers/mongoHelper"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"
	"go.mongodb.org/mongo-driver/mongo"

	aclrepo "github.com/NorskHelsenett/ror/internal/acl/repositories"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"

	"go.mongodb.org/mongo-driver/bson"
)

func GetForClusters(ctx context.Context) (*apicontracts.MetricList, error) {
	db := mongodb.GetMongoDb()
	groupQuery := getGroupBaseQuery()
	groupQuery["_id"] = "$clusterid"

	accessLists := aclrepo.GetACL2ByIdentityQuery(ctx, aclmodels.AclV2QueryAccessScope{Scope: aclmodels.Acl2ScopeCluster})
	accessQuery := mongoHelper.CreateClusterACLFilter(accessLists)

	queryCount := []bson.M{
		accessQuery,
		{"$group": groupQuery}, {"$sort": bson.M{"id": 1}},
	}

	results, err := db.Collection(CollectionName).Aggregate(ctx, queryCount)
	if err != nil {
		return nil, errors.New("could not fetch metrics for clusters")
	}

	if results.RemainingBatchLength() == 0 {
		return nil, nil
	}

	response := apicontracts.MetricList{}

	var acc []bson.M
	if err = results.All(ctx, &acc); err != nil {
		return nil, errors.New("could not fetch metrics for clusters")
	}

	if len(acc) == 0 {
		return nil, errors.New("could not fetch metrics for clusters")
	}

	for i := 0; i < len(acc); i++ {
		data := acc[i]
		metric := GetMetricFromPrimitivM(data)
		item := apicontracts.MetricItem{
			Id:      fmt.Sprint(data["_id"]),
			Metrics: metric,
		}

		response.Items = append(response.Items, item)
	}

	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	return &response, nil
}

func GetForClustersByWorkspace(ctx context.Context, workspaceName string) (*apicontracts.MetricList, error) {
	db := mongodb.GetMongoDb()
	groupQuery := getGroupBaseQuery()
	groupQuery["_id"] = "$clusterid"

	accessLists := aclrepo.GetACL2ByIdentityQuery(ctx, aclmodels.AclV2QueryAccessScope{Scope: aclmodels.Acl2ScopeCluster})
	accessQuery := mongoHelper.CreateClusterACLFilter(accessLists)

	queryCount := []bson.M{
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
		{"$match": bson.M{"workspace.name": workspaceName}},
		{"$group": groupQuery},
		{"$sort": bson.M{"id": 1}},
	}

	results, err := db.Collection(CollectionName).Aggregate(ctx, queryCount)
	if err != nil {
		return nil, errors.New("could not fetch metrics for clusters by workspace")
	}

	if results.RemainingBatchLength() == 0 {
		return nil, nil
	}

	response := apicontracts.MetricList{}

	var acc []bson.M
	if err = results.All(ctx, &acc); err != nil {
		return nil, errors.New("could not fetch metrics for clusters by workspace")
	}

	if len(acc) == 0 {
		return nil, errors.New("could not fetch metrics for clusters by workspace")
	}

	for i := 0; i < len(acc); i++ {
		data := acc[i]
		metric := GetMetricFromPrimitivM(data)
		item := apicontracts.MetricItem{
			Id:      fmt.Sprint(data["_id"]),
			Metrics: metric,
		}

		response.Items = append(response.Items, item)
	}

	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	return &response, nil
}

func GetForClusterid(ctx context.Context, clusterId string) (*apicontracts.MetricItem, error) {
	db := mongodb.GetMongoDb()
	groupQuery := getGroupBaseQuery()
	groupQuery["_id"] = "$clusterid"

	accessLists := aclrepo.GetACL2ByIdentityQuery(ctx, aclmodels.AclV2QueryAccessScope{Scope: aclmodels.Acl2ScopeCluster})
	accessQuery := mongoHelper.CreateClusterACLFilter(accessLists)

	queryCount := []bson.M{
		accessQuery,
		{"$match": bson.M{"clusterid": clusterId}},
		{"$group": groupQuery},
		{"$sort": bson.M{"id": 1}},
	}

	results, err := db.Collection(CollectionName).Aggregate(ctx, queryCount)
	if err != nil {
		return nil, errors.New("could not fetch metrics for clusterid")
	}

	if results.RemainingBatchLength() == 0 {
		return nil, nil
	}

	response := apicontracts.MetricItem{}

	var acc []bson.M
	if err = results.All(ctx, &acc); err != nil {
		return nil, errors.New("could not fetch metrics for clusterid")
	}

	if len(acc) > 1 {
		return nil, errors.New("could not fetch metrics for clusterid")
	}

	data := acc[0]
	metric := GetMetricFromPrimitivM(data)
	response.Id = fmt.Sprint(data["_id"])
	response.Metrics = metric

	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	return &response, nil
}

func ForClustersByProperty(ctx context.Context, property string) (*apicontracts.MetricsCustom, error) {
	db := mongodb.GetMongoDb()
	groupQuery := bson.M{
		"$group": bson.M{
			"_id":   "$" + property,
			"count": bson.M{"$sum": 1},
		},
	}

	accessLists := aclrepo.GetACL2ByIdentityQuery(ctx, aclmodels.AclV2QueryAccessScope{Scope: aclmodels.Acl2ScopeCluster})
	accessQuery := mongoHelper.CreateClusterACLFilter(accessLists)

	query := []bson.M{
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
			"$lookup": bson.M{
				"from":         "datacenters",
				"localField":   "workspaces.datacenterid",
				"foreignField": "_id",
				"as":           "datacenters",
			},
		},
		{
			"$set": bson.M{
				"workspace": bson.M{
					"datacenter": bson.M{
						"$first": "$datacenters",
					},
				},
			},
		},
		groupQuery,
		{"$sort": bson.M{"_id": 1}},
	}

	results, err := db.Collection(CollectionName).Aggregate(ctx, query)
	if err != nil {
		return nil, errors.New("could not fetch metrics for clusterid")
	}

	if results.RemainingBatchLength() == 0 {
		return nil, nil
	}

	var acc []bson.M
	if err = results.All(ctx, &acc); err != nil {
		return nil, errors.New("could not fetch metrics for clusterid")
	}

	if len(acc) < 1 {
		return nil, nil
	}

	data := make([]apicontracts.MetricsCustomItem, 0)
	for i := 0; i < len(acc); i++ {
		item := acc[i]
		text, _ := item["_id"].(string)
		value, _ := mapping.InterfaceToInt64(item["count"])
		data = append(data, apicontracts.MetricsCustomItem{
			Text:  text,
			Value: value,
		})
	}

	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	result := apicontracts.MetricsCustom{
		Data: data,
	}

	return &result, nil
}
