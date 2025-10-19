package rorconfig

import (
	"math"
	"os"
	"path/filepath"
	"testing"
)

type stubProvider struct {
	value string
}

func (s stubProvider) GetSecret() string {
	return s.value
}

func newConfigSet() *rorConfigSet {
	return &rorConfigSet{configs: make(configsMap)}
}

func withConfigConsts(t *testing.T, entries map[string]string) {
	t.Helper()
	original := ConfigConsts
	custom := make(ConfigconstsMap)
	for key, env := range entries {
		custom[ConfigConst(key)] = ConfigConstData{value: env}
	}

	ConfigConsts = custom
	t.Cleanup(func() {
		ConfigConsts = original
	})
}

func resetGlobalConfig(t *testing.T) {
	t.Helper()
	config = rorConfigSet{configs: make(configsMap)}
}

func TestConfigsMapSetAndUnset(t *testing.T) {
	cm := make(configsMap)
	key := ConfigConst("FOO")

	if cm.IsSet(key) {
		t.Fatalf("expected key %s to be unset", key)
	}

	cm.Set(key, "bar")

	if !cm.IsSet(key) {
		t.Fatalf("expected key %s to be set", key)
	}

	if cm.IsEmpty(key) {
		t.Fatalf("expected key %s to have a value", key)
	}

	if got := cm.Get(key).String(); got != "bar" {
		t.Fatalf("Get() returned %q, want %q", got, "bar")
	}

	cm.Unset(key)

	if cm.IsSet(key) {
		t.Fatalf("expected key %s to be removed", key)
	}
}

func TestConfigsMapGetAllReturnsCopy(t *testing.T) {
	cm := make(configsMap)
	key := ConfigConst("FOO")
	cm.Set(key, "bar")

	snapshot := cm.GetAll()

	cm.Set(key, "baz")

	if got := snapshot[key].String(); got != "bar" {
		t.Fatalf("snapshot mutated to %q, want %q", got, "bar")
	}
}

func TestConfigsMapIsEmptyBehavior(t *testing.T) {
	cm := make(configsMap)
	key := ConfigConst("EMPTY_TEST")

	if !cm.IsEmpty(key) {
		t.Fatalf("IsEmpty() should be true when key is missing")
	}

	cm.Set(key, "")
	if !cm.IsEmpty(key) {
		t.Fatalf("IsEmpty() should be true when value is empty string")
	}

	cm.Set(key, "value")
	if cm.IsEmpty(key) {
		t.Fatalf("IsEmpty() should be false when value is present")
	}
}

func TestRorConfigSet_LoadEnv(t *testing.T) {
	withConfigConsts(t, map[string]string{"APP_KEY": "APP_ENV"})

	t.Run("loads existing environment variable", func(t *testing.T) {
		t.Setenv("APP_ENV", "loaded")
		rc := newConfigSet()

		rc.LoadEnv(ConfigConst("APP_KEY"))

		if got := rc.GetString(ConfigConst("APP_KEY")); got != "loaded" {
			t.Fatalf("LoadEnv() = %q, want %q", got, "loaded")
		}
	})

	t.Run("missing environment results in empty value", func(t *testing.T) {
		rc := newConfigSet()
		rc.LoadEnv(ConfigConst("APP_KEY"))

		if got := rc.GetString(ConfigConst("APP_KEY")); got != "" {
			t.Fatalf("LoadEnv() for missing env = %q, want empty string", got)
		}
	})

	t.Run("deprecated key loads value", func(t *testing.T) {
		originalConfigConsts := ConfigConsts
		ConfigConsts = ConfigconstsMap{
			ConfigConst("DEPRECATED_KEY"): {
				value:       "DEPRECATED_ENV",
				deprecated:  true,
				description: "use NEW_KEY instead",
			},
		}
		t.Cleanup(func() {
			ConfigConsts = originalConfigConsts
		})

		t.Setenv("DEPRECATED_ENV", "deprecated_value")

		rc := newConfigSet()
		rc.LoadEnv(ConfigConst("DEPRECATED_KEY"))

		if got := rc.GetString(ConfigConst("DEPRECATED_KEY")); got != "deprecated_value" {
			t.Fatalf("LoadEnv() did not load value for deprecated key, got %q", got)
		}
	})
}

