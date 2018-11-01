// package: gomesh.widgets.v2
// file: proto/widgets/v2/widgets.proto

import * as proto_widgets_v2_widgets_pb from "../../../proto/widgets/v2/widgets_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "grpc-web-client";

type WidgetsGet = {
  readonly methodName: string;
  readonly service: typeof Widgets;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_widgets_v2_widgets_pb.GetRequest;
  readonly responseType: typeof proto_widgets_v2_widgets_pb.Widget;
};

type WidgetsCreate = {
  readonly methodName: string;
  readonly service: typeof Widgets;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_widgets_v2_widgets_pb.CreateRequest;
  readonly responseType: typeof proto_widgets_v2_widgets_pb.Widget;
};

type WidgetsUpdate = {
  readonly methodName: string;
  readonly service: typeof Widgets;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_widgets_v2_widgets_pb.UpdateRequest;
  readonly responseType: typeof proto_widgets_v2_widgets_pb.Widget;
};

type WidgetsDelete = {
  readonly methodName: string;
  readonly service: typeof Widgets;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_widgets_v2_widgets_pb.DeleteRequest;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

type WidgetsList = {
  readonly methodName: string;
  readonly service: typeof Widgets;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_widgets_v2_widgets_pb.ListRequest;
  readonly responseType: typeof proto_widgets_v2_widgets_pb.ListResponse;
};

type WidgetsBatchGet = {
  readonly methodName: string;
  readonly service: typeof Widgets;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_widgets_v2_widgets_pb.BatchGetRequest;
  readonly responseType: typeof proto_widgets_v2_widgets_pb.BatchGetResponse;
};

export class Widgets {
  static readonly serviceName: string;
  static readonly Get: WidgetsGet;
  static readonly Create: WidgetsCreate;
  static readonly Update: WidgetsUpdate;
  static readonly Delete: WidgetsDelete;
  static readonly List: WidgetsList;
  static readonly BatchGet: WidgetsBatchGet;
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

export class WidgetsClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  get(
    requestMessage: proto_widgets_v2_widgets_pb.GetRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_widgets_v2_widgets_pb.Widget|null) => void
  ): void;
  get(
    requestMessage: proto_widgets_v2_widgets_pb.GetRequest,
    callback: (error: ServiceError|null, responseMessage: proto_widgets_v2_widgets_pb.Widget|null) => void
  ): void;
  create(
    requestMessage: proto_widgets_v2_widgets_pb.CreateRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_widgets_v2_widgets_pb.Widget|null) => void
  ): void;
  create(
    requestMessage: proto_widgets_v2_widgets_pb.CreateRequest,
    callback: (error: ServiceError|null, responseMessage: proto_widgets_v2_widgets_pb.Widget|null) => void
  ): void;
  update(
    requestMessage: proto_widgets_v2_widgets_pb.UpdateRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_widgets_v2_widgets_pb.Widget|null) => void
  ): void;
  update(
    requestMessage: proto_widgets_v2_widgets_pb.UpdateRequest,
    callback: (error: ServiceError|null, responseMessage: proto_widgets_v2_widgets_pb.Widget|null) => void
  ): void;
  delete(
    requestMessage: proto_widgets_v2_widgets_pb.DeleteRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): void;
  delete(
    requestMessage: proto_widgets_v2_widgets_pb.DeleteRequest,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): void;
  list(
    requestMessage: proto_widgets_v2_widgets_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_widgets_v2_widgets_pb.ListResponse|null) => void
  ): void;
  list(
    requestMessage: proto_widgets_v2_widgets_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: proto_widgets_v2_widgets_pb.ListResponse|null) => void
  ): void;
  batchGet(
    requestMessage: proto_widgets_v2_widgets_pb.BatchGetRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_widgets_v2_widgets_pb.BatchGetResponse|null) => void
  ): void;
  batchGet(
    requestMessage: proto_widgets_v2_widgets_pb.BatchGetRequest,
    callback: (error: ServiceError|null, responseMessage: proto_widgets_v2_widgets_pb.BatchGetResponse|null) => void
  ): void;
}

