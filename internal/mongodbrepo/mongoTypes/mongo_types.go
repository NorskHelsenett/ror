package mongoTypes

import (
	"github.com/NorskHelsenett/ror/internal/models"
	"time"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"
	"github.com/NorskHelsenett/ror/pkg/models/providers"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoDatacenter struct {
	ID          primitive.ObjectID      `json:"id" bson:"_id,omitempty"`
	Name        string                  `json:"name"`
	Provider    providers.ProviderType  `json:"provider"`
	Location    MongoDatacenterLocation `json:"location"`
	APIEndpoint string                  `json:"apiEndpoint"`
}

type MongoDatacenterLocation struct {
	ID      string `json:"id" bson:"id,omitempty"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

type MongoWorkspace struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name"`
	DatacenterID primitive.ObjectID `json:"datacenterId" bson:"datacenterid,omitempty"`
}

type MongoCluster struct {
	ID            primitive.ObjectID     `json:"id" bson:"_id,omitempty"`
	Identifier    string                 `json:"identifier"`
	ACL           MongoAccessControlList `json:"acl"`
	ClusterId     string                 `json:"clusterId"`
	ClusterName   string                 `json:"clusterName"`
	WorkspaceId   primitive.ObjectID     `json:"workspaceId" bson:"workspaceid"`
	Environment   string                 `json:"environment"`
	Metrics       MongoMetrics           `json:"metrics"`
	Topology      MongoTopology          `json:"topology"`
	Versions      MongoVersions          `json:"versions"`
	Ingresses     []MongoIngress         `json:"ingresses"`
	Updated       time.Time              `json:"updated,omitempty"`
	Created       time.Time              `json:"created,omitempty"`
	LastObserved  time.Time              `json:"lastObserved,omitempty"`
	FirstObserved time.Time              `json:"firstObserved,omitempty"`
	HealthStatus  MongoHealthStatus      `json:"healthStatus"`
	CreatedBy     string                 `json:"createdBy"`
	SplunkIndex   string                 `json:"splunkIndex"`
	Config        MongoClusterConfig     `json:"config"`
	Metadata      ClusterMetadata        `json:"metadata"`
	Status        MongoClusterStatus     `json:"status"`
}

type MongoClusterStatus struct {
	State      apicontracts.ClusterState `json:"state"`
	Phase      apicontracts.ClusterPhase `json:"phase"`
	Conditions []MongoClusterCondition   `json:"conditions"`
}

type MongoClusterCondition struct {
	Type    apicontracts.ConditionType   `json:"type"`
	Status  apicontracts.ConditionStatus `json:"status"`
	Message string                       `json:"message"`
	Created time.Time                    `json:"created"`
	Updated time.Time                    `json:"updated"`
}

type ClusterMetadata struct {
	ProjectID   primitive.ObjectID            `json:"projectId,omitempty" bson:"projectid,omitempty"`
	Criticality apicontracts.CriticalityLevel `json:"criticality"`
	Sensitivity apicontracts.SensitivityLevel `json:"sensitivity"`
	Description string                        `json:"description"`
	ServiceTags map[string]string             `json:"serviceTags"`
	Billing     MongoBilling                  `json:"billing"`
	Roles       []MongoProjectRole            `json:"roles"`
}

type MongoClusterConfig struct {
	Versions  map[string]interface{} `json:"versions"`
	Overrides map[string]interface{} `json:"overrides"`
}

type MongoClusterWithWorkspace struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Workspace MongoWorkspace     `json:"workspace"`
}
type MongoClusterWithDatacenter struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Datacenter MongoDatacenter    `json:"datacenter"`
}

type MongoVersions struct {
	Kubernetes string          `json:"kubernetes"`
	NhnTooling MongoNhnTooling `json:"nhnTooling"`
	Agent      MongoAgent      `json:"agent"`
}

type MongoAgent struct {
	Version string `json:"version"`
	Sha     string `json:"sha"`
}

type MongoNodes struct {
	ControlPlane MongoControlPlane `json:"controlplane"`
}

type MongoArgoCdVersions struct {
	Version     string `json:"version"`
	HelmVersion string `json:"helmVersion"`
}

type MongoNhnTooling struct {
	Version     string `json:"version"`
	Branch      string `json:"branch"`
	Environment string `json:"environment"`
}

type MongoTopology struct {
	ControlPlaneEndpoint string            `json:"controlPlaneEndpoint"`
	EgressIp             string            `json:"egressIp"`
	ControlPlane         MongoControlPlane `json:"controlPlane"`
	NodePools            []MongoNodePool   `validate:"min=1" json:"nodePools"`
}

type MongoAccessControlList struct {
	AccessGroups []string `json:"accessGroups"`
}

type MongoControlPlane struct {
	Nodes   []MongoNode  `json:"nodes" bson:"nodes"`
	Metrics MongoMetrics `json:"metrics"`
}

