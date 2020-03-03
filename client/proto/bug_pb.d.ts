// package: pb
// file: proto/bug.proto

import * as jspb from "google-protobuf";

export class Bug extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getZone(): string;
  setZone(value: string): void;

  getClientversionid(): number;
  setClientversionid(value: number): void;

  getClientversionname(): string;
  setClientversionname(value: string): void;

  getAccountid(): number;
  setAccountid(value: number): void;

  getCharacterid(): number;
  setCharacterid(value: number): void;

  getCharactername(): string;
  setCharactername(value: string): void;

  getReporterspoof(): number;
  setReporterspoof(value: number): void;

  getCategoryid(): number;
  setCategoryid(value: number): void;

  getCategoryname(): string;
  setCategoryname(value: string): void;

  getReportername(): string;
  setReportername(value: string): void;

  getUipath(): string;
  setUipath(value: string): void;

  getPosx(): number;
  setPosx(value: number): void;

  getPosy(): number;
  setPosy(value: number): void;

  getPosz(): number;
  setPosz(value: number): void;

  getHeading(): number;
  setHeading(value: number): void;

  getTimeplayed(): number;
  setTimeplayed(value: number): void;

  getTargetid(): number;
  setTargetid(value: number): void;

  getTargetname(): string;
  setTargetname(value: string): void;

  getOptionalinfomask(): number;
  setOptionalinfomask(value: number): void;

  getCanduplicate(): number;
  setCanduplicate(value: number): void;

  getCrashbug(): number;
  setCrashbug(value: number): void;

  getTargetinfo(): number;
  setTargetinfo(value: number): void;

  getCharacterflags(): number;
  setCharacterflags(value: number): void;

  getUnknownvalue(): number;
  setUnknownvalue(value: number): void;

  getBugreport(): string;
  setBugreport(value: string): void;

  getSysteminfo(): string;
  setSysteminfo(value: string): void;

  getReportdatetime(): number;
  setReportdatetime(value: number): void;

  getBugstatus(): number;
  setBugstatus(value: number): void;

  getLastreview(): number;
  setLastreview(value: number): void;

  getLastreviewer(): string;
  setLastreviewer(value: string): void;

  getReviewernotes(): string;
  setReviewernotes(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Bug.AsObject;
  static toObject(includeInstance: boolean, msg: Bug): Bug.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Bug, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Bug;
  static deserializeBinaryFromReader(message: Bug, reader: jspb.BinaryReader): Bug;
}

export namespace Bug {
  export type AsObject = {
    id: number,
    zone: string,
    clientversionid: number,
    clientversionname: string,
    accountid: number,
    characterid: number,
    charactername: string,
    reporterspoof: number,
    categoryid: number,
    categoryname: string,
    reportername: string,
    uipath: string,
    posx: number,
    posy: number,
    posz: number,
    heading: number,
    timeplayed: number,
    targetid: number,
    targetname: string,
    optionalinfomask: number,
    canduplicate: number,
    crashbug: number,
    targetinfo: number,
    characterflags: number,
    unknownvalue: number,
    bugreport: string,
    systeminfo: string,
    reportdatetime: number,
    bugstatus: number,
    lastreview: number,
    lastreviewer: string,
    reviewernotes: string,
  }
}

