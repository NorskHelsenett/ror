package apiresourcecontracts

// Resource used by the switchboard microservice to determine notification routing// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceRoute struct {
	ApiVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   ResourceMetadata  `json:"metadata"`
	Spec       ResourceRouteSpec `json:"spec"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceRouteSpec struct {
	MessageType ResourceRouteMessageType `json:"messageType"`
	Receivers   ResourceRouteReceiver    `json:"receivers"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceRouteMessageType struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceRouteReceiver struct {
	Slack []ResourceRouteSlackReceiver `json:"slack"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceRouteSlackReceiver struct {
	ChannelId string `json:"channelId"`
}
