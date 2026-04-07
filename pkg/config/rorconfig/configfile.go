package rorconfig

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// LoadFromFile reads a YAML file at the given path, unmarshals it into T
// (which defines the file schema via yaml tags), and imports the rorconfig-tagged
// fields into the configuration store with ConfigSourceConfigFile.
// After loading, values are accessed via the rorconfig.Get* methods.
func LoadFromFile[T any](path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("rorconfig: reading %s: %w", path, err)
	}

	cfg := new(T)
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("rorconfig: unmarshaling %s: %w", path, err)
	}

	if err := config.ImportStruct(cfg); err != nil {
		return fmt.Errorf("rorconfig: importing struct from %s: %w", path, err)
	}

	return nil
}

// SaveToFile exports the current configuration store into a new instance of T
// (populated via rorconfig tags), marshals it as YAML, and writes it to the
// given path with 0600 permissions. Only struct fields participate in the
// marshal, so runtime-only keys that live only in the store are never
// persisted, and fields removed from the struct type are automatically pruned.
// fileSaveExcludeSources lists the config sources that should never be
// written to a config file. Env overrides and CLI flags are transient
// and must not leak into persisted YAML.
var fileSaveExcludeSources = []ConfigSource{ConfigSourceEnv, ConfigSourceFlag}

func SaveToFile[T any](path string) error {
	cfg, err := exportToStructFiltered[T](&config, fileSaveExcludeSources)
	if err != nil {
		return fmt.Errorf("rorconfig: exporting to struct: %w", err)
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("rorconfig: marshaling config: %w", err)
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("rorconfig: writing %s: %w", path, err)
	}

	return nil
}
