// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package rorresources

import (
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
)

// The Resource struct represents one resource in ror.
//
// It implement common and resource specific methods by providing interfaces to the underlying resources
type Resource struct {
	rortypes.CommonResource `json:",inline" bson:",inline"`

	NamespaceResource                       *rortypes.ResourceNamespace                       `json:"namespace,omitempty" bson:"namespace,omitempty"`
	NodeResource                            *rortypes.ResourceNode                            `json:"node,omitempty" bson:"node,omitempty"`
	PersistentVolumeClaimResource           *rortypes.ResourcePersistentVolumeClaim           `json:"persistentvolumeclaim,omitempty" bson:"persistentvolumeclaim,omitempty"`
	DeploymentResource                      *rortypes.ResourceDeployment                      `json:"deployment,omitempty" bson:"deployment,omitempty"`
	StorageClassResource                    *rortypes.ResourceStorageClass                    `json:"storageclass,omitempty" bson:"storageclass,omitempty"`
	PolicyReportResource                    *rortypes.ResourcePolicyReport                    `json:"policyreport,omitempty" bson:"policyreport,omitempty"`
	ApplicationResource                     *rortypes.ResourceApplication                     `json:"application,omitempty" bson:"application,omitempty"`
	AppProjectResource                      *rortypes.ResourceAppProject                      `json:"appproject,omitempty" bson:"appproject,omitempty"`
	CertificateResource                     *rortypes.ResourceCertificate                     `json:"certificate,omitempty" bson:"certificate,omitempty"`
	ServiceResource                         *rortypes.ResourceService                         `json:"service,omitempty" bson:"service,omitempty"`
	PodResource                             *rortypes.ResourcePod                             `json:"pod,omitempty" bson:"pod,omitempty"`
	ReplicaSetResource                      *rortypes.ResourceReplicaSet                      `json:"replicaset,omitempty" bson:"replicaset,omitempty"`
	StatefulSetResource                     *rortypes.ResourceStatefulSet                     `json:"statefulset,omitempty" bson:"statefulset,omitempty"`
	DaemonSetResource                       *rortypes.ResourceDaemonSet                       `json:"daemonset,omitempty" bson:"daemonset,omitempty"`
	IngressResource                         *rortypes.ResourceIngress                         `json:"ingress,omitempty" bson:"ingress,omitempty"`
	IngressClassResource                    *rortypes.ResourceIngressClass                    `json:"ingressclass,omitempty" bson:"ingressclass,omitempty"`
	SbomReportResource                      *rortypes.ResourceSbomReport                      `json:"sbomreport,omitempty" bson:"sbomreport,omitempty"`
	VulnerabilityReportResource             *rortypes.ResourceVulnerabilityReport             `json:"vulnerabilityreport,omitempty" bson:"vulnerabilityreport,omitempty"`
	ExposedSecretReportResource             *rortypes.ResourceExposedSecretReport             `json:"exposedsecretreport,omitempty" bson:"exposedsecretreport,omitempty"`
	ConfigAuditReportResource               *rortypes.ResourceConfigAuditReport               `json:"configauditreport,omitempty" bson:"configauditreport,omitempty"`
	RbacAssessmentReportResource            *rortypes.ResourceRbacAssessmentReport            `json:"rbacassessmentreport,omitempty" bson:"rbacassessmentreport,omitempty"`
	TanzuKubernetesClusterResource          *rortypes.ResourceTanzuKubernetesCluster          `json:"tanzukubernetescluster,omitempty" bson:"tanzukubernetescluster,omitempty"`
	TanzuKubernetesReleaseResource          *rortypes.ResourceTanzuKubernetesRelease          `json:"tanzukubernetesrelease,omitempty" bson:"tanzukubernetesrelease,omitempty"`
	VirtualMachineClassResource             *rortypes.ResourceVirtualMachineClass             `json:"virtualmachineclass,omitempty" bson:"virtualmachineclass,omitempty"`
	KubernetesClusterResource               *rortypes.ResourceKubernetesCluster               `json:"kubernetescluster,omitempty" bson:"kubernetescluster,omitempty"`
	ProviderResource                        *rortypes.ResourceProvider                        `json:"provider,omitempty" bson:"provider,omitempty"`
	WorkspaceResource                       *rortypes.ResourceWorkspace                       `json:"workspace,omitempty" bson:"workspace,omitempty"`
	KubernetesMachineClassResource          *rortypes.ResourceKubernetesMachineClass          `json:"kubernetesmachineclass,omitempty" bson:"kubernetesmachineclass,omitempty"`
	ClusterOrderResource                    *rortypes.ResourceClusterOrder                    `json:"clusterorder,omitempty" bson:"clusterorder,omitempty"`
	ProjectResource                         *rortypes.ResourceProject                         `json:"project,omitempty" bson:"project,omitempty"`
	ConfigurationResource                   *rortypes.ResourceConfiguration                   `json:"configuration,omitempty" bson:"configuration,omitempty"`
	ClusterComplianceReportResource         *rortypes.ResourceClusterComplianceReport         `json:"clustercompliancereport,omitempty" bson:"clustercompliancereport,omitempty"`
	ClusterVulnerabilityReportResource      *rortypes.ResourceClusterVulnerabilityReport      `json:"clustervulnerabilityreport,omitempty" bson:"clustervulnerabilityreport,omitempty"`
	RouteResource                           *rortypes.ResourceRoute                           `json:"route,omitempty" bson:"route,omitempty"`
	SlackMessageResource                    *rortypes.ResourceSlackMessage                    `json:"slackmessage,omitempty" bson:"slackmessage,omitempty"`
	VulnerabilityEventResource              *rortypes.ResourceVulnerabilityEvent              `json:"vulnerabilityevent,omitempty" bson:"vulnerabilityevent,omitempty"`
	VirtualMachineResource                  *rortypes.ResourceVirtualMachine                  `json:"virtualmachine,omitempty" bson:"virtualmachine,omitempty"`
	VirtualMachineVulnerabilityInfoResource *rortypes.ResourceVirtualMachineVulnerabilityInfo `json:"virtualmachinevulnerabilityinfo,omitempty" bson:"virtualmachinevulnerabilityinfo,omitempty"`
	EndpointsResource                       *rortypes.ResourceEndpoints                       `json:"endpoints,omitempty" bson:"endpoints,omitempty"`
	NetworkPolicyResource                   *rortypes.ResourceNetworkPolicy                   `json:"networkpolicy,omitempty" bson:"networkpolicy,omitempty"`
	DatacenterResource                      *rortypes.ResourceDatacenter                      `json:"datacenter,omitempty" bson:"datacenter,omitempty"`
	BackupJobResource                       *rortypes.ResourceBackupJob                       `json:"backupjob,omitempty" bson:"backupjob,omitempty"`
	BackupRunResource                       *rortypes.ResourceBackupRun                       `json:"backuprun,omitempty" bson:"backuprun,omitempty"`
	MachineResource                         *rortypes.ResourceMachine                         `json:"machine,omitempty" bson:"machine,omitempty"`
	UnknownResource                         *rortypes.ResourceUnknown                         `json:"unknown,omitempty" bson:"unknown,omitempty"`

	common rortypes.CommonResourceInterface
}

