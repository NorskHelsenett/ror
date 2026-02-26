package vitistackinterregator

import (
	"testing"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func makeNode(annotations, labels map[string]string) v1.Node {
	return v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: annotations,
			Labels:      labels,
		},
	}
}

func validKeys() map[string]string {
	return map[string]string{
		ClusterNameKey:        "test-cluster",
		ClusterWorkspaceKey:   "workspace-1",
		RegionKey:             "region-1",
		AzKey:                 "az-1",
		MachineProviderKey:    "vm-provider",
		KubernetesProviderKey: "kube-provider",
		ClusterIdKey:          "cluster-123",
	}
}

func TestInterregatorNewInterregatorValidAnnotations(t *testing.T) {
	nodes := []v1.Node{makeNode(validKeys(), nil)}

	inter := Interregator{}.NewInterregator(nodes)
	if inter == nil {
		t.Fatalf("expected interregator instance")
	}

	if got := inter.GetClusterName(); got != "test-cluster" {
		t.Fatalf("expected cluster name test-cluster, got %s", got)
	}
	if got := inter.GetClusterWorkspace(); got != "workspace-1" {
		t.Fatalf("expected workspace name workspace-1, got %s", got)
	}
	if got := inter.GetRegion(); got != "region-1" {
		t.Fatalf("expected region name region-1, got %s", got)
	}
	if got := inter.GetAz(); got != "az-1" {
		t.Fatalf("expected az name az-1, got %s", got)
	}
	if got := inter.GetMachineProvider(); got != "vm-provider" {
		t.Fatalf("expected machine provider vm-provider, got %s", got)
	}
	if got := inter.GetKubernetesProvider(); got != "kube-provider" {
		t.Fatalf("expected kubernetes provider kube-provider, got %s", got)
	}
	if got := inter.GetClusterId(); got != "cluster-123" {
		t.Fatalf("expected cluster id cluster-123, got %s", got)
	}
}
func TestInterregatorNewInterregatorValidAnnotationsEmptyNode(t *testing.T) {
	nodes := []v1.Node{
		makeNode(validKeys(), nil),
		makeNode(nil, nil),
	}
	inter := Interregator{}.NewInterregator(nodes)
	if inter == nil {
		t.Fatalf("expected interregator instance")
	}

	if got := inter.GetClusterName(); got != "test-cluster" {
		t.Fatalf("expected cluster name test-cluster, got %s", got)
	}
	if got := inter.GetClusterWorkspace(); got != "workspace-1" {
		t.Fatalf("expected workspace name workspace-1, got %s", got)
	}
	if got := inter.GetRegion(); got != "region-1" {
		t.Fatalf("expected region name region-1, got %s", got)
	}
	if got := inter.GetAz(); got != "az-1" {
		t.Fatalf("expected az name az-1, got %s", got)
	}
	if got := inter.GetMachineProvider(); got != "vm-provider" {
		t.Fatalf("expected machine provider vm-provider, got %s", got)
	}
	if got := inter.GetKubernetesProvider(); got != "kube-provider" {
		t.Fatalf("expected kubernetes provider kube-provider, got %s", got)
	}
	if got := inter.GetClusterId(); got != "cluster-123" {
		t.Fatalf("expected cluster id cluster-123, got %s", got)
	}
}

func TestInterregatorNewInterregatorValidLabels(t *testing.T) {
	nodes := []v1.Node{makeNode(nil, validKeys())}

	inter := Interregator{}.NewInterregator(nodes)
	if inter == nil {
		t.Fatalf("expected interregator instance")
	}

	if got := inter.GetClusterName(); got != "test-cluster" {
		t.Fatalf("expected cluster name test-cluster, got %s", got)
	}
	if got := inter.GetClusterWorkspace(); got != "workspace-1" {
		t.Fatalf("expected workspace name workspace-1, got %s", got)
	}
	if got := inter.GetRegion(); got != "region-1" {
		t.Fatalf("expected region name region-1, got %s", got)
	}
	if got := inter.GetAz(); got != "az-1" {
		t.Fatalf("expected az name az-1, got %s", got)
	}
	if got := inter.GetMachineProvider(); got != "vm-provider" {
		t.Fatalf("expected machine provider vm-provider, got %s", got)
	}
	if got := inter.GetKubernetesProvider(); got != "kube-provider" {
		t.Fatalf("expected kubernetes provider kube-provider, got %s", got)
	}
	if got := inter.GetClusterId(); got != "cluster-123" {
		t.Fatalf("expected cluster id cluster-123, got %s", got)
	}
}

