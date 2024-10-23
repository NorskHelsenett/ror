// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package rorkubernetes

import (
	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestNewResourceSetFromDynamicClientNamespace(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Namespace",
			"apiVersion": "v1",
			"metadata": map[string]interface{}{
				"name": "test-namespace",
			},
		},
	}

	expected := rorresources.NewRorResource("Namespace", "v1")
	expected.SetNamespace(newNamespaceFromDynamicClient(input))
	expected.SetCommonInterface(newNamespaceFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientNode(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Node",
			"apiVersion": "v1",
			"metadata": map[string]interface{}{
				"name": "test-node",
			},
		},
	}

	expected := rorresources.NewRorResource("Node", "v1")
	expected.SetNode(newNodeFromDynamicClient(input))
	expected.SetCommonInterface(newNodeFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientPersistentVolumeClaim(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "PersistentVolumeClaim",
			"apiVersion": "v1",
			"metadata": map[string]interface{}{
				"name": "test-persistentvolumeclaim",
			},
		},
	}

	expected := rorresources.NewRorResource("PersistentVolumeClaim", "v1")
	expected.SetPersistentVolumeClaim(newPersistentVolumeClaimFromDynamicClient(input))
	expected.SetCommonInterface(newPersistentVolumeClaimFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientDeployment(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Deployment",
			"apiVersion": "apps/v1",
			"metadata": map[string]interface{}{
				"name": "test-deployment",
			},
		},
	}

	expected := rorresources.NewRorResource("Deployment", "apps/v1")
	expected.SetDeployment(newDeploymentFromDynamicClient(input))
	expected.SetCommonInterface(newDeploymentFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientStorageClass(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "StorageClass",
			"apiVersion": "storage.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "test-storageclass",
			},
		},
	}

	expected := rorresources.NewRorResource("StorageClass", "storage.k8s.io/v1")
	expected.SetStorageClass(newStorageClassFromDynamicClient(input))
	expected.SetCommonInterface(newStorageClassFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientPolicyReport(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "PolicyReport",
			"apiVersion": "wgpolicyk8s.io/v1alpha2",
			"metadata": map[string]interface{}{
				"name": "test-policyreport",
			},
		},
	}

	expected := rorresources.NewRorResource("PolicyReport", "wgpolicyk8s.io/v1alpha2")
	expected.SetPolicyReport(newPolicyReportFromDynamicClient(input))
	expected.SetCommonInterface(newPolicyReportFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientApplication(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Application",
			"apiVersion": "argoproj.io/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-application",
			},
		},
	}

	expected := rorresources.NewRorResource("Application", "argoproj.io/v1alpha1")
	expected.SetApplication(newApplicationFromDynamicClient(input))
	expected.SetCommonInterface(newApplicationFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientAppProject(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "AppProject",
			"apiVersion": "argoproj.io/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-appproject",
			},
		},
	}

	expected := rorresources.NewRorResource("AppProject", "argoproj.io/v1alpha1")
	expected.SetAppProject(newAppProjectFromDynamicClient(input))
	expected.SetCommonInterface(newAppProjectFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientCertificate(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Certificate",
			"apiVersion": "cert-manager.io/v1",
			"metadata": map[string]interface{}{
				"name": "test-certificate",
			},
		},
	}

	expected := rorresources.NewRorResource("Certificate", "cert-manager.io/v1")
	expected.SetCertificate(newCertificateFromDynamicClient(input))
	expected.SetCommonInterface(newCertificateFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientService(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Service",
			"apiVersion": "v1",
			"metadata": map[string]interface{}{
				"name": "test-service",
			},
		},
	}

	expected := rorresources.NewRorResource("Service", "v1")
	expected.SetService(newServiceFromDynamicClient(input))
	expected.SetCommonInterface(newServiceFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientPod(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Pod",
			"apiVersion": "v1",
			"metadata": map[string]interface{}{
				"name": "test-pod",
			},
		},
	}

	expected := rorresources.NewRorResource("Pod", "v1")
	expected.SetPod(newPodFromDynamicClient(input))
	expected.SetCommonInterface(newPodFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientReplicaSet(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ReplicaSet",
			"apiVersion": "apps/v1",
			"metadata": map[string]interface{}{
				"name": "test-replicaset",
			},
		},
	}

	expected := rorresources.NewRorResource("ReplicaSet", "apps/v1")
	expected.SetReplicaSet(newReplicaSetFromDynamicClient(input))
	expected.SetCommonInterface(newReplicaSetFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientStatefulSet(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "StatefulSet",
			"apiVersion": "apps/v1",
			"metadata": map[string]interface{}{
				"name": "test-statefulset",
			},
		},
	}

	expected := rorresources.NewRorResource("StatefulSet", "apps/v1")
	expected.SetStatefulSet(newStatefulSetFromDynamicClient(input))
	expected.SetCommonInterface(newStatefulSetFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientDaemonSet(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "DaemonSet",
			"apiVersion": "apps/v1",
			"metadata": map[string]interface{}{
				"name": "test-daemonset",
			},
		},
	}

	expected := rorresources.NewRorResource("DaemonSet", "apps/v1")
	expected.SetDaemonSet(newDaemonSetFromDynamicClient(input))
	expected.SetCommonInterface(newDaemonSetFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientIngress(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Ingress",
			"apiVersion": "networking.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "test-ingress",
			},
		},
	}

	expected := rorresources.NewRorResource("Ingress", "networking.k8s.io/v1")
	expected.SetIngress(newIngressFromDynamicClient(input))
	expected.SetCommonInterface(newIngressFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientIngressClass(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "IngressClass",
			"apiVersion": "networking.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "test-ingressclass",
			},
		},
	}

	expected := rorresources.NewRorResource("IngressClass", "networking.k8s.io/v1")
	expected.SetIngressClass(newIngressClassFromDynamicClient(input))
	expected.SetCommonInterface(newIngressClassFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientVulnerabilityReport(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "VulnerabilityReport",
			"apiVersion": "aquasecurity.github.io/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-vulnerabilityreport",
			},
		},
	}

	expected := rorresources.NewRorResource("VulnerabilityReport", "aquasecurity.github.io/v1alpha1")
	expected.SetVulnerabilityReport(newVulnerabilityReportFromDynamicClient(input))
	expected.SetCommonInterface(newVulnerabilityReportFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientExposedSecretReport(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ExposedSecretReport",
			"apiVersion": "aquasecurity.github.io/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-exposedsecretreport",
			},
		},
	}

	expected := rorresources.NewRorResource("ExposedSecretReport", "aquasecurity.github.io/v1alpha1")
	expected.SetExposedSecretReport(newExposedSecretReportFromDynamicClient(input))
	expected.SetCommonInterface(newExposedSecretReportFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientConfigAuditReport(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ConfigAuditReport",
			"apiVersion": "aquasecurity.github.io/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-configauditreport",
			},
		},
	}

	expected := rorresources.NewRorResource("ConfigAuditReport", "aquasecurity.github.io/v1alpha1")
	expected.SetConfigAuditReport(newConfigAuditReportFromDynamicClient(input))
	expected.SetCommonInterface(newConfigAuditReportFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientRbacAssessmentReport(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "RbacAssessmentReport",
			"apiVersion": "aquasecurity.github.io/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-rbacassessmentreport",
			},
		},
	}

	expected := rorresources.NewRorResource("RbacAssessmentReport", "aquasecurity.github.io/v1alpha1")
	expected.SetRbacAssessmentReport(newRbacAssessmentReportFromDynamicClient(input))
	expected.SetCommonInterface(newRbacAssessmentReportFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientTanzuKubernetesCluster(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "TanzuKubernetesCluster",
			"apiVersion": "run.tanzu.vmware.com/v1alpha2",
			"metadata": map[string]interface{}{
				"name": "test-tanzukubernetescluster",
			},
		},
	}

	expected := rorresources.NewRorResource("TanzuKubernetesCluster", "run.tanzu.vmware.com/v1alpha2")
	expected.SetTanzuKubernetesCluster(newTanzuKubernetesClusterFromDynamicClient(input))
	expected.SetCommonInterface(newTanzuKubernetesClusterFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientTanzuKubernetesRelease(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "TanzuKubernetesRelease",
			"apiVersion": "run.tanzu.vmware.com/v1alpha2",
			"metadata": map[string]interface{}{
				"name": "test-tanzukubernetesrelease",
			},
		},
	}

	expected := rorresources.NewRorResource("TanzuKubernetesRelease", "run.tanzu.vmware.com/v1alpha2")
	expected.SetTanzuKubernetesRelease(newTanzuKubernetesReleaseFromDynamicClient(input))
	expected.SetCommonInterface(newTanzuKubernetesReleaseFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientVirtualMachineClass(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "VirtualMachineClass",
			"apiVersion": "vmoperator.vmware.com/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-virtualmachineclass",
			},
		},
	}

	expected := rorresources.NewRorResource("VirtualMachineClass", "vmoperator.vmware.com/v1alpha1")
	expected.SetVirtualMachineClass(newVirtualMachineClassFromDynamicClient(input))
	expected.SetCommonInterface(newVirtualMachineClassFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientVirtualMachineClassBinding(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "VirtualMachineClassBinding",
			"apiVersion": "vmoperator.vmware.com/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-virtualmachineclassbinding",
			},
		},
	}

	expected := rorresources.NewRorResource("VirtualMachineClassBinding", "vmoperator.vmware.com/v1alpha1")
	expected.SetVirtualMachineClassBinding(newVirtualMachineClassBindingFromDynamicClient(input))
	expected.SetCommonInterface(newVirtualMachineClassBindingFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientKubernetesCluster(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "KubernetesCluster",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-kubernetescluster",
			},
		},
	}

	expected := rorresources.NewRorResource("KubernetesCluster", "general.ror.internal/v1alpha1")
	expected.SetKubernetesCluster(newKubernetesClusterFromDynamicClient(input))
	expected.SetCommonInterface(newKubernetesClusterFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientClusterOrder(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterOrder",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-clusterorder",
			},
		},
	}

	expected := rorresources.NewRorResource("ClusterOrder", "general.ror.internal/v1alpha1")
	expected.SetClusterOrder(newClusterOrderFromDynamicClient(input))
	expected.SetCommonInterface(newClusterOrderFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientProject(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Project",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-project",
			},
		},
	}

	expected := rorresources.NewRorResource("Project", "general.ror.internal/v1alpha1")
	expected.SetProject(newProjectFromDynamicClient(input))
	expected.SetCommonInterface(newProjectFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientConfiguration(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Configuration",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-configuration",
			},
		},
	}

	expected := rorresources.NewRorResource("Configuration", "general.ror.internal/v1alpha1")
	expected.SetConfiguration(newConfigurationFromDynamicClient(input))
	expected.SetCommonInterface(newConfigurationFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientClusterComplianceReport(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterComplianceReport",
			"apiVersion": "aquasecurity.github.io/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-clustercompliancereport",
			},
		},
	}

	expected := rorresources.NewRorResource("ClusterComplianceReport", "aquasecurity.github.io/v1alpha1")
	expected.SetClusterComplianceReport(newClusterComplianceReportFromDynamicClient(input))
	expected.SetCommonInterface(newClusterComplianceReportFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientClusterVulnerabilityReport(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterVulnerabilityReport",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-clustervulnerabilityreport",
			},
		},
	}

	expected := rorresources.NewRorResource("ClusterVulnerabilityReport", "general.ror.internal/v1alpha1")
	expected.SetClusterVulnerabilityReport(newClusterVulnerabilityReportFromDynamicClient(input))
	expected.SetCommonInterface(newClusterVulnerabilityReportFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientRoute(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Route",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-route",
			},
		},
	}

	expected := rorresources.NewRorResource("Route", "general.ror.internal/v1alpha1")
	expected.SetRoute(newRouteFromDynamicClient(input))
	expected.SetCommonInterface(newRouteFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientSlackMessage(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "SlackMessage",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-slackmessage",
			},
		},
	}

	expected := rorresources.NewRorResource("SlackMessage", "general.ror.internal/v1alpha1")
	expected.SetSlackMessage(newSlackMessageFromDynamicClient(input))
	expected.SetCommonInterface(newSlackMessageFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientVulnerabilityEvent(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "VulnerabilityEvent",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-vulnerabilityevent",
			},
		},
	}

	expected := rorresources.NewRorResource("VulnerabilityEvent", "general.ror.internal/v1alpha1")
	expected.SetVulnerabilityEvent(newVulnerabilityEventFromDynamicClient(input))
	expected.SetCommonInterface(newVulnerabilityEventFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientVm(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Vm",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-vm",
			},
		},
	}

	expected := rorresources.NewRorResource("Vm", "general.ror.internal/v1alpha1")
	expected.SetVm(newVmFromDynamicClient(input))
	expected.SetCommonInterface(newVmFromDynamicClient(input))

	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientWrong(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "N00b",
			"apiVersion": "v900",
			"metadata": map[string]interface{}{
				"name": "test-wrong",
			},
		},
	}
	expected := new(rorresources.ResourceSet)
	result := NewResourceSetFromDynamicClient(input)

	if !cmp.Equal(result.Get(), expected.Get(), cmpopts.EquateEmpty()) {
		t.Errorf("Expected %v, but got %v", expected.Get(), result.Get())
	}
}
