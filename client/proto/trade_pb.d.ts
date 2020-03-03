// package: pb
// file: proto/trade.proto

import * as jspb from "google-protobuf";

export class Trade extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getTime(): number;
  setTime(value: number): void;

  getChar1id(): number;
  setChar1id(value: number): void;

  getChar1pp(): number;
  setChar1pp(value: number): void;

  getChar1gp(): number;
  setChar1gp(value: number): void;

  getChar1sp(): number;
  setChar1sp(value: number): void;

  getChar1cp(): number;
  setChar1cp(value: number): void;

  getChar1items(): number;
  setChar1items(value: number): void;

  getChar2id(): number;
  setChar2id(value: number): void;

  getChar2pp(): number;
  setChar2pp(value: number): void;

  getChar2gp(): number;
  setChar2gp(value: number): void;

  getChar2sp(): number;
  setChar2sp(value: number): void;

  getChar2cp(): number;
  setChar2cp(value: number): void;

  getChar2items(): number;
  setChar2items(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Trade.AsObject;
  static toObject(includeInstance: boolean, msg: Trade): Trade.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Trade, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Trade;
  static deserializeBinaryFromReader(message: Trade, reader: jspb.BinaryReader): Trade;
}

export namespace Trade {
  export type AsObject = {
    id: number,
    time: number,
    char1id: number,
    char1pp: number,
    char1gp: number,
    char1sp: number,
    char1cp: number,
    char1items: number,
    char2id: number,
    char2pp: number,
    char2gp: number,
    char2sp: number,
    char2cp: number,
    char2items: number,
  }
}

