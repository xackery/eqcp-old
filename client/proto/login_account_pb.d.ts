// package: pb
// file: proto/login_account.proto

import * as jspb from "google-protobuf";

export class LoginAccount extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getAccountname(): string;
  setAccountname(value: string): void;

  getAccountpassword(): string;
  setAccountpassword(value: string): void;

  getAccountemail(): string;
  setAccountemail(value: string): void;

  getSourceloginserver(): string;
  setSourceloginserver(value: string): void;

  getLastipaddress(): string;
  setLastipaddress(value: string): void;

  getLastlogindate(): number;
  setLastlogindate(value: number): void;

  getCreatedat(): number;
  setCreatedat(value: number): void;

  getUpdatedat(): number;
  setUpdatedat(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginAccount.AsObject;
  static toObject(includeInstance: boolean, msg: LoginAccount): LoginAccount.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginAccount, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginAccount;
  static deserializeBinaryFromReader(message: LoginAccount, reader: jspb.BinaryReader): LoginAccount;
}

export namespace LoginAccount {
  export type AsObject = {
    id: number,
    accountname: string,
    accountpassword: string,
    accountemail: string,
    sourceloginserver: string,
    lastipaddress: string,
    lastlogindate: number,
    createdat: number,
    updatedat: number,
  }
}

