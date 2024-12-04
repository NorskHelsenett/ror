package aclrepository

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/bson"
)

// dbcollection
var collectionName = "acl"
var denyallACL = aclmodels.AclV2ListItemAccess{Read: false, Create: false, Update: false, Delete: false, Owner: false}

// GetAllACL2 Gets all ACL2 items returns []aclmodels.AclV2ListItem
func GetAllACL2(ctx context.Context) ([]aclmodels.AclV2ListItem, error) {
	var aggregationPipeline []bson.M

	results := make([]aclmodels.AclV2ListItem, 0)
	err := mongodb.Aggregate(ctx, AclCollectionName, aggregationPipeline, &results)

	return results, err
}

// GetACL2ByIdentityQuery Gets ACL2 Access model for identity/scope returns aclmodels.AclV2ListItems
func GetACL2ByIdentityQuery(ctx context.Context, aclQuery aclmodels.AclV2QueryAccessScope) aclmodels.AclV2ListItems {
	identity := rorcontext.GetIdentityFromRorContext(ctx)
	denyall := denyallACL

	aclReturnArray := aclmodels.AclV2ListItems{
		Scope:   aclQuery.Scope,
		Subject: "NA",
		Global:  denyall,
	}

	if !identity.IsCluster() {
		dbResult := make([]aclmodels.AclV2ListItem, 0)

		var aggregationPipeline []bson.M
		aggregationPipeline = append(aggregationPipeline, createACLV2FilterByScope(identity, aclQuery.Scope)...)

		err := mongodb.Aggregate(ctx, AclCollectionName, aggregationPipeline, &dbResult)
		if err != nil {
			rlog.Error("could not query mongodb", err)
			return aclReturnArray
		}

		if len(dbResult) > 0 {
			for _, result := range dbResult {
				if result.Scope == aclmodels.Acl2ScopeRor {
					aclReturnArray.Global = compileAccessSum(aclReturnArray.Global, result.Access)
				}
				aclReturnArray.Items = append(aclReturnArray.Items, result)
			}
		}

		return aclReturnArray
	} else if identity.IsCluster() && aclQuery.Scope == aclmodels.Acl2ScopeCluster {
		aclReturn := aclmodels.AclV2ListItem{
			Version:    2,
			Group:      "NA",
			Scope:      aclmodels.Acl2ScopeCluster,
			Subject:    aclmodels.Acl2Subject(identity.GetId()),
			Access:     aclmodels.AclV2ListItemAccess{Read: true, Create: true, Update: true, Delete: false, Owner: false},
			Kubernetes: aclmodels.AclV2ListItemKubernetes{Logon: false},
		}
		aclReturnArray.Items = append(aclReturnArray.Items, aclReturn)
		return aclReturnArray
	}
	return aclReturnArray
}

// CheckAcl2ByIdentityQuery Gets ACL2 Access model for identity/scope/subject returns aclmodels.AclV2ListItemAccess
func CheckAcl2ByIdentityQuery(ctx context.Context, aclQuery aclmodels.AclV2QueryAccessScopeSubject) aclmodels.AclV2ListItemAccess {
	denyall := denyallACL
	identity := rorcontext.GetIdentityFromRorContext(ctx)

	if !identity.IsCluster() {
		dbResult := make([]aclmodels.AclV2ListItem, 0)
		var aggregationPipeline []bson.M
		aggregationPipeline = append(aggregationPipeline, createACLV2FilterByScopeSubject(identity, aclQuery.Scope, aclQuery.Subject)...)

		err := mongodb.Aggregate(ctx, AclCollectionName, aggregationPipeline, &dbResult)
		if err != nil {
			rlog.Error("could not query mongodb", err)
			return denyall
		}

		return compileAccess(dbResult)
	}

	if identity.IsCluster() && aclQuery.Subject == aclmodels.Acl2Subject(identity.GetId()) && aclQuery.Scope == aclmodels.Acl2ScopeCluster {
		return aclmodels.AclV2ListItemAccess{Read: true, Create: true, Update: true, Delete: false, Owner: false}
	}

	return denyall
}

