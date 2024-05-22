package apiresponses

type HealthStatusCode int

const (
	StatusOK           HealthStatusCode = 1
	StatusNotConnected HealthStatusCode = -1
	StatusUnableToPing HealthStatusCode = -2
)

type Services struct {
	Name   string           `json:"name"`
	Status HealthStatusCode `json:"status"`
}

type HealthStatus struct {
	Services []Services `json:"services"`
}
