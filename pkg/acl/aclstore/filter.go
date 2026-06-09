package aclstore

import (
	"cmp"
	"slices"
	"sort"

	"github.com/NorskHelsenett/ror/pkg/acl"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rordefs"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// DenyAllFilter is the filter returned when the user has no matching access.
// It matches an impossible scope+subject pair, ensuring no documents are returned.
var DenyAllFilter = bson.M{"$match": bson.M{"rormeta.ownerref.scope": "NA-UNKNOWN", "rormeta.ownerref.subject": "NA-UNKNOWN"}}

// OwnerrefsToFilter converts a list of ownerrefs from the resolver into a
// MongoDB aggregation pipeline stage ($match) that scopes resource queries
// to only the resources the user has access to.
//
// Behavior:
//   - nil refs (unrestricted) → bson.M{} (empty filter — matches everything)
//   - empty refs (no access) → DenyAllFilter (matches nothing)
//   - refs with scope "ror" + subject as a scope name → scope-level grants
//     (e.g. {Scope:"ror", Subject:"cluster"} → all resources where ownerref.scope = "cluster")
//   - other refs → specific scope+subject matches via $in
//   - scopes already covered by a ror-level grant are deduplicated
//
// The returned bson.M is intended to be appended as a stage in an aggregation pipeline.
func OwnerrefsToFilter(refs []acl.Ownerref) bson.M {
	// nil = unrestricted (global/all access)
	if refs == nil {
		return bson.M{}
	}

	// Empty = no matching access
	if len(refs) == 0 {
		return DenyAllFilter
	}

	orquery := bson.A{}

	// Collect ror-scope entries — these grant access to ALL resources of a given scope
	var globalScopes []aclscope.Scope
	for _, ref := range refs {
		if ref.Scope != aclscope.ScopeRor {
			continue
		}
		// scope=ror, subject=globalscope → unrestricted (defensive; resolver returns nil for this)
		if ref.Subject == aclscope.SubjectGlobal {
			return bson.M{}
		}
		scopeGrant := aclscope.Scope(ref.Subject)
		orquery = append(orquery, bson.M{
			"rormeta.ownerref.scope": string(scopeGrant),
		})
		globalScopes = append(globalScopes, scopeGrant)
	}

	// Collect specific scope+subject pairs, excluding those already covered by ror-level grants
	type scopeSubject struct {
		scope   string
		subject string
	}
	var specific []scopeSubject
	for _, ref := range refs {
		if ref.Scope == aclscope.ScopeRor {
			continue
		}
		if slices.Contains(globalScopes, ref.Scope) {
			continue
		}
		specific = append(specific, scopeSubject{scope: string(ref.Scope), subject: string(ref.Subject)})
	}

	// Sort for deterministic output
	slices.SortFunc(specific, func(a, b scopeSubject) int {
		if c := cmp.Compare(a.scope, b.scope); c != 0 {
			return c
		}
		return cmp.Compare(a.subject, b.subject)
	})

	if len(specific) > 0 {
		inquery := bson.A{}
		for _, ss := range specific {
			inquery = append(inquery, bson.D{
				{Key: "scope", Value: ss.scope},
				{Key: "subject", Value: ss.subject},
			})
		}
		orquery = append(orquery, bson.M{
			"rormeta.ownerref": bson.M{"$in": inquery},
		})
	}

	// Edge case: all refs were ror-scope but filtered out as global-scope covered
	if len(orquery) == 0 {
		return DenyAllFilter
	}

	return bson.M{
		"$match": bson.M{
			"$or": orquery,
		},
	}
}

// ClusterIdentityFilter returns a pipeline stage that scopes resource queries
// to resources owned by a specific cluster. Used when the identity is a cluster
// (which has implicit read/create/update access to its own resources).
func ClusterIdentityFilter(clusterID string) bson.M {
	return bson.M{
		"$match": bson.M{
			"rormeta.ownerref.scope":   string(aclscope.ScopeCluster),
			"rormeta.ownerref.subject": clusterID,
		},
	}
}

// ProtectedResourceTypes maps a Capability (without the verb) to the
// resource kinds it protects. The verb is appended at check time by
// ResourceTypeFilter (VerbRead) / ResourceTypeWriteFilter (VerbWrite).
//
// Example: CapRorConfig protects Configuration resources.
// A user needs CapRorConfig.WithVerb(VerbRead) to query them and
// CapRorConfig.WithVerb(VerbWrite) to mutate them.
//
// Resources not listed here are accessible with the standard ror:read / ror:write capabilities.
var ProtectedResourceTypes = map[aclmodels.Capability][]string{
	aclmodels.CapRorConfig: {
		rordefs.ResourceConfig.Kind,
	},
}

// ResourceTypeFilter builds a MongoDB aggregation pipeline stage that excludes
// resource kinds the user is not authorized to read.
//
// For each entry in ProtectedResourceTypes it checks whether the user has
// the capability + VerbRead. Missing → those kinds are added to $nin.
//
// Behavior:
//   - user has all protected read capabilities → bson.M{} (no restriction)
//   - user lacks some capabilities → {"$match": {"typemeta.kind": {"$nin": [...]}}}
//   - user lacks all capabilities → all protected kinds excluded
//   - empty/nil access list → all protected kinds excluded
func ResourceTypeFilter(access []aclmodels.AccessTypeV3) bson.M {
	return resourceTypeFilterByVerb(access, aclmodels.VerbRead)
}

// ResourceTypeWriteFilter builds a MongoDB aggregation pipeline stage that excludes
// resource kinds the user is not authorized to write.
//
// Same logic as ResourceTypeFilter but checks capability + VerbWrite.
func ResourceTypeWriteFilter(access []aclmodels.AccessTypeV3) bson.M {
	return resourceTypeFilterByVerb(access, aclmodels.VerbWrite)
}

func resourceTypeFilterByVerb(access []aclmodels.AccessTypeV3, verb aclmodels.Verb) bson.M {
	var excluded []string

	for cap, kinds := range ProtectedResourceTypes {
		required := cap.WithVerb(verb)
		if !slices.Contains(access, required) {
			excluded = append(excluded, kinds...)
		}
	}

	if len(excluded) == 0 {
		return bson.M{}
	}

	// Sort for deterministic output
	sort.Strings(excluded)

	return bson.M{
		"$match": bson.M{
			"typemeta.kind": bson.M{"$nin": excluded},
		},
	}
}
