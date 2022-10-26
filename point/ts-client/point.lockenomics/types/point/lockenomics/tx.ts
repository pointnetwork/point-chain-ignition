/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "point.lockenomics";

export interface MsgCreateLock {
  creator: string;
  lenght: number;
  validator: string;
}

export interface MsgCreateLockResponse {
}

export interface MsgCreateDelegatedAmount {
  creator: string;
  index: string;
  delegator: string;
  validator: string;
  amount: number;
}

export interface MsgCreateDelegatedAmountResponse {
}

export interface MsgUpdateDelegatedAmount {
  creator: string;
  index: string;
  delegator: string;
  validator: string;
  amount: number;
}

export interface MsgUpdateDelegatedAmountResponse {
}

export interface MsgDeleteDelegatedAmount {
  creator: string;
  index: string;
}

export interface MsgDeleteDelegatedAmountResponse {
}

function createBaseMsgCreateLock(): MsgCreateLock {
  return { creator: "", lenght: 0, validator: "" };
}

export const MsgCreateLock = {
  encode(message: MsgCreateLock, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.lenght !== 0) {
      writer.uint32(16).int32(message.lenght);
    }
    if (message.validator !== "") {
      writer.uint32(26).string(message.validator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateLock {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateLock();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.lenght = reader.int32();
          break;
        case 3:
          message.validator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateLock {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      lenght: isSet(object.lenght) ? Number(object.lenght) : 0,
      validator: isSet(object.validator) ? String(object.validator) : "",
    };
  },

  toJSON(message: MsgCreateLock): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.lenght !== undefined && (obj.lenght = Math.round(message.lenght));
    message.validator !== undefined && (obj.validator = message.validator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateLock>, I>>(object: I): MsgCreateLock {
    const message = createBaseMsgCreateLock();
    message.creator = object.creator ?? "";
    message.lenght = object.lenght ?? 0;
    message.validator = object.validator ?? "";
    return message;
  },
};

function createBaseMsgCreateLockResponse(): MsgCreateLockResponse {
  return {};
}

export const MsgCreateLockResponse = {
  encode(_: MsgCreateLockResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateLockResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateLockResponse();
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

  fromJSON(_: any): MsgCreateLockResponse {
    return {};
  },

  toJSON(_: MsgCreateLockResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateLockResponse>, I>>(_: I): MsgCreateLockResponse {
    const message = createBaseMsgCreateLockResponse();
    return message;
  },
};

function createBaseMsgCreateDelegatedAmount(): MsgCreateDelegatedAmount {
  return { creator: "", index: "", delegator: "", validator: "", amount: 0 };
}

export const MsgCreateDelegatedAmount = {
  encode(message: MsgCreateDelegatedAmount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.delegator !== "") {
      writer.uint32(26).string(message.delegator);
    }
    if (message.validator !== "") {
      writer.uint32(34).string(message.validator);
    }
    if (message.amount !== 0) {
      writer.uint32(40).uint64(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateDelegatedAmount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateDelegatedAmount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.delegator = reader.string();
          break;
        case 4:
          message.validator = reader.string();
          break;
        case 5:
          message.amount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDelegatedAmount {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      delegator: isSet(object.delegator) ? String(object.delegator) : "",
      validator: isSet(object.validator) ? String(object.validator) : "",
      amount: isSet(object.amount) ? Number(object.amount) : 0,
    };
  },

  toJSON(message: MsgCreateDelegatedAmount): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.delegator !== undefined && (obj.delegator = message.delegator);
    message.validator !== undefined && (obj.validator = message.validator);
    message.amount !== undefined && (obj.amount = Math.round(message.amount));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateDelegatedAmount>, I>>(object: I): MsgCreateDelegatedAmount {
    const message = createBaseMsgCreateDelegatedAmount();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.delegator = object.delegator ?? "";
    message.validator = object.validator ?? "";
    message.amount = object.amount ?? 0;
    return message;
  },
};

function createBaseMsgCreateDelegatedAmountResponse(): MsgCreateDelegatedAmountResponse {
  return {};
}

export const MsgCreateDelegatedAmountResponse = {
  encode(_: MsgCreateDelegatedAmountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateDelegatedAmountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateDelegatedAmountResponse();
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

  fromJSON(_: any): MsgCreateDelegatedAmountResponse {
    return {};
  },

  toJSON(_: MsgCreateDelegatedAmountResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateDelegatedAmountResponse>, I>>(
    _: I,
  ): MsgCreateDelegatedAmountResponse {
    const message = createBaseMsgCreateDelegatedAmountResponse();
    return message;
  },
};

function createBaseMsgUpdateDelegatedAmount(): MsgUpdateDelegatedAmount {
  return { creator: "", index: "", delegator: "", validator: "", amount: 0 };
}

export const MsgUpdateDelegatedAmount = {
  encode(message: MsgUpdateDelegatedAmount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.delegator !== "") {
      writer.uint32(26).string(message.delegator);
    }
    if (message.validator !== "") {
      writer.uint32(34).string(message.validator);
    }
    if (message.amount !== 0) {
      writer.uint32(40).uint64(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDelegatedAmount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDelegatedAmount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.delegator = reader.string();
          break;
        case 4:
          message.validator = reader.string();
          break;
        case 5:
          message.amount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateDelegatedAmount {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      delegator: isSet(object.delegator) ? String(object.delegator) : "",
      validator: isSet(object.validator) ? String(object.validator) : "",
      amount: isSet(object.amount) ? Number(object.amount) : 0,
    };
  },

  toJSON(message: MsgUpdateDelegatedAmount): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.delegator !== undefined && (obj.delegator = message.delegator);
    message.validator !== undefined && (obj.validator = message.validator);
    message.amount !== undefined && (obj.amount = Math.round(message.amount));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDelegatedAmount>, I>>(object: I): MsgUpdateDelegatedAmount {
    const message = createBaseMsgUpdateDelegatedAmount();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.delegator = object.delegator ?? "";
    message.validator = object.validator ?? "";
    message.amount = object.amount ?? 0;
    return message;
  },
};

function createBaseMsgUpdateDelegatedAmountResponse(): MsgUpdateDelegatedAmountResponse {
  return {};
}

export const MsgUpdateDelegatedAmountResponse = {
  encode(_: MsgUpdateDelegatedAmountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDelegatedAmountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDelegatedAmountResponse();
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

  fromJSON(_: any): MsgUpdateDelegatedAmountResponse {
    return {};
  },

  toJSON(_: MsgUpdateDelegatedAmountResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDelegatedAmountResponse>, I>>(
    _: I,
  ): MsgUpdateDelegatedAmountResponse {
    const message = createBaseMsgUpdateDelegatedAmountResponse();
    return message;
  },
};

function createBaseMsgDeleteDelegatedAmount(): MsgDeleteDelegatedAmount {
  return { creator: "", index: "" };
}

export const MsgDeleteDelegatedAmount = {
  encode(message: MsgDeleteDelegatedAmount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteDelegatedAmount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteDelegatedAmount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteDelegatedAmount {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
    };
  },

  toJSON(message: MsgDeleteDelegatedAmount): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteDelegatedAmount>, I>>(object: I): MsgDeleteDelegatedAmount {
    const message = createBaseMsgDeleteDelegatedAmount();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseMsgDeleteDelegatedAmountResponse(): MsgDeleteDelegatedAmountResponse {
  return {};
}

export const MsgDeleteDelegatedAmountResponse = {
  encode(_: MsgDeleteDelegatedAmountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteDelegatedAmountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteDelegatedAmountResponse();
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

  fromJSON(_: any): MsgDeleteDelegatedAmountResponse {
    return {};
  },

  toJSON(_: MsgDeleteDelegatedAmountResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteDelegatedAmountResponse>, I>>(
    _: I,
  ): MsgDeleteDelegatedAmountResponse {
    const message = createBaseMsgDeleteDelegatedAmountResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateLock(request: MsgCreateLock): Promise<MsgCreateLockResponse>;
  CreateDelegatedAmount(request: MsgCreateDelegatedAmount): Promise<MsgCreateDelegatedAmountResponse>;
  UpdateDelegatedAmount(request: MsgUpdateDelegatedAmount): Promise<MsgUpdateDelegatedAmountResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteDelegatedAmount(request: MsgDeleteDelegatedAmount): Promise<MsgDeleteDelegatedAmountResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateLock = this.CreateLock.bind(this);
    this.CreateDelegatedAmount = this.CreateDelegatedAmount.bind(this);
    this.UpdateDelegatedAmount = this.UpdateDelegatedAmount.bind(this);
    this.DeleteDelegatedAmount = this.DeleteDelegatedAmount.bind(this);
  }
  CreateLock(request: MsgCreateLock): Promise<MsgCreateLockResponse> {
    const data = MsgCreateLock.encode(request).finish();
    const promise = this.rpc.request("point.lockenomics.Msg", "CreateLock", data);
    return promise.then((data) => MsgCreateLockResponse.decode(new _m0.Reader(data)));
  }

  CreateDelegatedAmount(request: MsgCreateDelegatedAmount): Promise<MsgCreateDelegatedAmountResponse> {
    const data = MsgCreateDelegatedAmount.encode(request).finish();
    const promise = this.rpc.request("point.lockenomics.Msg", "CreateDelegatedAmount", data);
    return promise.then((data) => MsgCreateDelegatedAmountResponse.decode(new _m0.Reader(data)));
  }

  UpdateDelegatedAmount(request: MsgUpdateDelegatedAmount): Promise<MsgUpdateDelegatedAmountResponse> {
    const data = MsgUpdateDelegatedAmount.encode(request).finish();
    const promise = this.rpc.request("point.lockenomics.Msg", "UpdateDelegatedAmount", data);
    return promise.then((data) => MsgUpdateDelegatedAmountResponse.decode(new _m0.Reader(data)));
  }

  DeleteDelegatedAmount(request: MsgDeleteDelegatedAmount): Promise<MsgDeleteDelegatedAmountResponse> {
    const data = MsgDeleteDelegatedAmount.encode(request).finish();
    const promise = this.rpc.request("point.lockenomics.Msg", "DeleteDelegatedAmount", data);
    return promise.then((data) => MsgDeleteDelegatedAmountResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