// NewRorResource provides a empty resource of a given kind/apiversion
func NewRorResource(kind string, apiversion string) *Resource {
	r := Resource{}
	r.Kind = kind
	r.APIVersion = apiversion
	return &r
}

// NewRorNamespaceResource provides a empty resource of a given kind/apiversion
func NewRorNamespaceResource() *Resource {
	r := Resource{}
	r.Kind = "Namespace"
	r.APIVersion = "v1"
	r.NamespaceResource = &rortypes.ResourceNamespace{}
	r.common = r.NamespaceResource
	return &r
}

// NewRorNodeResource provides a empty resource of a given kind/apiversion
func NewRorNodeResource() *Resource {
	r := Resource{}
	r.Kind = "Node"
	r.APIVersion = "v1"
	r.NodeResource = &rortypes.ResourceNode{}
	r.common = r.NodeResource
	return &r
}

// NewRorPersistentVolumeClaimResource provides a empty resource of a given kind/apiversion
func NewRorPersistentVolumeClaimResource() *Resource {
	r := Resource{}
	r.Kind = "PersistentVolumeClaim"
	r.APIVersion = "v1"
	r.PersistentVolumeClaimResource = &rortypes.ResourcePersistentVolumeClaim{}
	r.common = r.PersistentVolumeClaimResource
	return &r
}

// NewRorDeploymentResource provides a empty resource of a given kind/apiversion
func NewRorDeploymentResource() *Resource {
	r := Resource{}
	r.Kind = "Deployment"
	r.APIVersion = "apps/v1"
	r.DeploymentResource = &rortypes.ResourceDeployment{}
	r.common = r.DeploymentResource
	return &r
}

// NewRorStorageClassResource provides a empty resource of a given kind/apiversion
func NewRorStorageClassResource() *Resource {
	r := Resource{}
	r.Kind = "StorageClass"
	r.APIVersion = "storage.k8s.io/v1"
	r.StorageClassResource = &rortypes.ResourceStorageClass{}
	r.common = r.StorageClassResource
	return &r
}

// NewRorPolicyReportResource provides a empty resource of a given kind/apiversion
func NewRorPolicyReportResource() *Resource {
	r := Resource{}
	r.Kind = "PolicyReport"
	r.APIVersion = "wgpolicyk8s.io/v1alpha2"
	r.PolicyReportResource = &rortypes.ResourcePolicyReport{}
	r.common = r.PolicyReportResource
	return &r
}

// NewRorApplicationResource provides a empty resource of a given kind/apiversion
func NewRorApplicationResource() *Resource {
	r := Resource{}
	r.Kind = "Application"
	r.APIVersion = "argoproj.io/v1alpha1"
	r.ApplicationResource = &rortypes.ResourceApplication{}
	r.common = r.ApplicationResource
	return &r
}

// NewRorAppProjectResource provides a empty resource of a given kind/apiversion
func NewRorAppProjectResource() *Resource {
	r := Resource{}
	r.Kind = "AppProject"
	r.APIVersion = "argoproj.io/v1alpha1"
	r.AppProjectResource = &rortypes.ResourceAppProject{}
	r.common = r.AppProjectResource
	return &r
}

