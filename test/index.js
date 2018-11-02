#!/usr/bin/env node

const util = require('util')

const { UsersClient } = require("../gen/js/users/v3/users_pb_service")
const { GetRequest }  = require("../gen/js/users/v3/users_pb")

const client = new UsersClient("http://localhost:80")

const req = new GetRequest()
req.setName("users/foo")

client.get(req, (err, user) => {
    log(user)
})

function log(o) {
    console.log(util.inspect(o.toObject(), false, null, true))
}