func TestInterregatorNewInterregatorValidLabelsEmptyNode(t *testing.T) {
	nodes := []v1.Node{
		makeNode(validKeys(), nil),
		makeNode(nil, nil),
	}
	inter := Interregator{}.NewInterregator(nodes)
	if inter == nil {
		t.Fatalf("expected interregator instance")
	}

	if got := inter.GetClusterName(); got != "test-cluster" {
		t.Fatalf("expected cluster name test-cluster, got %s", got)
	}
	if got := inter.GetClusterWorkspace(); got != "workspace-1" {
		t.Fatalf("expected workspace name workspace-1, got %s", got)
	}
	if got := inter.GetRegion(); got != "region-1" {
		t.Fatalf("expected region name region-1, got %s", got)
	}
	if got := inter.GetAz(); got != "az-1" {
		t.Fatalf("expected az name az-1, got %s", got)
	}
	if got := inter.GetMachineProvider(); got != "vm-provider" {
		t.Fatalf("expected machine provider vm-provider, got %s", got)
	}
	if got := inter.GetKubernetesProvider(); got != "kube-provider" {
		t.Fatalf("expected kubernetes provider kube-provider, got %s", got)
	}
	if got := inter.GetClusterId(); got != "cluster-123" {
		t.Fatalf("expected cluster id cluster-123, got %s", got)
	}
}

func TestInterregatorNewInterregatorEmptyNode(t *testing.T) {
	nodes := []v1.Node{
		makeNode(nil, nil),
		makeNode(nil, nil),
	}
	inter := Interregator{}.NewInterregator(nodes)
	if inter != nil {
		t.Fatalf("expected nil interregator for invalid nodes")
	}
}

