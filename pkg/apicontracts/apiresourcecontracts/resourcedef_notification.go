package apiresourcecontracts

import "github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

type ResourceNotification struct {
	ApiVersion string                   `json:"apiVersion"`
	Kind       string                   `json:"kind"`
	Metadata   ResourceMetadata         `json:"metadata"`
	Spec       ResourceNotificationSpec `json:"spec"`
}

type ResourceNotificationSpec struct {
	Owner   rortypes.RorResourceOwnerReference `json:"owner"`
	Message string                             `json:"message"`
}
