package database

func SetupDatabase() {
	{{- range .Models }}
	create{{ .Name }}Table()
	{{- end }}
}
