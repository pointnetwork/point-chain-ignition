/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "point.lockenomics";

export interface DelegatedAmount {
  index: string;
  delegator: string;
  validator: string;
  amount: number;
  creator: string;
}

function createBaseDelegatedAmount(): DelegatedAmount {
  return { index: "", delegator: "", validator: "", amount: 0, creator: "" };
}

export const DelegatedAmount = {
  encode(message: DelegatedAmount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.delegator !== "") {
      writer.uint32(18).string(message.delegator);
    }
    if (message.validator !== "") {
      writer.uint32(26).string(message.validator);
    }
    if (message.amount !== 0) {
      writer.uint32(32).uint64(message.amount);
    }
    if (message.creator !== "") {
      writer.uint32(42).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DelegatedAmount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDelegatedAmount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.delegator = reader.string();
          break;
        case 3:
          message.validator = reader.string();
          break;
        case 4:
          message.amount = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DelegatedAmount {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      delegator: isSet(object.delegator) ? String(object.delegator) : "",
      validator: isSet(object.validator) ? String(object.validator) : "",
      amount: isSet(object.amount) ? Number(object.amount) : 0,
      creator: isSet(object.creator) ? String(object.creator) : "",
    };
  },

  toJSON(message: DelegatedAmount): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.delegator !== undefined && (obj.delegator = message.delegator);
    message.validator !== undefined && (obj.validator = message.validator);
    message.amount !== undefined && (obj.amount = Math.round(message.amount));
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DelegatedAmount>, I>>(object: I): DelegatedAmount {
    const message = createBaseDelegatedAmount();
    message.index = object.index ?? "";
    message.delegator = object.delegator ?? "";
    message.validator = object.validator ?? "";
    message.amount = object.amount ?? 0;
    message.creator = object.creator ?? "";
    return message;
  },
};

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
