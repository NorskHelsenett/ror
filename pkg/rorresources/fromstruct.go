package rorresources

import (
	"github.com/NorskHelsenett/ror/pkg/rlog"
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
		r.SetNamespace(res.NamespaceResource)
		r.SetCommonInterface(res.NamespaceResource)

	case "/v1, Kind=Node":
		r.SetNode(res.NodeResource)
		r.SetCommonInterface(res.NodeResource)

	case "/v1, Kind=PersistentVolumeClaim":
		r.SetPersistentVolumeClaim(res.PersistentVolumeClaimResource)
		r.SetCommonInterface(res.PersistentVolumeClaimResource)

	case "apps/v1, Kind=Deployment":
		r.SetDeployment(res.DeploymentResource)
		r.SetCommonInterface(res.DeploymentResource)

	case "storage.k8s.io/v1, Kind=StorageClass":
		r.SetStorageClass(res.StorageClassResource)
		r.SetCommonInterface(res.StorageClassResource)

	case "wgpolicyk8s.io/v1alpha2, Kind=PolicyReport":
		r.SetPolicyReport(res.PolicyReportResource)
		r.SetCommonInterface(res.PolicyReportResource)

	case "argoproj.io/v1alpha1, Kind=Application":
		r.SetApplication(res.ApplicationResource)
		r.SetCommonInterface(res.ApplicationResource)

	case "argoproj.io/v1alpha1, Kind=AppProject":
		r.SetAppProject(res.AppProjectResource)
		r.SetCommonInterface(res.AppProjectResource)

	case "cert-manager.io/v1, Kind=Certificate":
		r.SetCertificate(res.CertificateResource)
		r.SetCommonInterface(res.CertificateResource)

	case "/v1, Kind=Service":
		r.SetService(res.ServiceResource)
		r.SetCommonInterface(res.ServiceResource)

	case "/v1, Kind=Pod":
		r.SetPod(res.PodResource)
		r.SetCommonInterface(res.PodResource)

	case "apps/v1, Kind=ReplicaSet":
		r.SetReplicaSet(res.ReplicaSetResource)
		r.SetCommonInterface(res.ReplicaSetResource)

	case "apps/v1, Kind=StatefulSet":
		r.SetStatefulSet(res.StatefulSetResource)
		r.SetCommonInterface(res.StatefulSetResource)

	case "apps/v1, Kind=DaemonSet":
		r.SetDaemonSet(res.DaemonSetResource)
		r.SetCommonInterface(res.DaemonSetResource)

	case "networking.k8s.io/v1, Kind=Ingress":
		r.SetIngress(res.IngressResource)
		r.SetCommonInterface(res.IngressResource)

	case "networking.k8s.io/v1, Kind=IngressClass":
		r.SetIngressClass(res.IngressClassResource)
		r.SetCommonInterface(res.IngressClassResource)

	case "aquasecurity.github.io/v1alpha1, Kind=VulnerabilityReport":
		r.SetVulnerabilityReport(res.VulnerabilityReportResource)
		r.SetCommonInterface(res.VulnerabilityReportResource)

	case "aquasecurity.github.io/v1alpha1, Kind=ExposedSecretReport":
		r.SetExposedSecretReport(res.ExposedSecretReportResource)
		r.SetCommonInterface(res.ExposedSecretReportResource)

	case "aquasecurity.github.io/v1alpha1, Kind=ConfigAuditReport":
		r.SetConfigAuditReport(res.ConfigAuditReportResource)
		r.SetCommonInterface(res.ConfigAuditReportResource)

	case "aquasecurity.github.io/v1alpha1, Kind=RbacAssessmentReport":
		r.SetRbacAssessmentReport(res.RbacAssessmentReportResource)
		r.SetCommonInterface(res.RbacAssessmentReportResource)

	case "run.tanzu.vmware.com/v1alpha2, Kind=TanzuKubernetesCluster":
		r.SetTanzuKubernetesCluster(res.TanzuKubernetesClusterResource)
		r.SetCommonInterface(res.TanzuKubernetesClusterResource)

	case "run.tanzu.vmware.com/v1alpha2, Kind=TanzuKubernetesRelease":
		r.SetTanzuKubernetesRelease(res.TanzuKubernetesReleaseResource)
		r.SetCommonInterface(res.TanzuKubernetesReleaseResource)

	case "vmoperator.vmware.com/v1alpha1, Kind=VirtualMachineClass":
		r.SetVirtualMachineClass(res.VirtualMachineClassResource)
		r.SetCommonInterface(res.VirtualMachineClassResource)

	case "vmoperator.vmware.com/v1alpha1, Kind=VirtualMachineClassBinding":
		r.SetVirtualMachineClassBinding(res.VirtualMachineClassBindingResource)
		r.SetCommonInterface(res.VirtualMachineClassBindingResource)

	case "general.ror.internal/v1alpha1, Kind=KubernetesCluster":
		r.SetKubernetesCluster(res.KubernetesClusterResource)
		r.SetCommonInterface(res.KubernetesClusterResource)

	case "general.ror.internal/v1alpha1, Kind=ClusterOrder":
		r.SetClusterOrder(res.ClusterOrderResource)
		r.SetCommonInterface(res.ClusterOrderResource)

	case "general.ror.internal/v1alpha1, Kind=Project":
		r.SetProject(res.ProjectResource)
		r.SetCommonInterface(res.ProjectResource)

	case "general.ror.internal/v1alpha1, Kind=Configuration":
		r.SetConfiguration(res.ConfigurationResource)
		r.SetCommonInterface(res.ConfigurationResource)

	case "aquasecurity.github.io/v1alpha1, Kind=ClusterComplianceReport":
		r.SetClusterComplianceReport(res.ClusterComplianceReportResource)
		r.SetCommonInterface(res.ClusterComplianceReportResource)

	case "general.ror.internal/v1alpha1, Kind=ClusterVulnerabilityReport":
		r.SetClusterVulnerabilityReport(res.ClusterVulnerabilityReportResource)
		r.SetCommonInterface(res.ClusterVulnerabilityReportResource)

	case "general.ror.internal/v1alpha1, Kind=Route":
		r.SetRoute(res.RouteResource)
		r.SetCommonInterface(res.RouteResource)

	case "general.ror.internal/v1alpha1, Kind=SlackMessage":
		r.SetSlackMessage(res.SlackMessageResource)
		r.SetCommonInterface(res.SlackMessageResource)

	case "general.ror.internal/v1alpha1, Kind=VulnerabilityEvent":
		r.SetVulnerabilityEvent(res.VulnerabilityEventResource)
		r.SetCommonInterface(res.VulnerabilityEventResource)

	case "general.ror.internal/v1alpha1, Kind=VirtualMachine":
		r.SetVirtualMachine(res.VirtualMachineResource)
		r.SetCommonInterface(res.VirtualMachineResource)

	default:
		rlog.Info("Unknown resource kind", rlog.String("gvk", gvk.String()), rlog.String("kind", res.Kind), rlog.String("apiVersion", res.APIVersion))
	}
	return r
}
