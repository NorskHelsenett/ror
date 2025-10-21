package rorconfig

import "testing"

func TestEnvironmentVariablesAdd(t *testing.T) {
	t.Run("adds new key", func(t *testing.T) {
		vars := EnvironmentVariables{}
		vars.Add("NEW_KEY")

		if !vars.IsSet("NEW_KEY") {
			t.Fatalf("expected NEW_KEY to be added")
		}

		cfg, _ := vars.GetEnvVariableConfigByKey("NEW_KEY")
		if cfg.description != "Local env variable not in central list" {
			t.Fatalf("unexpected description: %q", cfg.description)
		}
		if cfg.deprecated {
			t.Fatalf("expected deprecated to be false")
		}
	})

	t.Run("no-op when key exists", func(t *testing.T) {
		vars := EnvironmentVariables{{key: "EXISTING"}}

		vars.Add("EXISTING")

		if len(vars) != 1 {
			t.Fatalf("expected length 1, got %d", len(vars))
		}
	})
}

func TestEnvironmentVariablesGetDescription(t *testing.T) {
	vars := EnvironmentVariables{{key: "KNOWN", description: "some description"}}

	if got := vars.GetDescription("KNOWN"); got != "some description" {
		t.Fatalf("GetDescription() = %q, want %q", got, "some description")
	}

	if got := vars.GetDescription("UNKNOWN"); got != "" {
		t.Fatalf("GetDescription() for unknown key = %q, want empty", got)
	}
}
