#!/bin/bash
set -ex -o pipefail

ENVOY_CLUSTER=${ENVOY_CLUSTER:-$1}
ENVOY_OPTS=${ENVOY_OPTS:-"-c /etc/envoy/envoy.yaml"}

# run CMD in background
"$@" &

# run envoy in foreground
envoy $ENVOY_OPTS --service-node $HOSTNAME --service-cluster $ENVOY_CLUSTER
