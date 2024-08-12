package rorresources

import (
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	FilterTypeString FilterType = "string"
	FilterTypeInt    FilterType = "int"
	FilterTypeBool   FilterType = "bool"
)

type ResourceQueryFilter struct {
	Field    string     `json:"field,omitempty"`
	Value    string     `json:"value,omitempty"`
	Type     FilterType `json:"type,omitempty"`
	Operator string     `json:"operator,omitempty"`
}

type ResourceQueryOrder struct {
	Field      string `json:"field,omitempty"`
	Descending bool   `json:"descending,omitempty"`
}

type ResourceQuery struct {
	VersionKind         schema.GroupVersionKind              `json:"version_kind,omitempty"`         // memory
	Uids                []string                             `json:"uids,omitempty"`                 // memory
	OwnerRefs           []rortypes.RorResourceOwnerReference `json:"owner_refs,omitempty"`           // memory
	Fields              []string                             `json:"fields,omitempty"`               // post or db
	Order               map[int]ResourceQueryOrder           `json:"order,omitempty"`                // post or db
	Filters             []ResourceQueryFilter                `json:"filters,omitempty"`              // db
	Offset              int                                  `json:"offset,omitempty"`               // post or db
	Limit               int                                  `json:"limit,omitempty"`                // post or db
	AdditionalResources []schema.GroupVersionKind            `json:"additional_resources,omitempty"` // memory or db
}

func NewResourceQuery() *ResourceQuery {
	return &ResourceQuery{
		Fields:              make([]string, 0),
		Order:               make(map[int]ResourceQueryOrder),
		Filters:             make([]ResourceQueryFilter, 0),
		AdditionalResources: make([]schema.GroupVersionKind, 0),
	}
}

func (rq *ResourceQuery) WithUID(uid string) *ResourceQuery {
	if rq.Uids == nil {
		rq.Uids = make([]string, 0)
	}
	rq.Uids = append(rq.Uids, uid)
	return rq
}
