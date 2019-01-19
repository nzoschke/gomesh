#!/bin/bash
set -x

yamllint \
  -d "{extends: default, rules: {key-ordering: {}, line-length: {max: 140}}}" \
  .