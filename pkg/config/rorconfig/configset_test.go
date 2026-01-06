package rorconfig

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

type addStructSample struct {
	Host     string  `rorconfig:"HOST"`
	Port     int     `rorconfig:"PORT"`
	Enabled  *bool   `rorconfig:"ENABLED"`
	Optional *string `rorconfig:"OPTIONAL"`
	Ignored  string
}

type addStructWide struct {
	BoolValue  bool    `rorconfig:"BOOL_VALUE"`
	FloatValue float64 `rorconfig:"FLOAT_VALUE"`
	UintValue  uint32  `rorconfig:"UINT_VALUE"`
	StringPtr  *string `rorconfig:"STRING_PTR"`
	BoolPtr    *bool   `rorconfig:"BOOL_PTR"`
	Untagged   float64
}

func TestConfigsMapIsEmpty(t *testing.T) {
	cm := make(configsMap)

	if !cm.IsEmpty("missing") {
		t.Fatalf("expected missing key to be empty")
	}

	cm.Set("empty", "")
	if !cm.IsEmpty("empty") {
		t.Fatalf("expected key with empty string to be considered empty")
	}

	cm.Set("value", "something")
	if cm.IsEmpty("value") {
		t.Fatalf("expected key with value to be non-empty")
	}
}

func TestConfigsMapUnset(t *testing.T) {
	cm := make(configsMap)
	cm.Set("key", "value")

	cm.Unset("key")

	if cm.IsSet("key") {
		t.Fatalf("expected key to be removed")
	}
}

func TestRorConfigSetLoadEnvAddsUnknownKey(t *testing.T) {
	originalConsts := append(EnvironmentVariables(nil), ConfigConsts...)
	t.Cleanup(func() {
		ConfigConsts = append(EnvironmentVariables(nil), originalConsts...)
	})

	rc := rorConfigSet{configs: make(configsMap)}
	const key = "UNIT_TEST_LOAD_ENV"
	t.Setenv(key, "value-from-env")

	rc.LoadEnv(key)

	if got := rc.configs.Get(key).String(); got != "value-from-env" {
		t.Fatalf("LoadEnv() stored %q, want %q", got, "value-from-env")
	}

	if !ConfigConsts.IsSet(key) {
		t.Fatalf("expected key to be registered in ConfigConsts")
	}

	cfg, _ := ConfigConsts.GetEnvVariableConfigByKey(key)
	if cfg.description != "Local env variable not in central list" {
		t.Fatalf("unexpected description for added key: %q", cfg.description)
	}
	if cfg.deprecated {
		t.Fatalf("expected added key to be non-deprecated")
	}
}

func TestRorConfigSetLoadEnvDeprecatedKey(t *testing.T) {
	originalConsts := append(EnvironmentVariables(nil), ConfigConsts...)
	deprecatedKey := EnvironmentVariableConfig{key: "UNIT_TEST_DEPRECATED", deprecated: true, description: "use something else"}
	ConfigConsts = EnvironmentVariables{deprecatedKey}
	t.Cleanup(func() {
		ConfigConsts = append(EnvironmentVariables(nil), originalConsts...)
	})

	rc := rorConfigSet{configs: make(configsMap)}
	t.Setenv(deprecatedKey.key, "deprecated-value")

	rc.LoadEnv(deprecatedKey.key)

	if got := rc.GetString(deprecatedKey.key); got != "deprecated-value" {
		t.Fatalf("LoadEnv() stored %q, want %q", got, "deprecated-value")
	}

	if desc := ConfigConsts.GetDescription(deprecatedKey.key); desc != "use something else" {
		t.Fatalf("expected description to remain unchanged, got %q", desc)
	}
}

func TestRorConfigSetSetDefaultOverridesEmpty(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}
	rc.Set("EMPTY_KEY", "")

	rc.SetDefault("EMPTY_KEY", "fallback")

	if got := rc.GetString("EMPTY_KEY"); got != "fallback" {
		t.Fatalf("SetDefault() stored %q, want %q", got, "fallback")
	}
}

