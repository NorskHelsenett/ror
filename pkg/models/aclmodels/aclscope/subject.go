package aclscope

import (
	"context"
	"fmt"
	"slices"
)

// Subject represents the subject of an ACL entry.
// This is the identifier of the specific object, e.g. a cluster ID, project ID, or "All".
type Subject string

const (
	SubjectUnknown        Subject = "UNKNOWN"
	SubjectCluster        Subject = "cluster"
	SubjectProject        Subject = "project"
	SubjectGlobal         Subject = "globalscope"
	SubjectAcl            Subject = "acl"
	SubjectApiKey         Subject = "apikey"
	SubjectDatacenter     Subject = "datacenter"
	SubjectWorkspace      Subject = "workspace"
	SubjectPrice          Subject = "price"
	SubjectVirtualMachine Subject = "virtualmachine"
	SubjectBackup         Subject = "backup"
	SubjectDatabase       Subject = "database"
	SubjectAll            Subject = "all"
	SubjectSpamGit        Subject = "spamgit"
)

// ParseSubject resolves a string to a Subject and validates it for the given scope,
// returning an error if the subject is not valid for that scope.
func ParseSubject(scope Scope, inputSubject string) (Subject, error) {
	resolvedSubject := Subject(inputSubject)
	if !resolvedSubject.HasValidScope(scope) {
		return SubjectUnknown, fmt.Errorf("invalid subject %q for scope %q", inputSubject, scope)
	}
	return resolvedSubject, nil
}

// String returns the string representation of the subject.
func (s Subject) String() string {
	return string(s)
}

// IsValid validates the subject for the given scope.
func (s Subject) IsValid(scope Scope) bool {
	return s.HasValidScope(scope)
}

// GetValidSubjects returns all valid subjects for the "ror" scope.
func GetValidSubjects() []Subject {
	return []Subject{
		SubjectGlobal,
		SubjectCluster,
		SubjectProject,
		SubjectAcl,
		SubjectDatacenter,
		SubjectWorkspace,
		SubjectPrice,
		SubjectVirtualMachine,
		SubjectBackup,
		SubjectDatabase,
		SubjectAll,
	}
}

// HasValidScope checks if the subject is valid for the given scope.
func (s Subject) HasValidScope(scope Scope) bool {
	switch scope {
	case ScopeRor:
		return slices.Contains(GetValidSubjects(), s)
	default:
		return true
	}
}

// GetSubjects returns valid subjects for a given scope.
func (s Scope) GetSubjects(ctx context.Context) []Subject {
	switch s {
	case ScopeRor:
		return GetValidSubjects()
	default:
		return []Subject{}
	}
}
