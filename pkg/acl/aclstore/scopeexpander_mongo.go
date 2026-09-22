package aclstore

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/acl"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/telemetry/rortracer"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.opentelemetry.io/otel/attribute"
)

const resourceV2Collection = "resourcesv2"

// defaultOwnerUidsTTL bounds how stale the memoized owner-uid set may be when no
// explicit TTL is configured. A newly created owner relationship becomes visible
// to scope expansion within this window. It only gates leaf-pruning of the
// traversal, and per-seed expansion results are themselves cached
// (CachedScopeExpander), so this staleness is consistent with the access layer's
// existing eventual-consistency window. Override with WithOwnerUidsTTL.
const defaultOwnerUidsTTL = 30 * time.Second

// ownerUidsRefreshTimeout bounds a single background owner-uid refresh. It is
// generous — the Distinct scans the whole collection — but it runs off the
// request path, so it never adds request latency.
const ownerUidsRefreshTimeout = 60 * time.Second

// MongoScopeExpander implements acl.ScopeExpander by walking the
// ownerref chain in the resourcesv2 collection. No hardcoded hierarchy —
// the tree is derived entirely from rormeta.ownerref data on each resource.
type MongoScopeExpander struct {
	// dbProvider returns the live *mongo.Database on every call. It must not be
	// cached: the underlying mongo client is reconnected (and the previous one
	// disconnected) whenever its credentials are rotated, so a captured handle
	// would start failing with "client is disconnected" after the first renewal.
	dbProvider func() *mongo.Database

	// ownerUidsTTL is how long the memoized owner-uid set is reused before a
	// refresh. Set via WithOwnerUidsTTL; defaults to defaultOwnerUidsTTL.
	ownerUidsTTL time.Duration

	// ownerUids memoizes the set of owner subject uids (a Distinct over the whole
	// collection, identical for every seed) so scope expansion does not rescan
	// resourcesv2 on every authorized read. After the initial (cold) load it is
	// refreshed off the request path: a stale set is served immediately while a
	// single background goroutine repopulates it, so no request ever pays the
	// Distinct latency. See ownerUidList.
	ownerUidsMu         sync.Mutex
	ownerUids           bson.A
	ownerUidsAt         time.Time
	ownerUidsRefreshing bool // a background refresh is in flight (guarded by ownerUidsMu)

	// ownerUidsColdMu single-flights the initial synchronous load so concurrent
	// first callers issue one Distinct, not one each.
	ownerUidsColdMu sync.Mutex
}

// MongoScopeExpanderOption configures a MongoScopeExpander at construction.
type MongoScopeExpanderOption func(*MongoScopeExpander)

// WithOwnerUidsTTL sets how long the memoized owner-uid set is reused before a
// refresh. Values <= 0 are ignored (defaultOwnerUidsTTL is kept).
func WithOwnerUidsTTL(ttl time.Duration) MongoScopeExpanderOption {
	return func(e *MongoScopeExpander) {
		if ttl > 0 {
			e.ownerUidsTTL = ttl
		}
	}
}

