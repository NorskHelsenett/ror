// The package provides the models and variables needed to generate code and endpoints for the implemented rorresources
package rordefs // Package resourcegeneratormodels

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Individual resource definitions.
// Each is a named variable so it can be referenced directly elsewhere
// (e.g. ResourceConfiguration.Kind in ACL filters).
//
// When changed the generator must be run, and the files generated checked in with the code.
//
//	$ go run build/generator/main.go

var ResourceNamespace = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Namespace",
		APIVersion: "v1",
	},
	Plural:     "namespaces",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceNode = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Node",
		APIVersion: "v1",
	},
	Plural:     "nodes",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourcePersistentVolumeClaim = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "PersistentVolumeClaim",
		APIVersion: "v1",
	},
	Plural:     "persistentvolumeclaims",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceDeployment = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Deployment",
		APIVersion: "apps/v1",
	},
	Plural:     "deployments",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceStorageClass = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "StorageClass",
		APIVersion: "storage.k8s.io/v1",
	},
	Plural:     "storageclasses",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourcePolicyReport = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "PolicyReport",
		APIVersion: "wgpolicyk8s.io/v1alpha2",
	},
	Plural:     "policyreports",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceApplication = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Application",
		APIVersion: "argoproj.io/v1alpha1",
	},
	Plural:     "applications",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceAppProject = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "AppProject",
		APIVersion: "argoproj.io/v1alpha1",
	},
	Plural:     "appprojects",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceCertificate = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Certificate",
		APIVersion: "cert-manager.io/v1",
	},
	Plural:     "certificates",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceService = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Service",
		APIVersion: "v1",
	},
	Plural:     "services",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourcePod = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Pod",
		APIVersion: "v1",
	},
	Plural:     "pods",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceReplicaSet = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "ReplicaSet",
		APIVersion: "apps/v1",
	},
	Plural:     "replicasets",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceStatefulSet = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "StatefulSet",
		APIVersion: "apps/v1",
	},
	Plural:     "statefulsets",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceDaemonSet = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "DaemonSet",
		APIVersion: "apps/v1",
	},
	Plural:     "daemonsets",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceIngress = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Ingress",
		APIVersion: "networking.k8s.io/v1",
	},
	Plural:     "ingresses",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceIngressClass = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "IngressClass",
		APIVersion: "networking.k8s.io/v1",
	},
	Plural:     "ingressclasses",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceSbomReport = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "SbomReport",
		APIVersion: "aquasecurity.github.io/v1alpha1",
	},
	Plural:     "sbomreports",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeInternal},
	Versions:   []ApiVersions{ApiVersionV2},
}

var ResourceVulnerabilityReport = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "VulnerabilityReport",
		APIVersion: "aquasecurity.github.io/v1alpha1",
	},
	Plural:     "vulnerabilityreports",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceExposedSecretReport = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "ExposedSecretReport",
		APIVersion: "aquasecurity.github.io/v1alpha1",
	},
	Plural:     "exposedsecretreports",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceConfigAuditReport = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "ConfigAuditReport",
		APIVersion: "aquasecurity.github.io/v1alpha1",
	},
	Plural:     "configauditreports",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceRbacAssessmentReport = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "RbacAssessmentReport",
		APIVersion: "aquasecurity.github.io/v1alpha1",
	},
	Plural:     "rbacassessmentreports",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceTanzuKubernetesCluster = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "TanzuKubernetesCluster",
		APIVersion: "run.tanzu.vmware.com/v1alpha3",
	},
	Plural:     "tanzukubernetesclusters",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeTanzuAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceTanzuKubernetesRelease = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "TanzuKubernetesRelease",
		APIVersion: "run.tanzu.vmware.com/v1alpha3",
	},
	Plural:     "tanzukubernetesreleases",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeTanzuAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceVirtualMachineClass = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "VirtualMachineClass",
		APIVersion: "vmoperator.vmware.com/v1alpha2",
	},
	Plural:     "virtualmachineclasses",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeTanzuAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceKubernetesCluster = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "KubernetesCluster",
		APIVersion: "general.ror.internal/v1alpha1",
	},
	Plural:     "kubernetesclusters",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeInternal, ApiResourceTypeTanzuAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceProvider = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Provider",
		APIVersion: "general.ror.internal/v1alpha1",
	},
	Plural:     "providers",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeInternal, ApiResourceTypeTanzuAgent},
	Versions:   []ApiVersions{ApiVersionV2},
}

var ResourceWorkspace = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Workspace",
		APIVersion: "general.ror.internal/v1alpha1",
	},
	Plural:     "workspaces",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeInternal, ApiResourceTypeTanzuAgent},
	Versions:   []ApiVersions{ApiVersionV2},
}

var ResourceKubernetesMachineClass = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "KubernetesMachineClass",
		APIVersion: "general.ror.internal/v1alpha1",
	},
	Plural:     "Kubernetesmachineclasses",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeInternal, ApiResourceTypeTanzuAgent},
	Versions:   []ApiVersions{ApiVersionV2},
}

