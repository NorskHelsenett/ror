package apiresourcecontracts

import (
	"github.com/NorskHelsenett/ror/pkg/models/providers"
)

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceClusterOrder struct {
	ApiVersion string                     `json:"apiVersion"`
	Kind       string                     `json:"kind"`
	Metadata   ResourceMetadata           `json:"metadata"`
	Spec       ResourceClusterOrderSpec   `json:"spec"`
	Status     ResourceClusterOrderStatus `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceClusterOrderSpec struct {
	OrderType ResourceActionType     `json:"orderType" validate:"required,min=1,ne=' '"`
	Provider  providers.ProviderType `json:"provider,omitempty"`
	Cluster   string                 `json:"cluster" validate:"required,min=1,ne=' '"`
	ProjectId string                 `json:"projectId,omitempty"`
	OrderBy   string                 `json:"orderBy" validate:"required,min=1,ne=' '"`

	Environment      EnvironmentType                    `json:"environment,omitempty" validate:"min=1,max=4"`
	Criticality      CriticalityLevel                   `json:"criticality,omitempty" validate:"min=1,max=4"`
	Sensitivity      SensitivityLevel                   `json:"sensitivity,omitempty" validate:"min=1,max=4"`
	HighAvailability bool                               `json:"highAvailability,omitempty" validate:"boolean"`
	NodePools        []ResourceClusterOrderSpecNodePool `json:"nodePools,omitempty" validate:"min=1,dive,required"`
	ServiceTags      map[string]string                  `json:"serviceTags,omitempty"`
	ProviderConfig   map[string]interface{}             `json:"providerConfig,omitempty"`
	OwnerGroup       string                             `json:"ownerGroup,omitempty" validate:"min=1,ne=' '"`
	K8sVersion       string                             `json:"k8sVersion,omitempty"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceProviderConfigTanzu struct {
	DatacenterId string `json:"datacenterId" validate:"required,min=1,ne=' '"`
	NamespaceId  string `json:"namespaceId" validate:"required,min=1,ne=' '"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceProviderConfigAks struct {
	Region        string `json:"region" validate:"required,min=1,ne=' '"`
	Subscription  string `json:"subscription" validate:"required,min=1,ne=' '"`
	ResourceGroup string `json:"resourceGroup" validate:"required,min=1,ne=' '"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceProviderConfigKind struct {
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceClusterOrderSpecNodePool struct {
	Name         string `json:"name" validate:"required,min=1,ne=' '"`
	MachineClass string `json:"machineClass" validate:"required,min=1,ne=' '"`
	Count        int    `json:"count" validate:"required,min=1"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceClusterOrderStatus struct {
	Status           string                                     `json:"status"`
	Phase            ResourceClusterOrderStatusPhase            `json:"phase"`
	Conditions       []ResourceKubernetesClusterStatusCondition `json:"conditions"`
	CreatedTime      string                                     `json:"createdTime"`
	UpdatedTime      string                                     `json:"updatedTime"`
	LastObservedTime string                                     `json:"lastObservedTime"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceClusterOrderCondition struct {
	Type               string `json:"type"`
	Status             string `json:"status"`
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Message            string `json:"message"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceClusterOrderStatusPhase string

const (
	ResourceClusterOrderStatusPhasePending ResourceClusterOrderStatusPhase = "Pending"

	ResourceClusterOrderStatusPhaseCreating ResourceClusterOrderStatusPhase = "Creating"
	ResourceClusterOrderStatusPhaseUpdating ResourceClusterOrderStatusPhase = "Updating"
	ResourceClusterOrderStatusPhaseDeleting ResourceClusterOrderStatusPhase = "Deleting"

	ResourceClusterOrderStatusPhaseCompleted ResourceClusterOrderStatusPhase = "Completed"
	ResourceClusterOrderStatusPhaseFailed    ResourceClusterOrderStatusPhase = "Failed"
)

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceActionType string

const (
	ResourceActionTypeUnknown ResourceActionType = ""
	ResourceActionTypeCreate  ResourceActionType = "Create"
	ResourceActionTypeUpdate  ResourceActionType = "Update"
	ResourceActionTypeDelete  ResourceActionType = "Delete"
)

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type EnvironmentType int

const (
	EnvironmentUnknown     EnvironmentType = iota
	EnvironmentDevelopment EnvironmentType = 1
	EnvironmentTesting     EnvironmentType = 2
	EnvironmentQA          EnvironmentType = 3
	EnvironmentProduction  EnvironmentType = 4
)

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type SensitivityLevel int

const (
	SensitivityLevelUnknown  SensitivityLevel = iota
	SensitivityLevelNormal   SensitivityLevel = 1
	SensitivityLevelModerate SensitivityLevel = 2
	SensitivityLevelHigh     SensitivityLevel = 3
	SensitivityLevelCritical SensitivityLevel = 4
)

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type CriticalityLevel int

const (
	CriticalityLevelUnknown        CriticalityLevel = iota
	CriticalityLevelOpen           CriticalityLevel = 1
	CriticalityLevelIntern         CriticalityLevel = 2
	CriticalityLevelShielded       CriticalityLevel = 3
	CriticalityLevelHighlyShielded CriticalityLevel = 4
)
