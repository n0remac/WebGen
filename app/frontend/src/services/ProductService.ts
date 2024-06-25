import { ProductService } from "../rpc/proto/product/product_connect";
import { transport } from "./service";
import { createPromiseClient } from "@bufbuild/connect";

export const productService = createPromiseClient(ProductService, transport);
