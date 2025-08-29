package aclmodels

import (
	"context"
	"slices"
)

type Acl2Subject string

const (
	Acl2RorSubjecUnknown         = "UNKNOWN"
	Acl2RorSubjectCluster        = "cluster"
	Acl2RorSubjectProject        = "project"
	Acl2RorSubjectGlobal         = "globalscope" // for subject, not scope, TODO: new const
	Acl2RorSubjectAcl            = "acl"         // for subject, not scope, TODO: new const
	Acl2RorSubjectApiKey         = "apikey"      //api key
	Acl2RorSubjectDatacenter     = "datacenter"
	Acl2RorSubjectWorkspace      = "workspace"
	Acl2RorSubjectPrice          = "price"
	Acl2RorSubjectVirtualMachine = "virtualmachine"
)

var (
	Acl2RorValidSubjects []Acl2Subject = []Acl2Subject{
		Acl2RorSubjectGlobal,
		Acl2RorSubjectCluster,
		Acl2RorSubjectProject,
		Acl2RorSubjectAcl,
		Acl2RorSubjectDatacenter,
		Acl2RorSubjectWorkspace,
		Acl2RorSubjectPrice,
		Acl2RorSubjectVirtualMachine,
	}
)

// TODO: implement
func (s Acl2Subject) HasValidScope(scope Acl2Scope) bool {
	switch scope {
	case Acl2ScopeRor:
		return slices.Contains(Acl2RorValidSubjects, s)
	// case Acl2ScopeCluster:
	// 	return false
	// case Acl2ScopeProject:
	// 	return false
	default:
		return true
	}
}

// TODO: implement
func (s Acl2Scope) GetSubjects(ctx context.Context) []Acl2Subject {
	switch s {
	case Acl2ScopeRor:
		return Acl2RorValidSubjects
	// case Acl2ScopeCluster:
	// 	return []Acl2Subject{}
	// case Acl2ScopeProject:
	// 	return []Acl2Subject{}
	default:
		return []Acl2Subject{}
	}
}
