// The Generator package provides a way to generate code for collecting,
// transfering and saving resources in the agent and api.
// It also provides functions to fetch the resources.
//
//		go run build/generator/main.go
//
//	  - pkg/apicontracts/apiresourcecontracts/resource_models_generated.go
//	  - pkg/apicontracts/apiresourcecontracts/resource_models_methods_generated.go
//	  - pkg/rorresources/fromstruct.go
//	  - pkg/rorresources/resource.go
//	  - pkg/rorresources/rorkubernetes/k8s_test.go
//	  - pkg/rorresources/rorkubernetes/k8s.go
//	  - pkg/rorresources/rortypes/resource_interfaces.go
//	  - pkg/rorresources/rortypes/resource_models_methods.go
//	  - pkg/rorresources/rortypes/resource_input_filter_*.go
//	  - pkg/rorresources/rortypes/resourcepatchdef_<definition>.go
//	  - pkg/rorresources/resourcepatch.go
//
// The repo github.com/NorskHelsenett/ROR, should be placed in the ../ror folder or the ROR_PATH environment variable should be set to the path.
//
// The input of the package is the []rordefs.ApiResource provided by the "github.com/NorskHelsenett/ror/pkg/rorresources/rordefs" package rordefs in the variable Resources
//
// If new structs are added to the resources, add new structs in the pkg/apicontracts/apiresourcecontracts/* files
//
// TODO: Provide docslink
package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"

	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rordefs"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rorgenerator"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

var generator *rorgenerator.Generator

var v1TemplateFiles = []string{
	"pkg/apicontracts/apiresourcecontracts/resource_models_generated.go.tmpl",
	"pkg/apicontracts/apiresourcecontracts/resource_models_methods_generated.go.tmpl",
}

var v2TemplateFiles = []string{
	"pkg/rorresources/fromstruct.go.tmpl",
	"pkg/rorresources/resource.go.tmpl",
	"pkg/rorresources/rorkubernetes/k8s_test.go.tmpl",
	"pkg/rorresources/rorkubernetes/k8s.go.tmpl",
	"pkg/rorresources/rorkubernetes/k8s_new.go.tmpl",
	"pkg/rorresources/rortypes/resource_types_generated.go.tmpl",
}

func main() {

	generator = rorgenerator.NewGenerator()

	// Resource models v1
	for _, templateFile := range v1TemplateFiles {
		generator.TemplateFile(templateFile, rordefs.Resourcedefs.GetResourcesByVersion(rordefs.ApiVersionV1))
	}

	// Resource models v2
	for _, templateFile := range v2TemplateFiles {
		generator.TemplateFile(templateFile, rordefs.Resourcedefs.GetResourcesByVersion(rordefs.ApiVersionV2))
	}

	generatePatchMirrors()
	generator.TemplateFile("pkg/rorresources/resourcepatch.go.tmpl", rordefs.Resourcedefs.GetPatchableResources())

	generateTypescriptModels()
}

// generatePatchMirrors generates the patch mirrors of the patchable resource(s) definitions.
// see pkg/rorresources/rorgenerator/patchmirror.go for more information.
func generatePatchMirrors() {
	var roots []rorgenerator.PatchMirrorRoot
	for _, resource := range rordefs.Resourcedefs.GetPatchableResources() {
		roots = append(roots, rorgenerator.PatchMirrorRoot{Kind: resource.Kind, Type: patchMirrorRoot(resource.Kind)})
	}

	err := generator.PatchMirrorFiles(roots, rorgenerator.PatchMirrorConfig{
		Package:          "rortypes",
		PackagePath:      "github.com/NorskHelsenett/ror/pkg/rorresources/rortypes",
		SourceDir:        "pkg/rorresources/rortypes",
		Prefix:           "ResourcePatch",
		SourceFilePrefix: "resourcedef_",
		FilePrefix:       "resourcepatchdef_",
		OptionalPackage:  "github.com/NorskHelsenett/ror/pkg/helpers/optional",
		NilTargetError:   "ErrPatchNilTarget",
		GeneratedBy:      "cmd/generator/main.go",
		// Quantity wraps a kubernetes quantity and marshals itself as a
		// string, so it is replaced as a whole rather than taken apart.
		Opaque: map[string]bool{"Quantity": true},
	})
	if err != nil {
		panic(err.Error())
	}
}

// patchMirrorRoot finds the reflect.Type of the resource patch mirror of the given kind.
// It panics if the kind is marked patchable but rorresources.Resource has no field for it, which means the generator needs to be run again to add it.
func patchMirrorRoot(kind string) reflect.Type {
	field, ok := reflect.TypeOf(rorresources.Resource{}).FieldByName(kind + "Resource")
	if !ok {
		panic(fmt.Sprintf("kind %q is marked patchable but rorresources.Resource has no %sResource field, run the generator again to add it", kind, kind))
	}
	return field.Type.Elem()
}

func generateTypescriptModels() {
	workingDirectory, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	resourceV2TypescriptFilePath := fmt.Sprintf("%s/typescript/models/src/resources.ts", workingDirectory)
	if _, err = os.Stat(resourceV2TypescriptFilePath); errors.Is(err, os.ErrNotExist) {
		err = generator.TouchFile(resourceV2TypescriptFilePath)
		if err != nil {
			panic(err.Error())
		}
	}

	converter := typescriptify.New()
	converter.CreateInterface = true
	converter.CreateConstructor = true
	converter.AddEnum(rortypes.AllVulnerabilityStatuses)
	converter.AddEnum(rortypes.AllVulnerabilityDismissalReasons)
	converter.AddEnum(rortypes.AllResourceTagProperties)

	converter.Add(rorresources.ResourceSet{})
	converter.Add(rorresources.ResourceQuery{})

	err = converter.ConvertToFile(resourceV2TypescriptFilePath)
	if err != nil {
		panic(err.Error())
	}

	formatTypescript()
}

func formatTypescript() {
	workingDirectory, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	resourceV2TypescriptFilePath := fmt.Sprintf("%s/typescript/models", workingDirectory)

	getNodeDependenciesCmd := exec.Command("npm", "install")
	getNodeDependenciesCmd.Dir = resourceV2TypescriptFilePath
	npminstalloutput, err := getNodeDependenciesCmd.CombinedOutput()
	if err != nil {
		_, _ = fmt.Println("npm install failed on path", resourceV2TypescriptFilePath, "with err:", err.Error())
		fmt.Println("output", string(npminstalloutput))
		fmt.Println(err)
	}

	formatCmd := exec.Command("npm", "run", "format")
	formatCmd.Dir = resourceV2TypescriptFilePath
	npmformatoutput, err := formatCmd.CombinedOutput()
	if err != nil {
		_, _ = fmt.Println("prettier failed on path", resourceV2TypescriptFilePath, "with err:", err.Error())
		fmt.Println("output", string(npmformatoutput))
		fmt.Println(err)
	}
}
