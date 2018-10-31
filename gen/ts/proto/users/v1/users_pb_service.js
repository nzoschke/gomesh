// package: gomesh.users.v1
// file: proto/users/v1/users.proto

var proto_users_v1_users_pb = require("../../../proto/users/v1/users_pb");
var grpc = require("grpc-web-client").grpc;

var Users = (function () {
  function Users() {}
  Users.serviceName = "gomesh.users.v1.Users";
  return Users;
}());

Users.Get = {
  methodName: "Get",
  service: Users,
  requestStream: false,
  responseStream: false,
  requestType: proto_users_v1_users_pb.GetRequest,
  responseType: proto_users_v1_users_pb.User
};

Users.Create = {
  methodName: "Create",
  service: Users,
  requestStream: false,
  responseStream: false,
  requestType: proto_users_v1_users_pb.CreateRequest,
  responseType: proto_users_v1_users_pb.User
};

exports.Users = Users;

function UsersClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

UsersClient.prototype.get = function get(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Users.Get, {
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

UsersClient.prototype.create = function create(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Users.Create, {
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

exports.UsersClient = UsersClient;

