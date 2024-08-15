package rortypes

// K8s application struct used with ArgoCD
type ResourceApplication struct {
	Spec   ResourceApplicationSpec   `json:"spec"`
	Status ResourceApplicationStatus `json:"status"`
}

type ResourceApplicationStatus struct {
	SourceType     string                                  `json:"sourceType"`
	ReconciledAt   string                                  `json:"reconciledAt"`
	OperationState ResourceApplicationStatusOperationstate `json:"operationState"`
	Health         ResourceApplicationStatusHealth         `json:"health"`
	Sync           ResourceApplicationStatusSync           `json:"sync"`
}
type ResourceApplicationStatusOperationstate struct {
	StartedAt  string `json:"startedAt"`
	FinishedAt string `json:"finishedAt"`
	Phase      string `json:"phase"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

type ResourceApplicationStatusHealth struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
type ResourceApplicationStatusSync struct {
	Revision string `json:"revision"`
	Status   string `json:"status"`
}

type ResourceApplicationSpec struct {
	Destination ResourceApplicationSpecDestination `json:"destination"`
	Project     string                             `json:"project"`
	Source      ResourceApplicationSpecSource      `json:"source"`
	SyncPolicy  ResourceApplicationSpecSyncpolicy  `json:"syncPolicy"`
}
type ResourceApplicationSpecDestination struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Server    string `json:"server"`
}
type ResourceApplicationSpecSource struct {
	Chart          string `json:"chart,omitempty"`
	Path           string `json:"path,omitempty"`
	RepoURL        string `json:"repoURL"`
	TargetRevision string `json:"targetRevision"`
}

type ResourceApplicationSpecSyncpolicy struct {
	Automated   *ResourceApplicationSpecSyncpolicyAutomated `json:"automated"`
	Retry       *ResourceApplicationSpecSyncpolicyRetry     `json:"retry"`
	SyncOptions []string                                    `json:"syncOptions"`
}

type ResourceApplicationSpecSyncpolicyAutomated struct {
	AllowEmpty bool `json:"allowEmpty"`
	Prune      bool `json:"prune"`
	SelfHeal   bool `json:"selfHeal"`
}

type ResourceApplicationSpecSyncpolicyRetry struct {
	Backoff *ResourceApplicationSpecSyncpolicyRetryBackoff `json:"backoff"`
	Limit   int                                            `json:"limit"`
}

type ResourceApplicationSpecSyncpolicyRetryBackoff struct {
	Duration    string `json:"duration"`
	Factor      int    `json:"factor"`
	MaxDuration string `json:"maxDuration"`
}
