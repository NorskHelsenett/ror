package rorconfig

import (
	"os"
	"path/filepath"
	"testing"
)

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
