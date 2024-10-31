package globalconfig

// Lists the argo Projects that will be marked as internal for use in filtering functions
var InternalAppProjects map[string]bool = map[string]bool{
	"nhn-tooling": true,
}

// Lists the namspaces that will be marked as internal for use in filtering functions
var InternalNamespaces map[string]bool = map[string]bool{
	"argocd":                       true,
	"avi-system":                   true,
	"cert-manager":                 true,
	"jaeger":                       true,
	"kube-node-lease":              true,
	"kube-public":                  true,
	"kube-system":                  true,
	"kyverno":                      true,
	"monitoring":                   true,
	"prometheus-operator":          true,
	"vmware-system-auth":           true,
	"vmware-system-cloud-provider": true,
	"vmware-system-csi":            true,
	"tooling-falco":                true,
	"trident-operator":             true,
	"fluent":                       true,
	"nhn-ror":                      true,
	"ror":                          true,
	"trivy-system":                 true,
	"prometheus-blackbox-exporter": true,
	"eventus":                      true,
}
