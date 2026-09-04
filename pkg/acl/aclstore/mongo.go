package aclstore

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/telemetry/rortracer"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.opentelemetry.io/otel/attribute"
)

const aclCollectionName = "acl"

// MongoStore implements acl.Store backed by MongoDB.
// It queries both V2 and V3 entries and converts to the requested format.
type MongoStore struct {
	// dbProvider returns the live *mongo.Database on every call. It must not be
	// cached: the underlying mongo client is reconnected (and the previous one
	// disconnected) whenever its credentials are rotated, so a captured handle
	// would start failing with "client is disconnected" after the first renewal.
	dbProvider func() *mongo.Database
}

// NewMongoStore creates a new MongoDB-backed ACL store. dbProvider must return
// the current *mongo.Database; it is called on every query so the store always
// uses the live connection (see the field doc for why this matters).
func NewMongoStore(dbProvider func() *mongo.Database) *MongoStore {
	return &MongoStore{dbProvider: dbProvider}
}

// aclRawEntry is used to decode both V2 and V3 entries from MongoDB.
// The Version field determines which typed decode to use.
type aclRawEntry struct {
	Version int `bson:"version"`
}

// GetByGroups returns all ACL entries as V3 items. V2 entries are converted via aclmodels.V2ToV3.
func (s *MongoStore) GetByGroups(ctx context.Context, groups []string) (map[string][]aclmodels.AclV3ListItem, error) {
	ctx, span := rortracer.StartSpan(ctx, "acl.MongoStore.GetByGroups")
	defer span.End()
	span.SetAttributes(attribute.Int("acl.groups", len(groups)))

	db := s.dbProvider()
	if db == nil {
		return nil, rortracer.SpanErrorf(span, "mongodb not initialized")
	}

	filter := bson.M{
		"version": bson.M{"$in": bson.A{2, 3}},
		"group":   bson.M{"$in": groups},
	}

	cursor, err := db.Collection(aclCollectionName).Find(ctx, filter)
	if err != nil {
		return nil, rortracer.SpanError(span, fmt.Errorf("failed to query ACL entries: %w", err))
	}
	defer func() {
		err := cursor.Close(ctx)
		if err != nil {
			rlog.Error("failed to close ACL cursor", err)
		}
	}()

	result := make(map[string][]aclmodels.AclV3ListItem, len(groups))
	entryCount := 0
	for cursor.Next(ctx) {
		var raw aclRawEntry
		if err := cursor.Decode(&raw); err != nil {
			return nil, rortracer.SpanError(span, fmt.Errorf("failed to decode ACL entry version: %w", err))
		}

		switch raw.Version {
		case 3:
			var entry aclmodels.AclV3ListItem
			if err := cursor.Decode(&entry); err != nil {
				return nil, rortracer.SpanError(span, fmt.Errorf("failed to decode V3 ACL entry: %w", err))
			}
			result[entry.Group] = append(result[entry.Group], entry)
			entryCount++
		case 2:
			var entry aclmodels.AclV2ListItem
			if err := cursor.Decode(&entry); err != nil {
				return nil, rortracer.SpanError(span, fmt.Errorf("failed to decode V2 ACL entry: %w", err))
			}
			converted := aclmodels.V2ToV3(entry)
			result[converted.Group] = append(result[converted.Group], converted)
			entryCount++
		}
	}
	if err := cursor.Err(); err != nil {
		return nil, rortracer.SpanError(span, fmt.Errorf("cursor error reading ACL entries: %w", err))
	}
	span.SetAttributes(attribute.Int("acl.entries", entryCount))
	return result, nil
}

// GetV2ByGroups returns all ACL entries as V2 items. V3 entries are converted via aclmodels.V3ToV2.
func (s *MongoStore) GetV2ByGroups(ctx context.Context, groups []string) (map[string][]aclmodels.AclV2ListItem, error) {
	db := s.dbProvider()
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
	defer func() {
		err := cursor.Close(ctx)
		if err != nil {
			rlog.Error("failed to close ACL cursor", err)
		}
	}()

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

// GetByScopeSubject returns all ACL entries for the given scope+subject as V3 items.
// V2 entries are converted via aclmodels.V2ToV3.
func (s *MongoStore) GetByScopeSubject(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject) ([]aclmodels.AclV3ListItem, error) {
	ctx, span := rortracer.StartSpan(ctx, "acl.MongoStore.GetByScopeSubject")
	defer span.End()
	span.SetAttributes(
		attribute.String("acl.scope", scope.String()),
		attribute.String("acl.subject", subject.String()),
	)

	db := s.dbProvider()
	if db == nil {
		return nil, rortracer.SpanErrorf(span, "mongodb not initialized")
	}

	filter := bson.M{
		"version": bson.M{"$in": bson.A{2, 3}},
		"scope":   scope,
		"subject": subject,
	}

	cursor, err := db.Collection(aclCollectionName).Find(ctx, filter)
	if err != nil {
		return nil, rortracer.SpanError(span, fmt.Errorf("failed to query ACL entries: %w", err))
	}
	defer func() {
		err := cursor.Close(ctx)
		if err != nil {
			rlog.Error("failed to close ACL cursor", err)
		}
	}()

	var result []aclmodels.AclV3ListItem
	for cursor.Next(ctx) {
		var raw aclRawEntry
		if err := cursor.Decode(&raw); err != nil {
			return nil, rortracer.SpanError(span, fmt.Errorf("failed to decode ACL entry version: %w", err))
		}

		switch raw.Version {
		case 3:
			var entry aclmodels.AclV3ListItem
			if err := cursor.Decode(&entry); err != nil {
				return nil, rortracer.SpanError(span, fmt.Errorf("failed to decode V3 ACL entry: %w", err))
			}
			result = append(result, entry)
		case 2:
			var entry aclmodels.AclV2ListItem
			if err := cursor.Decode(&entry); err != nil {
				return nil, rortracer.SpanError(span, fmt.Errorf("failed to decode V2 ACL entry: %w", err))
			}
			result = append(result, aclmodels.V2ToV3(entry))
		}
	}
	if err := cursor.Err(); err != nil {
		return nil, rortracer.SpanError(span, fmt.Errorf("cursor error reading ACL entries: %w", err))
	}
	span.SetAttributes(attribute.Int("acl.entries", len(result)))
	return result, nil
}

// GetV2ByScopeSubject returns all ACL entries for the given scope+subject as V2 items.
// V3 entries are converted via aclmodels.V3ToV2.
func (s *MongoStore) GetV2ByScopeSubject(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject) ([]aclmodels.AclV2ListItem, error) {
	db := s.dbProvider()
	if db == nil {
		return nil, fmt.Errorf("mongodb not initialized")
	}

	filter := bson.M{
		"version": bson.M{"$in": bson.A{2, 3}},
		"scope":   scope,
		"subject": subject,
	}

	cursor, err := db.Collection(aclCollectionName).Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query ACL entries: %w", err)
	}
	defer func() {
		err := cursor.Close(ctx)
		if err != nil {
			rlog.Error("failed to close ACL cursor", err)
		}
	}()

	var result []aclmodels.AclV2ListItem
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
			result = append(result, entry)
		case 3:
			var entry aclmodels.AclV3ListItem
			if err := cursor.Decode(&entry); err != nil {
				return nil, fmt.Errorf("failed to decode V3 ACL entry: %w", err)
			}
			result = append(result, aclmodels.V3ToV2(entry))
		}
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error reading ACL entries: %w", err)
	}
	return result, nil
}
