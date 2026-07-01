package aclscope

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
		ScopeDatabase,
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
		ScopeDatabase,
		ScopeAll,
		ScopeSpam,
	}
}
