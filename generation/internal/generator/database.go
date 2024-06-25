package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)


func minus1(n int) int {
	return n - 1
}

func GenerateDatabase(schema *Schema, appPath string) error {
	// Generate table files
	if err := generateTableFiles(schema, appPath); err != nil {
		return fmt.Errorf("generating table files: %w", err)
	}

	// Generate setup.go file
	if err := generateSetupFile(schema, appPath); err != nil {
		return fmt.Errorf("generating setup file: %w", err)
	}

	return nil
}

func generateTableFiles(schema *Schema, appPath string) error {
	tmplPath := "templates/go/database/table.go.tpl"
	tmpl, err := template.New("table.go.tpl").Funcs(template.FuncMap{
		"sqlType": func(fieldType string) string {
			switch fieldType {
			case "int32":
				return "INTEGER"
			case "string":
				return "TEXT"
			default:
				return fieldType
			}
		},
		"lower":  strings.ToLower,
		"title":  title,
		"minus1": minus1,
	}).ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	for _, model := range schema.Models {
		outputFilePath := filepath.Join(appPath, "pkg/database", strings.ToLower(model.Name)+".go")
		outputDir := filepath.Dir(outputFilePath)
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return err
		}

		file, err := os.Create(outputFilePath)
		if err != nil {
			return err
		}
		defer file.Close()

		data := map[string]interface{}{
			"ModelName": model.Name,
			"Fields":    model.Fields,
		}

		err = tmpl.Execute(file, data)
		if err != nil {
			return err
		}

		fmt.Printf("Generated %s\n", outputFilePath)
	}

	return nil
}

func generateSetupFile(schema *Schema, appPath string) error {
	setupTmplPath := "templates/go/database/setup.go.tpl"
	setupTmpl, err := template.ParseFiles(setupTmplPath)
	if err != nil {
		return err
	}

	setupFilePath := filepath.Join(appPath, "pkg/database", "setup.go")
	setupFile, err := os.Create(setupFilePath)
	if err != nil {
		return err
	}
	defer setupFile.Close()

	err = setupTmpl.Execute(setupFile, schema)
	if err != nil {
		return err
	}

	fmt.Printf("Generated %s\n", setupFilePath)
	return nil
}
