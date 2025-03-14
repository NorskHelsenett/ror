package rortypes

type ResourceClusterOrder struct {
	Spec   ResourceClusterOrderSpec   `json:"spec"`
	Status ResourceClusterOrderStatus `json:"status"`
}

type ResourceClusterOrderSpec struct {
	Provider    ProviderType `json:"provider" validate:"required,min=1,ne=' '"`
	ClusterName string       `json:"clusterName" validate:"required,min=1,ne=' '"`
	ProjectId   string       `json:"projectId" validate:"required,min=1,ne=' '"`
	OrderBy     string       `json:"orderBy" validate:"required,min=1,ne=' '"`

	Environment      EnvironmentType                    `json:"environment" validate:"required,min=1,max=4"`
	Criticality      CriticalityLevel                   `json:"criticality" validate:"required,min=1,max=4"`
	Sensitivity      SensitivityLevel                   `json:"sensitivity" validate:"required,min=1,max=4"`
	HighAvailability bool                               `json:"highAvailability" validate:"boolean"`
	NodePools        []ResourceClusterOrderSpecNodePool `json:"nodePools" validate:"required,min=1,dive,required"`
	ServiceTags      map[string]string                  `json:"serviceTags,omitempty"`
	ProviderConfig   map[string]interface{}             `json:"providerConfig,omitempty"`
	OwnerGroup       string                             `json:"ownerGroup" validate:"required,min=1,ne=' '"`
}

type ResourceProviderConfigTanzu struct {
	DatacenterId string `json:"datacenterId" validate:"required,min=1,ne=' '"`
	NamespaceId  string `json:"namespaceId" validate:"required,min=1,ne=' '"`
}
type ResourceProviderConfigAks struct {
	Region        string `json:"region" validate:"required,min=1,ne=' '"`
	Subscription  string `json:"subscription" validate:"required,min=1,ne=' '"`
	ResourceGroup string `json:"resourceGroup" validate:"required,min=1,ne=' '"`
}

type ResourceClusterOrderSpecNodePool struct {
	Name         string `json:"name" validate:"required,min=1,ne=' '"`
	MachineClass string `json:"machineClass" validate:"required,min=1,ne=' '"`
	Count        int    `json:"count" validate:"required,min=1"`
}

type ResourceClusterOrderStatus struct {
	Status           string                                          `json:"status"`
	Phase            string                                          `json:"phase"`
	Conditions       []ResourceKubernetesClusterOrderStatusCondition `json:"conditions"`
	CreatedTime      string                                          `json:"createdTime"`
	UpdatedTime      string                                          `json:"updatedTime"`
	LastObservedTime string                                          `json:"lastObservedTime"`
}

type ResourceClusterOrderCondition struct {
	Type               string `json:"type"`
	Status             string `json:"status"`
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Message            string `json:"message"`
}

type EnvironmentType int

const (
	EnvironmentUnknown     EnvironmentType = iota
	EnvironmentDevelopment EnvironmentType = 1
	EnvironmentTesting     EnvironmentType = 2
	EnvironmentQA          EnvironmentType = 3
	EnvironmentProduction  EnvironmentType = 4
)

type SensitivityLevel int

const (
	SensitivityLevelUnknown  SensitivityLevel = iota
	SensitivityLevelNormal   SensitivityLevel = 1
	SensitivityLevelModerate SensitivityLevel = 2
	SensitivityLevelHigh     SensitivityLevel = 3
	SensitivityLevelCritical SensitivityLevel = 4
)

type CriticalityLevel int

const (
	CriticalityLevelUnknown        CriticalityLevel = iota
	CriticalityLevelOpen           CriticalityLevel = 1
	CriticalityLevelIntern         CriticalityLevel = 2
	CriticalityLevelShielded       CriticalityLevel = 3
	CriticalityLevelHighlyShielded CriticalityLevel = 4
)

type ProviderType string

const (
	ProviderTypeUnknown ProviderType = ""
	ProviderTypeTanzu   ProviderType = "Tanzu"
	ProviderTypeAzure   ProviderType = "Azure"
	ProviderTypeK3d     ProviderType = "K3D"
)

type ResourceKubernetesClusterOrderStatusCondition struct {
	Type               string `json:"type"`
	Status             string `json:"status"`
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Message            string `json:"message"`
}
