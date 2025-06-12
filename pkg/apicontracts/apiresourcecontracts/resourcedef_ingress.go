package apiresourcecontracts

// K8s namepace struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngress struct {
	ApiVersion string                `json:"apiVersion"`
	Kind       string                `json:"kind"`
	Metadata   ResourceMetadata      `json:"metadata"`
	Spec       ResourceIngressSpec   `json:"spec"`
	Status     ResourceIngressStatus `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressSpec struct {
	DefaultBackend   ResourceIngressSpecRulesHttpPathsBackend `json:"defaultBackend,omitempty"`
	IngressClassName string                                   `json:"ingressClassName"`
	Rules            []ResourceIngressSpecRules               `json:"rules"`
	Tls              []ResourceIngressSpecTls                 `json:"tls"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressSpecRules struct {
	Host string                       `json:"apiGroup"`
	Http ResourceIngressSpecRulesHttp `json:"http"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressSpecRulesHttp struct {
	Paths []ResourceIngressSpecRulesHttpPaths `json:"paths"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressSpecRulesHttpPaths struct {
	Backend  ResourceIngressSpecRulesHttpPathsBackend `json:"backend"`
	Path     string                                   `json:"path"`
	PathType string                                   `json:"pathType"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressSpecRulesHttpPathsBackend struct {
	Resource ResourceIngressSpecBackendResource `json:"resource,omitempty"`
	Service  ResourceIngressSpecBackendService  `json:"service,omitempty"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressSpecBackendResource struct {
	ApiGroup string `json:"apiGroup,omitempty"`
	Kind     string `json:"kind,omitempty"`
	Name     string `json:"name,omitempty"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressSpecBackendService struct {
	Name string                                `json:"name,omitempty"`
	Port ResourceIngressSpecBackendServicePort `json:"port,omitempty"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressSpecBackendServicePort struct {
	Name   string `json:"name,omitempty"`
	Number int    `json:"number,omitempty"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressSpecTls struct {
	Hosts      []string `json:"hosts"`
	SecretName string   `json:"secretName"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressStatus struct {
	LoadBalancer ResourceIngressStatusLoadBalancer `json:"loadBalancer"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressStatusLoadBalancer struct {
	Ingress []ResourceIngressStatusLoadBalancerIngress `json:"ingress"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressStatusLoadBalancerIngress struct {
	Hostname string `json:"hostname"`
	Ip       string `json:"ip"`
}
