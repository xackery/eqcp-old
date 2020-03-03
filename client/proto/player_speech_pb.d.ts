// package: pb
// file: proto/player_speech.proto

import * as jspb from "google-protobuf";

export class PlayerSpeech extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getFrom(): string;
  setFrom(value: string): void;

  getTo(): string;
  setTo(value: string): void;

  getFrommessage(): string;
  setFrommessage(value: string): void;

  getMinstatus(): number;
  setMinstatus(value: number): void;

  getGuilddbid(): number;
  setGuilddbid(value: number): void;

  getType(): number;
  setType(value: number): void;

  getTimerecorded(): number;
  setTimerecorded(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerSpeech.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerSpeech): PlayerSpeech.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerSpeech, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerSpeech;
  static deserializeBinaryFromReader(message: PlayerSpeech, reader: jspb.BinaryReader): PlayerSpeech;
}

export namespace PlayerSpeech {
  export type AsObject = {
    id: number,
    from: string,
    to: string,
    frommessage: string,
    minstatus: number,
    guilddbid: number,
    type: number,
    timerecorded: number,
  }
}

