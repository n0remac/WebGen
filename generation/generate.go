package main

import (
	"log"
	"generation/internal/generator"
)

func main() {
	err := generator.GenerateAll("schema.yaml")
	if err != nil {
		log.Fatalf("Generation failed: %v", err)
	}
}
