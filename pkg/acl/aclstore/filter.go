package aclstore

import (
	"cmp"
	"slices"

	"github.com/NorskHelsenett/ror/pkg/acl"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"

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