func TestRorConfigSetAutoLoadEnv(t *testing.T) {
	originalConsts := append(EnvironmentVariables(nil), ConfigConsts...)
	t.Cleanup(func() {
		ConfigConsts = append(EnvironmentVariables(nil), originalConsts...)
	})

	rc := rorConfigSet{configs: make(configsMap)}
	const key = "UNIT_TEST_AUTOLOAD"
	rc.Set(key, "initial")
	t.Setenv(key, "from-env")

	rc.AutoLoadEnv()

	if !rc.autoload {
		t.Fatalf("expected autoload flag to be set")
	}

	if got := rc.GetString(key); got != "from-env" {
		t.Fatalf("AutoLoadEnv() stored %q, want %q", got, "from-env")
	}
}

func TestRorConfigSetGetStringTriggersAutoload(t *testing.T) {
	originalConsts := append(EnvironmentVariables(nil), ConfigConsts...)
	t.Cleanup(func() {
		ConfigConsts = append(EnvironmentVariables(nil), originalConsts...)
	})

	rc := rorConfigSet{configs: make(configsMap)}
	rc.autoload = true
	const key = "UNIT_TEST_LAZY_LOAD"
	t.Setenv(key, "lazy-value")

	if got := rc.GetString(key); got != "lazy-value" {
		t.Fatalf("GetString() returned %q, want %q", got, "lazy-value")
	}

	if !rc.IsSet(key) {
		t.Fatalf("expected key to be populated after lazy load")
	}
}

func TestRorConfigSetAutoLoadAllEnvWithLocalKeys(t *testing.T) {
	originalConsts := append(EnvironmentVariables(nil), ConfigConsts...)
	t.Cleanup(func() {
		ConfigConsts = append(EnvironmentVariables(nil), originalConsts...)
	})

	rc := rorConfigSet{configs: make(configsMap)}
	const key = "UNIT_TEST_AUTOLOAD_ALL"
	t.Setenv(key, "from-env-all")

	rc.AutoLoadAllEnv(key)

	if got := rc.GetString(key); got != "from-env-all" {
		t.Fatalf("AutoLoadAllEnv() stored %q, want %q", got, "from-env-all")
	}

	if !ConfigConsts.IsSet(key) {
		t.Fatalf("expected local key to be added to ConfigConsts")
	}
}

type fakeProvider struct {
	value string
}

func (f fakeProvider) GetSecret() string { return f.value }

func TestRorConfigSetSetWithProvider(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}
	const key = "PROVIDER_KEY"

	rc.SetWithProvider(key, fakeProvider{value: "secret"})

	if got := rc.GetString(key); got != "secret" {
		t.Fatalf("SetWithProvider() stored %q, want %q", got, "secret")
	}
}

func TestRorConfigSetAutoLoadAllEnvLoadsDotEnv(t *testing.T) {
	originalConsts := append(EnvironmentVariables(nil), ConfigConsts...)
	t.Cleanup(func() {
		ConfigConsts = append(EnvironmentVariables(nil), originalConsts...)
	})

	rc := rorConfigSet{configs: make(configsMap)}
	const key = "DOTENV_KEY"
	ConfigConsts = EnvironmentVariables{{key: key}}

	tempDir := t.TempDir()
	envFile := filepath.Join(tempDir, "test.env")
	if err := os.WriteFile(envFile, []byte("DOTENV_KEY=from-dotenv\n"), 0o600); err != nil {
		t.Fatalf("failed to write env file: %v", err)
	}

	t.Setenv("ENV_FILE", envFile)

	rc.AutoLoadAllEnv()

	if got := rc.GetString(key); got != "from-dotenv" {
		t.Fatalf("AutoLoadAllEnv() loaded %q, want %q", got, "from-dotenv")
	}
}

