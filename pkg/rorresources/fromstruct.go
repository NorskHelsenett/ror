package rorresources

import (
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// NewResourceSetFromStruct creates a new ResourceSet from a struct of the type ResourceSet
// the function restores common methods after transit eg by json.
func NewResourceSetFromStruct(in ResourceSet) *ResourceSet {
	out := ResourceSet{}
	if len(in.Resources) == 0 {
		rlog.Warn("ResourceSet has no resources")
		return &out
	}

	if in.query == nil {
		in.query = &ResourceQuery{}
	}

	query := *in.query

	out.query = &query

	for _, res := range in.Resources {
		r := NewResourceFromStruct(*res)
		out.Add(r)
	}
	return &out

}

func NewResourceFromStruct(res Resource) *Resource {

	r := NewRorResource(res.Kind, res.APIVersion)
	r.CommonResource = res.CommonResource

	gvk := schema.FromAPIVersionAndKind(res.APIVersion, res.Kind)
	switch gvk.String() {

	case "/v1, Kind=Namespace":
		if res.NamespaceResource == nil {
			res.NamespaceResource = &rortypes.ResourceNamespace{}
		}
		r.SetNamespace(res.NamespaceResource)
		r.SetCommonInterface(res.NamespaceResource)

	case "/v1, Kind=Node":
		if res.NodeResource == nil {
			res.NodeResource = &rortypes.ResourceNode{}
		}
		r.SetNode(res.NodeResource)
		r.SetCommonInterface(res.NodeResource)

	case "/v1, Kind=PersistentVolumeClaim":
		if res.PersistentVolumeClaimResource == nil {
			res.PersistentVolumeClaimResource = &rortypes.ResourcePersistentVolumeClaim{}
		}
		r.SetPersistentVolumeClaim(res.PersistentVolumeClaimResource)
		r.SetCommonInterface(res.PersistentVolumeClaimResource)

	case "apps/v1, Kind=Deployment":
		if res.DeploymentResource == nil {
			res.DeploymentResource = &rortypes.ResourceDeployment{}
		}
		r.SetDeployment(res.DeploymentResource)
		r.SetCommonInterface(res.DeploymentResource)

	case "storage.k8s.io/v1, Kind=StorageClass":
		if res.StorageClassResource == nil {
			res.StorageClassResource = &rortypes.ResourceStorageClass{}
		}
		r.SetStorageClass(res.StorageClassResource)
		r.SetCommonInterface(res.StorageClassResource)

	case "wgpolicyk8s.io/v1alpha2, Kind=PolicyReport":
		if res.PolicyReportResource == nil {
			res.PolicyReportResource = &rortypes.ResourcePolicyReport{}
		}
		r.SetPolicyReport(res.PolicyReportResource)
		r.SetCommonInterface(res.PolicyReportResource)

	case "argoproj.io/v1alpha1, Kind=Application":
		if res.ApplicationResource == nil {
			res.ApplicationResource = &rortypes.ResourceApplication{}
		}
		r.SetApplication(res.ApplicationResource)
		r.SetCommonInterface(res.ApplicationResource)

	case "argoproj.io/v1alpha1, Kind=AppProject":
		if res.AppProjectResource == nil {
			res.AppProjectResource = &rortypes.ResourceAppProject{}
		}
		r.SetAppProject(res.AppProjectResource)
		r.SetCommonInterface(res.AppProjectResource)

	case "cert-manager.io/v1, Kind=Certificate":
		if res.CertificateResource == nil {
			res.CertificateResource = &rortypes.ResourceCertificate{}
		}
		r.SetCertificate(res.CertificateResource)
		r.SetCommonInterface(res.CertificateResource)

	case "/v1, Kind=Service":
		if res.ServiceResource == nil {
			res.ServiceResource = &rortypes.ResourceService{}
		}
		r.SetService(res.ServiceResource)
		r.SetCommonInterface(res.ServiceResource)

	case "/v1, Kind=Pod":
		if res.PodResource == nil {
			res.PodResource = &rortypes.ResourcePod{}
		}
		r.SetPod(res.PodResource)
		r.SetCommonInterface(res.PodResource)

	case "apps/v1, Kind=ReplicaSet":
		if res.ReplicaSetResource == nil {
			res.ReplicaSetResource = &rortypes.ResourceReplicaSet{}
		}
		r.SetReplicaSet(res.ReplicaSetResource)
		r.SetCommonInterface(res.ReplicaSetResource)

	case "apps/v1, Kind=StatefulSet":
		if res.StatefulSetResource == nil {
			res.StatefulSetResource = &rortypes.ResourceStatefulSet{}
		}
		r.SetStatefulSet(res.StatefulSetResource)
		r.SetCommonInterface(res.StatefulSetResource)

	case "apps/v1, Kind=DaemonSet":
		if res.DaemonSetResource == nil {
			res.DaemonSetResource = &rortypes.ResourceDaemonSet{}
		}
		r.SetDaemonSet(res.DaemonSetResource)
		r.SetCommonInterface(res.DaemonSetResource)

	case "networking.k8s.io/v1, Kind=Ingress":
		if res.IngressResource == nil {
			res.IngressResource = &rortypes.ResourceIngress{}
		}
		r.SetIngress(res.IngressResource)
		r.SetCommonInterface(res.IngressResource)

	case "networking.k8s.io/v1, Kind=IngressClass":
		if res.IngressClassResource == nil {
			res.IngressClassResource = &rortypes.ResourceIngressClass{}
		}
		r.SetIngressClass(res.IngressClassResource)
		r.SetCommonInterface(res.IngressClassResource)

	case "aquasecurity.github.io/v1alpha1, Kind=SbomReport":
		if res.SbomReportResource == nil {
			res.SbomReportResource = &rortypes.ResourceSbomReport{}
		}
		r.SetSbomReport(res.SbomReportResource)
		r.SetCommonInterface(res.SbomReportResource)

	case "aquasecurity.github.io/v1alpha1, Kind=VulnerabilityReport":
		if res.VulnerabilityReportResource == nil {
			res.VulnerabilityReportResource = &rortypes.ResourceVulnerabilityReport{}
		}
		r.SetVulnerabilityReport(res.VulnerabilityReportResource)
		r.SetCommonInterface(res.VulnerabilityReportResource)

	case "aquasecurity.github.io/v1alpha1, Kind=ExposedSecretReport":
		if res.ExposedSecretReportResource == nil {
			res.ExposedSecretReportResource = &rortypes.ResourceExposedSecretReport{}
		}
		r.SetExposedSecretReport(res.ExposedSecretReportResource)
		r.SetCommonInterface(res.ExposedSecretReportResource)

	case "aquasecurity.github.io/v1alpha1, Kind=ConfigAuditReport":
		if res.ConfigAuditReportResource == nil {
			res.ConfigAuditReportResource = &rortypes.ResourceConfigAuditReport{}
		}
		r.SetConfigAuditReport(res.ConfigAuditReportResource)
		r.SetCommonInterface(res.ConfigAuditReportResource)

	case "aquasecurity.github.io/v1alpha1, Kind=RbacAssessmentReport":
		if res.RbacAssessmentReportResource == nil {
			res.RbacAssessmentReportResource = &rortypes.ResourceRbacAssessmentReport{}
		}
		r.SetRbacAssessmentReport(res.RbacAssessmentReportResource)
		r.SetCommonInterface(res.RbacAssessmentReportResource)

	case "run.tanzu.vmware.com/v1alpha3, Kind=TanzuKubernetesCluster":
		if res.TanzuKubernetesClusterResource == nil {
			res.TanzuKubernetesClusterResource = &rortypes.ResourceTanzuKubernetesCluster{}
		}
		r.SetTanzuKubernetesCluster(res.TanzuKubernetesClusterResource)
		r.SetCommonInterface(res.TanzuKubernetesClusterResource)

	case "run.tanzu.vmware.com/v1alpha3, Kind=TanzuKubernetesRelease":
		if res.TanzuKubernetesReleaseResource == nil {
			res.TanzuKubernetesReleaseResource = &rortypes.ResourceTanzuKubernetesRelease{}
		}
		r.SetTanzuKubernetesRelease(res.TanzuKubernetesReleaseResource)
		r.SetCommonInterface(res.TanzuKubernetesReleaseResource)

	case "vmoperator.vmware.com/v1alpha2, Kind=VirtualMachineClass":
		if res.VirtualMachineClassResource == nil {
			res.VirtualMachineClassResource = &rortypes.ResourceVirtualMachineClass{}
		}
		r.SetVirtualMachineClass(res.VirtualMachineClassResource)
		r.SetCommonInterface(res.VirtualMachineClassResource)

	case "general.ror.internal/v1alpha1, Kind=KubernetesCluster":
		if res.KubernetesClusterResource == nil {
			res.KubernetesClusterResource = &rortypes.ResourceKubernetesCluster{}
		}
		r.SetKubernetesCluster(res.KubernetesClusterResource)
		r.SetCommonInterface(res.KubernetesClusterResource)

	case "general.ror.internal/v1alpha1, Kind=Provider":
		if res.ProviderResource == nil {
			res.ProviderResource = &rortypes.ResourceProvider{}
		}
		r.SetProvider(res.ProviderResource)
		r.SetCommonInterface(res.ProviderResource)

	case "general.ror.internal/v1alpha1, Kind=Workspace":
		if res.WorkspaceResource == nil {
			res.WorkspaceResource = &rortypes.ResourceWorkspace{}
		}
		r.SetWorkspace(res.WorkspaceResource)
		r.SetCommonInterface(res.WorkspaceResource)

	case "general.ror.internal/v1alpha1, Kind=KubernetesMachineClass":
		if res.KubernetesMachineClassResource == nil {
			res.KubernetesMachineClassResource = &rortypes.ResourceKubernetesMachineClass{}
		}
		r.SetKubernetesMachineClass(res.KubernetesMachineClassResource)
		r.SetCommonInterface(res.KubernetesMachineClassResource)

	case "general.ror.internal/v1alpha1, Kind=ClusterOrder":
		if res.ClusterOrderResource == nil {
			res.ClusterOrderResource = &rortypes.ResourceClusterOrder{}
		}
		r.SetClusterOrder(res.ClusterOrderResource)
		r.SetCommonInterface(res.ClusterOrderResource)

	case "general.ror.internal/v1alpha1, Kind=Project":
		if res.ProjectResource == nil {
			res.ProjectResource = &rortypes.ResourceProject{}
		}
		r.SetProject(res.ProjectResource)
		r.SetCommonInterface(res.ProjectResource)

	case "general.ror.internal/v1alpha1, Kind=Configuration":
		if res.ConfigurationResource == nil {
			res.ConfigurationResource = &rortypes.ResourceConfiguration{}
		}
		r.SetConfiguration(res.ConfigurationResource)
		r.SetCommonInterface(res.ConfigurationResource)

	case "aquasecurity.github.io/v1alpha1, Kind=ClusterComplianceReport":
		if res.ClusterComplianceReportResource == nil {
			res.ClusterComplianceReportResource = &rortypes.ResourceClusterComplianceReport{}
		}
		r.SetClusterComplianceReport(res.ClusterComplianceReportResource)
		r.SetCommonInterface(res.ClusterComplianceReportResource)

	case "general.ror.internal/v1alpha1, Kind=ClusterVulnerabilityReport":
		if res.ClusterVulnerabilityReportResource == nil {
			res.ClusterVulnerabilityReportResource = &rortypes.ResourceClusterVulnerabilityReport{}
		}
		r.SetClusterVulnerabilityReport(res.ClusterVulnerabilityReportResource)
		r.SetCommonInterface(res.ClusterVulnerabilityReportResource)

	case "general.ror.internal/v1alpha1, Kind=Route":
		if res.RouteResource == nil {
			res.RouteResource = &rortypes.ResourceRoute{}
		}
		r.SetRoute(res.RouteResource)
		r.SetCommonInterface(res.RouteResource)

	case "general.ror.internal/v1alpha1, Kind=SlackMessage":
		if res.SlackMessageResource == nil {
			res.SlackMessageResource = &rortypes.ResourceSlackMessage{}
		}
		r.SetSlackMessage(res.SlackMessageResource)
		r.SetCommonInterface(res.SlackMessageResource)

	case "general.ror.internal/v1alpha1, Kind=VulnerabilityEvent":
		if res.VulnerabilityEventResource == nil {
			res.VulnerabilityEventResource = &rortypes.ResourceVulnerabilityEvent{}
		}
		r.SetVulnerabilityEvent(res.VulnerabilityEventResource)
		r.SetCommonInterface(res.VulnerabilityEventResource)

	case "general.ror.internal/v1alpha1, Kind=VirtualMachine":
		if res.VirtualMachineResource == nil {
			res.VirtualMachineResource = &rortypes.ResourceVirtualMachine{}
		}
		r.SetVirtualMachine(res.VirtualMachineResource)
		r.SetCommonInterface(res.VirtualMachineResource)

	case "general.ror.internal/v1alpha1, Kind=VirtualMachineVulnerabilityInfo":
		if res.VirtualMachineVulnerabilityInfoResource == nil {
			res.VirtualMachineVulnerabilityInfoResource = &rortypes.ResourceVirtualMachineVulnerabilityInfo{}
		}
		r.SetVirtualMachineVulnerabilityInfo(res.VirtualMachineVulnerabilityInfoResource)
		r.SetCommonInterface(res.VirtualMachineVulnerabilityInfoResource)

	case "/v1, Kind=Endpoints":
		if res.EndpointsResource == nil {
			res.EndpointsResource = &rortypes.ResourceEndpoints{}
		}
		r.SetEndpoints(res.EndpointsResource)
		r.SetCommonInterface(res.EndpointsResource)

	case "networking.k8s.io/v1, Kind=NetworkPolicy":
		if res.NetworkPolicyResource == nil {
			res.NetworkPolicyResource = &rortypes.ResourceNetworkPolicy{}
		}
		r.SetNetworkPolicy(res.NetworkPolicyResource)
		r.SetCommonInterface(res.NetworkPolicyResource)

	case "infrastructure.ror.internal/v1alpha1, Kind=Datacenter":
		if res.DatacenterResource == nil {
			res.DatacenterResource = &rortypes.ResourceDatacenter{}
		}
		r.SetDatacenter(res.DatacenterResource)
		r.SetCommonInterface(res.DatacenterResource)

	case "backup.ror.internal/v1alpha1, Kind=BackupJob":
		if res.BackupJobResource == nil {
			res.BackupJobResource = &rortypes.ResourceBackupJob{}
		}
		r.SetBackupJob(res.BackupJobResource)
		r.SetCommonInterface(res.BackupJobResource)

	case "backup.ror.internal/v1alpha1, Kind=BackupRun":
		if res.BackupRunResource == nil {
			res.BackupRunResource = &rortypes.ResourceBackupRun{}
		}
		r.SetBackupRun(res.BackupRunResource)
		r.SetCommonInterface(res.BackupRunResource)

	case "machine.ror.internal/v1alpha1, Kind=Machine":
		if res.MachineResource == nil {
			res.MachineResource = &rortypes.ResourceMachine{}
		}
		r.SetMachine(res.MachineResource)
		r.SetCommonInterface(res.MachineResource)

	case "unknown.ror.internal/v1, Kind=Unknown":
		if res.UnknownResource == nil {
			res.UnknownResource = &rortypes.ResourceUnknown{}
		}
		r.SetUnknown(res.UnknownResource)
		r.SetCommonInterface(res.UnknownResource)

	default:
		rlog.Info("Unknown resource kind", rlog.String("gvk", gvk.String()), rlog.String("kind", res.Kind), rlog.String("apiVersion", res.APIVersion))
		return nil
	}
	return r
}
