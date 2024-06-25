package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GenerateServices(schema *Schema, appPath string, projectName string) error {
	// Generate service files
	if err := generateServiceFiles(schema, appPath, projectName); err != nil {
		return fmt.Errorf("generating service files: %w", err)
	}

	// Generate setup.go file for services
	if err := generateServiceSetupFile(schema, appPath, projectName); err != nil {
		return fmt.Errorf("generating service setup file: %w", err)
	}

	return nil
}

func generateServiceFiles(schema *Schema, appPath string, projectName string) error {
	serviceTmplPath := "templates/go/service/service.go.tpl"
	serviceTmpl, err := template.New("service.go.tpl").Funcs(template.FuncMap{
		"lower": strings.ToLower,
	}).ParseFiles(serviceTmplPath)
	if err != nil {
		return err
	}

	for _, model := range schema.Models {
		outputFilePath := filepath.Join(appPath, "pkg/service", strings.ToLower(model.Name)+".go")
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
			"ServiceType": model.Name + "Service",
			"ProjectName": projectName,
			"Fields":      model.Fields,
		}

		err = serviceTmpl.Execute(file, data)
		if err != nil {
			return err
		}

		fmt.Printf("Generated %s\n", outputFilePath)
	}

	return nil
}

func generateServiceSetupFile(schema *Schema, appPath string, projectName string) error {
	setupServiceTmplPath := "templates/go/service/setup.go.tpl"
	setupServiceTmpl, err := template.New("setup.go.tpl").Funcs(template.FuncMap{
		"lower": strings.ToLower,
	}).ParseFiles(setupServiceTmplPath)
	if err != nil {
		return err
	}

	setupServiceFilePath := filepath.Join(appPath, "pkg/service", "setup.go")
	setupServiceFile, err := os.Create(setupServiceFilePath)
	if err != nil {
		return err
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
			ProjectName: projectName,
		})
	}

	err = setupServiceTmpl.Execute(setupServiceFile, map[string]interface{}{
		"Services":    serviceData,
		"ProjectName": projectName,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Generated %s\n", setupServiceFilePath)
	return nil
}