func TestRorConfigSet_SetDefault(t *testing.T) {
	t.Run("applies default when value not set", func(t *testing.T) {
		rc := newConfigSet()
		rc.SetDefault(ConfigConst("DEFAULT_KEY"), "fallback")

		if got := rc.GetString(ConfigConst("DEFAULT_KEY")); got != "fallback" {
			t.Fatalf("SetDefault() = %q, want %q", got, "fallback")
		}
	})

	t.Run("keeps existing value", func(t *testing.T) {
		rc := newConfigSet()
		rc.Set(ConfigConst("DEFAULT_KEY"), "existing")
		rc.SetDefault(ConfigConst("DEFAULT_KEY"), "fallback")

		if got := rc.GetString(ConfigConst("DEFAULT_KEY")); got != "existing" {
			t.Fatalf("SetDefault() overwrote existing value: got %q", got)
		}
	})

	t.Run("autoload pulls value from env", func(t *testing.T) {
		withConfigConsts(t, map[string]string{"AUTO_KEY": "AUTO_ENV"})
		t.Setenv("AUTO_ENV", "from_env")

		rc := newConfigSet()
		rc.autoload = true
		rc.SetDefault(ConfigConst("AUTO_KEY"), "fallback")

		if got := rc.GetString(ConfigConst("AUTO_KEY")); got != "from_env" {
			t.Fatalf("SetDefault() with autoload = %q, want %q", got, "from_env")
		}
	})
}

func TestRorConfigSet_SetWithProvider(t *testing.T) {
	rc := newConfigSet()
	rc.SetWithProvider(ConfigConst("SECRET_KEY"), stubProvider{value: "secret"})

	if got := rc.GetString(ConfigConst("SECRET_KEY")); got != "secret" {
		t.Fatalf("SetWithProvider() = %q, want %q", got, "secret")
	}
}

func TestRorConfigSet_AutoLoadEnv(t *testing.T) {
	withConfigConsts(t, map[string]string{"SERVICE_KEY": "SERVICE_ENV"})
	t.Setenv("SERVICE_ENV", "9000")

	rc := newConfigSet()
	rc.Set(ConfigConst("SERVICE_KEY"), "placeholder")

	rc.AutoLoadEnv()

	if got := rc.GetString(ConfigConst("SERVICE_KEY")); got != "9000" {
		t.Fatalf("AutoLoadEnv() = %q, want %q", got, "9000")
	}

	if !rc.autoload {
		t.Fatalf("AutoLoadEnv() did not mark config set as autoload")
	}
}

func TestRorConfigSet_AutoLoadAllEnv(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, "test.env")
	if err := os.WriteFile(envFile, []byte("DOT_KEY=dot_value\n"), 0o600); err != nil {
		t.Fatalf("failed to write env file: %v", err)
	}

	t.Setenv("ENV_FILE", envFile)
	t.Setenv("DIRECT_KEY", "direct_value")

	withConfigConsts(t, map[string]string{
		"DOT_CONFIG":    "DOT_KEY",
		"DIRECT_CONFIG": "DIRECT_KEY",
		"MISSING":       "MISSING_KEY",
	})

	rc := newConfigSet()
	rc.AutoLoadAllEnv()

	if got := rc.GetString(ConfigConst("DOT_CONFIG")); got != "dot_value" {
		t.Fatalf("AutoLoadAllEnv() dot env value = %q, want %q", got, "dot_value")
	}

	if got := rc.GetString(ConfigConst("DIRECT_CONFIG")); got != "direct_value" {
		t.Fatalf("AutoLoadAllEnv() existing env value = %q, want %q", got, "direct_value")
	}

	if rc.IsSet(ConfigConst("MISSING")) {
		t.Fatalf("AutoLoadAllEnv() loaded config for missing env variable")
	}
}

func TestRorConfigSet_GetValueTriggersAutoload(t *testing.T) {
	withConfigConsts(t, map[string]string{"AUTO_KEY": "AUTO_ENV"})
	t.Setenv("AUTO_ENV", "auto_loaded")

	rc := newConfigSet()
	rc.autoload = true

	if got := rc.GetString(ConfigConst("AUTO_KEY")); got != "auto_loaded" {
		t.Fatalf("GetString() with autoload = %q, want %q", got, "auto_loaded")
	}
}

