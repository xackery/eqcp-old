// package: pb
// file: proto/inventory.proto

import * as jspb from "google-protobuf";

export class Inventory extends jspb.Message {
  getCharid(): number;
  setCharid(value: number): void;

  getSlotid(): number;
  setSlotid(value: number): void;

  getItemid(): number;
  setItemid(value: number): void;

  getCharges(): number;
  setCharges(value: number): void;

  getColor(): number;
  setColor(value: number): void;

  getAugslot1(): number;
  setAugslot1(value: number): void;

  getAugslot2(): number;
  setAugslot2(value: number): void;

  getAugslot3(): number;
  setAugslot3(value: number): void;

  getAugslot4(): number;
  setAugslot4(value: number): void;

  getAugslot5(): number;
  setAugslot5(value: number): void;

  getAugslot6(): number;
  setAugslot6(value: number): void;

  getInstnodrop(): number;
  setInstnodrop(value: number): void;

  getCustomdata(): string;
  setCustomdata(value: string): void;

  getOrnamenticon(): number;
  setOrnamenticon(value: number): void;

  getOrnamentidfile(): number;
  setOrnamentidfile(value: number): void;

  getOrnamentheromodel(): number;
  setOrnamentheromodel(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Inventory.AsObject;
  static toObject(includeInstance: boolean, msg: Inventory): Inventory.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Inventory, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Inventory;
  static deserializeBinaryFromReader(message: Inventory, reader: jspb.BinaryReader): Inventory;
}

export namespace Inventory {
  export type AsObject = {
    charid: number,
    slotid: number,
    itemid: number,
    charges: number,
    color: number,
    augslot1: number,
    augslot2: number,
    augslot3: number,
    augslot4: number,
    augslot5: number,
    augslot6: number,
    instnodrop: number,
    customdata: string,
    ornamenticon: number,
    ornamentidfile: number,
    ornamentheromodel: number,
  }
}