var ResourceClusterOrder = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "ClusterOrder",
		APIVersion: "general.ror.internal/v1alpha1",
	},
	Plural:     "clusterorders",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeInternal},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceProject = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Project",
		APIVersion: "general.ror.internal/v1alpha1",
	},
	Plural:     "projects",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeInternal},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceConfiguration = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Configuration",
		APIVersion: "general.ror.internal/v1alpha1",
	},
	Plural:     "configurations",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeInternal},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceClusterComplianceReport = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "ClusterComplianceReport",
		APIVersion: "aquasecurity.github.io/v1alpha1",
	},
	Plural:     "clustercompliancereports",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceClusterVulnerabilityReport = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "ClusterVulnerabilityReport",
		APIVersion: "general.ror.internal/v1alpha1",
	},
	Plural:     "clustervulnerabilityreports",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeInternal},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceRoute = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Route",
		APIVersion: "general.ror.internal/v1alpha1",
	},
	Plural:     "routes",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeInternal},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceSlackMessage = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "SlackMessage",
		APIVersion: "general.ror.internal/v1alpha1",
	},
	Plural:     "slackmessages",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeInternal},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceVulnerabilityEvent = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "VulnerabilityEvent",
		APIVersion: "general.ror.internal/v1alpha1",
	},
	Plural:     "vulnerabilityevents",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeInternal},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceVirtualMachine = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "VirtualMachine",
		APIVersion: "general.ror.internal/v1alpha1",
	},
	Plural:     "VirtualMachines",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeVmAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceVirtualMachineVulnerabilityInfo = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "VirtualMachineVulnerabilityInfo",
		APIVersion: "general.ror.internal/v1alpha1",
	},
	Plural:     "VirtualMachinesVulnerabilityInfo",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeVmAgent},
	Versions:   []ApiVersions{ApiVersionV2},
}

var ResourceEndpoints = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Endpoints",
		APIVersion: "v1",
	},
	Plural:     "endpoints",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceNetworkPolicy = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "NetworkPolicy",
		APIVersion: "networking.k8s.io/v1",
	},
	Plural:     "networkpolicies",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeAgent},
	Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
}

var ResourceDatacenter = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Datacenter",
		APIVersion: "infrastructure.ror.internal/v1alpha1",
	},
	Plural:     "datacenters",
	Namespaced: true,
	Types:      []ApiResourceType{ApiResourceTypeInternal, ApiResourceTypeTanzuAgent},
	Versions:   []ApiVersions{ApiVersionV2},
}

var ResourceBackupJob = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "BackupJob",
		APIVersion: "backup.ror.internal/v1alpha1",
	},
	Plural:     "backupjobs",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeBackupAgent},
	Versions:   []ApiVersions{ApiVersionV2},
}

var ResourceBackupRun = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "BackupRun",
		APIVersion: "backup.ror.internal/v1alpha1",
	},
	Plural:     "backupruns",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeBackupAgent},
	Versions:   []ApiVersions{ApiVersionV2},
}

var ResourceMachine = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Machine",
		APIVersion: "machine.ror.internal/v1alpha1",
	},
	Plural:     "machines",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeMachineAgent},
	Versions:   []ApiVersions{ApiVersionV2},
}

var ResourceConfig = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Config",
		APIVersion: "ror.internal/v1",
	},
	Plural:     "configs",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeInternal},
	Versions:   []ApiVersions{ApiVersionV2},
}

var ResourceOrganizationalUnit = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "OrganizationalUnit",
		APIVersion: "ror.internal/v1",
	},
	Plural:     "OrganizationalUnits",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeInternal},
	Versions:   []ApiVersions{ApiVersionV2},
}

var ResourceUnknown = ApiResource{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Unknown",
		APIVersion: "unknown.ror.internal/v1",
	},
	Plural:     "unknowns",
	Namespaced: false,
	Types:      []ApiResourceType{ApiResourceTypeInternal},
	Versions:   []ApiVersions{ApiVersionV2},
}

// Resourcedefs is the complete registry of all resources implemented in ror.
// It is composed from the individual Resource* variables above.
var Resourcedefs = ApiResources{
	ResourceNamespace,
	ResourceNode,
	ResourcePersistentVolumeClaim,
	ResourceDeployment,
	ResourceStorageClass,
	ResourcePolicyReport,
	ResourceApplication,
	ResourceAppProject,
	ResourceCertificate,
	ResourceService,
	ResourcePod,
	ResourceReplicaSet,
	ResourceStatefulSet,
	ResourceDaemonSet,
	ResourceIngress,
	ResourceIngressClass,
	ResourceSbomReport,
	ResourceVulnerabilityReport,
	ResourceExposedSecretReport,
	ResourceConfigAuditReport,
	ResourceRbacAssessmentReport,
	ResourceTanzuKubernetesCluster,
	ResourceTanzuKubernetesRelease,
	ResourceVirtualMachineClass,
	ResourceKubernetesCluster,
	ResourceProvider,
	ResourceWorkspace,
	ResourceKubernetesMachineClass,
	ResourceClusterOrder,
	ResourceProject,
	ResourceConfiguration,
	ResourceClusterComplianceReport,
	ResourceClusterVulnerabilityReport,
	ResourceRoute,
	ResourceSlackMessage,
	ResourceVulnerabilityEvent,
	ResourceVirtualMachine,
	ResourceVirtualMachineVulnerabilityInfo,
	ResourceEndpoints,
	ResourceNetworkPolicy,
	ResourceDatacenter,
	ResourceBackupJob,
	ResourceBackupRun,
	ResourceMachine,
	ResourceConfig,
	ResourceOrganizationalUnit,
	ResourceUnknown,
}
