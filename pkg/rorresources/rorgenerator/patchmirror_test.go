package rorgenerator

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResourceFixture struct {
	Name     string             `json:"name,omitempty"`
	Nested   ResourceFixtureSub `json:"nested"`
	Labels   map[string]string  `json:"labels"`
	Items    []string           `json:"items"`
	Pointer  *string            `json:"pointer"`
	Stamp    time.Time          `json:"stamp"`
	Meta     metav1.ObjectMeta  `json:"meta"`
	Opaque   ResourceFixtureRaw `json:"opaque"`
	Untagged string
	Excluded string `json:"excluded"`
	Secret   string `json:"-"`
	Callback func() `json:"callback"`
	hidden   string
}

type ResourceFixtureSub struct {
	Value string `json:"value"`
}

// ResourceFixtureRaw marshals itself, so the generator must treat it as a leaf
// instead of taking its fields apart.
type ResourceFixtureRaw struct {
	Raw string `json:"raw"`
}

func (ResourceFixtureRaw) MarshalJSON() ([]byte, error) { return []byte(`""`), nil }

type FixtureSub struct {
	Other string `json:"other"`
}

func fixtureConfig() PatchMirrorConfig {
	return PatchMirrorConfig{
		Package:          "rorgenerator",
		PackagePath:      "github.com/NorskHelsenett/ror/pkg/rorresources/rorgenerator",
		Prefix:           "Patch",
		SourceFilePrefix: "resourcedef_",
		FilePrefix:       "resourcepatchdef_",
		OptionalPackage:  "github.com/NorskHelsenett/ror/pkg/helpers/optional",
		NilTargetError:   "ErrNilTarget",
		GeneratedBy:      "cmd/generator/main.go",
		Skip:             map[string]bool{"ResourceFixture.Excluded": true},
	}
}

// fixtureDeclarations stands in for parsing a package, so the rendering tests
// do not need a package on disk.
func fixtureDeclarations() map[string]string {
	return map[string]string{
		"ResourceFixture":    "resourcedef_fixture.go",
		"ResourceFixtureSub": "resourcedef_fixture.go",
		"ResourceFixtureRaw": "resourcedef_fixture.go",
		"FixtureSub":         "resourcedef_fixture.go",
	}
}

// renderFixture renders the fixture with every type declared in the same file,
// so that the whole output can be asserted on as one string.
func renderFixture(t *testing.T) string {
	t.Helper()

	files, err := renderPatchMirrorFiles([]PatchMirrorRoot{
		{Kind: "Fixture", Type: reflect.TypeOf(ResourceFixture{})},
	}, fixtureConfig(), fixtureDeclarations())
	require.NoError(t, err)
	require.Len(t, files, 1)
	return files["resourcepatchdef_fixture.go"]
}

func TestPatchMirrorFieldKinds(t *testing.T) {
	source := renderFixture(t)

	tests := []struct {
		name  string
		field string
		apply string
	}{
		{
			name:  "a scalar is a leaf",
			field: "Name optional.Optional[string] `json:\"name,omitzero\"`",
			apply: "optional.Apply(p.Name, &target.Name)",
		},
		{
			name:  "a struct of this package is mirrored",
			field: "Nested optional.Optional[PatchFixtureSub] `json:\"nested,omitzero\"`",
			apply: "optional.ApplyNested(p.Nested, &target.Nested)",
		},
		{
			name:  "a map is patched per key",
			field: "Labels optional.Optional[optional.Map[string, string]] `json:\"labels,omitzero\"`",
			apply: "optional.ApplyNested(p.Labels, &target.Labels)",
		},
		{
			name:  "a slice is replaced as a whole",
			field: "Items optional.Optional[[]string] `json:\"items,omitzero\"`",
			apply: "optional.Apply(p.Items, &target.Items)",
		},
		{
			name:  "a pointer is a leaf",
			field: "Pointer optional.Optional[*string] `json:\"pointer,omitzero\"`",
			apply: "optional.Apply(p.Pointer, &target.Pointer)",
		},
		{
			name:  "a struct of another package is a leaf",
			field: "Meta optional.Optional[metav1.ObjectMeta] `json:\"meta,omitzero\"`",
			apply: "optional.Apply(p.Meta, &target.Meta)",
		},
		{
			name:  "a struct that marshals itself is a leaf",
			field: "Opaque optional.Optional[ResourceFixtureRaw] `json:\"opaque,omitzero\"`",
			apply: "optional.Apply(p.Opaque, &target.Opaque)",
		},
		{
			name:  "an untagged field keeps its go name",
			field: "Untagged optional.Optional[string] `json:\"Untagged,omitzero\"`",
			apply: "optional.Apply(p.Untagged, &target.Untagged)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Contains(t, source, test.field)
			assert.Contains(t, source, test.apply)
		})
	}
}

