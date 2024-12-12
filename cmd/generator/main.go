// The Generator package provides a way to generate code for collecting,
// transfering and saving resources in the agent and api.
// It also provides functions to fetc the resources.
//
<<<<<<< HEAD
<<<<<<< HEAD
//		go run build/generator/main.go
=======
//	go run cmd/generator/main.go
>>>>>>> 550cbdd (Added resourceBackupJob types and generator stuff)
=======
//	go run build/generator/main.go
>>>>>>> aaf2f15 (Update main.go)
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
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	"github.com/NorskHelsenett/ror/pkg/rorresources/rordefs"
)

func main() {

	// Schema - Tanzu service
	//   - cmd/tanzu/agent/tanzuservice/schemas/schemas_generated.go
	//templateFile("cmd/tanzu/agent/tanzuservice/schemas/schemas_generated.go.tmpl", rordefs.GetResourcesByType(rordefs.ApiResourceTypeTanzuAgent))

	templateFile("pkg/apicontracts/apiresourcecontracts/resource_models_generated.go.tmpl", rordefs.Resourcedefs)
	templateFile("pkg/apicontracts/apiresourcecontracts/resource_models_methods_generated.go.tmpl", rordefs.Resourcedefs)

	// Resource models
	templateFile("pkg/rorresources/fromstruct.go.tmpl", rordefs.Resourcedefs)
	templateFile("pkg/rorresources/resource.go.tmpl", rordefs.Resourcedefs)
	templateFile("pkg/rorresources/rorkubernetes/k8s_test.go.tmpl", rordefs.Resourcedefs)
	templateFile("pkg/rorresources/rorkubernetes/k8s.go.tmpl", rordefs.Resourcedefs)
	templateFile("pkg/rorresources/rortypes/resource_interfaces.go.tmpl", rordefs.Resourcedefs)
	templateFile("pkg/rorresources/rortypes/resource_models_methods.go.tmpl", rordefs.Resourcedefs)

	// Resource models - input filters
	for _, res := range rordefs.Resourcedefs {
		filepath := fmt.Sprintf("pkg/rorresources/rortypes/resource_input_filter_%s.go", res.GetKind())
		filepath = strings.ToLower(filepath)
		templateFileOnce(filepath, "/pkg/rorresources/rortypes/resource_models_input_filter.go.tmpl", res)
	}

}

func templateFileOnce(filepath string, templatePath string, data any) {

	if fileExists(filepath) {
		fmt.Println("File exists: ", filepath)
		return
	}
	templateToFile(filepath, templatePath, data)
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
