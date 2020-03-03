// package: pb
// file: proto/inventory_service.proto

import * as jspb from "google-protobuf";

import * as proto_inventory_pb from "../proto/inventory_pb";

export class InventorySearchRequest extends jspb.Message {
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
  toObject(includeInstance?: boolean): InventorySearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: InventorySearchRequest): InventorySearchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: InventorySearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InventorySearchRequest;
  static deserializeBinaryFromReader(message: InventorySearchRequest, reader: jspb.BinaryReader): InventorySearchRequest;
}

export namespace InventorySearchRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
    limit: number,
    offset: number,
    orderby: string,
    orderdesc: boolean,
  }
}

export class InventorySearchResponse extends jspb.Message {
  clearInventorysList(): void;
  getInventorysList(): Array<proto_inventory_pb.Inventory>;
  setInventorysList(value: Array<proto_inventory_pb.Inventory>): void;
  addInventorys(value?: proto_inventory_pb.Inventory, index?: number): proto_inventory_pb.Inventory;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InventorySearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: InventorySearchResponse): InventorySearchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: InventorySearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InventorySearchResponse;
  static deserializeBinaryFromReader(message: InventorySearchResponse, reader: jspb.BinaryReader): InventorySearchResponse;
}

export namespace InventorySearchResponse {
  export type AsObject = {
    inventorysList: Array<proto_inventory_pb.Inventory.AsObject>,
    total: number,
  }
}

export class InventoryCreateRequest extends jspb.Message {
  getCharid(): number;
  setCharid(value: number): void;

  getSlotid(): number;
  setSlotid(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InventoryCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: InventoryCreateRequest): InventoryCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: InventoryCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InventoryCreateRequest;
  static deserializeBinaryFromReader(message: InventoryCreateRequest, reader: jspb.BinaryReader): InventoryCreateRequest;
}

export namespace InventoryCreateRequest {
  export type AsObject = {
    charid: number,
    slotid: number,
    valuesMap: Array<[string, string]>,
  }
}

export class InventoryCreateResponse extends jspb.Message {
  getCharid(): number;
  setCharid(value: number): void;

  getSlotid(): number;
  setSlotid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InventoryCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: InventoryCreateResponse): InventoryCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: InventoryCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InventoryCreateResponse;
  static deserializeBinaryFromReader(message: InventoryCreateResponse, reader: jspb.BinaryReader): InventoryCreateResponse;
}

export namespace InventoryCreateResponse {
  export type AsObject = {
    charid: number,
    slotid: number,
  }
}

export class InventoryReadRequest extends jspb.Message {
  getCharid(): number;
  setCharid(value: number): void;

  getSlotid(): number;
  setSlotid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InventoryReadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: InventoryReadRequest): InventoryReadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: InventoryReadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InventoryReadRequest;
  static deserializeBinaryFromReader(message: InventoryReadRequest, reader: jspb.BinaryReader): InventoryReadRequest;
}

export namespace InventoryReadRequest {
  export type AsObject = {
    charid: number,
    slotid: number,
  }
}

export class InventoryReadResponse extends jspb.Message {
  hasInventory(): boolean;
  clearInventory(): void;
  getInventory(): proto_inventory_pb.Inventory | undefined;
  setInventory(value?: proto_inventory_pb.Inventory): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InventoryReadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: InventoryReadResponse): InventoryReadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: InventoryReadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InventoryReadResponse;
  static deserializeBinaryFromReader(message: InventoryReadResponse, reader: jspb.BinaryReader): InventoryReadResponse;
}

export namespace InventoryReadResponse {
  export type AsObject = {
    inventory?: proto_inventory_pb.Inventory.AsObject,
  }
}

export class InventoryUpdateRequest extends jspb.Message {
  getCharid(): number;
  setCharid(value: number): void;

  getSlotid(): number;
  setSlotid(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InventoryUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: InventoryUpdateRequest): InventoryUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: InventoryUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InventoryUpdateRequest;
  static deserializeBinaryFromReader(message: InventoryUpdateRequest, reader: jspb.BinaryReader): InventoryUpdateRequest;
}

export namespace InventoryUpdateRequest {
  export type AsObject = {
    charid: number,
    slotid: number,
    valuesMap: Array<[string, string]>,
  }
}

export class InventoryUpdateResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InventoryUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: InventoryUpdateResponse): InventoryUpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: InventoryUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InventoryUpdateResponse;
  static deserializeBinaryFromReader(message: InventoryUpdateResponse, reader: jspb.BinaryReader): InventoryUpdateResponse;
}

export namespace InventoryUpdateResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class InventoryDeleteRequest extends jspb.Message {
  getCharid(): number;
  setCharid(value: number): void;

  getSlotid(): number;
  setSlotid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InventoryDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: InventoryDeleteRequest): InventoryDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: InventoryDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InventoryDeleteRequest;
  static deserializeBinaryFromReader(message: InventoryDeleteRequest, reader: jspb.BinaryReader): InventoryDeleteRequest;
}

export namespace InventoryDeleteRequest {
  export type AsObject = {
    charid: number,
    slotid: number,
  }
}

export class InventoryDeleteResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InventoryDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: InventoryDeleteResponse): InventoryDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: InventoryDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InventoryDeleteResponse;
  static deserializeBinaryFromReader(message: InventoryDeleteResponse, reader: jspb.BinaryReader): InventoryDeleteResponse;
}

export namespace InventoryDeleteResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class InventoryPatchRequest extends jspb.Message {
  getCharid(): number;
  setCharid(value: number): void;

  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  getSlotid(): number;
  setSlotid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InventoryPatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: InventoryPatchRequest): InventoryPatchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: InventoryPatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InventoryPatchRequest;
  static deserializeBinaryFromReader(message: InventoryPatchRequest, reader: jspb.BinaryReader): InventoryPatchRequest;
}

export namespace InventoryPatchRequest {
  export type AsObject = {
    charid: number,
    key: string,
    value: string,
    slotid: number,
  }
}

export class InventoryPatchResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InventoryPatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: InventoryPatchResponse): InventoryPatchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: InventoryPatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InventoryPatchResponse;
  static deserializeBinaryFromReader(message: InventoryPatchResponse, reader: jspb.BinaryReader): InventoryPatchResponse;
}

export namespace InventoryPatchResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

