package mongohelper

import (
	"encoding/json"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

var (
	customSortFields = []string{"ingresses.datacenter", "ingresses.health", "ingresses.internet"}
)

func CreateClusterACLFilter(acl aclmodels.AclV2ListItems) bson.M {
	if acl.Global.Read {
		return bson.M{"$match": bson.M{}}
	} else {
		var clusterFilter bson.A
		if len(acl.Items) == 0 {
			clusterFilter = bson.A{"Unknown-Unauthorized"}
		} else {
			clusterFilter = getClusterFromAccessListItems(acl.Items)
		}
		filter := bson.M{
			"$match": bson.M{
				"clusterid": bson.M{
					"$in": clusterFilter,
				},
			},
		}
		return filter
	}
}

func getClusterFromAccessListItems(items []aclmodels.AclV2ListItem) bson.A {
	returnArray := bson.A{}
	for _, item := range items {
		if item.Scope == aclmodels.Acl2ScopeCluster && item.Access.Read {
			returnArray = append(returnArray, item.Subject)
		}
	}
	return returnArray
}

func CreateAggregationPipeline(filter *apicontracts.Filter, baseSort apicontracts.SortMetadata, unwindFields []string) []bson.M {
	if filter == nil {
		var aggregationPipeline []bson.M
		aggregationPipeline = append(aggregationPipeline, bson.M{"$skip": 0})
		aggregationPipeline = append(aggregationPipeline, bson.M{"$limit": 25})
		return aggregationPipeline
	}
	if filter.Limit <= 0 {
		filter.Limit = 25
	}
	aggregationPipeline := CreateFilterPipeline(filter, unwindFields)
	aggregationPipeline = append(aggregationPipeline, CreateSort(filter, baseSort)...)
	aggregationPipeline = append(aggregationPipeline, CreatePaginationPipeline(filter)...)
	return aggregationPipeline
}

func CreatePaginationPipeline(filter *apicontracts.Filter) []bson.M {
	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, bson.M{"$skip": filter.Skip})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$limit": filter.Limit})
	return aggregationPipeline
}

func CreateSort(filter *apicontracts.Filter, baseSort apicontracts.SortMetadata) []bson.M {
	var aggregationPipeline []bson.M
	sort := bson.M{}
	for _, v := range filter.Sort {
		if Contains(customSortFields, v.SortField) {
			return customSort(v)
		}
		if v.SortField != "" && (v.SortOrder == 1 || v.SortOrder == -1) {
			sort[v.SortField] = v.SortOrder
		}
	}
	if len(sort) == 0 {
		sort[baseSort.SortField] = baseSort.SortOrder
	}
	aggregationPipeline = append(aggregationPipeline, bson.M{"$sort": sort})
	return aggregationPipeline
}

func CreateFilterPipeline(filter *apicontracts.Filter, unwindFields []string) []bson.M {
	var aggregationPipeline []bson.M
	for _, v := range unwindFields {
		aggregationPipeline = append(aggregationPipeline, bson.M{"$unwind": "$" + v})
	}
	var matchPipeline []bson.M
	for _, v := range filter.Filters {
		addFilterToPipeline(&matchPipeline, v)
	}
	if len(matchPipeline) > 1 {
		aggregationPipeline = append(aggregationPipeline, bson.M{"$match": bson.M{"$and": matchPipeline}})
	}
	if len(matchPipeline) == 1 {
		aggregationPipeline = append(aggregationPipeline, bson.M{"$match": matchPipeline[0]})
	}
	return aggregationPipeline
}

func addFilterToPipeline(aggregationPipeline *[]bson.M, filter apicontracts.FilterMetadata) {
	if !validFilter(filter, filter.MatchMode) {
		return
	}
	switch filter.MatchMode {
	case apicontracts.MatchModeContains:
		*aggregationPipeline = append(*aggregationPipeline, bson.M{filter.Field: bson.M{"$regex": filter.Value, "$options": "i"}})
	case apicontracts.MatchModeIn:
		*aggregationPipeline = append(*aggregationPipeline, bson.M{filter.Field: bson.M{"$in": filter.Value}})
	case apicontracts.MatchModeEquals:
		*aggregationPipeline = append(*aggregationPipeline, bson.M{filter.Field: bson.M{"$eq": filter.Value}})
	}
}

func validFilter(filter apicontracts.FilterMetadata, matchMode apicontracts.MatchModeType) bool {
	switch matchMode {
	case apicontracts.MatchModeContains, apicontracts.MatchModeEquals:
		return filter.Value != nil && filter.Value != "" && filter.Field != ""
	case apicontracts.MatchModeIn:
		filters, ok := filter.Value.([]any)
		return filter.Value != nil && ok && len(filters) > 0
	default:
		return false
	}
}

func customSort(sortMetadata apicontracts.SortMetadata) []bson.M {
	var aggregationPipeline []bson.M
	var comparison bson.A
	switch sortMetadata.SortField {
	case "ingresses.datacenter":
		comparison = bson.A{"$$ingressClass", "avi-ingress-class-datacenter"}
	case "ingresses.health":
		comparison = bson.A{"$$ingressClass", "avi-ingress-class-helsenett"}
	case "ingresses.internet":
		comparison = bson.A{"$$ingressClass", "avi-ingress-class-internett"}
	default:
		comparison = bson.A{}
	}
	set := bson.M{"$set": bson.M{"count": bson.M{"$size": bson.M{"$filter": bson.M{"input": "$ingresses.class", "as": "ingressClass", "cond": bson.M{"$eq": comparison}}}}}}
	sort := bson.M{"$sort": bson.M{"count": sortMetadata.SortOrder}}
	unset := bson.M{"$unset": "count"}
	aggregationPipeline = append(aggregationPipeline, set)
	aggregationPipeline = append(aggregationPipeline, sort)
	aggregationPipeline = append(aggregationPipeline, unset)
	return aggregationPipeline
}

func Contains[E comparable](s []E, v E) bool {
	return Index(s, v) >= 0
}

func Index[E comparable](s []E, v E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

func PrettyprintBSON(pipeline []primitive.M) {
	var prettyDocs []bson.M
	for _, doc := range pipeline {
		bsonDoc, err := bson.Marshal(doc)
		if err != nil {
			fmt.Println(err)
		}
		var prettyDoc bson.M
		err = bson.Unmarshal(bsonDoc, &prettyDoc)
		if err != nil {
			fmt.Println(err)
		}
		prettyDocs = append(prettyDocs, prettyDoc)
	}
	prettyJSON, err := json.MarshalIndent(prettyDocs, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(prettyJSON))
}
