package apiresourcecontracts

// Resource used by the switchboard microservice to determine notification routing
type ResourceRoute struct {
	ApiVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   ResourceMetadata  `json:"metadata"`
	Spec       ResourceRouteSpec `json:"spec"`
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
