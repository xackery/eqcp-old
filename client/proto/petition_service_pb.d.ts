// package: pb
// file: proto/petition_service.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../google/api/annotations_pb";
import * as proto_petition_pb from "../proto/petition_pb";

export class PetitionSearchRequest extends jspb.Message {
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
  toObject(includeInstance?: boolean): PetitionSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PetitionSearchRequest): PetitionSearchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PetitionSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PetitionSearchRequest;
  static deserializeBinaryFromReader(message: PetitionSearchRequest, reader: jspb.BinaryReader): PetitionSearchRequest;
}

export namespace PetitionSearchRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
    limit: number,
    offset: number,
    orderby: string,
    orderdesc: boolean,
  }
}

export class PetitionSearchResponse extends jspb.Message {
  clearPetitionsList(): void;
  getPetitionsList(): Array<proto_petition_pb.Petition>;
  setPetitionsList(value: Array<proto_petition_pb.Petition>): void;
  addPetitions(value?: proto_petition_pb.Petition, index?: number): proto_petition_pb.Petition;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PetitionSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PetitionSearchResponse): PetitionSearchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PetitionSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PetitionSearchResponse;
  static deserializeBinaryFromReader(message: PetitionSearchResponse, reader: jspb.BinaryReader): PetitionSearchResponse;
}

export namespace PetitionSearchResponse {
  export type AsObject = {
    petitionsList: Array<proto_petition_pb.Petition.AsObject>,
    total: number,
  }
}

export class PetitionCreateRequest extends jspb.Message {
  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PetitionCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PetitionCreateRequest): PetitionCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PetitionCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PetitionCreateRequest;
  static deserializeBinaryFromReader(message: PetitionCreateRequest, reader: jspb.BinaryReader): PetitionCreateRequest;
}

export namespace PetitionCreateRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
  }
}

export class PetitionCreateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PetitionCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PetitionCreateResponse): PetitionCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PetitionCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PetitionCreateResponse;
  static deserializeBinaryFromReader(message: PetitionCreateResponse, reader: jspb.BinaryReader): PetitionCreateResponse;
}

export namespace PetitionCreateResponse {
  export type AsObject = {
    id: number,
  }
}

export class PetitionReadRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PetitionReadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PetitionReadRequest): PetitionReadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PetitionReadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PetitionReadRequest;
  static deserializeBinaryFromReader(message: PetitionReadRequest, reader: jspb.BinaryReader): PetitionReadRequest;
}

export namespace PetitionReadRequest {
  export type AsObject = {
    id: number,
  }
}

export class PetitionReadResponse extends jspb.Message {
  hasPetition(): boolean;
  clearPetition(): void;
  getPetition(): proto_petition_pb.Petition | undefined;
  setPetition(value?: proto_petition_pb.Petition): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PetitionReadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PetitionReadResponse): PetitionReadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PetitionReadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PetitionReadResponse;
  static deserializeBinaryFromReader(message: PetitionReadResponse, reader: jspb.BinaryReader): PetitionReadResponse;
}

export namespace PetitionReadResponse {
  export type AsObject = {
    petition?: proto_petition_pb.Petition.AsObject,
  }
}

export class PetitionUpdateRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PetitionUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PetitionUpdateRequest): PetitionUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PetitionUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PetitionUpdateRequest;
  static deserializeBinaryFromReader(message: PetitionUpdateRequest, reader: jspb.BinaryReader): PetitionUpdateRequest;
}

export namespace PetitionUpdateRequest {
  export type AsObject = {
    id: number,
    valuesMap: Array<[string, string]>,
  }
}

export class PetitionUpdateResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PetitionUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PetitionUpdateResponse): PetitionUpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PetitionUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PetitionUpdateResponse;
  static deserializeBinaryFromReader(message: PetitionUpdateResponse, reader: jspb.BinaryReader): PetitionUpdateResponse;
}

export namespace PetitionUpdateResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class PetitionDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PetitionDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PetitionDeleteRequest): PetitionDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PetitionDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PetitionDeleteRequest;
  static deserializeBinaryFromReader(message: PetitionDeleteRequest, reader: jspb.BinaryReader): PetitionDeleteRequest;
}

export namespace PetitionDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class PetitionDeleteResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PetitionDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PetitionDeleteResponse): PetitionDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PetitionDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PetitionDeleteResponse;
  static deserializeBinaryFromReader(message: PetitionDeleteResponse, reader: jspb.BinaryReader): PetitionDeleteResponse;
}

export namespace PetitionDeleteResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class PetitionPatchRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PetitionPatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PetitionPatchRequest): PetitionPatchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PetitionPatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PetitionPatchRequest;
  static deserializeBinaryFromReader(message: PetitionPatchRequest, reader: jspb.BinaryReader): PetitionPatchRequest;
}

export namespace PetitionPatchRequest {
  export type AsObject = {
    id: number,
    key: string,
    value: string,
  }
}

export class PetitionPatchResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PetitionPatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PetitionPatchResponse): PetitionPatchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PetitionPatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PetitionPatchResponse;
  static deserializeBinaryFromReader(message: PetitionPatchResponse, reader: jspb.BinaryReader): PetitionPatchResponse;
}

export namespace PetitionPatchResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