// NewRorCertificateResource provides a empty resource of a given kind/apiversion
func NewRorCertificateResource() *Resource {
	r := Resource{}
	r.Kind = "Certificate"
	r.APIVersion = "cert-manager.io/v1"
	r.CertificateResource = &rortypes.ResourceCertificate{}
	r.common = r.CertificateResource
	return &r
}

// NewRorServiceResource provides a empty resource of a given kind/apiversion
func NewRorServiceResource() *Resource {
	r := Resource{}
	r.Kind = "Service"
	r.APIVersion = "v1"
	r.ServiceResource = &rortypes.ResourceService{}
	r.common = r.ServiceResource
	return &r
}

// NewRorPodResource provides a empty resource of a given kind/apiversion
func NewRorPodResource() *Resource {
	r := Resource{}
	r.Kind = "Pod"
	r.APIVersion = "v1"
	r.PodResource = &rortypes.ResourcePod{}
	r.common = r.PodResource
	return &r
}

// NewRorReplicaSetResource provides a empty resource of a given kind/apiversion
func NewRorReplicaSetResource() *Resource {
	r := Resource{}
	r.Kind = "ReplicaSet"
	r.APIVersion = "apps/v1"
	r.ReplicaSetResource = &rortypes.ResourceReplicaSet{}
	r.common = r.ReplicaSetResource
	return &r
}

// NewRorStatefulSetResource provides a empty resource of a given kind/apiversion
func NewRorStatefulSetResource() *Resource {
	r := Resource{}
	r.Kind = "StatefulSet"
	r.APIVersion = "apps/v1"
	r.StatefulSetResource = &rortypes.ResourceStatefulSet{}
	r.common = r.StatefulSetResource
	return &r
}

// NewRorDaemonSetResource provides a empty resource of a given kind/apiversion
func NewRorDaemonSetResource() *Resource {
	r := Resource{}
	r.Kind = "DaemonSet"
	r.APIVersion = "apps/v1"
	r.DaemonSetResource = &rortypes.ResourceDaemonSet{}
	r.common = r.DaemonSetResource
	return &r
}

// NewRorIngressResource provides a empty resource of a given kind/apiversion
func NewRorIngressResource() *Resource {
	r := Resource{}
	r.Kind = "Ingress"
	r.APIVersion = "networking.k8s.io/v1"
	r.IngressResource = &rortypes.ResourceIngress{}
	r.common = r.IngressResource
	return &r
}

// NewRorIngressClassResource provides a empty resource of a given kind/apiversion
func NewRorIngressClassResource() *Resource {
	r := Resource{}
	r.Kind = "IngressClass"
	r.APIVersion = "networking.k8s.io/v1"
	r.IngressClassResource = &rortypes.ResourceIngressClass{}
	r.common = r.IngressClassResource
	return &r
}

// NewRorSbomReportResource provides a empty resource of a given kind/apiversion
func NewRorSbomReportResource() *Resource {
	r := Resource{}
	r.Kind = "SbomReport"
	r.APIVersion = "aquasecurity.github.io/v1alpha1"
	r.SbomReportResource = &rortypes.ResourceSbomReport{}
	r.common = r.SbomReportResource
	return &r
}

// NewRorVulnerabilityReportResource provides a empty resource of a given kind/apiversion
func NewRorVulnerabilityReportResource() *Resource {
	r := Resource{}
	r.Kind = "VulnerabilityReport"
	r.APIVersion = "aquasecurity.github.io/v1alpha1"
	r.VulnerabilityReportResource = &rortypes.ResourceVulnerabilityReport{}
	r.common = r.VulnerabilityReportResource
	return &r
}

// NewRorExposedSecretReportResource provides a empty resource of a given kind/apiversion
func NewRorExposedSecretReportResource() *Resource {
	r := Resource{}
	r.Kind = "ExposedSecretReport"
	r.APIVersion = "aquasecurity.github.io/v1alpha1"
	r.ExposedSecretReportResource = &rortypes.ResourceExposedSecretReport{}
	r.common = r.ExposedSecretReportResource
	return &r
}

// NewRorConfigAuditReportResource provides a empty resource of a given kind/apiversion
func NewRorConfigAuditReportResource() *Resource {
	r := Resource{}
	r.Kind = "ConfigAuditReport"
	r.APIVersion = "aquasecurity.github.io/v1alpha1"
	r.ConfigAuditReportResource = &rortypes.ResourceConfigAuditReport{}
	r.common = r.ConfigAuditReportResource
	return &r
}

// NewRorRbacAssessmentReportResource provides a empty resource of a given kind/apiversion
func NewRorRbacAssessmentReportResource() *Resource {
	r := Resource{}
	r.Kind = "RbacAssessmentReport"
	r.APIVersion = "aquasecurity.github.io/v1alpha1"
	r.RbacAssessmentReportResource = &rortypes.ResourceRbacAssessmentReport{}
	r.common = r.RbacAssessmentReportResource
	return &r
}

// NewRorTanzuKubernetesClusterResource provides a empty resource of a given kind/apiversion
func NewRorTanzuKubernetesClusterResource() *Resource {
	r := Resource{}
	r.Kind = "TanzuKubernetesCluster"
	r.APIVersion = "run.tanzu.vmware.com/v1alpha3"
	r.TanzuKubernetesClusterResource = &rortypes.ResourceTanzuKubernetesCluster{}
	r.common = r.TanzuKubernetesClusterResource
	return &r
}

