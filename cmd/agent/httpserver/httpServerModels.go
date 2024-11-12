package httpserver

type HealthStatus struct {
	Status string       `json:"status"`
	Report HealthReport `json:"report"`
}

type HealthReport struct {
	Kubernetes K8sReport    `json:"kubernetes"`
	RorApi     RorApiReport `json:"rorApi"`
	ErrorCount int          `json:"errorCount"`
}

type K8sReport struct {
	HasConfig bool `json:"hasConfig"`
}

type RorApiReport struct {
	GotToken bool `json:"token"`
}
