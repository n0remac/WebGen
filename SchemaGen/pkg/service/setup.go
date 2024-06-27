package service

import (
	"fmt"
	"net/http"
	"SchemaGen/gen/proto/schema/schemaconnect"

	"github.com/bufbuild/connect-go"
)

func SetupServices(apiRoot *http.ServeMux, interceptors connect.Option) []string {
	fmt.Println("Setting up services")
	schemaService := &SchemaService{}
	apiRoot.Handle(schemaconnect.NewSchemaServiceHandler(schemaService, interceptors))

	reflectorServices := []string{
		"schemas.schemaService",
	}

	return reflectorServices
}