func TestRorConfigSetAddStruct(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}
	flag := true
	sample := addStructSample{
		Host:    "localhost",
		Port:    8080,
		Enabled: &flag,
		Ignored: "noop",
	}

	if err := rc.ImportStruct(&sample); err != nil {
		t.Fatalf("AddStruct() returned error: %v", err)
	}

	if got := rc.GetString("HOST"); got != "localhost" {
		t.Fatalf("AddStruct() stored HOST=%q, want %q", got, "localhost")
	}

	if got := rc.GetInt("PORT"); got != 8080 {
		t.Fatalf("AddStruct() stored PORT=%d, want %d", got, 8080)
	}

	if got := rc.GetBool("ENABLED"); !got {
		t.Fatalf("AddStruct() stored ENABLED=%t, want true", got)
	}

	if rc.IsSet("OPTIONAL") {
		t.Fatalf("expected OPTIONAL to remain unset when source field is nil")
	}

	second := addStructSample{Host: "example", Port: 9090}
	if err := rc.ImportStruct(second); err != nil {
		t.Fatalf("AddStruct() with struct value failed: %v", err)
	}

	if got := rc.GetString("HOST"); got != "example" {
		t.Fatalf("AddStruct() did not overwrite HOST, got %q", got)
	}

	if rc.IsSet("Ignored") {
		t.Fatalf("field without tag should not populate configs")
	}
}

func TestRorConfigSetAddStructWithWideTypes(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}
	text := "configured"
	boolValue := false
	wide := addStructWide{
		BoolValue:  true,
		FloatValue: 3.1415,
		UintValue:  23,
		StringPtr:  &text,
		BoolPtr:    &boolValue,
	}

	if err := rc.ImportStruct(wide); err != nil {
		t.Fatalf("AddStruct() returned error: %v", err)
	}

	if got := rc.GetBool("BOOL_VALUE"); !got {
		t.Fatalf("AddStruct() stored BOOL_VALUE=%t, want true", got)
	}

	if got := rc.GetFloat64("FLOAT_VALUE"); got != 3.1415 {
		t.Fatalf("AddStruct() stored FLOAT_VALUE=%f, want %f", got, 3.1415)
	}

	if got := rc.GetUint32("UINT_VALUE"); got != 23 {
		t.Fatalf("AddStruct() stored UINT_VALUE=%d, want %d", got, 23)
	}

	if got := rc.GetString("STRING_PTR"); got != "configured" {
		t.Fatalf("AddStruct() stored STRING_PTR=%q, want %q", got, "configured")
	}

	if rc.GetBool("BOOL_PTR") {
		t.Fatalf("AddStruct() should store false for bool pointer")
	}

	if rc.IsSet("Untagged") {
		t.Fatalf("untagged fields should not populate configs")
	}
}

func TestRorConfigSetAddStructNestedStruct(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}
	sample := struct {
		Database struct {
			Host string `rorconfig:"DB_HOST"`
			SSL  *bool  `rorconfig:"DB_SSL"`
		}
	}{}

	enabled := true
	sample.Database.Host = "db.local"
	sample.Database.SSL = &enabled

	if err := rc.ImportStruct(sample); err != nil {
		t.Fatalf("AddStruct() returned error: %v", err)
	}

	if got := rc.GetString("DB_HOST"); got != "db.local" {
		t.Fatalf("AddStruct() stored DB_HOST=%q, want %q", got, "db.local")
	}

	if got := rc.GetBool("DB_SSL"); !got {
		t.Fatalf("AddStruct() stored DB_SSL=%t, want true", got)
	}
}

