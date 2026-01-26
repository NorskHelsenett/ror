package kindclusternamehelper

import "testing"

func TestGetClusternameFromHostname(t *testing.T) {
	testCases := map[string]string{
		"":                          "",
		"kind-control-plane":        "kind",
		"kind-control-plane2":       "kind",
		"kind-worker":               "kind",
		"kind-worker3":              "kind",
		"dev-cluster-control-plane": "dev-cluster",
		"dev-cluster-worker5":       "dev-cluster",
		"single":                    "single",
		"-control-plane":            "",
		"cluster-other":             "cluster-other",
	}

	for hostname, expected := range testCases {
		actual := GetClusternameFromHostname(hostname)
		if actual != expected {
			t.Errorf("hostname %q: expected %q, got %q", hostname, expected, actual)
		}
	}
}
