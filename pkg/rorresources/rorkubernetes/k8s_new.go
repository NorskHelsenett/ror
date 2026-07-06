// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package rorkubernetes

import (
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewResourceSetFromMapInterface(input map[string]any) *rorresources.ResourceSet {
	var rs rorresources.ResourceSet
	r := NewResourceFromMapInterface(input)
	rs.Add(r)
	return &rs

}

func newCommonResourceFromMapInterface(input map[string]any) v1.ObjectMeta {
	metadata, ok := input["metadata"].(map[string]any)

	if !ok {
		rlog.Warn("could not convert input to metav1.ObjectMeta", rlog.Any("input", input))
		return v1.ObjectMeta{}
	}
	// Convert the metadata map to a v1.ObjectMeta struct
	metadataConverted := &v1.ObjectMeta{}
	err := convertUnstructuredToStruct(metadata, metadataConverted)
	if err != nil {
		rlog.Error("could not convert input to metav1.ObjectMeta", err)
		return v1.ObjectMeta{}
	}

	return *metadataConverted
}

// NewResourceFromMapInterface creates a new resource from a map[string]any
// type provided by the kubernetes universal client.
func NewResourceFromMapInterface(input map[string]any) *rorresources.Resource {
	r := rorresources.NewRorResource(input["kind"].(string), input["apiVersion"].(string))
	r.SetMetadata(newCommonResourceFromMapInterface(input))

	switch r.GroupVersionKind().String() {

	case "/v1, Kind=Namespace":
		res := newNamespaceFromMapInterface(input)
		r.SetNamespace(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "/v1, Kind=Node":
		res := newNodeFromMapInterface(input)
		r.SetNode(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "/v1, Kind=PersistentVolumeClaim":
		res := newPersistentVolumeClaimFromMapInterface(input)
		r.SetPersistentVolumeClaim(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "apps/v1, Kind=Deployment":
		res := newDeploymentFromMapInterface(input)
		r.SetDeployment(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "storage.k8s.io/v1, Kind=StorageClass":
		res := newStorageClassFromMapInterface(input)
		r.SetStorageClass(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "wgpolicyk8s.io/v1alpha2, Kind=PolicyReport":
		res := newPolicyReportFromMapInterface(input)
		r.SetPolicyReport(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "argoproj.io/v1alpha1, Kind=Application":
		res := newApplicationFromMapInterface(input)
		r.SetApplication(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "argoproj.io/v1alpha1, Kind=AppProject":
		res := newAppProjectFromMapInterface(input)
		r.SetAppProject(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "cert-manager.io/v1, Kind=Certificate":
		res := newCertificateFromMapInterface(input)
		r.SetCertificate(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "/v1, Kind=Service":
		res := newServiceFromMapInterface(input)
		r.SetService(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "/v1, Kind=Pod":
		res := newPodFromMapInterface(input)
		r.SetPod(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "apps/v1, Kind=ReplicaSet":
		res := newReplicaSetFromMapInterface(input)
		r.SetReplicaSet(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "apps/v1, Kind=StatefulSet":
		res := newStatefulSetFromMapInterface(input)
		r.SetStatefulSet(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "apps/v1, Kind=DaemonSet":
		res := newDaemonSetFromMapInterface(input)
		r.SetDaemonSet(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "networking.k8s.io/v1, Kind=Ingress":
		res := newIngressFromMapInterface(input)
		r.SetIngress(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "networking.k8s.io/v1, Kind=IngressClass":
		res := newIngressClassFromMapInterface(input)
		r.SetIngressClass(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "aquasecurity.github.io/v1alpha1, Kind=SbomReport":
		res := newSbomReportFromMapInterface(input)
		r.SetSbomReport(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "aquasecurity.github.io/v1alpha1, Kind=VulnerabilityReport":
		res := newVulnerabilityReportFromMapInterface(input)
		r.SetVulnerabilityReport(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "aquasecurity.github.io/v1alpha1, Kind=ExposedSecretReport":
		res := newExposedSecretReportFromMapInterface(input)
		r.SetExposedSecretReport(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "aquasecurity.github.io/v1alpha1, Kind=ConfigAuditReport":
		res := newConfigAuditReportFromMapInterface(input)
		r.SetConfigAuditReport(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "aquasecurity.github.io/v1alpha1, Kind=RbacAssessmentReport":
		res := newRbacAssessmentReportFromMapInterface(input)
		r.SetRbacAssessmentReport(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "run.tanzu.vmware.com/v1alpha3, Kind=TanzuKubernetesCluster":
		res := newTanzuKubernetesClusterFromMapInterface(input)
		r.SetTanzuKubernetesCluster(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "run.tanzu.vmware.com/v1alpha3, Kind=TanzuKubernetesRelease":
		res := newTanzuKubernetesReleaseFromMapInterface(input)
		r.SetTanzuKubernetesRelease(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "vmoperator.vmware.com/v1alpha2, Kind=VirtualMachineClass":
		res := newVirtualMachineClassFromMapInterface(input)
		r.SetVirtualMachineClass(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "general.ror.internal/v1alpha1, Kind=KubernetesCluster":
		res := newKubernetesClusterFromMapInterface(input)
		r.SetKubernetesCluster(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "general.ror.internal/v1alpha1, Kind=Provider":
		res := newProviderFromMapInterface(input)
		r.SetProvider(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "general.ror.internal/v1alpha1, Kind=Workspace":
		res := newWorkspaceFromMapInterface(input)
		r.SetWorkspace(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "general.ror.internal/v1alpha1, Kind=KubernetesMachineClass":
		res := newKubernetesMachineClassFromMapInterface(input)
		r.SetKubernetesMachineClass(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "general.ror.internal/v1alpha1, Kind=ClusterOrder":
		res := newClusterOrderFromMapInterface(input)
		r.SetClusterOrder(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "general.ror.internal/v1alpha1, Kind=Project":
		res := newProjectFromMapInterface(input)
		r.SetProject(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "general.ror.internal/v1alpha1, Kind=Configuration":
		res := newConfigurationFromMapInterface(input)
		r.SetConfiguration(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "aquasecurity.github.io/v1alpha1, Kind=ClusterComplianceReport":
		res := newClusterComplianceReportFromMapInterface(input)
		r.SetClusterComplianceReport(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "general.ror.internal/v1alpha1, Kind=ClusterVulnerabilityReport":
		res := newClusterVulnerabilityReportFromMapInterface(input)
		r.SetClusterVulnerabilityReport(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "general.ror.internal/v1alpha1, Kind=Route":
		res := newRouteFromMapInterface(input)
		r.SetRoute(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "general.ror.internal/v1alpha1, Kind=SlackMessage":
		res := newSlackMessageFromMapInterface(input)
		r.SetSlackMessage(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "general.ror.internal/v1alpha1, Kind=VulnerabilityEvent":
		res := newVulnerabilityEventFromMapInterface(input)
		r.SetVulnerabilityEvent(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "general.ror.internal/v1alpha1, Kind=VirtualMachine":
		res := newVirtualMachineFromMapInterface(input)
		r.SetVirtualMachine(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "general.ror.internal/v1alpha1, Kind=VirtualMachineVulnerabilityInfo":
		res := newVirtualMachineVulnerabilityInfoFromMapInterface(input)
		r.SetVirtualMachineVulnerabilityInfo(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "/v1, Kind=Endpoints":
		res := newEndpointsFromMapInterface(input)
		r.SetEndpoints(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "networking.k8s.io/v1, Kind=NetworkPolicy":
		res := newNetworkPolicyFromMapInterface(input)
		r.SetNetworkPolicy(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "infrastructure.ror.internal/v1alpha1, Kind=Datacenter":
		res := newDatacenterFromMapInterface(input)
		r.SetDatacenter(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "backup.ror.internal/v1alpha1, Kind=BackupJob":
		res := newBackupJobFromMapInterface(input)
		r.SetBackupJob(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "backup.ror.internal/v1alpha1, Kind=BackupRun":
		res := newBackupRunFromMapInterface(input)
		r.SetBackupRun(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "machine.ror.internal/v1alpha1, Kind=Machine":
		res := newMachineFromMapInterface(input)
		r.SetMachine(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "database.ror.internal/v1, Kind=ManagedDatabase":
		res := newManagedDatabaseFromMapInterface(input)
		r.SetManagedDatabase(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "ror.internal/v1, Kind=Config":
		res := newConfigFromMapInterface(input)
		r.SetConfig(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "ror.internal/v1, Kind=OrganizationalUnit":
		res := newOrganizationalUnitFromMapInterface(input)
		r.SetOrganizationalUnit(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	case "unknown.ror.internal/v1, Kind=Unknown":
		res := newUnknownFromMapInterface(input)
		r.SetUnknown(res)
		r.SetCommonInterface(rortypes.NewCommonFactory(res))

	default:
		rlog.Warn("could not create ResourceSet")
		return nil
	}
	return r
}

// newNamespaceFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newNamespaceFromMapInterface(input map[string]any) *rortypes.ResourceNamespace {
	result := rortypes.ResourceNamespace{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceNamespace", err)
		return nil
	}

	return &result
}

// newNodeFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newNodeFromMapInterface(input map[string]any) *rortypes.ResourceNode {
	result := rortypes.ResourceNode{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceNode", err)
		return nil
	}

	return &result
}

// newPersistentVolumeClaimFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newPersistentVolumeClaimFromMapInterface(input map[string]any) *rortypes.ResourcePersistentVolumeClaim {
	result := rortypes.ResourcePersistentVolumeClaim{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourcePersistentVolumeClaim", err)
		return nil
	}

	return &result
}

// newDeploymentFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newDeploymentFromMapInterface(input map[string]any) *rortypes.ResourceDeployment {
	result := rortypes.ResourceDeployment{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceDeployment", err)
		return nil
	}

	return &result
}

// newStorageClassFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newStorageClassFromMapInterface(input map[string]any) *rortypes.ResourceStorageClass {
	result := rortypes.ResourceStorageClass{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceStorageClass", err)
		return nil
	}

	return &result
}

// newPolicyReportFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newPolicyReportFromMapInterface(input map[string]any) *rortypes.ResourcePolicyReport {
	result := rortypes.ResourcePolicyReport{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourcePolicyReport", err)
		return nil
	}

	return &result
}

// newApplicationFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newApplicationFromMapInterface(input map[string]any) *rortypes.ResourceApplication {
	result := rortypes.ResourceApplication{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceApplication", err)
		return nil
	}

	return &result
}

// newAppProjectFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newAppProjectFromMapInterface(input map[string]any) *rortypes.ResourceAppProject {
	result := rortypes.ResourceAppProject{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceAppProject", err)
		return nil
	}

	return &result
}

// newCertificateFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newCertificateFromMapInterface(input map[string]any) *rortypes.ResourceCertificate {
	result := rortypes.ResourceCertificate{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceCertificate", err)
		return nil
	}

	return &result
}

// newServiceFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newServiceFromMapInterface(input map[string]any) *rortypes.ResourceService {
	result := rortypes.ResourceService{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceService", err)
		return nil
	}

	return &result
}

// newPodFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newPodFromMapInterface(input map[string]any) *rortypes.ResourcePod {
	result := rortypes.ResourcePod{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourcePod", err)
		return nil
	}

	return &result
}

// newReplicaSetFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newReplicaSetFromMapInterface(input map[string]any) *rortypes.ResourceReplicaSet {
	result := rortypes.ResourceReplicaSet{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceReplicaSet", err)
		return nil
	}

	return &result
}

// newStatefulSetFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newStatefulSetFromMapInterface(input map[string]any) *rortypes.ResourceStatefulSet {
	result := rortypes.ResourceStatefulSet{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceStatefulSet", err)
		return nil
	}

	return &result
}

// newDaemonSetFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newDaemonSetFromMapInterface(input map[string]any) *rortypes.ResourceDaemonSet {
	result := rortypes.ResourceDaemonSet{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceDaemonSet", err)
		return nil
	}

	return &result
}

// newIngressFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newIngressFromMapInterface(input map[string]any) *rortypes.ResourceIngress {
	result := rortypes.ResourceIngress{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceIngress", err)
		return nil
	}

	return &result
}

// newIngressClassFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newIngressClassFromMapInterface(input map[string]any) *rortypes.ResourceIngressClass {
	result := rortypes.ResourceIngressClass{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceIngressClass", err)
		return nil
	}

	return &result
}

// newSbomReportFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newSbomReportFromMapInterface(input map[string]any) *rortypes.ResourceSbomReport {
	result := rortypes.ResourceSbomReport{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceSbomReport", err)
		return nil
	}

	return &result
}

// newVulnerabilityReportFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newVulnerabilityReportFromMapInterface(input map[string]any) *rortypes.ResourceVulnerabilityReport {
	result := rortypes.ResourceVulnerabilityReport{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceVulnerabilityReport", err)
		return nil
	}

	return &result
}

// newExposedSecretReportFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newExposedSecretReportFromMapInterface(input map[string]any) *rortypes.ResourceExposedSecretReport {
	result := rortypes.ResourceExposedSecretReport{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceExposedSecretReport", err)
		return nil
	}

	return &result
}

// newConfigAuditReportFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newConfigAuditReportFromMapInterface(input map[string]any) *rortypes.ResourceConfigAuditReport {
	result := rortypes.ResourceConfigAuditReport{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceConfigAuditReport", err)
		return nil
	}

	return &result
}

// newRbacAssessmentReportFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newRbacAssessmentReportFromMapInterface(input map[string]any) *rortypes.ResourceRbacAssessmentReport {
	result := rortypes.ResourceRbacAssessmentReport{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceRbacAssessmentReport", err)
		return nil
	}

	return &result
}

// newTanzuKubernetesClusterFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newTanzuKubernetesClusterFromMapInterface(input map[string]any) *rortypes.ResourceTanzuKubernetesCluster {
	result := rortypes.ResourceTanzuKubernetesCluster{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceTanzuKubernetesCluster", err)
		return nil
	}

	return &result
}

// newTanzuKubernetesReleaseFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newTanzuKubernetesReleaseFromMapInterface(input map[string]any) *rortypes.ResourceTanzuKubernetesRelease {
	result := rortypes.ResourceTanzuKubernetesRelease{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceTanzuKubernetesRelease", err)
		return nil
	}

	return &result
}

// newVirtualMachineClassFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newVirtualMachineClassFromMapInterface(input map[string]any) *rortypes.ResourceVirtualMachineClass {
	result := rortypes.ResourceVirtualMachineClass{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceVirtualMachineClass", err)
		return nil
	}

	return &result
}

// newKubernetesClusterFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newKubernetesClusterFromMapInterface(input map[string]any) *rortypes.ResourceKubernetesCluster {
	result := rortypes.ResourceKubernetesCluster{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceKubernetesCluster", err)
		return nil
	}

	return &result
}

// newProviderFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newProviderFromMapInterface(input map[string]any) *rortypes.ResourceProvider {
	result := rortypes.ResourceProvider{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceProvider", err)
		return nil
	}

	return &result
}

// newWorkspaceFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newWorkspaceFromMapInterface(input map[string]any) *rortypes.ResourceWorkspace {
	result := rortypes.ResourceWorkspace{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceWorkspace", err)
		return nil
	}

	return &result
}

// newKubernetesMachineClassFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newKubernetesMachineClassFromMapInterface(input map[string]any) *rortypes.ResourceKubernetesMachineClass {
	result := rortypes.ResourceKubernetesMachineClass{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceKubernetesMachineClass", err)
		return nil
	}

	return &result
}

// newClusterOrderFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newClusterOrderFromMapInterface(input map[string]any) *rortypes.ResourceClusterOrder {
	result := rortypes.ResourceClusterOrder{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceClusterOrder", err)
		return nil
	}

	return &result
}

// newProjectFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newProjectFromMapInterface(input map[string]any) *rortypes.ResourceProject {
	result := rortypes.ResourceProject{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceProject", err)
		return nil
	}

	return &result
}

// newConfigurationFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newConfigurationFromMapInterface(input map[string]any) *rortypes.ResourceConfiguration {
	result := rortypes.ResourceConfiguration{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceConfiguration", err)
		return nil
	}

	return &result
}

// newClusterComplianceReportFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newClusterComplianceReportFromMapInterface(input map[string]any) *rortypes.ResourceClusterComplianceReport {
	result := rortypes.ResourceClusterComplianceReport{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceClusterComplianceReport", err)
		return nil
	}

	return &result
}

// newClusterVulnerabilityReportFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newClusterVulnerabilityReportFromMapInterface(input map[string]any) *rortypes.ResourceClusterVulnerabilityReport {
	result := rortypes.ResourceClusterVulnerabilityReport{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceClusterVulnerabilityReport", err)
		return nil
	}

	return &result
}

// newRouteFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newRouteFromMapInterface(input map[string]any) *rortypes.ResourceRoute {
	result := rortypes.ResourceRoute{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceRoute", err)
		return nil
	}

	return &result
}

// newSlackMessageFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newSlackMessageFromMapInterface(input map[string]any) *rortypes.ResourceSlackMessage {
	result := rortypes.ResourceSlackMessage{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceSlackMessage", err)
		return nil
	}

	return &result
}

// newVulnerabilityEventFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newVulnerabilityEventFromMapInterface(input map[string]any) *rortypes.ResourceVulnerabilityEvent {
	result := rortypes.ResourceVulnerabilityEvent{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceVulnerabilityEvent", err)
		return nil
	}

	return &result
}

// newVirtualMachineFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newVirtualMachineFromMapInterface(input map[string]any) *rortypes.ResourceVirtualMachine {
	result := rortypes.ResourceVirtualMachine{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceVirtualMachine", err)
		return nil
	}

	return &result
}

// newVirtualMachineVulnerabilityInfoFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newVirtualMachineVulnerabilityInfoFromMapInterface(input map[string]any) *rortypes.ResourceVirtualMachineVulnerabilityInfo {
	result := rortypes.ResourceVirtualMachineVulnerabilityInfo{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceVirtualMachineVulnerabilityInfo", err)
		return nil
	}

	return &result
}

// newEndpointsFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newEndpointsFromMapInterface(input map[string]any) *rortypes.ResourceEndpoints {
	result := rortypes.ResourceEndpoints{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceEndpoints", err)
		return nil
	}

	return &result
}

// newNetworkPolicyFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newNetworkPolicyFromMapInterface(input map[string]any) *rortypes.ResourceNetworkPolicy {
	result := rortypes.ResourceNetworkPolicy{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceNetworkPolicy", err)
		return nil
	}

	return &result
}

// newDatacenterFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newDatacenterFromMapInterface(input map[string]any) *rortypes.ResourceDatacenter {
	result := rortypes.ResourceDatacenter{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceDatacenter", err)
		return nil
	}

	return &result
}

// newBackupJobFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newBackupJobFromMapInterface(input map[string]any) *rortypes.ResourceBackupJob {
	result := rortypes.ResourceBackupJob{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceBackupJob", err)
		return nil
	}

	return &result
}

// newBackupRunFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newBackupRunFromMapInterface(input map[string]any) *rortypes.ResourceBackupRun {
	result := rortypes.ResourceBackupRun{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceBackupRun", err)
		return nil
	}

	return &result
}

// newMachineFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newMachineFromMapInterface(input map[string]any) *rortypes.ResourceMachine {
	result := rortypes.ResourceMachine{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceMachine", err)
		return nil
	}

	return &result
}

// newManagedDatabaseFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newManagedDatabaseFromMapInterface(input map[string]any) *rortypes.ResourceManagedDatabase {
	result := rortypes.ResourceManagedDatabase{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceManagedDatabase", err)
		return nil
	}

	return &result
}

// newConfigFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newConfigFromMapInterface(input map[string]any) *rortypes.ResourceConfig {
	result := rortypes.ResourceConfig{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceConfig", err)
		return nil
	}

	return &result
}

// newOrganizationalUnitFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newOrganizationalUnitFromMapInterface(input map[string]any) *rortypes.ResourceOrganizationalUnit {
	result := rortypes.ResourceOrganizationalUnit{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceOrganizationalUnit", err)
		return nil
	}

	return &result
}

// newUnknownFromMapInterface creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newUnknownFromMapInterface(input map[string]any) *rortypes.ResourceUnknown {
	result := rortypes.ResourceUnknown{}
	err := convertUnstructuredToStruct(input, &result)

	if err != nil {
		rlog.Error("could not convert input to ResourceUnknown", err)
		return nil
	}

	return &result
}
