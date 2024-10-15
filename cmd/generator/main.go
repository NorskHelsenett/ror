// The Generator package provides a way to generate code for collecting,
// transfering and saving resources in the agent and api.
// It also provides functions to fetc the resources.
//
//	go run build/generator/main.go
//
// The package updates the files:
//   - cmd/agent/clients/dynamicclient_generated.go
//   - cmd/agentv2/clients/dynamicclient/dynamicclient_generated.go
//   - cmd/api/controllers/resourcescontroller/resources_controller_read_generated.go
//   - cmd/api/services/resourcesService/resourceServices_generated.go
//   - cmd/tanzu/agent/tanzuservice/schemas/schemas_generated.go
//   - internal/models/rorResources/extractResource.go
//   - internal/mongodbrepo/repositories/resourcesmongodbrepo/resourcesinsertupdate_generated.go
//
// If the ror package is available, the package updates the files in the ror package:
//   - pkg/apicontracts/apiresourcecontracts/resource_models_generated.go
//   - pkg/apicontracts/apiresourcecontracts/resource_models_methods_generated.go
//   - pkg/rorresources/fromstruct.go
//   - pkg/rorresources/resource.go
//   - pkg/rorresources/rorkubernetes/k8s_test.go
//   - pkg/rorresources/rorkubernetes/k8s.go
//   - pkg/rorresources/rortypes/resource_interfaces.go
//   - pkg/rorresources/rortypes/resource_models_methods.go
//   - pkg/rorresources/rortypes/resource_input_filter_*.go
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
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rordefs"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {

	// Dynamic clients - agents
	//   - cmd/agent/clients/dynamicclient_generated.go
	templateFile("cmd/agent/clients/dynamicclient_generated.go.tmpl", rordefs.GetResourcesByType(rordefs.ApiResourceTypeAgent))
	//  - cmd/agentv2/clients/dynamicclient/dynamicclient_generated.go
	templateFile("cmd/agentv2/clients/dynamicclient/dynamicclient_generated.go.tmpl", rordefs.GetResourcesByType(rordefs.ApiResourceTypeAgent))

	// Resource controller - api
	//   - cmd/api/controllers/resourcescontroller/resources_controller_read_generated.go
	templateFile("cmd/api/controllers/resourcescontroller/resources_controller_read_generated.go.tmpl", rordefs.Resourcedefs)

	// Resource services - api
	//   - cmd/api/services/resourcesService/resourceServices_generated.go
	templateFile("cmd/api/services/resourcesService/resourceServices_generated.go.tmpl", rordefs.Resourcedefs)

	// Schema - Tanzu service
	//   - cmd/tanzu/agent/tanzuservice/schemas/schemas_generated.go
	templateFile("cmd/tanzu/agent/tanzuservice/schemas/schemas_generated.go.tmpl", rordefs.GetResourcesByType(rordefs.ApiResourceTypeTanzuAgent))

	// Internal - models
	//   - internal/models/rorResources/extractResource.go
	templateFile("internal/models/rorResources/extractResource.go.tmpl", rordefs.Resourcedefs)

	// Resource - mongorepo
	//   - internal/mongodbrepo/repositories/resourcesmongodbrepo/resourcesinsertupdate_generated.go
	templateFile("internal/mongodbrepo/repositories/resourcesmongodbrepo/resourcesinsertupdate_generated.go.tmpl", rordefs.Resourcedefs)

	// Resource models . Moved to github.com/NorskHelsenett/ROR, should be placed in the ../ror folder
	var folderPath string
	if os.Getenv("ROR_PATH") == "" {
		folderPath = "../ror"
	} else {
		folderPath = os.Getenv("ROR_PATH")
	}

	fmt.Println("ROR_PATH: ", folderPath)

	exists := folderExists(folderPath)
	if exists {
		templateFile(fmt.Sprintf("%s/pkg/apicontracts/apiresourcecontracts/resource_models_generated.go.tmpl", folderPath), rordefs.Resourcedefs)
		templateFile(fmt.Sprintf("%s/pkg/apicontracts/apiresourcecontracts/resource_models_methods_generated.go.tmpl", folderPath), rordefs.Resourcedefs)

		// Resource models
		templateFile(fmt.Sprintf("%s/pkg/rorresources/fromstruct.go.tmpl", folderPath), rordefs.Resourcedefs)
		templateFile(fmt.Sprintf("%s/pkg/rorresources/resource.go.tmpl", folderPath), rordefs.Resourcedefs)
		templateFile(fmt.Sprintf("%s/pkg/rorresources/rorkubernetes/k8s_test.go.tmpl", folderPath), rordefs.Resourcedefs)
		templateFile(fmt.Sprintf("%s/pkg/rorresources/rorkubernetes/k8s.go.tmpl", folderPath), rordefs.Resourcedefs)
		templateFile(fmt.Sprintf("%s/pkg/rorresources/rortypes/resource_interfaces.go.tmpl", folderPath), rordefs.Resourcedefs)
		templateFile(fmt.Sprintf("%s/pkg/rorresources/rortypes/resource_models_methods.go.tmpl", folderPath), rordefs.Resourcedefs)

		// Resource models - input filters
		for _, res := range rordefs.Resourcedefs {
			filepath := fmt.Sprintf("%s/pkg/rorresources/rortypes/resource_input_filter_%s.go", folderPath, res.Kind)
			filepath = strings.ToLower(filepath)
			templateFileOnce(filepath, fmt.Sprintf("%s/pkg/rorresources/rortypes/resource_models_input_filter.go.tmpl", folderPath), res)
		}

	} else {
		fmt.Println("Folder does not exist: ", folderPath)
	}

	// golang struct to typescript interfaces
	workingDirectory, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	resourceV2TypescriptFilePath := fmt.Sprintf("%s/web/admin/src/app/core/models/resources-v2.ts", workingDirectory)
	if _, err = os.Stat(resourceV2TypescriptFilePath); errors.Is(err, os.ErrNotExist) {
		err = touchFile(resourceV2TypescriptFilePath)
		if err != nil {
			panic(err.Error())
		}
	}

	converter := typescriptify.New()
	converter.CreateInterface = true
	converter.CreateConstructor = true
	converter.BackupDir = os.TempDir()
	converter.Add(rorresources.ResourceSet{})
	converter.Add(rorresources.ResourceQuery{})
	err = converter.ConvertToFile(resourceV2TypescriptFilePath)
	if err != nil {
		panic(err.Error())
	}

	formatTypescript()
}

