package rortypes

type ResourceNotification struct {
	CommonResource `json:",inline"`
	Spec           ResourceNotificationSpec `json:"spec"`
}

type ResourceNotificationSpec struct {
	Message string `json:"message"`
}
