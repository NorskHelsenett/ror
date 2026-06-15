package apikeystypes

// RegisterClusterRequest is the agent registration payload. The agent may
// identify the cluster by its clusterid, its uid, or both. At least one of the
// two must be provided. A caller-supplied Uid is treated as a hint only: the
// API verifies it against an existing cluster and never blindly trusts it.
type RegisterClusterRequest struct {
	ClusterId string `json:"clusterid,omitempty"`
	Uid       string `json:"uid,omitempty"`
}

type RegisterClusterResponse struct {
	ClusterId string `json:"clusterid"`
	ApiKey    string `json:"apikey"`
	Uid       string `json:"uid,omitempty"`
}
