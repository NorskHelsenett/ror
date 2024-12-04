package clustersrepo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	clusterhelper "github.com/NorskHelsenett/ror/internal/helpers/clusterHelper"
	"github.com/NorskHelsenett/ror/internal/helpers/mapping"
	mongoHelper "github.com/NorskHelsenett/ror/internal/helpers/mongoHelper"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/mongoTypes"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	aclrepo "github.com/NorskHelsenett/ror/internal/acl/repositories"
	workspacesRepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/workspacesRepo"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionName = "clusters"
)

func GetByClusterId(ctx context.Context, clusterId string) (*apicontracts.Cluster, error) {
	db := mongodb.GetMongoDb()
	accessLists := aclrepo.GetACL2ByIdentityQuery(ctx, aclmodels.AclV2QueryAccessScope{Scope: aclmodels.Acl2ScopeCluster})
	accessQuery := mongoHelper.CreateClusterACLFilter(accessLists)

	query := []bson.M{
		accessQuery,
		{
			"$match": bson.M{
				"clusterid": clusterId,
			},
		},
		{
			"$project": bson.M{"state": 0},
		},
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
		{
			"$lookup": bson.M{
				"from":         "projects",
				"localField":   "metadata.projectid",
				"foreignField": "_id",
				"as":           "projects",
			},
		},
		{
			"$set": bson.M{
				"metadata": bson.M{
					"project": bson.M{
						"$first": "$projects",
					},
				},
			},
		},
	}

	mongoCol := db.Collection(CollectionName)
	results, err := mongoCol.Aggregate(ctx, query)
	if err != nil {
		return nil, errors.New("could not get cluster")
	}
	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	if results.RemainingBatchLength() == 0 {
		return nil, nil
	}

	var clusters []apicontracts.Cluster
	if err = results.All(ctx, &clusters); err != nil {
		return nil, errors.New("could not get error")
	}

	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	if len(clusters) > 1 {
		rlog.Error("could not get cluster", fmt.Errorf("multiple clusters with same id"), rlog.String("clusterid", clusterId))
		return nil, errors.New("could not get cluster")
	}

	cluster := clusters[0]
	clusterhelper.SetStatus(&cluster)

	return &cluster, nil
}