func templateFileOnce(filepath string, templatePath string, data any) {

	if fileExists(filepath) {
		fmt.Println("File exists: ", filepath)
		return
	}
	templateToFile(filepath, templatePath, data)
}

func folderExists(folderPath string) bool {
	fileInfo, err := os.Stat(folderPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		// Handle other errors if needed
	}
	return fileInfo.IsDir()
}

func fileExists(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		// Handle other errors if needed
	}
	return !fileInfo.IsDir()
}

func templateFile(filepath string, data any) {

	outfile := strings.TrimSuffix(filepath, path.Ext(filepath))
	templateToFile(outfile, filepath, data)
}

func templateToFile(filepath string, templatePath string, data any) {

	t, err := os.ReadFile(templatePath) // #nosec G304 - This is a generator and the files are under control

	if err != nil {
		log.Print(err)
		return
	}
	funcMap := template.FuncMap{
		"lower": strings.ToLower,
	}
	tmpl, err := template.New("Template").Funcs(funcMap).Parse(string(t))
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filepath) // #nosec G304 - This is a generator and the files are under control

	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	err = tmpl.Execute(f, data)
	if err != nil {
		fmt.Println(err)
	}

	fmtcmd := exec.Command("go", "fmt", filepath)
	_, err = fmtcmd.Output()
	if err != nil {
		_, _ = fmt.Println("go formater failed with err: ", err.Error())
		fmt.Println(err)
	}
	fmt.Println("Generated file: ", filepath)
}

func touchFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0600) // #nosec G304 - This is a generator and the files are under control
	if err != nil {
		return err
	}
	return file.Close()
}

func formatTypescript() {
	workingDirectory, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	resourceV2TypescriptFilePath := fmt.Sprintf("%s/web/admin", workingDirectory)

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
