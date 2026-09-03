package aclstore

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/acl"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/NorskHelsenett/ror/pkg/telemetry/rortracer"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.opentelemetry.io/otel/attribute"
)

// MongoAncestorResolver implements acl.AncestorResolver by walking the ownerref
// chain in the resourcesv2 collection upward (child -> parent). It is the
// inverse of MongoScopeExpander; the tree is derived entirely from
// rormeta.ownerref data, with no hardcoded hierarchy.
type MongoAncestorResolver struct {
	// dbProvider returns the live *mongo.Database on every call. It must not be
	// cached: the underlying mongo client is reconnected (and the previous one
	// disconnected) whenever its credentials are rotated, so a captured handle
	// would start failing with "client is disconnected" after the first renewal.
	dbProvider func() *mongo.Database
}

// compile-time assurance that MongoAncestorResolver satisfies the interface.
var _ acl.AncestorResolver = (*MongoAncestorResolver)(nil)

// NewMongoAncestorResolver creates a new MongoDB-backed ancestor resolver.
// dbProvider must return the current *mongo.Database; it is called on every
// resolution so the resolver always uses the live connection (see the field doc).
func NewMongoAncestorResolver(dbProvider func() *mongo.Database) *MongoAncestorResolver {
	return &MongoAncestorResolver{dbProvider: dbProvider}
}

// ancestorNode is a minimal projection of an ancestor resourcesv2 document.
type ancestorNode struct {
	UID  string `bson:"uid"`
	Kind string `bson:"kind"`
}

// Ancestors returns the ancestor ownerrefs of {scope, subject}, nearest-first.
// The resource is located by its uid (subject); scope is accepted for interface
// symmetry. The queried resource itself is not included. Returns an empty slice
// when the resource has no ancestors or does not exist.
func (r *MongoAncestorResolver) Ancestors(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject) ([]acl.Ownerref, error) {
	ctx, span := rortracer.StartSpan(ctx, "acl.MongoAncestorResolver.Ancestors")
	defer span.End()
	span.SetAttributes(
		attribute.String("acl.scope", string(scope)),
		attribute.String("acl.subject", string(subject)),
	)

	db := r.dbProvider()
	if db == nil {
		return nil, fmt.Errorf("mongodb not initialized")
	}

	pipeline := mongo.Pipeline{
		// Locate the resource by its (globally unique) uid.
		bson.D{{Key: "$match", Value: bson.D{{Key: "uid", Value: string(subject)}}}},
		// Follow the ownerref chain upward: each node's parent is the resource
		// whose uid equals this node's rormeta.ownerref.subject.
		bson.D{{Key: "$graphLookup", Value: bson.D{
			{Key: "from", Value: resourceV2Collection},
			{Key: "startWith", Value: "$rormeta.ownerref.subject"},
			{Key: "connectFromField", Value: "rormeta.ownerref.subject"},
			{Key: "connectToField", Value: "uid"},
			{Key: "as", Value: "ancestors"},
			{Key: "depthField", Value: "depth"},
		}}},
		bson.D{{Key: "$unwind", Value: "$ancestors"}},
		// depth 0 is the immediate parent, so ascending depth is nearest-first.
		bson.D{{Key: "$sort", Value: bson.D{{Key: "ancestors.depth", Value: 1}}}},
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "uid", Value: "$ancestors.uid"},
			{Key: "kind", Value: "$ancestors.typemeta.kind"},
		}}},
	}

	cursor, err := db.Collection(resourceV2Collection).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, rortracer.SpanError(span, fmt.Errorf("failed to resolve ancestors: %w", err))
	}

	var nodes []ancestorNode
	if err := cursor.All(ctx, &nodes); err != nil {
		return nil, rortracer.SpanError(span, fmt.Errorf("failed to decode ancestors: %w", err))
	}

	result := make([]acl.Ownerref, 0, len(nodes))
	for _, n := range nodes {
		if n.UID == "" || n.Kind == "" {
			continue
		}
		result = append(result, acl.Ownerref{Scope: aclscope.Scope(n.Kind), Subject: aclscope.Subject(n.UID)})
	}
	span.SetAttributes(attribute.Int("acl.ancestors", len(result)))
	return result, nil
}
