// package: gomesh.widgets.v3
// file: widgets/v3/widgets.proto

import * as widgets_v3_widgets_pb from "../../widgets/v3/widgets_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "grpc-web-client";

type WidgetsGet = {
  readonly methodName: string;
  readonly service: typeof Widgets;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof widgets_v3_widgets_pb.GetRequest;
  readonly responseType: typeof widgets_v3_widgets_pb.Widget;
};

type WidgetsCreate = {
  readonly methodName: string;
  readonly service: typeof Widgets;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof widgets_v3_widgets_pb.CreateRequest;
  readonly responseType: typeof widgets_v3_widgets_pb.Widget;
};

type WidgetsUpdate = {
  readonly methodName: string;
  readonly service: typeof Widgets;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof widgets_v3_widgets_pb.UpdateRequest;
  readonly responseType: typeof widgets_v3_widgets_pb.Widget;
};

type WidgetsDelete = {
  readonly methodName: string;
  readonly service: typeof Widgets;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof widgets_v3_widgets_pb.DeleteRequest;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

type WidgetsList = {
  readonly methodName: string;
  readonly service: typeof Widgets;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof widgets_v3_widgets_pb.ListRequest;
  readonly responseType: typeof widgets_v3_widgets_pb.ListResponse;
};

type WidgetsBatchGet = {
  readonly methodName: string;
  readonly service: typeof Widgets;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof widgets_v3_widgets_pb.BatchGetRequest;
  readonly responseType: typeof widgets_v3_widgets_pb.BatchGetResponse;
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

export class WidgetsClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  get(
    requestMessage: widgets_v3_widgets_pb.GetRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: widgets_v3_widgets_pb.Widget|null) => void
  ): UnaryResponse;
  get(
    requestMessage: widgets_v3_widgets_pb.GetRequest,
    callback: (error: ServiceError|null, responseMessage: widgets_v3_widgets_pb.Widget|null) => void
  ): UnaryResponse;
  create(
    requestMessage: widgets_v3_widgets_pb.CreateRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: widgets_v3_widgets_pb.Widget|null) => void
  ): UnaryResponse;
  create(
    requestMessage: widgets_v3_widgets_pb.CreateRequest,
    callback: (error: ServiceError|null, responseMessage: widgets_v3_widgets_pb.Widget|null) => void
  ): UnaryResponse;
  update(
    requestMessage: widgets_v3_widgets_pb.UpdateRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: widgets_v3_widgets_pb.Widget|null) => void
  ): UnaryResponse;
  update(
    requestMessage: widgets_v3_widgets_pb.UpdateRequest,
    callback: (error: ServiceError|null, responseMessage: widgets_v3_widgets_pb.Widget|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: widgets_v3_widgets_pb.DeleteRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: widgets_v3_widgets_pb.DeleteRequest,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  list(
    requestMessage: widgets_v3_widgets_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: widgets_v3_widgets_pb.ListResponse|null) => void
  ): UnaryResponse;
  list(
    requestMessage: widgets_v3_widgets_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: widgets_v3_widgets_pb.ListResponse|null) => void
  ): UnaryResponse;
  batchGet(
    requestMessage: widgets_v3_widgets_pb.BatchGetRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: widgets_v3_widgets_pb.BatchGetResponse|null) => void
  ): UnaryResponse;
  batchGet(
    requestMessage: widgets_v3_widgets_pb.BatchGetRequest,
    callback: (error: ServiceError|null, responseMessage: widgets_v3_widgets_pb.BatchGetResponse|null) => void
  ): UnaryResponse;
}

