// package: pb
// file: proto/npc_service.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../google/api/annotations_pb";
import * as proto_npc_pb from "../proto/npc_pb";

export class NpcSearchRequest extends jspb.Message {
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
  toObject(includeInstance?: boolean): NpcSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: NpcSearchRequest): NpcSearchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NpcSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NpcSearchRequest;
  static deserializeBinaryFromReader(message: NpcSearchRequest, reader: jspb.BinaryReader): NpcSearchRequest;
}

export namespace NpcSearchRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
    limit: number,
    offset: number,
    orderby: string,
    orderdesc: boolean,
  }
}

export class NpcSearchResponse extends jspb.Message {
  clearNpcsList(): void;
  getNpcsList(): Array<proto_npc_pb.Npc>;
  setNpcsList(value: Array<proto_npc_pb.Npc>): void;
  addNpcs(value?: proto_npc_pb.Npc, index?: number): proto_npc_pb.Npc;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NpcSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: NpcSearchResponse): NpcSearchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NpcSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NpcSearchResponse;
  static deserializeBinaryFromReader(message: NpcSearchResponse, reader: jspb.BinaryReader): NpcSearchResponse;
}

export namespace NpcSearchResponse {
  export type AsObject = {
    npcsList: Array<proto_npc_pb.Npc.AsObject>,
    total: number,
  }
}

export class NpcCreateRequest extends jspb.Message {
  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NpcCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: NpcCreateRequest): NpcCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NpcCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NpcCreateRequest;
  static deserializeBinaryFromReader(message: NpcCreateRequest, reader: jspb.BinaryReader): NpcCreateRequest;
}

export namespace NpcCreateRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
  }
}

export class NpcCreateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NpcCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: NpcCreateResponse): NpcCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NpcCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NpcCreateResponse;
  static deserializeBinaryFromReader(message: NpcCreateResponse, reader: jspb.BinaryReader): NpcCreateResponse;
}

export namespace NpcCreateResponse {
  export type AsObject = {
    id: number,
  }
}

export class NpcReadRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NpcReadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: NpcReadRequest): NpcReadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NpcReadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NpcReadRequest;
  static deserializeBinaryFromReader(message: NpcReadRequest, reader: jspb.BinaryReader): NpcReadRequest;
}

export namespace NpcReadRequest {
  export type AsObject = {
    id: number,
  }
}

export class NpcReadResponse extends jspb.Message {
  hasNpc(): boolean;
  clearNpc(): void;
  getNpc(): proto_npc_pb.Npc | undefined;
  setNpc(value?: proto_npc_pb.Npc): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NpcReadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: NpcReadResponse): NpcReadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NpcReadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NpcReadResponse;
  static deserializeBinaryFromReader(message: NpcReadResponse, reader: jspb.BinaryReader): NpcReadResponse;
}

export namespace NpcReadResponse {
  export type AsObject = {
    npc?: proto_npc_pb.Npc.AsObject,
  }
}

export class NpcUpdateRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NpcUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: NpcUpdateRequest): NpcUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NpcUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NpcUpdateRequest;
  static deserializeBinaryFromReader(message: NpcUpdateRequest, reader: jspb.BinaryReader): NpcUpdateRequest;
}

export namespace NpcUpdateRequest {
  export type AsObject = {
    id: number,
    valuesMap: Array<[string, string]>,
  }
}

export class NpcUpdateResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NpcUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: NpcUpdateResponse): NpcUpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NpcUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NpcUpdateResponse;
  static deserializeBinaryFromReader(message: NpcUpdateResponse, reader: jspb.BinaryReader): NpcUpdateResponse;
}

export namespace NpcUpdateResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class NpcDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NpcDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: NpcDeleteRequest): NpcDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NpcDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NpcDeleteRequest;
  static deserializeBinaryFromReader(message: NpcDeleteRequest, reader: jspb.BinaryReader): NpcDeleteRequest;
}

export namespace NpcDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class NpcDeleteResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NpcDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: NpcDeleteResponse): NpcDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NpcDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NpcDeleteResponse;
  static deserializeBinaryFromReader(message: NpcDeleteResponse, reader: jspb.BinaryReader): NpcDeleteResponse;
}

export namespace NpcDeleteResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class NpcPatchRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NpcPatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: NpcPatchRequest): NpcPatchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NpcPatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NpcPatchRequest;
  static deserializeBinaryFromReader(message: NpcPatchRequest, reader: jspb.BinaryReader): NpcPatchRequest;
}

export namespace NpcPatchRequest {
  export type AsObject = {
    id: number,
    key: string,
    value: string,
  }
}

export class NpcPatchResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NpcPatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: NpcPatchResponse): NpcPatchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NpcPatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NpcPatchResponse;
  static deserializeBinaryFromReader(message: NpcPatchResponse, reader: jspb.BinaryReader): NpcPatchResponse;
}

export namespace NpcPatchResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

