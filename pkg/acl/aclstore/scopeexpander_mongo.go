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

const resourceV2Collection = "resourcesv2"

// MongoScopeExpander implements acl.ScopeExpander by walking the
// ownerref chain in the resourcesv2 collection. No hardcoded hierarchy —
// the tree is derived entirely from rormeta.ownerref data on each resource.
type MongoScopeExpander struct {
	db *mongo.Database
}

// NewMongoScopeExpander creates a new MongoDB-backed scope expander.
func NewMongoScopeExpander(db *mongo.Database) *MongoScopeExpander {
	return &MongoScopeExpander{db: db}
}

// ownerRef is a minimal projection of a resourcesv2 document, carrying only
// the fields needed to build an acl.Ownerref.
type ownerRef struct {
	UID  string `bson:"uid"`
	Kind string `bson:"kind"`
}

// ExpandScope recursively finds all descendant ownerrefs by walking the
// ownerref chain in resourcesv2. Returns nil if no resources have the given ownerref.
//
// The full descendant subtree is resolved in a single $graphLookup aggregation
// (resource.uid -> child.rormeta.ownerref.subject) instead of issuing one query
// per node.
//
// Only "owner" nodes are returned: a node is an owner iff at least one other
// resource in the subtree references its uid as rormeta.ownerref.subject. Leaf
// resources ("stubs", which own nothing) are excluded — they are never anyone's
// ownerref.subject, so they would match no resources in an OwnerrefsToFilter
// query, and they remain reachable through their (owning) parent regardless.
// This keeps the result small, which matters because the expander runs on every
// authorized read.
func (e *MongoScopeExpander) ExpandScope(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject) ([]acl.Ownerref, error) {
	ctx, span := rortracer.StartSpan(ctx, "acl.MongoScopeExpander.ExpandScope")
	defer span.End()
	span.SetAttributes(
		attribute.String("acl.scope", string(scope)),
		attribute.String("acl.subject", string(subject)),
	)

	seed := acl.Ownerref{Scope: scope, Subject: subject}
	expanded, err := e.expandSeeds(ctx, []acl.Ownerref{seed})
	if err != nil {
		return nil, rortracer.SpanError(span, err)
	}

	result := expanded[seed]
	span.SetAttributes(
		attribute.Int("acl.queries", 1),
		attribute.Int("acl.descendants", len(result)),
	)
	if len(result) == 0 {
		return nil, nil
	}
	return result, nil
}

// ExpandScopes expands several scope+subject seeds in a single aggregation,
// returning the owner descendants for each seed keyed by the seed ownerref.
// Batching collapses many per-entry round-trips into one. The same owners-only
// (stub-excluding) rules as ExpandScope apply to every seed.
func (e *MongoScopeExpander) ExpandScopes(ctx context.Context, seeds []acl.Ownerref) (map[acl.Ownerref][]acl.Ownerref, error) {
	ctx, span := rortracer.StartSpan(ctx, "acl.MongoScopeExpander.ExpandScopes")
	defer span.End()
	span.SetAttributes(attribute.Int("acl.seeds", len(seeds)))

	expanded, err := e.expandSeeds(ctx, seeds)
	if err != nil {
		return nil, rortracer.SpanError(span, err)
	}

	total := 0
	for _, refs := range expanded {
		total += len(refs)
	}
	span.SetAttributes(
		attribute.Int("acl.queries", 1),
		attribute.Int("acl.descendants", total),
	)
	return expanded, nil
}

