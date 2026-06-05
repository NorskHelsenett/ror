package aclscope

import (
	"context"
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
	SubjectAll            Subject = "all"
	SubjectSpamGit        Subject = "spamgit"
)

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