func CheckAcl2ByCluster(ctx context.Context, aclQuery aclmodels.AclV2QueryAccessScopeSubject) []aclmodels.AclV2ListItem {
	identity := rorcontext.GetIdentityFromRorContext(ctx)
	result := make([]aclmodels.AclV2ListItem, 0)

	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, createACLV2FilterByScopeSubject(identity, aclQuery.Scope, aclQuery.Subject)...)

	err := mongodb.Aggregate(ctx, AclCollectionName, aggregationPipeline, &result)
	if err != nil {
		rlog.Error("could not query mongodb", err)
		return result
	}

	return result
}

func compileAccess(acls []aclmodels.AclV2ListItem) aclmodels.AclV2ListItemAccess {
	if len(acls) == 0 {
		return denyallACL
	}

	compiledAccess := denyallACL
	for _, result := range acls {
		compiledAccess = compileAccessSum(compiledAccess, result.Access)
	}
	return compiledAccess
}

// createACLV2FilterByScopeSubject returns a mongodb query for querying the acl database based on the identiys groups.
func createACLV2FilterByScopeSubject(identity identitymodels.Identity, scope aclmodels.Acl2Scope, subject aclmodels.Acl2Subject) []bson.M {
	var filters []bson.M
	var filterGroups bson.A
	denyallGroups := []bson.M{{"$match": bson.M{"group": bson.M{"$in": bson.A{"Unknown-Unauthorized"}}}}}

	groups, err := identity.ReturnGroupQuery()
	if err != nil {
		rlog.Error("could not extract groups from user", err)
		return denyallGroups
	}

	subjectArr := []string{string(scope), string(aclmodels.Acl2RorSubjectGlobal)}

	filterGroups = groups

	if len(filterGroups) == 0 {

		filterGroups = bson.A{"Unknown-Unauthorized"}
	}

	if filterGroups[0] == "" {

		filterGroups = bson.A{"Unknown-Unauthorized"}
	}

	filters = append(filters, bson.M{
		"$match": bson.M{
			"group": bson.M{
				"$in": filterGroups,
			},
		},
	})

	filters = append(filters, bson.M{
		"$match": bson.M{
			"$or": bson.A{
				bson.M{
					"scope":   scope,
					"subject": subject,
				},
				bson.M{
					"scope": aclmodels.Acl2ScopeRor,
					"subject": bson.M{
						"$in": subjectArr,
					},
				},
			},
		},
	})
	return filters
}

func createACLV2FilterByScope(identity identitymodels.Identity, scope aclmodels.Acl2Scope) []bson.M {
	var filters []bson.M
	var filterGroups bson.A
	denyallGroups := []bson.M{{"$match": bson.M{"group": bson.M{"$in": bson.A{"Unknown-Unauthorized"}}}}}

	groups, err := identity.ReturnGroupQuery()
	if err != nil {
		rlog.Error("could not extract groups from user", err)
		return denyallGroups
	}

	filterGroups = groups

	if len(groups) == 0 {

		filterGroups = bson.A{"Unknown-Unauthorized"}
	}

	if filterGroups[0] == "" {

		filterGroups = bson.A{"Unknown-Unauthorized"}
	}

	subjectArr := []string{string(scope), string(aclmodels.Acl2RorSubjectGlobal)}

	filters = append(filters, bson.M{
		"$match": bson.M{
			"group": bson.M{
				"$in": filterGroups,
			},
		},
	})

	filters = append(filters, bson.M{
		"$match": bson.M{
			"$or": bson.A{
				bson.M{
					"scope": scope,
				},
				bson.M{
					"scope": aclmodels.Acl2ScopeRor,
					"subject": bson.M{
						"$in": subjectArr,
					},
				},
			},
		},
	})
	return filters
}

// Return the sum of two AclV2ListItemAccess
func compileAccessSum(existing aclmodels.AclV2ListItemAccess, added aclmodels.AclV2ListItemAccess) aclmodels.AclV2ListItemAccess {
	compiledAccess := existing
	if added.Read {
		compiledAccess.Read = true
	}
	if added.Create {
		compiledAccess.Create = true
	}
	if added.Update {
		compiledAccess.Update = true
	}
	if added.Delete {
		compiledAccess.Delete = true
	}
	if added.Owner {
		compiledAccess.Owner = true
	}
	return compiledAccess
}
