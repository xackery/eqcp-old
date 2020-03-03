// package: pb
// file: proto/trade_service.proto

import * as jspb from "google-protobuf";

import * as proto_trade_pb from "../proto/trade_pb";

export class TradeSearchRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getLimit(): number;
  setLimit(value: number): void;

  getOffset(): number;
  setOffset(value: number): void;

  getOrderby(): string;
  setOrderby(value: string): void;

  getOrderdesc(): boolean;
  setOrderdesc(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TradeSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TradeSearchRequest): TradeSearchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TradeSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TradeSearchRequest;
  static deserializeBinaryFromReader(message: TradeSearchRequest, reader: jspb.BinaryReader): TradeSearchRequest;
}

export namespace TradeSearchRequest {
  export type AsObject = {
    name: string,
    limit: number,
    offset: number,
    orderby: string,
    orderdesc: boolean,
  }
}

export class TradeSearchResponse extends jspb.Message {
  clearTradesList(): void;
  getTradesList(): Array<proto_trade_pb.Trade>;
  setTradesList(value: Array<proto_trade_pb.Trade>): void;
  addTrades(value?: proto_trade_pb.Trade, index?: number): proto_trade_pb.Trade;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TradeSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TradeSearchResponse): TradeSearchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TradeSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TradeSearchResponse;
  static deserializeBinaryFromReader(message: TradeSearchResponse, reader: jspb.BinaryReader): TradeSearchResponse;
}

export namespace TradeSearchResponse {
  export type AsObject = {
    tradesList: Array<proto_trade_pb.Trade.AsObject>,
    total: number,
  }
}

export class TradeCreateRequest extends jspb.Message {
  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TradeCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TradeCreateRequest): TradeCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TradeCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TradeCreateRequest;
  static deserializeBinaryFromReader(message: TradeCreateRequest, reader: jspb.BinaryReader): TradeCreateRequest;
}

export namespace TradeCreateRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
  }
}

export class TradeCreateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TradeCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TradeCreateResponse): TradeCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TradeCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TradeCreateResponse;
  static deserializeBinaryFromReader(message: TradeCreateResponse, reader: jspb.BinaryReader): TradeCreateResponse;
}

export namespace TradeCreateResponse {
  export type AsObject = {
    id: number,
  }
}

export class TradeReadRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TradeReadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TradeReadRequest): TradeReadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TradeReadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TradeReadRequest;
  static deserializeBinaryFromReader(message: TradeReadRequest, reader: jspb.BinaryReader): TradeReadRequest;
}

export namespace TradeReadRequest {
  export type AsObject = {
    id: number,
  }
}

export class TradeReadResponse extends jspb.Message {
  hasTrade(): boolean;
  clearTrade(): void;
  getTrade(): proto_trade_pb.Trade | undefined;
  setTrade(value?: proto_trade_pb.Trade): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TradeReadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TradeReadResponse): TradeReadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TradeReadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TradeReadResponse;
  static deserializeBinaryFromReader(message: TradeReadResponse, reader: jspb.BinaryReader): TradeReadResponse;
}

export namespace TradeReadResponse {
  export type AsObject = {
    trade?: proto_trade_pb.Trade.AsObject,
  }
}

export class TradeUpdateRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TradeUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TradeUpdateRequest): TradeUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TradeUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TradeUpdateRequest;
  static deserializeBinaryFromReader(message: TradeUpdateRequest, reader: jspb.BinaryReader): TradeUpdateRequest;
}

export namespace TradeUpdateRequest {
  export type AsObject = {
    id: number,
    valuesMap: Array<[string, string]>,
  }
}

export class TradeUpdateResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TradeUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TradeUpdateResponse): TradeUpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TradeUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TradeUpdateResponse;
  static deserializeBinaryFromReader(message: TradeUpdateResponse, reader: jspb.BinaryReader): TradeUpdateResponse;
}

export namespace TradeUpdateResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class TradeDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TradeDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TradeDeleteRequest): TradeDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TradeDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TradeDeleteRequest;
  static deserializeBinaryFromReader(message: TradeDeleteRequest, reader: jspb.BinaryReader): TradeDeleteRequest;
}

export namespace TradeDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class TradeDeleteResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TradeDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TradeDeleteResponse): TradeDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TradeDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TradeDeleteResponse;
  static deserializeBinaryFromReader(message: TradeDeleteResponse, reader: jspb.BinaryReader): TradeDeleteResponse;
}

export namespace TradeDeleteResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class TradePatchRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TradePatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TradePatchRequest): TradePatchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TradePatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TradePatchRequest;
  static deserializeBinaryFromReader(message: TradePatchRequest, reader: jspb.BinaryReader): TradePatchRequest;
}

export namespace TradePatchRequest {
  export type AsObject = {
    id: number,
    key: string,
    value: string,
  }
}

export class TradePatchResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TradePatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TradePatchResponse): TradePatchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TradePatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TradePatchResponse;
  static deserializeBinaryFromReader(message: TradePatchResponse, reader: jspb.BinaryReader): TradePatchResponse;
}

export namespace TradePatchResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

