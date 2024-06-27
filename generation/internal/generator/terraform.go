package generator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type TerraformVars struct {
	GithubToken string
	GithubOwner string
	RepoName    string
}

func GenerateTerraform(vars TerraformVars, appPath string) error {
	// Create the Terraform directory
	terraformPath := filepath.Join(appPath, "terraform")
	err := os.Mkdir(terraformPath, 0755)
	if err != nil {
		return fmt.Errorf("creating terraform directory: %w", err)
	}

	if err := generateMainTF(vars, appPath); err != nil {
		return fmt.Errorf("generating main.tf: %w", err)
	}

	if err := generateTerraformTFVars(vars, appPath); err != nil {
		return fmt.Errorf("generating terraform.tfvars: %w", err)
	}

	if err := generateVariablesTF(vars, appPath); err != nil {
		return fmt.Errorf("generating variables.tf: %w", err)
	}

	return nil
}

func generateMainTF(vars TerraformVars, appPath string) error {
	tmplPath := "templates/terraform/main.tf.tpl"
	return executeTemplate(tmplPath, vars, filepath.Join(appPath, "terraform/main.tf"))
}

func generateTerraformTFVars(vars TerraformVars, appPath string) error {
	tmplPath := "templates/terraform/terraform.tfvars.tpl"
	return executeTemplate(tmplPath, vars, filepath.Join(appPath, "terraform/terraform.tfvars"))
}

func generateVariablesTF(vars TerraformVars, appPath string) error {
	tmplPath := "templates/terraform/variables.tf.tpl"
	return executeTemplate(tmplPath, vars, filepath.Join(appPath, "terraform/variables.tf"))
}

func executeTemplate(tmplPath string, vars TerraformVars, outputPath string) error {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, vars)
	if err != nil {
		return err
	}

	err = os.WriteFile(outputPath, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}
