package rorgenerator

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

// The patch mirror generator turns a resource definition struct into a struct
// where every field is an optional.Optional, so that a partial update can tell
// "not set" from "set to null" from "set to a value".
//
// The walk follows a few rules:
//
//   - A struct declared in the same package is mirrored, so its fields can be
//     patched one by one.
//   - A map is patched per key through optional.Map, because a key is a stable
//     address for an element.
//   - Everything else is replaced as a whole: slices,
//     pointers, structs from other packages and structs with custom json
//     marshalling.
//
// A mirror is written to the file matching the one its source type is declared
// in, so resourcedef_kubernetescluster.go gets a
// resourcepatchdef_kubernetescluster.go beside it.

const generatedMarker = "// THIS FILE IS GENERATED, DO NOT EDIT"

var (
	jsonMarshalerType = reflect.TypeFor[json.Marshaler]()
	versionSegment    = regexp.MustCompile(`^v[0-9]`)
	nonIdentifier     = regexp.MustCompile(`[^a-zA-Z0-9_]`)
	// genericPathSegments are path elements that say nothing about the package
	// and are skipped when an import alias is derived from the path.
	genericPathSegments = map[string]bool{"pkg": true, "pkgs": true, "api": true, "apis": true, "internal": true, "src": true}
)

// PatchMirrorRoot names one resource definition to mirror.
type PatchMirrorRoot struct {
	// Type is the resource definition struct, e.g.
	// reflect.TypeOf(rortypes.ResourceKubernetesCluster{}).
	Type reflect.Type
	// Kind is the ror kind the definition belongs to, used to name the
	// generated interface.
	Kind string
}

// PatchMirrorConfig configures the generated files.
type PatchMirrorConfig struct {
	// Package is the name of the package the files are generated into.
	Package string
	// PackagePath is the import path of that package. Structs from this path
	// are mirrored, structs from anywhere else are leaves.
	PackagePath string
	// SourceDir is the directory holding the package, both the source that is
	// parsed to find where each type is declared and the destination of the
	// generated files.
	SourceDir string
	// Prefix is prepended to the mirrored type names, after a leading
	// "Resource" has been stripped, e.g. ResourceKubernetesClusterSpec becomes
	// ResourcePatchKubernetesClusterSpec.
	Prefix string
	// SourceFilePrefix is stripped from the name of the file a type is
	// declared in before the generated name is built, e.g. "resourcedef_".
	SourceFilePrefix string
	// FilePrefix is prepended to the generated file names, e.g.
	// "resourcepatchdef_".
	FilePrefix string
	// OptionalPackage is the import path of the optional package.
	OptionalPackage string
	// NilTargetError names the error returned when ApplyTo is given a nil
	// target. It must be declared in the generated package.
	NilTargetError string
	// Opaque lists type names in PackagePath that are leaves even though they
	// are structs declared in that package.
	Opaque map[string]bool
	// Skip lists fields that are not patchable, keyed as "TypeName.FieldName".
	Skip map[string]bool
	// Aliases overrides the derived import alias for an import path.
	Aliases map[string]string
	// GeneratedBy is written into the header of the generated files.
	GeneratedBy string
}

// PatchMirrorFiles generates the patch mirrors for every root, one file per
// file of the source package.
//
// Stale files from kinds that's no longer patchable are removed.
// A file that exists without being generated is not touched, so hand written mirrors can live next to the generated ones.
func (g Generator) PatchMirrorFiles(roots []PatchMirrorRoot, cfg PatchMirrorConfig) error {
	declaredIn, err := declaringFiles(cfg.SourceDir)
	if err != nil {
		return err
	}

	files, err := renderPatchMirrorFiles(roots, cfg, declaredIn)
	if err != nil {
		return err
	}

	if err := removeStaleFiles(cfg, files); err != nil {
		return err
	}

	for _, name := range sortedKeys(files) {
		path := filepath.Join(cfg.SourceDir, name)
		if err := refuseToClobber(path); err != nil {
			return err
		}
		if err := g.WriteGoFile(path, files[name]); err != nil {
			return err
		}
	}
	return nil
}

type mirrorType struct {
	name   string
	source reflect.Type
	fields []mirrorField
}

type mirrorField struct {
	name     string
	jsonName string
	// source is the type of the field on the mirrored struct.
	source reflect.Type
	// mirrorName is set when the field is a struct with a mirror of its own.
	mirrorName string
	// isMap is set when the field is patched per key through optional.Map.
	isMap bool
}

