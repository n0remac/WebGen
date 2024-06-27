package service

import (
	"SchemaGen/gen/proto/schema"
	"SchemaGen/pkg/model"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bufbuild/connect-go"
)

type SchemaService struct {
	// Add any fields if needed
}

func (s *SchemaService) CreateSchema(ctx context.Context, req *connect.Request[schema.CreateSchemaRequest]) (*connect.Response[schema.CreateSchemaResponse], error) {
	newSchema, err := model.CreateSchema(req.Msg.Schema)
	if err != nil {
		return nil, err
	}

	// Call function to save schema content to schema.yml file
	err = saveSchemaToFile(newSchema.Name, newSchema.Content)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&schema.CreateSchemaResponse{
		Schema: newSchema,
	}), nil
}

// Function to save schema content to a schema.yml file
func saveSchemaToFile(schemaName, schemaContent string) error {
	// Ensure the directory exists
	dir := "schemas"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// Create the file path
	filePath := filepath.Join(dir, fmt.Sprintf("%s.yml", schemaName))

	// Write the schema content to the file
	return os.WriteFile(filePath, []byte(schemaContent), 0644)
}


func (s *SchemaService) GetSchema(ctx context.Context, req *connect.Request[schema.GetSchemaRequest]) (*connect.Response[schema.GetSchemaResponse], error) {
	u, err := model.GetSchemaFromDB(req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&schema.GetSchemaResponse{
		Schema: u,
	}), nil
}

func (s *SchemaService) UpdateSchema(ctx context.Context, req *connect.Request[schema.UpdateSchemaRequest]) (*connect.Response[schema.UpdateSchemaResponse], error) {
	updatedSchema, err := model.UpdateSchemaInDB(req.Msg.Schema)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&schema.UpdateSchemaResponse{
		Schema: updatedSchema,
	}), nil
}

func (s *SchemaService) DeleteSchema(ctx context.Context, req *connect.Request[schema.DeleteSchemaRequest]) (*connect.Response[schema.DeleteSchemaResponse], error) {
	err := model.DeleteSchemaFromDB(req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&schema.DeleteSchemaResponse{
		Success: true,
	}), nil
}