// GetByFilter Get cluster by filter  *apicontracts.Filter
func GetByFilter(ctx context.Context, filter *apicontracts.Filter) (*apicontracts.PaginatedResult[*apicontracts.Cluster], error) {
	db := mongodb.GetMongoDb()
	ctx, span := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "datacenterRepo.GetByFilter")
	defer span.End()

	_, span2 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "BuildQuery")
	defer span2.End()

	aggregationPipeline := mongoHelper.CreateAggregationPipeline(filter, apicontracts.SortMetadata{SortField: "clusterid", SortOrder: 1}, []string{"workspace", "workspace.datacenter"})

	accessLists := aclrepo.GetACL2ByIdentityQuery(ctx, aclmodels.AclV2QueryAccessScope{Scope: aclmodels.Acl2ScopeCluster})
	accessQuery := mongoHelper.CreateClusterACLFilter(accessLists)
	var query []bson.M
	var totalCountQuery []bson.M

	query = append(query, accessQuery)
	query = append(query, []bson.M{
		{
			"$lookup": bson.M{
				"from":         "workspaces",
				"localField":   "workspaceid",
				"foreignField": "_id",
				"as":           "workspaces",
			},
		},
		{
			"$project": bson.M{"state": 0},
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
		{
			"$lookup": bson.M{
				"from":         "projects",
				"localField":   "metadata.projectid",
				"foreignField": "_id",
				"as":           "projects",
			},
		},
		{
			"$set": bson.M{
				"metadata": bson.M{
					"project": bson.M{
						"$first": "$projects",
					},
				},
			},
		},
	}...)
	query = append(query, aggregationPipeline...)

	totalCountQuery = []bson.M{
		accessQuery,
	}

	totalCountQuery = append(totalCountQuery, bson.M{"$project": bson.M{"_id": 1}})
	span2.End()

	_, span3 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Run query data")
	defer span3.End()

	clusterCollection := db.Collection(CollectionName)

	results, err := db.Collection(CollectionName).Aggregate(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("could not fetch clusters: %v", err)
	}
	span3.End()

	_, span4 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Run query total")
	defer span4.End()

	totalCountResult, err := clusterCollection.Aggregate(ctx, totalCountQuery)
	if err != nil {
		return nil, fmt.Errorf("could not fetch clusters: %v", err)
	}

	var totalCountAcc []bson.M
	if err = totalCountResult.All(ctx, &totalCountAcc); err != nil {
		return nil, fmt.Errorf("could not fetch total count for clusters: %v", err)
	}

	totalCount := len(totalCountAcc)
	span4.End()

	_, span5 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Read data from db")
	defer span5.End()

	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)
	defer func(totalCountResult *mongo.Cursor, ctx context.Context) {
		_ = totalCountResult.Close(ctx)
	}(totalCountResult, ctx)

	if results.RemainingBatchLength() == 0 {
		emptyResult := apicontracts.PaginatedResult[*apicontracts.Cluster]{}
		return &emptyResult, nil
	}

	var clusters = make([]*apicontracts.Cluster, 0)
	paginatedResult := apicontracts.PaginatedResult[*apicontracts.Cluster]{}
	for results.Next(ctx) {
		var singleCluster apicontracts.Cluster
		if err = results.Decode(&singleCluster); err != nil {
			return nil, fmt.Errorf("could not fetch clusters: %v", err)
		}

		clusterhelper.SetStatus(&singleCluster)
		clusters = append(clusters, &singleCluster)
	}
	span5.End()

	paginatedResult.Data = clusters
	paginatedResult.DataCount = int64(len(clusters))
	paginatedResult.Offset = int64(filter.Skip)
	paginatedResult.TotalCount = int64(totalCount)

	return &paginatedResult, nil
}

func GetMetadata(ctx context.Context) (map[string][]string, error) {
	db := mongodb.GetMongoDb()
	accessLists := aclrepo.GetACL2ByIdentityQuery(ctx, aclmodels.AclV2QueryAccessScope{Scope: aclmodels.Acl2ScopeCluster})
	accessQuery := mongoHelper.CreateClusterACLFilter(accessLists)

	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, accessQuery)
	aggregationPipeline = append(aggregationPipeline, bson.M{
		"$group": bson.M{
			"_id": nil,
			"kubernetesVersions": bson.M{
				"$addToSet": "$versions.kubernetes",
			},
			"nhnToolingVersions": bson.M{
				"$addToSet": "$versions.nhntooling.version",
			},
			"nhnToolingBranches": bson.M{
				"$addToSet": "$versions.nhntooling.branch",
			},
		},
	})
	aggregationPipeline = append(aggregationPipeline, bson.M{
		"$project": bson.M{
			"_id":                false,
			"kubernetesVersions": "$kubernetesVersions",
			"nhnToolingVersions": "$nhnToolingVersions",
			"nhnToolingBranches": "$nhnToolingBranches",
		},
	})
	clusterCollection := db.Collection(CollectionName)

	cursor, err := clusterCollection.Aggregate(ctx, aggregationPipeline)
	if err != nil {
		return nil, fmt.Errorf("could not perform aggregation: %v", err)
	}

	if cursor.RemainingBatchLength() == 0 {
		return nil, nil
	}

	var metadataSlice []map[string][]string
	for cursor.Next(ctx) {
		var metadata map[string][]string
		err = cursor.Decode(&metadata)
		if err != nil {
			return nil, fmt.Errorf("could not decode mongo document: %v", err)
		}
		metadataSlice = append(metadataSlice, metadata)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)
	if len(metadataSlice) > 0 {
		return metadataSlice[0], nil
	}

	return nil, errors.New("missing metadata")
}

