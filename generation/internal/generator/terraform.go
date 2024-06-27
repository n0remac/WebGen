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

type DigitalOceanVars struct {
	DigitalOceanToken string
	ProjectName       string
	DropletName       string
	Region            string
	Size              string
	Image             string
}

func GenerateTerraform(vars TerraformVars, appPath string) error {
	// Create the Terraform directory
	terraformPath := filepath.Join(appPath, "terraform")
	err := os.Mkdir(terraformPath, 0755)
	if err != nil {
		return fmt.Errorf("creating git terraform directory: %w", err)
	}

	if err := generateMainTF(vars, appPath); err != nil {
		return fmt.Errorf("generating git main.tf: %w", err)
	}

	if err := generateTerraformTFVars(vars, appPath); err != nil {
		return fmt.Errorf("generating git terraform.tfvars: %w", err)
	}

	if err := generateVariablesTF(vars, appPath); err != nil {
		return fmt.Errorf("generating git variables.tf: %w", err)
	}

	return nil
}

func GenerateDigitalOceanTerraform(vars DigitalOceanVars, appPath string) error {
	// Create the Terraform directory
	terraformPath := filepath.Join(appPath, "terraform")
	err := os.Mkdir(terraformPath, 0755)
	if err != nil {
		return fmt.Errorf("creating digitalocean terraform directory: %w", err)
	}

	if err := generateDigitalOceanMainTF(vars, appPath); err != nil {
		return fmt.Errorf("generating digitalocean main.tf: %w", err)
	}

	if err := generateDigitalOceanTFVars(vars, appPath); err != nil {
		return fmt.Errorf("generating digitalocean terraform.tfvars: %w", err)
	}

	if err := generateDigitalOceanVariablesTF(vars, appPath); err != nil {
		return fmt.Errorf("generating digitalocean variables.tf: %w", err)
	}

	return nil
}

// GitHub Terraform generation functions

func generateMainTF(vars TerraformVars, appPath string) error {
	tmplPath := "templates/git-terraform/main.tf.tpl"
	return executeTemplate(tmplPath, vars, filepath.Join(appPath, "main.tf"))
}

func generateTerraformTFVars(vars TerraformVars, appPath string) error {
	tmplPath := "templates/git-terraform/terraform.tfvars.tpl"
	return executeTemplate(tmplPath, vars, filepath.Join(appPath, "terraform.tfvars"))
}

func generateVariablesTF(vars TerraformVars, appPath string) error {
	tmplPath := "templates/git-terraform/variables.tf.tpl"
	return executeTemplate(tmplPath, vars, filepath.Join(appPath, "variables.tf"))
}

// DigitalOcean Terraform generation functions

func generateDigitalOceanMainTF(vars DigitalOceanVars, appPath string) error {
	tmplPath := "templates/digitalocean-terraform/main.tf.tpl"
	return executeTemplate(tmplPath, vars, filepath.Join(appPath, "main.tf"))
}

func generateDigitalOceanTFVars(vars DigitalOceanVars, appPath string) error {
	tmplPath := "templates/digitalocean-terraform/terraform.tfvars.tpl"
	return executeTemplate(tmplPath, vars, filepath.Join(appPath, "terraform.tfvars"))
}

func generateDigitalOceanVariablesTF(vars DigitalOceanVars, appPath string) error {
	tmplPath := "templates/digitalocean-terraform/variables.tf.tpl"
	return executeTemplate(tmplPath, vars, filepath.Join(appPath, "variables.tf"))
}

func executeTemplate(tmplPath string, vars interface{}, outputPath string) error {
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