func TestRorConfigSetAddStructNestedPointer(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}

	sample := struct {
		Telemetry *struct {
			Endpoint string `rorconfig:"TEL_ENDPOINT"`
		}
	}{}

	if err := rc.ImportStruct(sample); err != nil {
		t.Fatalf("AddStruct() returned error for nil nested pointer: %v", err)
	}

	if rc.IsSet("TEL_ENDPOINT") {
		t.Fatalf("expected TEL_ENDPOINT to remain unset for nil nested pointer")
	}

	sample.Telemetry = &struct {
		Endpoint string `rorconfig:"TEL_ENDPOINT"`
	}{Endpoint: "http://example"}

	if err := rc.ImportStruct(sample); err != nil {
		t.Fatalf("AddStruct() returned error for populated nested pointer: %v", err)
	}

	if got := rc.GetString("TEL_ENDPOINT"); got != "http://example" {
		t.Fatalf("AddStruct() stored TEL_ENDPOINT=%q, want %q", got, "http://example")
	}
}

func TestRorConfigSetAddStructUnsupportedKind(t *testing.T) {
	type invalid struct {
		Numbers []int `rorconfig:"NUMBERS"`
	}

	rc := rorConfigSet{configs: make(configsMap)}

	if err := rc.ImportStruct(invalid{Numbers: []int{1}}); err == nil {
		t.Fatalf("expected error for unsupported slice field")
	}
}

func TestRorConfigSetAddStructNilPointer(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}

	var sample *addStructSample
	if err := rc.ImportStruct(sample); err == nil {
		t.Fatalf("expected error when source is nil pointer")
	}
}

func TestRorConfigSetAddStructNonStruct(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}

	if err := rc.ImportStruct(123); err == nil {
		t.Fatalf("expected error when source is not a struct")
	}
}

func TestRorConfigSetAddStructMultiLevelPointer(t *testing.T) {
	type pointer struct {
		Flag **bool `rorconfig:"FLAG"`
	}

	value := true
	ptr := &value
	sample := pointer{Flag: &ptr}

	rc := rorConfigSet{configs: make(configsMap)}

	if err := rc.ImportStruct(sample); err != nil {
		t.Fatalf("AddStruct() returned error for multi-level pointer: %v", err)
	}

	if got := rc.GetBool("FLAG"); !got {
		t.Fatalf("AddStruct() stored FLAG=%t, want true", got)
	}
}

func TestRorConfigSetAddStructTaggedStructField(t *testing.T) {
	type invalid struct {
		Nested struct {
			Value string `rorconfig:"VALUE"`
		} `rorconfig:"NESTED"`
	}

	rc := rorConfigSet{configs: make(configsMap)}

	sample := invalid{}
	sample.Nested.Value = "test"

	if err := rc.ImportStruct(sample); err == nil {
		t.Fatalf("expected error when struct field has rorconfig tag")
	}
}

func TestRorConfigSetSetAndGetTime(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}
	timestamp := time.Date(2024, time.January, 15, 10, 30, 45, 123456789, time.UTC)

	rc.Set("TIMESTAMP", timestamp)

	if got := rc.configs.Get("TIMESTAMP").String(); got != timestamp.Format(time.RFC3339Nano) {
		t.Fatalf("Set() stored %q, want %q", got, timestamp.Format(time.RFC3339Nano))
	}

	if got := rc.GetTime("TIMESTAMP"); !got.Equal(timestamp) {
		t.Fatalf("GetTime() returned %v, want %v", got, timestamp)
	}
}

func TestRorConfigSetImportStructWithTime(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}
	timestamp := time.Date(2023, time.March, 3, 8, 0, 0, 0, time.UTC)

	sample := struct {
		StartedAt time.Time `rorconfig:"STARTED_AT"`
	}{StartedAt: timestamp}

	if err := rc.ImportStruct(sample); err != nil {
		t.Fatalf("ImportStruct() returned error: %v", err)
	}

	if got := rc.configs.Get("STARTED_AT").String(); got != timestamp.Format(time.RFC3339Nano) {
		t.Fatalf("ImportStruct() stored %q, want %q", got, timestamp.Format(time.RFC3339Nano))
	}

	if got := rc.GetTime("STARTED_AT"); !got.Equal(timestamp) {
		t.Fatalf("GetTime() returned %v, want %v", got, timestamp)
	}
}
