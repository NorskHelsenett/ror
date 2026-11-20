// THIS FILE IS GENERATED, DO NOT EDIT
// ref: cmd/generator/main.go

package rorresources

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/helpers/optional"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rordefs"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ErrPatchNil                  = errors.New("resource patch is nil")
	ErrPatchNilResource          = errors.New("resource to patch is nil")
	ErrPatchEmptyUID             = errors.New("resource patch has no uid")
	ErrPatchUIDMismatch          = errors.New("resource patch uid does not match the resource")
	ErrPatchKindMismatch         = errors.New("resource patch kind does not match the resource")
	ErrPatchEmpty                = errors.New("resource patch would change nothing")
	ErrPatchNoSubResource        = errors.New("resource has no sub resource to patch")
	ErrPatchKindNotAllowed       = errors.New("resource patch carries a sub resource of another kind")
	ErrPatchMultipleSubResources = errors.New("resource patch carries more than one sub resource")
)

// ResourcePatch is the partial update counterpart of Resource.
//
// Resource cannot fully express a partial update, as they lack the ability to distinguish between
// between unset and null. ResourcePatch carries patch mirrors
// instead, where every field is an optional.Optional and therefore holds the
// three states a partial update needs: unset, null and a value.
//
// A patch targets exactly one resource, identified by kind, apiVersion and uid.
// It carries two independent parts:
//
//   - the common part (labels, annotations, ror metadata and tags), embedded
//     here and patchable.
//   - the sub resource of the target kind, patchable on the kinds marked
//     Patchable in rordefs, which are the kinds with a generated patch mirror.
//
// Example usage:
//
//	patch := rorresources.NewResourcePatchKubernetesCluster(uid)
//	patch.SetKubernetesCluster(rortypes.ResourcePatchKubernetesCluster{
//	    Spec: optional.NewOptional(rortypes.ResourcePatchKubernetesClusterSpec{
//	        ClusterMetadata: optional.NewOptional(rortypes.ResourcePatchKubernetesClusterSpecMetadataDetails{
//	            Description: optional.NewOptional("a new description"),
//	            Contacts:    optional.Null[[]rortypes.ResourceContact](),
//	        }),
//	    }),
//	})
//	patched, err := patch.Apply(current)
type ResourcePatch struct {
	metav1.TypeMeta              `json:",inline"`
	rortypes.ResourcePatchCommon `json:",inline"`

	UID string `json:"uid"`

	KubernetesClusterResource  optional.Optional[rortypes.ResourcePatchKubernetesCluster]  `json:"kubernetescluster,omitzero"`
	OrganizationalUnitResource optional.Optional[rortypes.ResourcePatchOrganizationalUnit] `json:"organizationalunit,omitzero"`
}

// NewResourcePatch returns an empty patch for a resource of a given kind/apiversion.
func NewResourcePatch(kind string, apiversion string, uid string) *ResourcePatch {
	p := ResourcePatch{
		UID:        uid,
		Kind:       kind,
		APIVersion: apiversion,
	}
	return &p
}

// NewResourcePatchKubernetesCluster returns an empty patch for the kubernetescluster with the given uid.
func NewResourcePatchKubernetesCluster(uid string) *ResourcePatch {
	return NewResourcePatch(rordefs.ResourceKubernetesCluster.Kind, rordefs.ResourceKubernetesCluster.APIVersion, uid)
}

// SetKubernetesCluster sets the kubernetescluster part of the patch.
func (p *ResourcePatch) SetKubernetesCluster(patch rortypes.ResourcePatchKubernetesCluster) *ResourcePatch {
	p.KubernetesClusterResource.Set(patch)
	return p
}

// KubernetesCluster returns the kubernetescluster part of the patch and whether it holds a
// value. Use IsNull on the field to tell an unset sub resource from one that is
// explicitly set to null.
func (p *ResourcePatch) KubernetesCluster() (rortypes.ResourcePatchKubernetesCluster, bool) {
	return p.KubernetesClusterResource.Value()
}

// NewResourcePatchOrganizationalUnit returns an empty patch for the organizationalunit with the given uid.
func NewResourcePatchOrganizationalUnit(uid string) *ResourcePatch {
	return NewResourcePatch(rordefs.ResourceOrganizationalUnit.Kind, rordefs.ResourceOrganizationalUnit.APIVersion, uid)
}

// SetOrganizationalUnit sets the organizationalunit part of the patch.
func (p *ResourcePatch) SetOrganizationalUnit(patch rortypes.ResourcePatchOrganizationalUnit) *ResourcePatch {
	p.OrganizationalUnitResource.Set(patch)
	return p
}

// OrganizationalUnit returns the organizationalunit part of the patch and whether it holds a
// value. Use IsNull on the field to tell an unset sub resource from one that is
// explicitly set to null.
func (p *ResourcePatch) OrganizationalUnit() (rortypes.ResourcePatchOrganizationalUnit, bool) {
	return p.OrganizationalUnitResource.Value()
}

// subResourceKinds returns the kinds of the sub resources the patch carries.
func (p *ResourcePatch) subResourceKinds() []string {
	var kinds []string
	if p.KubernetesClusterResource.IsSet() {
		kinds = append(kinds, rordefs.ResourceKubernetesCluster.Kind)
	}
	if p.OrganizationalUnitResource.IsSet() {
		kinds = append(kinds, rordefs.ResourceOrganizationalUnit.Kind)
	}
	return kinds
}

// IsEmpty reports whether the patch would change nothing.
func (p *ResourcePatch) IsEmpty() bool {
	if p == nil {
		return true
	}
	return p.ResourcePatchCommon.IsEmpty() && len(p.subResourceKinds()) == 0
}

// Validate reports whether the patch identifies a resource and carries a sub
// resource that matches its kind.
func (p *ResourcePatch) Validate() error {
	if p == nil {
		return ErrPatchNil
	}
	if p.UID == "" {
		return ErrPatchEmptyUID
	}
	if p.Kind == "" || p.APIVersion == "" {
		return ErrPatchKindMismatch
	}
	if p.IsEmpty() {
		return ErrPatchEmpty
	}

	kinds := p.subResourceKinds()
	if len(kinds) > 1 {
		return fmt.Errorf("%w: %v", ErrPatchMultipleSubResources, kinds)
	}
	if len(kinds) == 1 && kinds[0] != p.Kind {
		return fmt.Errorf("%w: %q on a patch for %q", ErrPatchKindNotAllowed, kinds[0], p.Kind)
	}

	return nil
}

// cloneResource returns a deep copy of in with its common interface rebound, so
// that applying a patch never reaches the resource the caller passed in.
func cloneResource(in *Resource) (*Resource, error) {
	encoded, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("could not encode resource: %w", err)
	}
	var cloned Resource
	if err := json.Unmarshal(encoded, &cloned); err != nil {
		return nil, fmt.Errorf("could not decode resource: %w", err)
	}
	return NewResourceFromStruct(cloned)
}
