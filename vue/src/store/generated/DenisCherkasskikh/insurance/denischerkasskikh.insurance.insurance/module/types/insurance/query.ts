/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../insurance/params";
import { NextRec } from "../insurance/next_rec";
import { CrashRec } from "../insurance/crash_rec";
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

export interface QueryGetNextRecRequest {}

export interface QueryGetNextRecResponse {
  NextRec: NextRec | undefined;
}

export interface QueryGetCrashRecRequest {
  index: string;
}

export interface QueryGetCrashRecResponse {
  crashRec: CrashRec | undefined;
}

export interface QueryAllCrashRecRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllCrashRecResponse {
  crashRec: CrashRec[];
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

const baseQueryGetNextRecRequest: object = {};

export const QueryGetNextRecRequest = {
  encode(_: QueryGetNextRecRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNextRecRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetNextRecRequest } as QueryGetNextRecRequest;
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

  fromJSON(_: any): QueryGetNextRecRequest {
    const message = { ...baseQueryGetNextRecRequest } as QueryGetNextRecRequest;
    return message;
  },

  toJSON(_: QueryGetNextRecRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryGetNextRecRequest>): QueryGetNextRecRequest {
    const message = { ...baseQueryGetNextRecRequest } as QueryGetNextRecRequest;
    return message;
  },
};

const baseQueryGetNextRecResponse: object = {};

export const QueryGetNextRecResponse = {
  encode(
    message: QueryGetNextRecResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.NextRec !== undefined) {
      NextRec.encode(message.NextRec, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNextRecResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNextRecResponse,
    } as QueryGetNextRecResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.NextRec = NextRec.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNextRecResponse {
    const message = {
      ...baseQueryGetNextRecResponse,
    } as QueryGetNextRecResponse;
    if (object.NextRec !== undefined && object.NextRec !== null) {
      message.NextRec = NextRec.fromJSON(object.NextRec);
    } else {
      message.NextRec = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNextRecResponse): unknown {
    const obj: any = {};
    message.NextRec !== undefined &&
      (obj.NextRec = message.NextRec
        ? NextRec.toJSON(message.NextRec)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNextRecResponse>
  ): QueryGetNextRecResponse {
    const message = {
      ...baseQueryGetNextRecResponse,
    } as QueryGetNextRecResponse;
    if (object.NextRec !== undefined && object.NextRec !== null) {
      message.NextRec = NextRec.fromPartial(object.NextRec);
    } else {
      message.NextRec = undefined;
    }
    return message;
  },
};

const baseQueryGetCrashRecRequest: object = { index: "" };

export const QueryGetCrashRecRequest = {
  encode(
    message: QueryGetCrashRecRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetCrashRecRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetCrashRecRequest,
    } as QueryGetCrashRecRequest;
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

  fromJSON(object: any): QueryGetCrashRecRequest {
    const message = {
      ...baseQueryGetCrashRecRequest,
    } as QueryGetCrashRecRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetCrashRecRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetCrashRecRequest>
  ): QueryGetCrashRecRequest {
    const message = {
      ...baseQueryGetCrashRecRequest,
    } as QueryGetCrashRecRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetCrashRecResponse: object = {};

export const QueryGetCrashRecResponse = {
  encode(
    message: QueryGetCrashRecResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.crashRec !== undefined) {
      CrashRec.encode(message.crashRec, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetCrashRecResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetCrashRecResponse,
    } as QueryGetCrashRecResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.crashRec = CrashRec.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCrashRecResponse {
    const message = {
      ...baseQueryGetCrashRecResponse,
    } as QueryGetCrashRecResponse;
    if (object.crashRec !== undefined && object.crashRec !== null) {
      message.crashRec = CrashRec.fromJSON(object.crashRec);
    } else {
      message.crashRec = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetCrashRecResponse): unknown {
    const obj: any = {};
    message.crashRec !== undefined &&
      (obj.crashRec = message.crashRec
        ? CrashRec.toJSON(message.crashRec)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetCrashRecResponse>
  ): QueryGetCrashRecResponse {
    const message = {
      ...baseQueryGetCrashRecResponse,
    } as QueryGetCrashRecResponse;
    if (object.crashRec !== undefined && object.crashRec !== null) {
      message.crashRec = CrashRec.fromPartial(object.crashRec);
    } else {
      message.crashRec = undefined;
    }
    return message;
  },
};

const baseQueryAllCrashRecRequest: object = {};

export const QueryAllCrashRecRequest = {
  encode(
    message: QueryAllCrashRecRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllCrashRecRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllCrashRecRequest,
    } as QueryAllCrashRecRequest;
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

  fromJSON(object: any): QueryAllCrashRecRequest {
    const message = {
      ...baseQueryAllCrashRecRequest,
    } as QueryAllCrashRecRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllCrashRecRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllCrashRecRequest>
  ): QueryAllCrashRecRequest {
    const message = {
      ...baseQueryAllCrashRecRequest,
    } as QueryAllCrashRecRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllCrashRecResponse: object = {};

export const QueryAllCrashRecResponse = {
  encode(
    message: QueryAllCrashRecResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.crashRec) {
      CrashRec.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllCrashRecResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllCrashRecResponse,
    } as QueryAllCrashRecResponse;
    message.crashRec = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.crashRec.push(CrashRec.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllCrashRecResponse {
    const message = {
      ...baseQueryAllCrashRecResponse,
    } as QueryAllCrashRecResponse;
    message.crashRec = [];
    if (object.crashRec !== undefined && object.crashRec !== null) {
      for (const e of object.crashRec) {
        message.crashRec.push(CrashRec.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllCrashRecResponse): unknown {
    const obj: any = {};
    if (message.crashRec) {
      obj.crashRec = message.crashRec.map((e) =>
        e ? CrashRec.toJSON(e) : undefined
      );
    } else {
      obj.crashRec = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllCrashRecResponse>
  ): QueryAllCrashRecResponse {
    const message = {
      ...baseQueryAllCrashRecResponse,
    } as QueryAllCrashRecResponse;
    message.crashRec = [];
    if (object.crashRec !== undefined && object.crashRec !== null) {
      for (const e of object.crashRec) {
        message.crashRec.push(CrashRec.fromPartial(e));
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
  /** Queries a NextRec by index. */
  NextRec(request: QueryGetNextRecRequest): Promise<QueryGetNextRecResponse>;
  /** Queries a CrashRec by index. */
  CrashRec(request: QueryGetCrashRecRequest): Promise<QueryGetCrashRecResponse>;
  /** Queries a list of CrashRec items. */
  CrashRecAll(
    request: QueryAllCrashRecRequest
  ): Promise<QueryAllCrashRecResponse>;
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

  NextRec(request: QueryGetNextRecRequest): Promise<QueryGetNextRecResponse> {
    const data = QueryGetNextRecRequest.encode(request).finish();
    const promise = this.rpc.request(
      "denischerkasskikh.insurance.insurance.Query",
      "NextRec",
      data
    );
    return promise.then((data) =>
      QueryGetNextRecResponse.decode(new Reader(data))
    );
  }

  CrashRec(
    request: QueryGetCrashRecRequest
  ): Promise<QueryGetCrashRecResponse> {
    const data = QueryGetCrashRecRequest.encode(request).finish();
    const promise = this.rpc.request(
      "denischerkasskikh.insurance.insurance.Query",
      "CrashRec",
      data
    );
    return promise.then((data) =>
      QueryGetCrashRecResponse.decode(new Reader(data))
    );
  }

  CrashRecAll(
    request: QueryAllCrashRecRequest
  ): Promise<QueryAllCrashRecResponse> {
    const data = QueryAllCrashRecRequest.encode(request).finish();
    const promise = this.rpc.request(
      "denischerkasskikh.insurance.insurance.Query",
      "CrashRecAll",
      data
    );
    return promise.then((data) =>
      QueryAllCrashRecResponse.decode(new Reader(data))
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
