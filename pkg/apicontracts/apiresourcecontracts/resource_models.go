// package delivers apicontracts for resources
package apiresourcecontracts

import (
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

	"github.com/gin-gonic/gin"
)

// Allowed actions from the kubernetes dynamic client
type ResourceAction string

const (
	K8sActionAdd    ResourceAction = "Add"
	K8sActionDelete ResourceAction = "Delete"
	K8sActionUpdate ResourceAction = "Update"
)

// Deprecated: Use rortypes.RorResourceOwnerReference instead
// The ResourceOwnerReference or ownereref references the owner og a resource.
// Its used to chek acl and select resources for valid Scopes.

func NewResourceQueryFromClient(c *gin.Context) ResourceQuery {

	owner := rortypes.RorResourceOwnerReference{
		Scope:   aclmodels.Acl2Scope(c.Query("ownerScope")),
		Subject: aclmodels.Acl2Subject(c.Query("ownerSubject")),
	}

	query := ResourceQuery{
		Owner:      owner,
		Kind:       c.Query("kind"),
		ApiVersion: c.Query("apiversion"),
	}

	if query.Owner.Scope == aclmodels.Acl2ScopeRor {
		query.Global = true
	}

	return query
}

// Resource query used in the web facing resource services/repos
type ResourceQuery struct {
	Owner      rortypes.RorResourceOwnerReference `json:"owner"`
	ApiVersion string                             `json:"apiVersion"`
	Kind       string                             `json:"kind"`
	Uid        string                             `json:"uid"`
	Internal   bool
	Global     bool
}

// Resource update model to exchange resources. The value MUST  be casted to its explicit value before its saved to mongodb.
type ResourceUpdateModel struct {
	Owner      rortypes.RorResourceOwnerReference `json:"owner"`
	ApiVersion string                             `json:"apiVersion"`
	Kind       string                             `json:"kind"`
	Uid        string                             `json:"uid"`
	Action     ResourceAction                     `json:"action,omitempty"`
	Hash       string                             `json:"hash"`
	Resource   any                                `json:"resource"`
}

// Generic resourcemodels for single resource.
type ResourceModel[T Resourcetypes] struct {
	Owner      rortypes.RorResourceOwnerReference `json:"owner"`
	ApiVersion string                             `json:"apiVersion"`
	Kind       string                             `json:"kind"`
	Uid        string                             `json:"uid"`
	Hash       string                             `json:"hash"`
	Internal   bool                               `json:"internal"`
	Resource   T                                  `json:"resource"`
}

// Generic resourcemodels for multiple resources.
type ResourceModels[T Resourcetypes] struct {
	Owner      rortypes.RorResourceOwnerReference `json:"owner"`
	ApiVersion string                             `json:"apiVersion"`
	Kind       string                             `json:"kind"`
	Resources  []T                                `json:"resources"`
}

// K8s metadata struct
type ResourceMetadata struct {
	Name              string                           `json:"name"`
	ResourceVersion   string                           `json:"resourceVersion"`
	CreationTimestamp string                           `json:"creationTimestamp"`
	Labels            map[string]string                `json:"labels,omitempty"`
	Annotations       map[string]string                `json:"annotations,omitempty"`
	Uid               string                           `json:"uid"`
	Namespace         string                           `json:"namespace,omitempty"`
	Generation        int                              `json:"generation,omitempty"`
	OwnerReferences   []ResourceMetadataOwnerReference `json:"ownerReferences,omitempty"`
}

// K8s metadata.OwnerReferences struct
type ResourceMetadataOwnerReference struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Uid        string `json:"uid"`
}

// Hashlist for use in agent communication
// NB This struct has a counterpart in the agent.
type HashList struct {
	Items []HashItem `json:"items"`
}

// Item for use in the hashlist
// NB This struct has a counterpart in the agent.
type HashItem struct {
	Uid  string `json:"uid"`
	Hash string `json:"hash"`
}
