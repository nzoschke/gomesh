// package: gomesh.widgets.v3
// file: widgets/v3/widgets.proto

var widgets_v3_widgets_pb = require("../../widgets/v3/widgets_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("grpc-web-client").grpc;

var Widgets = (function () {
  function Widgets() {}
  Widgets.serviceName = "gomesh.widgets.v3.Widgets";
  return Widgets;
}());

Widgets.Get = {
  methodName: "Get",
  service: Widgets,
  requestStream: false,
  responseStream: false,
  requestType: widgets_v3_widgets_pb.GetRequest,
  responseType: widgets_v3_widgets_pb.Widget
};

Widgets.Create = {
  methodName: "Create",
  service: Widgets,
  requestStream: false,
  responseStream: false,
  requestType: widgets_v3_widgets_pb.CreateRequest,
  responseType: widgets_v3_widgets_pb.Widget
};

Widgets.Update = {
  methodName: "Update",
  service: Widgets,
  requestStream: false,
  responseStream: false,
  requestType: widgets_v3_widgets_pb.UpdateRequest,
  responseType: widgets_v3_widgets_pb.Widget
};

Widgets.Delete = {
  methodName: "Delete",
  service: Widgets,
  requestStream: false,
  responseStream: false,
  requestType: widgets_v3_widgets_pb.DeleteRequest,
  responseType: google_protobuf_empty_pb.Empty
};

Widgets.List = {
  methodName: "List",
  service: Widgets,
  requestStream: false,
  responseStream: false,
  requestType: widgets_v3_widgets_pb.ListRequest,
  responseType: widgets_v3_widgets_pb.ListResponse
};

Widgets.BatchGet = {
  methodName: "BatchGet",
  service: Widgets,
  requestStream: false,
  responseStream: false,
  requestType: widgets_v3_widgets_pb.BatchGetRequest,
  responseType: widgets_v3_widgets_pb.BatchGetResponse
};

exports.Widgets = Widgets;

function WidgetsClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

WidgetsClient.prototype.get = function get(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Widgets.Get, {
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
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

WidgetsClient.prototype.create = function create(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Widgets.Create, {
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
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

WidgetsClient.prototype.update = function update(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Widgets.Update, {
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
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

WidgetsClient.prototype.delete = function pb_delete(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Widgets.Delete, {
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
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

WidgetsClient.prototype.list = function list(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Widgets.List, {
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
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

WidgetsClient.prototype.batchGet = function batchGet(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Widgets.BatchGet, {
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
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.WidgetsClient = WidgetsClient;

