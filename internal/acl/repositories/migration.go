package aclrepository

import (
	"context"
	"time"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ClusterCollectionName = "clusters"
	AclCollectionName     = "acl"
)

// TODO: remove when migration is complete
func MigrateAcl1UpdateCreatreAcl2(ctx context.Context, aclListV1 []aclmodels.AclV1ListItem, aclListV2 []aclmodels.AclV2ListItem) error {
	db := mongodb.GetMongoDb()
	aclCollection := db.Collection(AclCollectionName)

	for _, listitem := range aclListV1 {
		newline := aclmodels.AclV2ListItem{
			Version:    2,
			Group:      listitem.Group,
			Scope:      aclmodels.Acl2ScopeCluster,
			Subject:    aclmodels.Acl2Subject(listitem.Cluster),
			Access:     aclmodels.AclV2ListItemAccess{Read: true, Create: false, Update: false, Delete: false, Owner: false},
			Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: true},
			Created:    time.Now(),
			IssuedBy:   "migrate@ror.nhn.no",
		}

		query := bson.M{"group": newline.Group, "scope": "cluster", "subject": newline.Subject}
		update := bson.M{"$set": newline}
		opts := options.Update().SetUpsert(true)
		result, err := aclCollection.UpdateOne(ctx, query, update, opts)

		if err != nil {
			return err
		}
		if result.ModifiedCount > 0 || result.UpsertedCount > 0 {
			rlog.Warn("ACLs updated/upserted", rlog.Int64("modified count", result.ModifiedCount), rlog.Int64("upserted count", result.UpsertedCount))
		}

	}

	return nil
}

// TODO: remove when migration is complete
func MigrateAcl1DeleteRemovedAcl1(ctx context.Context, aclListV1 []aclmodels.AclV1ListItem, aclListV2 []aclmodels.AclV2ListItem) {
	for _, listitem := range aclListV2 {
		if !migrateAcl1checkV1Exists(listitem.Group, string(listitem.Subject), listitem.Access, listitem.Kubernetes, aclListV1) {
			idPrimitive, _ := primitive.ObjectIDFromHex(listitem.Id)
			rlog.Warn("might need to remove element from acl collection", rlog.String("listitem", listitem.Id))
			rlog.Warn("db.acl.deleteone", rlog.String("id", idPrimitive.String()))
		}
	}
}

// TODO: remove when migration is complete
func migrateAcl1checkV1Exists(group string, cluster string, access aclmodels.AclV2ListItemAccess, kubernetes aclmodels.AclV2ListItemKubernetes, aclListV1 []aclmodels.AclV1ListItem) bool {

	defaultaccess := aclmodels.AclV2ListItemAccess{Read: true, Create: false, Update: false, Delete: false, Owner: false}
	defaultk8 := aclmodels.AclV2ListItemKubernetes{Logon: true}

	for _, item := range aclListV1 {
		if item.Group == group && item.Cluster == cluster && access == defaultaccess && kubernetes == defaultk8 {
			return true
		}
	}
	return false
}
