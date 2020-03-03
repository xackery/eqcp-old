// package: pb
// file: proto/handin_service.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../google/api/annotations_pb";
import * as protoc_gen_swagger_options_annotations_pb from "../protoc-gen-swagger/options/annotations_pb";
import * as proto_handin_pb from "../proto/handin_pb";

export class HandinSearchRequest extends jspb.Message {
  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  getLimit(): number;
  setLimit(value: number): void;

  getOffset(): number;
  setOffset(value: number): void;

  getOrderby(): string;
  setOrderby(value: string): void;

  getOrderdesc(): boolean;
  setOrderdesc(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HandinSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HandinSearchRequest): HandinSearchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HandinSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HandinSearchRequest;
  static deserializeBinaryFromReader(message: HandinSearchRequest, reader: jspb.BinaryReader): HandinSearchRequest;
}

export namespace HandinSearchRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
    limit: number,
    offset: number,
    orderby: string,
    orderdesc: boolean,
  }
}

export class HandinSearchResponse extends jspb.Message {
  clearHandinsList(): void;
  getHandinsList(): Array<proto_handin_pb.Handin>;
  setHandinsList(value: Array<proto_handin_pb.Handin>): void;
  addHandins(value?: proto_handin_pb.Handin, index?: number): proto_handin_pb.Handin;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HandinSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HandinSearchResponse): HandinSearchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HandinSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HandinSearchResponse;
  static deserializeBinaryFromReader(message: HandinSearchResponse, reader: jspb.BinaryReader): HandinSearchResponse;
}

export namespace HandinSearchResponse {
  export type AsObject = {
    handinsList: Array<proto_handin_pb.Handin.AsObject>,
    total: number,
  }
}

export class HandinCreateRequest extends jspb.Message {
  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HandinCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HandinCreateRequest): HandinCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HandinCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HandinCreateRequest;
  static deserializeBinaryFromReader(message: HandinCreateRequest, reader: jspb.BinaryReader): HandinCreateRequest;
}

export namespace HandinCreateRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
  }
}

export class HandinCreateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HandinCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HandinCreateResponse): HandinCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HandinCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HandinCreateResponse;
  static deserializeBinaryFromReader(message: HandinCreateResponse, reader: jspb.BinaryReader): HandinCreateResponse;
}

export namespace HandinCreateResponse {
  export type AsObject = {
    id: number,
  }
}

export class HandinReadRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HandinReadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HandinReadRequest): HandinReadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HandinReadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HandinReadRequest;
  static deserializeBinaryFromReader(message: HandinReadRequest, reader: jspb.BinaryReader): HandinReadRequest;
}

export namespace HandinReadRequest {
  export type AsObject = {
    id: number,
  }
}

export class HandinReadResponse extends jspb.Message {
  hasHandin(): boolean;
  clearHandin(): void;
  getHandin(): proto_handin_pb.Handin | undefined;
  setHandin(value?: proto_handin_pb.Handin): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HandinReadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HandinReadResponse): HandinReadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HandinReadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HandinReadResponse;
  static deserializeBinaryFromReader(message: HandinReadResponse, reader: jspb.BinaryReader): HandinReadResponse;
}

export namespace HandinReadResponse {
  export type AsObject = {
    handin?: proto_handin_pb.Handin.AsObject,
  }
}

export class HandinUpdateRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HandinUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HandinUpdateRequest): HandinUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HandinUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HandinUpdateRequest;
  static deserializeBinaryFromReader(message: HandinUpdateRequest, reader: jspb.BinaryReader): HandinUpdateRequest;
}

export namespace HandinUpdateRequest {
  export type AsObject = {
    id: number,
    valuesMap: Array<[string, string]>,
  }
}

export class HandinUpdateResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HandinUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HandinUpdateResponse): HandinUpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HandinUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HandinUpdateResponse;
  static deserializeBinaryFromReader(message: HandinUpdateResponse, reader: jspb.BinaryReader): HandinUpdateResponse;
}

export namespace HandinUpdateResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class HandinDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HandinDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HandinDeleteRequest): HandinDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HandinDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HandinDeleteRequest;
  static deserializeBinaryFromReader(message: HandinDeleteRequest, reader: jspb.BinaryReader): HandinDeleteRequest;
}

export namespace HandinDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class HandinDeleteResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HandinDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HandinDeleteResponse): HandinDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HandinDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HandinDeleteResponse;
  static deserializeBinaryFromReader(message: HandinDeleteResponse, reader: jspb.BinaryReader): HandinDeleteResponse;
}

export namespace HandinDeleteResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class HandinPatchRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HandinPatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HandinPatchRequest): HandinPatchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HandinPatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HandinPatchRequest;
  static deserializeBinaryFromReader(message: HandinPatchRequest, reader: jspb.BinaryReader): HandinPatchRequest;
}

export namespace HandinPatchRequest {
  export type AsObject = {
    id: number,
    key: string,
    value: string,
  }
}

export class HandinPatchResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HandinPatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HandinPatchResponse): HandinPatchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HandinPatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HandinPatchResponse;
  static deserializeBinaryFromReader(message: HandinPatchResponse, reader: jspb.BinaryReader): HandinPatchResponse;
}

export namespace HandinPatchResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

