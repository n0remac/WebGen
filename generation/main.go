package main

import (
	"log"
	"generation/internal/generator"

	"fmt"
	"os/exec"
	"os"
	"path/filepath"
	"strings"
)

var projectName string = "MyApp" 

func main() {
	generateApp()
	setupGit()
}

func generateApp() {
	appPath := "../MyApp"

	copyDirectory("../app", appPath)
	replaceWordInDirectory(appPath, "CodeGen", projectName)
	replaceWordInDirectory(appPath, "codegen", strings.ToLower(projectName))
	
	err := generator.GenerateAll("schema.yaml", appPath, projectName)
	if err != nil {
		log.Fatalf("Generation failed: %v", err)
	}
}

func setupGit() {
	// Define variables here
	githubOwner := "n0remac"

	// Create the Terraform variables
	vars := generator.TerraformVars{
		GithubOwner: githubOwner,
		RepoName:    projectName,
	}

	// Generate Terraform files
	err := generator.GenerateTerraform(vars, "..")
	if err != nil {
		log.Fatalf("Error generating Terraform files: %v", err)
	}

	// Run Terraform
	runTerraform("../terraform")

	// Run Git commands to push code
	runGitCommands(githubOwner, projectName)
}

func runTerraform(appPath string) {
	cmds := [][]string{
		{"terraform", "init"},
		{"terraform", "apply", "-auto-approve"},
	}

	for _, cmd := range cmds {
		runCommand(appPath, cmd[0], cmd[1:]...)
	}
}

func runGitCommands(owner, repo string) {
	commands := [][]string{
		{"sh", "-c", `echo "# ` + repo + `" >> README.md`},
		{"git", "init"},
		{"git", "add", "."},
		{"git", "commit", "-m", "first commit"},
		{"git", "branch", "-M", "main"},
		{"git", "remote", "add", "origin", "git@github.com:" + owner + "/" + repo + ".git"},
		{"git", "push", "-u", "origin", "main"},
	}

	for _, cmd := range commands {
		runCommand("../"+projectName, cmd[0], cmd[1:]...)
	}
}

func runCommand(dir, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command %s %v failed with %v", name, args, err)
	}
}

func replaceWordInDirectory(directoryPath, oldWord, newWord string) error {
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			read, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			content := string(read)
			newContent := strings.ReplaceAll(content, oldWord, newWord)
			if content != newContent {
				err = os.WriteFile(path, []byte(newContent), info.Mode())
				if err != nil {
					return err
				}
				fmt.Printf("Replaced in file: %s\n", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	return nil
}

func copyDirectory(source, destination string) error {
	// Construct the cp command
	cmd := exec.Command("cp", "-r", source, destination)

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	// Print the output of the command
	fmt.Printf("Output: %s\n", string(output))

	return nil
}
