package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Phase string

const (
	PhasePending Phase = "PENDING"
	PhaseRunning Phase = "RUNNING"
	PhaseDone    Phase = "DONE"
	PhaseError   Phase = "ERROR"
)

type TaskSuccess struct {
	Timestamp metav1.Time `json:"timestamp"`
	ConfigMd5 string      `json:"configMd5"`
}

type TaskFailure struct {
	Reason    string      `json:"reason"`
	Timestamp metav1.Time `json:"timestamp"`
}
