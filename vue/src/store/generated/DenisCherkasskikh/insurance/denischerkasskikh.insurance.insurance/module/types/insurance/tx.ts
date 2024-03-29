/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "denischerkasskikh.insurance.insurance";

export interface MsgAddRecord {
  creator: string;
  brand: string;
  model: string;
  year: string;
  owner: string;
  licensePlate: string;
  vinNumber: string;
  odometer: string;
  side: string;
  part: string;
  payout: string;
}

export interface MsgAddRecordResponse {
  recNum: string;
}

const baseMsgAddRecord: object = {
  creator: "",
  brand: "",
  model: "",
  year: "",
  owner: "",
  licensePlate: "",
  vinNumber: "",
  odometer: "",
  side: "",
  part: "",
  payout: "",
};

export const MsgAddRecord = {
  encode(message: MsgAddRecord, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.brand !== "") {
      writer.uint32(18).string(message.brand);
    }
    if (message.model !== "") {
      writer.uint32(26).string(message.model);
    }
    if (message.year !== "") {
      writer.uint32(34).string(message.year);
    }
    if (message.owner !== "") {
      writer.uint32(42).string(message.owner);
    }
    if (message.licensePlate !== "") {
      writer.uint32(50).string(message.licensePlate);
    }
    if (message.vinNumber !== "") {
      writer.uint32(58).string(message.vinNumber);
    }
    if (message.odometer !== "") {
      writer.uint32(66).string(message.odometer);
    }
    if (message.side !== "") {
      writer.uint32(74).string(message.side);
    }
    if (message.part !== "") {
      writer.uint32(82).string(message.part);
    }
    if (message.payout !== "") {
      writer.uint32(90).string(message.payout);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddRecord {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddRecord } as MsgAddRecord;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.brand = reader.string();
          break;
        case 3:
          message.model = reader.string();
          break;
        case 4:
          message.year = reader.string();
          break;
        case 5:
          message.owner = reader.string();
          break;
        case 6:
          message.licensePlate = reader.string();
          break;
        case 7:
          message.vinNumber = reader.string();
          break;
        case 8:
          message.odometer = reader.string();
          break;
        case 9:
          message.side = reader.string();
          break;
        case 10:
          message.part = reader.string();
          break;
        case 11:
          message.payout = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddRecord {
    const message = { ...baseMsgAddRecord } as MsgAddRecord;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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
    if (object.year !== undefined && object.year !== null) {
      message.year = String(object.year);
    } else {
      message.year = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.licensePlate !== undefined && object.licensePlate !== null) {
      message.licensePlate = String(object.licensePlate);
    } else {
      message.licensePlate = "";
    }
    if (object.vinNumber !== undefined && object.vinNumber !== null) {
      message.vinNumber = String(object.vinNumber);
    } else {
      message.vinNumber = "";
    }
    if (object.odometer !== undefined && object.odometer !== null) {
      message.odometer = String(object.odometer);
    } else {
      message.odometer = "";
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

  toJSON(message: MsgAddRecord): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.brand !== undefined && (obj.brand = message.brand);
    message.model !== undefined && (obj.model = message.model);
    message.year !== undefined && (obj.year = message.year);
    message.owner !== undefined && (obj.owner = message.owner);
    message.licensePlate !== undefined &&
      (obj.licensePlate = message.licensePlate);
    message.vinNumber !== undefined && (obj.vinNumber = message.vinNumber);
    message.odometer !== undefined && (obj.odometer = message.odometer);
    message.side !== undefined && (obj.side = message.side);
    message.part !== undefined && (obj.part = message.part);
    message.payout !== undefined && (obj.payout = message.payout);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddRecord>): MsgAddRecord {
    const message = { ...baseMsgAddRecord } as MsgAddRecord;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
    if (object.year !== undefined && object.year !== null) {
      message.year = object.year;
    } else {
      message.year = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.licensePlate !== undefined && object.licensePlate !== null) {
      message.licensePlate = object.licensePlate;
    } else {
      message.licensePlate = "";
    }
    if (object.vinNumber !== undefined && object.vinNumber !== null) {
      message.vinNumber = object.vinNumber;
    } else {
      message.vinNumber = "";
    }
    if (object.odometer !== undefined && object.odometer !== null) {
      message.odometer = object.odometer;
    } else {
      message.odometer = "";
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

const baseMsgAddRecordResponse: object = { recNum: "" };

export const MsgAddRecordResponse = {
  encode(
    message: MsgAddRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.recNum !== "") {
      writer.uint32(10).string(message.recNum);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddRecordResponse } as MsgAddRecordResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.recNum = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddRecordResponse {
    const message = { ...baseMsgAddRecordResponse } as MsgAddRecordResponse;
    if (object.recNum !== undefined && object.recNum !== null) {
      message.recNum = String(object.recNum);
    } else {
      message.recNum = "";
    }
    return message;
  },

  toJSON(message: MsgAddRecordResponse): unknown {
    const obj: any = {};
    message.recNum !== undefined && (obj.recNum = message.recNum);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddRecordResponse>): MsgAddRecordResponse {
    const message = { ...baseMsgAddRecordResponse } as MsgAddRecordResponse;
    if (object.recNum !== undefined && object.recNum !== null) {
      message.recNum = object.recNum;
    } else {
      message.recNum = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddRecord(request: MsgAddRecord): Promise<MsgAddRecordResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  AddRecord(request: MsgAddRecord): Promise<MsgAddRecordResponse> {
    const data = MsgAddRecord.encode(request).finish();
    const promise = this.rpc.request(
      "denischerkasskikh.insurance.insurance.Msg",
      "AddRecord",
      data
    );
    return promise.then((data) =>
      MsgAddRecordResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
