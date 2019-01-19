#!/bin/bash
set -x
STATUS=0

# lint protos
prototool lint      proto     || STATUS=$?
prototool format -l proto     || STATUS=$?

# generate protos and mocks
prototool generate  proto     || STATUS=$?
prototool generate  proto_ext || STATUS=$?
find gen -name 'mock_*.go' -delete
mockery -all -dir gen -inpkg  || STATUS=$?

# exit on any errors
[ $STATUS -eq 0 ] || exit $STATUS

# push gen

git clone https://nzoschke:${PUSH_TOKEN}@github.com/nzoschke/gomesh-interface.git && cd gomesh-interface

# sync gen, proto, proto_ext folders
git rm -r gen proto proto_ext
cp -r ../gen ../proto ../proto_ext .
git add -f .

[[ -z $(git status -uno --porcelain) ]] && echo "this branch is clean, no need to push..." && exit 0;

git config user.email "gen@example.com"
git config user.name  "gen"
git commit -m "gen ${GITHUB_SHA:0:7}"
git push -f origin HEAD:${GITHUB_REF}