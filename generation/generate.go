package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Define the directories containing the generate.go files
	directories := []string{
		"go/database",
		"go/model",
		"go/service",
		"proto",
		"react",
		"react/services",
	}

	// Iterate over the directories and run each generate.go file
	for _, dir := range directories {
		cmd := exec.Command("go", "run", "generate.go")
		cmd.Dir = dir // Set the working directory
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		log.Printf("Running generate.go in %s\n", dir)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error running generate.go in %s: %v", dir, err)
		}
	}
}
