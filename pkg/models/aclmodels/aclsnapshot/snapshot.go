// Package aclsnapshot defines the signed ACL snapshot payload contract.
// Decoding validates data but does not authenticate it.
package aclsnapshot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclgrants"
)

const (
	Schema    = "ror.acl.snapshot/1"
	MediaType = "application/ror-acl-snapshot+jws"
)

// Revision is a positive uint64 encoded as a canonical decimal JSON string.
type Revision uint64

func (r Revision) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return nil, fmt.Errorf("snapshot revision must be positive")
	}
	return json.Marshal(strconv.FormatUint(uint64(r), 10))
}

func (r *Revision) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return fmt.Errorf("snapshot revision must be a decimal string: %w", err)
	}
	n, err := strconv.ParseUint(value, 10, 64)
	if err != nil || n == 0 || strconv.FormatUint(n, 10) != value {
		return fmt.Errorf("invalid snapshot revision %q", value)
	}
	*r = Revision(n)
	return nil
}

type Row struct {
	Group string          `json:"group"`
	Grant aclgrants.Grant `json:"grant"`
}

type Payload struct {
	Schema      string    `json:"schema"`
	Revision    Revision  `json:"revision"`
	GeneratedAt time.Time `json:"generatedAt"`
	Rows        []Row     `json:"rows"`
}

// Limits are required deployment bounds, not suggested production defaults.
// MaxFieldBytes applies to group/scope/subject/access strings and signing key IDs.
type Limits struct {
	MaxPayloadBytes  int
	MaxEnvelopeBytes int
	MaxRows          int
	MaxAccessEntries int
	MaxFieldBytes    int
	MaxHeaderBytes   int
}

func (l Limits) Validate() error {
	if l.MaxRows == math.MaxInt || l.MaxHeaderBytes > math.MaxInt/4*3 {
		return fmt.Errorf("snapshot limits exceed representable query or encoding bounds")
	}
	for _, field := range []struct {
		name  string
		value int
	}{
		{"maxPayloadBytes", l.MaxPayloadBytes}, {"maxEnvelopeBytes", l.MaxEnvelopeBytes},
		{"maxRows", l.MaxRows}, {"maxAccessEntries", l.MaxAccessEntries},
		{"maxFieldBytes", l.MaxFieldBytes}, {"maxHeaderBytes", l.MaxHeaderBytes},
	} {
		if field.value <= 0 {
			return fmt.Errorf("%s must be positive", field.name)
		}
	}
	return nil
}

func (p Payload) Validate(l Limits) error {
	if err := l.Validate(); err != nil {
		return err
	}
	if p.Schema != Schema || p.Revision == 0 {
		return fmt.Errorf("invalid snapshot schema or revision")
	}
	t := p.GeneratedAt
	_, offset := t.Zone()
	if t.IsZero() || t.Year() < 1 || t.Year() > 9999 || offset != 0 || t.Nanosecond() != 0 {
		return fmt.Errorf("generatedAt must be nonzero UTC whole-second RFC3339 time")
	}
	return validateRows(p.Rows, l)
}

func validateRows(rows []Row, l Limits) error {
	if rows == nil || len(rows) > l.MaxRows {
		return fmt.Errorf("rows must be non-null and contain at most %d entries", l.MaxRows)
	}
	for i, row := range rows {
		for _, value := range []string{row.Group, string(row.Grant.Scope), string(row.Grant.Subject)} {
			if value == "" || !utf8.ValidString(value) || len(value) > l.MaxFieldBytes {
				return fmt.Errorf("row %d: invalid or oversized group, scope or subject", i)
			}
		}
		if row.Grant.Access == nil || len(row.Grant.Access) > l.MaxAccessEntries {
			return fmt.Errorf("row %d: access must be non-null and contain at most %d entries", i, l.MaxAccessEntries)
		}
		accessSeen := make(map[string]bool, len(row.Grant.Access))
		for _, access := range row.Grant.Access {
			value := string(access)
			if !utf8.ValidString(value) || len(value) > l.MaxFieldBytes || accessSeen[value] {
				return fmt.Errorf("row %d: invalid, oversized or duplicate access", i)
			}
			parts := strings.Split(value, ":")
			if len(parts) < 2 || slices.Contains(parts, "") {
				return fmt.Errorf("row %d: access requires nonempty colon-separated components", i)
			}
			accessSeen[value] = true
		}
	}
	return nil
}