// NewRorTanzuKubernetesReleaseResource provides a empty resource of a given kind/apiversion
func NewRorTanzuKubernetesReleaseResource() *Resource {
	r := Resource{}
	r.Kind = "TanzuKubernetesRelease"
	r.APIVersion = "run.tanzu.vmware.com/v1alpha3"
	r.TanzuKubernetesReleaseResource = &rortypes.ResourceTanzuKubernetesRelease{}
	r.common = r.TanzuKubernetesReleaseResource
	return &r
}

// NewRorVirtualMachineClassResource provides a empty resource of a given kind/apiversion
func NewRorVirtualMachineClassResource() *Resource {
	r := Resource{}
	r.Kind = "VirtualMachineClass"
	r.APIVersion = "vmoperator.vmware.com/v1alpha2"
	r.VirtualMachineClassResource = &rortypes.ResourceVirtualMachineClass{}
	r.common = r.VirtualMachineClassResource
	return &r
}

// NewRorKubernetesClusterResource provides a empty resource of a given kind/apiversion
func NewRorKubernetesClusterResource() *Resource {
	r := Resource{}
	r.Kind = "KubernetesCluster"
	r.APIVersion = "vitistack.io/v1alpha1"
	r.KubernetesClusterResource = &rortypes.ResourceKubernetesCluster{}
	r.common = r.KubernetesClusterResource
	return &r
}

// NewRorProviderResource provides a empty resource of a given kind/apiversion
func NewRorProviderResource() *Resource {
	r := Resource{}
	r.Kind = "Provider"
	r.APIVersion = "general.ror.internal/v1alpha1"
	r.ProviderResource = &rortypes.ResourceProvider{}
	r.common = r.ProviderResource
	return &r
}

// NewRorWorkspaceResource provides a empty resource of a given kind/apiversion
func NewRorWorkspaceResource() *Resource {
	r := Resource{}
	r.Kind = "Workspace"
	r.APIVersion = "general.ror.internal/v1alpha1"
	r.WorkspaceResource = &rortypes.ResourceWorkspace{}
	r.common = r.WorkspaceResource
	return &r
}

// NewRorKubernetesMachineClassResource provides a empty resource of a given kind/apiversion
func NewRorKubernetesMachineClassResource() *Resource {
	r := Resource{}
	r.Kind = "KubernetesMachineClass"
	r.APIVersion = "general.ror.internal/v1alpha1"
	r.KubernetesMachineClassResource = &rortypes.ResourceKubernetesMachineClass{}
	r.common = r.KubernetesMachineClassResource
	return &r
}

// NewRorClusterOrderResource provides a empty resource of a given kind/apiversion
func NewRorClusterOrderResource() *Resource {
	r := Resource{}
	r.Kind = "ClusterOrder"
	r.APIVersion = "general.ror.internal/v1alpha1"
	r.ClusterOrderResource = &rortypes.ResourceClusterOrder{}
	r.common = r.ClusterOrderResource
	return &r
}

// NewRorProjectResource provides a empty resource of a given kind/apiversion
func NewRorProjectResource() *Resource {
	r := Resource{}
	r.Kind = "Project"
	r.APIVersion = "general.ror.internal/v1alpha1"
	r.ProjectResource = &rortypes.ResourceProject{}
	r.common = r.ProjectResource
	return &r
}

// NewRorConfigurationResource provides a empty resource of a given kind/apiversion
func NewRorConfigurationResource() *Resource {
	r := Resource{}
	r.Kind = "Configuration"
	r.APIVersion = "general.ror.internal/v1alpha1"
	r.ConfigurationResource = &rortypes.ResourceConfiguration{}
	r.common = r.ConfigurationResource
	return &r
}

// NewRorClusterComplianceReportResource provides a empty resource of a given kind/apiversion
func NewRorClusterComplianceReportResource() *Resource {
	r := Resource{}
	r.Kind = "ClusterComplianceReport"
	r.APIVersion = "aquasecurity.github.io/v1alpha1"
	r.ClusterComplianceReportResource = &rortypes.ResourceClusterComplianceReport{}
	r.common = r.ClusterComplianceReportResource
	return &r
}

// NewRorClusterVulnerabilityReportResource provides a empty resource of a given kind/apiversion
func NewRorClusterVulnerabilityReportResource() *Resource {
	r := Resource{}
	r.Kind = "ClusterVulnerabilityReport"
	r.APIVersion = "general.ror.internal/v1alpha1"
	r.ClusterVulnerabilityReportResource = &rortypes.ResourceClusterVulnerabilityReport{}
	r.common = r.ClusterVulnerabilityReportResource
	return &r
}

// NewRorRouteResource provides a empty resource of a given kind/apiversion
func NewRorRouteResource() *Resource {
	r := Resource{}
	r.Kind = "Route"
	r.APIVersion = "general.ror.internal/v1alpha1"
	r.RouteResource = &rortypes.ResourceRoute{}
	r.common = r.RouteResource
	return &r
}

// NewRorSlackMessageResource provides a empty resource of a given kind/apiversion
func NewRorSlackMessageResource() *Resource {
	r := Resource{}
	r.Kind = "SlackMessage"
	r.APIVersion = "general.ror.internal/v1alpha1"
	r.SlackMessageResource = &rortypes.ResourceSlackMessage{}
	r.common = r.SlackMessageResource
	return &r
}

