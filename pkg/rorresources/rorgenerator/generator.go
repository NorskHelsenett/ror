package rorgenerator

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}
}

// templateFileOnce will only generate the file if it does not exist
// if the file exists it will print a message to the console
func (g Generator) TemplateFileOnce(filepath string, templatePath string, data any) {
	if g.FileExists(filepath) {
		fmt.Println("File exists: ", filepath)
		return
	}
	g.TemplateToFile(filepath, templatePath, data)
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func (g Generator) FileExists(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		// Handle other errors if needed
	}
	return !fileInfo.IsDir()
}

// templateFile will generate the file regardless if it exists or not
func (g Generator) TemplateFile(filepath string, data any) {
	outfile := strings.TrimSuffix(filepath, path.Ext(filepath))
	g.TemplateToFile(outfile, filepath, data)
}

// templateToFile will generate a file from a template to a specific path
func (g Generator) TemplateToFile(filepath string, templatePath string, data any) {

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

// touchFile will create a file if it does not exist
func (g Generator) TouchFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	return file.Close()
}
