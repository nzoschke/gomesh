// package: gomesh.widgets.v1
// file: proto/widgets/v1/widgets.proto

import * as proto_widgets_v1_widgets_pb from "../../../proto/widgets/v1/widgets_pb";
import {grpc} from "grpc-web-client";

type WidgetsList = {
  readonly methodName: string;
  readonly service: typeof Widgets;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_widgets_v1_widgets_pb.ListRequest;
  readonly responseType: typeof proto_widgets_v1_widgets_pb.ListResponse;
};

export class Widgets {
  static readonly serviceName: string;
  static readonly List: WidgetsList;
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
  list(
    requestMessage: proto_widgets_v1_widgets_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_widgets_v1_widgets_pb.ListResponse|null) => void
  ): void;
  list(
    requestMessage: proto_widgets_v1_widgets_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: proto_widgets_v1_widgets_pb.ListResponse|null) => void
  ): void;
}