// NewMongoScopeExpander creates a new MongoDB-backed scope expander. dbProvider
// must return the current *mongo.Database; it is called on every expansion so
// the expander always uses the live connection (see the field doc for why).
func NewMongoScopeExpander(dbProvider func() *mongo.Database, opts ...MongoScopeExpanderOption) *MongoScopeExpander {
	e := &MongoScopeExpander{dbProvider: dbProvider, ownerUidsTTL: defaultOwnerUidsTTL}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// ownerUidList returns the uids that are referenced by at least one resource as
// its rormeta.ownerref.subject — the "owner" nodes used to prune leaves from the
// graph traversal. That set is a Distinct over the entire collection (expensive,
// O(collection)) but is identical for every seed and changes slowly, so it is
// memoized for ownerUidsTTL and refreshed off the request path: once loaded, a
// stale set is returned immediately while a single background goroutine
// repopulates it (stale-while-revalidate). Only the initial cold load blocks.
func (e *MongoScopeExpander) ownerUidList(ctx context.Context, collection *mongo.Collection) (bson.A, error) {
	e.ownerUidsMu.Lock()
	if e.ownerUids != nil {
		if time.Since(e.ownerUidsAt) >= e.ownerUidsTTL && !e.ownerUidsRefreshing {
			e.ownerUidsRefreshing = true
			go e.refreshOwnerUids()
		}
		uids := e.ownerUids
		e.ownerUidsMu.Unlock()
		return uids, nil
	}
	e.ownerUidsMu.Unlock()

	// Cold start: nothing to serve yet, so block on a single-flighted load.
	return e.coldLoadOwnerUids(ctx, collection)
}

// coldLoadOwnerUids performs the first synchronous owner-uid load, ensuring only
// one Distinct runs even when many requests arrive before the set is populated.
func (e *MongoScopeExpander) coldLoadOwnerUids(ctx context.Context, collection *mongo.Collection) (bson.A, error) {
	e.ownerUidsColdMu.Lock()
	defer e.ownerUidsColdMu.Unlock()

	// Another goroutine may have populated the set while we waited for the lock.
	e.ownerUidsMu.Lock()
	if e.ownerUids != nil {
		uids := e.ownerUids
		e.ownerUidsMu.Unlock()
		return uids, nil
	}
	e.ownerUidsMu.Unlock()

	uids, err := e.fetchOwnerUids(ctx, collection)
	if err != nil {
		return nil, err
	}
	e.setOwnerUids(uids)
	return uids, nil
}

// refreshOwnerUids repopulates the owner-uid set in the background. It runs off
// the request path with its own timeout and a fresh db handle (credentials may
// have rotated). On error it leaves the previous set in place, so expansion
// keeps working on slightly staler data.
func (e *MongoScopeExpander) refreshOwnerUids() {
	defer func() {
		e.ownerUidsMu.Lock()
		e.ownerUidsRefreshing = false
		e.ownerUidsMu.Unlock()
	}()

	db := e.dbProvider()
	if db == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), ownerUidsRefreshTimeout)
	defer cancel()

	uids, err := e.fetchOwnerUids(ctx, db.Collection(resourceV2Collection))
	if err != nil {
		rlog.Warn("scope expander: background owner-uid refresh failed, keeping previous set", rlog.Any("error", err))
		return
	}
	e.setOwnerUids(uids)
}

// setOwnerUids stores a freshly loaded owner-uid set and stamps its load time.
func (e *MongoScopeExpander) setOwnerUids(uids bson.A) {
	e.ownerUidsMu.Lock()
	e.ownerUids = uids
	e.ownerUidsAt = time.Now()
	e.ownerUidsMu.Unlock()
}

