package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GenerateProto(schema *Schema, appPath string) error {
	tmplPath := "templates/proto/crud.proto.tpl"
	tmpl, err := template.New("crud.proto.tpl").Funcs(template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"lower": strings.ToLower,
	}).ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	for _, model := range schema.Models {
		outputFilePath := filepath.Join(appPath, "proto", strings.ToLower(model.Name), strings.ToLower(model.Name)+".proto")
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
			"PackageName": strings.ToLower(model.Name),
			"ModelName":   model.Name,
			"Fields":      model.Fields,
		}

		err = tmpl.Execute(file, data)
		if err != nil {
			return err
		}

		fmt.Printf("Generated %s\n", outputFilePath)
	}

	return nil
}
