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
	Result    SlackMessageResult `json:"result"`
	Timestamp time.Time          `json:"timestamp"`
	Error     any                `json:"error"`
}

type SlackMessageResult int

const (
	SLACK_MESSAGE_OK SlackMessageResult = iota
	SLACK_MESSAGE_ERROR
	SLACK_MESSAGE_UNKNOWN
)
