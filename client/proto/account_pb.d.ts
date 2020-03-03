// package: pb
// file: proto/account.proto

import * as jspb from "google-protobuf";

export class Account extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getName(): string;
  setName(value: string): void;

  getCharname(): string;
  setCharname(value: string): void;

  getSharedplat(): number;
  setSharedplat(value: number): void;

  getPassword(): string;
  setPassword(value: string): void;

  getStatus(): number;
  setStatus(value: number): void;

  getLsid(): string;
  setLsid(value: string): void;

  getLsaccountid(): number;
  setLsaccountid(value: number): void;

  getGmspeed(): number;
  setGmspeed(value: number): void;

  getRevoked(): number;
  setRevoked(value: number): void;

  getKarma(): number;
  setKarma(value: number): void;

  getMiniloginip(): string;
  setMiniloginip(value: string): void;

  getHideme(): number;
  setHideme(value: number): void;

  getRulesflag(): number;
  setRulesflag(value: number): void;

  getSuspendeduntil(): number;
  setSuspendeduntil(value: number): void;

  getTimecreation(): number;
  setTimecreation(value: number): void;

  getExpansion(): number;
  setExpansion(value: number): void;

  getBanreason(): string;
  setBanreason(value: string): void;

  getSuspendreason(): string;
  setSuspendreason(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Account.AsObject;
  static toObject(includeInstance: boolean, msg: Account): Account.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Account, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Account;
  static deserializeBinaryFromReader(message: Account, reader: jspb.BinaryReader): Account;
}

export namespace Account {
  export type AsObject = {
    id: number,
    name: string,
    charname: string,
    sharedplat: number,
    password: string,
    status: number,
    lsid: string,
    lsaccountid: number,
    gmspeed: number,
    revoked: number,
    karma: number,
    miniloginip: string,
    hideme: number,
    rulesflag: number,
    suspendeduntil: number,
    timecreation: number,
    expansion: number,
    banreason: string,
    suspendreason: string,
  }
}