// nested reports whether the field is applied with ApplyNested rather than
// Apply, which is the case for mirrored structs and for maps.
func (f mirrorField) nested() bool {
	return f.mirrorName != "" || f.isMap
}

type mirrorGenerator struct {
	visited map[reflect.Type]string
	cfg     PatchMirrorConfig
	byName  map[string]reflect.Type
	types   []mirrorType
}

// renderPatchMirrorFiles returns the generated source keyed by file name.
func renderPatchMirrorFiles(roots []PatchMirrorRoot, cfg PatchMirrorConfig, declaredIn map[string]string) (map[string]string, error) {
	gen := &mirrorGenerator{
		cfg:     cfg,
		byName:  map[string]reflect.Type{},
		visited: map[reflect.Type]string{},
	}

	rootNames := map[string]string{}
	for _, root := range roots {
		if root.Type == nil || root.Type.Kind() != reflect.Struct {
			return nil, fmt.Errorf("patch mirror root %q is not a struct", root.Kind)
		}
		name, err := gen.mirror(root.Type)
		if err != nil {
			return nil, err
		}
		rootNames[root.Kind] = name
	}

	// Mirrors and interfaces are grouped by the file their source type is
	// declared in, so a mirror lands beside the definition it mirrors.
	mirrorsByFile := map[string][]mirrorType{}
	for _, mirror := range gen.types {
		file, err := cfg.outputFile(mirror.source, declaredIn)
		if err != nil {
			return nil, err
		}
		mirrorsByFile[file] = append(mirrorsByFile[file], mirror)
	}

	interfacesByFile := map[string][]PatchMirrorRoot{}
	for _, root := range roots {
		file, err := cfg.outputFile(root.Type, declaredIn)
		if err != nil {
			return nil, err
		}
		interfacesByFile[file] = append(interfacesByFile[file], root)
	}

	files := map[string]string{}
	for file, mirrors := range mirrorsByFile {
		// Sorted by name rather than left in walk order, so that making
		// another kind patchable cannot reshuffle a file that already exists.
		sort.Slice(mirrors, func(i, j int) bool { return mirrors[i].name < mirrors[j].name })

		interfaces := interfacesByFile[file]
		sort.Slice(interfaces, func(i, j int) bool { return interfaces[i].Kind < interfaces[j].Kind })

		files[file] = gen.renderFile(mirrors, interfaces, rootNames)
	}
	return files, nil
}

// outputFile returns the name of the generated file a type belongs in, derived
// from the file it is declared in.
func (cfg PatchMirrorConfig) outputFile(t reflect.Type, declaredIn map[string]string) (string, error) {
	source, ok := declaredIn[t.Name()]
	if !ok {
		return "", fmt.Errorf("could not find where %s is declared in %s", t.Name(), cfg.SourceDir)
	}
	stem := strings.TrimSuffix(source, ".go")
	return cfg.FilePrefix + strings.TrimPrefix(stem, cfg.SourceFilePrefix) + ".go", nil
}

func (g *mirrorGenerator) renderFile(mirrors []mirrorType, interfaces []PatchMirrorRoot, rootNames map[string]string) string {
	imports := newImportSet(g.cfg.PackagePath, g.cfg.Aliases)

	// A mirror of a struct with nothing patchable in it holds no optional
	// field, so the import is only needed once a field is there to use it.
	var optional string
	for _, mirror := range mirrors {
		if len(mirror.fields) > 0 {
			optional = imports.add(g.cfg.OptionalPackage)
			break
		}
	}

	var body strings.Builder
	for _, mirror := range mirrors {
		g.writeType(&body, imports, optional, mirror)
	}
	for _, root := range interfaces {
		writeInterface(&body, root, rootNames[root.Kind])
	}

	var out strings.Builder
	fmt.Fprintf(&out, "%s\n// ref: %s\n\n", generatedMarker, g.cfg.GeneratedBy)
	fmt.Fprintf(&out, "package %s\n\n", g.cfg.Package)
	out.WriteString(imports.render())
	out.WriteString(patchMirrorDoc)
	out.WriteString(body.String())
	return out.String()
}