// fetchOwnerUids runs the Distinct and returns the owner-uid set. Require the
// subject to be a present, non-empty string: $type screens out missing fields,
// null, and non-string values, so the result never includes a spurious uid.
func (e *MongoScopeExpander) fetchOwnerUids(ctx context.Context, collection *mongo.Collection) (bson.A, error) {
	filter := bson.D{{Key: "rormeta.ownerref.subject", Value: bson.D{
		{Key: "$type", Value: "string"},
		{Key: "$ne", Value: ""},
	}}}
	var ownerSubjects []string
	if err := collection.Distinct(ctx, "rormeta.ownerref.subject", filter).Decode(&ownerSubjects); err != nil {
		return nil, fmt.Errorf("failed to list owner subjects for scope expansion: %w", err)
	}
	uids := make(bson.A, len(ownerSubjects))
	for i, s := range ownerSubjects {
		uids[i] = s
	}
	return uids, nil
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
// The descendant subtree is resolved in a single $graphLookup aggregation
// (resource.uid -> child.rormeta.ownerref.subject) instead of issuing one query
// per node.
//
// Only "owner" nodes are ever traversed or returned: a node is an owner iff at
// least one other resource references its uid as rormeta.ownerref.subject. Leaf
// resources (which own nothing — e.g. the in-cluster resources a cluster owns)
// are pruned from the traversal itself via restrictSearchWithMatch. This keeps
// the result small — a single cluster can own tens of thousands of leaves —
// which both avoids overflowing the $graphLookup memory limit and matters
// because the expander runs on every authorized read.
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
// (leaf-excluding) traversal as ExpandScope applies to every seed.
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

	db := e.dbProvider()
	if db == nil {
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

	collection := db.Collection(resourceV2Collection)

	// The graph traversal must only ever visit owner ("parent") resources:
	// resources whose uid is referenced by at least one other resource as its
	// rormeta.ownerref.subject. Leaf resources (e.g. the in-cluster resources a
	// cluster owns — Pods, PolicyReports, ...) own nothing and must never be
	// traversed or returned: there are no ownership relations between resources
	// inside a cluster, so a leaf can never lead to another scope. Crucially,
	// pruning them at traversal time (rather than after) keeps the $graphLookup
	// result small enough to stay within MongoDB's memory limit — a single
	// cluster can own tens of thousands of leaf resources, which would otherwise
	// overflow the traversal.
	//
	// Owner chains never pass through a leaf (a leaf has no children), so
	// restricting the search to owners loses no owner-descendant.
	//
	// The owner-uid set is a Distinct over the entire collection and is identical
	// for every seed, so it is memoized (see ownerUidList) rather than recomputed
	// on each expansion — this path runs on every authorized read.
	ownerUids, err := e.ownerUidList(ctx, collection)
	if err != nil {
		return nil, err
	}

	// Scope objects (KubernetesCluster, Project, ...) must stay traversable even
	// when childless: a cluster with no in-cluster resources is referenced by
	// nothing, so it is absent from ownerUids and would otherwise be pruned,
	// hiding it from a grant on its parent. Keep any resource whose kind is a
	// scope kind, but exclude cluster-owned CRDs that reuse a scope kind name
	// (KubeVirt VirtualMachine, CAPI Machine) via the ownerref.scope guard.
	scopeKinds := scopeResourceKinds()

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
		// Recursively gather each seed's owner-descendants by following the
		// ownerref chain (resource.uid -> child.rormeta.ownerref.subject). uids
		// are globally unique so connecting on subject alone is sufficient.
		// restrictSearchWithMatch prunes non-owner (leaf) resources from the
		// traversal entirely: they are neither returned nor recursed into.
		bson.D{{Key: "$graphLookup", Value: bson.D{
			{Key: "from", Value: resourceV2Collection},
			{Key: "startWith", Value: "$seed"},
			{Key: "connectFromField", Value: "uid"},
			{Key: "connectToField", Value: "rormeta.ownerref.subject"},
			{Key: "as", Value: "descendants"},
			{Key: "restrictSearchWithMatch", Value: bson.D{{Key: "$or", Value: bson.A{
				bson.D{{Key: "uid", Value: bson.D{{Key: "$in", Value: ownerUids}}}},
				bson.D{
					{Key: "typemeta.kind", Value: bson.D{{Key: "$in", Value: scopeKinds}}},
					{Key: "rormeta.ownerref.scope", Value: bson.D{{Key: "$ne", Value: string(aclscope.ScopeCluster)}}},
				},
			}}}},
		}}},
		// All descendants are owners by construction; trim to uid + kind.
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "seed", Value: 1},
			{Key: "owners", Value: bson.D{{Key: "$map", Value: bson.D{
				{Key: "input", Value: "$descendants"},
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

// scopeResourceKinds returns the resource-kind scopes (all valid scopes except
// the system scopes ror/all/spam), derived from aclscope so the set stays in
// sync with the scope model rather than being hardcoded here.
func scopeResourceKinds() bson.A {
	kinds := bson.A{}
	for _, s := range aclscope.GetScopes() {
		switch s {
		case aclscope.ScopeRor, aclscope.ScopeAll, aclscope.ScopeSpam:
			continue
		}
		kinds = append(kinds, string(s))
	}
	return kinds
}
