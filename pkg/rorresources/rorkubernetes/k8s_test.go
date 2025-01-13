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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
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

	expected := NewResourceFromDynamicClient(input)
	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientVirtualMachine(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "VirtualMachine",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]interface{}{
				"name": "test-virtualmachine",
			},
		},
	}

	expected := NewResourceFromDynamicClient(input)
	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientEndpoints(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Endpoints",
			"apiVersion": "v1",
			"metadata": map[string]interface{}{
				"name": "test-endpoints",
			},
		},
	}

	expected := NewResourceFromDynamicClient(input)
	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientNetworkPolicy(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "NetworkPolicy",
			"apiVersion": "networking.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "test-networkpolicy",
			},
		},
	}

	expected := NewResourceFromDynamicClient(input)
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
