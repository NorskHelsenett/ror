package rorconfig

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

// --- test struct types ---

type fileConfig struct {
	Host    string `yaml:"host" rorconfig:"FILE_HOST"`
	Port    int    `yaml:"port" rorconfig:"FILE_PORT"`
	Verbose bool   `yaml:"verbose" rorconfig:"FILE_VERBOSE"`
}

type nestedFileConfig struct {
	Database struct {
		DSN     string `yaml:"dsn" rorconfig:"DB_DSN"`
		MaxConn int    `yaml:"max_conn" rorconfig:"DB_MAX_CONN"`
	} `yaml:"database"`
}

type wideExportConfig struct {
	Name     string        `rorconfig:"EXP_NAME"`
	Count    int64         `rorconfig:"EXP_COUNT"`
	Rate     float64       `rorconfig:"EXP_RATE"`
	Active   bool          `rorconfig:"EXP_ACTIVE"`
	Size     uint32        `rorconfig:"EXP_SIZE"`
	Started  time.Time     `rorconfig:"EXP_STARTED"`
	Interval time.Duration `rorconfig:"EXP_INTERVAL"`
}

type pointerExportConfig struct {
	Host *string `rorconfig:"PTR_HOST"`
	Port *int    `rorconfig:"PTR_PORT"`
}

type nestedExportConfig struct {
	Database struct {
		Host string `rorconfig:"NEST_DB_HOST"`
	}
	Cache *struct {
		TTL string `rorconfig:"NEST_CACHE_TTL"`
	}
}

// ---- LoadFromFile tests ----

func TestLoadFromFile(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}
	origConfig := config
	config = rc
	t.Cleanup(func() { config = origConfig })

	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	if err := os.WriteFile(path, []byte("host: example.com\nport: 443\nverbose: true\n"), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}

	if err := LoadFromFile[fileConfig](path); err != nil {
		t.Fatalf("LoadFromFile() error: %v", err)
	}

	// All values accessible through the store.
	if got := config.GetString("FILE_HOST"); got != "example.com" {
		t.Fatalf("store FILE_HOST = %q, want %q", got, "example.com")
	}
	if got := config.GetInt("FILE_PORT"); got != 443 {
		t.Fatalf("store FILE_PORT = %d, want %d", got, 443)
	}
	if got := config.GetBool("FILE_VERBOSE"); !got {
		t.Fatalf("store FILE_VERBOSE = %t, want true", got)
	}
}

func TestLoadFromFileNested(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}
	origConfig := config
	config = rc
	t.Cleanup(func() { config = origConfig })

	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	content := "database:\n  dsn: postgres://localhost/db\n  max_conn: 25\n"
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}

	if err := LoadFromFile[nestedFileConfig](path); err != nil {
		t.Fatalf("LoadFromFile() error: %v", err)
	}

	if got := config.GetString("DB_DSN"); got != "postgres://localhost/db" {
		t.Fatalf("store DB_DSN = %q, want %q", got, "postgres://localhost/db")
	}
	if got := config.GetInt("DB_MAX_CONN"); got != 25 {
		t.Fatalf("store DB_MAX_CONN = %d, want %d", got, 25)
	}
}

func TestLoadFromFileMissingFile(t *testing.T) {
	if err := LoadFromFile[fileConfig]("/nonexistent/path.yaml"); err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestLoadFromFileInvalidYAML(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.yaml")
	if err := os.WriteFile(path, []byte(":\t:\n\t:bad"), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}

	if err := LoadFromFile[fileConfig](path); err == nil {
		t.Fatal("expected error for invalid YAML")
	}
}

// ---- SaveToFile tests ----

func TestSaveToFile(t *testing.T) {
	origConfig := config
	config = rorConfigSet{configs: make(configsMap)}
	t.Cleanup(func() { config = origConfig })

	config.Set("FILE_HOST", "saved.local")
	config.Set("FILE_PORT", 8080)
	config.Set("FILE_VERBOSE", true)

	dir := t.TempDir()
	path := filepath.Join(dir, "out.yaml")

	if err := SaveToFile[fileConfig](path); err != nil {
		t.Fatalf("SaveToFile() error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read: %v", err)
	}

	content := string(data)
	for _, want := range []string{"host: saved.local", "port: 8080", "verbose: true"} {
		if !contains(content, want) {
			t.Fatalf("SaveToFile() output missing %q:\n%s", want, content)
		}
	}

	// Verify file permissions are 0600.
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat: %v", err)
	}
	if perm := info.Mode().Perm(); perm != 0o600 {
		t.Fatalf("file permissions = %o, want 0600", perm)
	}
}

