package rorresources

import (
	"github.com/NorskHelsenett/ror/pkg/helpers/optional"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
)

// This file holds helpers for the common part of a patch like labels, annotations, tags.
//
// Example:
//	patch := rorresources.NewResourcePatchKubernetesCluster(uid).
//	    SetLabel("team", "platform").
//	    RemoveTag("deprecated").
//	    SetKubernetesCluster(clusterPatch)

// SetLabel adds a label to the patch.
func (p *ResourcePatch) SetLabel(key string, value string) *ResourcePatch {
	p.labels().Set(key, value)
	return p
}

// RemoveLabel removes a label.
func (p *ResourcePatch) RemoveLabel(key string) *ResourcePatch {
	p.labels().Delete(key)
	return p
}

// SetAnnotation adds an annotation to the patch.
func (p *ResourcePatch) SetAnnotation(key string, value string) *ResourcePatch {
	p.annotations().Set(key, value)
	return p
}

// RemoveAnnotation removes an annotation.
func (p *ResourcePatch) RemoveAnnotation(key string) *ResourcePatch {
	p.annotations().Delete(key)
	return p
}

// SetTag adds or updates a tag by key.
func (p *ResourcePatch) SetTag(key string, value string) *ResourcePatch {
	p.tags()[key] = optional.NewOptional(rortypes.ResourcePatchTag{
		Value: optional.NewOptional(value),
	})
	return p
}

// RemoveTag removes a tag by key.
func (p *ResourcePatch) RemoveTag(key string) *ResourcePatch {
	p.tags()[key] = optional.Null[rortypes.ResourcePatchTag]()
	return p
}

// SetOwnerref sets the owner reference of the resource.
func (p *ResourcePatch) SetOwnerref(ownerref rorresourceowner.RorResourceOwnerReference) *ResourcePatch {
	meta, _ := p.RorMeta.Value()
	meta.Ownerref.Set(ownerref)
	p.RorMeta.Set(meta)
	return p
}

// labels returns the label map, if the labels are empty, it initalizes it and returns the empty label map.
func (p *ResourcePatch) labels() optional.Map[string, string] {
	meta, _ := p.Metadata.Value()
	labels, ok := meta.Labels.Value()
	if !ok {
		labels = optional.Map[string, string]{}
		meta.Labels.Set(labels)
	}
	p.Metadata.Set(meta)
	return labels
}

// annotations returns the annotations map, if the annotations are empty, it initalizes it
// and returns the empty annotations map.
func (p *ResourcePatch) annotations() optional.Map[string, string] {
	meta, _ := p.Metadata.Value()
	annotations, ok := meta.Annotations.Value()
	if !ok {
		annotations = optional.Map[string, string]{}
		meta.Annotations.Set(annotations)
	}
	p.Metadata.Set(meta)
	return annotations
}

// tags returns the tags map, if the tags are empty, it initalizes it
// and returns the empty tags map.
func (p *ResourcePatch) tags() rortypes.ResourcePatchTags {
	meta, _ := p.RorMeta.Value()
	tags, ok := meta.Tags.Value()
	if !ok {
		tags = rortypes.ResourcePatchTags{}
		meta.Tags.Set(tags)
	}
	p.RorMeta.Set(meta)
	return tags
}
