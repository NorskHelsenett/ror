package k3dclusternamehelper

import "testing"

func TestGetClusternameFromHostname(t *testing.T) {
	testCases := map[string]string{
		"":                                    "",
		"k3d-mycluster-server-0":              "mycluster",
		"k3d-mycluster-agent-1":               "mycluster",
		"k3d-mycluster-serverlb":              "mycluster",
		"k3d-mycluster-control-plane-2":       "mycluster",
		"k3d-mycluster-controlplane-3":        "mycluster",
		"k3d-workspace-dev-mycluster-agent":   "workspace-dev-mycluster",
		"k3d-workspace-dev-mycluster-agent-5": "workspace-dev-mycluster",
		"mycluster-server-0":                  "mycluster",
		"k3d-nonstandard":                     "nonstandard",
		"k3d-":                                "",
		"k3d--server-0":                       "",
	}

	for hostname, expected := range testCases {
		actual := GetClusternameFromHostname(hostname)
		if actual != expected {
			t.Errorf("hostname %q: expected %q, got %q", hostname, expected, actual)
		}
	}
}
