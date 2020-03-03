// package: pb
// file: proto/spell_service.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../google/api/annotations_pb";
import * as protoc_gen_swagger_options_annotations_pb from "../protoc-gen-swagger/options/annotations_pb";
import * as proto_spell_pb from "../proto/spell_pb";

export class SpellSearchRequest extends jspb.Message {
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
  toObject(includeInstance?: boolean): SpellSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SpellSearchRequest): SpellSearchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SpellSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpellSearchRequest;
  static deserializeBinaryFromReader(message: SpellSearchRequest, reader: jspb.BinaryReader): SpellSearchRequest;
}

export namespace SpellSearchRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
    limit: number,
    offset: number,
    orderby: string,
    orderdesc: boolean,
  }
}

export class SpellSearchResponse extends jspb.Message {
  clearSpellsList(): void;
  getSpellsList(): Array<proto_spell_pb.Spell>;
  setSpellsList(value: Array<proto_spell_pb.Spell>): void;
  addSpells(value?: proto_spell_pb.Spell, index?: number): proto_spell_pb.Spell;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpellSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SpellSearchResponse): SpellSearchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SpellSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpellSearchResponse;
  static deserializeBinaryFromReader(message: SpellSearchResponse, reader: jspb.BinaryReader): SpellSearchResponse;
}

export namespace SpellSearchResponse {
  export type AsObject = {
    spellsList: Array<proto_spell_pb.Spell.AsObject>,
    total: number,
  }
}

export class SpellCreateRequest extends jspb.Message {
  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpellCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SpellCreateRequest): SpellCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SpellCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpellCreateRequest;
  static deserializeBinaryFromReader(message: SpellCreateRequest, reader: jspb.BinaryReader): SpellCreateRequest;
}

export namespace SpellCreateRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
  }
}

export class SpellCreateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpellCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SpellCreateResponse): SpellCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SpellCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpellCreateResponse;
  static deserializeBinaryFromReader(message: SpellCreateResponse, reader: jspb.BinaryReader): SpellCreateResponse;
}

export namespace SpellCreateResponse {
  export type AsObject = {
    id: number,
  }
}

export class SpellReadRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpellReadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SpellReadRequest): SpellReadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SpellReadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpellReadRequest;
  static deserializeBinaryFromReader(message: SpellReadRequest, reader: jspb.BinaryReader): SpellReadRequest;
}

export namespace SpellReadRequest {
  export type AsObject = {
    id: number,
  }
}

export class SpellReadResponse extends jspb.Message {
  hasSpell(): boolean;
  clearSpell(): void;
  getSpell(): proto_spell_pb.Spell | undefined;
  setSpell(value?: proto_spell_pb.Spell): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpellReadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SpellReadResponse): SpellReadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SpellReadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpellReadResponse;
  static deserializeBinaryFromReader(message: SpellReadResponse, reader: jspb.BinaryReader): SpellReadResponse;
}

export namespace SpellReadResponse {
  export type AsObject = {
    spell?: proto_spell_pb.Spell.AsObject,
  }
}

export class SpellUpdateRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpellUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SpellUpdateRequest): SpellUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SpellUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpellUpdateRequest;
  static deserializeBinaryFromReader(message: SpellUpdateRequest, reader: jspb.BinaryReader): SpellUpdateRequest;
}

export namespace SpellUpdateRequest {
  export type AsObject = {
    id: number,
    valuesMap: Array<[string, string]>,
  }
}

export class SpellUpdateResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpellUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SpellUpdateResponse): SpellUpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SpellUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpellUpdateResponse;
  static deserializeBinaryFromReader(message: SpellUpdateResponse, reader: jspb.BinaryReader): SpellUpdateResponse;
}

export namespace SpellUpdateResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class SpellDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpellDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SpellDeleteRequest): SpellDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SpellDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpellDeleteRequest;
  static deserializeBinaryFromReader(message: SpellDeleteRequest, reader: jspb.BinaryReader): SpellDeleteRequest;
}

export namespace SpellDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class SpellDeleteResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpellDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SpellDeleteResponse): SpellDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SpellDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpellDeleteResponse;
  static deserializeBinaryFromReader(message: SpellDeleteResponse, reader: jspb.BinaryReader): SpellDeleteResponse;
}

export namespace SpellDeleteResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class SpellPatchRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpellPatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SpellPatchRequest): SpellPatchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SpellPatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpellPatchRequest;
  static deserializeBinaryFromReader(message: SpellPatchRequest, reader: jspb.BinaryReader): SpellPatchRequest;
}

export namespace SpellPatchRequest {
  export type AsObject = {
    id: number,
    key: string,
    value: string,
  }
}

export class SpellPatchResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpellPatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SpellPatchResponse): SpellPatchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SpellPatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpellPatchResponse;
  static deserializeBinaryFromReader(message: SpellPatchResponse, reader: jspb.BinaryReader): SpellPatchResponse;
}

export namespace SpellPatchResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

