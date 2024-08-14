package rortypes

// K8s namepace struct
type ResourceIngress struct {
	Spec   ResourceIngressSpec   `json:"spec"`
	Status ResourceIngressStatus `json:"status"`
}

type ResourceIngressSpec struct {
	DefaultBackend   ResourceIngressSpecRulesHttpPathsBackend `json:"defaultBackend,omitempty"`
	IngressClassName string                                   `json:"ingressClassName"`
	Rules            []ResourceIngressSpecRules               `json:"rules"`
	Tls              []ResourceIngressSpecTls                 `json:"tls"`
}

type ResourceIngressSpecRules struct {
	Host string                       `json:"apiGroup"`
	Http ResourceIngressSpecRulesHttp `json:"http"`
}

type ResourceIngressSpecRulesHttp struct {
	Paths []ResourceIngressSpecRulesHttpPaths `json:"paths"`
}

type ResourceIngressSpecRulesHttpPaths struct {
	Backend  ResourceIngressSpecRulesHttpPathsBackend `json:"backend"`
	Path     string                                   `json:"path"`
	PathType string                                   `json:"pathType"`
}

type ResourceIngressSpecRulesHttpPathsBackend struct {
	Resource ResourceIngressSpecBackendResource `json:"resource,omitempty"`
	Service  ResourceIngressSpecBackendService  `json:"service,omitempty"`
}

type ResourceIngressSpecBackendResource struct {
	ApiGroup string `json:"apiGroup,omitempty"`
	Kind     string `json:"kind,omitempty"`
	Name     string `json:"name,omitempty"`
}

type ResourceIngressSpecBackendService struct {
	Name string                                `json:"name,omitempty"`
	Port ResourceIngressSpecBackendServicePort `json:"port,omitempty"`
}

type ResourceIngressSpecBackendServicePort struct {
	Name   string `json:"name,omitempty"`
	Number int    `json:"number,omitempty"`
}

type ResourceIngressSpecTls struct {
	Hosts      []string `json:"hosts"`
	SecretName string   `json:"secretName"`
}

type ResourceIngressStatus struct {
	LoadBalancer ResourceIngressStatusLoadBalancer `json:"loadBalancer"`
}

type ResourceIngressStatusLoadBalancer struct {
	Ingress []ResourceIngressStatusLoadBalancerIngress `json:"ingress"`
}

type ResourceIngressStatusLoadBalancerIngress struct {
	Hostname string `json:"hostname"`
	Ip       string `json:"ip"`
}
