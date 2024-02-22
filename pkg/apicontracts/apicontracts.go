package apicontracts

import (
	"time"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
	"github.com/NorskHelsenett/ror/pkg/models/providers"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Input models, for Creating and updating with validation

type DatacenterModel struct {
	ID          string                  `json:"id" bson:"_id,omitempty"`
	Name        string                  `json:"name" validate:"required,min=1,ne=' '" `
	Provider    string                  `json:"provider" validate:"required,min=1,ne=' '" `
	Location    DatacenterLocationModel `json:"location" validate:"required"`
	APIEndpoint string                  `json:"apiEndpoint"`
}

type DatacenterLocationModel struct {
	ID      string `json:"id" bson:"id,omitempty"`
	Region  string `json:"region" validate:"required,min=1,ne=' '" `
	Country string `json:"country" validate:"required,min=1,ne=' '" `
}

type ProjectModel struct {
	ID              string               `json:"id" bson:"_id,omitempty"`
	Name            string               `json:"name" validate:"required,min=1,ne=' ',rortext" `
	Description     string               `json:"description" validate:"required,min=1,ne=' '" `
	Active          bool                 `json:"active"`
	Created         time.Time            `json:"created"`
	Updated         time.Time            `json:"updated"`
	ProjectMetadata ProjectMetadataModel `json:"projectMetadata" validate:"required" `
}

type ProjectMetadataModel struct {
	Roles       []ProjectRoleModel `json:"roles" validate:"required,gt=1,dive,required"`
	Billing     BillingModel       `json:"billing" validate:"required"`
	ServiceTags map[string]string  `json:"serviceTags"`
}

type ProjectRoleModel struct {
	ContactInfo    ProjectContactInfoModel `json:"contactInfo" validate:"required"`
	RoleDefinition ProjectRoleDefinition   `json:"roleDefinition" validate:"required,oneof=Owner Responsible" `
}

type ProjectContactInfoModel struct {
	UPN   string `json:"upn" validate:"required,rortext"`
	Email string `json:"email" validate:"required,rortext"`
	Phone string `json:"phone" validate:"rortext"`
}

type BillingModel struct {
	Workorder string `json:"workorder" validate:"required,min=1,rortext"`
}

type ClusterMetadataModel struct {
	ProjectID   string            `json:"projectId" validate:"required,min=1"`
	Criticality CriticalityLevel  `json:"criticality" validate:"required,min=1,max=4"`
	Sensitivity SensitivityLevel  `json:"sensitivity" validate:"required,min=1,max=4"`
	Description string            `json:"description" validate:"omitempty,min=1,rortext"`
	ServiceTags map[string]string `json:"serviceTags" validate:"omitempty"`
	Billing     BillingModel      `json:"billing"`
	Roles       []ProjectRole     `json:"roles" validate:"required,gt=1,dive,required"`
}

type AgentApiKeyModel struct {
	Identifier     string                 `json:"identifier" validate:"required,min=1,ne=' ',max=100"`
	DatacenterName string                 `json:"datacenterName" validate:"required,min=1,ne=' '"`
	WorkspaceName  string                 `json:"workspaceName" validate:"required,min=1,ne=' '"`
	Provider       providers.ProviderType `json:"provider" validate:"required,min=1,ne=' ',max=20"`
	Type           ApiKeyType             `json:"type" validate:"required,min=1,ne='',eq=Cluster"`
}

/// Domain object

type Datacenter struct {
	ID          string                 `json:"id" bson:"_id,omitempty"`
	Name        string                 `json:"name"`
	Provider    providers.ProviderType `json:"provider"`
	Location    DatacenterLocation     `json:"location"`
	APIEndpoint string                 `json:"apiEndpoint"`
}

type DatacenterLocation struct {
	ID      string `json:"id" bson:"id,omitempty"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

type Workspace struct {
	ID           string     `json:"id" bson:"_id,omitempty"`
	Name         string     `json:"name"`
	DatacenterID string     `json:"datacenterId" bson:"datacenterid,omitempty"`
	Datacenter   Datacenter `json:"datacenter"`
}

type Cluster struct {
	ID            string            `json:"id" bson:"_id,omitempty"`
	Identifier    string            `json:"identifier"`
	ClusterIdOld  string            `json:"clusterIdOld" bson:"clusteridold"`
	ACL           AccessControlList `json:"acl"`
	ClusterId     string            `json:"clusterId"`
	ClusterName   string            `json:"clusterName"`
	WorkspaceId   string            `json:"workspaceId" bson:"workspaceid"`
	Workspace     Workspace         `json:"workspace" bson:"workspace"`
	Environment   string            `json:"environment"`
	Metrics       Metrics           `json:"metrics"`
	Topology      Topology          `json:"topology"`
	Versions      Versions          `json:"versions"`
	Ingresses     []Ingress         `json:"ingresses"`
	Updated       time.Time         `json:"updated,omitempty"`
	Created       time.Time         `json:"created,omitempty"`
	LastObserved  time.Time         `json:"lastObserved,omitempty"`
	FirstObserved time.Time         `json:"firstObserved,omitempty"`
	HealthStatus  HealthStatus      `json:"healthStatus"`
	CreatedBy     string            `json:"createdBy"`
	SplunkIndex   string            `json:"splunkIndex"`
	Config        ClusterConfig     `json:"config"`
	Metadata      ClusterMetadata   `json:"metadata"`
	Status        ClusterStatus     `json:"status"`
}

type ClusterInfo struct {
	Id          string          `json:"id" bson:"_id"`
	ClusterId   string          `json:"clusterId"`
	ClusterName string          `json:"clusterName"`
	Metadata    ClusterMetadata `json:"metadata"`
}

type ClusterStatus struct {
	State      ClusterState       `json:"state"`
	Phase      ClusterPhase       `json:"phase"`
	Conditions []ClusterCondition `json:"conditions"`
}

type ClusterSelf struct {
	ClusterId string `json:"clusterId"`
}

type ClusterCondition struct {
	Type    ConditionType   `json:"type"`
	Status  ConditionStatus `json:"status"`
	Message string          `json:"message"`
	Created time.Time       `json:"created"`
	Updated time.Time       `json:"updated"`
}

type ClusterMetadata struct {
	ProjectID   string            `json:"projectId,omitempty" bson:"projectid,omitempty"`
	Project     *Project          `json:"project,omitempty"`
	Criticality CriticalityLevel  `json:"criticality"`
	Sensitivity SensitivityLevel  `json:"sensitivity"`
	Description string            `json:"description"`
	ServiceTags map[string]string `json:"serviceTags"`
	Billing     Billing           `json:"billing"`
	Roles       []ProjectRole     `json:"roles"`
}

type ClusterConfig struct {
	Versions        map[string]interface{} `json:"versions"`
	Overrides       map[string]interface{} `json:"overrides"`
	ProjectMetadata ProjectMetadata        `json:"projectMetadata"`
}

type Versions struct {
	Kubernetes string     `json:"kubernetes"`
	NhnTooling NhnTooling `json:"nhnTooling"`
	Agent      Agent      `json:"agent"`
}

type Agent struct {
	Version string `json:"version"`
	Sha     string `json:"sha"`
}

type Nodes struct {
	ControlPlane ControlPlane `json:"controlplane"`
}

type ArgoCdVersions struct {
	Version     string `json:"version"`
	HelmVersion string `json:"helmVersion"`
}

type NhnTooling struct {
	Version     string `json:"version"`
	Branch      string `json:"branch"`
	Environment string `json:"environment"`
}

type Topology struct {
	ControlPlaneEndpoint string       `json:"controlPlaneEndpoint"`
	EgressIp             string       `json:"egressIp"`
	ControlPlane         ControlPlane `json:"controlPlane"`
	NodePools            []NodePool   `validate:"min=1" json:"nodePools"`
}

type AccessControlList struct {
	AccessGroups []string `json:"accessGroups"`
}

type ControlPlane struct {
	Nodes   []Node  `json:"nodes"`
	Metrics Metrics `json:"metrics"`
}

type Node struct {
	Name                    string    `json:"name"`
	Role                    string    `json:"role"`
	Created                 time.Time `json:"created"`
	OsImage                 string    `json:"osImage"`
	MachineName             string    `json:"machineName"`
	Metrics                 Metrics   `json:"metrics"`
	Architecture            string    `json:"architecture"`
	ContainerRuntimeVersion string    `json:"containerRuntimeVersion"`
	KernelVersion           string    `json:"kernelVersion"`
	KubeProxyVersion        string    `json:"kubeProxyVersion"`
	KubeletVersion          string    `json:"kubeletVersion"`
	OperatingSystem         string    `json:"operatingSystem"`
	MachineClass            string    `json:"machineClass"`
}

type NodeResources struct {
	Allocated ResourceAllocated `json:"allocated"`
	Consumed  ResourceConsumed  `json:"consumed"`
}

type ResourceAllocated struct {
	Cpu         int64 `json:"cpu"`
	MemoryBytes int64 `json:"memoryBytes"`
}

type ResourceConsumed struct {
	CpuMilliValue int64 `json:"cpuMilliValue"`
	MemoryBytes   int64 `json:"memoryBytes"`
}

type NodePool struct {
	Name         string  `json:"name"`
	MachineClass string  `json:"machineClass"`
	Metrics      Metrics `json:"metrics"`
	Nodes        []Node  `json:"nodes"`
}

type Ingress struct {
	UID       string        `json:"uid"`
	Health    Health        `json:"health"`
	Name      string        `json:"name"`
	Namespace string        `json:"namespace"`
	Class     string        `json:"class"`
	Rules     []IngressRule `json:"ingressrules" bson:"ingressrules"`
}

type IngressRule struct {
	Hostname    string        `json:"hostname"`
	IPAddresses []string      `json:"ipaddresses"`
	Paths       []IngressPath `json:"rules" bson:"rules"`
}

type IngressPath struct {
	Path    string  `json:"path"`
	Service Service `json:"service"`
}

type Service struct {
	Name      string            `json:"name"`
	Type      string            `json:"type"`
	Selector  string            `json:"selector" bson:"selector"`
	Ports     []ServicePort     `json:"ports"`
	Endpoints []EndpointAddress `json:"endpoints"`
}

type ServicePort struct {
	Name     string `json:"name"`
	NodePort string `json:"nodeport"`
	Protocol string `json:"protocol"`
}

type EndpointAddress struct {
	NodeName string `json:"nodename"`
	PodName  string `json:"podnamespace" bson:"podnamespace"`
}

type HealthStatus struct {
	Health   Health          `json:"health"`
	Messages []StatusMessage `json:"messages"`
}

type StatusMessage struct {
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

type ClusterControlPlaneMetadata struct {
	ClusterId                interface{}                           `json:"clusterId"`
	ClusterName              interface{}                           `json:"clusterName"`
	Environment              interface{}                           `json:"environment"`
	ProjectName              interface{}                           `json:"projectName"`
	ControlPlaneEndpoint     ClusterControlPlaneMetadataIp         `json:"controlPlaneEndpoint"`
	ControlPlaneEndpointPort interface{}                           `json:"controlPlaneEndpointPort"`
	Egress                   ClusterControlPlaneMetadataIp         `json:"egress"`
	Datacenter               ClusterControlPlaneMetadataDatacenter `json:"datacenter"`
}

type ClusterControlPlaneMetadataDatacenter struct {
	Name        interface{} `json:"name"`
	Provider    interface{} `json:"provider"`
	ApiEndpoint interface{} `json:"apiEndpoint"`
}

type ClusterControlPlaneMetadataIp struct {
	IpV4 interface{} `json:"ipv4"`
	IpV6 interface{} `json:"ipv6"`
}

type Metrics struct {
	PriceMonth       int64 `json:"priceMonth"`
	PriceYear        int64 `json:"priceYear"`
	Cpu              int64 `json:"cpu"`
	Memory           int64 `json:"memory"`
	CpuConsumed      int64 `json:"cpuConsumed"`
	MemoryConsumed   int64 `json:"memoryConsumed"`
	CpuPercentage    int64 `json:"cpuPercentage"`
	MemoryPercentage int64 `json:"memoryPercentage"`
	NodePoolCount    int64 `json:"nodePoolCount"`
	NodeCount        int64 `json:"nodeCount"`
	ClusterCount     int64 `json:"clusterCount"`
}

type MetricsTotal struct {
	Cpu              int64 `json:"cpu"`
	Memory           int64 `json:"memory"`
	CpuConsumed      int64 `json:"cpuConsumed"`
	MemoryConsumed   int64 `json:"memoryConsumed"`
	CpuPercentage    int64 `json:"cpuPercentage"`
	MemoryPercentage int64 `json:"memoryPercentage"`
	NodePoolCount    int64 `json:"nodePoolCount"`
	NodeCount        int64 `json:"nodeCount"`
	ClusterCount     int64 `json:"clusterCount"`
	WorkspaceCount   int64 `json:"workspaceCount"`
	DatacenterCount  int64 `json:"datacenterCount"`
}

type Project struct {
	ID              string          `json:"id" bson:"_id,omitempty"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Active          bool            `json:"active"`
	Created         time.Time       `json:"created"`
	Updated         time.Time       `json:"updated"`
	ProjectMetadata ProjectMetadata `json:"projectMetadata"`
}

type ProjectMetadata struct {
	Roles       []ProjectRole     `json:"roles"`
	Billing     Billing           `json:"billing"`
	ServiceTags map[string]string `json:"serviceTags"`
}

type ProjectRole struct {
	ContactInfo    ProjectContactInfo    `json:"contactInfo"`
	RoleDefinition ProjectRoleDefinition `json:"roleDefinition"`
}

type ProjectContactInfo struct {
	UPN   string `json:"upn"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Billing struct {
	Workorder string `json:"workorder"`
}

type Price struct {
	ID           string                 `json:"id" bson:"_id,omitempty"`
	Provider     providers.ProviderType `json:"provider"`
	MachineClass string                 `json:"machineClass"`
	Cpu          int                    `json:"cpu"`
	Memory       int64                  `json:"memory"`
	MemoryBytes  int64                  `json:"memoryBytes"`
	Price        int                    `json:"price"`
	From         time.Time              `json:"from"`
	To           time.Time              `json:"to,omitempty"`
}

type PaginatedResult[T any] struct {
	Data       []T   `json:"data"`
	DataCount  int64 `json:"dataCount"`
	TotalCount int64 `json:"totalCount"`
	Offset     int64 `json:"offset"`
}

type MetricList struct {
	Items []MetricItem `json:"items"`
}

type MetricItem struct {
	Id      string  `json:"id"`
	Metrics Metrics `json:"metrics"`
}

type MetricsCustom struct {
	Data []MetricsCustomItem `json:"data"`
}

type MetricsCustomItem struct {
	Value int64  `json:"value"`
	Text  string `json:"text"`
}

type Metric struct {
	Id               string `json:"id"`
	PriceMonth       int64  `json:"priceMonth"`
	PriceYear        int64  `json:"priceYear"`
	Cpu              int64  `json:"cpu"`
	Memory           int64  `json:"memory"`
	CpuConsumed      int64  `json:"cpuConsumed"`
	MemoryConsumed   int64  `json:"memoryConsumed"`
	CpuPercentage    int64  `json:"cpuPercentage"`
	MemoryPercentage int64  `json:"memoryPercentage"`
	NodePoolCount    int64  `json:"nodePoolCount"`
	NodeCount        int64  `json:"nodeCount"`
	ClusterCount     int64  `json:"clusterCount"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type DesiredVersion struct {
	Key   string      `json:"key" bson:"key"`
	Value interface{} `json:"value" bson:"value"`
}

type User struct {
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Groups []string `json:"groups"`
}

type AuditLogMetadata struct {
	Msg       string    `json:"msg"`
	Timestamp time.Time `json:"timestamp"`
	Category  string    `json:"category"`
	Action    string    `json:"action"`
	User      User      `json:"user"`
}

type AuditLog struct {
	ID       string           `json:"id" bson:"_id,omitempty"`
	Metadata AuditLogMetadata `json:"metadata"`
	Data     map[string]any   `json:"data"`
}

type FilterMetadata struct {
	Field     string        `json:"field"`
	Value     any           `json:"value"`
	MatchMode MatchModeType `json:"matchMode"`
}

type SortMetadata struct {
	SortField string `json:"sortField"`
	SortOrder int    `json:"sortOrder"`
}

type Filter struct {
	Skip         int              `json:"skip"`
	Limit        int              `json:"limit"`
	Sort         []SortMetadata   `json:"sort" validate:"dive"`
	Filters      []FilterMetadata `json:"filters"`
	GlobalFilter any              `json:"globalFilter"`
}

type Task struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" validate:"required"`
	Config TaskSpec           `json:"config" validate:"required"`
}

type TaskCollection struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Tasks []Task `json:"tasks"`
}

