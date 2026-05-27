package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// K8s applicationProject struct used with ArgoCD
type ResourceAppProject struct {
	Spec ResourceAppProjectSpec `json:"spec"`
}
type ResourceAppProjectSpec struct {
	Description  string                               `json:"description"`
	SourceRepos  []string                             `json:"sourceRepos"`
	Destinations []ResourceApplicationSpecDestination `json:"destinations"`
}

// (r *ResourceAppProject) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceAppProject) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalAppProjects[cr.Metadata.Name] {
		cr.RorMeta.Internal = true
	}
	return nil
}

// (r ResourceAppProject) Get returns a pointer to the resource of type ResourceAppProject
func (r *ResourceAppProject) Get() *ResourceAppProject {
	return r
}
