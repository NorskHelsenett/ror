package rortypes

// K8s namepace struct
type ResourceIngressClass struct {
	Spec ResourceIngressClassSpec `json:"spec"`
}

type ResourceIngressClassSpec struct {
	Controller string                             `json:"controller"`
	Parameters ResourceIngressClassSpecParameters `json:"parameters"`
}

type ResourceIngressClassSpecParameters struct {
	ApiGroup  string `json:"apiGroup"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Scope     string `json:"scope"`
}

// (r ResourceIngressClass) Get returns a pointer to the resource of type ResourceIngressClass
func (r *ResourceIngressClass) Get() *ResourceIngressClass {
	return r
}

// IngressClassinterface represents the interface for resources of the type ingressclass
type IngressClassinterface interface {
	Get() *ResourceIngressClass
}
