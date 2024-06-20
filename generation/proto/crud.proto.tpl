syntax = "proto3";

package {{.PackageName}};

// The {{.ModelName}} service definition.
service {{.ModelName}}Service {
  rpc Create{{.ModelName}}(Create{{.ModelName}}Request) returns (Create{{.ModelName}}Response);
  rpc Get{{.ModelName}}(Get{{.ModelName}}Request) returns (Get{{.ModelName}}Response);
  rpc Update{{.ModelName}}(Update{{.ModelName}}Request) returns (Update{{.ModelName}}Response);
  rpc Delete{{.ModelName}}(Delete{{.ModelName}}Request) returns (Delete{{.ModelName}}Response);
}

message Create{{.ModelName}}Request {
  {{.ModelName}} {{.ModelName | lower}} = 1;
}

message Create{{.ModelName}}Response {
  {{.ModelName}} {{.ModelName | lower}} = 1;
}

message Get{{.ModelName}}Request {
  int32 id = 1;
}

message Get{{.ModelName}}Response {
  {{.ModelName}} {{.ModelName | lower}} = 1;
}

message Update{{.ModelName}}Request {
  {{.ModelName}} {{.ModelName | lower}} = 1;
}

message Update{{.ModelName}}Response {
  {{.ModelName}} {{.ModelName | lower}} = 1;
}

message Delete{{.ModelName}}Request {
  int32 id = 1;
}

message Delete{{.ModelName}}Response {
  bool success = 1;
}

// The {{.ModelName}} message represents a {{.ModelName | lower}} in the system.
message {{.ModelName}} {
{{- range $i, $field := .Fields }}
  {{ $field.Type }} {{ $field.Name }} = {{ inc $i }};
{{- end }}
}
