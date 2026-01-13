// package delivers apicontracts for resources in the V1 API
// Deprecated: This package is deprecated. Use rortypes instead.
package apiresourcecontracts

import (
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

// Allowed actions from the kubernetes dynamic client// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceAction string

const (
	K8sActionAdd    ResourceAction = "Add"
	K8sActionDelete ResourceAction = "Delete"
	K8sActionUpdate ResourceAction = "Update"
)

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVersion string

const (
	ResourceVersionV1 ResourceVersion = "v1"
	ResourceVersionV2 ResourceVersion = "v2"
)

// Deprecated: Use rorresourceowner.RorResourceOwnerReference instead
// The ResourceOwnerReference or ownereref references the owner og a resource.
// Its used to chek acl and select resources for valid Scopes.// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceOwnerReference struct {
	Scope   aclmodels.Acl2Scope `json:"scope"`   // cluster, workspace,...
	Subject string              `json:"subject"` // ror id eg clusterId or workspaceName
}

// Returns a map to use in the `*Resty.Request.SetQueryParams(<ResourceOwnerReference>.GetQueryParams())“ function
func (m ResourceOwnerReference) GetQueryParams() map[string]string {
	response := make(map[string]string)

	response["ownerScope"] = string(m.Scope)
	response["ownerSubject"] = m.Subject
	return response
}

// Resource query used in the web facing resource services/repos// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceQuery struct {
	Owner      ResourceOwnerReference `json:"owner"`
	ApiVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Uid        string                 `json:"uid"`
	Internal   bool
	Global     bool
}

// Resource update model to exchange resources. The value MUST  be casted to its explicit value before its saved to mongodb.// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceUpdateModel struct {
	Owner      ResourceOwnerReference `json:"owner"`
	ApiVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Uid        string                 `json:"uid"`
	Action     ResourceAction         `json:"action,omitempty"`
	Hash       string                 `json:"hash"`
	Version    ResourceVersion        `json:"version"`
	Resource   any                    `json:"resource"`
}

// Generic resourcemodels for single resource.// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceModel[T Resourcetypes] struct {
	Owner      ResourceOwnerReference `json:"owner"`
	ApiVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Uid        string                 `json:"uid"`
	Hash       string                 `json:"hash"`
	Internal   bool                   `json:"internal"`
	Version    ResourceVersion        `json:"version"`
	Resource   T                      `json:"resource"`
}

// Generic resourcemodels for multiple resources.// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceModels[T Resourcetypes] struct {
	Owner      ResourceOwnerReference `json:"owner"`
	ApiVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Resources  []T                    `json:"resources"`
}

// K8s metadata struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
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

// K8s metadata.OwnerReferences struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceMetadataOwnerReference struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Uid        string `json:"uid"`
}

// Hashlist for use in agent communication
// NB This struct has a counterpart in the agent.// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type HashList struct {
	Items []HashItem `json:"items"`
}

// Item for use in the hashlist
// NB This struct has a counterpart in the agent.// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type HashItem struct {
	Uid  string `json:"uid"`
	Hash string `json:"hash"`
}
