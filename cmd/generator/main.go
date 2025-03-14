// The Generator package provides a way to generate code for collecting,
// transfering and saving resources in the agent and api.
// It also provides functions to fetc the resources.
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
	"strings"

	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rordefs"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rorgenerator"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

var generator *rorgenerator.Generator

func main() {

	generator = rorgenerator.NewGenerator()

	// Resource models v1
	generator.TemplateFile("pkg/apicontracts/apiresourcecontracts/resource_models_generated.go.tmpl", rordefs.Resourcedefs.GetResourcesByVersion(rordefs.ApiVersionV1))
	generator.TemplateFile("pkg/apicontracts/apiresourcecontracts/resource_models_methods_generated.go.tmpl", rordefs.Resourcedefs.GetResourcesByVersion(rordefs.ApiVersionV1))

	// Resource models v2
	generator.TemplateFile("pkg/rorresources/fromstruct.go.tmpl", rordefs.Resourcedefs.GetResourcesByVersion(rordefs.ApiVersionV2))
	generator.TemplateFile("pkg/rorresources/resource.go.tmpl", rordefs.Resourcedefs.GetResourcesByVersion(rordefs.ApiVersionV2))
	generator.TemplateFile("pkg/rorresources/rorkubernetes/k8s_test.go.tmpl", rordefs.Resourcedefs.GetResourcesByVersion(rordefs.ApiVersionV2))
	generator.TemplateFile("pkg/rorresources/rorkubernetes/k8s.go.tmpl", rordefs.Resourcedefs.GetResourcesByVersion(rordefs.ApiVersionV2))
	generator.TemplateFile("pkg/rorresources/rortypes/resource_interfaces.go.tmpl", rordefs.Resourcedefs.GetResourcesByVersion(rordefs.ApiVersionV2))
	generator.TemplateFile("pkg/rorresources/rortypes/resource_models_methods.go.tmpl", rordefs.Resourcedefs.GetResourcesByVersion(rordefs.ApiVersionV2))

	// Resource models - input filters
	for _, res := range rordefs.Resourcedefs.GetResourcesByVersion(rordefs.ApiVersionV2) {
		filepath := fmt.Sprintf("pkg/rorresources/rortypes/resource_input_filter_%s.go", res.GetKind())
		filepath = strings.ToLower(filepath)
		generator.TemplateFileOnce(filepath, "pkg/rorresources/rortypes/resource_models_input_filter.go.tmpl", res)
	}

	generateTypescriptModels()
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
	_, err = getNodeDependenciesCmd.CombinedOutput()
	if err != nil {
		_, _ = fmt.Println("npm install failed with err: ", err.Error())
		fmt.Println(err)
	}

	formatCmd := exec.Command("npm", "run", "format")
	formatCmd.Dir = resourceV2TypescriptFilePath
	_, err = formatCmd.CombinedOutput()
	if err != nil {
		_, _ = fmt.Println("prettier failed with err: ", err.Error())
		fmt.Println(err)
	}
}