func TestSaveToFileRoundTrip(t *testing.T) {
	origConfig := config
	config = rorConfigSet{configs: make(configsMap)}
	t.Cleanup(func() { config = origConfig })

	// Set values via the store.
	config.Set("FILE_HOST", "round.trip")
	config.Set("FILE_PORT", 9999)
	config.Set("FILE_VERBOSE", false)

	dir := t.TempDir()
	path := filepath.Join(dir, "roundtrip.yaml")

	if err := SaveToFile[fileConfig](path); err != nil {
		t.Fatalf("SaveToFile() error: %v", err)
	}

	// Reset store and reload from file.
	config = rorConfigSet{configs: make(configsMap)}
	if err := LoadFromFile[fileConfig](path); err != nil {
		t.Fatalf("LoadFromFile() error: %v", err)
	}

	if got := config.GetString("FILE_HOST"); got != "round.trip" {
		t.Fatalf("round-trip FILE_HOST = %q, want %q", got, "round.trip")
	}
	if got := config.GetInt("FILE_PORT"); got != 9999 {
		t.Fatalf("round-trip FILE_PORT = %d, want %d", got, 9999)
	}
}

func TestSaveToFileOmitsRuntimeOnlyKeys(t *testing.T) {
	origConfig := config
	config = rorConfigSet{configs: make(configsMap)}
	t.Cleanup(func() { config = origConfig })

	config.Set("FILE_HOST", "visible")
	config.Set("FILE_PORT", 80)
	config.Set("SECRET_TOKEN", "should-not-appear") // runtime-only, not in fileConfig

	dir := t.TempDir()
	path := filepath.Join(dir, "out.yaml")

	if err := SaveToFile[fileConfig](path); err != nil {
		t.Fatalf("SaveToFile() error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read: %v", err)
	}

	if contains(string(data), "should-not-appear") {
		t.Fatalf("runtime-only key leaked to file:\n%s", data)
	}
}

// ---- ExportToStruct tests ----

func TestExportToStructBasic(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}
	rc.Set("EXP_NAME", "test-service")
	rc.Set("EXP_COUNT", 42)
	rc.Set("EXP_RATE", 3.14)
	rc.Set("EXP_ACTIVE", true)
	rc.Set("EXP_SIZE", uint32(256))

	ts := time.Date(2025, time.June, 15, 12, 0, 0, 0, time.UTC)
	rc.Set("EXP_STARTED", ts)
	rc.Set("EXP_INTERVAL", time.Second*30)

	cfg, err := ExportToStruct[wideExportConfig](&rc)
	if err != nil {
		t.Fatalf("ExportToStruct() error: %v", err)
	}

	if cfg.Name != "test-service" {
		t.Fatalf("Name = %q, want %q", cfg.Name, "test-service")
	}
	if cfg.Count != 42 {
		t.Fatalf("Count = %d, want %d", cfg.Count, 42)
	}
	if cfg.Rate != 3.14 {
		t.Fatalf("Rate = %f, want %f", cfg.Rate, 3.14)
	}
	if !cfg.Active {
		t.Fatalf("Active = %t, want true", cfg.Active)
	}
	if cfg.Size != 256 {
		t.Fatalf("Size = %d, want %d", cfg.Size, 256)
	}
	if !cfg.Started.Equal(ts) {
		t.Fatalf("Started = %v, want %v", cfg.Started, ts)
	}
	if cfg.Interval != 30*time.Second {
		t.Fatalf("Interval = %v, want %v", cfg.Interval, 30*time.Second)
	}
}

