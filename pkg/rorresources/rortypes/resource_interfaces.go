// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package rortypes

// CommonResourceInterface represents the minimum interface for all resources
type CommonResourceInterface interface {
	GetRorHash() string
	ApplyInputFilter(cr *CommonResource) error
	ApplyOutputFilter(cr *CommonResource) error
}

// Namespaceinterface represents the interface for resources of the type namespace
type Namespaceinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceNamespace
}

// Nodeinterface represents the interface for resources of the type node
type Nodeinterface interface {
	Get() *ResourceNode
}

// PersistentVolumeClaiminterface represents the interface for resources of the type persistentvolumeclaim
type PersistentVolumeClaiminterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourcePersistentVolumeClaim
}

// Deploymentinterface represents the interface for resources of the type deployment
type Deploymentinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceDeployment
}

// StorageClassinterface represents the interface for resources of the type storageclass
type StorageClassinterface interface {
	Get() *ResourceStorageClass
}

// PolicyReportinterface represents the interface for resources of the type policyreport
type PolicyReportinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourcePolicyReport
}

// Applicationinterface represents the interface for resources of the type application
type Applicationinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceApplication
}

// AppProjectinterface represents the interface for resources of the type appproject
type AppProjectinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceAppProject
}

// Certificateinterface represents the interface for resources of the type certificate
type Certificateinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceCertificate
}

// Serviceinterface represents the interface for resources of the type service
type Serviceinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceService
}

// Podinterface represents the interface for resources of the type pod
type Podinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourcePod
}

// ReplicaSetinterface represents the interface for resources of the type replicaset
type ReplicaSetinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceReplicaSet
}

// StatefulSetinterface represents the interface for resources of the type statefulset
type StatefulSetinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceStatefulSet
}

// DaemonSetinterface represents the interface for resources of the type daemonset
type DaemonSetinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceDaemonSet
}

// Ingressinterface represents the interface for resources of the type ingress
type Ingressinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceIngress
}

// IngressClassinterface represents the interface for resources of the type ingressclass
type IngressClassinterface interface {
	Get() *ResourceIngressClass
}

// SbomReportinterface represents the interface for resources of the type sbomreport
type SbomReportinterface interface {
	Get() *ResourceSbomReport
}

// VulnerabilityReportinterface represents the interface for resources of the type vulnerabilityreport
type VulnerabilityReportinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceVulnerabilityReport
}

// ExposedSecretReportinterface represents the interface for resources of the type exposedsecretreport
type ExposedSecretReportinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceExposedSecretReport
}

// ConfigAuditReportinterface represents the interface for resources of the type configauditreport
type ConfigAuditReportinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceConfigAuditReport
}

// RbacAssessmentReportinterface represents the interface for resources of the type rbacassessmentreport
type RbacAssessmentReportinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceRbacAssessmentReport
}

// TanzuKubernetesClusterinterface represents the interface for resources of the type tanzukubernetescluster
type TanzuKubernetesClusterinterface interface {
	Get() *ResourceTanzuKubernetesCluster
}

// TanzuKubernetesReleaseinterface represents the interface for resources of the type tanzukubernetesrelease
type TanzuKubernetesReleaseinterface interface {
	Get() *ResourceTanzuKubernetesRelease
}

// VirtualMachineClassinterface represents the interface for resources of the type virtualmachineclass
type VirtualMachineClassinterface interface {
	Get() *ResourceVirtualMachineClass
}

// KubernetesClusterinterface represents the interface for resources of the type kubernetescluster
type KubernetesClusterinterface interface {
	Get() *ResourceKubernetesCluster
}

// Providerinterface represents the interface for resources of the type provider
type Providerinterface interface {
	Get() *ResourceProvider
}

// Workspaceinterface represents the interface for resources of the type workspace
type Workspaceinterface interface {
	Get() *ResourceWorkspace
}

// KubernetesMachineClassinterface represents the interface for resources of the type kubernetesmachineclass
type KubernetesMachineClassinterface interface {
	Get() *ResourceKubernetesMachineClass
}

// ClusterOrderinterface represents the interface for resources of the type clusterorder
type ClusterOrderinterface interface {
	Get() *ResourceClusterOrder
}

// Projectinterface represents the interface for resources of the type project
type Projectinterface interface {
	Get() *ResourceProject
}

// Configurationinterface represents the interface for resources of the type configuration
type Configurationinterface interface {
	Get() *ResourceConfiguration
}

// ClusterComplianceReportinterface represents the interface for resources of the type clustercompliancereport
type ClusterComplianceReportinterface interface {
	Get() *ResourceClusterComplianceReport
}

// ClusterVulnerabilityReportinterface represents the interface for resources of the type clustervulnerabilityreport
type ClusterVulnerabilityReportinterface interface {
	Get() *ResourceClusterVulnerabilityReport
}

// Routeinterface represents the interface for resources of the type route
type Routeinterface interface {
	Get() *ResourceRoute
}

// SlackMessageinterface represents the interface for resources of the type slackmessage
type SlackMessageinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceSlackMessage
}

// VulnerabilityEventinterface represents the interface for resources of the type vulnerabilityevent
type VulnerabilityEventinterface interface {
	Get() *ResourceVulnerabilityEvent
}

// VirtualMachineinterface represents the interface for resources of the type virtualmachine
type VirtualMachineinterface interface {
	Get() *ResourceVirtualMachine
}

// VirtualMachineVulnerabilityInfointerface represents the interface for resources of the type virtualmachinevulnerabilityinfo
type VirtualMachineVulnerabilityInfointerface interface {
	Get() *ResourceVirtualMachineVulnerabilityInfo
}

// Endpointsinterface represents the interface for resources of the type endpoints
type Endpointsinterface interface {
	Get() *ResourceEndpoints
}

// NetworkPolicyinterface represents the interface for resources of the type networkpolicy
type NetworkPolicyinterface interface {
	Get() *ResourceNetworkPolicy
}

// Datacenterinterface represents the interface for resources of the type datacenter
type Datacenterinterface interface {
	Get() *ResourceDatacenter
}

// BackupJobinterface represents the interface for resources of the type backupjob
type BackupJobinterface interface {
	Get() *ResourceBackupJob
}

// BackupRuninterface represents the interface for resources of the type backuprun
type BackupRuninterface interface {
	Get() *ResourceBackupRun
}

// Machineinterface represents the interface for resources of the type machine
type Machineinterface interface {
	Get() *ResourceMachine
}

// Unknowninterface represents the interface for resources of the type unknown
type Unknowninterface interface {
	Get() *ResourceUnknown
}
