package apicontracts

type Health int

const (
	HealthUnknown   Health = iota
	HealthHealthy          = 1
	HealthUnhealthy        = 2
	HealthBad              = 3
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

type ClusterPhase string

const (
	ClusterPhaseUnkown   ClusterPhase = "Unknown"
	ClusterPhaseCreating ClusterPhase = "Creating"
	ClusterPhaseReady    ClusterPhase = "Ready"
	ClusterPhaseUpdating ClusterPhase = "Updating"
	ClusterPhaseDeleting ClusterPhase = "Deleting"
	ClusterPhaseDeleted  ClusterPhase = "Deleted"
	ClusterPhaseError    ClusterPhase = "Error"
)

type ProjectRoleDefinition = string

const (
	ProjectRoleUnknown     ProjectRoleDefinition = ""
	ProjectRoleOwner       ProjectRoleDefinition = "Owner"
	ProjectRoleResponsible ProjectRoleDefinition = "Responsible"
)

type TaskSpecType = string

const (
	TaskSpecTypeUnknown   TaskSpecType = ""
	TaskSpecTypeSecret    TaskSpecType = "Secret"
	TaskSpecTypeConfigMap TaskSpecType = "ConfigMap"
)

type MatchModeType string

const (
	MatchModeUnknown  MatchModeType = "unknown"
	MatchModeContains MatchModeType = "contains"
	MatchModeEquals   MatchModeType = "equals"
	MatchModeIn       MatchModeType = "in"
)

type ConditionStatus string

const (
	ConditionStatusUnknown ConditionStatus = "Unknown"
	ConditionStatusTrue    ConditionStatus = "True"
	ConditionStatusFalse   ConditionStatus = "False"
)

type ConditionType string

const (
	ConditionTypeUnknown      ConditionType = "Unknown"
	ConditionTypeClusterReady ConditionType = "ClusterReady"
	ConditionTypeToolingOk    ConditionType = "ToolingOk"
)

type ClusterState string

const (
	ClusterStateUnknown ClusterState = "Unknown"
	ClusterStateReady   ClusterState = "Ready"
	ClusterStateWarning ClusterState = "Warning"
	ClusterStateError   ClusterState = "Error"
)

type EnvironmentType int

const (
	EnvironmentUnknown     EnvironmentType = iota
	EnvironmentDevelopment EnvironmentType = 1
	EnvironmentTesting     EnvironmentType = 2
	EnvironmentQA          EnvironmentType = 3
	EnvironmentProduction  EnvironmentType = 4
)
