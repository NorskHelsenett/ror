package rorresources

import (
	"encoding/json"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// ResourcePatchSet is the common way to present one or more patch resources in ror.
type ResourcePatchSet struct {
	query      *ResourceQuery
	Resources  []*ResourcePatch `json:"resources,omitempty"`
	nextcursor int
}

func (rs *ResourcePatchSet) SetQuery(query *ResourceQuery) {
	rs.query = query
}
func (rs *ResourcePatchSet) GetQuery() *ResourceQuery {
	return rs.query
}

func NewResourcePatchSet() *ResourcePatchSet {
	return &ResourcePatchSet{
		Resources:  make([]*ResourcePatch, 0),
		nextcursor: 0,
	}
}

// Add adds a patch resource to the ResourceSet
func (rs *ResourcePatchSet) Add(add *ResourcePatch) error {
	if add == nil {
		return ErrResourceIsNil
	}
	if rs.Resources == nil {
		rs.Resources = make([]*ResourcePatch, 0)
	}

	rs.Resources = append(rs.Resources, add)

	return nil
}

func (rs *ResourcePatchSet) Upsert(add *ResourcePatch) error {
	if add == nil {
		return ErrResourceIsNil
	}
	if rs.Resources == nil {
		rs.Resources = make([]*ResourcePatch, 0)
	}

	for i, patch := range rs.Resources {
		if patch.GetUID() == add.GetUID() {
			rs.Resources[i] = add
			return nil
		}
	}

	rs.Resources = append(rs.Resources, add)

	return nil
}

func (rs *ResourcePatchSet) ToJSON() ([]byte, error) {
	return json.Marshal(rs)
}

// Next moves the cursor along, use Get() to fetch the value eg:
//
//	    for resourceSet.Next(){
//		     stringhelper.PrettyprintStruct(resourceSet.Get())
//	    }
func (rs *ResourcePatchSet) Next() bool {
	if rs.nextcursor > (len(rs.Resources) - 1) {
		rs.nextcursor = 0
		return false
	}
	rs.nextcursor++
	return true
}

// Get returns the value of the current patch resource. can be used without moving the
// pointer in case the resourceset only contains one patch resource.
func (rs *ResourcePatchSet) Get() *ResourcePatch {
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

// GetAll returns a slice with all patch resources
func (rs *ResourcePatchSet) GetAll() []*ResourcePatch {
	return rs.Resources
}

// Len returns the number of patch resources in the ResourceSet
func (rs *ResourcePatchSet) Len() int {
	return len(rs.Resources)
}

// // Function to return resource by name.
// func (r *ResourcePatchSet) GetByName(search string) *ResourcePatch {
// 	for _, resource := range r.Resources {
// 		if resource.GetName() == search {
// 			return resource
// 		}
// 	}
// 	return nil
// }

// GetByUID returns a resource by uid.
// Since the UID is supposed to be unique, this function will return the first resource that matches the uid.
// If there's multiple resources with the same uid, only the first one will be returned.
func (rs *ResourcePatchSet) GetByUID(search string) *ResourcePatch {
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

// DeleteByUID removes a resource from the ResourceSet by its UID.
// Since the UID is supposed to be unique, this function will remove the first resource that matches the uid.
// If there's multiple resources with the same uid, only the first one will be removed.
func (rs *ResourcePatchSet) DeleteByUID(search string) {
	if search == "" {
		return
	}
	var newResources []*ResourcePatch
	for _, resource := range rs.Resources {
		if resource.GetUID() != search {
			newResources = append(newResources, resource)
		}
	}

	rs.Resources = newResources
	newResources = nil
}

// FilterByLabels returns a ResourceSet filtered by label.
func (rs *ResourcePatchSet) FilterByLabels(search map[string]string) *ResourcePatchSet {
	var response ResourcePatchSet

	for _, resource := range rs.Resources {
		labels := resource.labels()
		if len(labels) != 0 {
			if stringhelper.CompareOptionalLabels(search, labels) {
				response.Add(resource)
			}
		}
	}
	return &response
}

// FilterByAPIVersionKind returns a ResourceSet filtered by apiversion and kind.
func (rs *ResourcePatchSet) FilterByAPIVersionKind(apiVersion string, kind string) *ResourcePatchSet {
	var response ResourcePatchSet

	for _, resource := range rs.Resources {
		if resource.GetAPIVersion() == apiVersion && resource.GetKind() == kind {
			response.Add(resource)
		}
	}
	return &response
}

// FilterByOwnerReference returns a ResourceSet filtered by ownerreference.
func (rs *ResourcePatchSet) FilterByOwnerReference(ownerRef rorresourceowner.RorResourceOwnerReference) *ResourcePatchSet {
	var response ResourcePatchSet

	for _, resource := range rs.Resources {
		metadata := resource.GetRorMeta()
		if rormeta, ok := metadata.Value(); ok {
			ownerref, ok := rormeta.Ownerref.Value()
			if ok && ownerref.Scope == ownerRef.Scope && ownerref.Subject == ownerRef.Subject {
				response.Add(resource)
			}
		}
	}
	return &response
}
