// package: gomesh.users.v3
// file: users/v3/users.proto

var users_v3_users_pb = require("../../users/v3/users_pb");
var grpc = require("grpc-web-client").grpc;

var Users = (function () {
  function Users() {}
  Users.serviceName = "gomesh.users.v3.Users";
  return Users;
}());

Users.Get = {
  methodName: "Get",
  service: Users,
  requestStream: false,
  responseStream: false,
  requestType: users_v3_users_pb.GetRequest,
  responseType: users_v3_users_pb.User
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
  var client = grpc.unary(Users.Get, {
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

exports.UsersClient = UsersClient;

