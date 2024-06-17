package apiresourcecontracts

import "time"

type ResourceSlackMessage struct {
	ApiVersion string                       `json:"apiVersion"`
	Kind       string                       `json:"kind"`
	Metadata   ResourceMetadata             `json:"metadata"`
	Spec       ResourceSlackMessageSpec     `json:"spec"`
	Status     []ResourceSlackMessageStatus `json:"status"`
}

type ResourceSlackMessageSpec struct {
	ChannelId string `json:"channelId"`
	Message   string `json:"message"`
}

type ResourceSlackMessageStatus struct {
	Result    ResourceSlackMessageResult `json:"result"`
	Timestamp time.Time                  `json:"timestamp"`
	Error     any                        `json:"error"`
}

type ResourceSlackMessageResult int

const (
	SLACK_MESSAGE_OK ResourceSlackMessageResult = iota
	SLACK_MESSAGE_ERROR
	SLACK_MESSAGE_UNKNOWN
)
