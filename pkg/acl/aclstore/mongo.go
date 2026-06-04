package aclstore

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const aclCollectionName = "acl"

// MongoStore implements acl.Store backed by MongoDB.
// When Redis caching is added, wrap this store with a caching layer
// that implements the same acl.Store interface.
type MongoStore struct{}

// NewMongoStore creates a new MongoDB-backed V3 ACL store.
func NewMongoStore() *MongoStore {
	return &MongoStore{}
}

// GetByGroups returns all V3 ACL entries for the given groups in a single MongoDB query.
// Results are returned as a map keyed by group name for cache-friendly consumption.
func (s *MongoStore) GetByGroups(ctx context.Context, groups []string) (map[string][]aclmodels.AclV3ListItem, error) {
	db := mongodb.GetMongoDb()
	if db == nil {
		return nil, fmt.Errorf("mongodb not initialized")
	}

	filter := bson.M{
		"version": 3,
		"group":   bson.M{"$in": groups},
	}

	cursor, err := db.Collection(aclCollectionName).Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query ACL entries: %w", err)
	}
	defer cursor.Close(ctx)

	var entries []aclmodels.AclV3ListItem
	if err := cursor.All(ctx, &entries); err != nil {
		return nil, fmt.Errorf("failed to decode ACL entries: %w", err)
	}

	// Group results by group name
	result := make(map[string][]aclmodels.AclV3ListItem, len(groups))
	for i := range entries {
		g := entries[i].Group
		result[g] = append(result[g], entries[i])
	}
	return result, nil
}
