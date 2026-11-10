package rortypes

import gatewayapiv1 "sigs.k8s.io/gateway-api/apis/v1"

// K8s namepace struct
type ResourceGateway struct {
	Spec   gatewayapiv1.GatewaySpec   `json:"spec"`
	Status gatewayapiv1.GatewayStatus `json:"status"`
}

// (r ResourceIngress) Get returns a pointer to the resource of type ResourceIngress
func (r *ResourceGateway) Get() *ResourceGateway {
	return r
}

// Ingressinterface represents the interface for resources of the type ingress
type Gatewayinterface interface {
	Get() *ResourceGateway
}
