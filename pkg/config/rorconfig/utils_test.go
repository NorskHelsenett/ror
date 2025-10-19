package rorconfig

import "testing"

func TestAnyToStringSupportedTypes(t *testing.T) {
	cases := []struct {
		name     string
		value    any
		expected string
	}{
		{name: "string", value: "hello", expected: "hello"},
		{name: "int", value: int(12), expected: "12"},
		{name: "int64", value: int64(34), expected: "34"},
		{name: "float64", value: float64(3.14), expected: "3.14"},
		{name: "float32", value: float32(2.5), expected: "2.5"},
		{name: "uint", value: uint(8), expected: "8"},
		{name: "uint64", value: uint64(16), expected: "16"},
		{name: "uint32", value: uint32(4), expected: "4"},
		{name: "bool true", value: true, expected: "true"},
		{name: "bool false", value: false, expected: "false"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := anyToString(tc.value); got != tc.expected {
				t.Fatalf("anyToString(%v) = %q, want %q", tc.value, got, tc.expected)
			}
		})
	}
}

func TestAnyToStringUnsupportedType(t *testing.T) {
	result := anyToString(struct{ field string }{field: "value"})

	if result != "" {
		t.Fatalf("anyToString() unsupported type returned %q, want empty string", result)
	}
}
