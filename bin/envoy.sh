#!/bin/bash
set -ex -o pipefail

ENVOY_CLUSTER=${ENVOY_CLUSTER:-$1}
ENVOY_CONFIG=${ENVOY_CONFIG:-/etc/envoy/envoy.yaml}

# run CMD in background
"$@" &

# run envoy in foreground
envoy $ENVOY_OPTS -c $ENVOY_CONFIG --service-node $HOSTNAME --service-cluster $ENVOY_CLUSTER