func TestGetterConversions(t *testing.T) {
	rc := newConfigSet()
	rc.Set(ConfigConst("STRING"), "value")
	rc.Set(ConfigConst("BOOL_TRUE"), "true")
	rc.Set(ConfigConst("INT"), "123")
	rc.Set(ConfigConst("INT64"), "9223372036854775807")
	rc.Set(ConfigConst("FLOAT64"), "3.25")
	rc.Set(ConfigConst("FLOAT32"), "2.5")
	rc.Set(ConfigConst("UINT"), "42")
	rc.Set(ConfigConst("UINT64"), "100")
	rc.Set(ConfigConst("UINT32"), "24")

	if got := rc.GetString(ConfigConst("STRING")); got != "value" {
		t.Fatalf("GetString() = %q, want %q", got, "value")
	}

	if !rc.GetBool(ConfigConst("BOOL_TRUE")) {
		t.Fatalf("GetBool() expected true")
	}

	if got := rc.GetInt(ConfigConst("INT")); got != 123 {
		t.Fatalf("GetInt() = %d, want %d", got, 123)
	}

	if got := rc.GetInt64(ConfigConst("INT64")); got != 9223372036854775807 {
		t.Fatalf("GetInt64() = %d, want %d", got, 9223372036854775807)
	}

	if got := rc.GetFloat64(ConfigConst("FLOAT64")); got != 3.25 {
		t.Fatalf("GetFloat64() = %f, want %f", got, 3.25)
	}

	if got := rc.GetFloat32(ConfigConst("FLOAT32")); got != 2.5 {
		t.Fatalf("GetFloat32() = %f, want %f", got, 2.5)
	}

	if got := rc.GetUint(ConfigConst("UINT")); got != 42 {
		t.Fatalf("GetUint() = %d, want %d", got, 42)
	}

	if got := rc.GetUint64(ConfigConst("UINT64")); got != 100 {
		t.Fatalf("GetUint64() = %d, want %d", got, 100)
	}

	if got := rc.GetUint32(ConfigConst("UINT32")); got != 24 {
		t.Fatalf("GetUint32() = %d, want %d", got, 24)
	}
}

func TestGetterConversionsInvalidInput(t *testing.T) {
	rc := newConfigSet()
	rc.Set(ConfigConst("BAD_BOOL"), "not-bool")
	rc.Set(ConfigConst("BAD_INT"), "not-int")
	rc.Set(ConfigConst("BAD_FLOAT64"), math.NaN())
	rc.Set(ConfigConst("BAD_FLOAT32"), math.Inf(1))
	rc.Set(ConfigConst("BAD_UINT"), "-1")

	if rc.GetBool(ConfigConst("BAD_BOOL")) {
		t.Fatalf("GetBool() should return false for invalid input")
	}

	if got := rc.GetInt(ConfigConst("BAD_INT")); got != 0 {
		t.Fatalf("GetInt() = %d, want 0", got)
	}

	if got := rc.GetFloat64(ConfigConst("BAD_FLOAT64")); got != 0 {
		t.Fatalf("GetFloat64() for NaN = %f, want 0", got)
	}

	if got := rc.GetFloat32(ConfigConst("BAD_FLOAT32")); got != 0 {
		t.Fatalf("GetFloat32() for Inf = %f, want 0", got)
	}

	if got := rc.GetUint(ConfigConst("BAD_UINT")); got != 0 {
		t.Fatalf("GetUint() for invalid input = %d, want 0", got)
	}
}

func TestSetHandlesUnsupportedType(t *testing.T) {
	rc := newConfigSet()
	rc.Set(ConfigConst("UNSUPPORTED"), struct{}{})

	if got := rc.GetString(ConfigConst("UNSUPPORTED")); got != "" {
		t.Fatalf("Set() with unsupported type should store empty string, got %q", got)
	}
}

func TestPackageLevelHelpers(t *testing.T) {
	resetGlobalConfig(t)

	Set(ConfigConst("STRING_KEY"), "value")
	if !IsSet(ConfigConst("STRING_KEY")) {
		t.Fatalf("IsSet() reported false for existing value")
	}

	if got := GetString(ConfigConst("STRING_KEY")); got != "value" {
		t.Fatalf("GetString() = %q, want %q", got, "value")
	}

	SetDefault(ConfigConst("DEFAULT_KEY"), "fallback")
	if got := GetString(ConfigConst("DEFAULT_KEY")); got != "fallback" {
		t.Fatalf("SetDefault() via package function failed, got %q", got)
	}

	AutomaticEnv()

	SetWithProvider(ConfigConst("SECRET"), stubProvider{value: "secret"})
	if got := GetString(ConfigConst("SECRET")); got != "secret" {
		t.Fatalf("SetWithProvider() via package function = %q, want %q", got, "secret")
	}
}

func TestAutomaticEnvReloadsExistingKeys(t *testing.T) {
	resetGlobalConfig(t)

	withConfigConsts(t, map[string]string{"GLOBAL_KEY": "GLOBAL_ENV"})
	Set(ConfigConst("GLOBAL_KEY"), "initial")
	t.Setenv("GLOBAL_ENV", "from_env")

	AutomaticEnv()

	if got := GetString(ConfigConst("GLOBAL_KEY")); got != "from_env" {
		t.Fatalf("AutomaticEnv() did not refresh value from environment: got %q", got)
	}
}

func TestInitConfigLoadsEnvironment(t *testing.T) {
	resetGlobalConfig(t)

	withConfigConsts(t, map[string]string{"APP_KEY": "APP_ENV"})
	t.Setenv("APP_ENV", "from_env")

	InitConfig()

	if got := GetString(ConfigConst("APP_KEY")); got != "from_env" {
		t.Fatalf("InitConfig() = %q, want %q", got, "from_env")
	}
}

