// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package resourcesservice

import (
	"context"
	"errors"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/resourcesmongodbrepo"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// Functions to get Namespaces by uid,ownerref
// The function is intended for use by internal functions
func GetNamespaceByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceNamespace, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceNamespace{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Namespace",
		ApiVersion: "v1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceNamespace](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceNamespace{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Nodes by uid,ownerref
// The function is intended for use by internal functions
func GetNodeByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceNode, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceNode{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Node",
		ApiVersion: "v1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceNode](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceNode{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Persistentvolumeclaims by uid,ownerref
// The function is intended for use by internal functions
func GetPersistentVolumeClaimByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourcePersistentVolumeClaim, error) {
	if uid == "" {
		return apiresourcecontracts.ResourcePersistentVolumeClaim{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "PersistentVolumeClaim",
		ApiVersion: "v1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourcePersistentVolumeClaim](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourcePersistentVolumeClaim{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Deployments by uid,ownerref
// The function is intended for use by internal functions
func GetDeploymentByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceDeployment, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceDeployment{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Deployment",
		ApiVersion: "apps/v1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceDeployment](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceDeployment{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Storageclasses by uid,ownerref
// The function is intended for use by internal functions
func GetStorageClassByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceStorageClass, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceStorageClass{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "StorageClass",
		ApiVersion: "storage.k8s.io/v1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceStorageClass](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceStorageClass{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Policyreports by uid,ownerref
// The function is intended for use by internal functions
func GetPolicyReportByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourcePolicyReport, error) {
	if uid == "" {
		return apiresourcecontracts.ResourcePolicyReport{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "PolicyReport",
		ApiVersion: "wgpolicyk8s.io/v1alpha2",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourcePolicyReport](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourcePolicyReport{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Applications by uid,ownerref
// The function is intended for use by internal functions
func GetApplicationByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceApplication, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceApplication{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Application",
		ApiVersion: "argoproj.io/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceApplication](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceApplication{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Appprojects by uid,ownerref
// The function is intended for use by internal functions
func GetAppProjectByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceAppProject, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceAppProject{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "AppProject",
		ApiVersion: "argoproj.io/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceAppProject](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceAppProject{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Certificates by uid,ownerref
// The function is intended for use by internal functions
func GetCertificateByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceCertificate, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceCertificate{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Certificate",
		ApiVersion: "cert-manager.io/v1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceCertificate](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceCertificate{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Services by uid,ownerref
// The function is intended for use by internal functions
func GetServiceByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceService, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceService{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Service",
		ApiVersion: "v1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceService](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceService{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Pods by uid,ownerref
// The function is intended for use by internal functions
func GetPodByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourcePod, error) {
	if uid == "" {
		return apiresourcecontracts.ResourcePod{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Pod",
		ApiVersion: "v1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourcePod](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourcePod{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Replicasets by uid,ownerref
// The function is intended for use by internal functions
func GetReplicaSetByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceReplicaSet, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceReplicaSet{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "ReplicaSet",
		ApiVersion: "apps/v1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceReplicaSet](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceReplicaSet{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Statefulsets by uid,ownerref
// The function is intended for use by internal functions
func GetStatefulSetByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceStatefulSet, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceStatefulSet{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "StatefulSet",
		ApiVersion: "apps/v1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceStatefulSet](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceStatefulSet{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Daemonsets by uid,ownerref
// The function is intended for use by internal functions
func GetDaemonSetByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceDaemonSet, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceDaemonSet{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "DaemonSet",
		ApiVersion: "apps/v1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceDaemonSet](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceDaemonSet{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Ingresses by uid,ownerref
// The function is intended for use by internal functions
func GetIngressByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceIngress, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceIngress{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Ingress",
		ApiVersion: "networking.k8s.io/v1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceIngress](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceIngress{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Ingressclasses by uid,ownerref
// The function is intended for use by internal functions
func GetIngressClassByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceIngressClass, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceIngressClass{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "IngressClass",
		ApiVersion: "networking.k8s.io/v1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceIngressClass](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceIngressClass{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Vulnerabilityreports by uid,ownerref
// The function is intended for use by internal functions
func GetVulnerabilityReportByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceVulnerabilityReport, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceVulnerabilityReport{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "VulnerabilityReport",
		ApiVersion: "aquasecurity.github.io/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceVulnerabilityReport](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceVulnerabilityReport{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Exposedsecretreports by uid,ownerref
// The function is intended for use by internal functions
func GetExposedSecretReportByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceExposedSecretReport, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceExposedSecretReport{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "ExposedSecretReport",
		ApiVersion: "aquasecurity.github.io/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceExposedSecretReport](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceExposedSecretReport{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Configauditreports by uid,ownerref
// The function is intended for use by internal functions
func GetConfigAuditReportByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceConfigAuditReport, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceConfigAuditReport{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "ConfigAuditReport",
		ApiVersion: "aquasecurity.github.io/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceConfigAuditReport](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceConfigAuditReport{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Rbacassessmentreports by uid,ownerref
// The function is intended for use by internal functions
func GetRbacAssessmentReportByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceRbacAssessmentReport, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceRbacAssessmentReport{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "RbacAssessmentReport",
		ApiVersion: "aquasecurity.github.io/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceRbacAssessmentReport](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceRbacAssessmentReport{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Tanzukubernetesclusters by uid,ownerref
// The function is intended for use by internal functions
func GetTanzuKubernetesClusterByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceTanzuKubernetesCluster, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceTanzuKubernetesCluster{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "TanzuKubernetesCluster",
		ApiVersion: "run.tanzu.vmware.com/v1alpha2",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceTanzuKubernetesCluster](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceTanzuKubernetesCluster{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Tanzukubernetesreleases by uid,ownerref
// The function is intended for use by internal functions
func GetTanzuKubernetesReleaseByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceTanzuKubernetesRelease, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceTanzuKubernetesRelease{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "TanzuKubernetesRelease",
		ApiVersion: "run.tanzu.vmware.com/v1alpha2",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceTanzuKubernetesRelease](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceTanzuKubernetesRelease{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Virtualmachineclasses by uid,ownerref
// The function is intended for use by internal functions
func GetVirtualMachineClassByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceVirtualMachineClass, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceVirtualMachineClass{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "VirtualMachineClass",
		ApiVersion: "vmoperator.vmware.com/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceVirtualMachineClass](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceVirtualMachineClass{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Virtualmachineclassbindings by uid,ownerref
// The function is intended for use by internal functions
func GetVirtualMachineClassBindingByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceVirtualMachineClassBinding, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceVirtualMachineClassBinding{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "VirtualMachineClassBinding",
		ApiVersion: "vmoperator.vmware.com/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceVirtualMachineClassBinding](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceVirtualMachineClassBinding{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Kubernetesclusters by uid,ownerref
// The function is intended for use by internal functions
func GetKubernetesClusterByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceKubernetesCluster, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceKubernetesCluster{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "KubernetesCluster",
		ApiVersion: "general.ror.internal/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceKubernetesCluster](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceKubernetesCluster{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Clusterorders by uid,ownerref
// The function is intended for use by internal functions
func GetClusterOrderByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceClusterOrder, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceClusterOrder{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "ClusterOrder",
		ApiVersion: "general.ror.internal/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceClusterOrder](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceClusterOrder{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Projects by uid,ownerref
// The function is intended for use by internal functions
func GetProjectByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceProject, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceProject{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Project",
		ApiVersion: "general.ror.internal/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceProject](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceProject{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Configurations by uid,ownerref
// The function is intended for use by internal functions
func GetConfigurationByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceConfiguration, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceConfiguration{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Configuration",
		ApiVersion: "general.ror.internal/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceConfiguration](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceConfiguration{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Clustercompliancereports by uid,ownerref
// The function is intended for use by internal functions
func GetClusterComplianceReportByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceClusterComplianceReport, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceClusterComplianceReport{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "ClusterComplianceReport",
		ApiVersion: "aquasecurity.github.io/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceClusterComplianceReport](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceClusterComplianceReport{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Clustervulnerabilityreports by uid,ownerref
// The function is intended for use by internal functions
func GetClusterVulnerabilityReportByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceClusterVulnerabilityReport, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceClusterVulnerabilityReport{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "ClusterVulnerabilityReport",
		ApiVersion: "general.ror.internal/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceClusterVulnerabilityReport](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceClusterVulnerabilityReport{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Routes by uid,ownerref
// The function is intended for use by internal functions
func GetRouteByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceRoute, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceRoute{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Route",
		ApiVersion: "general.ror.internal/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceRoute](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceRoute{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Slackmessages by uid,ownerref
// The function is intended for use by internal functions
func GetSlackMessageByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceSlackMessage, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceSlackMessage{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "SlackMessage",
		ApiVersion: "general.ror.internal/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceSlackMessage](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceSlackMessage{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Vulnerabilityevents by uid,ownerref
// The function is intended for use by internal functions
func GetVulnerabilityEventByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceVulnerabilityEvent, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceVulnerabilityEvent{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "VulnerabilityEvent",
		ApiVersion: "general.ror.internal/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceVulnerabilityEvent](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceVulnerabilityEvent{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Vms by uid,ownerref
// The function is intended for use by internal functions
func GetVmByUid(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference, uid string) (apiresourcecontracts.ResourceVm, error) {
	if uid == "" {
		return apiresourcecontracts.ResourceVm{}, errors.New("uid is empty")
	}
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Vm",
		ApiVersion: "general.ror.internal/v1alpha1",
		Internal:   true,
		Uid:        uid,
	}

	resource, err := GetResource[apiresourcecontracts.ResourceVm](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceVm{}, errors.New("could not get resource")
	}

	return resource, nil

}

// Functions to get Namespaces by ownerref
// The function is intended for use by internal functions
func GetNamespaces(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceNamespaces, error) {
	var resources apiresourcecontracts.ResourceNamespaces
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Namespace",
		ApiVersion: "v1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceNamespace](ctx, query)
	resources.Owner = ownerref
	resources.Namespaces = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource Namespace")
	}
	return resources, nil
}

// Functions to get Nodes by ownerref
// The function is intended for use by internal functions
func GetNodes(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceNodes, error) {
	var resources apiresourcecontracts.ResourceNodes
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Node",
		ApiVersion: "v1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceNode](ctx, query)
	resources.Owner = ownerref
	resources.Nodes = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource Node")
	}
	return resources, nil
}

// Functions to get Persistentvolumeclaims by ownerref
// The function is intended for use by internal functions
func GetPersistentvolumeclaims(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourcePersistentvolumeclaims, error) {
	var resources apiresourcecontracts.ResourcePersistentvolumeclaims
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "PersistentVolumeClaim",
		ApiVersion: "v1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourcePersistentVolumeClaim](ctx, query)
	resources.Owner = ownerref
	resources.Persistentvolumeclaims = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource PersistentVolumeClaim")
	}
	return resources, nil
}

// Functions to get Deployments by ownerref
// The function is intended for use by internal functions
func GetDeployments(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceDeployments, error) {
	var resources apiresourcecontracts.ResourceDeployments
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Deployment",
		ApiVersion: "apps/v1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceDeployment](ctx, query)
	resources.Owner = ownerref
	resources.Deployments = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource Deployment")
	}
	return resources, nil
}

// Functions to get Storageclasses by ownerref
// The function is intended for use by internal functions
func GetStorageclasses(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceStorageclasses, error) {
	var resources apiresourcecontracts.ResourceStorageclasses
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "StorageClass",
		ApiVersion: "storage.k8s.io/v1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceStorageClass](ctx, query)
	resources.Owner = ownerref
	resources.Storageclasses = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource StorageClass")
	}
	return resources, nil
}

// Functions to get Policyreports by ownerref
// The function is intended for use by internal functions
func GetPolicyreports(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourcePolicyreports, error) {
	var resources apiresourcecontracts.ResourcePolicyreports
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "PolicyReport",
		ApiVersion: "wgpolicyk8s.io/v1alpha2",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourcePolicyReport](ctx, query)
	resources.Owner = ownerref
	resources.Policyreports = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource PolicyReport")
	}
	return resources, nil
}

// Functions to get Applications by ownerref
// The function is intended for use by internal functions
func GetApplications(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceApplications, error) {
	var resources apiresourcecontracts.ResourceApplications
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Application",
		ApiVersion: "argoproj.io/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceApplication](ctx, query)
	resources.Owner = ownerref
	resources.Applications = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource Application")
	}
	return resources, nil
}

// Functions to get Appprojects by ownerref
// The function is intended for use by internal functions
func GetAppprojects(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceAppprojects, error) {
	var resources apiresourcecontracts.ResourceAppprojects
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "AppProject",
		ApiVersion: "argoproj.io/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceAppProject](ctx, query)
	resources.Owner = ownerref
	resources.Appprojects = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource AppProject")
	}
	return resources, nil
}

// Functions to get Certificates by ownerref
// The function is intended for use by internal functions
func GetCertificates(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceCertificates, error) {
	var resources apiresourcecontracts.ResourceCertificates
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Certificate",
		ApiVersion: "cert-manager.io/v1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceCertificate](ctx, query)
	resources.Owner = ownerref
	resources.Certificates = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource Certificate")
	}
	return resources, nil
}

// Functions to get Services by ownerref
// The function is intended for use by internal functions
func GetServices(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceServices, error) {
	var resources apiresourcecontracts.ResourceServices
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Service",
		ApiVersion: "v1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceService](ctx, query)
	resources.Owner = ownerref
	resources.Services = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource Service")
	}
	return resources, nil
}

// Functions to get Pods by ownerref
// The function is intended for use by internal functions
func GetPods(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourcePods, error) {
	var resources apiresourcecontracts.ResourcePods
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Pod",
		ApiVersion: "v1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourcePod](ctx, query)
	resources.Owner = ownerref
	resources.Pods = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource Pod")
	}
	return resources, nil
}

// Functions to get Replicasets by ownerref
// The function is intended for use by internal functions
func GetReplicasets(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceReplicasets, error) {
	var resources apiresourcecontracts.ResourceReplicasets
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "ReplicaSet",
		ApiVersion: "apps/v1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceReplicaSet](ctx, query)
	resources.Owner = ownerref
	resources.Replicasets = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource ReplicaSet")
	}
	return resources, nil
}

// Functions to get Statefulsets by ownerref
// The function is intended for use by internal functions
func GetStatefulsets(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceStatefulsets, error) {
	var resources apiresourcecontracts.ResourceStatefulsets
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "StatefulSet",
		ApiVersion: "apps/v1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceStatefulSet](ctx, query)
	resources.Owner = ownerref
	resources.Statefulsets = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource StatefulSet")
	}
	return resources, nil
}

// Functions to get Daemonsets by ownerref
// The function is intended for use by internal functions
func GetDaemonsets(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceDaemonsets, error) {
	var resources apiresourcecontracts.ResourceDaemonsets
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "DaemonSet",
		ApiVersion: "apps/v1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceDaemonSet](ctx, query)
	resources.Owner = ownerref
	resources.Daemonsets = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource DaemonSet")
	}
	return resources, nil
}

// Functions to get Ingresses by ownerref
// The function is intended for use by internal functions
func GetIngresses(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceIngresses, error) {
	var resources apiresourcecontracts.ResourceIngresses
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Ingress",
		ApiVersion: "networking.k8s.io/v1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceIngress](ctx, query)
	resources.Owner = ownerref
	resources.Ingresses = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource Ingress")
	}
	return resources, nil
}

// Functions to get Ingressclasses by ownerref
// The function is intended for use by internal functions
func GetIngressclasses(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceIngressclasses, error) {
	var resources apiresourcecontracts.ResourceIngressclasses
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "IngressClass",
		ApiVersion: "networking.k8s.io/v1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceIngressClass](ctx, query)
	resources.Owner = ownerref
	resources.Ingressclasses = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource IngressClass")
	}
	return resources, nil
}

// Functions to get Vulnerabilityreports by ownerref
// The function is intended for use by internal functions
func GetVulnerabilityreports(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceVulnerabilityreports, error) {
	var resources apiresourcecontracts.ResourceVulnerabilityreports
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "VulnerabilityReport",
		ApiVersion: "aquasecurity.github.io/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceVulnerabilityReport](ctx, query)
	resources.Owner = ownerref
	resources.Vulnerabilityreports = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource VulnerabilityReport")
	}
	return resources, nil
}

// Functions to get Exposedsecretreports by ownerref
// The function is intended for use by internal functions
func GetExposedsecretreports(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceExposedsecretreports, error) {
	var resources apiresourcecontracts.ResourceExposedsecretreports
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "ExposedSecretReport",
		ApiVersion: "aquasecurity.github.io/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceExposedSecretReport](ctx, query)
	resources.Owner = ownerref
	resources.Exposedsecretreports = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource ExposedSecretReport")
	}
	return resources, nil
}

// Functions to get Configauditreports by ownerref
// The function is intended for use by internal functions
func GetConfigauditreports(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceConfigauditreports, error) {
	var resources apiresourcecontracts.ResourceConfigauditreports
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "ConfigAuditReport",
		ApiVersion: "aquasecurity.github.io/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceConfigAuditReport](ctx, query)
	resources.Owner = ownerref
	resources.Configauditreports = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource ConfigAuditReport")
	}
	return resources, nil
}

// Functions to get Rbacassessmentreports by ownerref
// The function is intended for use by internal functions
func GetRbacassessmentreports(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceRbacassessmentreports, error) {
	var resources apiresourcecontracts.ResourceRbacassessmentreports
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "RbacAssessmentReport",
		ApiVersion: "aquasecurity.github.io/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceRbacAssessmentReport](ctx, query)
	resources.Owner = ownerref
	resources.Rbacassessmentreports = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource RbacAssessmentReport")
	}
	return resources, nil
}

// Functions to get Tanzukubernetesclusters by ownerref
// The function is intended for use by internal functions
func GetTanzukubernetesclusters(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceTanzukubernetesclusters, error) {
	var resources apiresourcecontracts.ResourceTanzukubernetesclusters
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "TanzuKubernetesCluster",
		ApiVersion: "run.tanzu.vmware.com/v1alpha2",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceTanzuKubernetesCluster](ctx, query)
	resources.Owner = ownerref
	resources.Tanzukubernetesclusters = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource TanzuKubernetesCluster")
	}
	return resources, nil
}

// Functions to get Tanzukubernetesreleases by ownerref
// The function is intended for use by internal functions
func GetTanzukubernetesreleases(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceTanzukubernetesreleases, error) {
	var resources apiresourcecontracts.ResourceTanzukubernetesreleases
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "TanzuKubernetesRelease",
		ApiVersion: "run.tanzu.vmware.com/v1alpha2",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceTanzuKubernetesRelease](ctx, query)
	resources.Owner = ownerref
	resources.Tanzukubernetesreleases = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource TanzuKubernetesRelease")
	}
	return resources, nil
}

// Functions to get Virtualmachineclasses by ownerref
// The function is intended for use by internal functions
func GetVirtualmachineclasses(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceVirtualmachineclasses, error) {
	var resources apiresourcecontracts.ResourceVirtualmachineclasses
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "VirtualMachineClass",
		ApiVersion: "vmoperator.vmware.com/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceVirtualMachineClass](ctx, query)
	resources.Owner = ownerref
	resources.Virtualmachineclasses = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource VirtualMachineClass")
	}
	return resources, nil
}

// Functions to get Virtualmachineclassbindings by ownerref
// The function is intended for use by internal functions
func GetVirtualmachineclassbindings(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceVirtualmachineclassbindings, error) {
	var resources apiresourcecontracts.ResourceVirtualmachineclassbindings
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "VirtualMachineClassBinding",
		ApiVersion: "vmoperator.vmware.com/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceVirtualMachineClassBinding](ctx, query)
	resources.Owner = ownerref
	resources.Virtualmachineclassbindings = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource VirtualMachineClassBinding")
	}
	return resources, nil
}

// Functions to get Kubernetesclusters by ownerref
// The function is intended for use by internal functions
func GetKubernetesclusters(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceKubernetesclusters, error) {
	var resources apiresourcecontracts.ResourceKubernetesclusters
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "KubernetesCluster",
		ApiVersion: "general.ror.internal/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceKubernetesCluster](ctx, query)
	resources.Owner = ownerref
	resources.Kubernetesclusters = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource KubernetesCluster")
	}
	return resources, nil
}

// Functions to get Clusterorders by ownerref
// The function is intended for use by internal functions
func GetClusterorders(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceClusterorders, error) {
	var resources apiresourcecontracts.ResourceClusterorders
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "ClusterOrder",
		ApiVersion: "general.ror.internal/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceClusterOrder](ctx, query)
	resources.Owner = ownerref
	resources.Clusterorders = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource ClusterOrder")
	}
	return resources, nil
}

// Functions to get Projects by ownerref
// The function is intended for use by internal functions
func GetProjects(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceProjects, error) {
	var resources apiresourcecontracts.ResourceProjects
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Project",
		ApiVersion: "general.ror.internal/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceProject](ctx, query)
	resources.Owner = ownerref
	resources.Projects = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource Project")
	}
	return resources, nil
}

// Functions to get Configurations by ownerref
// The function is intended for use by internal functions
func GetConfigurations(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceConfigurations, error) {
	var resources apiresourcecontracts.ResourceConfigurations
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Configuration",
		ApiVersion: "general.ror.internal/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceConfiguration](ctx, query)
	resources.Owner = ownerref
	resources.Configurations = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource Configuration")
	}
	return resources, nil
}

// Functions to get Clustercompliancereports by ownerref
// The function is intended for use by internal functions
func GetClustercompliancereports(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceClustercompliancereports, error) {
	var resources apiresourcecontracts.ResourceClustercompliancereports
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "ClusterComplianceReport",
		ApiVersion: "aquasecurity.github.io/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceClusterComplianceReport](ctx, query)
	resources.Owner = ownerref
	resources.Clustercompliancereports = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource ClusterComplianceReport")
	}
	return resources, nil
}

// Functions to get Clustervulnerabilityreports by ownerref
// The function is intended for use by internal functions
func GetClustervulnerabilityreports(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceClustervulnerabilityreports, error) {
	var resources apiresourcecontracts.ResourceClustervulnerabilityreports
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "ClusterVulnerabilityReport",
		ApiVersion: "general.ror.internal/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceClusterVulnerabilityReport](ctx, query)
	resources.Owner = ownerref
	resources.Clustervulnerabilityreports = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource ClusterVulnerabilityReport")
	}
	return resources, nil
}

// Functions to get Routes by ownerref
// The function is intended for use by internal functions
func GetRoutes(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceRoutes, error) {
	var resources apiresourcecontracts.ResourceRoutes
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Route",
		ApiVersion: "general.ror.internal/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceRoute](ctx, query)
	resources.Owner = ownerref
	resources.Routes = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource Route")
	}
	return resources, nil
}

// Functions to get Slackmessages by ownerref
// The function is intended for use by internal functions
func GetSlackmessages(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceSlackmessages, error) {
	var resources apiresourcecontracts.ResourceSlackmessages
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "SlackMessage",
		ApiVersion: "general.ror.internal/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceSlackMessage](ctx, query)
	resources.Owner = ownerref
	resources.Slackmessages = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource SlackMessage")
	}
	return resources, nil
}

// Functions to get Vulnerabilityevents by ownerref
// The function is intended for use by internal functions
func GetVulnerabilityevents(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceVulnerabilityevents, error) {
	var resources apiresourcecontracts.ResourceVulnerabilityevents
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "VulnerabilityEvent",
		ApiVersion: "general.ror.internal/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceVulnerabilityEvent](ctx, query)
	resources.Owner = ownerref
	resources.Vulnerabilityevents = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource VulnerabilityEvent")
	}
	return resources, nil
}

// Functions to get Vms by ownerref
// The function is intended for use by internal functions
func GetVms(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (apiresourcecontracts.ResourceVms, error) {
	var resources apiresourcecontracts.ResourceVms
	query := apiresourcecontracts.ResourceQuery{
		Owner:      ownerref,
		Kind:       "Vm",
		ApiVersion: "general.ror.internal/v1alpha1",
	}
	resourceset, err := resourcesmongodbrepo.GetResourcesByQuery[apiresourcecontracts.ResourceVm](ctx, query)
	resources.Owner = ownerref
	resources.Vms = resourceset
	if err != nil {
		return resources, errors.New("could not fetch resource Vm")
	}
	return resources, nil
}

// Function to creates a resource by the 'apiresourcecontracts.ResourceUpdateModel'
func ResourceCreateService(ctx context.Context, resourceUpdate apiresourcecontracts.ResourceUpdateModel) error {
	var err error

	if resourceUpdate.ApiVersion == "v1" && resourceUpdate.Kind == "Namespace" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNamespace]](resourceUpdate)
		resource = filterInNamespace(resource)
		err = resourcesmongodbrepo.CreateResourceNamespace(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "v1" && resourceUpdate.Kind == "Node" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNode]](resourceUpdate)
		resource = filterInNode(resource)
		err = resourcesmongodbrepo.CreateResourceNode(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "v1" && resourceUpdate.Kind == "PersistentVolumeClaim" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePersistentVolumeClaim]](resourceUpdate)
		resource = filterInPersistentVolumeClaim(resource)
		err = resourcesmongodbrepo.CreateResourcePersistentVolumeClaim(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "apps/v1" && resourceUpdate.Kind == "Deployment" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDeployment]](resourceUpdate)
		resource = filterInDeployment(resource)
		err = resourcesmongodbrepo.CreateResourceDeployment(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "storage.k8s.io/v1" && resourceUpdate.Kind == "StorageClass" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStorageClass]](resourceUpdate)
		resource = filterInStorageClass(resource)
		err = resourcesmongodbrepo.CreateResourceStorageClass(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "wgpolicyk8s.io/v1alpha2" && resourceUpdate.Kind == "PolicyReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePolicyReport]](resourceUpdate)
		resource = filterInPolicyReport(resource)
		err = resourcesmongodbrepo.CreateResourcePolicyReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "argoproj.io/v1alpha1" && resourceUpdate.Kind == "Application" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceApplication]](resourceUpdate)
		resource = filterInApplication(resource)
		err = resourcesmongodbrepo.CreateResourceApplication(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "argoproj.io/v1alpha1" && resourceUpdate.Kind == "AppProject" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceAppProject]](resourceUpdate)
		resource = filterInAppProject(resource)
		err = resourcesmongodbrepo.CreateResourceAppProject(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "cert-manager.io/v1" && resourceUpdate.Kind == "Certificate" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceCertificate]](resourceUpdate)
		resource = filterInCertificate(resource)
		err = resourcesmongodbrepo.CreateResourceCertificate(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "v1" && resourceUpdate.Kind == "Service" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceService]](resourceUpdate)
		resource = filterInService(resource)
		err = resourcesmongodbrepo.CreateResourceService(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "v1" && resourceUpdate.Kind == "Pod" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePod]](resourceUpdate)
		resource = filterInPod(resource)
		err = resourcesmongodbrepo.CreateResourcePod(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "apps/v1" && resourceUpdate.Kind == "ReplicaSet" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceReplicaSet]](resourceUpdate)
		resource = filterInReplicaSet(resource)
		err = resourcesmongodbrepo.CreateResourceReplicaSet(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "apps/v1" && resourceUpdate.Kind == "StatefulSet" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStatefulSet]](resourceUpdate)
		resource = filterInStatefulSet(resource)
		err = resourcesmongodbrepo.CreateResourceStatefulSet(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "apps/v1" && resourceUpdate.Kind == "DaemonSet" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDaemonSet]](resourceUpdate)
		resource = filterInDaemonSet(resource)
		err = resourcesmongodbrepo.CreateResourceDaemonSet(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "networking.k8s.io/v1" && resourceUpdate.Kind == "Ingress" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngress]](resourceUpdate)
		resource = filterInIngress(resource)
		err = resourcesmongodbrepo.CreateResourceIngress(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "networking.k8s.io/v1" && resourceUpdate.Kind == "IngressClass" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngressClass]](resourceUpdate)
		resource = filterInIngressClass(resource)
		err = resourcesmongodbrepo.CreateResourceIngressClass(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "aquasecurity.github.io/v1alpha1" && resourceUpdate.Kind == "VulnerabilityReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityReport]](resourceUpdate)
		resource = filterInVulnerabilityReport(resource)
		err = resourcesmongodbrepo.CreateResourceVulnerabilityReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "aquasecurity.github.io/v1alpha1" && resourceUpdate.Kind == "ExposedSecretReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceExposedSecretReport]](resourceUpdate)
		resource = filterInExposedSecretReport(resource)
		err = resourcesmongodbrepo.CreateResourceExposedSecretReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "aquasecurity.github.io/v1alpha1" && resourceUpdate.Kind == "ConfigAuditReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfigAuditReport]](resourceUpdate)
		resource = filterInConfigAuditReport(resource)
		err = resourcesmongodbrepo.CreateResourceConfigAuditReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "aquasecurity.github.io/v1alpha1" && resourceUpdate.Kind == "RbacAssessmentReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRbacAssessmentReport]](resourceUpdate)
		resource = filterInRbacAssessmentReport(resource)
		err = resourcesmongodbrepo.CreateResourceRbacAssessmentReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "run.tanzu.vmware.com/v1alpha2" && resourceUpdate.Kind == "TanzuKubernetesCluster" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesCluster]](resourceUpdate)
		resource = filterInTanzuKubernetesCluster(resource)
		err = resourcesmongodbrepo.CreateResourceTanzuKubernetesCluster(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "run.tanzu.vmware.com/v1alpha2" && resourceUpdate.Kind == "TanzuKubernetesRelease" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesRelease]](resourceUpdate)
		resource = filterInTanzuKubernetesRelease(resource)
		err = resourcesmongodbrepo.CreateResourceTanzuKubernetesRelease(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "vmoperator.vmware.com/v1alpha1" && resourceUpdate.Kind == "VirtualMachineClass" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClass]](resourceUpdate)
		resource = filterInVirtualMachineClass(resource)
		err = resourcesmongodbrepo.CreateResourceVirtualMachineClass(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "vmoperator.vmware.com/v1alpha1" && resourceUpdate.Kind == "VirtualMachineClassBinding" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClassBinding]](resourceUpdate)
		resource = filterInVirtualMachineClassBinding(resource)
		err = resourcesmongodbrepo.CreateResourceVirtualMachineClassBinding(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "KubernetesCluster" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceKubernetesCluster]](resourceUpdate)
		resource = filterInKubernetesCluster(resource)
		err = resourcesmongodbrepo.CreateResourceKubernetesCluster(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "ClusterOrder" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterOrder]](resourceUpdate)
		resource = filterInClusterOrder(resource)
		err = resourcesmongodbrepo.CreateResourceClusterOrder(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "Project" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceProject]](resourceUpdate)
		resource = filterInProject(resource)
		err = resourcesmongodbrepo.CreateResourceProject(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "Configuration" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfiguration]](resourceUpdate)
		resource = filterInConfiguration(resource)
		err = resourcesmongodbrepo.CreateResourceConfiguration(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "aquasecurity.github.io/v1alpha1" && resourceUpdate.Kind == "ClusterComplianceReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterComplianceReport]](resourceUpdate)
		resource = filterInClusterComplianceReport(resource)
		err = resourcesmongodbrepo.CreateResourceClusterComplianceReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "ClusterVulnerabilityReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterVulnerabilityReport]](resourceUpdate)
		resource = filterInClusterVulnerabilityReport(resource)
		err = resourcesmongodbrepo.CreateResourceClusterVulnerabilityReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "Route" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRoute]](resourceUpdate)
		resource = filterInRoute(resource)
		err = resourcesmongodbrepo.CreateResourceRoute(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "SlackMessage" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceSlackMessage]](resourceUpdate)
		resource = filterInSlackMessage(resource)
		err = resourcesmongodbrepo.CreateResourceSlackMessage(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "VulnerabilityEvent" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityEvent]](resourceUpdate)
		resource = filterInVulnerabilityEvent(resource)
		err = resourcesmongodbrepo.CreateResourceVulnerabilityEvent(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "Vm" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVm]](resourceUpdate)
		resource = filterInVm(resource)
		err = resourcesmongodbrepo.CreateResourceVm(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionAdd)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if err != nil {
		rlog.Errorc(ctx, "could not create resource", err)
		return err
	}

	return nil

}

