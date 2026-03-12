// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package rortypes

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	ResourceNamespaceGVK = schema.GroupVersionKind{
		Group:   "",
		Version: "v1",
		Kind:    "Namespace",
	}

	ResourceNodeGVK = schema.GroupVersionKind{
		Group:   "",
		Version: "v1",
		Kind:    "Node",
	}

	ResourcePersistentVolumeClaimGVK = schema.GroupVersionKind{
		Group:   "",
		Version: "v1",
		Kind:    "PersistentVolumeClaim",
	}

	ResourceDeploymentGVK = schema.GroupVersionKind{
		Group:   "apps",
		Version: "v1",
		Kind:    "Deployment",
	}

	ResourceStorageClassGVK = schema.GroupVersionKind{
		Group:   "storage.k8s.io",
		Version: "v1",
		Kind:    "StorageClass",
	}

	ResourcePolicyReportGVK = schema.GroupVersionKind{
		Group:   "wgpolicyk8s.io",
		Version: "v1alpha2",
		Kind:    "PolicyReport",
	}

	ResourceApplicationGVK = schema.GroupVersionKind{
		Group:   "argoproj.io",
		Version: "v1alpha1",
		Kind:    "Application",
	}

	ResourceAppProjectGVK = schema.GroupVersionKind{
		Group:   "argoproj.io",
		Version: "v1alpha1",
		Kind:    "AppProject",
	}

	ResourceCertificateGVK = schema.GroupVersionKind{
		Group:   "cert-manager.io",
		Version: "v1",
		Kind:    "Certificate",
	}

	ResourceServiceGVK = schema.GroupVersionKind{
		Group:   "",
		Version: "v1",
		Kind:    "Service",
	}

	ResourcePodGVK = schema.GroupVersionKind{
		Group:   "",
		Version: "v1",
		Kind:    "Pod",
	}

	ResourceReplicaSetGVK = schema.GroupVersionKind{
		Group:   "apps",
		Version: "v1",
		Kind:    "ReplicaSet",
	}

	ResourceStatefulSetGVK = schema.GroupVersionKind{
		Group:   "apps",
		Version: "v1",
		Kind:    "StatefulSet",
	}

	ResourceDaemonSetGVK = schema.GroupVersionKind{
		Group:   "apps",
		Version: "v1",
		Kind:    "DaemonSet",
	}

	ResourceIngressGVK = schema.GroupVersionKind{
		Group:   "networking.k8s.io",
		Version: "v1",
		Kind:    "Ingress",
	}

	ResourceIngressClassGVK = schema.GroupVersionKind{
		Group:   "networking.k8s.io",
		Version: "v1",
		Kind:    "IngressClass",
	}

	ResourceSbomReportGVK = schema.GroupVersionKind{
		Group:   "aquasecurity.github.io",
		Version: "v1alpha1",
		Kind:    "SbomReport",
	}

	ResourceVulnerabilityReportGVK = schema.GroupVersionKind{
		Group:   "aquasecurity.github.io",
		Version: "v1alpha1",
		Kind:    "VulnerabilityReport",
	}

	ResourceExposedSecretReportGVK = schema.GroupVersionKind{
		Group:   "aquasecurity.github.io",
		Version: "v1alpha1",
		Kind:    "ExposedSecretReport",
	}

	ResourceConfigAuditReportGVK = schema.GroupVersionKind{
		Group:   "aquasecurity.github.io",
		Version: "v1alpha1",
		Kind:    "ConfigAuditReport",
	}

	ResourceRbacAssessmentReportGVK = schema.GroupVersionKind{
		Group:   "aquasecurity.github.io",
		Version: "v1alpha1",
		Kind:    "RbacAssessmentReport",
	}

	ResourceTanzuKubernetesClusterGVK = schema.GroupVersionKind{
		Group:   "run.tanzu.vmware.com",
		Version: "v1alpha3",
		Kind:    "TanzuKubernetesCluster",
	}

	ResourceTanzuKubernetesReleaseGVK = schema.GroupVersionKind{
		Group:   "run.tanzu.vmware.com",
		Version: "v1alpha3",
		Kind:    "TanzuKubernetesRelease",
	}

	ResourceVirtualMachineClassGVK = schema.GroupVersionKind{
		Group:   "vmoperator.vmware.com",
		Version: "v1alpha2",
		Kind:    "VirtualMachineClass",
	}

	ResourceKubernetesClusterGVK = schema.GroupVersionKind{
		Group:   "vitistack.io",
		Version: "v1alpha1",
		Kind:    "KubernetesCluster",
	}

	ResourceProviderGVK = schema.GroupVersionKind{
		Group:   "general.ror.internal",
		Version: "v1alpha1",
		Kind:    "Provider",
	}

	ResourceWorkspaceGVK = schema.GroupVersionKind{
		Group:   "general.ror.internal",
		Version: "v1alpha1",
		Kind:    "Workspace",
	}

	ResourceKubernetesMachineClassGVK = schema.GroupVersionKind{
		Group:   "general.ror.internal",
		Version: "v1alpha1",
		Kind:    "KubernetesMachineClass",
	}

	ResourceClusterOrderGVK = schema.GroupVersionKind{
		Group:   "general.ror.internal",
		Version: "v1alpha1",
		Kind:    "ClusterOrder",
	}

	ResourceProjectGVK = schema.GroupVersionKind{
		Group:   "general.ror.internal",
		Version: "v1alpha1",
		Kind:    "Project",
	}

	ResourceConfigurationGVK = schema.GroupVersionKind{
		Group:   "general.ror.internal",
		Version: "v1alpha1",
		Kind:    "Configuration",
	}

	ResourceClusterComplianceReportGVK = schema.GroupVersionKind{
		Group:   "aquasecurity.github.io",
		Version: "v1alpha1",
		Kind:    "ClusterComplianceReport",
	}

	ResourceClusterVulnerabilityReportGVK = schema.GroupVersionKind{
		Group:   "general.ror.internal",
		Version: "v1alpha1",
		Kind:    "ClusterVulnerabilityReport",
	}

	ResourceRouteGVK = schema.GroupVersionKind{
		Group:   "general.ror.internal",
		Version: "v1alpha1",
		Kind:    "Route",
	}

	ResourceSlackMessageGVK = schema.GroupVersionKind{
		Group:   "general.ror.internal",
		Version: "v1alpha1",
		Kind:    "SlackMessage",
	}

	ResourceVulnerabilityEventGVK = schema.GroupVersionKind{
		Group:   "general.ror.internal",
		Version: "v1alpha1",
		Kind:    "VulnerabilityEvent",
	}

	ResourceVirtualMachineGVK = schema.GroupVersionKind{
		Group:   "general.ror.internal",
		Version: "v1alpha1",
		Kind:    "VirtualMachine",
	}

	ResourceVirtualMachineVulnerabilityInfoGVK = schema.GroupVersionKind{
		Group:   "general.ror.internal",
		Version: "v1alpha1",
		Kind:    "VirtualMachineVulnerabilityInfo",
	}

	ResourceEndpointsGVK = schema.GroupVersionKind{
		Group:   "",
		Version: "v1",
		Kind:    "Endpoints",
	}

	ResourceNetworkPolicyGVK = schema.GroupVersionKind{
		Group:   "networking.k8s.io",
		Version: "v1",
		Kind:    "NetworkPolicy",
	}

	ResourceDatacenterGVK = schema.GroupVersionKind{
		Group:   "infrastructure.ror.internal",
		Version: "v1alpha1",
		Kind:    "Datacenter",
	}

	ResourceBackupJobGVK = schema.GroupVersionKind{
		Group:   "backup.ror.internal",
		Version: "v1alpha1",
		Kind:    "BackupJob",
	}

	ResourceBackupRunGVK = schema.GroupVersionKind{
		Group:   "backup.ror.internal",
		Version: "v1alpha1",
		Kind:    "BackupRun",
	}

	ResourceUnknownGVK = schema.GroupVersionKind{
		Group:   "unknown.ror.internal",
		Version: "v1",
		Kind:    "Unknown",
	}
)
