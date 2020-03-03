// package: pb
// file: proto/item_service.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../google/api/annotations_pb";
import * as proto_item_pb from "../proto/item_pb";

export class ItemSearchRequest extends jspb.Message {
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
  toObject(includeInstance?: boolean): ItemSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ItemSearchRequest): ItemSearchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemSearchRequest;
  static deserializeBinaryFromReader(message: ItemSearchRequest, reader: jspb.BinaryReader): ItemSearchRequest;
}

export namespace ItemSearchRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
    limit: number,
    offset: number,
    orderby: string,
    orderdesc: boolean,
  }
}

export class ItemSearchResponse extends jspb.Message {
  clearItemsList(): void;
  getItemsList(): Array<proto_item_pb.Item>;
  setItemsList(value: Array<proto_item_pb.Item>): void;
  addItems(value?: proto_item_pb.Item, index?: number): proto_item_pb.Item;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ItemSearchResponse): ItemSearchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemSearchResponse;
  static deserializeBinaryFromReader(message: ItemSearchResponse, reader: jspb.BinaryReader): ItemSearchResponse;
}

export namespace ItemSearchResponse {
  export type AsObject = {
    itemsList: Array<proto_item_pb.Item.AsObject>,
    total: number,
  }
}

export class ItemCreateRequest extends jspb.Message {
  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ItemCreateRequest): ItemCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemCreateRequest;
  static deserializeBinaryFromReader(message: ItemCreateRequest, reader: jspb.BinaryReader): ItemCreateRequest;
}

export namespace ItemCreateRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
  }
}

export class ItemCreateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ItemCreateResponse): ItemCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemCreateResponse;
  static deserializeBinaryFromReader(message: ItemCreateResponse, reader: jspb.BinaryReader): ItemCreateResponse;
}

export namespace ItemCreateResponse {
  export type AsObject = {
    id: number,
  }
}

export class ItemReadRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemReadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ItemReadRequest): ItemReadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemReadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemReadRequest;
  static deserializeBinaryFromReader(message: ItemReadRequest, reader: jspb.BinaryReader): ItemReadRequest;
}

export namespace ItemReadRequest {
  export type AsObject = {
    id: number,
  }
}

export class ItemReadResponse extends jspb.Message {
  hasItem(): boolean;
  clearItem(): void;
  getItem(): proto_item_pb.Item | undefined;
  setItem(value?: proto_item_pb.Item): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemReadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ItemReadResponse): ItemReadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemReadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemReadResponse;
  static deserializeBinaryFromReader(message: ItemReadResponse, reader: jspb.BinaryReader): ItemReadResponse;
}

export namespace ItemReadResponse {
  export type AsObject = {
    item?: proto_item_pb.Item.AsObject,
  }
}

export class ItemUpdateRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ItemUpdateRequest): ItemUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemUpdateRequest;
  static deserializeBinaryFromReader(message: ItemUpdateRequest, reader: jspb.BinaryReader): ItemUpdateRequest;
}

export namespace ItemUpdateRequest {
  export type AsObject = {
    id: number,
    valuesMap: Array<[string, string]>,
  }
}

export class ItemUpdateResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ItemUpdateResponse): ItemUpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemUpdateResponse;
  static deserializeBinaryFromReader(message: ItemUpdateResponse, reader: jspb.BinaryReader): ItemUpdateResponse;
}

export namespace ItemUpdateResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class ItemDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ItemDeleteRequest): ItemDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemDeleteRequest;
  static deserializeBinaryFromReader(message: ItemDeleteRequest, reader: jspb.BinaryReader): ItemDeleteRequest;
}

export namespace ItemDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class ItemDeleteResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ItemDeleteResponse): ItemDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemDeleteResponse;
  static deserializeBinaryFromReader(message: ItemDeleteResponse, reader: jspb.BinaryReader): ItemDeleteResponse;
}

export namespace ItemDeleteResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class ItemPatchRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemPatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ItemPatchRequest): ItemPatchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemPatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemPatchRequest;
  static deserializeBinaryFromReader(message: ItemPatchRequest, reader: jspb.BinaryReader): ItemPatchRequest;
}

export namespace ItemPatchRequest {
  export type AsObject = {
    id: number,
    key: string,
    value: string,
  }
}

export class ItemPatchResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ItemPatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ItemPatchResponse): ItemPatchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ItemPatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ItemPatchResponse;
  static deserializeBinaryFromReader(message: ItemPatchResponse, reader: jspb.BinaryReader): ItemPatchResponse;
}

export namespace ItemPatchResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

