// package: pb
// file: proto/zone_service.proto

import * as jspb from "google-protobuf";

import * as proto_zone_pb from "../proto/zone_pb";

export class ZoneSearchRequest extends jspb.Message {
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
  toObject(includeInstance?: boolean): ZoneSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ZoneSearchRequest): ZoneSearchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ZoneSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZoneSearchRequest;
  static deserializeBinaryFromReader(message: ZoneSearchRequest, reader: jspb.BinaryReader): ZoneSearchRequest;
}

export namespace ZoneSearchRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
    limit: number,
    offset: number,
    orderby: string,
    orderdesc: boolean,
  }
}

export class ZoneSearchResponse extends jspb.Message {
  clearZonesList(): void;
  getZonesList(): Array<proto_zone_pb.Zone>;
  setZonesList(value: Array<proto_zone_pb.Zone>): void;
  addZones(value?: proto_zone_pb.Zone, index?: number): proto_zone_pb.Zone;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ZoneSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ZoneSearchResponse): ZoneSearchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ZoneSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZoneSearchResponse;
  static deserializeBinaryFromReader(message: ZoneSearchResponse, reader: jspb.BinaryReader): ZoneSearchResponse;
}

export namespace ZoneSearchResponse {
  export type AsObject = {
    zonesList: Array<proto_zone_pb.Zone.AsObject>,
    total: number,
  }
}

export class ZoneCreateRequest extends jspb.Message {
  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ZoneCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ZoneCreateRequest): ZoneCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ZoneCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZoneCreateRequest;
  static deserializeBinaryFromReader(message: ZoneCreateRequest, reader: jspb.BinaryReader): ZoneCreateRequest;
}

export namespace ZoneCreateRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
  }
}

export class ZoneCreateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ZoneCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ZoneCreateResponse): ZoneCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ZoneCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZoneCreateResponse;
  static deserializeBinaryFromReader(message: ZoneCreateResponse, reader: jspb.BinaryReader): ZoneCreateResponse;
}

export namespace ZoneCreateResponse {
  export type AsObject = {
    id: number,
  }
}

export class ZoneReadRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ZoneReadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ZoneReadRequest): ZoneReadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ZoneReadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZoneReadRequest;
  static deserializeBinaryFromReader(message: ZoneReadRequest, reader: jspb.BinaryReader): ZoneReadRequest;
}

export namespace ZoneReadRequest {
  export type AsObject = {
    id: number,
  }
}

export class ZoneReadResponse extends jspb.Message {
  hasZone(): boolean;
  clearZone(): void;
  getZone(): proto_zone_pb.Zone | undefined;
  setZone(value?: proto_zone_pb.Zone): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ZoneReadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ZoneReadResponse): ZoneReadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ZoneReadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZoneReadResponse;
  static deserializeBinaryFromReader(message: ZoneReadResponse, reader: jspb.BinaryReader): ZoneReadResponse;
}

export namespace ZoneReadResponse {
  export type AsObject = {
    zone?: proto_zone_pb.Zone.AsObject,
  }
}

export class ZoneUpdateRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ZoneUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ZoneUpdateRequest): ZoneUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ZoneUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZoneUpdateRequest;
  static deserializeBinaryFromReader(message: ZoneUpdateRequest, reader: jspb.BinaryReader): ZoneUpdateRequest;
}

export namespace ZoneUpdateRequest {
  export type AsObject = {
    id: number,
    valuesMap: Array<[string, string]>,
  }
}

export class ZoneUpdateResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ZoneUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ZoneUpdateResponse): ZoneUpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ZoneUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZoneUpdateResponse;
  static deserializeBinaryFromReader(message: ZoneUpdateResponse, reader: jspb.BinaryReader): ZoneUpdateResponse;
}

export namespace ZoneUpdateResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class ZoneDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ZoneDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ZoneDeleteRequest): ZoneDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ZoneDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZoneDeleteRequest;
  static deserializeBinaryFromReader(message: ZoneDeleteRequest, reader: jspb.BinaryReader): ZoneDeleteRequest;
}

export namespace ZoneDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class ZoneDeleteResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ZoneDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ZoneDeleteResponse): ZoneDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ZoneDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZoneDeleteResponse;
  static deserializeBinaryFromReader(message: ZoneDeleteResponse, reader: jspb.BinaryReader): ZoneDeleteResponse;
}

export namespace ZoneDeleteResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class ZonePatchRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ZonePatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ZonePatchRequest): ZonePatchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ZonePatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZonePatchRequest;
  static deserializeBinaryFromReader(message: ZonePatchRequest, reader: jspb.BinaryReader): ZonePatchRequest;
}

export namespace ZonePatchRequest {
  export type AsObject = {
    id: number,
    key: string,
    value: string,
  }
}

export class ZonePatchResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ZonePatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ZonePatchResponse): ZonePatchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ZonePatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZonePatchResponse;
  static deserializeBinaryFromReader(message: ZonePatchResponse, reader: jspb.BinaryReader): ZonePatchResponse;
}

export namespace ZonePatchResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

