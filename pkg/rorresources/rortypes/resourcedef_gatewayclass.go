package rortypes

import gatewayapiv1 "sigs.k8s.io/gateway-api/apis/v1"

// K8s namepace struct
type ResourceGatewayClass struct {
	Spec   gatewayapiv1.GatewayClassSpec   `json:"spec"`
	Status gatewayapiv1.GatewayClassStatus `json:"status"`
}

// (r ResourceIngress) Get returns a pointer to the resource of type ResourceIngress
func (r *ResourceGatewayClass) Get() *ResourceGatewayClass {
	return r
}

// Ingressinterface represents the interface for resources of the type ingress
type GatewayClassinterface interface {
	Get() *ResourceGatewayClass
}
