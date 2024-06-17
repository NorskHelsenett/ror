package rortypes

import "time"

type ResourceSlackMessage struct {
	CommonResource `json:",inline"`
	Spec           ResourceSlackMessageSpec     `json:"spec"`
	Status         []ResourceSlackMessageStatus `json:"status"`
}

type ResourceSlackMessageSpec struct {
	ChannelId string `json:"channelId"`
	Message   string `json:"message"`
}

type ResourceSlackMessageStatus struct {
	Result    string    `json:"result"`
	Timestamp time.Time `json:"timestamp"`
	Error     any       `json:"error"`
}
