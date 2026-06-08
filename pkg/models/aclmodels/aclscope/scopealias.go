package aclscope

// LegacyToKind maps legacy V2 scope names to resource Kind names.
// Used to translate V2 API calls to the Kind-based scope values stored in the database
// after the scope migration.
var LegacyToKind = map[Scope]Scope{
	"cluster":        "KubernetesCluster",
	"project":        "Project",
	"workspace":      "Workspace",
	"virtualmachine": "VirtualMachine",
	"backup":         "BackupJob",
	"datacenter":     "Datacenter",
	"machine":        "Machine",
}

// KindToLegacy maps resource Kind names back to legacy V2 scope names.
// Used when returning data to V2 API consumers that expect the old naming.
var KindToLegacy = map[Scope]Scope{
	"KubernetesCluster": "cluster",
	"Project":           "project",
	"Workspace":         "workspace",
	"VirtualMachine":    "virtualmachine",
	"BackupJob":         "backup",
	"Datacenter":        "datacenter",
	"Machine":           "machine",
}

// LegacySubjectToKind maps legacy V2 subject names (used with scope "ror") to
// resource Kind names. These represent type-level grants (e.g. "can manage all clusters").
var LegacySubjectToKind = map[Subject]Subject{
	"cluster":        "KubernetesCluster",
	"project":        "Project",
	"workspace":      "Workspace",
	"virtualmachine": "VirtualMachine",
	"backup":         "BackupJob",
	"datacenter":     "Datacenter",
	"machine":        "Machine",
}

// KindToLegacySubject maps resource Kind subject names back to legacy V2 subject names.
var KindToLegacySubject = map[Subject]Subject{
	"KubernetesCluster": "cluster",
	"Project":           "project",
	"Workspace":         "workspace",
	"VirtualMachine":    "virtualmachine",
	"BackupJob":         "backup",
	"Datacenter":        "datacenter",
	"Machine":           "machine",
}

// ToKind translates a legacy scope to its Kind equivalent.
// If no mapping exists, returns the scope unchanged (it may already be a Kind).
func (s Scope) ToKind() Scope {
	if mapped, ok := LegacyToKind[s]; ok {
		return mapped
	}
	return s
}

// ToLegacy translates a Kind-based scope to its legacy V2 equivalent.
// If no mapping exists, returns the scope unchanged.
func (s Scope) ToLegacy() Scope {
	if mapped, ok := KindToLegacy[s]; ok {
		return mapped
	}
	return s
}

// ToKind translates a legacy subject to its Kind equivalent.
// If no mapping exists, returns the subject unchanged.
func (s Subject) ToKind() Subject {
	if mapped, ok := LegacySubjectToKind[s]; ok {
		return mapped
	}
	return s
}

// ToLegacy translates a Kind-based subject to its legacy V2 equivalent.
// If no mapping exists, returns the subject unchanged.
func (s Subject) ToLegacy() Subject {
	if mapped, ok := KindToLegacySubject[s]; ok {
		return mapped
	}
	return s
}
