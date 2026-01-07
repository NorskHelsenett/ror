package rorconfig

import (
	"math"
	"testing"
)

type staticSecretProvider struct {
	secret string
}

func (s staticSecretProvider) GetSecret() string {
	return s.secret
}

func resetConfigState(t *testing.T) {
	t.Helper()
	config = rorConfigSet{configs: make(configsMap)}
	t.Cleanup(func() {
		config = rorConfigSet{configs: make(configsMap)}
	})
}

func TestSetAndGetString(t *testing.T) {
	resetConfigState(t)

	if IsSet(HTTP_PORT) {
		t.Fatalf("expected key to be unset")
	}

	Set(HTTP_PORT, "8080")

	if !IsSet(HTTP_PORT) {
		t.Fatalf("expected key to be set")
	}

	if got := GetString(HTTP_PORT); got != "8080" {
		t.Fatalf("GetString() = %q, want %q", got, "8080")
	}
}

func TestSetDefault(t *testing.T) {
	resetConfigState(t)

	SetDefault(HTTP_HOST, "localhost")

	if got := GetString(HTTP_HOST); got != "localhost" {
		t.Fatalf("GetString() = %q, want %q", got, "localhost")
	}
}

func TestSetDefaultDoesNotOverrideExplicitValue(t *testing.T) {
	resetConfigState(t)

	Set(HTTP_HOST, "explicit")
	SetDefault(HTTP_HOST, "default")

	if got := GetString(HTTP_HOST); got != "explicit" {
		t.Fatalf("GetString() = %q, want %q", got, "explicit")
	}
}

func TestSetDefaultUsesEnvWhenAutoloadEnabled(t *testing.T) {
	resetConfigState(t)
	AutomaticEnv()
	t.Setenv(string(HTTP_TIMEOUT), "15")

	SetDefault(HTTP_TIMEOUT, "30")

	if got := GetString(HTTP_TIMEOUT); got != "15" {
		t.Fatalf("GetString() = %q, want %q", got, "15")
	}
}

func TestSetWithProvider(t *testing.T) {
	resetConfigState(t)

	SetWithProvider(HTTP_PORT, staticSecretProvider{secret: "from-provider"})

	if got := GetString(HTTP_PORT); got != "from-provider" {
		t.Fatalf("GetString() = %q, want %q", got, "from-provider")
	}
}

func TestAutomaticEnvLoadsValues(t *testing.T) {
	resetConfigState(t)
	Set(HTTP_PORT, "default")
	t.Setenv(string(HTTP_PORT), "9090")

	AutomaticEnv()

	if got := GetString(HTTP_PORT); got != "9090" {
		t.Fatalf("GetString() = %q, want %q", got, "9090")
	}
}

func TestGetTypedValues(t *testing.T) {
	resetConfigState(t)

	Set("TEST_BOOL", "true")
	Set("TEST_INT", "42")
	Set("TEST_INT64", "9223372036854775806")
	Set("TEST_FLOAT64", "3.14")
	Set("TEST_FLOAT32", "2.5")
	Set("TEST_UINT", "7")
	Set("TEST_UINT64", "8")
	Set("TEST_UINT32", "9")

	if !GetBool("TEST_BOOL") {
		t.Fatalf("GetBool() = false, want true")
	}

	if got := GetInt("TEST_INT"); got != 42 {
		t.Fatalf("GetInt() = %d, want %d", got, 42)
	}

	if got := GetInt64("TEST_INT64"); got != 9223372036854775806 {
		t.Fatalf("GetInt64() = %d, want %d", got, 9223372036854775806)
	}

	if got := GetFloat64("TEST_FLOAT64"); math.Abs(got-3.14) > 1e-9 {
		t.Fatalf("GetFloat64() = %f, want 3.14", got)
	}

	if got := GetFloat32("TEST_FLOAT32"); math.Abs(float64(got)-2.5) > 1e-6 {
		t.Fatalf("GetFloat32() = %f, want 2.5", got)
	}

	if got := GetUint("TEST_UINT"); got != uint(7) {
		t.Fatalf("GetUint() = %d, want %d", got, 7)
	}

	if got := GetUint64("TEST_UINT64"); got != uint64(8) {
		t.Fatalf("GetUint64() = %d, want %d", got, 8)
	}

	if got := GetUint32("TEST_UINT32"); got != uint32(9) {
		t.Fatalf("GetUint32() = %d, want %d", got, 9)
	}
}

