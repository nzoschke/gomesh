// package: gomesh.widgets.v1
// file: proto/widgets/v1/widgets.proto

var proto_widgets_v1_widgets_pb = require("../../../proto/widgets/v1/widgets_pb");
var grpc = require("grpc-web-client").grpc;

var Widgets = (function () {
  function Widgets() {}
  Widgets.serviceName = "gomesh.widgets.v1.Widgets";
  return Widgets;
}());

Widgets.List = {
  methodName: "List",
  service: Widgets,
  requestStream: false,
  responseStream: false,
  requestType: proto_widgets_v1_widgets_pb.ListRequest,
  responseType: proto_widgets_v1_widgets_pb.ListResponse
};

exports.Widgets = Widgets;

function WidgetsClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

WidgetsClient.prototype.list = function list(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Widgets.List, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

exports.WidgetsClient = WidgetsClient;

