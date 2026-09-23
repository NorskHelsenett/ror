package aclscope

// legacyKindNames is the single source of truth for legacy V2 name <-> resource
// Kind name pairs. The Scope and Subject alias maps below are all derived from
// it, so the mapping is defined exactly once.
var legacyKindNames = map[string]string{
	"cluster":        "KubernetesCluster",
	"project":        "Project",
	"workspace":      "Workspace",
	"virtualmachine": "VirtualMachine",
	"backup":         "BackupJob",
	"datacenter":     "Datacenter",
	"machine":        "Machine",
}

// aliasMap builds a legacy<->kind alias map of the given string type from
// legacyKindNames, inverting (kind -> legacy) when invert is true.
func aliasMap[T ~string](invert bool) map[T]T {
	out := make(map[T]T, len(legacyKindNames))
	for legacy, kind := range legacyKindNames {
		if invert {
			out[T(kind)] = T(legacy)
		} else {
			out[T(legacy)] = T(kind)
		}
	}
	return out
}

// LegacyToKind maps legacy V2 scope names to resource Kind names.
var LegacyToKind = aliasMap[Scope](false)

// KindToLegacy maps resource Kind names back to legacy V2 scope names.
var KindToLegacy = aliasMap[Scope](true)

// LegacySubjectToKind maps legacy V2 subject names (used with scope "ror") to
// resource Kind names. These represent type-level grants (e.g. "can manage all clusters").
var LegacySubjectToKind = aliasMap[Subject](false)

// KindToLegacySubject maps resource Kind subject names back to legacy V2 subject names.
var KindToLegacySubject = aliasMap[Subject](true)

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
