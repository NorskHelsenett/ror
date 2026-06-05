package aclmodels

import "github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"

// Acl2Subject is an alias for backward compatibility. Use aclscope.Subject for new code.
type Acl2Subject = aclscope.Subject

const (
	Acl2RorSubjecUnknown         = aclscope.SubjectUnknown
	Acl2RorSubjectCluster        = aclscope.SubjectCluster
	Acl2RorSubjectProject        = aclscope.SubjectProject
	Acl2RorSubjectGlobal         = aclscope.SubjectGlobal
	Acl2RorSubjectAcl            = aclscope.SubjectAcl
	Acl2RorSubjectApiKey         = aclscope.SubjectApiKey
	Acl2RorSubjectDatacenter     = aclscope.SubjectDatacenter
	Acl2RorSubjectWorkspace      = aclscope.SubjectWorkspace
	Acl2RorSubjectPrice          = aclscope.SubjectPrice
	Acl2RorSubjectVirtualMachine = aclscope.SubjectVirtualMachine
	Acl2RorSubjectBackup         = aclscope.SubjectBackup
	Acl2RorSubjectAll            = aclscope.SubjectAll
	Acl2SpamSubjectGit           = aclscope.SubjectSpamGit
)

// Deprecated: Use function GetAcl2RorValidSubjects() as dropin replacement instead.
// This variable gives the possiblity of being overwritten on accident.
var Acl2RorValidSubjects []Acl2Subject = GetAcl2RorValidSubjects()

// GetAcl2RorValidSubjects is an alias for backward compatibility. Use aclscope.GetValidSubjects for new code.
func GetAcl2RorValidSubjects() []Acl2Subject { return aclscope.GetValidSubjects() }
