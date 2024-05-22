package apiresponses

type DatacentersResponse struct {
	Datacenters []string `json:"datacenters"`
	Success     bool     `json:"success"`
	Message     string   `json:"message"`
}
