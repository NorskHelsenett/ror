package aclstore

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/acl"
	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const resourceV2Collection = "resourcesv2"

// MongoScopeExpander implements acl.ScopeExpander by walking the
// ownerref chain in the resourcesv2 collection. No hardcoded hierarchy —
// the tree is derived entirely from rormeta.ownerref data on each resource.
type MongoScopeExpander struct{}

// NewMongoScopeExpander creates a new MongoDB-backed scope expander.
func NewMongoScopeExpander() *MongoScopeExpander {
	return &MongoScopeExpander{}
}

// resourceRef is a minimal projection of a resourcesv2 document.
type resourceRef struct {
	UID      string          `bson:"uid"`
	TypeMeta resourceRefType `bson:"typemeta"`
}

type resourceRefType struct {
	Kind string `bson:"kind"`
}

// ExpandScope recursively finds all descendant ownerrefs by walking the
// ownerref chain in resourcesv2. Returns nil if no resources have the given ownerref.
func (e *MongoScopeExpander) ExpandScope(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject) ([]acl.Ownerref, error) {
	db := mongodb.GetMongoDb()
	if db == nil {
		return nil, fmt.Errorf("mongodb not initialized")
	}

	var result []acl.Ownerref
	seen := make(map[acl.Ownerref]struct{})

	// BFS queue: start with the given scope+subject
	type expandItem struct {
		scope   aclscope.Scope
		subject aclscope.Subject
	}
	queue := []expandItem{{scope: scope, subject: subject}}

	collection := db.Collection(resourceV2Collection)
	projection := bson.M{"uid": 1, "typemeta.kind": 1, "_id": 0}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		filter := bson.M{
			"rormeta.ownerref.scope":   string(item.scope),
			"rormeta.ownerref.subject": string(item.subject),
		}

		cursor, err := collection.Find(ctx, filter, options.Find().SetProjection(projection))
		if err != nil {
			return nil, fmt.Errorf("failed to query resourcesv2 for scope expansion: %w", err)
		}

		var refs []resourceRef
		if err := cursor.All(ctx, &refs); err != nil {
			return nil, fmt.Errorf("failed to decode resourcesv2 scope expansion results: %w", err)
		}

		for _, ref := range refs {
			ownerref := acl.Ownerref{
				Scope:   aclscope.Scope(ref.TypeMeta.Kind),
				Subject: aclscope.Subject(ref.UID),
			}
			if _, ok := seen[ownerref]; ok {
				continue
			}
			seen[ownerref] = struct{}{}
			result = append(result, ownerref)
			// Enqueue for further expansion (children of this resource)
			queue = append(queue, expandItem{scope: ownerref.Scope, subject: ownerref.Subject})
		}
	}

	if len(result) == 0 {
		return nil, nil
	}
	return result, nil
}
