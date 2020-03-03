// package: pb
// file: proto/character_service.proto

import * as jspb from "google-protobuf";


import * as proto_character_pb from "../proto/character_pb";

export class CharacterSearchRequest extends jspb.Message {
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
  toObject(includeInstance?: boolean): CharacterSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CharacterSearchRequest): CharacterSearchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharacterSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharacterSearchRequest;
  static deserializeBinaryFromReader(message: CharacterSearchRequest, reader: jspb.BinaryReader): CharacterSearchRequest;
}

export namespace CharacterSearchRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
    limit: number,
    offset: number,
    orderby: string,
    orderdesc: boolean,
  }
}

export class CharacterSearchResponse extends jspb.Message {
  clearCharactersList(): void;
  getCharactersList(): Array<proto_character_pb.Character>;
  setCharactersList(value: Array<proto_character_pb.Character>): void;
  addCharacters(value?: proto_character_pb.Character, index?: number): proto_character_pb.Character;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharacterSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CharacterSearchResponse): CharacterSearchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharacterSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharacterSearchResponse;
  static deserializeBinaryFromReader(message: CharacterSearchResponse, reader: jspb.BinaryReader): CharacterSearchResponse;
}

export namespace CharacterSearchResponse {
  export type AsObject = {
    charactersList: Array<proto_character_pb.Character.AsObject>,
    total: number,
  }
}

export class CharacterCreateRequest extends jspb.Message {
  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharacterCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CharacterCreateRequest): CharacterCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharacterCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharacterCreateRequest;
  static deserializeBinaryFromReader(message: CharacterCreateRequest, reader: jspb.BinaryReader): CharacterCreateRequest;
}

export namespace CharacterCreateRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
  }
}

export class CharacterCreateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharacterCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CharacterCreateResponse): CharacterCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharacterCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharacterCreateResponse;
  static deserializeBinaryFromReader(message: CharacterCreateResponse, reader: jspb.BinaryReader): CharacterCreateResponse;
}

export namespace CharacterCreateResponse {
  export type AsObject = {
    id: number,
  }
}

export class CharacterReadRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharacterReadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CharacterReadRequest): CharacterReadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharacterReadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharacterReadRequest;
  static deserializeBinaryFromReader(message: CharacterReadRequest, reader: jspb.BinaryReader): CharacterReadRequest;
}

export namespace CharacterReadRequest {
  export type AsObject = {
    id: number,
  }
}

export class CharacterReadResponse extends jspb.Message {
  hasCharacter(): boolean;
  clearCharacter(): void;
  getCharacter(): proto_character_pb.Character | undefined;
  setCharacter(value?: proto_character_pb.Character): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharacterReadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CharacterReadResponse): CharacterReadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharacterReadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharacterReadResponse;
  static deserializeBinaryFromReader(message: CharacterReadResponse, reader: jspb.BinaryReader): CharacterReadResponse;
}

export namespace CharacterReadResponse {
  export type AsObject = {
    character?: proto_character_pb.Character.AsObject,
  }
}

export class CharacterUpdateRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharacterUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CharacterUpdateRequest): CharacterUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharacterUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharacterUpdateRequest;
  static deserializeBinaryFromReader(message: CharacterUpdateRequest, reader: jspb.BinaryReader): CharacterUpdateRequest;
}

export namespace CharacterUpdateRequest {
  export type AsObject = {
    id: number,
    valuesMap: Array<[string, string]>,
  }
}

export class CharacterUpdateResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharacterUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CharacterUpdateResponse): CharacterUpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharacterUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharacterUpdateResponse;
  static deserializeBinaryFromReader(message: CharacterUpdateResponse, reader: jspb.BinaryReader): CharacterUpdateResponse;
}

export namespace CharacterUpdateResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class CharacterDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharacterDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CharacterDeleteRequest): CharacterDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharacterDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharacterDeleteRequest;
  static deserializeBinaryFromReader(message: CharacterDeleteRequest, reader: jspb.BinaryReader): CharacterDeleteRequest;
}

export namespace CharacterDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class CharacterDeleteResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharacterDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CharacterDeleteResponse): CharacterDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharacterDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharacterDeleteResponse;
  static deserializeBinaryFromReader(message: CharacterDeleteResponse, reader: jspb.BinaryReader): CharacterDeleteResponse;
}

export namespace CharacterDeleteResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class CharacterPatchRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharacterPatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CharacterPatchRequest): CharacterPatchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharacterPatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharacterPatchRequest;
  static deserializeBinaryFromReader(message: CharacterPatchRequest, reader: jspb.BinaryReader): CharacterPatchRequest;
}

export namespace CharacterPatchRequest {
  export type AsObject = {
    id: number,
    key: string,
    value: string,
  }
}

export class CharacterPatchResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharacterPatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CharacterPatchResponse): CharacterPatchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharacterPatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharacterPatchResponse;
  static deserializeBinaryFromReader(message: CharacterPatchResponse, reader: jspb.BinaryReader): CharacterPatchResponse;
}

export namespace CharacterPatchResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

