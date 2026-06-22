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

// newTestInterregator creates a VitistackProviderinterregator from nodes for testing,
// bypassing the Kubernetes client requirement of NewInterregator.
// Returns nil if the nodes don't satisfy vitistack requirements.
func newTestInterregator(nodes []v1.Node) *VitistackProviderinterregator {
	vt := &VitistackProviderinterregator{nodes: nodes}
	if !vt.DetectProvider() {
		return nil
	}
	return vt
}

func TestInterregatorNewInterregatorValidAnnotations(t *testing.T) {
	nodes := []v1.Node{makeNode(validKeys(), nil)}

	inter := newTestInterregator(nodes)
	if inter == nil {
		t.Fatalf("expected interregator instance")
	}

	if got := inter.GetProvider(); got != providermodels.ProviderTypeVitistack {
		t.Fatalf("expected provider vitistack, got %v", got)
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
	if got := inter.GetCountry(); got != providermodels.DefaultCountry {
		t.Fatalf("expected default country %s, got %s", providermodels.DefaultCountry, got)
	}
	if got := inter.GetDatacenter(); got != "az-1.region-1."+providermodels.DefaultCountry {
		t.Fatalf("expected datacenter az-1.region-1.%s, got %s", providermodels.DefaultCountry, got)
	}
}
func TestInterregatorNewInterregatorValidAnnotationsEmptyNode(t *testing.T) {
	nodes := []v1.Node{
		makeNode(validKeys(), nil),
		makeNode(nil, nil),
	}
	inter := newTestInterregator(nodes)
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

	inter := newTestInterregator(nodes)
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
	inter := newTestInterregator(nodes)
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
	inter := newTestInterregator(nodes)
	if inter != nil {
		t.Fatalf("expected nil interregator for invalid nodes")
	}
}

func TestInterregatorNewInterregatorInvalidAnnotations(t *testing.T) {
	annotations := validKeys()
	delete(annotations, ClusterIdKey)
	nodes := []v1.Node{makeNode(annotations, nil)}

	inter := newTestInterregator(nodes)
	if inter != nil {
		t.Fatalf("expected nil interregator for invalid nodes")
	}
}
func TestInterregatorNewInterregatorInvalidLabels(t *testing.T) {
	labels := validKeys()
	delete(labels, ClusterIdKey)
	nodes := []v1.Node{makeNode(nil, labels)}

	inter := newTestInterregator(nodes)
	if inter != nil {
		t.Fatalf("expected nil interregator for invalid nodes")
	}
}
func TestDetectProvider(t *testing.T) {
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
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := tc.vt.DetectProvider(); got != tc.want {
				t.Fatalf("DetectProvider() = %v, want %v", got, tc.want)
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
	for _, key := range mustBeSet {
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

func TestGetProvider(t *testing.T) {
	vt := newTestInterregator([]v1.Node{makeNode(validKeys(), nil)})
	if vt == nil {
		t.Fatalf("expected interregator instance")
	}
	if got := vt.GetProvider(); got != providermodels.ProviderTypeVitistack {
		t.Fatalf("expected vitistack provider, got %v", got)
	}
}

func TestGetClusterId(t *testing.T) {
	vt := newTestInterregator([]v1.Node{makeNode(validKeys(), nil)})
	if vt == nil {
		t.Fatalf("expected interregator instance")
	}
	if got := vt.GetClusterId(); got != "cluster-123" {
		t.Fatalf("expected cluster id cluster-123, got %s", got)
	}
}

func TestGetClusterName(t *testing.T) {
	vt := newTestInterregator([]v1.Node{makeNode(validKeys(), nil)})
	if vt == nil {
		t.Fatalf("expected interregator instance")
	}
	if got := vt.GetClusterName(); got != "test-cluster" {
		t.Fatalf("expected cluster name test-cluster, got %s", got)
	}
}

func TestGetClusterWorkspace(t *testing.T) {
	vt := newTestInterregator([]v1.Node{makeNode(validKeys(), nil)})
	if vt == nil {
		t.Fatalf("expected interregator instance")
	}
	if got := vt.GetClusterWorkspace(); got != "workspace-1" {
		t.Fatalf("expected workspace workspace-1, got %s", got)
	}
}

func TestGetDatacenter(t *testing.T) {
	vt := newTestInterregator([]v1.Node{makeNode(validKeys(), nil)})
	if vt == nil {
		t.Fatalf("expected interregator instance")
	}
	if got := vt.GetDatacenter(); got != "az-1.region-1.no" {
		t.Fatalf("expected datacenter az-1.region-1.no, got %s", got)
	}
}

func TestGetRegion(t *testing.T) {
	vt := newTestInterregator([]v1.Node{makeNode(validKeys(), nil)})
	if vt == nil {
		t.Fatalf("expected interregator instance")
	}
	if got := vt.GetRegion(); got != "region-1" {
		t.Fatalf("expected region region-1, got %s", got)
	}
}

func TestGetAz(t *testing.T) {
	vt := newTestInterregator([]v1.Node{makeNode(validKeys(), nil)})
	if vt == nil {
		t.Fatalf("expected interregator instance")
	}
	if got := vt.GetAz(); got != "az-1" {
		t.Fatalf("expected az az-1, got %s", got)
	}
}

func TestGetMachineProvider(t *testing.T) {
	vt := newTestInterregator([]v1.Node{makeNode(validKeys(), nil)})
	if vt == nil {
		t.Fatalf("expected interregator instance")
	}
	if got := vt.GetMachineProvider(); got != "vm-provider" {
		t.Fatalf("expected vm provider vm-provider, got %s", got)
	}
}

func TestGetKubernetesProvider(t *testing.T) {
	vt := newTestInterregator([]v1.Node{makeNode(validKeys(), nil)})
	if vt == nil {
		t.Fatalf("expected interregator instance")
	}
	if got := vt.GetKubernetesProvider(); got != "kube-provider" {
		t.Fatalf("expected kubernetes provider kube-provider, got %s", got)
	}
}

func TestGetCountry(t *testing.T) {
	// CountryKey is optional; when missing it should default.
	vtDefault := newTestInterregator([]v1.Node{makeNode(validKeys(), nil)})
	if vtDefault == nil {
		t.Fatalf("expected interregator instance")
	}
	if got := vtDefault.GetCountry(); got != providermodels.DefaultCountry {
		t.Fatalf("expected default country %s, got %s", providermodels.DefaultCountry, got)
	}

	withCountry := validKeys()
	withCountry[CountryKey] = "se"
	vt := newTestInterregator([]v1.Node{makeNode(withCountry, nil)})
	if vt == nil {
		t.Fatalf("expected interregator instance")
	}
	if got := vt.GetCountry(); got != "se" {
		t.Fatalf("expected country se, got %s", got)
	}
}

func TestInterregatorNewInterregatorNoNodesSlice(t *testing.T) {
	inter := newTestInterregator(nil)
	if inter != nil {
		t.Fatalf("expected nil interregator for nil nodes slice")
	}
}

func TestDetectProvider_SetsStateAndFields(t *testing.T) {
	vt := &VitistackProviderinterregator{nodes: []v1.Node{makeNode(validKeys(), nil)}}
	if got := vt.DetectProvider(); !got {
		t.Fatalf("expected DetectProvider to return true")
	}
	if !vt.initialized {
		t.Fatalf("expected initialized=true")
	}
	if !vt.isOfType {
		t.Fatalf("expected isOfType=true")
	}
	if vt.clustername != "test-cluster" {
		t.Fatalf("expected clustername test-cluster, got %s", vt.clustername)
	}
	if vt.clusterworkspace != "workspace-1" {
		t.Fatalf("expected clusterworkspace workspace-1, got %s", vt.clusterworkspace)
	}
	if vt.region != "region-1" {
		t.Fatalf("expected region region-1, got %s", vt.region)
	}
	if vt.az != "az-1" {
		t.Fatalf("expected az az-1, got %s", vt.az)
	}
	if vt.machineprovider != "vm-provider" {
		t.Fatalf("expected machineprovider vm-provider, got %s", vt.machineprovider)
	}
	if vt.kubernetesprovider != "kube-provider" {
		t.Fatalf("expected kubernetesprovider kube-provider, got %s", vt.kubernetesprovider)
	}
	if vt.clusterId != "cluster-123" {
		t.Fatalf("expected clusterId cluster-123, got %s", vt.clusterId)
	}
}

func TestDetectProvider_InvalidNodesMarksInitialized(t *testing.T) {
	vt := &VitistackProviderinterregator{nodes: []v1.Node{makeNode(map[string]string{}, nil)}}
	if got := vt.DetectProvider(); got {
		t.Fatalf("expected DetectProvider to return false")
	}
	if !vt.initialized {
		t.Fatalf("expected initialized=true")
	}
	if vt.isOfType {
		t.Fatalf("expected isOfType=false")
	}
}
