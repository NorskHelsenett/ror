package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := "docs/code"
	nav := generateNav(dir)
	fmt.Println(nav)
}

func generateNav(dir string) string {
	var navBuilder strings.Builder
	navBuilder.WriteString(fmt.Sprintf("  - %s:\n", "Code"))
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(dir, path)
		indentLevel := strings.Count(relPath, string(os.PathSeparator))

		if info.IsDir() && relPath != "." {
			navBuilder.WriteString(strings.Repeat("  ", indentLevel+2))
			navBuilder.WriteString(fmt.Sprintf("- %s:\n", info.Name()))
		} else if !info.IsDir() && filepath.Ext(path) == ".md" {
			navBuilder.WriteString(strings.Repeat("  ", indentLevel+1))
			navBuilder.WriteString(fmt.Sprintf("  - code/%s\n", relPath))
		}

		return nil
	})

	return navBuilder.String()
}
