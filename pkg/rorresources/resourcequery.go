package rorresources

import (
	"cmp"
	"slices"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
)

const (
	FilterTypeString     FilterType = "string"
	FilterTypeInt        FilterType = "int"
	FilterTupeIntString  FilterType = "intstring"
	FilterTypeBool       FilterType = "bool"
	FilterTypeTime       FilterType = "time"
	FilterTypeTimeString FilterType = "timestring"

	// FilterOperatorEq is the equal operator.
	// It looks for an exact value.
	FilterOperatorEq FilterOperator = "eq"

	// FilterOperatorEq is the not equal operator.
	// It looks for not exact value.
	FilterOperatorNe FilterOperator = "ne"

	// FilterOperatorRegexp is the regex operator.
	// It which searched with regex based on the value.
	FilterOperatorRegexp FilterOperator = "regexp"

	// FilterOperatorGt is the greater than operator.
	// It searches for the value in the Field paramter that is greater than the Value parameter.
	FilterOperatorGt FilterOperator = "gt"

	// FilterOperatorLt is the less than operator.
	// It searches for the value in the Field paramter that is less than the Value parameter.
	FilterOperatorLt FilterOperator = "lt"

	// FilterOperatorGe is the greater or equal operator.
	// It searches for the value in the Field paramter that is greater or equal to the Value parameter.
	FilterOperatorGe FilterOperator = "ge"

	// FilterOperatorLe is the less or equal operator.
	// It searches for the value in the Field paramter that is less or equal to the Value parameter.
	FilterOperatorLe FilterOperator = "le"
)

// FilterType is the type of the value you're looking for.
type FilterType string

// FilterOperator is the operator in how the search is done.
type FilterOperator string

// ResourceQueryFilter is a definition of a filter.
//
// For an truncated example resource:
//
//	{
//	    _id: ObjectId('68d3daefc30906fc9314d1d5'),
//	    uid: '6035f8f2-eb39-56c0-a5dc-0bc1c3d3ff07',
//	    backuprun: {
//	        id: '16835:1758207900971270',
//	        status: {
//	            backupjobid: '4923908281402464:1614676439887:16835',
//	            backuptargets: [
//	                {
//	                    name: 'something-name',
//	                    id: '706',
//	                    externalid: '503f1b69-7998-b7bb-5cc2-56297656e04d',
//	                }
//	            ]
//	        },
//	    },
//	    metadata: {
//	        uid: '6035f8f2-eb39-56c0-a5dc-0bc1c3d3ff07',
//	        creationtimestamp: {
//	            time: ISODate('0001-01-01T00:00:00.000Z')
//	        },
//	    },
//	    typemeta: {
//	        kind: 'BackupRun',
//	        apiversion: 'backup.ror.internal/v1alpha1'
//	    }
//	}
//
// Exmaple to query based on a field within the object:
//
//	filter := rorresources.ResourceQueryFilter{
//	        Field:    "backupjob.status.id",
//	        Value:    "16835:1758207900971270",
//	        Type:     rorresources.FilterTypeString,
//	        Operator: rorresources.FilterOperatorEq,
//	}
//
// Exmaple to query based on a field within the metadata:
//
//	filter := rorresources.ResourceQueryFilter{
//	        Field:    "metadata.uid",
//	        Value:    "6035f8f2-eb39-56c0-a5dc-0bc1c3d3ff07",
//	        Type:     rorresources.FilterTypeString,
//	        Operator: rorresources.FilterOperatorEq,
//	}.
type ResourceQueryFilter struct {

	// Field parameter starts based on the base of the object from the database view.
	// This parameter is case sensitive.
	Field string `json:"field,omitempty"`

	// Value parameter is the value you're searching with based on which Operator you use.
	// This parameter is case sensitive.
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
