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
    # get last arg, e.g. /Users/noah/dev/gomesh/proto/users/v1/users.proto
    IN=$(echo $x | grep -oE "[^ ]+$")

    # translate to output, e.g. /Users/noah/dev/gomesh/gen/pb/proto/users/v1/users.pb
    OUT=${IN/$PWD/$PWD/gen/pb}
    OUT=${OUT/.proto/.pb}

    # replace /dev/null with output
    CMD=${x/\/dev\/null/$OUT}

    # make dir and compile
    mkdir -p $(dirname $OUT)
    $CMD --include_imports
done