// expandSeeds resolves the owner-descendants of every seed in one aggregation.
// Each seed becomes a row (via $unwind of a literal seed array) so that
// $graphLookup keeps each subtree separate and results can be attributed back to
// (and cached per) individual seeds. Subjects (uids) are globally unique, so the
// result rows are keyed by subject and mapped back to the seed ownerref.
func (e *MongoScopeExpander) expandSeeds(ctx context.Context, seeds []acl.Ownerref) (map[acl.Ownerref][]acl.Ownerref, error) {
	out := make(map[acl.Ownerref][]acl.Ownerref, len(seeds))
	if len(seeds) == 0 {
		return out, nil
	}

	if e.db == nil {
		return nil, fmt.Errorf("mongodb not initialized")
	}

	bySubject := make(map[string]acl.Ownerref, len(seeds))
	seedVals := bson.A{}
	for _, s := range seeds {
		subj := string(s.Subject)
		if _, ok := bySubject[subj]; ok {
			continue
		}
		bySubject[subj] = s
		seedVals = append(seedVals, subj)
		out[s] = nil // ensure every queried seed is present in the result
	}

	collection := e.db.Collection(resourceV2Collection)

	pipeline := mongo.Pipeline{
		// Emit one synthetic row per seed subject. Seeds need not exist as
		// documents; one graph traversal runs per seed instead of one per
		// direct child (a scope can have thousands of direct children).
		bson.D{{Key: "$limit", Value: 1}},
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "seed", Value: bson.D{{Key: "$literal", Value: seedVals}}},
		}}},
		bson.D{{Key: "$unwind", Value: "$seed"}},
		// Recursively gather each seed's subtree by following the ownerref chain
		// (resource.uid -> child.rormeta.ownerref.subject). uids are globally
		// unique so connecting on subject alone is sufficient.
		bson.D{{Key: "$graphLookup", Value: bson.D{
			{Key: "from", Value: resourceV2Collection},
			{Key: "startWith", Value: "$seed"},
			{Key: "connectFromField", Value: "uid"},
			{Key: "connectToField", Value: "rormeta.ownerref.subject"},
			{Key: "as", Value: "descendants"},
		}}},
		// Build the set of uids that own at least one resource within this
		// subtree: every uid that appears as some descendant's ownerref.subject.
		// A leaf resource ("stub") never appears here.
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "ownerUids", Value: bson.D{{Key: "$setUnion", Value: bson.A{
				bson.D{{Key: "$map", Value: bson.D{
					{Key: "input", Value: "$descendants"},
					{Key: "as", Value: "d"},
					{Key: "in", Value: "$$d.rormeta.ownerref.subject"},
				}}},
			}}}},
		}}},
		// Keep only owner nodes (exclude stubs) and trim to uid + kind.
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "seed", Value: 1},
			{Key: "owners", Value: bson.D{{Key: "$map", Value: bson.D{
				{Key: "input", Value: bson.D{{Key: "$filter", Value: bson.D{
					{Key: "input", Value: "$descendants"},
					{Key: "as", Value: "d"},
					{Key: "cond", Value: bson.D{{Key: "$in", Value: bson.A{"$$d.uid", "$ownerUids"}}}},
				}}}},
				{Key: "as", Value: "d"},
				{Key: "in", Value: bson.D{
					{Key: "uid", Value: "$$d.uid"},
					{Key: "kind", Value: "$$d.typemeta.kind"},
				}},
			}}}},
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to expand scopes via graph lookup: %w", err)
	}

	var docs []struct {
		Seed   string     `bson:"seed"`
		Owners []ownerRef `bson:"owners"`
	}
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, fmt.Errorf("failed to decode resourcesv2 scope expansion results: %w", err)
	}

	for _, d := range docs {
		seed, ok := bySubject[d.Seed]
		if !ok {
			continue
		}
		var refs []acl.Ownerref
		seen := make(map[acl.Ownerref]struct{})
		for _, c := range d.Owners {
			ref := acl.Ownerref{Scope: aclscope.Scope(c.Kind), Subject: aclscope.Subject(c.UID)}
			if _, dup := seen[ref]; dup {
				continue
			}
			seen[ref] = struct{}{}
			refs = append(refs, ref)
		}
		out[seed] = refs
	}

	return out, nil
}
