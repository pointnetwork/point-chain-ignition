/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "point.lockenomics";

export interface MsgCreateLock {
  creator: string;
  lenght: number;
  validator: string;
}

export interface MsgCreateLockResponse {
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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateLock(request: MsgCreateLock): Promise<MsgCreateLockResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateLock = this.CreateLock.bind(this);
  }
  CreateLock(request: MsgCreateLock): Promise<MsgCreateLockResponse> {
    const data = MsgCreateLock.encode(request).finish();
    const promise = this.rpc.request("point.lockenomics.Msg", "CreateLock", data);
    return promise.then((data) => MsgCreateLockResponse.decode(new _m0.Reader(data)));
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
