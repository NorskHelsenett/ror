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
// It queries both V2 and V3 entries and converts to the requested format.
type MongoStore struct{}

// NewMongoStore creates a new MongoDB-backed ACL store.
func NewMongoStore() *MongoStore {
	return &MongoStore{}
}

// aclRawEntry is used to decode both V2 and V3 entries from MongoDB.
// The Version field determines which typed decode to use.
type aclRawEntry struct {
	Version int `bson:"version"`
}

// GetByGroups returns all ACL entries as V3 items. V2 entries are converted via aclmodels.V2ToV3.
func (s *MongoStore) GetByGroups(ctx context.Context, groups []string) (map[string][]aclmodels.AclV3ListItem, error) {
	db := mongodb.GetMongoDb()
	if db == nil {
		return nil, fmt.Errorf("mongodb not initialized")
	}

	filter := bson.M{
		"version": bson.M{"$in": bson.A{2, 3}},
		"group":   bson.M{"$in": groups},
	}

	cursor, err := db.Collection(aclCollectionName).Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query ACL entries: %w", err)
	}
	defer cursor.Close(ctx)

	result := make(map[string][]aclmodels.AclV3ListItem, len(groups))
	for cursor.Next(ctx) {
		var raw aclRawEntry
		if err := cursor.Decode(&raw); err != nil {
			return nil, fmt.Errorf("failed to decode ACL entry version: %w", err)
		}

		switch raw.Version {
		case 3:
			var entry aclmodels.AclV3ListItem
			if err := cursor.Decode(&entry); err != nil {
				return nil, fmt.Errorf("failed to decode V3 ACL entry: %w", err)
			}
			result[entry.Group] = append(result[entry.Group], entry)
		case 2:
			var entry aclmodels.AclV2ListItem
			if err := cursor.Decode(&entry); err != nil {
				return nil, fmt.Errorf("failed to decode V2 ACL entry: %w", err)
			}
			converted := aclmodels.V2ToV3(entry)
			result[converted.Group] = append(result[converted.Group], converted)
		}
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error reading ACL entries: %w", err)
	}
	return result, nil
}

// GetV2ByGroups returns all ACL entries as V2 items. V3 entries are converted via aclmodels.V3ToV2.
func (s *MongoStore) GetV2ByGroups(ctx context.Context, groups []string) (map[string][]aclmodels.AclV2ListItem, error) {
	db := mongodb.GetMongoDb()
	if db == nil {
		return nil, fmt.Errorf("mongodb not initialized")
	}

	filter := bson.M{
		"version": bson.M{"$in": bson.A{2, 3}},
		"group":   bson.M{"$in": groups},
	}

	cursor, err := db.Collection(aclCollectionName).Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query ACL entries: %w", err)
	}
	defer cursor.Close(ctx)

	result := make(map[string][]aclmodels.AclV2ListItem, len(groups))
	for cursor.Next(ctx) {
		var raw aclRawEntry
		if err := cursor.Decode(&raw); err != nil {
			return nil, fmt.Errorf("failed to decode ACL entry version: %w", err)
		}

		switch raw.Version {
		case 2:
			var entry aclmodels.AclV2ListItem
			if err := cursor.Decode(&entry); err != nil {
				return nil, fmt.Errorf("failed to decode V2 ACL entry: %w", err)
			}
			result[entry.Group] = append(result[entry.Group], entry)
		case 3:
			var entry aclmodels.AclV3ListItem
			if err := cursor.Decode(&entry); err != nil {
				return nil, fmt.Errorf("failed to decode V3 ACL entry: %w", err)
			}
			converted := aclmodels.V3ToV2(entry)
			result[converted.Group] = append(result[converted.Group], converted)
		}
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error reading ACL entries: %w", err)
	}
	return result, nil
}
