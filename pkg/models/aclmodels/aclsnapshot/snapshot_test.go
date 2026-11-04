package aclsnapshot

import (
	"encoding/json"
	"math"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclcaps"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclgrants"
)

func testLimits() Limits {
	return Limits{MaxPayloadBytes: 8192, MaxEnvelopeBytes: 16384, MaxRows: 10,
		MaxAccessEntries: 10, MaxFieldBytes: 128, MaxHeaderBytes: 256}
}

func testPayload() Payload {
	return Payload{Schema: Schema, Revision: Revision(math.MaxUint64),
		GeneratedAt: time.Date(2026, 11, 4, 14, 0, 0, 0, time.UTC),
		Rows: []Row{{Group: "service@ror.io", Grant: aclgrants.Grant{
			Scope: "ror", Subject: "acl", Access: []aclcaps.AccessTypeV3{"ror:read"},
		}}}}
}

func TestEncodeDecode(t *testing.T) {
	p := testPayload()
	data, err := Encode(p, testLimits())
	if err != nil {
		t.Fatal(err)
	}
	want := `{"schema":"ror.acl.snapshot/1","revision":"18446744073709551615","generatedAt":"2026-11-04T14:00:00Z","rows":[{"group":"service@ror.io","grant":{"scope":"ror","subject":"acl","access":["ror:read"]}}]}`
	if string(data) != want {
		t.Fatalf("got %s, want %s", data, want)
	}
	decoded, err := Decode(data, testLimits())
	if err != nil || !reflect.DeepEqual(decoded, p) {
		t.Fatalf("round trip: %#v, %v", decoded, err)
	}
	for _, rows := range [][]Row{{}, {{Group: "inert", Grant: aclgrants.Grant{Scope: "*", Subject: "*", Access: []aclcaps.AccessTypeV3{}}}}} {
		p.Rows = rows
		data, err := Encode(p, testLimits())
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(string(data), "null") {
			t.Fatalf("empty arrays became null: %s", data)
		}
		if _, err := Decode(data, testLimits()); err != nil {
			t.Fatal(err)
		}
	}
}

func TestDeterministicAndOwnership(t *testing.T) {
	p := testPayload()
	p.Rows[0].Grant.Access = []aclcaps.AccessTypeV3{"ror:write", "ror:read"}
	p.Rows = append(p.Rows, Row{Group: "a", Grant: p.Rows[0].Grant})
	before, _ := json.Marshal(p)
	first, err := Encode(p, testLimits())
	if err != nil {
		t.Fatal(err)
	}
	after, _ := json.Marshal(p)
	if string(before) != string(after) {
		t.Fatal("caller data mutated")
	}
	p.Rows[0], p.Rows[1] = p.Rows[1], p.Rows[0]
	p.Rows[0].Grant.Access = []aclcaps.AccessTypeV3{"ror:read", "ror:write"}
	second, err := Encode(p, testLimits())
	if err != nil || string(first) != string(second) {
		t.Fatalf("encoding not deterministic: %v", err)
	}
}

func TestValidationFailures(t *testing.T) {
	for _, tt := range []struct {
		name string
		edit func(*Payload)
	}{
		{"schema", func(p *Payload) { p.Schema = "other" }},
		{"revision", func(p *Payload) { p.Revision = 0 }},
		{"zero time", func(p *Payload) { p.GeneratedAt = time.Time{} }},
		{"fractional", func(p *Payload) { p.GeneratedAt = p.GeneratedAt.Add(time.Nanosecond) }},
		{"offset", func(p *Payload) { p.GeneratedAt = p.GeneratedAt.In(time.FixedZone("x", 3600)) }},
		{"year", func(p *Payload) { p.GeneratedAt = time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC) }},
		{"nil rows", func(p *Payload) { p.Rows = nil }},
		{"too many rows", func(p *Payload) { p.Rows = make([]Row, 11) }},
		{"group", func(p *Payload) { p.Rows[0].Group = "" }},
		{"scope", func(p *Payload) { p.Rows[0].Grant.Scope = "" }},
		{"subject", func(p *Payload) { p.Rows[0].Grant.Subject = "" }},
		{"utf8", func(p *Payload) { p.Rows[0].Group = "\xff" }},
		{"field bound", func(p *Payload) { p.Rows[0].Group = strings.Repeat("a", 129) }},
		{"nil access", func(p *Payload) { p.Rows[0].Grant.Access = nil }},
		{"access bound", func(p *Payload) { p.Rows[0].Grant.Access = make([]aclcaps.AccessTypeV3, 11) }},
		{"duplicate access", func(p *Payload) { p.Rows[0].Grant.Access = []aclcaps.AccessTypeV3{"a:b", "a:b"} }},
		{"empty access", func(p *Payload) { p.Rows[0].Grant.Access = []aclcaps.AccessTypeV3{""} }},
		{"malformed access", func(p *Payload) { p.Rows[0].Grant.Access = []aclcaps.AccessTypeV3{"a::b"} }},
		{"bare access", func(p *Payload) { p.Rows[0].Grant.Access = []aclcaps.AccessTypeV3{"read"} }},
	} {
		t.Run(tt.name, func(t *testing.T) {
			p := testPayload()
			tt.edit(&p)
			if _, err := Encode(p, testLimits()); err == nil {
				t.Fatal("invalid payload accepted")
			}
		})
	}
}

