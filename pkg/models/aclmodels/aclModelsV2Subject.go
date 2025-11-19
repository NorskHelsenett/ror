package aclmodels

import (
	"context"
	"slices"
)

type Acl2Subject string

const (
	Acl2RorSubjecUnknown         Acl2Subject = "UNKNOWN"
	Acl2RorSubjectCluster        Acl2Subject = "cluster"
	Acl2RorSubjectProject        Acl2Subject = "project"
	Acl2RorSubjectGlobal         Acl2Subject = "globalscope" // for subject, not scope, TODO: new const
	Acl2RorSubjectAcl            Acl2Subject = "acl"         // for subject, not scope, TODO: new const
	Acl2RorSubjectApiKey         Acl2Subject = "apikey"      //api key
	Acl2RorSubjectDatacenter     Acl2Subject = "datacenter"
	Acl2RorSubjectWorkspace      Acl2Subject = "workspace"
	Acl2RorSubjectPrice          Acl2Subject = "price"
	Acl2RorSubjectVirtualMachine Acl2Subject = "virtualmachine"
)

// Deprecated: Use function GetAcl2RorValidSubjects() as dropin replacement instead.
// This variable gives the possiblity of being overwritten on accident.
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

// GetAcl2RorValidSubjects returns all possible Acl2Subject values.
func GetAcl2RorValidSubjects() []Acl2Subject {
	return []Acl2Subject{
		Acl2RorSubjectGlobal,
		Acl2RorSubjectCluster,
		Acl2RorSubjectProject,
		Acl2RorSubjectAcl,
		Acl2RorSubjectDatacenter,
		Acl2RorSubjectWorkspace,
		Acl2RorSubjectPrice,
		Acl2RorSubjectVirtualMachine,
	}
}

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
