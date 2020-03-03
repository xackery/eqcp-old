// package: pb
// file: proto/petition.proto

import * as jspb from "google-protobuf";

export class Petition extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getPetid(): number;
  setPetid(value: number): void;

  getCharname(): string;
  setCharname(value: string): void;

  getAccountname(): string;
  setAccountname(value: string): void;

  getLastgm(): string;
  setLastgm(value: string): void;

  getPetitiontext(): string;
  setPetitiontext(value: string): void;

  getGmtext(): string;
  setGmtext(value: string): void;

  getZone(): string;
  setZone(value: string): void;

  getUrgency(): number;
  setUrgency(value: number): void;

  getCharclass(): number;
  setCharclass(value: number): void;

  getCharrace(): number;
  setCharrace(value: number): void;

  getCharlevel(): number;
  setCharlevel(value: number): void;

  getCheckouts(): number;
  setCheckouts(value: number): void;

  getUnavailables(): number;
  setUnavailables(value: number): void;

  getIscheckedout(): number;
  setIscheckedout(value: number): void;

  getSenttime(): number;
  setSenttime(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Petition.AsObject;
  static toObject(includeInstance: boolean, msg: Petition): Petition.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Petition, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Petition;
  static deserializeBinaryFromReader(message: Petition, reader: jspb.BinaryReader): Petition;
}

export namespace Petition {
  export type AsObject = {
    id: number,
    petid: number,
    charname: string,
    accountname: string,
    lastgm: string,
    petitiontext: string,
    gmtext: string,
    zone: string,
    urgency: number,
    charclass: number,
    charrace: number,
    charlevel: number,
    checkouts: number,
    unavailables: number,
    ischeckedout: number,
    senttime: number,
  }
}

