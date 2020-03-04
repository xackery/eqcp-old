// package: pb
// file: proto/account_service.proto

import * as jspb from "google-protobuf";


import * as proto_account_pb from "../proto/account_pb";

export class AccountSearchRequest extends jspb.Message {
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
  toObject(includeInstance?: boolean): AccountSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AccountSearchRequest): AccountSearchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountSearchRequest;
  static deserializeBinaryFromReader(message: AccountSearchRequest, reader: jspb.BinaryReader): AccountSearchRequest;
}

export namespace AccountSearchRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
    limit: number,
    offset: number,
    orderby: string,
    orderdesc: boolean,
  }
}

export class AccountSearchResponse extends jspb.Message {
  clearAccountsList(): void;
  getAccountsList(): Array<proto_account_pb.Account>;
  setAccountsList(value: Array<proto_account_pb.Account>): void;
  addAccounts(value?: proto_account_pb.Account, index?: number): proto_account_pb.Account;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccountSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AccountSearchResponse): AccountSearchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountSearchResponse;
  static deserializeBinaryFromReader(message: AccountSearchResponse, reader: jspb.BinaryReader): AccountSearchResponse;
}

export namespace AccountSearchResponse {
  export type AsObject = {
    accountsList: Array<proto_account_pb.Account.AsObject>,
    total: number,
  }
}

export class AccountCreateRequest extends jspb.Message {
  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccountCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AccountCreateRequest): AccountCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountCreateRequest;
  static deserializeBinaryFromReader(message: AccountCreateRequest, reader: jspb.BinaryReader): AccountCreateRequest;
}

export namespace AccountCreateRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
  }
}

export class AccountCreateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccountCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AccountCreateResponse): AccountCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountCreateResponse;
  static deserializeBinaryFromReader(message: AccountCreateResponse, reader: jspb.BinaryReader): AccountCreateResponse;
}

export namespace AccountCreateResponse {
  export type AsObject = {
    id: number,
  }
}

export class AccountReadRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccountReadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AccountReadRequest): AccountReadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountReadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountReadRequest;
  static deserializeBinaryFromReader(message: AccountReadRequest, reader: jspb.BinaryReader): AccountReadRequest;
}

export namespace AccountReadRequest {
  export type AsObject = {
    id: number,
  }
}

export class AccountReadResponse extends jspb.Message {
  hasAccount(): boolean;
  clearAccount(): void;
  getAccount(): proto_account_pb.Account | undefined;
  setAccount(value?: proto_account_pb.Account): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccountReadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AccountReadResponse): AccountReadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountReadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountReadResponse;
  static deserializeBinaryFromReader(message: AccountReadResponse, reader: jspb.BinaryReader): AccountReadResponse;
}

export namespace AccountReadResponse {
  export type AsObject = {
    account?: proto_account_pb.Account.AsObject,
  }
}

export class AccountUpdateRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccountUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AccountUpdateRequest): AccountUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountUpdateRequest;
  static deserializeBinaryFromReader(message: AccountUpdateRequest, reader: jspb.BinaryReader): AccountUpdateRequest;
}

export namespace AccountUpdateRequest {
  export type AsObject = {
    id: number,
    valuesMap: Array<[string, string]>,
  }
}

export class AccountUpdateResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccountUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AccountUpdateResponse): AccountUpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountUpdateResponse;
  static deserializeBinaryFromReader(message: AccountUpdateResponse, reader: jspb.BinaryReader): AccountUpdateResponse;
}

export namespace AccountUpdateResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class AccountDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccountDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AccountDeleteRequest): AccountDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountDeleteRequest;
  static deserializeBinaryFromReader(message: AccountDeleteRequest, reader: jspb.BinaryReader): AccountDeleteRequest;
}

export namespace AccountDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class AccountDeleteResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccountDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AccountDeleteResponse): AccountDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountDeleteResponse;
  static deserializeBinaryFromReader(message: AccountDeleteResponse, reader: jspb.BinaryReader): AccountDeleteResponse;
}

export namespace AccountDeleteResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class AccountPatchRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccountPatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AccountPatchRequest): AccountPatchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountPatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountPatchRequest;
  static deserializeBinaryFromReader(message: AccountPatchRequest, reader: jspb.BinaryReader): AccountPatchRequest;
}

export namespace AccountPatchRequest {
  export type AsObject = {
    id: number,
    key: string,
    value: string,
  }
}

export class AccountPatchResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccountPatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AccountPatchResponse): AccountPatchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountPatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountPatchResponse;
  static deserializeBinaryFromReader(message: AccountPatchResponse, reader: jspb.BinaryReader): AccountPatchResponse;
}

export namespace AccountPatchResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

