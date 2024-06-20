package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

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
	if (err != nil) {
		panic(err)
	}

	// Generate service files
	serviceTmplPath := "service.go.tpl"
	serviceTmpl, err := template.New("service.go.tpl").Funcs(template.FuncMap{
		"lower": strings.ToLower,
	}).ParseFiles(serviceTmplPath)
	if err != nil {
		panic(err)
	}

	for _, model := range schema.Models {
		outputFilePath := filepath.Join("../../../pkg/service", strings.ToLower(model.Name)+".go")
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
			"ServiceName": model.Name,
			"ServiceType": model.Name + "Service",
			"ProjectName": "Go-gRPC-React-starter",
			"Fields":      model.Fields,
		}

		err = serviceTmpl.Execute(file, data)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Generated %s\n", outputFilePath)
	}

	// Generate setup.go file for services
	setupServiceTmplPath := "setup.go.tpl"
	setupServiceTmpl, err := template.New("setup.go.tpl").Funcs(template.FuncMap{
		"lower": strings.ToLower,
	}).ParseFiles(setupServiceTmplPath)
	if err != nil {
		panic(err)
	}

	setupServiceFilePath := filepath.Join("../../../pkg/service", "setup.go")
	setupServiceFile, err := os.Create(setupServiceFilePath)
	if err != nil {
		panic(err)
	}
	defer setupServiceFile.Close()

	type ServiceData struct {
		ServiceName string
		ServiceType string
		ProjectName string
	}

	var serviceData []ServiceData
	for _, model := range schema.Models {
		serviceData = append(serviceData, ServiceData{
			ServiceName: model.Name,
			ServiceType: model.Name + "Service",
			ProjectName: "Go-gRPC-React-starter",
		})
	}

	err = setupServiceTmpl.Execute(setupServiceFile, map[string]interface{}{
		"Services":    serviceData,
		"ProjectName": "Go-gRPC-React-starter",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Generated %s\n", setupServiceFilePath)
}
