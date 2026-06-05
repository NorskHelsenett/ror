package aclscope

// Scope represents the scope of an ACL entry.
// Valid values are known resource kinds (e.g. "cluster", "project")
// or system identifiers (e.g. "ror", "all").
type Scope string

const (
	ScopeUnknown        Scope = "UNKNOWN"
	ScopeRor            Scope = "ror"
	ScopeCluster        Scope = "cluster"
	ScopeProject        Scope = "project"
	ScopeDatacenter     Scope = "datacenter"
	ScopeVirtualMachine Scope = "virtualmachine"
	ScopeMachine        Scope = "machine"
	ScopeBackup         Scope = "backup"
	ScopeAll            Scope = "all"
	ScopeSpam           Scope = "spam"
)

// IsValid validates the scope
func (s Scope) IsValid() bool {
	switch s {
	case ScopeRor,
		ScopeCluster,
		ScopeProject,
		ScopeDatacenter,
		ScopeVirtualMachine,
		ScopeMachine,
		ScopeBackup,
		ScopeAll,
		ScopeSpam:
		return true
	default:
		return false
	}
}

// GetScopes returns all valid scopes.
func GetScopes() []Scope {
	return []Scope{
		ScopeRor,
		ScopeCluster,
		ScopeVirtualMachine,
		ScopeBackup,
		ScopeProject,
		ScopeMachine,
		ScopeDatacenter,
		ScopeAll,
		ScopeSpam,
	}
}
