// package: pb
// file: proto/login_account_service.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../google/api/annotations_pb";
import * as proto_login_account_pb from "../proto/login_account_pb";

export class LoginAccountSearchRequest extends jspb.Message {
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
  toObject(includeInstance?: boolean): LoginAccountSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginAccountSearchRequest): LoginAccountSearchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginAccountSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginAccountSearchRequest;
  static deserializeBinaryFromReader(message: LoginAccountSearchRequest, reader: jspb.BinaryReader): LoginAccountSearchRequest;
}

export namespace LoginAccountSearchRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
    limit: number,
    offset: number,
    orderby: string,
    orderdesc: boolean,
  }
}

export class LoginAccountSearchResponse extends jspb.Message {
  clearLoginaccountsList(): void;
  getLoginaccountsList(): Array<proto_login_account_pb.LoginAccount>;
  setLoginaccountsList(value: Array<proto_login_account_pb.LoginAccount>): void;
  addLoginaccounts(value?: proto_login_account_pb.LoginAccount, index?: number): proto_login_account_pb.LoginAccount;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginAccountSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginAccountSearchResponse): LoginAccountSearchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginAccountSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginAccountSearchResponse;
  static deserializeBinaryFromReader(message: LoginAccountSearchResponse, reader: jspb.BinaryReader): LoginAccountSearchResponse;
}

export namespace LoginAccountSearchResponse {
  export type AsObject = {
    loginaccountsList: Array<proto_login_account_pb.LoginAccount.AsObject>,
    total: number,
  }
}

export class LoginAccountCreateRequest extends jspb.Message {
  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginAccountCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginAccountCreateRequest): LoginAccountCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginAccountCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginAccountCreateRequest;
  static deserializeBinaryFromReader(message: LoginAccountCreateRequest, reader: jspb.BinaryReader): LoginAccountCreateRequest;
}

export namespace LoginAccountCreateRequest {
  export type AsObject = {
    valuesMap: Array<[string, string]>,
  }
}

export class LoginAccountCreateResponse extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginAccountCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginAccountCreateResponse): LoginAccountCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginAccountCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginAccountCreateResponse;
  static deserializeBinaryFromReader(message: LoginAccountCreateResponse, reader: jspb.BinaryReader): LoginAccountCreateResponse;
}

export namespace LoginAccountCreateResponse {
  export type AsObject = {
    id: number,
  }
}

export class LoginAccountReadRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginAccountReadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginAccountReadRequest): LoginAccountReadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginAccountReadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginAccountReadRequest;
  static deserializeBinaryFromReader(message: LoginAccountReadRequest, reader: jspb.BinaryReader): LoginAccountReadRequest;
}

export namespace LoginAccountReadRequest {
  export type AsObject = {
    id: number,
  }
}

export class LoginAccountReadResponse extends jspb.Message {
  hasLoginaccount(): boolean;
  clearLoginaccount(): void;
  getLoginaccount(): proto_login_account_pb.LoginAccount | undefined;
  setLoginaccount(value?: proto_login_account_pb.LoginAccount): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginAccountReadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginAccountReadResponse): LoginAccountReadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginAccountReadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginAccountReadResponse;
  static deserializeBinaryFromReader(message: LoginAccountReadResponse, reader: jspb.BinaryReader): LoginAccountReadResponse;
}

export namespace LoginAccountReadResponse {
  export type AsObject = {
    loginaccount?: proto_login_account_pb.LoginAccount.AsObject,
  }
}

export class LoginAccountUpdateRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getValuesMap(): jspb.Map<string, string>;
  clearValuesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginAccountUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginAccountUpdateRequest): LoginAccountUpdateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginAccountUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginAccountUpdateRequest;
  static deserializeBinaryFromReader(message: LoginAccountUpdateRequest, reader: jspb.BinaryReader): LoginAccountUpdateRequest;
}

export namespace LoginAccountUpdateRequest {
  export type AsObject = {
    id: number,
    valuesMap: Array<[string, string]>,
  }
}

export class LoginAccountUpdateResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginAccountUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginAccountUpdateResponse): LoginAccountUpdateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginAccountUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginAccountUpdateResponse;
  static deserializeBinaryFromReader(message: LoginAccountUpdateResponse, reader: jspb.BinaryReader): LoginAccountUpdateResponse;
}

export namespace LoginAccountUpdateResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class LoginAccountDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginAccountDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginAccountDeleteRequest): LoginAccountDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginAccountDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginAccountDeleteRequest;
  static deserializeBinaryFromReader(message: LoginAccountDeleteRequest, reader: jspb.BinaryReader): LoginAccountDeleteRequest;
}

export namespace LoginAccountDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class LoginAccountDeleteResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginAccountDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginAccountDeleteResponse): LoginAccountDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginAccountDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginAccountDeleteResponse;
  static deserializeBinaryFromReader(message: LoginAccountDeleteResponse, reader: jspb.BinaryReader): LoginAccountDeleteResponse;
}

export namespace LoginAccountDeleteResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

export class LoginAccountPatchRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginAccountPatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginAccountPatchRequest): LoginAccountPatchRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginAccountPatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginAccountPatchRequest;
  static deserializeBinaryFromReader(message: LoginAccountPatchRequest, reader: jspb.BinaryReader): LoginAccountPatchRequest;
}

export namespace LoginAccountPatchRequest {
  export type AsObject = {
    id: number,
    key: string,
    value: string,
  }
}

export class LoginAccountPatchResponse extends jspb.Message {
  getRowsaffected(): number;
  setRowsaffected(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginAccountPatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginAccountPatchResponse): LoginAccountPatchResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginAccountPatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginAccountPatchResponse;
  static deserializeBinaryFromReader(message: LoginAccountPatchResponse, reader: jspb.BinaryReader): LoginAccountPatchResponse;
}

export namespace LoginAccountPatchResponse {
  export type AsObject = {
    rowsaffected: number,
  }
}

