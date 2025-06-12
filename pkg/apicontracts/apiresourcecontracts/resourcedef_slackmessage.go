package apiresourcecontracts

import "time"

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceSlackMessage struct {
	ApiVersion string                       `json:"apiVersion"`
	Kind       string                       `json:"kind"`
	Metadata   ResourceMetadata             `json:"metadata"`
	Spec       ResourceSlackMessageSpec     `json:"spec"`
	Status     []ResourceSlackMessageStatus `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceSlackMessageSpec struct {
	ChannelId string `json:"channelId"`
	Message   string `json:"message"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceSlackMessageStatus struct {
	Result    ResourceSlackMessageResult `json:"result"`
	Timestamp time.Time                  `json:"timestamp"`
	Error     any                        `json:"error"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceSlackMessageResult int

const (
	SLACK_MESSAGE_OK ResourceSlackMessageResult = iota
	SLACK_MESSAGE_ERROR
	SLACK_MESSAGE_UNKNOWN
)
