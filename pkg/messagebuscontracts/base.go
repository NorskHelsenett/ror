package messagebuscontracts

type EventBase struct {
	TraceId string `json:"traceId"`
}

type EventClusterBase struct {
	ClusterId string `json:"clusterId"`
}
