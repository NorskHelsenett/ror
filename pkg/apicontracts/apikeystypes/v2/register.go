package apikeystypes

type RegisterClusterRequest struct {
	ClusterId string `json:"clusterid"`
}

type RegisterClusterResponse struct {
	ClusterId string `json:"clusterid"`
	ApiKey    string `json:"apikey"`
}
