// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file proto/product/product.proto (package product, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message product.CreateProductRequest
 */
export class CreateProductRequest extends Message<CreateProductRequest> {
  /**
   * @generated from field: product.Product product = 1;
   */
  product?: Product;

  constructor(data?: PartialMessage<CreateProductRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "product.CreateProductRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "product", kind: "message", T: Product },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateProductRequest {
    return new CreateProductRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateProductRequest {
    return new CreateProductRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateProductRequest {
    return new CreateProductRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateProductRequest | PlainMessage<CreateProductRequest> | undefined, b: CreateProductRequest | PlainMessage<CreateProductRequest> | undefined): boolean {
    return proto3.util.equals(CreateProductRequest, a, b);
  }
}

/**
 * @generated from message product.CreateProductResponse
 */
export class CreateProductResponse extends Message<CreateProductResponse> {
  /**
   * @generated from field: product.Product product = 1;
   */
  product?: Product;

  constructor(data?: PartialMessage<CreateProductResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "product.CreateProductResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "product", kind: "message", T: Product },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateProductResponse {
    return new CreateProductResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateProductResponse {
    return new CreateProductResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateProductResponse {
    return new CreateProductResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateProductResponse | PlainMessage<CreateProductResponse> | undefined, b: CreateProductResponse | PlainMessage<CreateProductResponse> | undefined): boolean {
    return proto3.util.equals(CreateProductResponse, a, b);
  }
}

/**
 * @generated from message product.GetProductRequest
 */
export class GetProductRequest extends Message<GetProductRequest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<GetProductRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "product.GetProductRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetProductRequest {
    return new GetProductRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetProductRequest {
    return new GetProductRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetProductRequest {
    return new GetProductRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetProductRequest | PlainMessage<GetProductRequest> | undefined, b: GetProductRequest | PlainMessage<GetProductRequest> | undefined): boolean {
    return proto3.util.equals(GetProductRequest, a, b);
  }
}

/**
 * @generated from message product.GetProductResponse
 */
export class GetProductResponse extends Message<GetProductResponse> {
  /**
   * @generated from field: product.Product product = 1;
   */
  product?: Product;

  constructor(data?: PartialMessage<GetProductResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "product.GetProductResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "product", kind: "message", T: Product },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetProductResponse {
    return new GetProductResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetProductResponse {
    return new GetProductResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetProductResponse {
    return new GetProductResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetProductResponse | PlainMessage<GetProductResponse> | undefined, b: GetProductResponse | PlainMessage<GetProductResponse> | undefined): boolean {
    return proto3.util.equals(GetProductResponse, a, b);
  }
}

/**
 * @generated from message product.UpdateProductRequest
 */
export class UpdateProductRequest extends Message<UpdateProductRequest> {
  /**
   * @generated from field: product.Product product = 1;
   */
  product?: Product;

  constructor(data?: PartialMessage<UpdateProductRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "product.UpdateProductRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "product", kind: "message", T: Product },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateProductRequest {
    return new UpdateProductRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateProductRequest {
    return new UpdateProductRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateProductRequest {
    return new UpdateProductRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateProductRequest | PlainMessage<UpdateProductRequest> | undefined, b: UpdateProductRequest | PlainMessage<UpdateProductRequest> | undefined): boolean {
    return proto3.util.equals(UpdateProductRequest, a, b);
  }
}

/**
 * @generated from message product.UpdateProductResponse
 */
export class UpdateProductResponse extends Message<UpdateProductResponse> {
  /**
   * @generated from field: product.Product product = 1;
   */
  product?: Product;

  constructor(data?: PartialMessage<UpdateProductResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "product.UpdateProductResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "product", kind: "message", T: Product },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateProductResponse {
    return new UpdateProductResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateProductResponse {
    return new UpdateProductResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateProductResponse {
    return new UpdateProductResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateProductResponse | PlainMessage<UpdateProductResponse> | undefined, b: UpdateProductResponse | PlainMessage<UpdateProductResponse> | undefined): boolean {
    return proto3.util.equals(UpdateProductResponse, a, b);
  }
}

/**
 * @generated from message product.DeleteProductRequest
 */
export class DeleteProductRequest extends Message<DeleteProductRequest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<DeleteProductRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "product.DeleteProductRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteProductRequest {
    return new DeleteProductRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteProductRequest {
    return new DeleteProductRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteProductRequest {
    return new DeleteProductRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteProductRequest | PlainMessage<DeleteProductRequest> | undefined, b: DeleteProductRequest | PlainMessage<DeleteProductRequest> | undefined): boolean {
    return proto3.util.equals(DeleteProductRequest, a, b);
  }
}

/**
 * @generated from message product.DeleteProductResponse
 */
export class DeleteProductResponse extends Message<DeleteProductResponse> {
  /**
   * @generated from field: bool success = 1;
   */
  success = false;

  constructor(data?: PartialMessage<DeleteProductResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "product.DeleteProductResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "success", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteProductResponse {
    return new DeleteProductResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteProductResponse {
    return new DeleteProductResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteProductResponse {
    return new DeleteProductResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteProductResponse | PlainMessage<DeleteProductResponse> | undefined, b: DeleteProductResponse | PlainMessage<DeleteProductResponse> | undefined): boolean {
    return proto3.util.equals(DeleteProductResponse, a, b);
  }
}

/**
 * The Product message represents a product in the system.
 *
 * @generated from message product.Product
 */
export class Product extends Message<Product> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: int32 amount = 3;
   */
  amount = 0;

  /**
   * @generated from field: string description = 4;
   */
  description = "";

  constructor(data?: PartialMessage<Product>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "product.Product";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "amount", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Product {
    return new Product().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Product {
    return new Product().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Product {
    return new Product().fromJsonString(jsonString, options);
  }

  static equals(a: Product | PlainMessage<Product> | undefined, b: Product | PlainMessage<Product> | undefined): boolean {
    return proto3.util.equals(Product, a, b);
  }
}
