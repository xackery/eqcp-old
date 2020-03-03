// package: pb
// file: proto/login_server_service.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../google/api/annotations_pb";
import * as proto_login_server_pb from "../proto/login_server_pb";

export class LoginServerListRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginServerListRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginServerListRequest): LoginServerListRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginServerListRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginServerListRequest;
  static deserializeBinaryFromReader(message: LoginServerListRequest, reader: jspb.BinaryReader): LoginServerListRequest;
}

export namespace LoginServerListRequest {
  export type AsObject = {
  }
}

export class LoginServerListResponse extends jspb.Message {
  clearServersList(): void;
  getServersList(): Array<proto_login_server_pb.Server>;
  setServersList(value: Array<proto_login_server_pb.Server>): void;
  addServers(value?: proto_login_server_pb.Server, index?: number): proto_login_server_pb.Server;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginServerListResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginServerListResponse): LoginServerListResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginServerListResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginServerListResponse;
  static deserializeBinaryFromReader(message: LoginServerListResponse, reader: jspb.BinaryReader): LoginServerListResponse;
}

export namespace LoginServerListResponse {
  export type AsObject = {
    serversList: Array<proto_login_server_pb.Server.AsObject>,
  }
}

export class LoginServerLoginRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginServerLoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginServerLoginRequest): LoginServerLoginRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginServerLoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginServerLoginRequest;
  static deserializeBinaryFromReader(message: LoginServerLoginRequest, reader: jspb.BinaryReader): LoginServerLoginRequest;
}

export namespace LoginServerLoginRequest {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class LoginServerLoginResponse extends jspb.Message {
  getToken(): string;
  setToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginServerLoginResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginServerLoginResponse): LoginServerLoginResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginServerLoginResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginServerLoginResponse;
  static deserializeBinaryFromReader(message: LoginServerLoginResponse, reader: jspb.BinaryReader): LoginServerLoginResponse;
}

export namespace LoginServerLoginResponse {
  export type AsObject = {
    token: string,
  }
}

export class LoginServerLogoutRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginServerLogoutRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginServerLogoutRequest): LoginServerLogoutRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginServerLogoutRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginServerLogoutRequest;
  static deserializeBinaryFromReader(message: LoginServerLogoutRequest, reader: jspb.BinaryReader): LoginServerLogoutRequest;
}

export namespace LoginServerLogoutRequest {
  export type AsObject = {
  }
}

export class LoginServerLogoutResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginServerLogoutResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginServerLogoutResponse): LoginServerLogoutResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginServerLogoutResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginServerLogoutResponse;
  static deserializeBinaryFromReader(message: LoginServerLogoutResponse, reader: jspb.BinaryReader): LoginServerLogoutResponse;
}

export namespace LoginServerLogoutResponse {
  export type AsObject = {
  }
}

export class LoginServerCreateRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  getEmail(): string;
  setEmail(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginServerCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginServerCreateRequest): LoginServerCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginServerCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginServerCreateRequest;
  static deserializeBinaryFromReader(message: LoginServerCreateRequest, reader: jspb.BinaryReader): LoginServerCreateRequest;
}

export namespace LoginServerCreateRequest {
  export type AsObject = {
    username: string,
    password: string,
    email: string,
  }
}

export class LoginServerCreateResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginServerCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginServerCreateResponse): LoginServerCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginServerCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginServerCreateResponse;
  static deserializeBinaryFromReader(message: LoginServerCreateResponse, reader: jspb.BinaryReader): LoginServerCreateResponse;
}

export namespace LoginServerCreateResponse {
  export type AsObject = {
  }
}

