package aclmodels

import (
	"cmp"
	"slices"
	"time"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

// FieldValue returns the string value of a top-level field by its json/bson name,
// used for sorting and filtering. Unknown fields return "" (so a filter on an
// absent field matches nothing, mirroring the storage query behavior).
func (i AclV3ListItem) FieldValue(field string) string {
	switch field {
	case "group":
		return i.Group
	case "scope":
		return string(i.Scope)
	case "subject":
		return string(i.Subject)
	case "issuedBy":
		return i.IssuedBy
	case "id", "_id":
		return i.Id
	case "created":
		return i.Created.UTC().Format(time.RFC3339Nano)
	default:
		return ""
	}
}

// Filter returns a new list containing the entries for which keep returns true.
func (l AclV3List) Filter(keep func(AclV3ListItem) bool) AclV3List {
	out := make(AclV3List, 0, len(l))
	for _, e := range l {
		if keep(e) {
			out = append(out, e)
		}
	}
	return out
}

// FilterScopeSubject returns the entries that apply to the given scope and
// subject: exact scope+subject, the "all" scope/subject wildcards, or a
// ror-scoped grant whose subject is the queried scope or the global subject.
// Scope and subject must be in Kind form (as stored).
func (l AclV3List) FilterScopeSubject(scope aclscope.Scope, subject aclscope.Subject) AclV3List {
	return l.Filter(func(e AclV3ListItem) bool {
		rorLevel := e.Scope == aclscope.ScopeRor &&
			(e.Subject == aclscope.Subject(scope) || e.Subject == aclscope.SubjectGlobal)
		switch {
		case scope == aclscope.ScopeAll && subject == aclscope.SubjectAll:
			return true
		case scope == aclscope.ScopeAll:
			return e.Subject == subject
		case subject == aclscope.SubjectAll:
			return e.Scope == scope || rorLevel
		default:
			return (e.Scope == scope && e.Subject == subject) || rorLevel
		}
	})
}

// Sorted returns a sorted copy ordered by the given field, with the entry id as a
// unique tiebreaker so the order is a stable total order. The receiver is never
// mutated. Unknown fields sort by id only.
func (l AclV3List) Sorted(field string, asc bool) AclV3List {
	out := make(AclV3List, len(l))
	copy(out, l)
	slices.SortFunc(out, func(a, b AclV3ListItem) int {
		if c := cmp.Compare(a.FieldValue(field), b.FieldValue(field)); c != 0 {
			if asc {
				return c
			}
			return -c
		}
		return cmp.Compare(a.Id, b.Id)
	})
	return out
}

// Page returns the offset/limit window of the list. A limit <= 0 returns all
// entries from offset onward. It returns a sub-slice of the receiver (no copy),
// so callers should page a list they own (e.g. the result of Sorted or Filter).
func (l AclV3List) Page(offset, limit int) AclV3List {
	if offset < 0 {
		offset = 0
	}
	if offset >= len(l) {
		return AclV3List{}
	}
	end := len(l)
	if limit > 0 && offset+limit < end {
		end = offset + limit
	}
	return l[offset:end]
}