// Function to update a resource by the 'apiresourcecontracts.ResourceUpdateModel' struct
func ResourceUpdateService(ctx context.Context, resourceUpdate apiresourcecontracts.ResourceUpdateModel) error {
	var err error

	if resourceUpdate.ApiVersion == "v1" && resourceUpdate.Kind == "Namespace" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNamespace]](resourceUpdate)
		resource = filterInNamespace(resource)
		err = resourcesmongodbrepo.UpdateResourceNamespace(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "v1" && resourceUpdate.Kind == "Node" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceNode]](resourceUpdate)
		resource = filterInNode(resource)
		err = resourcesmongodbrepo.UpdateResourceNode(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "v1" && resourceUpdate.Kind == "PersistentVolumeClaim" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePersistentVolumeClaim]](resourceUpdate)
		resource = filterInPersistentVolumeClaim(resource)
		err = resourcesmongodbrepo.UpdateResourcePersistentVolumeClaim(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "apps/v1" && resourceUpdate.Kind == "Deployment" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDeployment]](resourceUpdate)
		resource = filterInDeployment(resource)
		err = resourcesmongodbrepo.UpdateResourceDeployment(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "storage.k8s.io/v1" && resourceUpdate.Kind == "StorageClass" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStorageClass]](resourceUpdate)
		resource = filterInStorageClass(resource)
		err = resourcesmongodbrepo.UpdateResourceStorageClass(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "wgpolicyk8s.io/v1alpha2" && resourceUpdate.Kind == "PolicyReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePolicyReport]](resourceUpdate)
		resource = filterInPolicyReport(resource)
		err = resourcesmongodbrepo.UpdateResourcePolicyReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "argoproj.io/v1alpha1" && resourceUpdate.Kind == "Application" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceApplication]](resourceUpdate)
		resource = filterInApplication(resource)
		err = resourcesmongodbrepo.UpdateResourceApplication(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "argoproj.io/v1alpha1" && resourceUpdate.Kind == "AppProject" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceAppProject]](resourceUpdate)
		resource = filterInAppProject(resource)
		err = resourcesmongodbrepo.UpdateResourceAppProject(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "cert-manager.io/v1" && resourceUpdate.Kind == "Certificate" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceCertificate]](resourceUpdate)
		resource = filterInCertificate(resource)
		err = resourcesmongodbrepo.UpdateResourceCertificate(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "v1" && resourceUpdate.Kind == "Service" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceService]](resourceUpdate)
		resource = filterInService(resource)
		err = resourcesmongodbrepo.UpdateResourceService(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "v1" && resourceUpdate.Kind == "Pod" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourcePod]](resourceUpdate)
		resource = filterInPod(resource)
		err = resourcesmongodbrepo.UpdateResourcePod(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "apps/v1" && resourceUpdate.Kind == "ReplicaSet" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceReplicaSet]](resourceUpdate)
		resource = filterInReplicaSet(resource)
		err = resourcesmongodbrepo.UpdateResourceReplicaSet(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "apps/v1" && resourceUpdate.Kind == "StatefulSet" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceStatefulSet]](resourceUpdate)
		resource = filterInStatefulSet(resource)
		err = resourcesmongodbrepo.UpdateResourceStatefulSet(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "apps/v1" && resourceUpdate.Kind == "DaemonSet" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceDaemonSet]](resourceUpdate)
		resource = filterInDaemonSet(resource)
		err = resourcesmongodbrepo.UpdateResourceDaemonSet(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "networking.k8s.io/v1" && resourceUpdate.Kind == "Ingress" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngress]](resourceUpdate)
		resource = filterInIngress(resource)
		err = resourcesmongodbrepo.UpdateResourceIngress(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "networking.k8s.io/v1" && resourceUpdate.Kind == "IngressClass" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceIngressClass]](resourceUpdate)
		resource = filterInIngressClass(resource)
		err = resourcesmongodbrepo.UpdateResourceIngressClass(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "aquasecurity.github.io/v1alpha1" && resourceUpdate.Kind == "VulnerabilityReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityReport]](resourceUpdate)
		resource = filterInVulnerabilityReport(resource)
		err = resourcesmongodbrepo.UpdateResourceVulnerabilityReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "aquasecurity.github.io/v1alpha1" && resourceUpdate.Kind == "ExposedSecretReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceExposedSecretReport]](resourceUpdate)
		resource = filterInExposedSecretReport(resource)
		err = resourcesmongodbrepo.UpdateResourceExposedSecretReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "aquasecurity.github.io/v1alpha1" && resourceUpdate.Kind == "ConfigAuditReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfigAuditReport]](resourceUpdate)
		resource = filterInConfigAuditReport(resource)
		err = resourcesmongodbrepo.UpdateResourceConfigAuditReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "aquasecurity.github.io/v1alpha1" && resourceUpdate.Kind == "RbacAssessmentReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRbacAssessmentReport]](resourceUpdate)
		resource = filterInRbacAssessmentReport(resource)
		err = resourcesmongodbrepo.UpdateResourceRbacAssessmentReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "run.tanzu.vmware.com/v1alpha2" && resourceUpdate.Kind == "TanzuKubernetesCluster" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesCluster]](resourceUpdate)
		resource = filterInTanzuKubernetesCluster(resource)
		err = resourcesmongodbrepo.UpdateResourceTanzuKubernetesCluster(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "run.tanzu.vmware.com/v1alpha2" && resourceUpdate.Kind == "TanzuKubernetesRelease" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceTanzuKubernetesRelease]](resourceUpdate)
		resource = filterInTanzuKubernetesRelease(resource)
		err = resourcesmongodbrepo.UpdateResourceTanzuKubernetesRelease(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "vmoperator.vmware.com/v1alpha1" && resourceUpdate.Kind == "VirtualMachineClass" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClass]](resourceUpdate)
		resource = filterInVirtualMachineClass(resource)
		err = resourcesmongodbrepo.UpdateResourceVirtualMachineClass(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "vmoperator.vmware.com/v1alpha1" && resourceUpdate.Kind == "VirtualMachineClassBinding" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVirtualMachineClassBinding]](resourceUpdate)
		resource = filterInVirtualMachineClassBinding(resource)
		err = resourcesmongodbrepo.UpdateResourceVirtualMachineClassBinding(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "KubernetesCluster" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceKubernetesCluster]](resourceUpdate)
		resource = filterInKubernetesCluster(resource)
		err = resourcesmongodbrepo.UpdateResourceKubernetesCluster(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "ClusterOrder" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterOrder]](resourceUpdate)
		resource = filterInClusterOrder(resource)
		err = resourcesmongodbrepo.UpdateResourceClusterOrder(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "Project" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceProject]](resourceUpdate)
		resource = filterInProject(resource)
		err = resourcesmongodbrepo.UpdateResourceProject(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "Configuration" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceConfiguration]](resourceUpdate)
		resource = filterInConfiguration(resource)
		err = resourcesmongodbrepo.UpdateResourceConfiguration(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "aquasecurity.github.io/v1alpha1" && resourceUpdate.Kind == "ClusterComplianceReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterComplianceReport]](resourceUpdate)
		resource = filterInClusterComplianceReport(resource)
		err = resourcesmongodbrepo.UpdateResourceClusterComplianceReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "ClusterVulnerabilityReport" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceClusterVulnerabilityReport]](resourceUpdate)
		resource = filterInClusterVulnerabilityReport(resource)
		err = resourcesmongodbrepo.UpdateResourceClusterVulnerabilityReport(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "Route" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceRoute]](resourceUpdate)
		resource = filterInRoute(resource)
		err = resourcesmongodbrepo.UpdateResourceRoute(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "SlackMessage" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceSlackMessage]](resourceUpdate)
		resource = filterInSlackMessage(resource)
		err = resourcesmongodbrepo.UpdateResourceSlackMessage(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "VulnerabilityEvent" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVulnerabilityEvent]](resourceUpdate)
		resource = filterInVulnerabilityEvent(resource)
		err = resourcesmongodbrepo.UpdateResourceVulnerabilityEvent(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if resourceUpdate.ApiVersion == "general.ror.internal/v1alpha1" && resourceUpdate.Kind == "Vm" {
		resource := resourcesmongodbrepo.MapToResourceModel[apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceVm]](resourceUpdate)
		resource = filterInVm(resource)
		err = resourcesmongodbrepo.UpdateResourceVm(resource, ctx)
		if err == nil {
			err = sendToMessageBus(ctx, resource, apiresourcecontracts.K8sActionUpdate)
			if err != nil {
				rlog.Errorc(ctx, "could not send to message bus", err)
			}
		}
	}

	if err != nil {
		rlog.Errorc(ctx, "could not update resource", err)
		return err
	}

	return nil
}
