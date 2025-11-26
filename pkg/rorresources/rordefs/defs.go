// The package provides the models and variables needed to generate code and endpoints for the implemented rorresources
package rordefs // Package resourcegeneratormodels

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Resources implemented in ror
//
// When changed the generator must be run, and the files generated checked in with the code.
//
//	$ go run build/generator/main.go
var Resourcedefs = ApiResources{
	{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		Plural:     "namespaces",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		Plural:     "nodes",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "PersistentVolumeClaim",
			APIVersion: "v1",
		},
		Plural:     "persistentvolumeclaims",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		Plural:     "deployments",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "StorageClass",
			APIVersion: "storage.k8s.io/v1",
		},
		Plural:     "storageclasses",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "PolicyReport",
			APIVersion: "wgpolicyk8s.io/v1alpha2",
		},
		Plural:     "policyreports",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "Application",
			APIVersion: "argoproj.io/v1alpha1",
		},
		Plural:     "applications",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "AppProject",
			APIVersion: "argoproj.io/v1alpha1",
		},
		Plural:     "appprojects",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "Certificate",
			APIVersion: "cert-manager.io/v1",
		},
		Plural:     "certificates",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		Plural:     "services",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		Plural:     "pods",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "ReplicaSet",
			APIVersion: "apps/v1",
		},
		Plural:     "replicasets",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		Plural:     "statefulsets",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "DaemonSet",
			APIVersion: "apps/v1",
		},
		Plural:     "daemonsets",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "networking.k8s.io/v1",
		},
		Plural:     "ingresses",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "IngressClass",
			APIVersion: "networking.k8s.io/v1",
		},
		Plural:     "ingressclasses",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "VulnerabilityReport",
			APIVersion: "aquasecurity.github.io/v1alpha1",
		},
		Plural:     "vulnerabilityreports",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "ExposedSecretReport",
			APIVersion: "aquasecurity.github.io/v1alpha1",
		},
		Plural:     "exposedsecretreports",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigAuditReport",
			APIVersion: "aquasecurity.github.io/v1alpha1",
		},
		Plural:     "configauditreports",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "RbacAssessmentReport",
			APIVersion: "aquasecurity.github.io/v1alpha1",
		},
		Plural:     "rbacassessmentreports",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "TanzuKubernetesCluster",
			APIVersion: "run.tanzu.vmware.com/v1alpha3",
		},
		Plural:     "tanzukubernetesclusters",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeTanzuAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "TanzuKubernetesRelease",
			APIVersion: "run.tanzu.vmware.com/v1alpha3",
		},
		Plural:     "tanzukubernetesreleases",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeTanzuAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "VirtualMachineClass",
			APIVersion: "vmoperator.vmware.com/v1alpha2",
		},
		Plural:     "virtualmachineclasses",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeTanzuAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "KubernetesCluster",
			APIVersion: "vitistack.io/v1alpha1",
		},
		Plural:     "kubernetesclusters",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeInternal, ApiResourceTypeTanzuAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	},
	{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Provider",
			APIVersion: "general.ror.internal/v1alpha1",
		},
		Plural:     "providers",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeInternal, ApiResourceTypeTanzuAgent},
		Versions:   []ApiVersions{ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "Workspace",
			APIVersion: "general.ror.internal/v1alpha1",
		},
		Plural:     "workspaces",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeInternal, ApiResourceTypeTanzuAgent},
		Versions:   []ApiVersions{ApiVersionV2},
	},
	{
		TypeMeta: metav1.TypeMeta{
			Kind:       "KubernetesMachineClass",
			APIVersion: "general.ror.internal/v1alpha1",
		},
		Plural:     "Kubernetesmachineclasses",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeInternal, ApiResourceTypeTanzuAgent},
		Versions:   []ApiVersions{ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterOrder",
			APIVersion: "general.ror.internal/v1alpha1",
		},
		Plural:     "clusterorders",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeInternal},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "Project",
			APIVersion: "general.ror.internal/v1alpha1",
		},
		Plural:     "projects",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeInternal},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "Configuration",
			APIVersion: "general.ror.internal/v1alpha1",
		},
		Plural:     "configurations",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeInternal},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterComplianceReport",
			APIVersion: "aquasecurity.github.io/v1alpha1",
		},
		Plural:     "clustercompliancereports",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterVulnerabilityReport",
			APIVersion: "general.ror.internal/v1alpha1",
		},
		Plural:     "clustervulnerabilityreports",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeInternal},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "Route",
			APIVersion: "general.ror.internal/v1alpha1",
		},
		Plural:     "routes",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeInternal},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "SlackMessage",
			APIVersion: "general.ror.internal/v1alpha1",
		},
		Plural:     "slackmessages",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeInternal},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "VulnerabilityEvent",
			APIVersion: "general.ror.internal/v1alpha1",
		},
		Plural:     "vulnerabilityevents",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeInternal},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "VirtualMachine",
			APIVersion: "general.ror.internal/v1alpha1",
		},
		Plural:     "VirtualMachines",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeVmAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	},
	{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Endpoints",
			APIVersion: "v1",
		},
		Plural:     "endpoints",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "NetworkPolicy",
			APIVersion: "networking.k8s.io/v1",
		},
		Plural:     "networkpolicies",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeAgent},
		Versions:   []ApiVersions{ApiVersionV1, ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "Datacenter",
			APIVersion: "infrastructure.ror.internal/v1alpha1",
		},
		Plural:     "datacenters",
		Namespaced: true,
		Types:      []ApiResourceType{ApiResourceTypeInternal, ApiResourceTypeTanzuAgent},
		Versions:   []ApiVersions{ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "BackupJob",
			APIVersion: "backup.ror.internal/v1alpha1",
		},
		Plural:     "backupjobs",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeBackupAgent},
		Versions:   []ApiVersions{ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "BackupRun",
			APIVersion: "backup.ror.internal/v1alpha1",
		},
		Plural:     "backupruns",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeBackupAgent},
		Versions:   []ApiVersions{ApiVersionV2},
	}, {
		TypeMeta: metav1.TypeMeta{
			Kind:       "Unknown",
			APIVersion: "unknown.ror.internal/v1",
		},
		Plural:     "unknowns",
		Namespaced: false,
		Types:      []ApiResourceType{ApiResourceTypeInternal},
		Versions:   []ApiVersions{ApiVersionV2},
	},
}