// NewRorVulnerabilityEventResource provides a empty resource of a given kind/apiversion
func NewRorVulnerabilityEventResource() *Resource {
	r := Resource{}
	r.Kind = "VulnerabilityEvent"
	r.APIVersion = "general.ror.internal/v1alpha1"
	r.VulnerabilityEventResource = &rortypes.ResourceVulnerabilityEvent{}
	r.common = r.VulnerabilityEventResource
	return &r
}

// NewRorVirtualMachineResource provides a empty resource of a given kind/apiversion
func NewRorVirtualMachineResource() *Resource {
	r := Resource{}
	r.Kind = "VirtualMachine"
	r.APIVersion = "general.ror.internal/v1alpha1"
	r.VirtualMachineResource = &rortypes.ResourceVirtualMachine{}
	r.common = r.VirtualMachineResource
	return &r
}

// NewRorVirtualMachineVulnerabilityInfoResource provides a empty resource of a given kind/apiversion
func NewRorVirtualMachineVulnerabilityInfoResource() *Resource {
	r := Resource{}
	r.Kind = "VirtualMachineVulnerabilityInfo"
	r.APIVersion = "general.ror.internal/v1alpha1"
	r.VirtualMachineVulnerabilityInfoResource = &rortypes.ResourceVirtualMachineVulnerabilityInfo{}
	r.common = r.VirtualMachineVulnerabilityInfoResource
	return &r
}

// NewRorEndpointsResource provides a empty resource of a given kind/apiversion
func NewRorEndpointsResource() *Resource {
	r := Resource{}
	r.Kind = "Endpoints"
	r.APIVersion = "v1"
	r.EndpointsResource = &rortypes.ResourceEndpoints{}
	r.common = r.EndpointsResource
	return &r
}

// NewRorNetworkPolicyResource provides a empty resource of a given kind/apiversion
func NewRorNetworkPolicyResource() *Resource {
	r := Resource{}
	r.Kind = "NetworkPolicy"
	r.APIVersion = "networking.k8s.io/v1"
	r.NetworkPolicyResource = &rortypes.ResourceNetworkPolicy{}
	r.common = r.NetworkPolicyResource
	return &r
}

// NewRorDatacenterResource provides a empty resource of a given kind/apiversion
func NewRorDatacenterResource() *Resource {
	r := Resource{}
	r.Kind = "Datacenter"
	r.APIVersion = "infrastructure.ror.internal/v1alpha1"
	r.DatacenterResource = &rortypes.ResourceDatacenter{}
	r.common = r.DatacenterResource
	return &r
}

// NewRorBackupJobResource provides a empty resource of a given kind/apiversion
func NewRorBackupJobResource() *Resource {
	r := Resource{}
	r.Kind = "BackupJob"
	r.APIVersion = "backup.ror.internal/v1alpha1"
	r.BackupJobResource = &rortypes.ResourceBackupJob{}
	r.common = r.BackupJobResource
	return &r
}

// NewRorBackupRunResource provides a empty resource of a given kind/apiversion
func NewRorBackupRunResource() *Resource {
	r := Resource{}
	r.Kind = "BackupRun"
	r.APIVersion = "backup.ror.internal/v1alpha1"
	r.BackupRunResource = &rortypes.ResourceBackupRun{}
	r.common = r.BackupRunResource
	return &r
}

// NewRorUnknownResource provides a empty resource of a given kind/apiversion
func NewRorUnknownResource() *Resource {
	r := Resource{}
	r.Kind = "Unknown"
	r.APIVersion = "unknown.ror.internal/v1"
	r.UnknownResource = &rortypes.ResourceUnknown{}
	r.common = r.UnknownResource
	return &r
}

// SetCommonResource sets the common resource of the resource, the common resource implements common metadata of the resource
func (r *Resource) SetCommonResource(common rortypes.CommonResource) {
	r.CommonResource = common
}

// SetCommonInterface sets the common interface of the resource, the common interface implements common methods of the resource
func (r *Resource) SetCommonInterface(common rortypes.CommonResourceInterface) {
	r.common = common
}

func (r *Resource) SetNamespace(res *rortypes.ResourceNamespace) {
	r.NamespaceResource = res
}

func (r *Resource) SetNode(res *rortypes.ResourceNode) {
	r.NodeResource = res
}

func (r *Resource) SetPersistentVolumeClaim(res *rortypes.ResourcePersistentVolumeClaim) {
	r.PersistentVolumeClaimResource = res
}

func (r *Resource) SetDeployment(res *rortypes.ResourceDeployment) {
	r.DeploymentResource = res
}

func (r *Resource) SetStorageClass(res *rortypes.ResourceStorageClass) {
	r.StorageClassResource = res
}

func (r *Resource) SetPolicyReport(res *rortypes.ResourcePolicyReport) {
	r.PolicyReportResource = res
}

func (r *Resource) SetApplication(res *rortypes.ResourceApplication) {
	r.ApplicationResource = res
}

func (r *Resource) SetAppProject(res *rortypes.ResourceAppProject) {
	r.AppProjectResource = res
}

func (r *Resource) SetCertificate(res *rortypes.ResourceCertificate) {
	r.CertificateResource = res
}

func (r *Resource) SetService(res *rortypes.ResourceService) {
	r.ServiceResource = res
}

