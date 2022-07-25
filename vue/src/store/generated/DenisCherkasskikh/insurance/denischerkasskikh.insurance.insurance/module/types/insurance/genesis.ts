/* eslint-disable */
import { Params } from "../insurance/params";
import { CrashRecord } from "../insurance/crash_record";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "denischerkasskikh.insurance.insurance";

/** GenesisState defines the insurance module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  crashRecordList: CrashRecord[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.crashRecordList) {
      CrashRecord.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.crashRecordList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.crashRecordList.push(
            CrashRecord.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.crashRecordList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.crashRecordList !== undefined &&
      object.crashRecordList !== null
    ) {
      for (const e of object.crashRecordList) {
        message.crashRecordList.push(CrashRecord.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.crashRecordList) {
      obj.crashRecordList = message.crashRecordList.map((e) =>
        e ? CrashRecord.toJSON(e) : undefined
      );
    } else {
      obj.crashRecordList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.crashRecordList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.crashRecordList !== undefined &&
      object.crashRecordList !== null
    ) {
      for (const e of object.crashRecordList) {
        message.crashRecordList.push(CrashRecord.fromPartial(e));
      }
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