func TestConfigConstsHelpers(t *testing.T) {
	original := ConfigConsts
	custom := ConfigconstsMap{
		ConfigConst("KNOWN"): {
			value:       "KNOWN_ENV",
			deprecated:  false,
			description: "known entry",
		},
		ConfigConst("DEPRECATED"): {
			value:       "DEPRECATED_ENV",
			deprecated:  true,
			description: "deprecated entry",
		},
	}

	ConfigConsts = custom
	t.Cleanup(func() {
		ConfigConsts = original
	})

	if got := ConfigConsts.GetConfigConstByName("KNOWN_ENV"); got != ConfigConst("KNOWN") {
		t.Fatalf("GetConfigConstByName() = %q, want %q", got, ConfigConst("KNOWN"))
	}

	if got := ConfigConsts.GetConfigConstByKey(ConfigConst("KNOWN")); got != ConfigConst("KNOWN") {
		t.Fatalf("GetConfigConstByKey() = %q, want %q", got, ConfigConst("KNOWN"))
	}

	ConfigConsts.Add(ConfigConst("KNOWN"), ConfigConstData{value: "SHOULD_NOT_OVERRIDE"})
	if got := ConfigConsts[ConfigConst("KNOWN")].value; got != "KNOWN_ENV" {
		t.Fatalf("Add() should not override existing entry, got %q", got)
	}

	localKey := ConfigConst("LOCAL_ONLY")
	if got := ConfigConsts.GetEnvVariable(localKey); got != string(localKey) {
		t.Fatalf("GetEnvVariable() for local key = %q, want %q", got, string(localKey))
	}

	if got := ConfigConsts.GetDescription(localKey); got != "Local env variable not in central list" {
		t.Fatalf("GetDescription() for local key = %q, want default description", got)
	}

	if _, exists := ConfigConsts[localKey]; !exists {
		t.Fatalf("ensureKeyExists() should add local key to ConfigConsts")
	}

	if !ConfigConsts.IsDeprecated(ConfigConst("DEPRECATED")) {
		t.Fatalf("IsDeprecated() should return true for deprecated entries")
	}
}

func TestPackageLevelGettersCoverAllTypes(t *testing.T) {
	resetGlobalConfig(t)

	Set(ConfigConst("BOOL_KEY"), "true")
	Set(ConfigConst("INT_KEY"), "101")
	Set(ConfigConst("INT64_KEY"), "9223372036854775806")
	Set(ConfigConst("FLOAT64_KEY"), "6.28")
	Set(ConfigConst("FLOAT32_KEY"), "1.5")
	Set(ConfigConst("UINT_KEY"), "7")
	Set(ConfigConst("UINT64_KEY"), "42")
	Set(ConfigConst("UINT32_KEY"), "24")

	if !GetBool(ConfigConst("BOOL_KEY")) {
		t.Fatalf("GetBool() should return true")
	}

	if got := GetInt(ConfigConst("INT_KEY")); got != 101 {
		t.Fatalf("GetInt() = %d, want %d", got, 101)
	}

	if got := GetInt64(ConfigConst("INT64_KEY")); got != 9223372036854775806 {
		t.Fatalf("GetInt64() = %d, want %d", got, 9223372036854775806)
	}

	if got := GetFloat64(ConfigConst("FLOAT64_KEY")); got != 6.28 {
		t.Fatalf("GetFloat64() = %f, want %f", got, 6.28)
	}

	if got := GetFloat32(ConfigConst("FLOAT32_KEY")); got != 1.5 {
		t.Fatalf("GetFloat32() = %f, want %f", got, 1.5)
	}

	if got := GetUint(ConfigConst("UINT_KEY")); got != 7 {
		t.Fatalf("GetUint() = %d, want %d", got, 7)
	}

	if got := GetUint64(ConfigConst("UINT64_KEY")); got != 42 {
		t.Fatalf("GetUint64() = %d, want %d", got, 42)
	}

	if got := GetUint32(ConfigConst("UINT32_KEY")); got != 24 {
		t.Fatalf("GetUint32() = %d, want %d", got, 24)
	}

	snapshot := GetConfigs()
	if got := snapshot[ConfigConst("INT_KEY")]; got.String() != "101" {
		t.Fatalf("GetConfigs() snapshot missing expected value, got %q", got.String())
	}

	snapshot[ConfigConst("INT_KEY")] = ConfigData("mutated")
	if got := GetInt(ConfigConst("INT_KEY")); got != 101 {
		t.Fatalf("GetConfigs() should return copy; mutation leaked with value %d", got)
	}
}