type ApiKey struct {
	Id          string     `json:"id,omitempty" bson:"_id,omitempty"`
	Identifier  string     `json:"identifier" validate:"required,min=1,ne='',max=100,rortext"`
	DisplayName string     `json:"displayName,omitempty" validate:"required,min=3,ne='',max=20,rortext"`
	Type        ApiKeyType `json:"type" validate:"required,min=1,ne=''"`
	ReadOnly    bool       `json:"readOnly"`
	Expires     time.Time  `json:"expires,omitempty"`
	Created     time.Time  `json:"created,omitempty"`
	LastUsed    time.Time  `json:"lastUsed,omitempty"`
	Hash        string
}

type SSEMessage struct {
	Event string `json:"event" validate:"required,min=1,ne=' '"`
	Data  any    `json:"data" validate:"required"`
}

type SSEClient struct {
	Identity   identitymodels.Identity `json:"identity"`
	Connection chan string             `json:"connection"`
}

type TanzuKubeConfigPayload struct {
	User          string `json:"user"`
	Password      string `json:"pwd"`
	DatacenterUrl string `json:"datacenterUrl"`
	ClusterName   string `json:"clusterName,omitempty"`
	ClusterId     string `json:"clusterId,omitempty"`
	WorkspaceName string `json:"workspaceName"`
	WorkspaceOnly bool   `json:"workspaceOnly"`
}

type KubeconfigCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ClusterKubeconfig struct {
	Status   string `json:"status,omitempty"`
	Message  string `json:"message,omitempty"`
	Data     string `json:"data,omitempty"`
	DataType string `json:"dataType,omitempty"`
}

type ApiKeyType string

const (
	ApiKeyTypeUnknown ApiKeyType = ""
	ApiKeyTypeCluster ApiKeyType = "Cluster"
	ApiKeyTypeUser    ApiKeyType = "User"
	ApiKeyTypeService ApiKeyType = "Service"
)

func (key ApiKey) IsExpired() bool {
	if key.Expires.IsZero() {
		return false
	}
	if key.Expires.Local().After(time.Now()) {
		return false
	}
	return true
}