func TestInterregatorNewInterregatorInvalidAnnotations(t *testing.T) {
	annotations := validKeys()
	delete(annotations, ClusterIdKey)
	nodes := []v1.Node{makeNode(annotations, nil)}

	inter := Interregator{}.NewInterregator(nodes)
	if inter != nil {
		t.Fatalf("expected nil interregator for invalid nodes")
	}
}
func TestInterregatorNewInterregatorInvalidLabels(t *testing.T) {
	labels := validKeys()
	delete(labels, ClusterIdKey)
	nodes := []v1.Node{makeNode(nil, labels)}

	inter := Interregator{}.NewInterregator(nodes)
	if inter != nil {
		t.Fatalf("expected nil interregator for invalid nodes")
	}
}
func TestVitistacktypesMustInitialize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		vt   VitistackProviderinterregator
		want bool
	}{
		{
			name: "already typed short-circuits",
			vt:   VitistackProviderinterregator{isOfType: true},
			want: true,
		},
		{
			name: "initialized but not typed stays false",
			vt:   VitistackProviderinterregator{initialized: true},
			want: false,
		},
		{
			name: "valid nodes initialize",
			vt: VitistackProviderinterregator{nodes: []v1.Node{
				makeNode(validKeys(), nil),
			}},
			want: true,
		},
		{
			name: "invalid nodes fail",
			vt: VitistackProviderinterregator{nodes: []v1.Node{
				makeNode(map[string]string{}, nil),
			}},
			want: false,
		},
		{
			name: "mix of nodes picks first valid",
			vt: VitistackProviderinterregator{nodes: []v1.Node{
				makeNode(map[string]string{}, nil),
				makeNode(validKeys(), nil),
			}},
			want: true,
		},
		{
			name: "no nodes returns false",
			vt:   VitistackProviderinterregator{},
			want: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := tc.vt.MustInitialize(); got != tc.want {
				t.Fatalf("MustInitialize() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCheckIfValid(t *testing.T) {
	vt := VitistackProviderinterregator{}
	node := makeNode(validKeys(), nil)
	if !vt.checkIfValid(&node) {
		t.Fatalf("expected node with annotations to be valid")
	}

	labels := map[string]string{}
	for _, key := range MustBeSet {
		labels[key] = "value"
	}
	nodeWithLabels := makeNode(nil, labels)
	if !vt.checkIfValid(&nodeWithLabels) {
		t.Fatalf("expected node with labels to be valid")
	}

	missing := makeNode(map[string]string{ClusterNameKey: "name"}, nil)
	if vt.checkIfValid(&missing) {
		t.Fatalf("expected node missing keys to be invalid")
	}
}

func TestCheckIfKeyPresent(t *testing.T) {
	node := makeNode(map[string]string{ClusterNameKey: "name"}, nil)
	if !checkIfKeyPresent(&node, ClusterNameKey) {
		t.Fatalf("expected annotation key to be present")
	}

	nodeLabels := makeNode(nil, map[string]string{ClusterNameKey: "name"})
	if !checkIfKeyPresent(&nodeLabels, ClusterNameKey) {
		t.Fatalf("expected label key to be present")
	}

	empty := makeNode(nil, nil)
	if checkIfKeyPresent(&empty, ClusterNameKey) {
		t.Fatalf("expected missing key to return false")
	}
}

func TestGetValueByKey(t *testing.T) {
	node := makeNode(map[string]string{RegionKey: "annotation"}, nil)
	if val, _ := getValueByKey(&node, RegionKey); val != "annotation" {
		t.Fatalf("expected annotation value, got %s", val)
	}

	nodeLabel := makeNode(nil, map[string]string{RegionKey: "label"})
	if val, _ := getValueByKey(&nodeLabel, RegionKey); val != "label" {
		t.Fatalf("expected label value, got %s", val)
	}

	empty := makeNode(nil, nil)
	if val, _ := getValueByKey(&empty, RegionKey); val != "" {
		t.Fatalf("expected empty string, got %s", val)
	}
}

func TestIsTypeOf(t *testing.T) {
	vt := VitistackProviderinterregator{isOfType: true}
	if !vt.IsTypeOf() {
		t.Fatalf("expected true when already typed")
	}

	invalid := VitistackProviderinterregator{nodes: []v1.Node{makeNode(map[string]string{}, nil)}}
	if invalid.IsTypeOf() {
		t.Fatalf("expected false for non vitistack nodes")
	}
}

func TestGetProvider(t *testing.T) {
	vt := VitistackProviderinterregator{nodes: []v1.Node{}}
	if got := vt.GetProvider(); got != providermodels.ProviderTypeUnknown {
		t.Fatalf("expected unknown provider, got %v", got)
	}

	valid := VitistackProviderinterregator{nodes: []v1.Node{makeNode(validKeys(), nil)}}
	if got := valid.GetProvider(); got != providermodels.ProviderTypeVitistack {
		t.Fatalf("expected vitistack provider, got %v", got)
	}
}

func TestGetClusterId(t *testing.T) {
	node := makeNode(validKeys(), nil)
	vt := VitistackProviderinterregator{nodes: []v1.Node{node}}
	if got := vt.GetClusterId(); got != "cluster-123" {
		t.Fatalf("expected cluster id cluster-123, got %s", got)
	}

	vtMissing := VitistackProviderinterregator{nodes: []v1.Node{makeNode(map[string]string{}, nil)}}
	if got := vtMissing.GetClusterId(); got != providermodels.UNKNOWN_CLUSTER_ID {
		t.Fatalf("expected unknown cluster id, got %s", got)
	}
}

func TestGetClusterName(t *testing.T) {
	node := makeNode(validKeys(), nil)
	vt := VitistackProviderinterregator{nodes: []v1.Node{node}}
	if got := vt.GetClusterName(); got != "test-cluster" {
		t.Fatalf("expected cluster name test-cluster, got %s", got)
	}

	vtMissing := VitistackProviderinterregator{nodes: []v1.Node{makeNode(map[string]string{}, nil)}}
	if got := vtMissing.GetClusterName(); got != providermodels.UNKNOWN_CLUSTER {
		t.Fatalf("expected unknown cluster name, got %s", got)
	}
}

func TestGetClusterWorkspace(t *testing.T) {
	node := makeNode(validKeys(), nil)
	vt := VitistackProviderinterregator{nodes: []v1.Node{node}}
	if got := vt.GetClusterWorkspace(); got != "workspace-1" {
		t.Fatalf("expected workspace workspace-1, got %s", got)
	}

	vtMissing := VitistackProviderinterregator{nodes: []v1.Node{makeNode(map[string]string{}, nil)}}
	if got := vtMissing.GetClusterWorkspace(); got != "Vitistack" {
		t.Fatalf("expected default workspace Vitistack, got %s", got)
	}
}

func TestGetDatacenter(t *testing.T) {
	node := makeNode(validKeys(), nil)
	vt := VitistackProviderinterregator{nodes: []v1.Node{node}}
	if got := vt.GetDatacenter(); got != "az-1.region-1.no" {
		t.Fatalf("expected datacenter az-1.region-1.no, got %s", got)
	}

	vtMissing := VitistackProviderinterregator{nodes: []v1.Node{makeNode(map[string]string{}, nil)}}
	expected := providermodels.UNKNOWN_DATACENTER
	if got := vtMissing.GetDatacenter(); got != expected {
		t.Fatalf("expected unknown datacenter, got %s", got)
	}
}

func TestGetRegion(t *testing.T) {
	node := makeNode(validKeys(), nil)
	vt := VitistackProviderinterregator{nodes: []v1.Node{node}}
	if got := vt.GetRegion(); got != "region-1" {
		t.Fatalf("expected region region-1, got %s", got)
	}

	vtMissing := VitistackProviderinterregator{nodes: []v1.Node{makeNode(map[string]string{}, nil)}}
	if got := vtMissing.GetRegion(); got != providermodels.UNKNOWN_REGION {
		t.Fatalf("expected unknown region, got %s", got)
	}
}

func TestGetAz(t *testing.T) {
	node := makeNode(validKeys(), nil)
	vt := VitistackProviderinterregator{nodes: []v1.Node{node}}
	if got := vt.GetAz(); got != "az-1" {
		t.Fatalf("expected az az-1, got %s", got)
	}

	vtMissing := VitistackProviderinterregator{nodes: []v1.Node{makeNode(map[string]string{}, nil)}}
	if got := vtMissing.GetAz(); got != providermodels.UNKNOWN_AZ {
		t.Fatalf("expected unknown az, got %s", got)
	}
}

func TestGetMachineProvider(t *testing.T) {
	node := makeNode(validKeys(), nil)
	vt := VitistackProviderinterregator{nodes: []v1.Node{node}}
	if got := vt.GetMachineProvider(); got != "vm-provider" {
		t.Fatalf("expected vm provider vm-provider, got %s", got)
	}

	vtMissing := VitistackProviderinterregator{nodes: []v1.Node{makeNode(map[string]string{}, nil)}}
	if got := vtMissing.GetMachineProvider(); got != providermodels.ProviderTypeUnknown {
		t.Fatalf("expected unknown vm provider, got %s", got)
	}
}

func TestGetKubernetesProvider(t *testing.T) {
	node := makeNode(validKeys(), nil)
	vt := VitistackProviderinterregator{nodes: []v1.Node{node}}
	if got := vt.GetKubernetesProvider(); got != "kube-provider" {
		t.Fatalf("expected kubernetes provider kube-provider, got %s", got)
	}

	vtMissing := VitistackProviderinterregator{nodes: []v1.Node{makeNode(map[string]string{}, nil)}}
	if got := vtMissing.GetKubernetesProvider(); got != providermodels.ProviderTypeUnknown {
		t.Fatalf("expected unknown kubernetes provider, got %s", got)
	}
}
