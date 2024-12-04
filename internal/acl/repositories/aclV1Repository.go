package aclrepository

import (
	"context"
	"errors"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	CollectionName = "clusters"
)

func GetAllACL1(ctx context.Context) ([]aclmodels.AclV1ListItem, error) {
	db := mongodb.GetMongoDb()
	var aclList []aclmodels.AclV1ListItem
	//db.clusters.aggregate({$project: {"clusterid":1, "acl.accessgroups":1}},{$unwind: "$acl.accessgroups"},{$project: { _id: 0, "cluster": "$clusterid", "group": "$acl.accessgroups" }})

	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, bson.M{"$project": bson.M{"clusterid": 1, "acl.accessgroups": 1}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$unwind": "$acl.accessgroups"})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$project": bson.M{"_id": 0, "cluster": "$clusterid", "group": "$acl.accessgroups"}})
	clusterCollection := db.Collection(ClusterCollectionName)
	results, _ := clusterCollection.Aggregate(ctx, aggregationPipeline)

	if err := results.All(ctx, &aclList); err != nil {
		return aclList, errors.New("could not get error")
	}

	return aclList, nil

}