func TestPatchMirrorSkipsUnpatchableFields(t *testing.T) {
	source := renderFixture(t)

	assert.NotContains(t, source, "Secret", "a field tagged json:\"-\" is not patchable")
	assert.NotContains(t, source, "hidden", "an unexported field is not patchable")
	assert.NotContains(t, source, "Callback", "a func field has no shape a patch can describe")
	assert.NotContains(t, source, "Excluded", "a field listed in Skip is not patchable")
	assert.NotContains(t, source, "PatchFixtureRaw", "a leaf struct must not get a mirror of its own")
}

func TestPatchMirrorHeaderAndImports(t *testing.T) {
	source := renderFixture(t)

	assert.Contains(t, source, "// THIS FILE IS GENERATED, DO NOT EDIT\n// ref: cmd/generator/main.go")
	assert.Contains(t, source, "package rorgenerator")
	assert.Contains(t, source, "\t\"time\"\n")
	assert.Contains(t, source, "\t\"github.com/NorskHelsenett/ror/pkg/helpers/optional\"\n")
	assert.Contains(t, source, "\tmetav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"\n", "a trailing version segment is prefixed with the segment before it")
	assert.Contains(t, source, "type FixturePatchinterface interface {\n\tApplyTo(target *ResourceFixture) error\n}")
	assert.Contains(t, source, "func (p PatchFixture) ApplyTo(target *ResourceFixture) error {\n\tif target == nil {\n\t\treturn ErrNilTarget\n\t}")
}

func TestPatchMirrorFollowsTheDeclaringFile(t *testing.T) {
	declared := fixtureDeclarations()
	declared["ResourceFixtureSub"] = "resourcedef_fixturesub.go"

	files, err := renderPatchMirrorFiles([]PatchMirrorRoot{
		{Kind: "Fixture", Type: reflect.TypeOf(ResourceFixture{})},
	}, fixtureConfig(), declared)
	require.NoError(t, err)

	require.Len(t, files, 2, "a nested struct declared elsewhere gets its own file")
	assert.Contains(t, files["resourcepatchdef_fixture.go"], "type PatchFixture struct")
	assert.Contains(t, files["resourcepatchdef_fixture.go"], "type FixturePatchinterface interface")
	assert.Contains(t, files["resourcepatchdef_fixturesub.go"], "type PatchFixtureSub struct")
	assert.NotContains(t, files["resourcepatchdef_fixture.go"], "type PatchFixtureSub struct")
	assert.NotContains(t, files["resourcepatchdef_fixturesub.go"], "metav1", "a file only imports what it uses")
}

func TestPatchMirrorNeedsToKnowWhereATypeIsDeclared(t *testing.T) {
	_, err := renderPatchMirrorFiles([]PatchMirrorRoot{
		{Kind: "Fixture", Type: reflect.TypeOf(ResourceFixture{})},
	}, fixtureConfig(), map[string]string{})

	assert.ErrorContains(t, err, "could not find where ResourceFixtureSub is declared")
}

func TestPatchMirrorRejectsCollidingNames(t *testing.T) {
	// ResourceFixtureSub and FixtureSub would both mirror to PatchFixtureSub.
	_, err := renderPatchMirrorFiles([]PatchMirrorRoot{
		{Kind: "FixtureSub", Type: reflect.TypeOf(ResourceFixtureSub{})},
		{Kind: "Other", Type: reflect.TypeOf(FixtureSub{})},
	}, fixtureConfig(), fixtureDeclarations())

	assert.ErrorContains(t, err, "claimed by both")
}

func TestPatchMirrorDeclaresASharedStructOnce(t *testing.T) {
	declared := fixtureDeclarations()
	declared["ResourceFixtureSub"] = "resourcedef_fixturesub.go"

	files, err := renderPatchMirrorFiles([]PatchMirrorRoot{
		{Kind: "Fixture", Type: reflect.TypeOf(ResourceFixture{})},
		{Kind: "FixtureSub", Type: reflect.TypeOf(ResourceFixtureSub{})},
	}, fixtureConfig(), declared)
	require.NoError(t, err)

	var declarations int
	for _, source := range files {
		declarations += countOccurrences(source, "type PatchFixtureSub struct")
	}
	assert.Equal(t, 1, declarations, "a struct reached from two kinds is declared once")
	assert.Contains(t, files["resourcepatchdef_fixturesub.go"], "type FixtureSubPatchinterface interface")
}

type ResourceFixtureEmpty struct{}

func TestPatchMirrorOfAnEmptyStructCompiles(t *testing.T) {
	files, err := renderPatchMirrorFiles([]PatchMirrorRoot{
		{Kind: "FixtureEmpty", Type: reflect.TypeOf(ResourceFixtureEmpty{})},
	}, fixtureConfig(), map[string]string{"ResourceFixtureEmpty": "resourcedef_fixtureempty.go"})
	require.NoError(t, err)

	source := files["resourcepatchdef_fixtureempty.go"]
	assert.Contains(t, source, "type PatchFixtureEmpty struct {\n}")
	assert.NotContains(t, source, "optional\"", "a mirror with no field must not import optional")
}

func TestDeclaringFiles(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "resourcedef_one.go", "package p\n\ntype One struct{}\n\ntype AlsoOne int\n")
	write(t, dir, "resourcedef_two.go", "package p\n\ntype Two struct{}\n")
	write(t, dir, "resourcedef_one_test.go", "package p\n\ntype NotIncluded struct{}\n")

	declared, err := declaringFiles(dir)
	require.NoError(t, err)

	assert.Equal(t, map[string]string{
		"One":     "resourcedef_one.go",
		"AlsoOne": "resourcedef_one.go",
		"Two":     "resourcedef_two.go",
	}, declared, "a test file holds no definition to mirror")
}

