// Package rorresources contains the resource and resource set definitions used in the ror eco-system.
package rorresources

import (
	"encoding/json"
	"errors"
	"math"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// ResourceSet is the common way to present one or more resources in rors.
type ResourceSet struct {
	query      *ResourceQuery
	Resources  []*Resource `json:"resources,omitempty"`
	nextcursor int
}

func (rs *ResourceSet) SetQuery(query *ResourceQuery) {
	rs.query = query
}
func (rs *ResourceSet) GetQuery() *ResourceQuery {
	return rs.query
}

func NewResourceSet() *ResourceSet {
	return &ResourceSet{
		Resources:  make([]*Resource, 0),
		nextcursor: 0,
	}
}

// Add adds a resource to the ResourceSet
// If resource already exists it will be replaced.
func (rs *ResourceSet) Add(add *Resource) {
	if add == nil {
		return
	}
	if rs.GetByUid(add.GetUID()) != nil {
		rs.DeleteByUid(add.GetUID())
	}
	if rs.Resources == nil {
		rs.Resources = make([]*Resource, 0)
	}

	rs.Resources = append(rs.Resources, add)
}

var (
	ErrResourceIsNil = errors.New("resource is nil")
)

func (rs *ResourceSet) ToJson() ([]byte, error) {
	return json.Marshal(rs)
}

// Next moves the cursor along, use Get() to fetch the value eg:
//
//	    for resourceSet.Next(){
//		     stringhelpers.PrettyprintStruct(resourceSet.Get())
//	    }
func (rs *ResourceSet) Next() bool {
	if rs.nextcursor > (len(rs.Resources) - 1) {
		rs.nextcursor = 0
		return false
	}
	rs.nextcursor++
	return true
}

// Get returns the value of the current resource. can be used without moving the
// pointer in case the resourceset only contains one resource.
func (rs *ResourceSet) Get() *Resource {
	cursor := rs.nextcursor - 1
	if rs.nextcursor == 0 {
		cursor = 0
	}
	if cursor > len(rs.Resources) {
		rlog.Error("ResourceSet.Get() cursor out of bounds", nil)
		return nil
	}
	if rs.Resources == nil {
		return nil
	}
	return rs.Resources[cursor]
}

// All returns a slice with all resources
func (rs *ResourceSet) GetAll() []*Resource {
	return rs.Resources
}

// Len returns the number of resources in the ResourceSetr
func (rs *ResourceSet) Len() int {
	return len(rs.Resources)
}

// Function to return resource by name.
func (rs *ResourceSet) GetByName(search string) *Resource {
	for _, resource := range rs.Resources {
		if resource.GetName() == search {
			return resource
		}
	}
	return nil
}

// GetByUid to return resource by uid.
func (rs *ResourceSet) GetByUid(search string) *Resource {
	if rs == nil {
		return nil
	}
	if rs.Resources == nil {
		return nil
	}
	for _, resource := range rs.Resources {
		if resource.GetUID() == search {
			return resource
		}
	}
	return nil
}

// DeleteByUid to delete a resource by uid.
func (rs *ResourceSet) DeleteByUid(search string) {
	if search == "" {
		return
	}
	var newResources []*Resource
	for _, resource := range rs.Resources {
		if resource.GetUID() != search {
			newResources = append(newResources, resource)
		}
	}

	rs.Resources = newResources
	newResources = nil
}

// FilterByLabels returns a ResourceSet filtered by label.
func (rs *ResourceSet) FilterByLabels(search map[string]string) *ResourceSet {
	var response ResourceSet

	for _, resource := range rs.Resources {
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
func (rs *ResourceSet) FilterByAPIVersionKind(apiVersion string, kind string) *ResourceSet {
	var response ResourceSet

	for _, resource := range rs.Resources {
		if resource.GetAPIVersion() == apiVersion && resource.GetKind() == kind {
			response.Add(resource)
		}
	}
	return &response
}

// FilterByOwnerReference returns a ResourceSet filtered by ownerreference.
func (rs *ResourceSet) FilterByOwnerReference(ownerRef rorresourceowner.RorResourceOwnerReference) *ResourceSet {
	var response ResourceSet

	for _, resource := range rs.Resources {
		metadata := resource.GetRorMeta()
		if metadata.Ownerref.Scope == ownerRef.Scope && metadata.Ownerref.Subject == ownerRef.Subject {
			response.Add(resource)
		}
	}
	return &response
}

type ResourceUpdateResult struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

func NewResourceUpdateResult(status int, message string) *ResourceUpdateResult {
	return &ResourceUpdateResult{
		Status:  status,
		Message: message,
	}
}

type (
	config struct {
		size int
	}

	ConfigOption func(*config)
)

func newConfig() *config {
	return &config{
		size: 0,
	}
}

type ResourceUpdateResults struct {
	Results map[string]ResourceUpdateResult `json:"results,omitempty"`
}

// WithCapacity sets the map's initial capacity to the specified amount.
func WithCapacity(size uint) ConfigOption {
	if size > math.MaxInt {
		size = math.MaxInt
	}
	return func(c *config) {
		c.size = int(size)
	}
}

func NewResourceUpdateResults(opts ...ConfigOption) *ResourceUpdateResults {
	config := newConfig()
	for _, opt := range opts {
		opt(config)
	}
	return &ResourceUpdateResults{
		Results: make(map[string]ResourceUpdateResult, config.size),
	}
}

func (rs *ResourceUpdateResults) Add(key string, entry *ResourceUpdateResult) {
	rs.Results[key] = *entry
}

// GetFailedResources is a method to return a list of failed resources.
func (rs *ResourceUpdateResults) GetFailedResources() map[string]ResourceUpdateResult {
	failedResources := make(map[string]ResourceUpdateResult, 0)
	for key, value := range rs.Results {
		if value.Status > 399 || value.Status < 200 {
			failedResources[key] = value
		}
	}
	return failedResources
}
