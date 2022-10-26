/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { DelegatedAmount } from "./delegated_amount";
import { DelegationLock } from "./delegation_lock";
import { Params } from "./params";

export const protobufPackage = "point.lockenomics";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetDelegationLockRequest {
  index: string;
}

export interface QueryGetDelegationLockResponse {
  delegationLock: DelegationLock | undefined;
}

export interface QueryAllDelegationLockRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDelegationLockResponse {
  delegationLock: DelegationLock[];
  pagination: PageResponse | undefined;
}

export interface QueryGetDelegatedAmountRequest {
  index: string;
}

export interface QueryGetDelegatedAmountResponse {
  delegatedAmount: DelegatedAmount | undefined;
}

export interface QueryAllDelegatedAmountRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDelegatedAmountResponse {
  delegatedAmount: DelegatedAmount[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
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
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetDelegationLockRequest(): QueryGetDelegationLockRequest {
  return { index: "" };
}

export const QueryGetDelegationLockRequest = {
  encode(message: QueryGetDelegationLockRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDelegationLockRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDelegationLockRequest();
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

  fromJSON(object: any): QueryGetDelegationLockRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetDelegationLockRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDelegationLockRequest>, I>>(
    object: I,
  ): QueryGetDelegationLockRequest {
    const message = createBaseQueryGetDelegationLockRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetDelegationLockResponse(): QueryGetDelegationLockResponse {
  return { delegationLock: undefined };
}

export const QueryGetDelegationLockResponse = {
  encode(message: QueryGetDelegationLockResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.delegationLock !== undefined) {
      DelegationLock.encode(message.delegationLock, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDelegationLockResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDelegationLockResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.delegationLock = DelegationLock.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDelegationLockResponse {
    return {
      delegationLock: isSet(object.delegationLock) ? DelegationLock.fromJSON(object.delegationLock) : undefined,
    };
  },

  toJSON(message: QueryGetDelegationLockResponse): unknown {
    const obj: any = {};
    message.delegationLock !== undefined
      && (obj.delegationLock = message.delegationLock ? DelegationLock.toJSON(message.delegationLock) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDelegationLockResponse>, I>>(
    object: I,
  ): QueryGetDelegationLockResponse {
    const message = createBaseQueryGetDelegationLockResponse();
    message.delegationLock = (object.delegationLock !== undefined && object.delegationLock !== null)
      ? DelegationLock.fromPartial(object.delegationLock)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDelegationLockRequest(): QueryAllDelegationLockRequest {
  return { pagination: undefined };
}

export const QueryAllDelegationLockRequest = {
  encode(message: QueryAllDelegationLockRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDelegationLockRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDelegationLockRequest();
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

  fromJSON(object: any): QueryAllDelegationLockRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllDelegationLockRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDelegationLockRequest>, I>>(
    object: I,
  ): QueryAllDelegationLockRequest {
    const message = createBaseQueryAllDelegationLockRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDelegationLockResponse(): QueryAllDelegationLockResponse {
  return { delegationLock: [], pagination: undefined };
}

export const QueryAllDelegationLockResponse = {
  encode(message: QueryAllDelegationLockResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.delegationLock) {
      DelegationLock.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDelegationLockResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDelegationLockResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.delegationLock.push(DelegationLock.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllDelegationLockResponse {
    return {
      delegationLock: Array.isArray(object?.delegationLock)
        ? object.delegationLock.map((e: any) => DelegationLock.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllDelegationLockResponse): unknown {
    const obj: any = {};
    if (message.delegationLock) {
      obj.delegationLock = message.delegationLock.map((e) => e ? DelegationLock.toJSON(e) : undefined);
    } else {
      obj.delegationLock = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDelegationLockResponse>, I>>(
    object: I,
  ): QueryAllDelegationLockResponse {
    const message = createBaseQueryAllDelegationLockResponse();
    message.delegationLock = object.delegationLock?.map((e) => DelegationLock.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetDelegatedAmountRequest(): QueryGetDelegatedAmountRequest {
  return { index: "" };
}

export const QueryGetDelegatedAmountRequest = {
  encode(message: QueryGetDelegatedAmountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDelegatedAmountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDelegatedAmountRequest();
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

  fromJSON(object: any): QueryGetDelegatedAmountRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetDelegatedAmountRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDelegatedAmountRequest>, I>>(
    object: I,
  ): QueryGetDelegatedAmountRequest {
    const message = createBaseQueryGetDelegatedAmountRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetDelegatedAmountResponse(): QueryGetDelegatedAmountResponse {
  return { delegatedAmount: undefined };
}

export const QueryGetDelegatedAmountResponse = {
  encode(message: QueryGetDelegatedAmountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.delegatedAmount !== undefined) {
      DelegatedAmount.encode(message.delegatedAmount, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDelegatedAmountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDelegatedAmountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.delegatedAmount = DelegatedAmount.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDelegatedAmountResponse {
    return {
      delegatedAmount: isSet(object.delegatedAmount) ? DelegatedAmount.fromJSON(object.delegatedAmount) : undefined,
    };
  },

  toJSON(message: QueryGetDelegatedAmountResponse): unknown {
    const obj: any = {};
    message.delegatedAmount !== undefined
      && (obj.delegatedAmount = message.delegatedAmount ? DelegatedAmount.toJSON(message.delegatedAmount) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDelegatedAmountResponse>, I>>(
    object: I,
  ): QueryGetDelegatedAmountResponse {
    const message = createBaseQueryGetDelegatedAmountResponse();
    message.delegatedAmount = (object.delegatedAmount !== undefined && object.delegatedAmount !== null)
      ? DelegatedAmount.fromPartial(object.delegatedAmount)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDelegatedAmountRequest(): QueryAllDelegatedAmountRequest {
  return { pagination: undefined };
}

export const QueryAllDelegatedAmountRequest = {
  encode(message: QueryAllDelegatedAmountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDelegatedAmountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDelegatedAmountRequest();
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

  fromJSON(object: any): QueryAllDelegatedAmountRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllDelegatedAmountRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDelegatedAmountRequest>, I>>(
    object: I,
  ): QueryAllDelegatedAmountRequest {
    const message = createBaseQueryAllDelegatedAmountRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDelegatedAmountResponse(): QueryAllDelegatedAmountResponse {
  return { delegatedAmount: [], pagination: undefined };
}

export const QueryAllDelegatedAmountResponse = {
  encode(message: QueryAllDelegatedAmountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.delegatedAmount) {
      DelegatedAmount.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDelegatedAmountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDelegatedAmountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.delegatedAmount.push(DelegatedAmount.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllDelegatedAmountResponse {
    return {
      delegatedAmount: Array.isArray(object?.delegatedAmount)
        ? object.delegatedAmount.map((e: any) => DelegatedAmount.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllDelegatedAmountResponse): unknown {
    const obj: any = {};
    if (message.delegatedAmount) {
      obj.delegatedAmount = message.delegatedAmount.map((e) => e ? DelegatedAmount.toJSON(e) : undefined);
    } else {
      obj.delegatedAmount = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDelegatedAmountResponse>, I>>(
    object: I,
  ): QueryAllDelegatedAmountResponse {
    const message = createBaseQueryAllDelegatedAmountResponse();
    message.delegatedAmount = object.delegatedAmount?.map((e) => DelegatedAmount.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a DelegationLock by index. */
  DelegationLock(request: QueryGetDelegationLockRequest): Promise<QueryGetDelegationLockResponse>;
  /** Queries a list of DelegationLock items. */
  DelegationLockAll(request: QueryAllDelegationLockRequest): Promise<QueryAllDelegationLockResponse>;
  /** Queries a DelegatedAmount by index. */
  DelegatedAmount(request: QueryGetDelegatedAmountRequest): Promise<QueryGetDelegatedAmountResponse>;
  /** Queries a list of DelegatedAmount items. */
  DelegatedAmountAll(request: QueryAllDelegatedAmountRequest): Promise<QueryAllDelegatedAmountResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.DelegationLock = this.DelegationLock.bind(this);
    this.DelegationLockAll = this.DelegationLockAll.bind(this);
    this.DelegatedAmount = this.DelegatedAmount.bind(this);
    this.DelegatedAmountAll = this.DelegatedAmountAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("point.lockenomics.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  DelegationLock(request: QueryGetDelegationLockRequest): Promise<QueryGetDelegationLockResponse> {
    const data = QueryGetDelegationLockRequest.encode(request).finish();
    const promise = this.rpc.request("point.lockenomics.Query", "DelegationLock", data);
    return promise.then((data) => QueryGetDelegationLockResponse.decode(new _m0.Reader(data)));
  }

  DelegationLockAll(request: QueryAllDelegationLockRequest): Promise<QueryAllDelegationLockResponse> {
    const data = QueryAllDelegationLockRequest.encode(request).finish();
    const promise = this.rpc.request("point.lockenomics.Query", "DelegationLockAll", data);
    return promise.then((data) => QueryAllDelegationLockResponse.decode(new _m0.Reader(data)));
  }

  DelegatedAmount(request: QueryGetDelegatedAmountRequest): Promise<QueryGetDelegatedAmountResponse> {
    const data = QueryGetDelegatedAmountRequest.encode(request).finish();
    const promise = this.rpc.request("point.lockenomics.Query", "DelegatedAmount", data);
    return promise.then((data) => QueryGetDelegatedAmountResponse.decode(new _m0.Reader(data)));
  }

  DelegatedAmountAll(request: QueryAllDelegatedAmountRequest): Promise<QueryAllDelegatedAmountResponse> {
    const data = QueryAllDelegatedAmountRequest.encode(request).finish();
    const promise = this.rpc.request("point.lockenomics.Query", "DelegatedAmountAll", data);
    return promise.then((data) => QueryAllDelegatedAmountResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
