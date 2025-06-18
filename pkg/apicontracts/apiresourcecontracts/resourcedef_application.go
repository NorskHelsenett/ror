package apiresourcecontracts

// K8s application struct used with ArgoCD// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceApplication struct {
	ApiVersion string                    `json:"apiVersion"`
	Kind       string                    `json:"kind"`
	Metadata   ResourceMetadata          `json:"metadata"`
	Spec       ResourceApplicationSpec   `json:"spec"`
	Status     ResourceApplicationStatus `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceApplicationStatus struct {
	SourceType     string                                  `json:"sourceType"`
	ReconciledAt   string                                  `json:"reconciledAt"`
	OperationState ResourceApplicationStatusOperationstate `json:"operationState"`
	Health         ResourceApplicationStatusHealth         `json:"health"`
	Sync           ResourceApplicationStatusSync           `json:"sync"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceApplicationStatusOperationstate struct {
	StartedAt  string `json:"startedAt"`
	FinishedAt string `json:"finishedAt"`
	Phase      string `json:"phase"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceApplicationStatusHealth struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceApplicationStatusSync struct {
	Revision string `json:"revision"`
	Status   string `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceApplicationSpec struct {
	Destination ResourceApplicationSpecDestination `json:"destination"`
	Project     string                             `json:"project"`
	Source      ResourceApplicationSpecSource      `json:"source"`
	SyncPolicy  ResourceApplicationSpecSyncpolicy  `json:"syncPolicy"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceApplicationSpecDestination struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Server    string `json:"server"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceApplicationSpecSource struct {
	Chart          string `json:"chart,omitempty"`
	Path           string `json:"path,omitempty"`
	RepoURL        string `json:"repoURL"`
	TargetRevision string `json:"targetRevision"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceApplicationSpecSyncpolicy struct {
	Automated   *ResourceApplicationSpecSyncpolicyAutomated `json:"automated"`
	Retry       *ResourceApplicationSpecSyncpolicyRetry     `json:"retry"`
	SyncOptions []string                                    `json:"syncOptions"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceApplicationSpecSyncpolicyAutomated struct {
	AllowEmpty bool `json:"allowEmpty"`
	Prune      bool `json:"prune"`
	SelfHeal   bool `json:"selfHeal"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceApplicationSpecSyncpolicyRetry struct {
	Backoff *ResourceApplicationSpecSyncpolicyRetryBackoff `json:"backoff"`
	Limit   int                                            `json:"limit"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceApplicationSpecSyncpolicyRetryBackoff struct {
	Duration    string `json:"duration"`
	Factor      int    `json:"factor"`
	MaxDuration string `json:"maxDuration"`
}
