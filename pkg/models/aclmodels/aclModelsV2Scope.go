package aclmodels

type Acl2Scope string

const (
	Acl2ScopeUnknown        Acl2Scope = "UNKNOWN" // unknown
	Acl2ScopeRor            Acl2Scope = "ror"     // ROR
	Acl2ScopeCluster        Acl2Scope = "cluster"
	Acl2ScopeProject        Acl2Scope = "project"
	Acl2ScopeDatacenter     Acl2Scope = "datacenter"
	Acl2ScopeVirtualMachine Acl2Scope = "virtualmachine"
	Acl2ScopeBackup         Acl2Scope = "backup"
)

// IsValid validates the scope
func (s Acl2Scope) IsValid() bool {
	switch s {
	case Acl2ScopeRor:
		return true
	case Acl2ScopeCluster:
		return true
	case Acl2ScopeProject:
		return true
	case Acl2ScopeDatacenter:
		return true
	case Acl2ScopeVirtualMachine:
		return true
	case Acl2ScopeBackup:
		return true
	case Acl2ScopeUnknown:
		return false
	default:
		return false
	}
}

func GetScopes() []Acl2Scope {
	return []Acl2Scope{
		Acl2ScopeRor,
		Acl2ScopeCluster,
		Acl2ScopeVirtualMachine,
		Acl2ScopeBackup,
		Acl2ScopeProject,
	}
}
