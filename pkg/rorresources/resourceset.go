package rorresources

import (
	"encoding/json"

	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type FilterType string

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
	VersionKind         schema.GroupVersionKind              `json:"version_kind,omitempty"`
	Uids                []string                             `json:"uids,omitempty"`
	OwnerRefs           []rortypes.RorResourceOwnerReference `json:"owner_refs,omitempty"`
	Fields              []string                             `json:"fields,omitempty"`
	Order               map[int]ResourceQueryOrder           `json:"order,omitempty"`
	Filters             []ResourceQueryFilter                `json:"filters,omitempty"`
	AdditionalResources []schema.GroupVersionKind            `json:"additional_resources,omitempty"`
}

type ResourceUpdateResults struct {
	Results map[string]ResourceUpdateResult `json:"results,omitempty"`
}
type ResourceUpdateResult struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

// FailedResources is a method to return a list of failed resources.
func (r *ResourceUpdateResults) GetFailedResources() map[string]ResourceUpdateResult {
	failedResources := make(map[string]ResourceUpdateResult, 0)
	for key, value := range r.Results {
		if value.Status > 399 || value.Status < 200 {
			failedResources[key] = value
		}
	}
	return failedResources
}

// ResourceSet is the common way to present one or more resources in ror.
type ResourceSet struct {
	nextcursor int
	query      *ResourceQuery
	Resources  []*Resource `json:"resources,omitempty"`
}

func NewResourceSet() *ResourceSet {
	return &ResourceSet{
		Resources:  make([]*Resource, 0),
		nextcursor: 0,
	}
}

// Add adds a resource to the ResourceSet
// If resource already exists it will be replaced.
func (r *ResourceSet) Add(add *Resource) {
	if add == nil {
		return
	}
	if r.GetByUid(add.GetUID()) != nil {
		r.DeleteByUid(add.GetUID())
	}
	if r.Resources == nil {
		r.Resources = make([]*Resource, 0)
	}

	r.Resources = append(r.Resources, add)
}

func (r *ResourceSet) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

// Next moves the cursor along, use Get() to fetch the value eg:
//
//	    for resourceSet.Next(){
//		     stringhelper.PrettyprintStruct(resourceSet.Get())
//	    }
func (r *ResourceSet) Next() bool {
	if r.nextcursor > (len(r.Resources) - 1) {
		r.nextcursor = 0
		return false
	}
	r.nextcursor++
	return true
}

// Get returns the value of the current resource. can be used without moving the
// pointer in case the resourceset only contains one resource.
func (r *ResourceSet) Get() *Resource {
	cursor := r.nextcursor - 1
	if r.nextcursor == 0 {
		cursor = 0
	}
	if cursor > len(r.Resources) {
		rlog.Error("ResourceSet.Get() cursor out of bounds", nil)
		return nil
	}
	return r.Resources[cursor]
}

// All retuns a slice with all resources
func (r *ResourceSet) GetAll() []*Resource {
	return r.Resources
}

// Len returns the number of resources in the ResourceSetr
func (r *ResourceSet) Len() int {
	return len(r.Resources)
}

// Function to return resource by name.
func (r *ResourceSet) GetByName(search string) *Resource {
	for _, resource := range r.Resources {
		if resource.GetName() == search {
			return resource
		}
	}
	return nil
}

// Function to return resource by uid.
func (r *ResourceSet) GetByUid(search string) *Resource {
	if r == nil {
		return nil
	}
	if r.Resources == nil {
		return nil
	}
	for _, resource := range r.Resources {
		if resource.GetUID() == search {
			return resource
		}
	}
	return nil
}

// Function to delete resource by uid.
func (r *ResourceSet) DeleteByUid(search string) {
	if search == "" {
		return
	}
	var newResources []*Resource
	for _, resource := range r.Resources {
		if resource.GetUID() != search {
			newResources = append(newResources, resource)
		}
	}
	r.Resources = newResources
}

// FilterByLabels returns a ResourceSet filtered by label.
func (r *ResourceSet) FilterByLabels(search map[string]string) *ResourceSet {
	var response ResourceSet

	for _, resource := range r.Resources {
		metadata := resource.GetMetadata()
		if len(metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, metadata.Labels) {
				response.Add(resource)
			}
		}
	}
	return &response
}

// FilterByAPIVersionKind returns a ResourceSet filtered by apiversion and kind.
func (r *ResourceSet) FilterByAPIVersionKind(apiVersion string, kind string) *ResourceSet {
	var response ResourceSet

	for _, resource := range r.Resources {
		if resource.GetAPIVersion() == apiVersion && resource.GetKind() == kind {
			response.Add(resource)
		}
	}
	return &response
}

// FilterByOwnerReference returns a ResourceSet filtered by ownerreference.
func (r *ResourceSet) FilterByOwnerReference(ownerRef rortypes.RorResourceOwnerReference) *ResourceSet {
	var response ResourceSet

	for _, resource := range r.Resources {
		metadata := resource.GetRorMeta()
		if metadata.Ownerref.Scope == ownerRef.Scope && metadata.Ownerref.Subject == ownerRef.Subject {
			response.Add(resource)
		}
	}
	return &response
}
