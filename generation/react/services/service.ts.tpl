import { {{.ServiceName}}Service } from "../rpc/proto/{{.ServiceName | lower}}/{{.ServiceName | lower}}_connect";
import { transport } from "./service";
import { createPromiseClient } from "@bufbuild/connect";

export const {{.ServiceName | lower}}Service = createPromiseClient({{.ServiceName}}Service, transport);