func GetByWorkspace(ctx context.Context,
	filter *apicontracts.Filter,
	workspaceName string) (*apicontracts.PaginatedResult[apicontracts.Cluster], error) {
	db := mongodb.GetMongoDb()
	bsonSort := bson.M{}
	for i := 0; i < len(filter.Sort); i++ {
		sort := filter.Sort[i]
		if sort.SortField == "" {
			continue
		}
		bsonSort[sort.SortField] = sort.SortOrder

	}

	if len(bsonSort) == 0 {
		bsonSort = bson.M{"clusterid": 1}
	}

	var queryCount []bson.M
	var query []bson.M

	accessLists := aclrepo.GetACL2ByIdentityQuery(ctx, aclmodels.AclV2QueryAccessScope{Scope: aclmodels.Acl2ScopeCluster})
	accessQuery := mongoHelper.CreateClusterACLFilter(accessLists)
	queryCount = []bson.M{
		accessQuery,
		{"$project": bson.M{"_id": 1}},
		{
			"$lookup": bson.M{
				"from":         "workspaces",
				"localField":   "workspaceid",
				"foreignField": "_id",
				"as":           "workspaces",
			},
		},
		{
			"$project": bson.M{"state": 0},
		},
		{
			"$set": bson.M{
				"workspace": bson.M{
					"$first": "$workspaces",
				},
			},
		},
		{"$match": bson.M{"workspace.name": workspaceName}},
	}
	query = []bson.M{
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
			"$project": bson.M{"state": 0},
		},
		{
			"$set": bson.M{
				"workspace": bson.M{
					"$first": "$workspaces",
				},
			},
		},
		{"$match": bson.M{"workspace.name": workspaceName}},
		{"$sort": bsonSort},
		{"$skip": filter.Skip},
		{"$limit": filter.Limit},
	}

	clusterCollection := db.Collection(CollectionName)
	countResult, err := clusterCollection.Aggregate(ctx, queryCount)
	if err != nil {
		return nil, errors.New("could not get cluster by workspace")
	}

	totalQueryCount := countResult.RemainingBatchLength()
	results, err := clusterCollection.Aggregate(ctx, query)
	if err != nil {
		return nil, errors.New("could not get cluster by workspace")
	}

	clusters := make([]apicontracts.Cluster, 0)

	//reading from the db in an optimal way
	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)
	if results.RemainingBatchLength() == 0 {
		return nil, nil
	}
	for results.Next(ctx) {
		var cluster apicontracts.Cluster
		if err = results.Decode(&cluster); err != nil {
			return nil, errors.New("could not get cluster by workspace")
		}

		clusterhelper.SetStatus(&cluster)
		clusters = append(clusters, cluster)
	}

	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)
	defer func(countResult *mongo.Cursor, ctx context.Context) {
		_ = countResult.Close(ctx)
	}(countResult, ctx)

	paginatedResult := apicontracts.PaginatedResult[apicontracts.Cluster]{}
	paginatedResult.Data = clusters
	paginatedResult.DataCount = int64(len(clusters))
	paginatedResult.Offset = int64(filter.Skip)
	paginatedResult.TotalCount = int64(totalQueryCount)

	return &paginatedResult, nil
}

func FindByClusterId(ctx context.Context, clusterId string) (*apicontracts.Cluster, error) {
	var clusterResult mongoTypes.MongoCluster
	db := mongodb.GetMongoDb()
	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"clusterid": clusterId}).Decode(&clusterResult); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		msg := "could not find cluster"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	if clusterResult.ClusterId == "" {
		return nil, nil
	}

	var mapped apicontracts.Cluster
	err := mapping.Map(clusterResult, &mapped)
	if err != nil {
		return nil, nil
	}

	return &mapped, nil
}

func FindByName(ctx context.Context, clusterName string) (*apicontracts.Cluster, error) {
	var clusterResult mongoTypes.MongoCluster
	db := mongodb.GetMongoDb()
	if err := db.Collection(CollectionName).FindOne(ctx, bson.M{"clustername": clusterName}).Decode(&clusterResult); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		msg := "could not find cluster"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	if clusterResult.ClusterId == "" {
		return nil, nil
	}

	var mapped apicontracts.Cluster
	err := mapping.Map(clusterResult, &mapped)
	if err != nil {
		return nil, nil
	}

	return &mapped, nil
}

