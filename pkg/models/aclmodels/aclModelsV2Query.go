package aclmodels

import (
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// v2 querymodel for access
type AclV2QueryAccessScopeSubject struct {
	Scope   Acl2Scope
	Subject Acl2Subject
}
type AclV2QueryAccessScope struct {
	Scope Acl2Scope
}

func NewAclV2QueryAccessScopeSubject(scope any, subject any) AclV2QueryAccessScopeSubject {
	var returnQuery AclV2QueryAccessScopeSubject

	switch s := scope.(type) {
	case Acl2Scope:
		returnQuery.Scope = s
	case string:
		returnQuery.Scope = Acl2Scope(s)
	default:
		input, ok := s.(string)
		if !ok && input == "" {
			returnQuery.Scope = Acl2ScopeUnknown
		} else {
			returnQuery.Scope = Acl2Scope(input)
		}
	}

	switch s := subject.(type) {
	case Acl2Subject:
		returnQuery.Subject = s
	case string:
		returnQuery.Subject = Acl2Subject(s)
	default:
		inputsubject, ok := s.(string)
		if !ok && inputsubject == "" {
			returnQuery.Subject = Acl2Subject("ErrorSubject")
		} else {
			returnQuery.Subject = Acl2Subject(inputsubject)
		}
	}

	if !returnQuery.IsValid() {
		returnQuery.Scope = Acl2ScopeUnknown
		returnQuery.Subject = Acl2Subject("ErrorSubject")
	}

	return returnQuery
}

func (q AclV2QueryAccessScopeSubject) IsValid() bool {

	if !q.Scope.IsValid() {
		rlog.Debug("Scope is invalid", rlog.Any("scope", q.Scope))
	}
	if !q.Subject.HasValidScope(q.Scope) {
		rlog.Debug("subject is invalid", rlog.Any("subject", q.Subject), rlog.Any("scope", q.Scope))
	}

	return q.Subject.HasValidScope(q.Scope) && q.Scope.IsValid()

}
