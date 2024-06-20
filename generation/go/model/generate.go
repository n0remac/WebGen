package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

type Field struct {
	Name          string `yaml:"name"`
	Type          string `yaml:"type"`
	PrimaryKey    bool   `yaml:"primary_key,omitempty"`
	AutoIncrement bool   `yaml:"auto_increment,omitempty"`
}

type Model struct {
	Name   string  `yaml:"name"`
	Fields []Field `yaml:"fields"`
}

type Schema struct {
	Models []Model `yaml:"models"`
}

func title(s string) string {
	return cases.Title(language.English).String(s)
}

func main() {
	file, err := os.Open("../../schema.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	yamlData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var schema Schema
	err = yaml.Unmarshal(yamlData, &schema)
	if err != nil {
		panic(err)
	}

	tmplPath := "model.go.tpl"
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
		panic(err)
	}

	for _, model := range schema.Models {
		outputFilePath := filepath.Join("../../../pkg/model", strings.ToLower(model.Name)+".go")
		outputDir := filepath.Dir(outputFilePath)
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			panic(err)
		}

		file, err := os.Create(outputFilePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		data := map[string]interface{}{
			"ModelName":   model.Name,
			"Fields":      model.Fields,
			"ProjectName": "Go-gRPC-React-starter", // Update this as per your project name
		}

		err = tmpl.Execute(file, data)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Generated %s\n", outputFilePath)
	}
}
