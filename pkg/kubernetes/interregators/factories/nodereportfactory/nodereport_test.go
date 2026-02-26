package nodereportfactory

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func makeNodes() []v1.Node {
	return []v1.Node{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "node-a",
				UID:  "uid-a",
				Labels: map[string]string{
					"kubernetes.io/hostname": "hostname-a",
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "node-b",
				UID:  "uid-b",
				Labels: map[string]string{
					"kubernetes.io/hostname": "hostname-b",
				},
			},
		},
	}
}

func TestGetByName_ReturnsPointerToSliceElement(t *testing.T) {
	nodes := makeNodes()
	r := NodeReport{nodes: nodes}

	got := r.GetByName("node-a")
	if got == nil {
		t.Fatal("expected non-nil result")
	}
	if got.Name != "node-a" {
		t.Errorf("expected node-a, got %s", got.Name)
	}
	// Verify pointer points into the slice, not a copy
	if got != &r.nodes[0] {
		t.Error("GetByName returned a pointer to a copy, not to the slice element")
	}
}

func TestGetByUid_ReturnsPointerToSliceElement(t *testing.T) {
	nodes := makeNodes()
	r := NodeReport{nodes: nodes}

	got := r.GetByUid("uid-b")
	if got == nil {
		t.Fatal("expected non-nil result")
	}
	if string(got.UID) != "uid-b" {
		t.Errorf("expected uid-b, got %s", got.UID)
	}
	if got != &r.nodes[1] {
		t.Error("GetByUid returned a pointer to a copy, not to the slice element")
	}
}

func TestGetByHostname_ReturnsPointerToSliceElement(t *testing.T) {
	nodes := makeNodes()
	r := NodeReport{nodes: nodes}

	got := r.GetByHostname("hostname-a")
	if got == nil {
		t.Fatal("expected non-nil result")
	}
	if got.Labels["kubernetes.io/hostname"] != "hostname-a" {
		t.Errorf("expected hostname-a, got %s", got.Labels["kubernetes.io/hostname"])
	}
	if got != &r.nodes[0] {
		t.Error("GetByHostname returned a pointer to a copy, not to the slice element")
	}
}

func TestGetByName_NotFound(t *testing.T) {
	r := NodeReport{nodes: makeNodes()}
	if r.GetByName("missing") != nil {
		t.Error("expected nil for missing node")
	}
}

func TestGetByUid_NotFound(t *testing.T) {
	r := NodeReport{nodes: makeNodes()}
	if r.GetByUid("missing") != nil {
		t.Error("expected nil for missing uid")
	}
}

func TestGetByHostname_NotFound(t *testing.T) {
	r := NodeReport{nodes: makeNodes()}
	if r.GetByHostname("missing") != nil {
		t.Error("expected nil for missing hostname")
	}
}