func TestRepeatedRowsPreservedAndDeterministic(t *testing.T) {
	p := testPayload()
	grant := p.Rows[0]
	p.Rows = nil
	for _, access := range [][]aclcaps.AccessTypeV3{
		{"ror:write", "ror:read"}, {"ror:read"}, {}, {"ror:read"}, {"ror:write"},
	} {
		row := grant
		row.Grant.Access = access
		p.Rows = append(p.Rows, row)
	}
	before, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	first, err := Encode(p, testLimits())
	if err != nil {
		t.Fatal(err)
	}
	after, err := json.Marshal(p)
	if err != nil || string(before) != string(after) {
		t.Fatal("canonicalization mutated caller-owned repeated rows")
	}
	decoded, err := Decode(first, testLimits())
	if err != nil {
		t.Fatal(err)
	}
	want := [][]aclcaps.AccessTypeV3{
		{}, {"ror:read"}, {"ror:read"}, {"ror:read", "ror:write"}, {"ror:write"},
	}
	if len(decoded.Rows) != len(want) {
		t.Fatalf("repeated rows were lost: %#v", decoded.Rows)
	}
	for i, row := range decoded.Rows {
		if !reflect.DeepEqual(row.Grant.Access, want[i]) {
			t.Fatalf("row %d: access = %v, want %v", i, row.Grant.Access, want[i])
		}
	}
	// Exercise all source-row orders, including identical repetitions and an
	// empty access list, to guard the content hash against MongoDB scan order.
	var permute func(int)
	permute = func(i int) {
		if i == len(p.Rows) {
			encoded, err := Encode(p, testLimits())
			if err != nil || string(encoded) != string(first) {
				t.Fatalf("source order changed canonical bytes: %v", err)
			}
			return
		}
		for j := i; j < len(p.Rows); j++ {
			p.Rows[i], p.Rows[j] = p.Rows[j], p.Rows[i]
			permute(i + 1)
			p.Rows[i], p.Rows[j] = p.Rows[j], p.Rows[i]
		}
	}
	permute(0)
	limits := testLimits()
	limits.MaxRows = len(p.Rows) - 1
	if _, err := Encode(p, limits); err == nil {
		t.Fatal("repeated rows bypassed the row bound")
	}
	if _, err := Decode(first, limits); err == nil {
		t.Fatal("repeated rows bypassed the decoded row bound")
	}
}

func TestDecodeRejections(t *testing.T) {
	valid, _ := Encode(testPayload(), testLimits())
	for _, bad := range []string{
		`null`, `{`, string(valid) + `{}`, string(valid) + ` true`,
		strings.Replace(string(valid), `"rows":[`, `"rows":null,"ignored":[`, 1),
		strings.Replace(string(valid), `"group":`, `"group":"x","group":`, 1),
		strings.Replace(string(valid), `"scope":`, `"scope":"x","scope":`, 1),
		strings.Replace(string(valid), `"schema":`, `"unknown":true,"schema":`, 1),
		strings.Replace(string(valid), `"2026-11-04T14:00:00Z"`, `"2026-11-04T14:00:00+00:00"`, 1),
		strings.Replace(string(valid), `"2026-11-04T14:00:00Z"`, `"2026-11-04T14:00:00.000Z"`, 1),
		strings.Replace(string(valid), `"service@ror.io"`, `null`, 1),
		string([]byte{0xff}), strings.Repeat("[", 10) + strings.Repeat("]", 10),
	} {
		t.Run(bad, func(t *testing.T) {
			if _, err := Decode([]byte(bad), testLimits()); err == nil {
				t.Fatalf("accepted %s", bad)
			}
		})
	}
	for _, value := range []string{`0`, `"0"`, `"-1"`, `"+1"`, `"01"`, `" 1"`, `"1e2"`, `"18446744073709551616"`, `null`} {
		var revision Revision
		if err := json.Unmarshal([]byte(value), &revision); err == nil {
			t.Fatalf("accepted revision %s", value)
		}
	}
	if _, err := json.Marshal(Revision(0)); err == nil {
		t.Fatal("marshaled zero revision")
	}
}

func TestBounds(t *testing.T) {
	p := testPayload()
	limits := testLimits()
	if _, err := CanonicalRows(nil, limits); err == nil {
		t.Fatal("canonicalization accepted nil rows")
	}
	if _, err := CanonicalRows(p.Rows, Limits{}); err == nil {
		t.Fatal("canonicalization accepted missing limits")
	}
	unrepresentable := limits
	unrepresentable.MaxRows = math.MaxInt
	if err := unrepresentable.Validate(); err == nil {
		t.Fatal("query bound overflow accepted")
	}
	data, _ := Encode(p, limits)
	limits.MaxPayloadBytes = len(data)
	limits.MaxRows = 1
	limits.MaxAccessEntries = 1
	limits.MaxFieldBytes = len(p.Rows[0].Group)
	if _, err := Encode(p, limits); err != nil {
		t.Fatal(err)
	}
	if _, err := Decode(data, limits); err != nil {
		t.Fatal(err)
	}
	limits.MaxPayloadBytes--
	if _, err := Encode(p, limits); err == nil {
		t.Fatal("over-limit encode succeeded")
	}
	if _, err := Decode(data, limits); err == nil {
		t.Fatal("over-limit decode succeeded")
	}
	for i := 0; i < 6; i++ {
		l := testLimits()
		fields := []*int{&l.MaxPayloadBytes, &l.MaxEnvelopeBytes, &l.MaxRows, &l.MaxAccessEntries, &l.MaxFieldBytes, &l.MaxHeaderBytes}
		*fields[i] = 0
		if l.Validate() == nil {
			t.Fatal("missing limit accepted")
		}
		if _, err := Encode(p, l); err == nil {
			t.Fatal("encode accepted missing limit")
		}
		if _, err := Decode(data, l); err == nil {
			t.Fatal("decode accepted missing limit")
		}
	}
}
