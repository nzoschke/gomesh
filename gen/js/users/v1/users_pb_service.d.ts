// package: gomesh.users.v1
// file: users/v1/users.proto

import * as users_v1_users_pb from "../../users/v1/users_pb";
import {grpc} from "grpc-web-client";

type UsersGet = {
  readonly methodName: string;
  readonly service: typeof Users;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof users_v1_users_pb.GetRequest;
  readonly responseType: typeof users_v1_users_pb.User;
};

type UsersCreate = {
  readonly methodName: string;
  readonly service: typeof Users;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof users_v1_users_pb.CreateRequest;
  readonly responseType: typeof users_v1_users_pb.User;
};

export class Users {
  static readonly serviceName: string;
  static readonly Get: UsersGet;
  static readonly Create: UsersCreate;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: () => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: () => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class UsersClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  get(
    requestMessage: users_v1_users_pb.GetRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: users_v1_users_pb.User|null) => void
  ): UnaryResponse;
  get(
    requestMessage: users_v1_users_pb.GetRequest,
    callback: (error: ServiceError|null, responseMessage: users_v1_users_pb.User|null) => void
  ): UnaryResponse;
  create(
    requestMessage: users_v1_users_pb.CreateRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: users_v1_users_pb.User|null) => void
  ): UnaryResponse;
  create(
    requestMessage: users_v1_users_pb.CreateRequest,
    callback: (error: ServiceError|null, responseMessage: users_v1_users_pb.User|null) => void
  ): UnaryResponse;
}

