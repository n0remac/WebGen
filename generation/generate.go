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

func main() {
	appPath := "../MyApp"
    projectName := "MyApp"

	copyDirectory("../app", appPath)
	replaceWordInDirectory(appPath, "CodeGen", projectName)
	replaceWordInDirectory(appPath, "codegen", strings.ToLower(projectName))
	
	err := generator.GenerateAll("schema.yaml", appPath, projectName)
	if err != nil {
		log.Fatalf("Generation failed: %v", err)
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
