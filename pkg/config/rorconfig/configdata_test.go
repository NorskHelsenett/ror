package rorconfig

import "testing"

func TestConfigDataNumericFallback(t *testing.T) {
	cases := []struct {
		name       string
		value      ConfigData
		wantInt    int
		wantInt64  int64
		wantUint   uint
		wantUint64 uint64
		wantUint32 uint32
	}{
		{name: "non numeric", value: ConfigData("abc"), wantInt: 0, wantInt64: 0, wantUint: 0, wantUint64: 0, wantUint32: 0},
		{name: "empty", value: ConfigData(""), wantInt: 0, wantInt64: 0, wantUint: 0, wantUint64: 0, wantUint32: 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.value.Int(); got != tc.wantInt {
				t.Fatalf("Int() = %d, want %d", got, tc.wantInt)
			}
			if got := tc.value.Int64(); got != tc.wantInt64 {
				t.Fatalf("Int64() = %d, want %d", got, tc.wantInt64)
			}
			if got := tc.value.Uint(); got != tc.wantUint {
				t.Fatalf("Uint() = %d, want %d", got, tc.wantUint)
			}
			if got := tc.value.Uint64(); got != tc.wantUint64 {
				t.Fatalf("Uint64() = %d, want %d", got, tc.wantUint64)
			}
			if got := tc.value.Uint32(); got != tc.wantUint32 {
				t.Fatalf("Uint32() = %d, want %d", got, tc.wantUint32)
			}
		})
	}
}

func TestConfigDataBoolFallback(t *testing.T) {
	if ConfigData("notbool").Bool() {
		t.Fatalf("Bool() returned true for invalid input")
	}
}

func TestConfigDataFloat64SpecialValues(t *testing.T) {
	if got := ConfigData("NaN").Float64(); got != 0 {
		t.Fatalf("Float64() for NaN = %f, want 0", got)
	}
	if got := ConfigData("Inf").Float64(); got != 0 {
		t.Fatalf("Float64() for Inf = %f, want 0", got)
	}
}

func TestConfigDataFloat32SpecialValues(t *testing.T) {
	if got := ConfigData("NaN").Float32(); got != 0 {
		t.Fatalf("Float32() for NaN = %f, want 0", got)
	}
	if got := ConfigData("Inf").Float32(); got != 0 {
		t.Fatalf("Float32() for Inf = %f, want 0", got)
	}
}