func TestGetConfigsReturnsCopy(t *testing.T) {
	resetConfigState(t)

	Set("TEST_KEY", "value")

	snapshot := GetConfigs()
	snapshot["TEST_KEY"] = ConfigData{Value: "mutated"}
	snapshot["OTHER_KEY"] = ConfigData{Value: "new"}

	if got := GetString("TEST_KEY"); got != "value" {
		t.Fatalf("GetString() = %q, want %q", got, "value")
	}

	if IsSet("OTHER_KEY") {
		t.Fatalf("expected OTHER_KEY to be absent in original config")
	}
}

func TestInitConfigLoadsEnvironment(t *testing.T) {
	resetConfigState(t)
	t.Setenv(string(HTTP_PORT), "6060")

	InitConfig()

	if got := GetString(HTTP_PORT); got != "6060" {
		t.Fatalf("GetString() = %q, want %q", got, "6060")
	}
}

func TestAddStruct(t *testing.T) {
	resetConfigState(t)

	type sample struct {
		Host string `rorconfig:"HOST"`
		Port int    `rorconfig:"PORT"`
	}

	if err := SetConfigFromStruct(sample{Host: "srv", Port: 7000}); err != nil {
		t.Fatalf("AddStruct() returned error: %v", err)
	}

	if got := GetString("HOST"); got != "srv" {
		t.Fatalf("GetString(HOST) = %q, want %q", got, "srv")
	}

	if got := GetInt("PORT"); got != 7000 {
		t.Fatalf("GetInt(PORT) = %d, want %d", got, 7000)
	}
}

func TestAddStructWithPointerInput(t *testing.T) {
	resetConfigState(t)

	type sample struct {
		Enabled *bool   `rorconfig:"ENABLED"`
		Rate    float32 `rorconfig:"RATE"`
	}

	flag := false
	s := &sample{Enabled: &flag, Rate: 2.5}

	if err := SetConfigFromStruct(s); err != nil {
		t.Fatalf("AddStruct() with pointer input returned error: %v", err)
	}

	if GetBool("ENABLED") {
		t.Fatalf("GetBool(ENABLED) = true, want false")
	}

	if got := GetFloat32("RATE"); got != 2.5 {
		t.Fatalf("GetFloat32(RATE) = %f, want %f", got, 2.5)
	}
}

func TestAddStructInvalidSource(t *testing.T) {
	resetConfigState(t)

	if err := SetConfigFromStruct(123); err == nil {
		t.Fatalf("expected AddStruct() to fail for non-struct input")
	}

	var ptr *struct {
		Value string `rorconfig:"VALUE"`
	}
	if err := SetConfigFromStruct(ptr); err == nil {
		t.Fatalf("expected AddStruct() to fail for nil pointer")
	}
}

func TestAddStructNestedFields(t *testing.T) {
	resetConfigState(t)

	type config struct {
		Logging struct {
			Level   string `rorconfig:"LOG_LEVEL"`
			Enabled bool   `rorconfig:"LOG_ENABLED"`
		}
	}

	sample := config{}
	sample.Logging.Level = "debug"
	sample.Logging.Enabled = true

	if err := SetConfigFromStruct(sample); err != nil {
		t.Fatalf("AddStruct() returned error: %v", err)
	}

	if got := GetString("LOG_LEVEL"); got != "debug" {
		t.Fatalf("GetString(LOG_LEVEL) = %q, want %q", got, "debug")
	}

	if got := GetBool("LOG_ENABLED"); !got {
		t.Fatalf("GetBool(LOG_ENABLED) = %t, want true", got)
	}
}

func TestAddStructNestedPointerFields(t *testing.T) {
	resetConfigState(t)

	type config struct {
		Telemetry *struct {
			Endpoint string `rorconfig:"TEL_ENDPOINT"`
		}
	}

	var sample config
	if err := SetConfigFromStruct(sample); err != nil {
		t.Fatalf("AddStruct() returned error for nil nested pointer: %v", err)
	}

	if IsSet("TEL_ENDPOINT") {
		t.Fatalf("expected TEL_ENDPOINT to remain unset for nil nested pointer")
	}

	sample.Telemetry = &struct {
		Endpoint string `rorconfig:"TEL_ENDPOINT"`
	}{Endpoint: "http://example"}

	if err := SetConfigFromStruct(sample); err != nil {
		t.Fatalf("AddStruct() returned error for populated nested pointer: %v", err)
	}

	if got := GetString("TEL_ENDPOINT"); got != "http://example" {
		t.Fatalf("GetString(TEL_ENDPOINT) = %q, want %q", got, "http://example")
	}
}
