package rortypes

import gatewayapiv1 "sigs.k8s.io/gateway-api/apis/v1"

// K8s namepace struct
type ResourceHTTPRoute struct {
	Spec   gatewayapiv1.HTTPRouteSpec   `json:"spec"`
	Status gatewayapiv1.HTTPRouteStatus `json:"status"`
}

// (r ResourceIngress) Get returns a pointer to the resource of type ResourceIngress
func (r *ResourceHTTPRoute) Get() *ResourceHTTPRoute {
	return r
}

// Ingressinterface represents the interface for resources of the type ingress
type HTTPRouteinterface interface {
	Get() *ResourceHTTPRoute
}
