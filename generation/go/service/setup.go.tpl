package service

import (
	"fmt"
	"net/http"
	{{- range .Services }}
	"{{ .ProjectName }}/gen/proto/{{ lower .ServiceName }}/{{ lower .ServiceName }}connect"
	{{- end }}

	"github.com/bufbuild/connect-go"
)

func SetupServices(apiRoot *http.ServeMux, interceptors connect.Option) []string {
	fmt.Println("Setting up services")

	{{- range .Services }}
	{{ lower .ServiceName }}Service := &{{ .ServiceType }}{}
	apiRoot.Handle({{ lower .ServiceName }}connect.New{{ .ServiceType }}Handler({{ lower .ServiceName }}Service, interceptors))
	{{- end }}

	reflectorServices := []string{
		{{- range .Services }}
		"{{ lower .ServiceName }}s.{{ lower .ServiceName }}Service",
		{{- end }}
	}

	return reflectorServices
}
