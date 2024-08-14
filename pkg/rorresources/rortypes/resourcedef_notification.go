package rortypes

type ResourceNotification struct {
	Spec ResourceNotificationSpec `json:"spec"`
}

type ResourceNotificationSpec struct {
	Owner   RorResourceOwnerReference `json:"owner"`
	Message string                    `json:"message"`
}