func (r *Resource) SetPod(res *rortypes.ResourcePod) {
	r.PodResource = res
}

func (r *Resource) SetReplicaSet(res *rortypes.ResourceReplicaSet) {
	r.ReplicaSetResource = res
}

func (r *Resource) SetStatefulSet(res *rortypes.ResourceStatefulSet) {
	r.StatefulSetResource = res
}

func (r *Resource) SetDaemonSet(res *rortypes.ResourceDaemonSet) {
	r.DaemonSetResource = res
}

func (r *Resource) SetIngress(res *rortypes.ResourceIngress) {
	r.IngressResource = res
}

func (r *Resource) SetIngressClass(res *rortypes.ResourceIngressClass) {
	r.IngressClassResource = res
}

func (r *Resource) SetSbomReport(res *rortypes.ResourceSbomReport) {
	r.SbomReportResource = res
}

func (r *Resource) SetVulnerabilityReport(res *rortypes.ResourceVulnerabilityReport) {
	r.VulnerabilityReportResource = res
}

func (r *Resource) SetExposedSecretReport(res *rortypes.ResourceExposedSecretReport) {
	r.ExposedSecretReportResource = res
}

func (r *Resource) SetConfigAuditReport(res *rortypes.ResourceConfigAuditReport) {
	r.ConfigAuditReportResource = res
}

func (r *Resource) SetRbacAssessmentReport(res *rortypes.ResourceRbacAssessmentReport) {
	r.RbacAssessmentReportResource = res
}

func (r *Resource) SetTanzuKubernetesCluster(res *rortypes.ResourceTanzuKubernetesCluster) {
	r.TanzuKubernetesClusterResource = res
}

func (r *Resource) SetTanzuKubernetesRelease(res *rortypes.ResourceTanzuKubernetesRelease) {
	r.TanzuKubernetesReleaseResource = res
}

func (r *Resource) SetVirtualMachineClass(res *rortypes.ResourceVirtualMachineClass) {
	r.VirtualMachineClassResource = res
}

func (r *Resource) SetKubernetesCluster(res *rortypes.ResourceKubernetesCluster) {
	r.KubernetesClusterResource = res
}

func (r *Resource) SetProvider(res *rortypes.ResourceProvider) {
	r.ProviderResource = res
}

func (r *Resource) SetWorkspace(res *rortypes.ResourceWorkspace) {
	r.WorkspaceResource = res
}

func (r *Resource) SetKubernetesMachineClass(res *rortypes.ResourceKubernetesMachineClass) {
	r.KubernetesMachineClassResource = res
}

func (r *Resource) SetClusterOrder(res *rortypes.ResourceClusterOrder) {
	r.ClusterOrderResource = res
}

func (r *Resource) SetProject(res *rortypes.ResourceProject) {
	r.ProjectResource = res
}

func (r *Resource) SetConfiguration(res *rortypes.ResourceConfiguration) {
	r.ConfigurationResource = res
}

func (r *Resource) SetClusterComplianceReport(res *rortypes.ResourceClusterComplianceReport) {
	r.ClusterComplianceReportResource = res
}

func (r *Resource) SetClusterVulnerabilityReport(res *rortypes.ResourceClusterVulnerabilityReport) {
	r.ClusterVulnerabilityReportResource = res
}

func (r *Resource) SetRoute(res *rortypes.ResourceRoute) {
	r.RouteResource = res
}

func (r *Resource) SetSlackMessage(res *rortypes.ResourceSlackMessage) {
	r.SlackMessageResource = res
}

func (r *Resource) SetVulnerabilityEvent(res *rortypes.ResourceVulnerabilityEvent) {
	r.VulnerabilityEventResource = res
}

func (r *Resource) SetVirtualMachine(res *rortypes.ResourceVirtualMachine) {
	r.VirtualMachineResource = res
}

func (r *Resource) SetVirtualMachineVulnerabilityInfo(res *rortypes.ResourceVirtualMachineVulnerabilityInfo) {
	r.VirtualMachineVulnerabilityInfoResource = res
}

func (r *Resource) SetEndpoints(res *rortypes.ResourceEndpoints) {
	r.EndpointsResource = res
}

func (r *Resource) SetNetworkPolicy(res *rortypes.ResourceNetworkPolicy) {
	r.NetworkPolicyResource = res
}

func (r *Resource) SetDatacenter(res *rortypes.ResourceDatacenter) {
	r.DatacenterResource = res
}

func (r *Resource) SetBackupJob(res *rortypes.ResourceBackupJob) {
	r.BackupJobResource = res
}

func (r *Resource) SetBackupRun(res *rortypes.ResourceBackupRun) {
	r.BackupRunResource = res
}

func (r *Resource) SetMachine(res *rortypes.ResourceMachine) {
	r.MachineResource = res
}

func (r *Resource) SetUnknown(res *rortypes.ResourceUnknown) {
	r.UnknownResource = res
}

// Namespace is a wrapper for the underlying resource, it provides a Namespaceinterface to work with namespaces
func (r *Resource) Namespace() rortypes.Namespaceinterface {
	return r.NamespaceResource
}

// Node is a wrapper for the underlying resource, it provides a Nodeinterface to work with nodes
func (r *Resource) Node() rortypes.Nodeinterface {
	return r.NodeResource
}

