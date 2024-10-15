/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TaskStatus defines the observed state of Task
type TaskStatus struct {
	// +kubebuilder:validation:Optional
	Success *TaskSuccess `json:"success"`

	// +kubebuilder:validation:Optional
	Failure *TaskFailure `json:"failure"`
	Phase   Phase        `json:"phase"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Version",type=string,JSONPath=`.spec.version`
// +kubebuilder:printcolumn:name="Index",type=integer,JSONPath=`.spec.index`
// +kubebuilder:printcolumn:name="RunOnce",type=boolean,JSONPath=`.spec.runOnce`,priority=50
// +kubebuilder:printcolumn:name="FailureTimestamp",type=string,JSONPath=`.status.failure.timestamp`,priority=100
// +kubebuilder:printcolumn:name="FailureReason",type=string,JSONPath=`.status.failure.reason`,priority=100

// Task is the Schema for the tasks API
type Task struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   apicontracts.OperatorJob `json:"spec,omitempty"`
	Status TaskStatus               `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TaskList contains a list of Task
type TaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Task `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Task{}, &TaskList{})
}
