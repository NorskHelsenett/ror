// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package rortypes

// CommonResourceInterface represents the minimum interface for all resources
type CommonResourceInterface interface {
	GetRorHash() string
	ApplyInputFilter(cr *CommonResource) error
}

// Namespaceinterface represents the interface for resources of the type namespace
type Namespaceinterface interface {
	CommonResourceInterface
	Get() *ResourceNamespace
}

// Nodeinterface represents the interface for resources of the type node
type Nodeinterface interface {
	CommonResourceInterface
	Get() *ResourceNode
}

// PersistentVolumeClaiminterface represents the interface for resources of the type persistentvolumeclaim
type PersistentVolumeClaiminterface interface {
	CommonResourceInterface
	Get() *ResourcePersistentVolumeClaim
}

// Deploymentinterface represents the interface for resources of the type deployment
type Deploymentinterface interface {
	CommonResourceInterface
	Get() *ResourceDeployment
}

// StorageClassinterface represents the interface for resources of the type storageclass
type StorageClassinterface interface {
	CommonResourceInterface
	Get() *ResourceStorageClass
}

// PolicyReportinterface represents the interface for resources of the type policyreport
type PolicyReportinterface interface {
	CommonResourceInterface
	Get() *ResourcePolicyReport
}

// Applicationinterface represents the interface for resources of the type application
type Applicationinterface interface {
	CommonResourceInterface
	Get() *ResourceApplication
}

// AppProjectinterface represents the interface for resources of the type appproject
type AppProjectinterface interface {
	CommonResourceInterface
	Get() *ResourceAppProject
}

// Certificateinterface represents the interface for resources of the type certificate
type Certificateinterface interface {
	CommonResourceInterface
	Get() *ResourceCertificate
}

// Serviceinterface represents the interface for resources of the type service
type Serviceinterface interface {
	CommonResourceInterface
	Get() *ResourceService
}

// Podinterface represents the interface for resources of the type pod
type Podinterface interface {
	CommonResourceInterface
	Get() *ResourcePod
}

// ReplicaSetinterface represents the interface for resources of the type replicaset
type ReplicaSetinterface interface {
	CommonResourceInterface
	Get() *ResourceReplicaSet
}

// StatefulSetinterface represents the interface for resources of the type statefulset
type StatefulSetinterface interface {
	CommonResourceInterface
	Get() *ResourceStatefulSet
}

// DaemonSetinterface represents the interface for resources of the type daemonset
type DaemonSetinterface interface {
	CommonResourceInterface
	Get() *ResourceDaemonSet
}

// Ingressinterface represents the interface for resources of the type ingress
type Ingressinterface interface {
	CommonResourceInterface
	Get() *ResourceIngress
}

// IngressClassinterface represents the interface for resources of the type ingressclass
type IngressClassinterface interface {
	CommonResourceInterface
	Get() *ResourceIngressClass
}

// VulnerabilityReportinterface represents the interface for resources of the type vulnerabilityreport
type VulnerabilityReportinterface interface {
	CommonResourceInterface
	Get() *ResourceVulnerabilityReport
}

// ExposedSecretReportinterface represents the interface for resources of the type exposedsecretreport
type ExposedSecretReportinterface interface {
	CommonResourceInterface
	Get() *ResourceExposedSecretReport
}

// ConfigAuditReportinterface represents the interface for resources of the type configauditreport
type ConfigAuditReportinterface interface {
	CommonResourceInterface
	Get() *ResourceConfigAuditReport
}

// RbacAssessmentReportinterface represents the interface for resources of the type rbacassessmentreport
type RbacAssessmentReportinterface interface {
	CommonResourceInterface
	Get() *ResourceRbacAssessmentReport
}

// TanzuKubernetesClusterinterface represents the interface for resources of the type tanzukubernetescluster
type TanzuKubernetesClusterinterface interface {
	CommonResourceInterface
	Get() *ResourceTanzuKubernetesCluster
}

// TanzuKubernetesReleaseinterface represents the interface for resources of the type tanzukubernetesrelease
type TanzuKubernetesReleaseinterface interface {
	CommonResourceInterface
	Get() *ResourceTanzuKubernetesRelease
}

// VirtualMachineClassinterface represents the interface for resources of the type virtualmachineclass
type VirtualMachineClassinterface interface {
	CommonResourceInterface
	Get() *ResourceVirtualMachineClass
}

// VirtualMachineClassBindinginterface represents the interface for resources of the type virtualmachineclassbinding
type VirtualMachineClassBindinginterface interface {
	CommonResourceInterface
	Get() *ResourceVirtualMachineClassBinding
}

// KubernetesClusterinterface represents the interface for resources of the type kubernetescluster
type KubernetesClusterinterface interface {
	CommonResourceInterface
	Get() *ResourceKubernetesCluster
}

// ClusterOrderinterface represents the interface for resources of the type clusterorder
type ClusterOrderinterface interface {
	CommonResourceInterface
	Get() *ResourceClusterOrder
}

// Projectinterface represents the interface for resources of the type project
type Projectinterface interface {
	CommonResourceInterface
	Get() *ResourceProject
}

// Configurationinterface represents the interface for resources of the type configuration
type Configurationinterface interface {
	CommonResourceInterface
	Get() *ResourceConfiguration
}

// ClusterComplianceReportinterface represents the interface for resources of the type clustercompliancereport
type ClusterComplianceReportinterface interface {
	CommonResourceInterface
	Get() *ResourceClusterComplianceReport
}

// ClusterVulnerabilityReportinterface represents the interface for resources of the type clustervulnerabilityreport
type ClusterVulnerabilityReportinterface interface {
	CommonResourceInterface
	Get() *ResourceClusterVulnerabilityReport
}

// Routeinterface represents the interface for resources of the type route
type Routeinterface interface {
	CommonResourceInterface
	Get() *ResourceRoute
}

// SlackMessageinterface represents the interface for resources of the type slackmessage
type SlackMessageinterface interface {
	CommonResourceInterface
	Get() *ResourceSlackMessage
}

// VulnerabilityEventinterface represents the interface for resources of the type vulnerabilityevent
type VulnerabilityEventinterface interface {
	CommonResourceInterface
	Get() *ResourceVulnerabilityEvent
}

// VirtualMachineinterface represents the interface for resources of the type virtualmachine
type VirtualMachineinterface interface {
	CommonResourceInterface
	Get() *ResourceVirtualMachine
}

// Endpointsinterface represents the interface for resources of the type endpoints
type Endpointsinterface interface {
	CommonResourceInterface
	Get() *ResourceEndpoints
}

// NetworkPolicyinterface represents the interface for resources of the type networkpolicy
type NetworkPolicyinterface interface {
	CommonResourceInterface
	Get() *ResourceNetworkPolicy
}

// FirewallRuleinterface represents the interface for resources of the type firewallrule
type FirewallRuleinterface interface {
	CommonResourceInterface
	Get() *ResourceFirewallRule
}

// VirtualMachineinterface represents the interface for resources of the type virtualmachine
type VirtualMachineinterface interface {
	CommonResourceInterface
	Get() *ResourceVirtualMachine
}

// FirewallPolicyinterface represents the interface for resources of the type firewallpolicy
type FirewallPolicyinterface interface {
	CommonResourceInterface
	Get() *ResourceFirewallPolicy
}