// PersistentVolumeClaim is a wrapper for the underlying resource, it provides a PersistentVolumeClaiminterface to work with persistentvolumeclaims
func (r *Resource) PersistentVolumeClaim() rortypes.PersistentVolumeClaiminterface {
	return r.PersistentVolumeClaimResource
}

// Deployment is a wrapper for the underlying resource, it provides a Deploymentinterface to work with deployments
func (r *Resource) Deployment() rortypes.Deploymentinterface {
	return r.DeploymentResource
}

// StorageClass is a wrapper for the underlying resource, it provides a StorageClassinterface to work with storageclasses
func (r *Resource) StorageClass() rortypes.StorageClassinterface {
	return r.StorageClassResource
}

// PolicyReport is a wrapper for the underlying resource, it provides a PolicyReportinterface to work with policyreports
func (r *Resource) PolicyReport() rortypes.PolicyReportinterface {
	return r.PolicyReportResource
}

// Application is a wrapper for the underlying resource, it provides a Applicationinterface to work with applications
func (r *Resource) Application() rortypes.Applicationinterface {
	return r.ApplicationResource
}

// AppProject is a wrapper for the underlying resource, it provides a AppProjectinterface to work with appprojects
func (r *Resource) AppProject() rortypes.AppProjectinterface {
	return r.AppProjectResource
}

// Certificate is a wrapper for the underlying resource, it provides a Certificateinterface to work with certificates
func (r *Resource) Certificate() rortypes.Certificateinterface {
	return r.CertificateResource
}

// Service is a wrapper for the underlying resource, it provides a Serviceinterface to work with services
func (r *Resource) Service() rortypes.Serviceinterface {
	return r.ServiceResource
}

// Pod is a wrapper for the underlying resource, it provides a Podinterface to work with pods
func (r *Resource) Pod() rortypes.Podinterface {
	return r.PodResource
}

// ReplicaSet is a wrapper for the underlying resource, it provides a ReplicaSetinterface to work with replicasets
func (r *Resource) ReplicaSet() rortypes.ReplicaSetinterface {
	return r.ReplicaSetResource
}

// StatefulSet is a wrapper for the underlying resource, it provides a StatefulSetinterface to work with statefulsets
func (r *Resource) StatefulSet() rortypes.StatefulSetinterface {
	return r.StatefulSetResource
}

// DaemonSet is a wrapper for the underlying resource, it provides a DaemonSetinterface to work with daemonsets
func (r *Resource) DaemonSet() rortypes.DaemonSetinterface {
	return r.DaemonSetResource
}

// Ingress is a wrapper for the underlying resource, it provides a Ingressinterface to work with ingresses
func (r *Resource) Ingress() rortypes.Ingressinterface {
	return r.IngressResource
}

// IngressClass is a wrapper for the underlying resource, it provides a IngressClassinterface to work with ingressclasses
func (r *Resource) IngressClass() rortypes.IngressClassinterface {
	return r.IngressClassResource
}

// SbomReport is a wrapper for the underlying resource, it provides a SbomReportinterface to work with sbomreports
func (r *Resource) SbomReport() rortypes.SbomReportinterface {
	return r.SbomReportResource
}

// VulnerabilityReport is a wrapper for the underlying resource, it provides a VulnerabilityReportinterface to work with vulnerabilityreports
func (r *Resource) VulnerabilityReport() rortypes.VulnerabilityReportinterface {
	return r.VulnerabilityReportResource
}

// ExposedSecretReport is a wrapper for the underlying resource, it provides a ExposedSecretReportinterface to work with exposedsecretreports
func (r *Resource) ExposedSecretReport() rortypes.ExposedSecretReportinterface {
	return r.ExposedSecretReportResource
}

// ConfigAuditReport is a wrapper for the underlying resource, it provides a ConfigAuditReportinterface to work with configauditreports
func (r *Resource) ConfigAuditReport() rortypes.ConfigAuditReportinterface {
	return r.ConfigAuditReportResource
}

// RbacAssessmentReport is a wrapper for the underlying resource, it provides a RbacAssessmentReportinterface to work with rbacassessmentreports
func (r *Resource) RbacAssessmentReport() rortypes.RbacAssessmentReportinterface {
	return r.RbacAssessmentReportResource
}

// TanzuKubernetesCluster is a wrapper for the underlying resource, it provides a TanzuKubernetesClusterinterface to work with tanzukubernetesclusters
func (r *Resource) TanzuKubernetesCluster() rortypes.TanzuKubernetesClusterinterface {
	return r.TanzuKubernetesClusterResource
}

// TanzuKubernetesRelease is a wrapper for the underlying resource, it provides a TanzuKubernetesReleaseinterface to work with tanzukubernetesreleases
func (r *Resource) TanzuKubernetesRelease() rortypes.TanzuKubernetesReleaseinterface {
	return r.TanzuKubernetesReleaseResource
}

// VirtualMachineClass is a wrapper for the underlying resource, it provides a VirtualMachineClassinterface to work with virtualmachineclasses
func (r *Resource) VirtualMachineClass() rortypes.VirtualMachineClassinterface {
	return r.VirtualMachineClassResource
}

