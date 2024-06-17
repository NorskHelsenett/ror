package rortypes

type ResourceNotification struct {
	CommonResource `json:",inline"`
	Spec           ResourceNotificationSpec `json:"spec"`
}

type ResourceNotificationSpec struct {
	Owner   RorResourceOwnerReference `json:"owner"`
	Message string                    `json:"message"`
}
