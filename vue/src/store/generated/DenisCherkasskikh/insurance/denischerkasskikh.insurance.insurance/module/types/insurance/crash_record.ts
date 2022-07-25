/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "denischerkasskikh.insurance.insurance";

export interface CrashRecord {
  index: string;
  brand: string;
  model: string;
  owner: string;
  licencePlate: string;
  vinNumber: string;
  mileage: string;
  side: string;
  part: string;
  payout: string;
}

const baseCrashRecord: object = {
  index: "",
  brand: "",
  model: "",
  owner: "",
  licencePlate: "",
  vinNumber: "",
  mileage: "",
  side: "",
  part: "",
  payout: "",
};

export const CrashRecord = {
  encode(message: CrashRecord, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.brand !== "") {
      writer.uint32(18).string(message.brand);
    }
    if (message.model !== "") {
      writer.uint32(26).string(message.model);
    }
    if (message.owner !== "") {
      writer.uint32(34).string(message.owner);
    }
    if (message.licencePlate !== "") {
      writer.uint32(42).string(message.licencePlate);
    }
    if (message.vinNumber !== "") {
      writer.uint32(50).string(message.vinNumber);
    }
    if (message.mileage !== "") {
      writer.uint32(58).string(message.mileage);
    }
    if (message.side !== "") {
      writer.uint32(66).string(message.side);
    }
    if (message.part !== "") {
      writer.uint32(74).string(message.part);
    }
    if (message.payout !== "") {
      writer.uint32(82).string(message.payout);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CrashRecord {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCrashRecord } as CrashRecord;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.brand = reader.string();
          break;
        case 3:
          message.model = reader.string();
          break;
        case 4:
          message.owner = reader.string();
          break;
        case 5:
          message.licencePlate = reader.string();
          break;
        case 6:
          message.vinNumber = reader.string();
          break;
        case 7:
          message.mileage = reader.string();
          break;
        case 8:
          message.side = reader.string();
          break;
        case 9:
          message.part = reader.string();
          break;
        case 10:
          message.payout = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CrashRecord {
    const message = { ...baseCrashRecord } as CrashRecord;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.brand !== undefined && object.brand !== null) {
      message.brand = String(object.brand);
    } else {
      message.brand = "";
    }
    if (object.model !== undefined && object.model !== null) {
      message.model = String(object.model);
    } else {
      message.model = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.licencePlate !== undefined && object.licencePlate !== null) {
      message.licencePlate = String(object.licencePlate);
    } else {
      message.licencePlate = "";
    }
    if (object.vinNumber !== undefined && object.vinNumber !== null) {
      message.vinNumber = String(object.vinNumber);
    } else {
      message.vinNumber = "";
    }
    if (object.mileage !== undefined && object.mileage !== null) {
      message.mileage = String(object.mileage);
    } else {
      message.mileage = "";
    }
    if (object.side !== undefined && object.side !== null) {
      message.side = String(object.side);
    } else {
      message.side = "";
    }
    if (object.part !== undefined && object.part !== null) {
      message.part = String(object.part);
    } else {
      message.part = "";
    }
    if (object.payout !== undefined && object.payout !== null) {
      message.payout = String(object.payout);
    } else {
      message.payout = "";
    }
    return message;
  },

  toJSON(message: CrashRecord): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.brand !== undefined && (obj.brand = message.brand);
    message.model !== undefined && (obj.model = message.model);
    message.owner !== undefined && (obj.owner = message.owner);
    message.licencePlate !== undefined &&
      (obj.licencePlate = message.licencePlate);
    message.vinNumber !== undefined && (obj.vinNumber = message.vinNumber);
    message.mileage !== undefined && (obj.mileage = message.mileage);
    message.side !== undefined && (obj.side = message.side);
    message.part !== undefined && (obj.part = message.part);
    message.payout !== undefined && (obj.payout = message.payout);
    return obj;
  },

  fromPartial(object: DeepPartial<CrashRecord>): CrashRecord {
    const message = { ...baseCrashRecord } as CrashRecord;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.brand !== undefined && object.brand !== null) {
      message.brand = object.brand;
    } else {
      message.brand = "";
    }
    if (object.model !== undefined && object.model !== null) {
      message.model = object.model;
    } else {
      message.model = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.licencePlate !== undefined && object.licencePlate !== null) {
      message.licencePlate = object.licencePlate;
    } else {
      message.licencePlate = "";
    }
    if (object.vinNumber !== undefined && object.vinNumber !== null) {
      message.vinNumber = object.vinNumber;
    } else {
      message.vinNumber = "";
    }
    if (object.mileage !== undefined && object.mileage !== null) {
      message.mileage = object.mileage;
    } else {
      message.mileage = "";
    }
    if (object.side !== undefined && object.side !== null) {
      message.side = object.side;
    } else {
      message.side = "";
    }
    if (object.part !== undefined && object.part !== null) {
      message.part = object.part;
    } else {
      message.part = "";
    }
    if (object.payout !== undefined && object.payout !== null) {
      message.payout = object.payout;
    } else {
      message.payout = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