func GetClusterIdByProjectId(ctx context.Context, projectId string) ([]*apicontracts.ClusterInfo, error) {
	db := mongodb.GetMongoDb()
	projectObjectId, _ := primitive.ObjectIDFromHex(projectId)

	var query []bson.M
	query = []bson.M{}

	query = append(query, bson.M{
		"$match": bson.M{
			"metadata.projectid": projectObjectId,
		},
	})

	bsonSort := bson.M{"clusterid": -1}
	query = append(query, bson.M{"$sort": bsonSort})

	query = append(query, bson.M{
		"$project": bson.M{
			"clusterid":   1,
			"clustername": 1,
			"metadata":    1,
			"environment": 1,
		},
	})

	mongoCol := db.Collection(CollectionName)
	results, err := mongoCol.Aggregate(ctx, query)
	if err != nil {
		return nil, errors.New("could not get cluster")
	}
	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	if results.RemainingBatchLength() == 0 {
		return nil, nil
	}

	var clusters []*apicontracts.ClusterInfo
	if err = results.All(ctx, &clusters); err != nil {
		return nil, errors.New("could not get error")
	}

	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	return clusters, nil
}

func Create(ctx context.Context, clusterInput *apicontracts.Cluster) error {
	db := mongodb.GetMongoDb()
	mongoInput, err := mapToMongo(clusterInput)
	if err != nil {
		return errors.New("could not map to mongo cluster")
	}

	_, err = db.Collection(CollectionName).InsertOne(ctx, mongoInput)
	if err != nil {
		msg := "could not create cluster"
		rlog.Error(msg, err)
		return errors.New(msg)
	}

	return nil
}

func Update(ctx context.Context, clusterInput *apicontracts.Cluster) error {
	ctx, span := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Repository: clustersrepo.Update")
	defer span.End()
	_, span1 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Get mongoclient")
	defer span1.End()
	db := mongodb.GetMongoDb()
	span1.End()

	_, span2 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Prepare data for mongo")
	defer span2.End()
	mongoInput, err := mapToMongo(clusterInput)
	if err != nil {
		return errors.New("could not map to mongo cluster")
	}

	if strings.Contains(clusterInput.WorkspaceId, "000000") || clusterInput.WorkspaceId == "" {
		workspaceNameArray := strings.Split(clusterInput.ClusterId, ".")
		if len(workspaceNameArray) == 2 {
			workspaceName := workspaceNameArray[1]
			workspace, _ := workspacesRepo.FindByName(ctx, workspaceName)

			if workspace != nil {
				workspaceId, _ := primitive.ObjectIDFromHex(workspace.ID)
				mongoInput.WorkspaceId = workspaceId
			} else {
				rlog.Error("could not update", fmt.Errorf("workspace is nil"))
			}
		} else {
			rlog.Error("could not update", fmt.Errorf("could not get workspace name"))
		}
	}
	span2.End()

	_, span3 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Update data in mongo")
	defer span3.End()
	result, err := db.Collection(CollectionName).UpdateOne(ctx, bson.M{"clusterid": clusterInput.ClusterId}, bson.M{"$set": mongoInput})
	if err != nil {
		return fmt.Errorf("could not update cluster with id: %s", clusterInput.ClusterId)
	}

	if result.MatchedCount != 1 {
		return fmt.Errorf("could not update cluster with id: %s", clusterInput.ClusterId)
	}
	span3.End()

	span.End()

	return nil
}

func mapToMongo(source *apicontracts.Cluster) (*mongoTypes.MongoCluster, error) {
	var mongoInput mongoTypes.MongoCluster
	err := mapping.Map(source, &mongoInput)
	if err != nil {
		return nil, errors.New("could not map to mongo cluster")
	}

	mongoInput.WorkspaceId, _ = primitive.ObjectIDFromHex(source.WorkspaceId)

	return &mongoInput, nil
}

