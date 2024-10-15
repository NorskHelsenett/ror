package models

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Applications struct {
	Items []Application `json:"items"`
}

type Application struct {
	ApiVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   metav1.ObjectMeta `json:"metadata"`
	//Operation  ApplicationOperation `json:"operation"`
	Spec   ApplicationSpec   `json:"spec"`
	Status ApplicationStatus `json:"status"`
}

type ApplicationSpec struct {
	Destination        ApplicationDestination `json:"destination"`
	IgnoredDifferences []string               `json:"ignoredDifferences"`
	Info               []ApplicationInfo      `json:"info"`
	Source             ApplicationSource      `json:"source"`
}

type ApplicationDestination struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Server    string `json:"server"`
}

type ApplicationInfo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ApplicationSource struct {
	Chart          string                     `json:"chart"`
	Directory      ApplicationSourceDirectory `json:"directory"`
	Helm           ApplicationSourceHelm      `json:"helm"`
	Path           string                     `json:"path"`
	Ref            string                     `json:"ref"`
	RepoUrl        string                     `json:"repoUrl"`
	TargetRevision string                     `json:"targetRevision"`
}

type ApplicationSourceDirectory struct {
	Exclude string `json:"exclude"`
	Include string `json:"include"`
	Recurse bool   `json:"recurse"`
	//Jsonnet ApplicationSourceDirectoryJsonnet `json:"jsonnet"`
}

type ApplicationSourceHelm struct {
}

type ApplicationStatus struct {
	//Conditions          []ApplicationStatusCondition `json:"conditions"`
	ControllerNamespace string `json:"controllerNamespace"`
	//Health ApplicationStatusHealth `json:"health"`
	//History 		   []ApplicationStatusHistory   `json:"history"`
	ObservedAt string `json:"observedAt"`
	//OperationState	  ApplicationStatusOperationState `json:"operationState"`
	ReconciledAt         string `json:"reconciledAt"`
	ResourceHealthSource string `json:"resourceHealthSource"`
	//Resources 		 []ApplicationStatusResource  `json:"resources"`
	SourceType string                   `json:"sourceType"`
	Summary    ApplicationStatusSummary `json:"summary"`
	Sync       ApplicationStatusSync    `json:"sync"`
}

type ApplicationStatusSummary struct {
	ExternalUrls []string `json:"externalUrls"`
	Images       []string `json:"images"`
}

type ApplicationStatusSync struct {
	//ComparedTo ApplicationStatusSyncComparedTo   `json:"comparedTo"`
	Revision  string   `json:"revision"`
	Revisions []string `json:"revisions"`
	Status    string   `json:"status"`
}