type MongoNode struct {
	Name                    string       `json:"name"`
	Role                    string       `json:"role"`
	Created                 time.Time    `json:"created"`
	OsImage                 string       `json:"osImage"`
	MachineName             string       `json:"machineName"`
	MachineClass            string       `json:"machineClass"`
	Metrics                 MongoMetrics `json:"metrics"`
	Architecture            string       `json:"architecture"`
	ContainerRuntimeVersion string       `json:"containerRuntimeVersion"`
	KernelVersion           string       `json:"kernelVersion"`
	KubeProxyVersion        string       `json:"kubeProxyVersion"`
	KubeletVersion          string       `json:"kubeletVersion"`
	OperatingSystem         string       `json:"operatingSystem"`
}

type MongoNodeResources struct {
	Allocated MongoResourceAllocated `json:"allocated"`
	Consumed  MongoResourceConsumed  `json:"consumed"`
}

type MongoResourceAllocated struct {
	Cpu         int64 `json:"cpu"`
	MemoryBytes int64 `json:"memoryBytes"`
}

type MongoResourceConsumed struct {
	CpuMilliValue int64 `json:"cpuMilliValue"`
	MemoryBytes   int64 `json:"memoryBytes"`
}

type MongoNodePool struct {
	Name         string       `json:"name"`
	MachineClass string       `json:"machineClass"`
	Metrics      MongoMetrics `json:"metrics"`
	Nodes        []MongoNode  `json:"nodes"`
}

type MongoIngress struct {
	UID       string              `json:"uid"`
	Health    apicontracts.Health `json:"health"`
	Name      string              `json:"name"`
	Namespace string              `json:"namespace"`
	Class     string              `json:"class"`
	Rules     []MongoIngressRule  `json:"ingressrules" bson:"ingressrules"`
}

type MongoIngressRule struct {
	Hostname    string             `json:"hostname"`
	IPAddresses []string           `json:"ipaddresses"`
	Paths       []MongoIngressPath `json:"rules" bson:"rules"`
}

type MongoIngressPath struct {
	Path    string       `json:"path"`
	Service MongoService `json:"service"`
}

type MongoService struct {
	Name      string                 `json:"name"`
	Type      string                 `json:"type"`
	Selector  string                 `json:"selector" bson:"selector"`
	Ports     []MongoServicePort     `json:"ports" bson:"ports"`
	Endpoints []MongoEndpointAddress `json:"endpoints" bson:"endpoints"`
}

type MongoServicePort struct {
	Name     string `json:"name"`
	NodePort string `json:"nodeport"`
	Protocol string `json:"protocol"`
}

type MongoEndpointAddress struct {
	NodeName string `json:"nodename"`
	PodName  string `json:"podnamespace" bson:"podnamespace"`
}

type MongoHealthStatus struct {
	Health   apicontracts.Health  `json:"health"`
	Messages []MongoStatusMessage `json:"messages" bson:"messages"`
}

type MongoStatusMessage struct {
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

type MongoMetrics struct {
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

type MongoPrice struct {
	ID           primitive.ObjectID     `json:"id" bson:"_id,omitempty"`
	Provider     providers.ProviderType `json:"provider"`
	MachineClass string                 `json:"machineClass"`
	Cpu          int                    `json:"cpu"`
	Memory       int64                  `json:"memory"`
	MemoryBytes  int64                  `json:"memoryBytes"`
	Price        int                    `json:"price"`
	From         time.Time              `json:"from"`
	To           time.Time              `json:"to,omitempty"`
}

type MongoAuditLogMetadata struct {
	Msg       string               `json:"msg"`
	Timestamp time.Time            `json:"timestamp"`
	Category  models.AuditCategory `json:"category"`
	Action    models.AuditAction   `json:"action"`
	User      identitymodels.User  `json:"user"`
}
type MongoAuditLog struct {
	ID       string                `json:"id" bson:"_id,omitempty"`
	Metadata MongoAuditLogMetadata `json:"metadata"`
	Data     map[string]any        `json:"data"`
}

type MongoTaskCollection struct {
	ID    primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Tasks []MongoOperatorTask `json:"tasks"`
}

type MongoOperatorConfig struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	ApiVersion string             `json:"apiVersion"`
	Kind       string             `json:"kind"`
	Spec       *MongoOperatorSpec `json:"spec"`
}

type MongoOperatorSpec struct {
	ImagePostfix string              `json:"imagePostfix"`
	Tasks        []MongoOperatorTask `json:"tasks"`
}

type MongoOperatorTask struct {
	Index   uint   `json:"index"`
	Name    string `json:"name"`
	Version string `json:"version"`
	RunOnce bool   `json:"runOnce"`
}

type MongoProject struct {
	ID              primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	Active          bool                 `json:"active"`
	Created         time.Time            `json:"created"`
	Updated         time.Time            `json:"updated"`
	ProjectMetadata MongoProjectMetadata `json:"projectmetadata"`
}

type MongoProjectMetadata struct {
	Roles       []MongoProjectRole `json:"roles"`
	Billing     MongoBilling       `json:"billing"`
	ServiceTags map[string]string  `json:"serviceTags"`
}

type MongoProjectRole struct {
	ContactInfo    MongoProjectContactInfo            `json:"contactInfo"`
	RoleDefinition apicontracts.ProjectRoleDefinition `json:"roleDefinition"`
}

type MongoProjectContactInfo struct {
	UPN   string `json:"upn"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type MongoBilling struct {
	Workorder string `json:"workorder"`
}
