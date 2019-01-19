#!/bin/bash
set -ex

CONFIG=${1:-/etc/envoy/envoy.yaml}
SERVICE=${2:-$HOSTNAME}

# template config file
sed                                \
    -e "s/\${HOSTNAME}/$HOSTNAME/" \
    -e "s/\${SERVICE}/$SERVICE/"   \
    $CONFIG > /tmp/envoy.yaml

exec envoy -c /tmp/envoy.yaml