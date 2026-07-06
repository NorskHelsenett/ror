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
		Object: map[string]any{
			"kind":       "Namespace",
			"apiVersion": "v1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "Node",
			"apiVersion": "v1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "PersistentVolumeClaim",
			"apiVersion": "v1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "Deployment",
			"apiVersion": "apps/v1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "StorageClass",
			"apiVersion": "storage.k8s.io/v1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "PolicyReport",
			"apiVersion": "wgpolicyk8s.io/v1alpha2",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "Application",
			"apiVersion": "argoproj.io/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "AppProject",
			"apiVersion": "argoproj.io/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "Certificate",
			"apiVersion": "cert-manager.io/v1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "Service",
			"apiVersion": "v1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "Pod",
			"apiVersion": "v1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "ReplicaSet",
			"apiVersion": "apps/v1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "StatefulSet",
			"apiVersion": "apps/v1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "DaemonSet",
			"apiVersion": "apps/v1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "Ingress",
			"apiVersion": "networking.k8s.io/v1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "IngressClass",
			"apiVersion": "networking.k8s.io/v1",
			"metadata": map[string]any{
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

func TestNewResourceSetFromDynamicClientSbomReport(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]any{
			"kind":       "SbomReport",
			"apiVersion": "aquasecurity.github.io/v1alpha1",
			"metadata": map[string]any{
				"name": "test-sbomreport",
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
		Object: map[string]any{
			"kind":       "VulnerabilityReport",
			"apiVersion": "aquasecurity.github.io/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "ExposedSecretReport",
			"apiVersion": "aquasecurity.github.io/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "ConfigAuditReport",
			"apiVersion": "aquasecurity.github.io/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "RbacAssessmentReport",
			"apiVersion": "aquasecurity.github.io/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "TanzuKubernetesCluster",
			"apiVersion": "run.tanzu.vmware.com/v1alpha3",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "TanzuKubernetesRelease",
			"apiVersion": "run.tanzu.vmware.com/v1alpha3",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "VirtualMachineClass",
			"apiVersion": "vmoperator.vmware.com/v1alpha2",
			"metadata": map[string]any{
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

func TestNewResourceSetFromDynamicClientKubernetesCluster(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]any{
			"kind":       "KubernetesCluster",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]any{
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

func TestNewResourceSetFromDynamicClientProvider(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]any{
			"kind":       "Provider",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]any{
				"name": "test-provider",
			},
		},
	}

	expected := NewResourceFromDynamicClient(input)
	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientWorkspace(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]any{
			"kind":       "Workspace",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]any{
				"name": "test-workspace",
			},
		},
	}

	expected := NewResourceFromDynamicClient(input)
	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientKubernetesMachineClass(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]any{
			"kind":       "KubernetesMachineClass",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]any{
				"name": "test-kubernetesmachineclass",
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
		Object: map[string]any{
			"kind":       "ClusterOrder",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "Project",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "Configuration",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "ClusterComplianceReport",
			"apiVersion": "aquasecurity.github.io/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "ClusterVulnerabilityReport",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "Route",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "SlackMessage",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "VulnerabilityEvent",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "VirtualMachine",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]any{
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

func TestNewResourceSetFromDynamicClientVirtualMachineVulnerabilityInfo(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]any{
			"kind":       "VirtualMachineVulnerabilityInfo",
			"apiVersion": "general.ror.internal/v1alpha1",
			"metadata": map[string]any{
				"name": "test-virtualmachinevulnerabilityinfo",
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
		Object: map[string]any{
			"kind":       "Endpoints",
			"apiVersion": "v1",
			"metadata": map[string]any{
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
		Object: map[string]any{
			"kind":       "NetworkPolicy",
			"apiVersion": "networking.k8s.io/v1",
			"metadata": map[string]any{
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

func TestNewResourceSetFromDynamicClientDatacenter(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]any{
			"kind":       "Datacenter",
			"apiVersion": "infrastructure.ror.internal/v1alpha1",
			"metadata": map[string]any{
				"name": "test-datacenter",
			},
		},
	}

	expected := NewResourceFromDynamicClient(input)
	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientBackupJob(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]any{
			"kind":       "BackupJob",
			"apiVersion": "backup.ror.internal/v1alpha1",
			"metadata": map[string]any{
				"name": "test-backupjob",
			},
		},
	}

	expected := NewResourceFromDynamicClient(input)
	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientBackupRun(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]any{
			"kind":       "BackupRun",
			"apiVersion": "backup.ror.internal/v1alpha1",
			"metadata": map[string]any{
				"name": "test-backuprun",
			},
		},
	}

	expected := NewResourceFromDynamicClient(input)
	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientMachine(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]any{
			"kind":       "Machine",
			"apiVersion": "machine.ror.internal/v1alpha1",
			"metadata": map[string]any{
				"name": "test-machine",
			},
		},
	}

	expected := NewResourceFromDynamicClient(input)
	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientConfig(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]any{
			"kind":       "Config",
			"apiVersion": "ror.internal/v1",
			"metadata": map[string]any{
				"name": "test-config",
			},
		},
	}

	expected := NewResourceFromDynamicClient(input)
	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientOrganizationalUnit(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]any{
			"kind":       "OrganizationalUnit",
			"apiVersion": "ror.internal/v1",
			"metadata": map[string]any{
				"name": "test-organizationalunit",
			},
		},
	}

	expected := NewResourceFromDynamicClient(input)
	result := NewResourceSetFromDynamicClient(input)

	if !reflect.DeepEqual(result.Get(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNewResourceSetFromDynamicClientUnknown(t *testing.T) {
	input := &unstructured.Unstructured{
		Object: map[string]any{
			"kind":       "Unknown",
			"apiVersion": "unknown.ror.internal/v1",
			"metadata": map[string]any{
				"name": "test-unknown",
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
		Object: map[string]any{
			"kind":       "N00b",
			"apiVersion": "v900",
			"metadata": map[string]any{
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