const patchMirrorDoc = `
// This file holds a patch mirror of a resource. a patch mirror mirrors every field
// but the field are a generic type of optional.Optional, which carries the three
// states a partial update needs: unset, null, and a value. 
// See the optional package for why every field is tagged omitzero.

`

// mirror returns the name of the mirror of t, generating it when it is new.
func (g *mirrorGenerator) mirror(t reflect.Type) (string, error) {
	if name, ok := g.visited[t]; ok {
		return name, nil
	}

	name := g.cfg.Prefix + strings.TrimPrefix(t.Name(), "Resource")
	if existing, taken := g.byName[name]; taken {
		return "", fmt.Errorf("patch mirror name %q is claimed by both %s and %s", name, existing.Name(), t.Name())
	}
	g.byName[name] = t
	// Recorded before the fields are walked, so a type that refers to itself
	// resolves to the mirror being built instead of recursing forever.
	g.visited[t] = name

	mirror := mirrorType{name: name, source: t}
	for field := range t.Fields() {
		if field.PkgPath != "" {
			continue
		}

		jsonName, ok := jsonFieldName(field)
		if !ok {
			continue
		}
		if g.cfg.Skip[t.Name()+"."+field.Name] {
			continue
		}

		mirrored, ok, err := g.field(field.Type)
		if err != nil {
			return "", fmt.Errorf("%s.%s: %w", t.Name(), field.Name, err)
		}
		if !ok {
			continue
		}

		mirrored.name = field.Name
		mirrored.jsonName = jsonName
		mirror.fields = append(mirror.fields, mirrored)
	}

	g.types = append(g.types, mirror)
	return name, nil
}

// field decides how a field is patched, and whether it is patchable at all.
func (g *mirrorGenerator) field(t reflect.Type) (mirrorField, bool, error) {
	if g.isMirrored(t) {
		name, err := g.mirror(t)
		if err != nil {
			return mirrorField{}, false, err
		}
		return mirrorField{source: t, mirrorName: name}, true, nil
	}

	if t.Kind() == reflect.Map && t.Name() == "" {
		if !canRender(t.Key(), g.cfg) || !canRender(t.Elem(), g.cfg) {
			return mirrorField{}, false, nil
		}
		return mirrorField{source: t, isMap: true}, true, nil
	}

	return mirrorField{source: t}, canRender(t, g.cfg), nil
}

// isMirrored reports whether t is a struct this package owns and can take
// apart field by field.
func (g *mirrorGenerator) isMirrored(t reflect.Type) bool {
	if t.Kind() != reflect.Struct || t.Name() == "" || t.PkgPath() != g.cfg.PackagePath {
		return false
	}
	if g.cfg.Opaque[t.Name()] {
		return false
	}
	// A type that marshals itself decides its own json shape, so mirroring its
	// fields would produce a patch document that does not match it.
	return !t.Implements(jsonMarshalerType) && !reflect.PointerTo(t).Implements(jsonMarshalerType)
}

// canRender reports whether a type can be written as a go type expression,
// without registering the imports it would need.
func canRender(t reflect.Type, cfg PatchMirrorConfig) bool {
	_, ok := renderType(t, newImportSet(cfg.PackagePath, cfg.Aliases), cfg.PackagePath)
	return ok
}

// renderType writes t as a go type expression, registering the imports it
// needs. It reports false for types a patch cannot carry.
func renderType(t reflect.Type, imports *importSet, selfPath string) (string, bool) {
	if t.Name() != "" {
		if t.PkgPath() == "" {
			// A builtin such as string or int.
			return t.Name(), true
		}
		if t.PkgPath() == selfPath {
			return t.Name(), true
		}
		return imports.add(t.PkgPath()) + "." + t.Name(), true
	}

	switch t.Kind() {
	case reflect.Pointer:
		inner, ok := renderType(t.Elem(), imports, selfPath)
		return "*" + inner, ok
	case reflect.Slice:
		inner, ok := renderType(t.Elem(), imports, selfPath)
		return "[]" + inner, ok
	case reflect.Array:
		inner, ok := renderType(t.Elem(), imports, selfPath)
		return fmt.Sprintf("[%d]%s", t.Len(), inner), ok
	case reflect.Map:
		key, keyOk := renderType(t.Key(), imports, selfPath)
		value, valueOk := renderType(t.Elem(), imports, selfPath)
		return fmt.Sprintf("map[%s]%s", key, value), keyOk && valueOk
	default:
		// Anonymous structs, interfaces, channels and functions have no shape
		// a patch can describe.
		return "", false
	}
}

