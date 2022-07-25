/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../insurance/params";
import { CrashRecord } from "../insurance/crash_record";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "denischerkasskikh.insurance.insurance";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetCrashRecordRequest {
  index: string;
}

export interface QueryGetCrashRecordResponse {
  crashRecord: CrashRecord | undefined;
}

export interface QueryAllCrashRecordRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllCrashRecordResponse {
  crashRecord: CrashRecord[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetCrashRecordRequest: object = { index: "" };

export const QueryGetCrashRecordRequest = {
  encode(
    message: QueryGetCrashRecordRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetCrashRecordRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetCrashRecordRequest,
    } as QueryGetCrashRecordRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCrashRecordRequest {
    const message = {
      ...baseQueryGetCrashRecordRequest,
    } as QueryGetCrashRecordRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetCrashRecordRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetCrashRecordRequest>
  ): QueryGetCrashRecordRequest {
    const message = {
      ...baseQueryGetCrashRecordRequest,
    } as QueryGetCrashRecordRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetCrashRecordResponse: object = {};

export const QueryGetCrashRecordResponse = {
  encode(
    message: QueryGetCrashRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.crashRecord !== undefined) {
      CrashRecord.encode(
        message.crashRecord,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetCrashRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetCrashRecordResponse,
    } as QueryGetCrashRecordResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.crashRecord = CrashRecord.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCrashRecordResponse {
    const message = {
      ...baseQueryGetCrashRecordResponse,
    } as QueryGetCrashRecordResponse;
    if (object.crashRecord !== undefined && object.crashRecord !== null) {
      message.crashRecord = CrashRecord.fromJSON(object.crashRecord);
    } else {
      message.crashRecord = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetCrashRecordResponse): unknown {
    const obj: any = {};
    message.crashRecord !== undefined &&
      (obj.crashRecord = message.crashRecord
        ? CrashRecord.toJSON(message.crashRecord)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetCrashRecordResponse>
  ): QueryGetCrashRecordResponse {
    const message = {
      ...baseQueryGetCrashRecordResponse,
    } as QueryGetCrashRecordResponse;
    if (object.crashRecord !== undefined && object.crashRecord !== null) {
      message.crashRecord = CrashRecord.fromPartial(object.crashRecord);
    } else {
      message.crashRecord = undefined;
    }
    return message;
  },
};

const baseQueryAllCrashRecordRequest: object = {};

export const QueryAllCrashRecordRequest = {
  encode(
    message: QueryAllCrashRecordRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllCrashRecordRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllCrashRecordRequest,
    } as QueryAllCrashRecordRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllCrashRecordRequest {
    const message = {
      ...baseQueryAllCrashRecordRequest,
    } as QueryAllCrashRecordRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllCrashRecordRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllCrashRecordRequest>
  ): QueryAllCrashRecordRequest {
    const message = {
      ...baseQueryAllCrashRecordRequest,
    } as QueryAllCrashRecordRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllCrashRecordResponse: object = {};

export const QueryAllCrashRecordResponse = {
  encode(
    message: QueryAllCrashRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.crashRecord) {
      CrashRecord.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllCrashRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllCrashRecordResponse,
    } as QueryAllCrashRecordResponse;
    message.crashRecord = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.crashRecord.push(CrashRecord.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllCrashRecordResponse {
    const message = {
      ...baseQueryAllCrashRecordResponse,
    } as QueryAllCrashRecordResponse;
    message.crashRecord = [];
    if (object.crashRecord !== undefined && object.crashRecord !== null) {
      for (const e of object.crashRecord) {
        message.crashRecord.push(CrashRecord.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllCrashRecordResponse): unknown {
    const obj: any = {};
    if (message.crashRecord) {
      obj.crashRecord = message.crashRecord.map((e) =>
        e ? CrashRecord.toJSON(e) : undefined
      );
    } else {
      obj.crashRecord = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllCrashRecordResponse>
  ): QueryAllCrashRecordResponse {
    const message = {
      ...baseQueryAllCrashRecordResponse,
    } as QueryAllCrashRecordResponse;
    message.crashRecord = [];
    if (object.crashRecord !== undefined && object.crashRecord !== null) {
      for (const e of object.crashRecord) {
        message.crashRecord.push(CrashRecord.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a CrashRecord by index. */
  CrashRecord(
    request: QueryGetCrashRecordRequest
  ): Promise<QueryGetCrashRecordResponse>;
  /** Queries a list of CrashRecord items. */
  CrashRecordAll(
    request: QueryAllCrashRecordRequest
  ): Promise<QueryAllCrashRecordResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "denischerkasskikh.insurance.insurance.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  CrashRecord(
    request: QueryGetCrashRecordRequest
  ): Promise<QueryGetCrashRecordResponse> {
    const data = QueryGetCrashRecordRequest.encode(request).finish();
    const promise = this.rpc.request(
      "denischerkasskikh.insurance.insurance.Query",
      "CrashRecord",
      data
    );
    return promise.then((data) =>
      QueryGetCrashRecordResponse.decode(new Reader(data))
    );
  }

  CrashRecordAll(
    request: QueryAllCrashRecordRequest
  ): Promise<QueryAllCrashRecordResponse> {
    const data = QueryAllCrashRecordRequest.encode(request).finish();
    const promise = this.rpc.request(
      "denischerkasskikh.insurance.insurance.Query",
      "CrashRecordAll",
      data
    );
    return promise.then((data) =>
      QueryAllCrashRecordResponse.decode(new Reader(data))
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
