package rortypes

import (
	"sort"

	"github.com/NorskHelsenett/ror/pkg/helpers/optional"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// This file holds the patch mirror of CommonResource, the part every resource
// shares. Unlike the per kind mirrors it is hand written, because it is
// curated rather than a one to one reflection of the source struct:
//
//   - metav1.ObjectMeta is owned by kubernetes and most of it identifies the
//     resource. Only labels and annotations are patchable, the rest (name,
//     uid, namespace, resourceVersion, timestamps) is identity or bookkeeping
//     that a partial update must not rewrite.
//   - ResourceRorMeta.Hash is derived from the resource by GenRorHash, so it
//     is deliberately absent: patching it would only let a caller store a hash
//     that does not match the resource.
//   - TypeMeta is absent for the same reason. The kind and apiVersion of a
//     patch identify the target, they are not a payload.

// ResourcePatchCommon is the patch mirror of CommonResource.
type ResourcePatchCommon struct {
	Metadata optional.Optional[ResourcePatchObjectMeta] `json:"metadata,omitzero"`
	RorMeta  optional.Optional[ResourcePatchRorMeta]    `json:"rormeta,omitzero"`
}

// IsEmpty reports whether the patch would change nothing.
func (p ResourcePatchCommon) IsEmpty() bool {
	return !p.Metadata.IsSet() && !p.RorMeta.IsSet()
}

// ApplyTo writes the patch into target, leaving the fields the patch does not
// mention untouched.
func (p ResourcePatchCommon) ApplyTo(target *CommonResource) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	if err := optional.ApplyNested(p.Metadata, &target.Metadata); err != nil {
		return err
	}
	return optional.ApplyNested(p.RorMeta, &target.RorMeta)
}

// ResourcePatchObjectMeta is the patchable subset of metav1.ObjectMeta.
//
// Labels and annotations are patched per key, so removing one label does not
// require resending the others. Setting the field itself to null clears every
// label or annotation.
type ResourcePatchObjectMeta struct {
	Labels      optional.Optional[optional.Map[string, string]] `json:"labels,omitzero"`
	Annotations optional.Optional[optional.Map[string, string]] `json:"annotations,omitzero"`
}

// ApplyTo writes the patch into target, leaving the fields the patch does not
// mention untouched.
func (p ResourcePatchObjectMeta) ApplyTo(target *metav1.ObjectMeta) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	if err := optional.ApplyNested(p.Labels, &target.Labels); err != nil {
		return err
	}
	return optional.ApplyNested(p.Annotations, &target.Annotations)
}

// ResourcePatchRorMeta is the patch mirror of ResourceRorMeta, without the
// generated Hash field.
type ResourcePatchRorMeta struct {
	Version      optional.Optional[string]                                     `json:"version,omitzero"`
	LastReported optional.Optional[string]                                     `json:"lastReported,omitzero"`
	Internal     optional.Optional[bool]                                       `json:"internal,omitzero"`
	Ownerref     optional.Optional[rorresourceowner.RorResourceOwnerReference] `json:"ownerref,omitzero"`
	Action       optional.Optional[ResourceAction]                             `json:"action,omitzero"`
	Tags         optional.Optional[ResourcePatchTags]                          `json:"tags,omitzero"`
}

// ApplyTo writes the patch into target, leaving the fields the patch does not
// mention untouched.
func (p ResourcePatchRorMeta) ApplyTo(target *ResourceRorMeta) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	optional.Apply(p.Version, &target.Version)
	optional.Apply(p.LastReported, &target.LastReported)
	optional.Apply(p.Internal, &target.Internal)
	optional.Apply(p.Ownerref, &target.Ownerref)
	optional.Apply(p.Action, &target.Action)
	return optional.ApplyNested(p.Tags, &target.Tags)
}

// ResourcePatchTags patches the tag list of a resource by tag key.
//
// Tags are stored as a slice, but they behave like a map: the key of a tag
// identifies it. Patching them per key means a caller can add or drop a single
// tag without resending the whole list, which a slice replacement would
// require. Setting the Tags field itself to null removes every tag.
//
//	"tags": {"env": {"value": "test"}, "owner": null}
//
// adds or updates the tag "env" and removes the tag "owner".
type ResourcePatchTags map[string]optional.Optional[ResourcePatchTag]

// ApplyTo merges the patch into target. Tags the patch does not mention keep
// their position and value, updated tags keep their position and new tags are
// appended in key order.
func (p ResourcePatchTags) ApplyTo(target *[]ResourceTag) error {
	if target == nil || len(p) == 0 {
		return nil
	}

	patched := make([]ResourceTag, 0, len(*target)+len(p))
	seen := make(map[string]bool, len(*target))

	for _, tag := range *target {
		seen[tag.Key] = true

		entry, mentioned := p[tag.Key]
		if !mentioned || !entry.IsSet() {
			patched = append(patched, tag)
			continue
		}
		if entry.IsNull() {
			continue
		}

		tagPatch, _ := entry.Value()
		if err := tagPatch.ApplyTo(&tag); err != nil {
			return err
		}
		patched = append(patched, tag)
	}

	for _, key := range sortedKeys(p) {
		entry := p[key]
		if seen[key] || !entry.IsSet() || entry.IsNull() {
			continue
		}

		tagPatch, _ := entry.Value()
		tag := ResourceTag{Key: key}
		if err := tagPatch.ApplyTo(&tag); err != nil {
			return err
		}
		patched = append(patched, tag)
	}

	*target = patched
	return nil
}

// ResourcePatchTag is the patch mirror of ResourceTag, without the Key field
// which is the key of the ResourcePatchTags map.
type ResourcePatchTag struct {
	Value      optional.Optional[string]                                      `json:"value,omitzero"`
	Properties optional.Optional[optional.Map[ResourceTagProperties, string]] `json:"properties,omitzero"`
}

// ApplyTo writes the patch into target, leaving the fields the patch does not
// mention untouched.
func (p ResourcePatchTag) ApplyTo(target *ResourceTag) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	optional.Apply(p.Value, &target.Value)
	return optional.ApplyNested(p.Properties, &target.Properties)
}

// sortedKeys returns the keys of m in a stable order, so that tags added by a
// patch always land in the same order.
func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
