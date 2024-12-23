// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package rorkubernetes

import (
	"encoding/json"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func NewResourceSetFromDynamicClient(input *unstructured.Unstructured) *rorresources.ResourceSet {
	var rs rorresources.ResourceSet
	r := NewResourceFromDynamicClient(input)
	rs.Add(r)
	return &rs

}

func NewCommonResourceFromDynamicClient(input *unstructured.Unstructured) rortypes.CommonResource {
	r := rortypes.CommonResource{}
	rjson, err := input.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(rjson, &r)
	if err != nil {
		rlog.Error("Could not unmarshal json to Common", err)
	}
	return r

}

// NewResourceFromDynamicClient creates a new resource from a unstructured.Unstructured
// type provided by the kubernetes universal client.
func NewResourceFromDynamicClient(input *unstructured.Unstructured) *rorresources.Resource {
	r := rorresources.NewRorResource(input.GetKind(), input.GetAPIVersion())
	r.CommonResource = NewCommonResourceFromDynamicClient(input)

	switch input.GroupVersionKind().String() {

	case "/v1, Kind=Namespace":
		res := newNamespaceFromDynamicClient(input)
		r.SetNamespace(res)
		r.SetCommonInterface(res)

	case "/v1, Kind=Node":
		res := newNodeFromDynamicClient(input)
		r.SetNode(res)
		r.SetCommonInterface(res)

	case "/v1, Kind=PersistentVolumeClaim":
		res := newPersistentVolumeClaimFromDynamicClient(input)
		r.SetPersistentVolumeClaim(res)
		r.SetCommonInterface(res)

	case "apps/v1, Kind=Deployment":
		res := newDeploymentFromDynamicClient(input)
		r.SetDeployment(res)
		r.SetCommonInterface(res)

	case "storage.k8s.io/v1, Kind=StorageClass":
		res := newStorageClassFromDynamicClient(input)
		r.SetStorageClass(res)
		r.SetCommonInterface(res)

	case "wgpolicyk8s.io/v1alpha2, Kind=PolicyReport":
		res := newPolicyReportFromDynamicClient(input)
		r.SetPolicyReport(res)
		r.SetCommonInterface(res)

	case "argoproj.io/v1alpha1, Kind=Application":
		res := newApplicationFromDynamicClient(input)
		r.SetApplication(res)
		r.SetCommonInterface(res)

	case "argoproj.io/v1alpha1, Kind=AppProject":
		res := newAppProjectFromDynamicClient(input)
		r.SetAppProject(res)
		r.SetCommonInterface(res)

	case "cert-manager.io/v1, Kind=Certificate":
		res := newCertificateFromDynamicClient(input)
		r.SetCertificate(res)
		r.SetCommonInterface(res)

	case "/v1, Kind=Service":
		res := newServiceFromDynamicClient(input)
		r.SetService(res)
		r.SetCommonInterface(res)

	case "/v1, Kind=Pod":
		res := newPodFromDynamicClient(input)
		r.SetPod(res)
		r.SetCommonInterface(res)

	case "apps/v1, Kind=ReplicaSet":
		res := newReplicaSetFromDynamicClient(input)
		r.SetReplicaSet(res)
		r.SetCommonInterface(res)

	case "apps/v1, Kind=StatefulSet":
		res := newStatefulSetFromDynamicClient(input)
		r.SetStatefulSet(res)
		r.SetCommonInterface(res)

	case "apps/v1, Kind=DaemonSet":
		res := newDaemonSetFromDynamicClient(input)
		r.SetDaemonSet(res)
		r.SetCommonInterface(res)

	case "networking.k8s.io/v1, Kind=Ingress":
		res := newIngressFromDynamicClient(input)
		r.SetIngress(res)
		r.SetCommonInterface(res)

	case "networking.k8s.io/v1, Kind=IngressClass":
		res := newIngressClassFromDynamicClient(input)
		r.SetIngressClass(res)
		r.SetCommonInterface(res)

	case "aquasecurity.github.io/v1alpha1, Kind=VulnerabilityReport":
		res := newVulnerabilityReportFromDynamicClient(input)
		r.SetVulnerabilityReport(res)
		r.SetCommonInterface(res)

	case "aquasecurity.github.io/v1alpha1, Kind=ExposedSecretReport":
		res := newExposedSecretReportFromDynamicClient(input)
		r.SetExposedSecretReport(res)
		r.SetCommonInterface(res)

	case "aquasecurity.github.io/v1alpha1, Kind=ConfigAuditReport":
		res := newConfigAuditReportFromDynamicClient(input)
		r.SetConfigAuditReport(res)
		r.SetCommonInterface(res)

	case "aquasecurity.github.io/v1alpha1, Kind=RbacAssessmentReport":
		res := newRbacAssessmentReportFromDynamicClient(input)
		r.SetRbacAssessmentReport(res)
		r.SetCommonInterface(res)

	case "run.tanzu.vmware.com/v1alpha2, Kind=TanzuKubernetesCluster":
		res := newTanzuKubernetesClusterFromDynamicClient(input)
		r.SetTanzuKubernetesCluster(res)
		r.SetCommonInterface(res)

	case "run.tanzu.vmware.com/v1alpha2, Kind=TanzuKubernetesRelease":
		res := newTanzuKubernetesReleaseFromDynamicClient(input)
		r.SetTanzuKubernetesRelease(res)
		r.SetCommonInterface(res)

	case "vmoperator.vmware.com/v1alpha1, Kind=VirtualMachineClass":
		res := newVirtualMachineClassFromDynamicClient(input)
		r.SetVirtualMachineClass(res)
		r.SetCommonInterface(res)

	case "vmoperator.vmware.com/v1alpha1, Kind=VirtualMachineClassBinding":
		res := newVirtualMachineClassBindingFromDynamicClient(input)
		r.SetVirtualMachineClassBinding(res)
		r.SetCommonInterface(res)

	case "general.ror.internal/v1alpha1, Kind=KubernetesCluster":
		res := newKubernetesClusterFromDynamicClient(input)
		r.SetKubernetesCluster(res)
		r.SetCommonInterface(res)

	case "general.ror.internal/v1alpha1, Kind=ClusterOrder":
		res := newClusterOrderFromDynamicClient(input)
		r.SetClusterOrder(res)
		r.SetCommonInterface(res)

	case "general.ror.internal/v1alpha1, Kind=Project":
		res := newProjectFromDynamicClient(input)
		r.SetProject(res)
		r.SetCommonInterface(res)

	case "general.ror.internal/v1alpha1, Kind=Configuration":
		res := newConfigurationFromDynamicClient(input)
		r.SetConfiguration(res)
		r.SetCommonInterface(res)

	case "aquasecurity.github.io/v1alpha1, Kind=ClusterComplianceReport":
		res := newClusterComplianceReportFromDynamicClient(input)
		r.SetClusterComplianceReport(res)
		r.SetCommonInterface(res)

	case "general.ror.internal/v1alpha1, Kind=ClusterVulnerabilityReport":
		res := newClusterVulnerabilityReportFromDynamicClient(input)
		r.SetClusterVulnerabilityReport(res)
		r.SetCommonInterface(res)

	case "general.ror.internal/v1alpha1, Kind=Route":
		res := newRouteFromDynamicClient(input)
		r.SetRoute(res)
		r.SetCommonInterface(res)

	case "general.ror.internal/v1alpha1, Kind=SlackMessage":
		res := newSlackMessageFromDynamicClient(input)
		r.SetSlackMessage(res)
		r.SetCommonInterface(res)

	case "general.ror.internal/v1alpha1, Kind=VulnerabilityEvent":
		res := newVulnerabilityEventFromDynamicClient(input)
		r.SetVulnerabilityEvent(res)
		r.SetCommonInterface(res)

	case "general.ror.internal/v1alpha1, Kind=VirtualMachine":
		res := newVirtualMachineFromDynamicClient(input)
		r.SetVirtualMachine(res)
		r.SetCommonInterface(res)

	case "/v1, Kind=Endpoints":
		res := newEndpointsFromDynamicClient(input)
		r.SetEndpoints(res)
		r.SetCommonInterface(res)

	case "networking.k8s.io/v1, Kind=NetworkPolicy":
		res := newNetworkPolicyFromDynamicClient(input)
		r.SetNetworkPolicy(res)
		r.SetCommonInterface(res)

	case "general.ror.internal/v1alpha1, Kind=FirewallPolicy":
		res := newFirewallPolicyFromDynamicClient(input)
		r.SetFirewallPolicy(res)
		r.SetCommonInterface(res)

	default:
		rlog.Warn("could not create ResourceSet")
		return r
	}
	return r
}

// newNamespaceFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newNamespaceFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceNamespace {
	nr := rortypes.ResourceNamespace{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to Namespace", err)
	}
	return &nr
}

// newNodeFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newNodeFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceNode {
	nr := rortypes.ResourceNode{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to Node", err)
	}
	return &nr
}

// newPersistentVolumeClaimFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newPersistentVolumeClaimFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourcePersistentVolumeClaim {
	nr := rortypes.ResourcePersistentVolumeClaim{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to PersistentVolumeClaim", err)
	}
	return &nr
}

// newDeploymentFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newDeploymentFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceDeployment {
	nr := rortypes.ResourceDeployment{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to Deployment", err)
	}
	return &nr
}

// newStorageClassFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newStorageClassFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceStorageClass {
	nr := rortypes.ResourceStorageClass{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to StorageClass", err)
	}
	return &nr
}

// newPolicyReportFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newPolicyReportFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourcePolicyReport {
	nr := rortypes.ResourcePolicyReport{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to PolicyReport", err)
	}
	return &nr
}

// newApplicationFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newApplicationFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceApplication {
	nr := rortypes.ResourceApplication{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to Application", err)
	}
	return &nr
}

// newAppProjectFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newAppProjectFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceAppProject {
	nr := rortypes.ResourceAppProject{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to AppProject", err)
	}
	return &nr
}

// newCertificateFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newCertificateFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceCertificate {
	nr := rortypes.ResourceCertificate{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to Certificate", err)
	}
	return &nr
}

// newServiceFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newServiceFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceService {
	nr := rortypes.ResourceService{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to Service", err)
	}
	return &nr
}

// newPodFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newPodFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourcePod {
	nr := rortypes.ResourcePod{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to Pod", err)
	}
	return &nr
}

// newReplicaSetFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newReplicaSetFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceReplicaSet {
	nr := rortypes.ResourceReplicaSet{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to ReplicaSet", err)
	}
	return &nr
}

// newStatefulSetFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newStatefulSetFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceStatefulSet {
	nr := rortypes.ResourceStatefulSet{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to StatefulSet", err)
	}
	return &nr
}

// newDaemonSetFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newDaemonSetFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceDaemonSet {
	nr := rortypes.ResourceDaemonSet{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to DaemonSet", err)
	}
	return &nr
}

// newIngressFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newIngressFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceIngress {
	nr := rortypes.ResourceIngress{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to Ingress", err)
	}
	return &nr
}

// newIngressClassFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newIngressClassFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceIngressClass {
	nr := rortypes.ResourceIngressClass{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to IngressClass", err)
	}
	return &nr
}

// newVulnerabilityReportFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newVulnerabilityReportFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceVulnerabilityReport {
	nr := rortypes.ResourceVulnerabilityReport{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to VulnerabilityReport", err)
	}
	return &nr
}

// newExposedSecretReportFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newExposedSecretReportFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceExposedSecretReport {
	nr := rortypes.ResourceExposedSecretReport{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to ExposedSecretReport", err)
	}
	return &nr
}

// newConfigAuditReportFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newConfigAuditReportFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceConfigAuditReport {
	nr := rortypes.ResourceConfigAuditReport{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to ConfigAuditReport", err)
	}
	return &nr
}

// newRbacAssessmentReportFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newRbacAssessmentReportFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceRbacAssessmentReport {
	nr := rortypes.ResourceRbacAssessmentReport{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to RbacAssessmentReport", err)
	}
	return &nr
}

// newTanzuKubernetesClusterFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newTanzuKubernetesClusterFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceTanzuKubernetesCluster {
	nr := rortypes.ResourceTanzuKubernetesCluster{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to TanzuKubernetesCluster", err)
	}
	return &nr
}

// newTanzuKubernetesReleaseFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newTanzuKubernetesReleaseFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceTanzuKubernetesRelease {
	nr := rortypes.ResourceTanzuKubernetesRelease{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to TanzuKubernetesRelease", err)
	}
	return &nr
}

// newVirtualMachineClassFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newVirtualMachineClassFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceVirtualMachineClass {
	nr := rortypes.ResourceVirtualMachineClass{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to VirtualMachineClass", err)
	}
	return &nr
}

// newVirtualMachineClassBindingFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newVirtualMachineClassBindingFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceVirtualMachineClassBinding {
	nr := rortypes.ResourceVirtualMachineClassBinding{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to VirtualMachineClassBinding", err)
	}
	return &nr
}

// newKubernetesClusterFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newKubernetesClusterFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceKubernetesCluster {
	nr := rortypes.ResourceKubernetesCluster{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to KubernetesCluster", err)
	}
	return &nr
}

// newClusterOrderFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newClusterOrderFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceClusterOrder {
	nr := rortypes.ResourceClusterOrder{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to ClusterOrder", err)
	}
	return &nr
}

// newProjectFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newProjectFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceProject {
	nr := rortypes.ResourceProject{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to Project", err)
	}
	return &nr
}

// newConfigurationFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newConfigurationFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceConfiguration {
	nr := rortypes.ResourceConfiguration{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to Configuration", err)
	}
	return &nr
}

// newClusterComplianceReportFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newClusterComplianceReportFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceClusterComplianceReport {
	nr := rortypes.ResourceClusterComplianceReport{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to ClusterComplianceReport", err)
	}
	return &nr
}

// newClusterVulnerabilityReportFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newClusterVulnerabilityReportFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceClusterVulnerabilityReport {
	nr := rortypes.ResourceClusterVulnerabilityReport{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to ClusterVulnerabilityReport", err)
	}
	return &nr
}

// newRouteFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newRouteFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceRoute {
	nr := rortypes.ResourceRoute{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to Route", err)
	}
	return &nr
}

// newSlackMessageFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newSlackMessageFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceSlackMessage {
	nr := rortypes.ResourceSlackMessage{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to SlackMessage", err)
	}
	return &nr
}

// newVulnerabilityEventFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newVulnerabilityEventFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceVulnerabilityEvent {
	nr := rortypes.ResourceVulnerabilityEvent{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to VulnerabilityEvent", err)
	}
	return &nr
}

// newVirtualMachineFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newVirtualMachineFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceVirtualMachine {
	nr := rortypes.ResourceVirtualMachine{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to VirtualMachine", err)
	}
	return &nr
}

// newEndpointsFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newEndpointsFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceEndpoints {
	nr := rortypes.ResourceEndpoints{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to Endpoints", err)
	}
	return &nr
}

// newNetworkPolicyFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newNetworkPolicyFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceNetworkPolicy {
	nr := rortypes.ResourceNetworkPolicy{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to NetworkPolicy", err)
	}
	return &nr
}

// newFirewallPolicyFromDynamicClient creates the underlying resource from a unstructured.Unstructured type provided
// by the kubernetes universal client.
func newFirewallPolicyFromDynamicClient(obj *unstructured.Unstructured) *rortypes.ResourceFirewallPolicy {
	nr := rortypes.ResourceFirewallPolicy{}
	nrjson, err := obj.MarshalJSON()
	if err != nil {
		rlog.Error("Could not mashal unstructired to json", err)
	}

	err = json.Unmarshal(nrjson, &nr)
	if err != nil {
		rlog.Error("Could not unmarshal json to FirewallPolicy", err)
	}
	return &nr
}
