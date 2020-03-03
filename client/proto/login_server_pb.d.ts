// package: pb
// file: proto/login_server.proto

import * as jspb from "google-protobuf";

export class Server extends jspb.Message {
  getLocalip(): string;
  setLocalip(value: string): void;

  getPlayersonline(): number;
  setPlayersonline(value: number): void;

  getRemoteip(): string;
  setRemoteip(value: string): void;

  getServerlistid(): number;
  setServerlistid(value: number): void;

  getServerlongname(): string;
  setServerlongname(value: string): void;

  getServershortname(): string;
  setServershortname(value: string): void;

  getServerstatus(): number;
  setServerstatus(value: number): void;

  getZonesbooted(): number;
  setZonesbooted(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Server.AsObject;
  static toObject(includeInstance: boolean, msg: Server): Server.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Server, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Server;
  static deserializeBinaryFromReader(message: Server, reader: jspb.BinaryReader): Server;
}

export namespace Server {
  export type AsObject = {
    localip: string,
    playersonline: number,
    remoteip: string,
    serverlistid: number,
    serverlongname: string,
    servershortname: string,
    serverstatus: number,
    zonesbooted: number,
  }
}

