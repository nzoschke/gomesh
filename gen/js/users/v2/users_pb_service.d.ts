// package: gomesh.users.v2
// file: users/v2/users.proto

import * as users_v2_users_pb from "../../users/v2/users_pb";
import {grpc} from "grpc-web-client";

type UsersGet = {
  readonly methodName: string;
  readonly service: typeof Users;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof users_v2_users_pb.GetRequest;
  readonly responseType: typeof users_v2_users_pb.User;
};

type UsersCreate = {
  readonly methodName: string;
  readonly service: typeof Users;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof users_v2_users_pb.CreateRequest;
  readonly responseType: typeof users_v2_users_pb.User;
};

export class Users {
  static readonly serviceName: string;
  static readonly Get: UsersGet;
  static readonly Create: UsersCreate;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }
export type ServiceClientOptions = { transport: grpc.TransportConstructor; debug?: boolean }

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
interface BidirectionalStream<T> {
  write(message: T): BidirectionalStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): BidirectionalStream<T>;
  on(type: 'end', handler: () => void): BidirectionalStream<T>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<T>;
}

export class UsersClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  get(
    requestMessage: users_v2_users_pb.GetRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.User|null) => void
  ): void;
  get(
    requestMessage: users_v2_users_pb.GetRequest,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.User|null) => void
  ): void;
  create(
    requestMessage: users_v2_users_pb.CreateRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.User|null) => void
  ): void;
  create(
    requestMessage: users_v2_users_pb.CreateRequest,
    callback: (error: ServiceError|null, responseMessage: users_v2_users_pb.User|null) => void
  ): void;
}