// fieldExpr returns the type to put inside the optional for a field.
func (g *mirrorGenerator) fieldExpr(field mirrorField, imports *importSet, optional string) string {
	if field.mirrorName != "" {
		return field.mirrorName
	}
	if field.isMap {
		key, _ := renderType(field.source.Key(), imports, g.cfg.PackagePath)
		value, _ := renderType(field.source.Elem(), imports, g.cfg.PackagePath)
		return fmt.Sprintf("%s.Map[%s, %s]", optional, key, value)
	}
	expr, _ := renderType(field.source, imports, g.cfg.PackagePath)
	return expr
}

func (g *mirrorGenerator) writeType(out *strings.Builder, imports *importSet, optional string, mirror mirrorType) {
	fmt.Fprintf(out, "// %s is the patch mirror of %s.\n", mirror.name, mirror.source.Name())
	fmt.Fprintf(out, "type %s struct {\n", mirror.name)
	for _, field := range mirror.fields {
		expr := g.fieldExpr(field, imports, optional)
		fmt.Fprintf(out, "\t%s %s.Optional[%s] `json:\"%s,omitzero\"`\n", field.name, optional, expr, field.jsonName)
	}
	out.WriteString("}\n\n")

	// TODO: Figure out if ApplyTo is necessary
	fmt.Fprintf(out, "// ApplyTo writes the patch into target, leaving the fields the patch does\n// not mention untouched.\n")
	fmt.Fprintf(out, "func (p %s) ApplyTo(target *%s) error {\n", mirror.name, mirror.source.Name())
	fmt.Fprintf(out, "\tif target == nil {\n\t\treturn %s\n\t}\n", g.cfg.NilTargetError)
	for _, field := range mirror.fields {
		if field.nested() {
			fmt.Fprintf(out, "\tif err := %s.ApplyNested(p.%s, &target.%s); err != nil {\n\t\treturn err\n\t}\n", optional, field.name, field.name)
			continue
		}
		fmt.Fprintf(out, "\t%s.Apply(p.%s, &target.%s)\n", optional, field.name, field.name)
	}
	out.WriteString("\treturn nil\n}\n\n")
}

func writeInterface(out *strings.Builder, root PatchMirrorRoot, mirrorName string) {
	fmt.Fprintf(out, "// %sPatchinterface represents the interface for patches of the type %s,\n// implemented by %s.\n", root.Kind, strings.ToLower(root.Kind), mirrorName)
	fmt.Fprintf(out, "type %sPatchinterface interface {\n", root.Kind)
	fmt.Fprintf(out, "\tApplyTo(target *%s) error\n}\n\n", root.Type.Name())
}

// declaringFiles maps every type name declared in a package to the base name
// of the file it is declared in.
func declaringFiles(dir string) (map[string]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %w", dir, err)
	}

	fset := token.NewFileSet()
	declared := map[string]string{}
	for _, entry := range entries {
		name := entry.Name()

		// Skip directories, non-go files and test files
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}

		file, err := parser.ParseFile(fset, filepath.Join(dir, name), nil, parser.SkipObjectResolution)
		if err != nil {
			return nil, fmt.Errorf("could not parse %s: %w", name, err)
		}
		for _, decl := range file.Decls {
			generic, ok := decl.(*ast.GenDecl)
			if !ok || generic.Tok != token.TYPE {
				continue
			}
			for _, spec := range generic.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					declared[typeSpec.Name.Name] = name
				}
			}
		}
	}
	return declared, nil
}

// removeStaleFiles deletes generated files that this run did not produce.
// This is to delete generated files from kinds that are no longer patchable to avoid stale files being left behind.
func removeStaleFiles(cfg PatchMirrorConfig, produced map[string]string) error {
	entries, err := os.ReadDir(cfg.SourceDir)
	if err != nil {
		return fmt.Errorf("could not read %s: %w", cfg.SourceDir, err)
	}

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasPrefix(name, cfg.FilePrefix) || !strings.HasSuffix(name, ".go") {
			continue
		}
		if _, ok := produced[name]; ok {
			continue
		}

		path := filepath.Join(cfg.SourceDir, name)
		generated, err := isGenerated(path)
		if err != nil {
			return err
		}
		if !generated {
			continue
		}
		if err := os.Remove(path); err != nil {
			return fmt.Errorf("could not remove stale %s: %w", path, err)
		}
		fmt.Println("Removed stale file: ", path)
	}
	return nil
}