// CanonicalRows validates and sorts copied rows/access lists without changing
// caller-owned slices. It preserves repeated grants, empty policy and empty access arrays.
func CanonicalRows(rows []Row, l Limits) ([]Row, error) {
	if err := l.Validate(); err != nil {
		return nil, err
	}
	if err := validateRows(rows, l); err != nil {
		return nil, err
	}
	result := make([]Row, len(rows))
	for i, row := range rows {
		result[i] = row
		result[i].Grant.Access = append(row.Grant.Access[:0:0], row.Grant.Access...)
		slices.Sort(result[i].Grant.Access)
	}
	slices.SortFunc(result, func(a, b Row) int {
		for _, pair := range [][2]string{
			{a.Group, b.Group}, {string(a.Grant.Scope), string(b.Grant.Scope)},
			{string(a.Grant.Subject), string(b.Grant.Subject)},
		} {
			if cmp := strings.Compare(pair[0], pair[1]); cmp != 0 {
				return cmp
			}
		}
		return slices.Compare(a.Grant.Access, b.Grant.Access)
	})
	return result, nil
}

// Encode emits deterministic compact JSON without a trailing newline.
func Encode(p Payload, l Limits) ([]byte, error) {
	if err := p.Validate(l); err != nil {
		return nil, err
	}
	rows, err := CanonicalRows(p.Rows, l)
	if err != nil {
		return nil, err
	}
	p.Rows = rows
	p.GeneratedAt = p.GeneratedAt.UTC()
	data, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("encode snapshot: %w", err)
	}
	if len(data) > l.MaxPayloadBytes {
		return nil, fmt.Errorf("snapshot exceeds %d payload bytes", l.MaxPayloadBytes)
	}
	return data, nil
}

// Decode must only be used as trusted policy after signature verification.
// It rejects unknown fields, duplicate object keys and noncanonical timestamps.
func Decode(data []byte, l Limits) (Payload, error) {
	var p Payload
	if err := l.Validate(); err != nil {
		return p, err
	}
	if len(data) > l.MaxPayloadBytes || !utf8.Valid(data) {
		return p, fmt.Errorf("snapshot payload is oversized or invalid UTF-8")
	}
	check := json.NewDecoder(bytes.NewReader(data))
	if err := checkValue(check, 0); err != nil {
		return p, err
	}
	if _, err := check.Token(); err != io.EOF {
		return p, fmt.Errorf("snapshot contains trailing data")
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&p); err != nil {
		return Payload{}, fmt.Errorf("decode snapshot: %w", err)
	}
	var timestamp struct {
		GeneratedAt string `json:"generatedAt"`
	}
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return Payload{}, fmt.Errorf("decode generatedAt: %w", err)
	}
	if timestamp.GeneratedAt != p.GeneratedAt.UTC().Format(time.RFC3339) {
		return Payload{}, fmt.Errorf("generatedAt must use whole-second UTC Z format")
	}
	if err := p.Validate(l); err != nil {
		return Payload{}, err
	}
	return p, nil
}

func checkValue(decoder *json.Decoder, depth int) error {
	if depth > 8 {
		return fmt.Errorf("snapshot JSON is nested too deeply")
	}
	token, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("invalid snapshot JSON: %w", err)
	}
	delim, composite := token.(json.Delim)
	if !composite {
		return nil
	}
	seen := map[string]bool{}
	for decoder.More() {
		if delim == '{' {
			key, err := decoder.Token()
			if err != nil {
				return fmt.Errorf("invalid snapshot object key: %w", err)
			}
			name, ok := key.(string)
			if !ok || seen[name] {
				return fmt.Errorf("invalid or duplicate snapshot object key")
			}
			if !allowedSnapshotField(depth, name) {
				return fmt.Errorf("invalid snapshot object key %q", name)
			}
			seen[name] = true
		}

		if err := checkValue(decoder, depth+1); err != nil {
			return err
		}
	}
	if _, err := decoder.Token(); err != nil {
		return fmt.Errorf("invalid snapshot JSON closing delimiter: %w", err)
	}
	return nil
}

func allowedSnapshotField(depth int, name string) bool {
	switch depth {
	case 0:
		return name == "schema" || name == "revision" || name == "generatedAt" || name == "rows"
	case 2:
		return name == "group" || name == "grant"
	case 3:
		return name == "scope" || name == "subject" || name == "access"
	}
	return false
}
