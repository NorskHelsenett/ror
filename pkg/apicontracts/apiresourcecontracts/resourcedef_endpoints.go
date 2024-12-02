package apiresourcecontracts

type ResourceEndpoints struct {
	ApiVersion string                        `json:"apiVersion"`
	Kind       string                        `json:"kind"`
	Metadata   ResourceMetadata              `json:"metadata"`
	Subsets    []ResourceEndpointSpecSubsets `json:"subsets,omitempty"`
}

type ResourceEndpointSpecSubsets struct {
	Addresses         []ResourceEndpointSpecSubsetsAddresses         `json:"addresses,omitempty"`
	NotReadyAddresses []ResourceEndpointSpecSubsetsNotReadyAddresses `json:"notReadyAddresses,omitempty"`
	Ports             []ResourceEndpointSpecSubsetsPorts             `json:"ports,omitempty"`
}

type ResourceEndpointSpecSubsetsAddresses struct {
	Hostname  string                                        `json:"hostname,omitempty"`
	Ip        string                                        `json:"ip,omitempty"`
	NodeName  string                                        `json:"nodeName,omitempty"`
	TargetRef ResourceEndpointSpecSubsetsAddressesTargetRef `json:"targetRef,omitempty"`
}

type ResourceEndpointSpecSubsetsAddressesTargetRef struct {
	ApiVersion      string `json:"apiVersion,omitempty"`
	FieldPath       string `json:"fieldPath,omitempty"`
	Kind            string `json:"kind,omitempty"`
	Name            string `json:"name,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
	Uid             string `json:"uid,omitempty"`
}

type ResourceEndpointSpecSubsetsNotReadyAddresses struct {
	Hostname  string                                                `json:"hostname,omitempty"`
	Ip        string                                                `json:"ip,omitempty"`
	NodeName  string                                                `json:"nodeName,omitempty"`
	TargetRef ResourceEndpointSpecSubsetsNotReadyAddressesTargetRef `json:"targetRef,omitempty"`
}

type ResourceEndpointSpecSubsetsNotReadyAddressesTargetRef struct {
	ApiVersion      string `json:"apiVersion,omitempty"`
	FieldPath       string `json:"fieldPath,omitempty"`
	Kind            string `json:"kind,omitempty"`
	Name            string `json:"name,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
	Uid             string `json:"uid,omitempty"`
}

type ResourceEndpointSpecSubsetsPorts struct {
	AppProtocol string `json:"appProtocol,omitempty"`
	Name        string `json:"name,omitempty"`
	Port        int32  `json:"port,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
}
