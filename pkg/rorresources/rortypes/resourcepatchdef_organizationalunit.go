// THIS FILE IS GENERATED, DO NOT EDIT
// ref: cmd/generator/main.go

package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/helpers/optional"
)

// This file holds a patch mirror of a resource. a patch mirror mirrors every field
// but the field are a generic type of optional.Optional, which carries the three
// states a partial update needs: unset, null, and a value.
// See the optional package for why every field is tagged omitzero.

// ResourcePatchOrganizationalUnit is the patch mirror of ResourceOrganizationalUnit.
type ResourcePatchOrganizationalUnit struct {
	ID     optional.Optional[string]                                `json:"id,omitzero"`
	Spec   optional.Optional[ResourcePatchOrganizationalUnitSpec]   `json:"spec,omitzero"`
	Status optional.Optional[ResourcePatchOrganizationalUnitStatus] `json:"status,omitzero"`
}

// ApplyTo writes the patch into target, leaving the fields the patch does
// not mention untouched.
func (p ResourcePatchOrganizationalUnit) ApplyTo(target *ResourceOrganizationalUnit) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	optional.Apply(p.ID, &target.ID)
	if err := optional.ApplyNested(p.Spec, &target.Spec); err != nil {
		return err
	}
	if err := optional.ApplyNested(p.Status, &target.Status); err != nil {
		return err
	}
	return nil
}

// ResourcePatchOrganizationalUnitSpec is the patch mirror of ResourceOrganizationalUnitSpec.
type ResourcePatchOrganizationalUnitSpec struct {
	Name optional.Optional[string]                 `json:"name,omitzero"`
	Type optional.Optional[OrganizationalUnitType] `json:"type,omitzero"`
}

// ApplyTo writes the patch into target, leaving the fields the patch does
// not mention untouched.
func (p ResourcePatchOrganizationalUnitSpec) ApplyTo(target *ResourceOrganizationalUnitSpec) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	optional.Apply(p.Name, &target.Name)
	optional.Apply(p.Type, &target.Type)
	return nil
}

// ResourcePatchOrganizationalUnitStatus is the patch mirror of ResourceOrganizationalUnitStatus.
type ResourcePatchOrganizationalUnitStatus struct {
	Name optional.Optional[string]                 `json:"name,omitzero"`
	Type optional.Optional[OrganizationalUnitType] `json:"type,omitzero"`
}

// ApplyTo writes the patch into target, leaving the fields the patch does
// not mention untouched.
func (p ResourcePatchOrganizationalUnitStatus) ApplyTo(target *ResourceOrganizationalUnitStatus) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	optional.Apply(p.Name, &target.Name)
	optional.Apply(p.Type, &target.Type)
	return nil
}

// OrganizationalUnitPatchinterface represents the interface for patches of the type organizationalunit,
// implemented by ResourcePatchOrganizationalUnit.
type OrganizationalUnitPatchinterface interface {
	ApplyTo(target *ResourceOrganizationalUnit) error
}
