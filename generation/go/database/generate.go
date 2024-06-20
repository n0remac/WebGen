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

func minus1(n int) int {
	return n - 1
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

	// Generate table files
	tmplPath := "table.go.tpl"
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
		panic(err)
	}

	for _, model := range schema.Models {
		outputFilePath := filepath.Join("../../../pkg/database", strings.ToLower(model.Name)+".go")
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
			"ModelName": model.Name,
			"Fields":    model.Fields,
		}

		err = tmpl.Execute(file, data)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Generated %s\n", outputFilePath)
	}

	// Generate setup.go file
	setupTmplPath := "setup.go.tpl"
	setupTmpl, err := template.ParseFiles(setupTmplPath)
	if err != nil {
		panic(err)
	}

	setupFilePath := filepath.Join("../../../pkg/database", "setup.go")
	setupFile, err := os.Create(setupFilePath)
	if err != nil {
		panic(err)
	}
	defer setupFile.Close()

	err = setupTmpl.Execute(setupFile, schema)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Generated %s\n", setupFilePath)
}
