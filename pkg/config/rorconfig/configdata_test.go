package rorconfig

import "testing"

func TestConfigDataNumericConversions(t *testing.T) {
	cd := ConfigData("314")

	if got := cd.Int(); got != 314 {
		t.Fatalf("Int() = %d, want %d", got, 314)
	}

	cd64 := ConfigData("9223372036854775807")
	if got := cd64.Int64(); got != 9223372036854775807 {
		t.Fatalf("Int64() = %d, want %d", got, 9223372036854775807)
	}

	cdFloat32 := ConfigData("12.75")
	if got := cdFloat32.Float32(); got != 12.75 {
		t.Fatalf("Float32() = %f, want %f", got, 12.75)
	}

	cdUint := ConfigData("123")
	if got := cdUint.Uint(); got != 123 {
		t.Fatalf("Uint() = %d, want %d", got, 123)
	}

	cdUint64 := ConfigData("9876543210")
	if got := cdUint64.Uint64(); got != 9876543210 {
		t.Fatalf("Uint64() = %d, want %d", got, 9876543210)
	}

	cdUint32 := ConfigData("4294967295")
	if got := cdUint32.Uint32(); got != 4294967295 {
		t.Fatalf("Uint32() = %d, want %d", got, 4294967295)
	}
}

func TestConfigDataBoolConversion(t *testing.T) {
	trueValues := []string{"true", "TRUE", "1"}
	for _, val := range trueValues {
		if got := ConfigData(val).Bool(); !got {
			t.Fatalf("Bool() = false for value %q, want true", val)
		}
	}

	falseValues := []string{"false", "FALSE", "0", "not-bool"}
	for _, val := range falseValues {
		if got := ConfigData(val).Bool(); got {
			t.Fatalf("Bool() = true for value %q, want false", val)
		}
	}
}
