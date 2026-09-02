package aclscope

import "sort"

// Scope represents the scope of an ACL entry.
// Valid values are known resource kinds (e.g. "cluster", "project")
// or system identifiers (e.g. "ror", "all").
type Scope string

const (
	ScopeUnknown        Scope = "UNKNOWN"
	ScopeRor            Scope = "ror"
	ScopeCluster        Scope = "KubernetesCluster"
	ScopeProject        Scope = "Project"
	ScopeDatacenter     Scope = "Datacenter"
	ScopeVirtualMachine Scope = "VirtualMachine"
	ScopeMachine        Scope = "Machine"
	ScopeBackup         Scope = "BackupJob"
	ScopeDatabase       Scope = "Database"
	ScopeAll            Scope = "all"
	ScopeSpam           Scope = "spam"
)

// scopeAliases maps every accepted string representation to its Scope.
// This is the single source of truth for ScopeFromString, IsValid and GetScopes.
var scopeAliases = map[string]Scope{
	"ror":               ScopeRor,
	"cluster":           ScopeCluster,
	"KubernetesCluster": ScopeCluster,
	"project":           ScopeProject,
	"Project":           ScopeProject,
	"datacenter":        ScopeDatacenter,
	"Datacenter":        ScopeDatacenter,
	"virtualmachine":    ScopeVirtualMachine,
	"VirtualMachine":    ScopeVirtualMachine,
	"machine":           ScopeMachine,
	"Machine":           ScopeMachine,
	"backup":            ScopeBackup,
	"BackupJob":         ScopeBackup,
	"database":          ScopeDatabase,
	"Database":          ScopeDatabase,
	"all":               ScopeAll,
	"spam":              ScopeSpam,
}

func ScopeFromString(s string) (Scope, bool) {
	scope, ok := scopeAliases[s]
	if !ok {
		return ScopeUnknown, false
	}
	return scope, true
}

// String returns the string representation of the scope.
func (s Scope) String() string {
	return string(s)
}

// IsValid validates the scope
func (s Scope) IsValid() bool {
	scope, ok := scopeAliases[string(s)]
	return ok && scope == s
}

// GetScopes returns all valid scopes.
func GetScopes() []Scope {
	scopes := make([]Scope, 0)
	for alias, scope := range scopeAliases {
		// only the canonical alias (key == value) to avoid duplicates
		if alias == string(scope) {
			scopes = append(scopes, scope)
		}
	}
	sort.Slice(scopes, func(i, j int) bool {
		return scopes[i] < scopes[j]
	})
	return scopes
}