func TestRefuseToClobberProtectsHandWrittenFiles(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "generated.go", generatedMarker+"\npackage p\n")
	write(t, dir, "handwritten.go", "package p\n")

	assert.NoError(t, refuseToClobber(filepath.Join(dir, "missing.go")))
	assert.NoError(t, refuseToClobber(filepath.Join(dir, "generated.go")))
	assert.ErrorContains(t, refuseToClobber(filepath.Join(dir, "handwritten.go")), "exists and is not generated")
}

func TestRemoveStaleFilesKeepsWhatItDidNotWrite(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "resourcepatchdef_current.go", generatedMarker+"\npackage p\n")
	write(t, dir, "resourcepatchdef_stale.go", generatedMarker+"\npackage p\n")
	write(t, dir, "resourcepatchdef_common.go", "package p\n")
	write(t, dir, "resourcedef_one.go", generatedMarker+"\npackage p\n")

	cfg := fixtureConfig()
	cfg.SourceDir = dir
	require.NoError(t, removeStaleFiles(cfg, map[string]string{"resourcepatchdef_current.go": ""}))

	assert.FileExists(t, filepath.Join(dir, "resourcepatchdef_current.go"))
	assert.FileExists(t, filepath.Join(dir, "resourcepatchdef_common.go"), "a hand written mirror is not removed")
	assert.FileExists(t, filepath.Join(dir, "resourcedef_one.go"), "a file outside the prefix is not removed")
	assert.NoFileExists(t, filepath.Join(dir, "resourcepatchdef_stale.go"))
}

func write(t *testing.T, dir string, name string, content string) {
	t.Helper()
	require.NoError(t, os.WriteFile(filepath.Join(dir, name), []byte(content), 0o600))
}

func countOccurrences(haystack string, needle string) int {
	count := 0
	for i := 0; i+len(needle) <= len(haystack); i++ {
		if haystack[i:i+len(needle)] == needle {
			count++
		}
	}
	return count
}
