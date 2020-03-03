// package: pb
// file: proto/player_speech_service.proto

import * as jspb from "google-protobuf";

import * as proto_player_speech_pb from "../proto/player_speech_pb";

export class PlayerSpeechSearchRequest extends jspb.Message {
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
  toObject(includeInstance?: boolean): PlayerSpeechSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerSpeechSearchRequest): PlayerSpeechSearchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerSpeechSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerSpeechSearchRequest;
  static deserializeBinaryFromReader(message: PlayerSpeechSearchRequest, reader: jspb.BinaryReader): PlayerSpeechSearchRequest;
}

export namespace PlayerSpeechSearchRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
    limit: number,
    offset: number,
    orderby: string,
    orderdesc: boolean,
  }
}

export class PlayerSpeechSearchResponse extends jspb.Message {
  clearPlayerspeechsList(): void;
  getPlayerspeechsList(): Array<proto_player_speech_pb.PlayerSpeech>;
  setPlayerspeechsList(value: Array<proto_player_speech_pb.PlayerSpeech>): void;
  addPlayerspeechs(value?: proto_player_speech_pb.PlayerSpeech, index?: number): proto_player_speech_pb.PlayerSpeech;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerSpeechSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerSpeechSearchResponse): PlayerSpeechSearchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerSpeechSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerSpeechSearchResponse;
  static deserializeBinaryFromReader(message: PlayerSpeechSearchResponse, reader: jspb.BinaryReader): PlayerSpeechSearchResponse;
}

export namespace PlayerSpeechSearchResponse {
  export type AsObject = {
    playerspeechsList: Array<proto_player_speech_pb.PlayerSpeech.AsObject>,
    total: number,
  }
}

export class PlayerSpeechCreateRequest extends jspb.Message {
  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerSpeechCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerSpeechCreateRequest): PlayerSpeechCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerSpeechCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerSpeechCreateRequest;
  static deserializeBinaryFromReader(message: PlayerSpeechCreateRequest, reader: jspb.BinaryReader): PlayerSpeechCreateRequest;
}

export namespace PlayerSpeechCreateRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
  }
}

export class PlayerSpeechCreateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerSpeechCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerSpeechCreateResponse): PlayerSpeechCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerSpeechCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerSpeechCreateResponse;
  static deserializeBinaryFromReader(message: PlayerSpeechCreateResponse, reader: jspb.BinaryReader): PlayerSpeechCreateResponse;
}

export namespace PlayerSpeechCreateResponse {
  export type AsObject = {
    id: number,
  }
}

export class PlayerSpeechReadRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerSpeechReadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerSpeechReadRequest): PlayerSpeechReadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerSpeechReadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerSpeechReadRequest;
  static deserializeBinaryFromReader(message: PlayerSpeechReadRequest, reader: jspb.BinaryReader): PlayerSpeechReadRequest;
}

export namespace PlayerSpeechReadRequest {
  export type AsObject = {
    id: number,
  }
}

export class PlayerSpeechReadResponse extends jspb.Message {
  hasPlayerspeech(): boolean;
  clearPlayerspeech(): void;
  getPlayerspeech(): proto_player_speech_pb.PlayerSpeech | undefined;
  setPlayerspeech(value?: proto_player_speech_pb.PlayerSpeech): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerSpeechReadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerSpeechReadResponse): PlayerSpeechReadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerSpeechReadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerSpeechReadResponse;
  static deserializeBinaryFromReader(message: PlayerSpeechReadResponse, reader: jspb.BinaryReader): PlayerSpeechReadResponse;
}

export namespace PlayerSpeechReadResponse {
  export type AsObject = {
    playerspeech?: proto_player_speech_pb.PlayerSpeech.AsObject,
  }
}

export class PlayerSpeechUpdateRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerSpeechUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerSpeechUpdateRequest): PlayerSpeechUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerSpeechUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerSpeechUpdateRequest;
  static deserializeBinaryFromReader(message: PlayerSpeechUpdateRequest, reader: jspb.BinaryReader): PlayerSpeechUpdateRequest;
}

export namespace PlayerSpeechUpdateRequest {
  export type AsObject = {
    id: number,
    valuesMap: Array<[string, string]>,
  }
}

export class PlayerSpeechUpdateResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerSpeechUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerSpeechUpdateResponse): PlayerSpeechUpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerSpeechUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerSpeechUpdateResponse;
  static deserializeBinaryFromReader(message: PlayerSpeechUpdateResponse, reader: jspb.BinaryReader): PlayerSpeechUpdateResponse;
}

export namespace PlayerSpeechUpdateResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class PlayerSpeechDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerSpeechDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerSpeechDeleteRequest): PlayerSpeechDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerSpeechDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerSpeechDeleteRequest;
  static deserializeBinaryFromReader(message: PlayerSpeechDeleteRequest, reader: jspb.BinaryReader): PlayerSpeechDeleteRequest;
}

export namespace PlayerSpeechDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class PlayerSpeechDeleteResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerSpeechDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerSpeechDeleteResponse): PlayerSpeechDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerSpeechDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerSpeechDeleteResponse;
  static deserializeBinaryFromReader(message: PlayerSpeechDeleteResponse, reader: jspb.BinaryReader): PlayerSpeechDeleteResponse;
}

export namespace PlayerSpeechDeleteResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class PlayerSpeechPatchRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerSpeechPatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerSpeechPatchRequest): PlayerSpeechPatchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerSpeechPatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerSpeechPatchRequest;
  static deserializeBinaryFromReader(message: PlayerSpeechPatchRequest, reader: jspb.BinaryReader): PlayerSpeechPatchRequest;
}

export namespace PlayerSpeechPatchRequest {
  export type AsObject = {
    id: number,
    key: string,
    value: string,
  }
}

export class PlayerSpeechPatchResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerSpeechPatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerSpeechPatchResponse): PlayerSpeechPatchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerSpeechPatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerSpeechPatchResponse;
  static deserializeBinaryFromReader(message: PlayerSpeechPatchResponse, reader: jspb.BinaryReader): PlayerSpeechPatchResponse;
}

export namespace PlayerSpeechPatchResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

