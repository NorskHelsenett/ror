package models

import (
	"time"
)

type Event struct {
	Id         string    `json:"id"` //mongodb id (auto magical)
	Created    time.Time `json:"created,omitempty"`
	Updated    time.Time `json:"updated,omitempty"`
	Type       string    `json:"type"`  //, sos, alert, security
	Scope      string    `json:"scope"` //(cluster, namespace, ingress, data center, etc)
	Resolved   bool      `json:"resolved"`
	ResolvedBy string    `json:"resolvedBy"`
	ClusterId  string    `json:"clusterId"`
	Status     string    `json:"status,omitempty"`
	Topic      string    `json:"topic,omitempty"`
	Message    string    `json:"message,omitempty"`
}