func TestExportToStructPointerFields(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}
	rc.Set("PTR_HOST", "ptr.example")
	rc.Set("PTR_PORT", 5432)

	cfg, err := ExportToStruct[pointerExportConfig](&rc)
	if err != nil {
		t.Fatalf("ExportToStruct() error: %v", err)
	}

	if cfg.Host == nil || *cfg.Host != "ptr.example" {
		t.Fatalf("Host = %v, want %q", cfg.Host, "ptr.example")
	}
	// Port is *int — the conversion goes through Int64 then SetInt, verify via pointer.
	if cfg.Port == nil {
		t.Fatalf("Port is nil, want 5432")
	}
}

func TestExportToStructNestedStruct(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}
	rc.Set("NEST_DB_HOST", "db.example")
	rc.Set("NEST_CACHE_TTL", "30s")

	cfg, err := ExportToStruct[nestedExportConfig](&rc)
	if err != nil {
		t.Fatalf("ExportToStruct() error: %v", err)
	}

	if cfg.Database.Host != "db.example" {
		t.Fatalf("Database.Host = %q, want %q", cfg.Database.Host, "db.example")
	}
	if cfg.Cache == nil || cfg.Cache.TTL != "30s" {
		t.Fatalf("Cache.TTL = %v, want %q", cfg.Cache, "30s")
	}
}

func TestExportToStructMissingKeys(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}
	// Only set one of the fields.
	rc.Set("EXP_NAME", "partial")

	cfg, err := ExportToStruct[wideExportConfig](&rc)
	if err != nil {
		t.Fatalf("ExportToStruct() error: %v", err)
	}

	if cfg.Name != "partial" {
		t.Fatalf("Name = %q, want %q", cfg.Name, "partial")
	}
	// Unset fields should be zero values.
	if cfg.Count != 0 {
		t.Fatalf("Count = %d, want 0", cfg.Count)
	}
}

func TestExportToStructImportRoundTrip(t *testing.T) {
	rc := rorConfigSet{configs: make(configsMap)}

	// Set values via the store (simulating what LoadFromFile or Set would do).
	rc.Set("FILE_HOST", "rt.local")
	rc.Set("FILE_PORT", 7777)
	rc.Set("FILE_VERBOSE", true)

	exported, err := ExportToStruct[fileConfig](&rc)
	if err != nil {
		t.Fatalf("ExportToStruct() error: %v", err)
	}

	if exported.Host != "rt.local" || exported.Port != 7777 || !exported.Verbose {
		t.Fatalf("round-trip mismatch: got %+v", exported)
	}
}

func TestGetConfigToStruct(t *testing.T) {
	origConfig := config
	config = rorConfigSet{configs: make(configsMap)}
	t.Cleanup(func() { config = origConfig })

	config.Set("FILE_HOST", "public-api")
	config.Set("FILE_PORT", 443)
	config.Set("FILE_VERBOSE", true)

	cfg, err := GetConfigToStruct[fileConfig]()
	if err != nil {
		t.Fatalf("GetConfigToStruct() error: %v", err)
	}

	if cfg.Host != "public-api" {
		t.Fatalf("Host = %q, want %q", cfg.Host, "public-api")
	}
	if cfg.Port != 443 {
		t.Fatalf("Port = %d, want %d", cfg.Port, 443)
	}
	if !cfg.Verbose {
		t.Fatalf("Verbose = %t, want true", cfg.Verbose)
	}
}

// ---- helper ----

func contains(s, substr string) bool {
	return len(s) >= len(substr) && searchSubstring(s, substr)
}

func searchSubstring(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
