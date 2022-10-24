/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "point.lockenomics";

export interface DelegationLock {
  index: string;
  start: number;
  length: number;
  delegator: string;
  validator: string;
}

function createBaseDelegationLock(): DelegationLock {
  return { index: "", start: 0, length: 0, delegator: "", validator: "" };
}

export const DelegationLock = {
  encode(message: DelegationLock, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.start !== 0) {
      writer.uint32(16).int32(message.start);
    }
    if (message.length !== 0) {
      writer.uint32(24).int32(message.length);
    }
    if (message.delegator !== "") {
      writer.uint32(34).string(message.delegator);
    }
    if (message.validator !== "") {
      writer.uint32(42).string(message.validator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DelegationLock {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDelegationLock();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.start = reader.int32();
          break;
        case 3:
          message.length = reader.int32();
          break;
        case 4:
          message.delegator = reader.string();
          break;
        case 5:
          message.validator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DelegationLock {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      start: isSet(object.start) ? Number(object.start) : 0,
      length: isSet(object.length) ? Number(object.length) : 0,
      delegator: isSet(object.delegator) ? String(object.delegator) : "",
      validator: isSet(object.validator) ? String(object.validator) : "",
    };
  },

  toJSON(message: DelegationLock): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.start !== undefined && (obj.start = Math.round(message.start));
    message.length !== undefined && (obj.length = Math.round(message.length));
    message.delegator !== undefined && (obj.delegator = message.delegator);
    message.validator !== undefined && (obj.validator = message.validator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DelegationLock>, I>>(object: I): DelegationLock {
    const message = createBaseDelegationLock();
    message.index = object.index ?? "";
    message.start = object.start ?? 0;
    message.length = object.length ?? 0;
    message.delegator = object.delegator ?? "";
    message.validator = object.validator ?? "";
    return message;
  },
};

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