// KubernetesCluster is a wrapper for the underlying resource, it provides a KubernetesClusterinterface to work with kubernetesclusters
func (r *Resource) KubernetesCluster() rortypes.KubernetesClusterinterface {
	return r.KubernetesClusterResource
}

// Provider is a wrapper for the underlying resource, it provides a Providerinterface to work with providers
func (r *Resource) Provider() rortypes.Providerinterface {
	return r.ProviderResource
}

// Workspace is a wrapper for the underlying resource, it provides a Workspaceinterface to work with workspaces
func (r *Resource) Workspace() rortypes.Workspaceinterface {
	return r.WorkspaceResource
}

// KubernetesMachineClass is a wrapper for the underlying resource, it provides a KubernetesMachineClassinterface to work with Kubernetesmachineclasses
func (r *Resource) KubernetesMachineClass() rortypes.KubernetesMachineClassinterface {
	return r.KubernetesMachineClassResource
}

// ClusterOrder is a wrapper for the underlying resource, it provides a ClusterOrderinterface to work with clusterorders
func (r *Resource) ClusterOrder() rortypes.ClusterOrderinterface {
	return r.ClusterOrderResource
}

// Project is a wrapper for the underlying resource, it provides a Projectinterface to work with projects
func (r *Resource) Project() rortypes.Projectinterface {
	return r.ProjectResource
}

// Configuration is a wrapper for the underlying resource, it provides a Configurationinterface to work with configurations
func (r *Resource) Configuration() rortypes.Configurationinterface {
	return r.ConfigurationResource
}

// ClusterComplianceReport is a wrapper for the underlying resource, it provides a ClusterComplianceReportinterface to work with clustercompliancereports
func (r *Resource) ClusterComplianceReport() rortypes.ClusterComplianceReportinterface {
	return r.ClusterComplianceReportResource
}

// ClusterVulnerabilityReport is a wrapper for the underlying resource, it provides a ClusterVulnerabilityReportinterface to work with clustervulnerabilityreports
func (r *Resource) ClusterVulnerabilityReport() rortypes.ClusterVulnerabilityReportinterface {
	return r.ClusterVulnerabilityReportResource
}

// Route is a wrapper for the underlying resource, it provides a Routeinterface to work with routes
func (r *Resource) Route() rortypes.Routeinterface {
	return r.RouteResource
}

// SlackMessage is a wrapper for the underlying resource, it provides a SlackMessageinterface to work with slackmessages
func (r *Resource) SlackMessage() rortypes.SlackMessageinterface {
	return r.SlackMessageResource
}

// VulnerabilityEvent is a wrapper for the underlying resource, it provides a VulnerabilityEventinterface to work with vulnerabilityevents
func (r *Resource) VulnerabilityEvent() rortypes.VulnerabilityEventinterface {
	return r.VulnerabilityEventResource
}

// VirtualMachine is a wrapper for the underlying resource, it provides a VirtualMachineinterface to work with VirtualMachines
func (r *Resource) VirtualMachine() rortypes.VirtualMachineinterface {
	return r.VirtualMachineResource
}

// VirtualMachineVulnerabilityInfo is a wrapper for the underlying resource, it provides a VirtualMachineVulnerabilityInfointerface to work with VirtualMachinesVulnerabilityInfo
func (r *Resource) VirtualMachineVulnerabilityInfo() rortypes.VirtualMachineVulnerabilityInfointerface {
	return r.VirtualMachineVulnerabilityInfoResource
}

// Endpoints is a wrapper for the underlying resource, it provides a Endpointsinterface to work with endpoints
func (r *Resource) Endpoints() rortypes.Endpointsinterface {
	return r.EndpointsResource
}

// NetworkPolicy is a wrapper for the underlying resource, it provides a NetworkPolicyinterface to work with networkpolicies
func (r *Resource) NetworkPolicy() rortypes.NetworkPolicyinterface {
	return r.NetworkPolicyResource
}

// Datacenter is a wrapper for the underlying resource, it provides a Datacenterinterface to work with datacenters
func (r *Resource) Datacenter() rortypes.Datacenterinterface {
	return r.DatacenterResource
}

// BackupJob is a wrapper for the underlying resource, it provides a BackupJobinterface to work with backupjobs
func (r *Resource) BackupJob() rortypes.BackupJobinterface {
	return r.BackupJobResource
}

// BackupRun is a wrapper for the underlying resource, it provides a BackupRuninterface to work with backupruns
func (r *Resource) BackupRun() rortypes.BackupRuninterface {
	return r.BackupRunResource
}

// Machine is a wrapper for the underlying resource, it provides a Machineinterface to work with machines
func (r *Resource) Machine() rortypes.Machineinterface {
	return r.MachineResource
}

// Unknown is a wrapper for the underlying resource, it provides a Unknowninterface to work with unknowns
func (r *Resource) Unknown() rortypes.Unknowninterface {
	return r.UnknownResource
}

// (r *Resource) GetRorHash() returns the hash from the common interface
func (r *Resource) GetRorHash() string {
	return r.common.GetRorHash()
}

// (r *Resource) GenRorHash() calculates the hash of the resource and set the metadata header
func (r *Resource) GenRorHash() {
	hash := r.common.GetRorHash()
	r.CommonResource.RorMeta.Hash = hash
}

func (r *Resource) ApplyInputFilter() error {
	return r.common.ApplyInputFilter(&r.CommonResource)
}
