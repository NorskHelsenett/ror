package rorresources

import (
	"cmp"
	"slices"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
)

const (
	FilterTypeString     FilterType     = "string"
	FilterTypeInt        FilterType     = "int"
	FilterTupeIntString  FilterType     = "intstring"
	FilterTypeBool       FilterType     = "bool"
	FilterTypeTime       FilterType     = "time"
	FilterTypeTimeString FilterType     = "timestring"
	FilterOperatorEq     FilterOperator = "eq"
	FilterOperatorNe     FilterOperator = "ne"
	FilterOperatorRegexp FilterOperator = "regexp"
	FilterOperatorGt     FilterOperator = "gt"
	FilterOperatorLt     FilterOperator = "lt"
	FilterOperatorGe     FilterOperator = "ge"
	FilterOperatorLe     FilterOperator = "le"
)

type FilterType string

type FilterOperator string

type ResourceQueryFilter struct {
	Field    string         `json:"field,omitempty"`
	Value    string         `json:"value,omitempty"`
	Type     FilterType     `json:"type,omitempty"`
	Operator FilterOperator `json:"operator,omitempty"`
}

type ResourceQueryOrder struct {
	Field      string `json:"field,omitempty"`
	Descending bool   `json:"descending,omitempty"`
	Index      int    `json:"index,omitempty"`
}

type ResourceQuery struct {
	VersionKind      schema.GroupVersionKind                      `json:"versionkind,omitempty"`      // memory getparam: apiversion, kind
	Uids             []string                                     `json:"uids,omitempty"`             // memory
	OwnerRefs        []rorresourceowner.RorResourceOwnerReference `json:"ownerrefs,omitempty"`        // memory
	Fields           []string                                     `json:"fields,omitempty"`           // post or db
	Order            []ResourceQueryOrder                         `json:"order,omitempty"`            // post or db
	Filters          []ResourceQueryFilter                        `json:"filters,omitempty"`          // db
	Offset           int                                          `json:"offset,omitempty"`           // post or db
	Limit            int                                          `json:"limit,omitempty"`            // post or db
	RelatedResources []ResourceQuery                              `json:"relatedresources,omitempty"` // memory or db
}

// func ParseResourceQuery(c *gin.Context) (*ResourceQuery, error) {
//     rq := NewResourceQuery()

//     // Parse apiversion and kind
//     rq.VersionKind.Group = schema.GroupVersion{Group: "", Version: c.Query("apiversion")} // Wrong
//     rq.VersionKind.Kind = c.Query("kind")

//     // Parse UIDs
//     if uids := c.Query("uids"); uids != "" {
//         rq.Uids = strings.Split(uids, ",")
//     }

//     // Parse Fields
//     if fields := c.Query("fields"); fields != "" {
//         rq.Fields = strings.Split(fields, ",")
//     }

//     // Parse Order
//     if order := c.Query("order"); order != "" {
//         var orders []ResourceQueryOrder
//         if err := json.Unmarshal([]byte(order), &orders); err != nil {
//             return nil, fmt.Errorf("invalid order parameter: %v", err)
//         }
//         rq.Order = orders
//     }

//     // Parse Filters
//     if filters := c.Query("filters"); filters != "" {
//         var filterList []ResourceQueryFilter
//         if err := json.Unmarshal([]byte(filters), &filterList); err != nil {
//             return nil, fmt.Errorf("invalid filters parameter: %v", err)
//         }
//         rq.Filters = filterList
//     }

//     // Parse Offset
//     if offset := c.Query("offset"); offset != "" {
//         if off, err := strconv.Atoi(offset); err == nil {
//             rq.Offset = off
//         }
//     }

//     // Parse Limit
//     if limit := c.Query("limit"); limit != "" {
//         if lim, err := strconv.Atoi(limit); err == nil {
//             rq.Limit = lim
//         }
//     }

//     // Parse RelatedResources
//     if relatedResources := c.Query("relatedresources"); relatedResources != "" {
//         var related []ResourceQuery
//         if err := json.Unmarshal([]byte(relatedResources), &related); err != nil {
//             return nil, fmt.Errorf("invalid relatedresources parameter: %v", err)
//         }
//         rq.RelatedResources = related
//     }

//	    return rq, nil
//	}
func NewResourceQuery() *ResourceQuery {
	return &ResourceQuery{
		Fields:           make([]string, 0),
		Order:            make([]ResourceQueryOrder, 0),
		Filters:          make([]ResourceQueryFilter, 0),
		RelatedResources: make([]ResourceQuery, 0),
		Limit:            100,
	}
}

func (rq *ResourceQuery) WithUID(uid string) *ResourceQuery {
	if rq.Uids == nil {
		rq.Uids = make([]string, 0)
	}
	rq.Uids = append(rq.Uids, uid)
	return rq
}

func (rq *ResourceQuery) SetLimit(limit int) *ResourceQuery {
	rq.Limit = limit
	return rq
}

func (rq ResourceQuery) GetOrderSorted() []ResourceQueryOrder {
	rqo := rq.Order
	slices.SortFunc(rqo, func(a, b ResourceQueryOrder) int {
		return cmp.Compare(a.Index, b.Index)
	})

	return rqo
}