// refuseToClobber reports an error when a file exists and was not generated,
// so that a hand written mirror is never overwritten.
func refuseToClobber(path string) error {
	if _, err := os.Stat(path); err != nil {
		return nil
	}
	generated, err := isGenerated(path)
	if err != nil {
		return err
	}
	if !generated {
		return fmt.Errorf("%s exists and is not generated, rename it or change the file prefix", path)
	}
	return nil
}

func isGenerated(path string) (bool, error) {
	content, err := os.ReadFile(path) // #nosec G304 - the path is built from the configured source dir
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("could not read %s: %w", path, err)
	}
	return strings.HasPrefix(string(content), generatedMarker), nil
}

// jsonFieldName returns the name the field is encoded under, and whether the
// field is encoded at all.
func jsonFieldName(field reflect.StructField) (string, bool) {
	tag, ok := field.Tag.Lookup("json")
	if !ok {
		return field.Name, true
	}

	name, _, _ := strings.Cut(tag, ",")
	if name == "-" {
		return "", false
	}
	if name == "" {
		// Covers `json:",inline"` on an embedded field too: a patch document
		// is its own contract, so the field is nested under its own name
		// rather than inlined.
		return field.Name, true
	}
	return name, true
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// importSet collects the imports a generated file needs and hands out a unique
// alias per import path.
type importSet struct {
	taken     map[string]bool
	aliases   map[string]string
	overrides map[string]string
	self      string
}

func newImportSet(self string, overrides map[string]string) *importSet {
	return &importSet{
		self:      self,
		overrides: overrides,
		aliases:   map[string]string{},
		taken:     map[string]bool{},
	}
}

// add registers an import path and returns the name it is referenced by.
func (s *importSet) add(path string) string {
	if path == "" || path == s.self {
		return ""
	}
	if alias, ok := s.aliases[path]; ok {
		return alias
	}

	alias := s.overrides[path]
	if alias == "" {
		alias = deriveAlias(path)
	}
	for suffix := 2; s.taken[alias]; suffix++ {
		alias = fmt.Sprintf("%s%d", deriveAlias(path), suffix)
	}

	s.taken[alias] = true
	s.aliases[path] = alias
	return alias
}

// deriveAlias builds a package name from an import path. A trailing version
// segment carries no meaning on its own, so it is prefixed with the first
// meaningful segment before it, turning .../meta/v1 into metav1.
func deriveAlias(path string) string {
	segments := strings.Split(path, "/")
	last := segments[len(segments)-1]
	if !versionSegment.MatchString(last) {
		return nonIdentifier.ReplaceAllString(last, "")
	}

	for i := len(segments) - 2; i >= 0; i-- {
		segment := nonIdentifier.ReplaceAllString(segments[i], "")
		if segment == "" || genericPathSegments[segments[i]] {
			continue
		}
		return segment + last
	}
	return nonIdentifier.ReplaceAllString(last, "")
}

// render writes the import block, with the standard library in its own group.
func (s *importSet) render() string {
	if len(s.aliases) == 0 {
		return ""
	}

	var stdlib, external []string
	for path := range s.aliases {
		if isStdlib(path) {
			stdlib = append(stdlib, path)
			continue
		}
		external = append(external, path)
	}
	sort.Strings(stdlib)
	sort.Strings(external)

	var out strings.Builder
	out.WriteString("import (\n")
	for _, group := range [][]string{stdlib, external} {
		if len(group) == 0 {
			continue
		}
		if out.Len() > len("import (\n") {
			out.WriteString("\n")
		}
		for _, path := range group {
			if alias := s.aliases[path]; alias != defaultPackageName(path) {
				fmt.Fprintf(&out, "\t%s %q\n", alias, path)
				continue
			}
			fmt.Fprintf(&out, "\t%q\n", path)
		}
	}
	out.WriteString(")\n\n")
	return out.String()
}

func defaultPackageName(path string) string {
	segments := strings.Split(path, "/")
	return segments[len(segments)-1]
}

// isStdlib reports whether an import path belongs to the standard library,
// which has no dot in its first path segment.
func isStdlib(path string) bool {
	first, _, _ := strings.Cut(path, "/")
	return !strings.Contains(first, ".")
}