func UpdateMetadata(ctx context.Context, input *apicontracts.ClusterMetadataModel, existing *apicontracts.Cluster) error {
	db := mongodb.GetMongoDb()
	var mongoInput mongoTypes.ClusterMetadata
	err := mapping.Map(input, &mongoInput)

	if err != nil {
		return fmt.Errorf("could not map data from cluster metadata model to cluster metadata for cluster: %s", existing.ClusterId)
	}

	mongoInput.ProjectID, _ = primitive.ObjectIDFromHex(input.ProjectID)

	filter := bson.M{"clusterid": existing.ClusterId}
	update := bson.M{"$set": bson.M{"metadata": mongoInput}}
	updateResult, err := db.Collection(CollectionName).UpdateOne(ctx, filter, update)
	if err != nil {
		msg := "could not set new metadata"
		rlog.Error(msg, err)
		return errors.New(msg)
	}

	if updateResult.MatchedCount == 1 {
		return nil
	}

	return errors.New("could not update metadata")
}

func GetControlPlaneMetadata(ctx context.Context) ([]apicontracts.ClusterControlPlaneMetadata, error) {
	db := mongodb.GetMongoDb()

	var query []bson.M
	query = []bson.M{}

	bsonSort := bson.M{"clusterid": 1}
	query = append(query, bson.M{"$sort": bsonSort})

	query = append(query,
		bson.M{
			"$lookup": bson.M{
				"from":         "projects",
				"localField":   "metadata.projectid",
				"foreignField": "_id",
				"as":           "projects",
			},
		},
		bson.M{
			"$set": bson.M{
				"metadata": bson.M{
					"project": bson.M{
						"$first": "$projects",
					},
				},
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "workspaces",
				"localField":   "workspaceid",
				"foreignField": "_id",
				"as":           "workspaces",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "workspaces",
				"localField":   "workspaceid",
				"foreignField": "_id",
				"as":           "workspaces",
			},
		},
		bson.M{
			"$set": bson.M{
				"workspace": bson.M{
					"$first": "$workspaces",
				},
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "datacenters",
				"localField":   "workspaces.datacenterid",
				"foreignField": "_id",
				"as":           "datacenters",
			},
		},
		bson.M{
			"$set": bson.M{
				"workspace": bson.M{
					"datacenter": bson.M{
						"$first": "$datacenters",
					},
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":         0,
				"clustername": 1,
				"clusterid":   1,
				"environment": 1,
				"projectname": "$metadata.project.name",
				"egress": bson.M{
					"ipv4": "$topology.egressip",
					"ipv6": nil,
				},
				"controlplaneendpoint": bson.M{
					"ipv4": bson.M{
						"$concat": bson.A{
							bson.M{
								"$arrayElemAt": bson.A{
									bson.M{
										"$split": bson.A{
											"$topology.controlplaneendpoint",
											":",
										},
									},
									0,
								},
							},
						},
					},
					"ipv6": nil,
				},
				"controlplaneendpointport": bson.M{
					"$concat": bson.A{
						bson.M{
							"$arrayElemAt": bson.A{
								bson.M{
									"$split": bson.A{
										"$topology.controlplaneendpoint",
										":",
									},
								},
								1,
							},
						},
					},
				},
				"datacenter": bson.M{
					"name":        "$workspace.datacenter.name",
					"provider":    "$workspace.datacenter.provider",
					"apiEndpoint": "$workspace.datacenter.apiendpoint",
				},
			},
		},
	)

	mongoCol := db.Collection(CollectionName)
	results, err := mongoCol.Aggregate(ctx, query)
	if err != nil {
		return nil, errors.New("could not get cluster control plane metadata")
	}
	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	if results.RemainingBatchLength() == 0 {
		return nil, nil
	}

	var metadata []apicontracts.ClusterControlPlaneMetadata
	if err = results.All(ctx, &metadata); err != nil {
		return nil, errors.New("could not get error")
	}

	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	return metadata, nil
}
