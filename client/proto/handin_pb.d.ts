// package: pb
// file: proto/handin.proto

import * as jspb from "google-protobuf";

export class Handin extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getTime(): number;
  setTime(value: number): void;

  getQuestid(): number;
  setQuestid(value: number): void;

  getCharid(): number;
  setCharid(value: number): void;

  getCharpp(): number;
  setCharpp(value: number): void;

  getChargp(): number;
  setChargp(value: number): void;

  getCharsp(): number;
  setCharsp(value: number): void;

  getCharcp(): number;
  setCharcp(value: number): void;

  getCharitems(): number;
  setCharitems(value: number): void;

  getNpcid(): number;
  setNpcid(value: number): void;

  getNpcpp(): number;
  setNpcpp(value: number): void;

  getNpcgp(): number;
  setNpcgp(value: number): void;

  getNpcsp(): number;
  setNpcsp(value: number): void;

  getNpccp(): number;
  setNpccp(value: number): void;

  getNpcitems(): number;
  setNpcitems(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Handin.AsObject;
  static toObject(includeInstance: boolean, msg: Handin): Handin.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Handin, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Handin;
  static deserializeBinaryFromReader(message: Handin, reader: jspb.BinaryReader): Handin;
}

export namespace Handin {
  export type AsObject = {
    id: number,
    time: number,
    questid: number,
    charid: number,
    charpp: number,
    chargp: number,
    charsp: number,
    charcp: number,
    charitems: number,
    npcid: number,
    npcpp: number,
    npcgp: number,
    npcsp: number,
    npccp: number,
    npcitems: number,
  }
}

