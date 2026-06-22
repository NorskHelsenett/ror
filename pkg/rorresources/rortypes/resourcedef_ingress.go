package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

// K8s namepace struct
type ResourceIngress struct {
	Spec   ResourceIngressSpec   `json:"spec"`
	Status ResourceIngressStatus `json:"status"`
}

type ResourceIngressSpec struct {
	DefaultBackend   ResourceIngressSpecRulesHttpPathsBackend `json:"defaultBackend"`
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
	Resource ResourceIngressSpecBackendResource `json:"resource"`
	Service  ResourceIngressSpecBackendService  `json:"service"`
}

type ResourceIngressSpecBackendResource struct {
	ApiGroup string `json:"apiGroup,omitempty"`
	Kind     string `json:"kind,omitempty"`
	Name     string `json:"name,omitempty"`
}

type ResourceIngressSpecBackendService struct {
	Name string                                `json:"name,omitempty"`
	Port ResourceIngressSpecBackendServicePort `json:"port"`
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

// (r *ResourceIngress) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceIngress) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Namespace] {
		cr.RorMeta.Internal = true
	}
	return nil
}

// (r ResourceIngress) Get returns a pointer to the resource of type ResourceIngress
func (r *ResourceIngress) Get() *ResourceIngress {
	return r
}

// Ingressinterface represents the interface for resources of the type ingress
type Ingressinterface interface {
	ApplyInputFilter(cr *CommonResource) error
	Get() *ResourceIngress
}
