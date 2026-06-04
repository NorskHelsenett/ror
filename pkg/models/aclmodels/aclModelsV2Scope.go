package aclmodels

import "github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"

// Acl2Scope is an alias for backward compatibility. Use aclscope.Scope for new code.
type Acl2Scope = aclscope.Scope

const (
	Acl2ScopeUnknown        = aclscope.ScopeUnknown
	Acl2ScopeRor            = aclscope.ScopeRor
	Acl2ScopeCluster        = aclscope.ScopeCluster
	Acl2ScopeProject        = aclscope.ScopeProject
	Acl2ScopeDatacenter     = aclscope.ScopeDatacenter
	Acl2ScopeVirtualMachine = aclscope.ScopeVirtualMachine
	Acl2ScopeMachine        = aclscope.ScopeMachine
	Acl2ScopeBackup         = aclscope.ScopeBackup
	Acl2ScopeAll            = aclscope.ScopeAll
	Acl2ScopeSpam           = aclscope.ScopeSpam
)

// GetScopes is an alias for backward compatibility. Use aclscope.GetScopes for new code.
func GetScopes() []Acl2Scope { return aclscope.GetScopes() }
