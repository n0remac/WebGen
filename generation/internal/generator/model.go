package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func title(s string) string {
	return cases.Title(language.English).String(s)
}

func GenerateModels(schema *Schema, appPath string, projectName string) error {
	tmplPath := "templates/go/model/model.go.tpl"
	tmpl, err := template.New("model.go.tpl").Funcs(template.FuncMap{
		"sqlType": func(fieldType string) string {
			switch fieldType {
			case "int32":
				return "int"
			case "string":
				return "string"
			default:
				return fieldType
			}
		},
		"lower": strings.ToLower,
		"title": title,
	}).ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	for _, model := range schema.Models {
		outputFilePath := filepath.Join(appPath, "pkg/model", strings.ToLower(model.Name)+".go")
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
			"ModelName":   model.Name,
			"Fields":      model.Fields,
			"ProjectName": projectName,
		}

		if err := tmpl.Execute(file, data); err != nil {
			return err
		}

		fmt.Printf("Generated %s\n", outputFilePath)
	}

	return nil
}
