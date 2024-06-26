package generator

import (
    "fmt"
    "io"
    "os"

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

func LoadSchema(path string) (*Schema, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    yamlData, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }

    var schema Schema
    err = yaml.Unmarshal(yamlData, &schema)
    if err != nil {
        return nil, err
    }

    return &schema, nil
}

func GenerateAll(schemaPath string, appPath string, projectName string, vars TerraformVars) error {
    schema, err := LoadSchema(schemaPath)
    if err != nil {
        return fmt.Errorf("loading schema: %w", err)
    }

    if err := GenerateModels(schema, appPath, projectName); err != nil {
        return fmt.Errorf("generating models: %w", err)
    }

    if err := GenerateDatabase(schema, appPath); err != nil {
        return fmt.Errorf("generating database: %w", err)
    }

    if err := GenerateServices(schema, appPath, projectName); err != nil {
        return fmt.Errorf("generating services: %w", err)
    }

    if err := GenerateProto(schema, appPath); err != nil {
        return fmt.Errorf("generating proto: %w", err)
    }

    if err := GenerateReact(schema, appPath); err != nil {
        return fmt.Errorf("generating react: %w", err)
    }

    if err := GenerateTerraform(vars, appPath); err != nil {
        return fmt.Errorf("generating terraform: %w", err)
    }

    return nil
}
