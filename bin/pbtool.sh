#!/bin/bash
set -e -o pipefail
PWD=$(pwd)

# Example output from `prototool compile --dry-run`:
#
# /Users/noah/Library/Caches/prototool/Darwin/x86_64/protobuf/3.6.1/bin/protoc \
#   -I /Users/noah/Library/Caches/prototool/Darwin/x86_64/protobuf/3.6.1/include \
#   -I /Users/noah/dev/gomesh \
#   -o /dev/null \
#   /Users/noah/dev/gomesh/proto/users/v1/users.proto

prototool compile --dry-run | while read x; do
    # get last arg
    IN=$(echo $x | grep -oE "[^ ]+$")      # /Users/noah/dev/gomesh/proto/users/v1/users.proto

    # get arg components
    IN_FILE=${IN/$PWD\//}                  # proto/users/v1/users.proto
    IN_DIR=$(dirname $IN_FILE)             # proto/users/v1

    # make dir and compile .pb
    OUT_DIR=gen/pb/${IN_DIR}               # gen/pb/proto/users/v1
    mkdir -p $OUT_DIR

    OUT_FILE=gen/pb/${IN_FILE/.proto/.pb}  # gen/pb/proto/users/v1/users.pb
    CMD=${x/\/dev\/null/$OUT_FILE}         # replace /dev/null with output
    $CMD --include_imports

    # make dir and generate mock
    OUT_DIR=gen/go/${IN_DIR}/mock          # gen/go/proto/users/v1/mock
    mkdir -p $OUT_DIR

    SOURCE=gen/go/${IN_FILE/.proto/.pb.go} # gen/go/proto/users/v1/users.pb.go
    OUT_FILE=$OUT_DIR/mock.go              # gen/go/proto/users/v1/mock/mock.go
    mockgen -source=$SOURCE > $OUT_FILE
done
