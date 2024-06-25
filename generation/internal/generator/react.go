package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GenerateReact(schema *Schema, appPath string) error {
	if err := generateReactServices(schema, appPath); err != nil {
		return fmt.Errorf("generating react services: %w", err)
	}

	if err := generateReactAdminPages(schema, appPath); err != nil {
		return fmt.Errorf("generating react admin pages: %w", err)
	}

	return nil
}

func generateReactServices(schema *Schema, appPath string) error {
	tmplPath := "templates/react/services/service.ts.tpl"
	tmpl, err := template.New("service.ts.tpl").Funcs(template.FuncMap{
		"title": title,
		"lower": strings.ToLower,
	}).ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	for _, model := range schema.Models {
		outputFilePath := filepath.Join(appPath, "frontend/src/services", model.Name+"Service.ts")
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
			"ServiceName": model.Name,
		}

		err = tmpl.Execute(file, data)
		if err != nil {
			return err
		}

		fmt.Printf("Generated %s\n", outputFilePath)
	}

	return nil
}

func generateReactAdminPages(schema *Schema, appPath string) error {
	tmplPath := "templates/react/admin.tsx.tpl"
	tmpl, err := template.New("admin.tsx.tpl").Funcs(template.FuncMap{
		"lower":         strings.ToLower,
		"title":         title,
		"wrapInBraces":  wrapInBraces,
		"isIntegerType": isIntegerType,
	}).ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	for _, model := range schema.Models {
		outputFilePath := filepath.Join(appPath, "frontend/src/pages/admin", strings.ToLower(model.Name)+"AdminPage.tsx")
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

func wrapInBraces(s string) string {
	return "{" + s + "}"
}

func isIntegerType(fieldType string) bool {
	return fieldType == "int32" || fieldType == "int64" || fieldType == "uint32" || fieldType == "uint64"
}
