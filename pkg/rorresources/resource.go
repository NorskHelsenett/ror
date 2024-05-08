// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package rorresources

import (
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// The Resource struct represents one resource in ror.
//
// It implement common and resource specific methods by providing interfaces to the underlying resources
type Resource struct {
	Kind       string `json:"kind,omitempty"`
	ApiVersion string `json:"api_version,omitempty"`

	NamespaceResource                  *rortypes.ResourceNamespace                  `json:"namespace,omitempty"`
	NodeResource                       *rortypes.ResourceNode                       `json:"node,omitempty"`
	PersistentVolumeClaimResource      *rortypes.ResourcePersistentVolumeClaim      `json:"persistentvolumeclaim,omitempty"`
	DeploymentResource                 *rortypes.ResourceDeployment                 `json:"deployment,omitempty"`
	StorageClassResource               *rortypes.ResourceStorageClass               `json:"storageclass,omitempty"`
	PolicyReportResource               *rortypes.ResourcePolicyReport               `json:"policyreport,omitempty"`
	ApplicationResource                *rortypes.ResourceApplication                `json:"application,omitempty"`
	AppProjectResource                 *rortypes.ResourceAppProject                 `json:"appproject,omitempty"`
	CertificateResource                *rortypes.ResourceCertificate                `json:"certificate,omitempty"`
	ServiceResource                    *rortypes.ResourceService                    `json:"service,omitempty"`
	PodResource                        *rortypes.ResourcePod                        `json:"pod,omitempty"`
	ReplicaSetResource                 *rortypes.ResourceReplicaSet                 `json:"replicaset,omitempty"`
	StatefulSetResource                *rortypes.ResourceStatefulSet                `json:"statefulset,omitempty"`
	DaemonSetResource                  *rortypes.ResourceDaemonSet                  `json:"daemonset,omitempty"`
	IngressResource                    *rortypes.ResourceIngress                    `json:"ingress,omitempty"`
	IngressClassResource               *rortypes.ResourceIngressClass               `json:"ingressclass,omitempty"`
	VulnerabilityReportResource        *rortypes.ResourceVulnerabilityReport        `json:"vulnerabilityreport,omitempty"`
	ExposedSecretReportResource        *rortypes.ResourceExposedSecretReport        `json:"exposedsecretreport,omitempty"`
	ConfigAuditReportResource          *rortypes.ResourceConfigAuditReport          `json:"configauditreport,omitempty"`
	RbacAssessmentReportResource       *rortypes.ResourceRbacAssessmentReport       `json:"rbacassessmentreport,omitempty"`
	TanzuKubernetesClusterResource     *rortypes.ResourceTanzuKubernetesCluster     `json:"tanzukubernetescluster,omitempty"`
	TanzuKubernetesReleaseResource     *rortypes.ResourceTanzuKubernetesRelease     `json:"tanzukubernetesrelease,omitempty"`
	VirtualMachineClassResource        *rortypes.ResourceVirtualMachineClass        `json:"virtualmachineclass,omitempty"`
	VirtualMachineClassBindingResource *rortypes.ResourceVirtualMachineClassBinding `json:"virtualmachineclassbinding,omitempty"`
	KubernetesClusterResource          *rortypes.ResourceKubernetesCluster          `json:"kubernetescluster,omitempty"`
	ClusterOrderResource               *rortypes.ResourceClusterOrder               `json:"clusterorder,omitempty"`
	ProjectResource                    *rortypes.ResourceProject                    `json:"project,omitempty"`
	ConfigurationResource              *rortypes.ResourceConfiguration              `json:"configuration,omitempty"`
	ClusterComplianceReportResource    *rortypes.ResourceClusterComplianceReport    `json:"clustercompliancereport,omitempty"`

	common rortypes.CommonResourceInterface
}

// NewRorResource provides a empty resource of a given kind/apiversion
func NewRorResource(kind string, apiversion string) *Resource {
	r := Resource{Kind: kind, ApiVersion: apiversion}
	return &r
}

// SetCommon sets the common interface of the resource, the common interface implements common methods of the resource
func (r *Resource) SetCommon(common rortypes.CommonResourceInterface) {
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

func (r *Resource) SetVirtualMachineClassBinding(res *rortypes.ResourceVirtualMachineClassBinding) {
	r.VirtualMachineClassBindingResource = res
}

func (r *Resource) SetKubernetesCluster(res *rortypes.ResourceKubernetesCluster) {
	r.KubernetesClusterResource = res
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

// VirtualMachineClassBinding is a wrapper for the underlying resource, it provides a VirtualMachineClassBindinginterface to work with virtualmachineclassbindings
func (r *Resource) VirtualMachineClassBinding() rortypes.VirtualMachineClassBindinginterface {
	return r.VirtualMachineClassBindingResource
}

// KubernetesCluster is a wrapper for the underlying resource, it provides a KubernetesClusterinterface to work with kubernetesclusters
func (r *Resource) KubernetesCluster() rortypes.KubernetesClusterinterface {
	return r.KubernetesClusterResource
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

// (r *Resource) GetName() returns the name from the common interface
func (r *Resource) GetName() string {
	return r.common.GetName()
}

// (r *Resource) GetUID() returns the UID from the common interface
func (r *Resource) GetUID() string {
	return r.common.GetUID()
}

// (r *Resource) GetKind() returns the Kind from the common interface
func (r *Resource) GetKind() string {
	return r.common.GetKind()
}

// (r *Resource) GetAPIVersion() returns the APIVersion from the common interface
func (r *Resource) GetAPIVersion() string {
	return r.common.GetAPIVersion()
}

// (r *Resource) GetMetadata() returns the Metadata from the common interface
func (r *Resource) GetMetadata() metav1.ObjectMeta {
	return r.common.GetMetadata()
}

// (r *Resource) GetRorMeta() returns the RorMetadata from the common interface
func (r *Resource) GetRorMeta() rortypes.ResourceRorMeta {
	return r.common.GetRorMeta()
}

// (r *Resource) SetRorMeta() returns the RorMetadata from the common interface
func (r *Resource) SetRorMeta(input rortypes.ResourceRorMeta) error {
	return r.common.SetRorMeta(input)
}

// (r *Resource) GetRorHash() returns the hash from the common interface
func (r *Resource) GetRorHash() string {
	return r.common.GetRorHash()
}

func (r *Resource) ApplyInputFilter() error {
	return r.common.ApplyInputFilter()
}
