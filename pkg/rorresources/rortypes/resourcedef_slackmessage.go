package rortypes

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ResourceSlackMessage struct {
	Spec   ResourceSlackMessageSpec     `json:"spec"`
	Status []ResourceSlackMessageStatus `json:"status"`
}

type ResourceSlackMessageSpec struct {
	ChannelId string `json:"channelId"`
	Message   string `json:"message"`
}

type ResourceSlackMessageStatus struct {
	Result    ResourceSlackMessageResult `json:"result"`
	Timestamp metav1.Time                `json:"timestamp"`
	Error     any                        `json:"error"`
}

type ResourceSlackMessageResult int

const (
	SLACK_MESSAGE_OK ResourceSlackMessageResult = iota
	SLACK_MESSAGE_ERROR
	SLACK_MESSAGE_UNKNOWN
)

// (r *ResourceSlackMessage) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceSlackMessage) ApplyInputFilter(cr *CommonResource) error {
	return nil
}

// (r ResourceSlackMessage) Get returns a pointer to the resource of type ResourceSlackMessage
func (r *ResourceSlackMessage) Get() *ResourceSlackMessage {
	return r
}
