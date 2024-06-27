import { SchemaService } from "../rpc/proto/schema/schema_connect";
import { transport } from "./service";
import { createPromiseClient } from "@bufbuild/connect";

export const schemaService = createPromiseClient(SchemaService, transport);
