package rortypes

// Resource used by the switchboard microservice to determine notification routing
type ResourceRoute struct {
	Spec ResourceRouteSpec `json:"spec"`
}

type ResourceRouteSpec struct {
	MessageType ResourceRouteMessageType `json:"messageType"`
	Receivers   ResourceRouteReceiver    `json:"receivers"`
}

type ResourceRouteMessageType struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

type ResourceRouteReceiver struct {
	Slack []ResourceRouteSlackReceiver `json:"slack"`
}

type ResourceRouteSlackReceiver struct {
	ChannelId string `json:"channelId"`
}

// (r ResourceRoute) Get returns a pointer to the resource of type ResourceRoute
func (r *ResourceRoute) Get() *ResourceRoute {
	return r
}

// Routeinterface represents the interface for resources of the type route
type Routeinterface interface {
	Get() *ResourceRoute
}
