// package: pb
// file: proto/bug_service.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../google/api/annotations_pb";
import * as protoc_gen_swagger_options_annotations_pb from "../protoc-gen-swagger/options/annotations_pb";
import * as proto_bug_pb from "../proto/bug_pb";

export class BugSearchRequest extends jspb.Message {
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
  toObject(includeInstance?: boolean): BugSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: BugSearchRequest): BugSearchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BugSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BugSearchRequest;
  static deserializeBinaryFromReader(message: BugSearchRequest, reader: jspb.BinaryReader): BugSearchRequest;
}

export namespace BugSearchRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
    limit: number,
    offset: number,
    orderby: string,
    orderdesc: boolean,
  }
}

export class BugSearchResponse extends jspb.Message {
  clearBugsList(): void;
  getBugsList(): Array<proto_bug_pb.Bug>;
  setBugsList(value: Array<proto_bug_pb.Bug>): void;
  addBugs(value?: proto_bug_pb.Bug, index?: number): proto_bug_pb.Bug;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BugSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: BugSearchResponse): BugSearchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BugSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BugSearchResponse;
  static deserializeBinaryFromReader(message: BugSearchResponse, reader: jspb.BinaryReader): BugSearchResponse;
}

export namespace BugSearchResponse {
  export type AsObject = {
    bugsList: Array<proto_bug_pb.Bug.AsObject>,
    total: number,
  }
}

export class BugCreateRequest extends jspb.Message {
  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BugCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: BugCreateRequest): BugCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BugCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BugCreateRequest;
  static deserializeBinaryFromReader(message: BugCreateRequest, reader: jspb.BinaryReader): BugCreateRequest;
}

export namespace BugCreateRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
  }
}

export class BugCreateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BugCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: BugCreateResponse): BugCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BugCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BugCreateResponse;
  static deserializeBinaryFromReader(message: BugCreateResponse, reader: jspb.BinaryReader): BugCreateResponse;
}

export namespace BugCreateResponse {
  export type AsObject = {
    id: number,
  }
}

export class BugReadRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BugReadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: BugReadRequest): BugReadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BugReadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BugReadRequest;
  static deserializeBinaryFromReader(message: BugReadRequest, reader: jspb.BinaryReader): BugReadRequest;
}

export namespace BugReadRequest {
  export type AsObject = {
    id: number,
  }
}

export class BugReadResponse extends jspb.Message {
  hasBug(): boolean;
  clearBug(): void;
  getBug(): proto_bug_pb.Bug | undefined;
  setBug(value?: proto_bug_pb.Bug): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BugReadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: BugReadResponse): BugReadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BugReadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BugReadResponse;
  static deserializeBinaryFromReader(message: BugReadResponse, reader: jspb.BinaryReader): BugReadResponse;
}

export namespace BugReadResponse {
  export type AsObject = {
    bug?: proto_bug_pb.Bug.AsObject,
  }
}

export class BugUpdateRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BugUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: BugUpdateRequest): BugUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BugUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BugUpdateRequest;
  static deserializeBinaryFromReader(message: BugUpdateRequest, reader: jspb.BinaryReader): BugUpdateRequest;
}

export namespace BugUpdateRequest {
  export type AsObject = {
    id: number,
    valuesMap: Array<[string, string]>,
  }
}

export class BugUpdateResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BugUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: BugUpdateResponse): BugUpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BugUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BugUpdateResponse;
  static deserializeBinaryFromReader(message: BugUpdateResponse, reader: jspb.BinaryReader): BugUpdateResponse;
}

export namespace BugUpdateResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class BugDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BugDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: BugDeleteRequest): BugDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BugDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BugDeleteRequest;
  static deserializeBinaryFromReader(message: BugDeleteRequest, reader: jspb.BinaryReader): BugDeleteRequest;
}

export namespace BugDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class BugDeleteResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BugDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: BugDeleteResponse): BugDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BugDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BugDeleteResponse;
  static deserializeBinaryFromReader(message: BugDeleteResponse, reader: jspb.BinaryReader): BugDeleteResponse;
}

export namespace BugDeleteResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class BugPatchRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BugPatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: BugPatchRequest): BugPatchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BugPatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BugPatchRequest;
  static deserializeBinaryFromReader(message: BugPatchRequest, reader: jspb.BinaryReader): BugPatchRequest;
}

export namespace BugPatchRequest {
  export type AsObject = {
    id: number,
    key: string,
    value: string,
  }
}

export class BugPatchResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BugPatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: BugPatchResponse): BugPatchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BugPatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BugPatchResponse;
  static deserializeBinaryFromReader(message: BugPatchResponse, reader: jspb.BinaryReader): BugPatchResponse;
}

export namespace BugPatchResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

