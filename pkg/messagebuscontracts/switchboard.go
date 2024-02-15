package messagebuscontracts

import "github.com/NorskHelsenett/ror/pkg/apicontracts/messages"

type SwitchboardPostSeverity string
type SwitchboardSourceType int

const (
	SwitchboardPostSeverityFatal SwitchboardPostSeverity = "fatal"
	SwitchboardPostSeverityError SwitchboardPostSeverity = "error"
	SwitchboardPostSeverityWarn  SwitchboardPostSeverity = "warn"
	SwitchboardPostSeverityInfo  SwitchboardPostSeverity = "info"
)

const (
	SwitchboardSourceTypeCluster SwitchboardSourceType = iota
	SwitchboardSourceTypeRor
)

type SwitchboardSource struct {
	Type SwitchboardSourceType `json:"type"`

	// name-hash. You will only find this one when source type is cluster. Since there is no need to identify ROR. For sanity's sake. It should be ror by default
	ClusterId string `json:"clusterId,omitempty"`

	// These fields are used to perform a lookup on the resource. If uid is specified we only use that in a fast path solution,
	// if it is a wildcard we use a combination of apiversion and kind to look up the resource.

	// Uid of the resource we are looking up
	Uid string `json:"uid"`

	// ApiVersion of the resource we are looking up
	ApiVersion string `json:"apiVersion"`

	// Kind the kind of resource we are looking up
	Kind string `json:"kind"`

	// Namespace the namespace in which the resource lives
	Namespace string `json:"namespace"`
}

type SwitchboardSubject struct {
	Resource string `json:"resource"`
	Event    string `json:"event"`
}

type SwitchboardPost struct {
	Source SwitchboardSource `json:"source"`

	Severity SwitchboardPostSeverity  `json:"severity"`
	Event    messages.RulesetRuleType `json:"event"`

	Attributes map[string]string `json:"attributes"`
}
